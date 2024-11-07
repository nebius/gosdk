package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

type ExchangeTokenRequester interface {
	GetExchangeTokenRequest(context.Context) (*iampb.ExchangeTokenRequest, error)
}

// ExchangeableBearerTokener sends request to [iampb.TokenExchangeServiceClient.Exchange] to get a [BearerToken].
type ExchangeableBearerTokener struct {
	creds  ExchangeTokenRequester
	client iampb.TokenExchangeServiceClient
	now    func() time.Time
}

var _ BearerTokener = (*ExchangeableBearerTokener)(nil)

func NewExchangeableBearerTokener(
	creds ExchangeTokenRequester,
	client iampb.TokenExchangeServiceClient,
) *ExchangeableBearerTokener {
	return &ExchangeableBearerTokener{
		creds:  creds,
		client: client,
		now:    time.Now,
	}
}

func (t *ExchangeableBearerTokener) SetClient(client iampb.TokenExchangeServiceClient) {
	t.client = client
}

func (t *ExchangeableBearerTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	if t.client == nil {
		return BearerToken{}, errors.New("tokenService client is not defined")
	}

	now := t.now()

	request, err := t.creds.GetExchangeTokenRequest(ctx)
	if err != nil {
		return BearerToken{}, err
	}

	resp, err := t.client.Exchange(ctx, request, Disable{})
	if err != nil {
		return BearerToken{}, fmt.Errorf("exchange token: %w", err)
	}

	if strings.ToLower(resp.GetTokenType()) != "bearer" {
		return BearerToken{}, fmt.Errorf("unsupported token received: expected Bearer, received %q", resp.GetTokenType())
	}

	return BearerToken{
		Token:     resp.GetAccessToken(),
		ExpiresAt: now.Add(time.Duration(resp.GetExpiresIn()) * time.Second),
	}, nil
}

func (t *ExchangeableBearerTokener) HandleError(context.Context, BearerToken, error) error {
	return nil
}

// ServiceAccountExchangeTokenRequester generates a signed JWT using the credentials of a [ServiceAccount],
// intended for exchange to obtain a [BearerToken].
type ServiceAccountExchangeTokenRequester struct {
	account ServiceAccountReader
	now     func() time.Time
}

var _ ExchangeTokenRequester = ServiceAccountExchangeTokenRequester{}

func NewServiceAccountExchangeTokenRequester(account ServiceAccountReader) ServiceAccountExchangeTokenRequester {
	return ServiceAccountExchangeTokenRequester{
		account: account,
		now:     time.Now,
	}
}

func (s ServiceAccountExchangeTokenRequester) GetExchangeTokenRequest(
	ctx context.Context,
) (*iampb.ExchangeTokenRequest, error) {
	serviceAccount, err := s.account.ServiceAccount(ctx)
	if err != nil {
		return nil, err
	}

	now := s.now()
	const jwtTTL = 1 * time.Minute // it should be alive only to exchange to the IAM token

	claims := jwt.RegisteredClaims{
		Issuer:    serviceAccount.ServiceAccountID,
		Subject:   serviceAccount.ServiceAccountID,
		Audience:  jwt.ClaimStrings{"token-service.iam.new.nebiuscloud.net"},
		ExpiresAt: jwt.NewNumericDate(now.Add(jwtTTL)),
		IssuedAt:  jwt.NewNumericDate(now),
	}
	jwToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwToken.Header["kid"] = serviceAccount.PublicKeyID

	signedJWT, err := jwToken.SignedString(serviceAccount.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("sign JWT: %w", err)
	}

	return &iampb.ExchangeTokenRequest{
		GrantType:          "urn:ietf:params:oauth:grant-type:token-exchange",
		RequestedTokenType: "urn:ietf:params:oauth:token-type:access_token",
		SubjectToken:       signedJWT,
		SubjectTokenType:   "urn:ietf:params:oauth:token-type:jwt",
	}, nil
}
