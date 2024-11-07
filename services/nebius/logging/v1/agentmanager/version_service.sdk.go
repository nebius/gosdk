package agentmanager

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	agentmanager "github.com/nebius/gosdk/proto/nebius/logging/v1/agentmanager"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[VersionServiceID] = "observability-agent-manager"
}

func (s Services) Version() VersionService {
	return NewVersionService(s.sdk)
}

const VersionServiceID conn.ServiceID = "nebius.logging.agentmanager.v1.VersionService"

type VersionService interface {
	GetVersion(context.Context, *agentmanager.GetVersionRequest, ...grpc.CallOption) (*agentmanager.GetVersionResponse, error)
}

type versionService struct {
	sdk iface.SDK
}

func NewVersionService(sdk iface.SDK) VersionService {
	return versionService{
		sdk: sdk,
	}
}

func (s versionService) GetVersion(ctx context.Context, request *agentmanager.GetVersionRequest, opts ...grpc.CallOption) (*agentmanager.GetVersionResponse, error) {
	address, err := s.sdk.Resolve(ctx, VersionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return agentmanager.NewVersionServiceClient(con).GetVersion(ctx, request, opts...)
}
