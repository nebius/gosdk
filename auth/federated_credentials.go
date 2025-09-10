package auth

import (
	"context"
	"fmt"
	"os"

	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

type FederatedCredentials string

type FederatedCredentialsReader interface {
	GetFederatedCredentials() (FederatedCredentials, error)
}

type FederatedCredentialsTokenRequester struct {
	serviceAccountID string
	reader           FederatedCredentialsReader
}

func NewFederatedCredentialsTokenRequester(serviceAccountID string, reader FederatedCredentialsReader) *FederatedCredentialsTokenRequester {
	return &FederatedCredentialsTokenRequester{
		serviceAccountID: serviceAccountID,
		reader:           reader,
	}
}

// GetExchangeTokenRequest implements ExchangeTokenRequester.
func (f *FederatedCredentialsTokenRequester) GetExchangeTokenRequest(ctx context.Context) (*iampb.ExchangeTokenRequest, error) {
	federatedCredentials, err := f.reader.GetFederatedCredentials()
	if err != nil {
		return nil, fmt.Errorf("get federated credentials: %w", err)
	}

	return &iampb.ExchangeTokenRequest{
		GrantType:          "urn:ietf:params:oauth:grant-type:token-exchange",
		RequestedTokenType: "urn:ietf:params:oauth:token-type:access_token",
		SubjectToken:       f.serviceAccountID,
		SubjectTokenType:   "urn:nebius:params:oauth:token-type:subject_identifier",
		ActorTokenType:     "urn:ietf:params:oauth:token-type:jwt",
		ActorToken:         string(federatedCredentials),
	}, nil
}

var _ ExchangeTokenRequester = (*FederatedCredentialsTokenRequester)(nil)

type StaticFederatedCredentialsReader struct {
	credentials FederatedCredentials
}

func NewStaticFederatedCredentialsReader(credentials FederatedCredentials) *StaticFederatedCredentialsReader {
	return &StaticFederatedCredentialsReader{
		credentials: credentials,
	}
}

func (s *StaticFederatedCredentialsReader) GetFederatedCredentials() (FederatedCredentials, error) {
	return s.credentials, nil
}

type FileFederatedCredentialsReader struct {
	filePath string
}

var _ FederatedCredentialsReader = (*FileFederatedCredentialsReader)(nil)

func NewFileFederatedCredentialsReader(filePath string) *FileFederatedCredentialsReader {
	return &FileFederatedCredentialsReader{
		filePath: filePath,
	}
}

func (f *FileFederatedCredentialsReader) GetFederatedCredentials() (FederatedCredentials, error) {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}
	return FederatedCredentials(data), nil
}

// NewFederatedCredentialsTokener creates a [BearerTokener] serving token from
// federated credentials stored in a file.
// Don't forget to wrap in NewAsynchronouslyRenewableFileCachedTokener or
// NewCachedServiceTokener
func NewFederatedCredentialsTokener(
	serviceAccountID string,
	fileName string,
	clientFunc func() (iampb.TokenExchangeServiceClient, error),
) NamedTokener {
	reader := NewFileFederatedCredentialsReader(fileName)
	requester := NewFederatedCredentialsTokenRequester(serviceAccountID, reader)
	return NewNameWrapper(
		"federated-credentials/"+serviceAccountID+"/"+fileName,
		NewExchangeableBearerTokenerWithDeferredClient(
			requester,
			clientFunc,
		),
	)
}
