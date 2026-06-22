package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

const (
	tokenExchangeAccessTokenType       = "urn:ietf:params:oauth:token-type:access_token"         //nolint:gosec // OAuth token type URN, not a credential.
	tokenExchangeGrantType             = "urn:ietf:params:oauth:grant-type:token-exchange"       //nolint:gosec // OAuth grant type URN, not a credential.
	tokenExchangeSubjectIdentifierType = "urn:nebius:params:oauth:token-type:subject_identifier" //nolint:gosec // OAuth token type URN, not a credential.
)

// ExchangeImpersonatedBearerTokener creates a service account token by calling
// an IAM token exchange RPC.
type ExchangeImpersonatedBearerTokener struct {
	metrics          atomicMetrics
	serviceAccountID string
	tokener          BearerTokener
	client           iampb.TokenExchangeServiceClient
	clientFunc       func() (iampb.TokenExchangeServiceClient, error)
	now              func() time.Time
}

var _ BearerTokener = (*ExchangeImpersonatedBearerTokener)(nil)
var _ MetricsSetter = (*ExchangeImpersonatedBearerTokener)(nil)
var _ TypedTokener = (*ExchangeImpersonatedBearerTokener)(nil)

func NewExchangeImpersonatedBearerTokener(
	serviceAccountID string,
	tokener BearerTokener,
	clientFunc func() (iampb.TokenExchangeServiceClient, error),
) *ExchangeImpersonatedBearerTokener {
	return &ExchangeImpersonatedBearerTokener{
		serviceAccountID: serviceAccountID,
		tokener:          tokener,
		clientFunc:       clientFunc,
		now:              time.Now,
	}
}

func (t *ExchangeImpersonatedBearerTokener) SetMetrics(metrics Metrics) {
	t.metrics.Store(metrics)
	if setter, ok := t.tokener.(MetricsSetter); ok {
		setter.SetMetrics(metrics)
	}
}

func (t *ExchangeImpersonatedBearerTokener) Unwrap() BearerTokener {
	return t.tokener
}

func (t *ExchangeImpersonatedBearerTokener) Type() string {
	return "impersonated"
}

func (t *ExchangeImpersonatedBearerTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	start := t.now()
	if t.client == nil && t.clientFunc != nil {
		client, err := t.clientFunc()
		if err != nil {
			t.metrics.tokenAcquireError(ctx, t, t.now().Sub(start), 0)
			return BearerToken{}, fmt.Errorf("get token exchange client: %w", err)
		}
		t.client = client
	}
	if t.client == nil {
		t.metrics.tokenAcquireError(ctx, t, t.now().Sub(start), 0)
		return BearerToken{}, errors.New("token exchange client is not set")
	}
	token, err := t.tokener.BearerToken(ctx)
	if err != nil {
		t.metrics.tokenAcquireError(ctx, t, t.now().Sub(start), 0)
		return BearerToken{}, err
	}

	now := t.now()
	response, err := t.client.Exchange(ctx, t.request(token.Token), Disable{})

	if shouldRetryExchangeActorToken(err) {
		errX := t.tokener.HandleError(ctx, token, err)
		if errX != nil {
			t.metrics.tokenAcquireError(ctx, t, t.now().Sub(start), 0)
			return BearerToken{}, errX
		}

		newToken, errX := t.tokener.BearerToken(ctx)
		if errX != nil {
			t.metrics.tokenAcquireError(ctx, t, t.now().Sub(start), 0)
			return BearerToken{}, errX
		}

		now = t.now()
		response, err = t.client.Exchange(ctx, t.request(newToken.Token), Disable{})
	}

	if err != nil {
		t.metrics.tokenAcquireError(ctx, t, t.now().Sub(start), 0)
		return BearerToken{}, fmt.Errorf("exchange impersonated token: %w", err)
	}
	bearerToken := BearerToken{
		Token:     response.GetAccessToken(),
		ExpiresAt: now.Add(time.Duration(response.GetExpiresIn()) * time.Second),
	}
	t.metrics.tokenAcquireSuccess(ctx, t, t.now().Sub(start), 0, bearerToken, now)
	return bearerToken, nil
}

func shouldRetryExchangeActorToken(err error) bool {
	return status.Code(err) == codes.Unauthenticated
}

func (t *ExchangeImpersonatedBearerTokener) request(actorToken string) *iampb.ExchangeTokenRequest {
	return &iampb.ExchangeTokenRequest{
		GrantType:          tokenExchangeGrantType,
		RequestedTokenType: tokenExchangeAccessTokenType,
		SubjectToken:       t.serviceAccountID,
		SubjectTokenType:   tokenExchangeSubjectIdentifierType,
		ActorToken:         actorToken,
		ActorTokenType:     tokenExchangeAccessTokenType,
	}
}

func (t *ExchangeImpersonatedBearerTokener) HandleError(context.Context, BearerToken, error) error {
	return nil
}
