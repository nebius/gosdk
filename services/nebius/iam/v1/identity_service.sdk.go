package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[IdentityServiceID] = "identity.iam"
}

func (s Services) Identity() IdentityService {
	return NewIdentityService(s.sdk)
}

const IdentityServiceID conn.ServiceID = "nebius.iam.v1.IdentityService"

type IdentityService interface {
	ExchangeToken(context.Context, *v1.ExchangeTokenRequest, ...grpc.CallOption) (*v1.CreateTokenResponse, error)
}

type identityService struct {
	sdk iface.SDK
}

func NewIdentityService(sdk iface.SDK) IdentityService {
	return identityService{
		sdk: sdk,
	}
}

func (s identityService) ExchangeToken(ctx context.Context, request *v1.ExchangeTokenRequest, opts ...grpc.CallOption) (*v1.CreateTokenResponse, error) {
	address, err := s.sdk.Resolve(ctx, IdentityServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewIdentityServiceClient(con).ExchangeToken(ctx, request, opts...)
}
