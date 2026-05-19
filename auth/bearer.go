package auth

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/metadata"
)

const AuthorizationHeader = "Authorization"

// BearerTokener is an interface for providing a [BearerToken] for authentication.
// Most implementations provided in this package are decorators,
// allowing you to layer additional behavior on top of a base implementation.
// These can be combined or extended to fit custom authentication requirements.
type BearerTokener interface {
	// BearerToken returns a [BearerToken] for use in the [AuthorizationHeader].
	BearerToken(context.Context) (BearerToken, error)

	// HandleError processes errors from gRPC calls that fail with the Unauthenticated status code.
	// It receives the [BearerToken] returned by [BearerTokener.BearerToken] and the error from the gRPC call.
	// If HandleError returns nil, the system will attempt to re-authenticate and retry the call.
	// If the call should not be retried, HandleError should return the original error.
	HandleError(context.Context, BearerToken, error) error
}

// StaticBearerToken is a [BearerTokener] that always returns a fixed [BearerToken].
type StaticBearerToken string

var _ BearerTokener = StaticBearerToken("")
var _ TypedTokener = StaticBearerToken("")

// NewStaticBearerToken returns a [BearerTokener] that always returns a fixed [BearerToken].
func NewStaticBearerToken(token string) StaticBearerToken {
	return StaticBearerToken(token)
}

func (t StaticBearerToken) BearerToken(context.Context) (BearerToken, error) {
	return BearerToken{
		Token:     string(t),
		ExpiresAt: time.Time{},
	}, nil
}

func (t StaticBearerToken) HandleError(_ context.Context, _ BearerToken, err error) error {
	return err // Unauthenticated error should not be retried
}

func (t StaticBearerToken) Type() string {
	return "static"
}

// InstrumentedBearerTokener wraps a [BearerTokener] and emits acquire/lifetime
// metrics around calls to the wrapped provider.
//
// If the wrapped tokener already emits metrics and receives the same metrics
// through SetMetrics, both layers will report observations. Use this wrapper
// when that extra outer observation is desired.
type InstrumentedBearerTokener struct {
	tokener BearerTokener
	metrics atomicMetrics
}

var _ BearerTokener = (*InstrumentedBearerTokener)(nil)
var _ MetricsSetter = (*InstrumentedBearerTokener)(nil)
var _ Wrapper = (*InstrumentedBearerTokener)(nil)

func NewInstrumentedBearerTokener(tokener BearerTokener) *InstrumentedBearerTokener {
	return &InstrumentedBearerTokener{tokener: tokener}
}

// NewStaticTokener returns a [BearerTokener] that always returns a fixed
// [BearerToken] and can participate in auth metrics.
func NewStaticTokener(token string, opts ...Option) *InstrumentedBearerTokener {
	tokener := NewInstrumentedBearerTokener(StaticBearerToken(token))
	applyOptions(tokener, opts...)
	return tokener
}

func (t *InstrumentedBearerTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	if t == nil || t.tokener == nil {
		return BearerToken{}, errors.New("instrumented bearer tokener has no wrapped tokener")
	}
	start := time.Now()
	token, err := t.tokener.BearerToken(ctx)
	if err != nil {
		t.metrics.tokenAcquireError(ctx, t, time.Since(start), 0)
		return BearerToken{}, err
	}
	t.metrics.tokenAcquireSuccess(ctx, t, time.Since(start), 0, token, time.Now())
	return token, nil
}

func (t *InstrumentedBearerTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	if t == nil || t.tokener == nil {
		if err != nil {
			return err
		}
		return errors.New("instrumented bearer tokener has no wrapped tokener")
	}
	return t.tokener.HandleError(ctx, token, err)
}

func (t *InstrumentedBearerTokener) SetMetrics(metrics Metrics) {
	if t == nil {
		return
	}
	t.metrics.Store(metrics)
	if setter, ok := t.tokener.(MetricsSetter); ok {
		setter.SetMetrics(metrics)
	}
}

func (t *InstrumentedBearerTokener) Unwrap() BearerTokener {
	if t == nil {
		return nil
	}
	return t.tokener
}

// AuthenticatorFromBearerTokener is an [Authenticator] that uses a [BearerTokener]
// to populate the [AuthorizationHeader] with a "Bearer " prefix.
type AuthenticatorFromBearerTokener struct {
	tokener BearerTokener
}

var _ Authenticator = AuthenticatorFromBearerTokener{}

// NewAuthenticatorFromBearerTokener returns an [Authenticator] that uses a [BearerTokener]
// to populate the [AuthorizationHeader] with a "Bearer " prefix.
func NewAuthenticatorFromBearerTokener(tokener BearerTokener) AuthenticatorFromBearerTokener {
	return AuthenticatorFromBearerTokener{
		tokener: tokener,
	}
}

func (a AuthenticatorFromBearerTokener) Auth(ctx context.Context) (context.Context, error) {
	token, err := a.tokener.BearerToken(ctx)
	if err != nil {
		return nil, err
	}
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = make(metadata.MD)
	}
	md.Set(AuthorizationHeader, "Bearer "+token.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)
	ctx = context.WithValue(ctx, ctxKeyBearerToken{}, token)
	return ctx, nil
}

func (a AuthenticatorFromBearerTokener) HandleError(ctx context.Context, err error) error {
	token, ok := ctx.Value(ctxKeyBearerToken{}).(BearerToken)
	if !ok {
		return err // Unauthenticated error should not be retried
	}
	return a.tokener.HandleError(ctx, token, err)
}

type ctxKeyBearerToken struct{}

// PropagateAuthorizationHeader is an [Authenticator] that transfers the [AuthorizationHeader]
// from incoming gRPC metadata to outgoing metadata.
type PropagateAuthorizationHeader struct{}

var _ Authenticator = PropagateAuthorizationHeader{}

// NewPropagateAuthorizationHeader returns an [Authenticator] that transfers the [AuthorizationHeader]
// from incoming gRPC metadata to outgoing metadata.
func NewPropagateAuthorizationHeader() PropagateAuthorizationHeader {
	return PropagateAuthorizationHeader{}
}

func (PropagateAuthorizationHeader) Auth(ctx context.Context) (context.Context, error) {
	incoming := metadata.ValueFromIncomingContext(ctx, AuthorizationHeader)
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = make(metadata.MD)
	}
	md.Set(AuthorizationHeader, incoming...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, nil
}

func (PropagateAuthorizationHeader) HandleError(_ context.Context, err error) error {
	return err // Unauthenticated error should not be retried
}
