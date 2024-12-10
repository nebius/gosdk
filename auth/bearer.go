package auth

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"
)

const AuthorizationHeader = "Authorization"

// BearerToken is a token used in the Bearer authentication scheme.
type BearerToken struct {
	Token     string
	ExpiresAt time.Time
}

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
