package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[TokenExchangeServiceID] = "tokens.iam"
}

func (s Services) TokenExchange() TokenExchangeService {
	return NewTokenExchangeService(s.sdk)
}

const TokenExchangeServiceID conn.ServiceID = "nebius.iam.v1.TokenExchangeService"

type TokenExchangeService interface {
	Exchange(context.Context, *v1.ExchangeTokenRequest, ...grpc.CallOption) (*v1.CreateTokenResponse, error)
}

type tokenExchangeService struct {
	sdk iface.SDK
}

func NewTokenExchangeService(sdk iface.SDK) TokenExchangeService {
	return tokenExchangeService{
		sdk: sdk,
	}
}

func (s tokenExchangeService) Exchange(ctx context.Context, request *v1.ExchangeTokenRequest, opts ...grpc.CallOption) (*v1.CreateTokenResponse, error) {
	address, err := s.sdk.Resolve(ctx, TokenExchangeServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTokenExchangeServiceClient(con).Exchange(ctx, request, opts...)
}
