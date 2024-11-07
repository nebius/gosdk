package auth

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/metadata"
)

const AuthorizationHeader = "Authorization"

type BearerToken struct {
	Token     string
	ExpiresAt time.Time
}

type BearerTokener interface {
	// BearerToken returns a token to be used in "Authorization" header.
	BearerToken(context.Context) (BearerToken, error)

	// HandleError is called with the [BearerToken] received from [BearerTokener.BearerToken]
	// and an error from a gRPC call if it has the Unauthenticated code.
	// If HandleError returns nil, a new auth will be requested to retry the gRPC call.
	// If the gRPC call should not be retried, HandleError must return the incoming error.
	HandleError(context.Context, BearerToken, error) error
}

// StaticBearerToken implement [BearerTokener] with constant token.
type StaticBearerToken string

var _ BearerTokener = StaticBearerToken("")

func (t StaticBearerToken) BearerToken(context.Context) (BearerToken, error) {
	return BearerToken{
		Token:     string(t),
		ExpiresAt: time.Time{},
	}, nil
}

func (t StaticBearerToken) HandleError(_ context.Context, _ BearerToken, err error) error {
	return err
}

type AuthenticatorFromBearerTokener struct {
	tokener BearerTokener
}

var _ Authenticator = AuthenticatorFromBearerTokener{}

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
		return err
	}
	return a.tokener.HandleError(ctx, token, err)
}

type ctxKeyBearerToken struct{}

type PropagateAuthorizationHeader struct{}

var _ Authenticator = PropagateAuthorizationHeader{}

func NewPropagateAuthorizationHeader() PropagateAuthorizationHeader {
	return PropagateAuthorizationHeader{}
}

func (PropagateAuthorizationHeader) Auth(ctx context.Context) (context.Context, error) {
	incoming := metadata.ValueFromIncomingContext(ctx, AuthorizationHeader)
	if len(incoming) == 0 {
		return nil, errors.New("missing authorization header in the incoming metadata")
	}
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = make(metadata.MD)
	}
	md.Set(AuthorizationHeader, incoming...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, nil
}

func (PropagateAuthorizationHeader) HandleError(_ context.Context, err error) error {
	return err
}
