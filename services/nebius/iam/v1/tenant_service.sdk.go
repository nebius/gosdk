package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[TenantServiceID] = "cpl.iam"
}

func (s Services) Tenant() TenantService {
	return NewTenantService(s.sdk)
}

const TenantServiceID conn.ServiceID = "nebius.iam.v1.TenantService"

type TenantService interface {
	Get(context.Context, *v1.GetTenantRequest, ...grpc.CallOption) (*v1.Container, error)
	List(context.Context, *v1.ListTenantsRequest, ...grpc.CallOption) (*v1.ListTenantsResponse, error)
	Filter(context.Context, *v1.ListTenantsRequest, ...grpc.CallOption) iter.Seq2[*v1.Container, error]
}

type tenantService struct {
	sdk iface.SDK
}

func NewTenantService(sdk iface.SDK) TenantService {
	return tenantService{
		sdk: sdk,
	}
}

func (s tenantService) Get(ctx context.Context, request *v1.GetTenantRequest, opts ...grpc.CallOption) (*v1.Container, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantServiceClient(con).Get(ctx, request, opts...)
}

func (s tenantService) List(ctx context.Context, request *v1.ListTenantsRequest, opts ...grpc.CallOption) (*v1.ListTenantsResponse, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantServiceClient(con).List(ctx, request, opts...)
}

func (s tenantService) Filter(ctx context.Context, request *v1.ListTenantsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Container, error] {
	req := proto.Clone(request).(*v1.ListTenantsRequest)
	return func(yield func(*v1.Container, error) bool) {
		for {
			res, err := s.List(ctx, req, opts...)
			if err != nil {
				yield(nil, err)
				return
			}

			for _, item := range res.GetItems() {
				if !yield(item, nil) {
					return
				}
			}

			if res.GetNextPageToken() == "" {
				return
			}

			req.PageToken = res.GetNextPageToken()
		}
	}
}
