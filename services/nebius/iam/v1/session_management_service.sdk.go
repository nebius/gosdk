package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[SessionManagementServiceID] = "cpl.iam"
}

func (s Services) SessionManagement() SessionManagementService {
	return NewSessionManagementService(s.sdk)
}

const SessionManagementServiceID conn.ServiceID = "nebius.iam.v1.SessionManagementService"

type SessionManagementService interface {
	Revoke(context.Context, *v1.RevokeSessionRequest, ...grpc.CallOption) (*v1.RevokeSessionResponse, error)
}

type sessionManagementService struct {
	sdk iface.SDK
}

func NewSessionManagementService(sdk iface.SDK) SessionManagementService {
	return sessionManagementService{
		sdk: sdk,
	}
}

func (s sessionManagementService) Revoke(ctx context.Context, request *v1.RevokeSessionRequest, opts ...grpc.CallOption) (*v1.RevokeSessionResponse, error) {
	address, err := s.sdk.Resolve(ctx, SessionManagementServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewSessionManagementServiceClient(con).Revoke(ctx, request, opts...)
}
