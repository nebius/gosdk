// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v2

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	grpcheader "github.com/nebius/gosdk/proto/fieldmask/grpcheader"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v2 "github.com/nebius/gosdk/proto/nebius/iam/v2"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[TenantServiceID] = "cpl.iam"
}

func (s Services) Tenant() TenantService {
	return NewTenantService(s.sdk)
}

const TenantServiceID conn.ServiceID = "nebius.iam.v2.TenantService"

type TenantService interface {
	Get(context.Context, *v2.GetTenantRequest, ...grpc.CallOption) (*v2.Tenant, error)
	GetByName(context.Context, *v2.GetTenantByNameRequest, ...grpc.CallOption) (*v2.Tenant, error)
	List(context.Context, *v2.ListTenantsRequest, ...grpc.CallOption) (*v2.ListTenantsResponse, error)
	Filter(context.Context, *v2.ListTenantsRequest, ...grpc.CallOption) iter.Seq2[*v2.Tenant, error]
	Update(context.Context, *v2.UpdateTenantRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)
}

type tenantService struct {
	sdk iface.SDK
}

func NewTenantService(sdk iface.SDK) TenantService {
	return tenantService{
		sdk: sdk,
	}
}

func (s tenantService) Get(ctx context.Context, request *v2.GetTenantRequest, opts ...grpc.CallOption) (*v2.Tenant, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v2.NewTenantServiceClient(con).Get(ctx, request, opts...)
}

func (s tenantService) GetByName(ctx context.Context, request *v2.GetTenantByNameRequest, opts ...grpc.CallOption) (*v2.Tenant, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v2.NewTenantServiceClient(con).GetByName(ctx, request, opts...)
}

func (s tenantService) List(ctx context.Context, request *v2.ListTenantsRequest, opts ...grpc.CallOption) (*v2.ListTenantsResponse, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v2.NewTenantServiceClient(con).List(ctx, request, opts...)
}

func (s tenantService) Filter(ctx context.Context, request *v2.ListTenantsRequest, opts ...grpc.CallOption) iter.Seq2[*v2.Tenant, error] {
	req := proto.Clone(request).(*v2.ListTenantsRequest)
	return func(yield func(*v2.Tenant, error) bool) {
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

func (s tenantService) Update(ctx context.Context, request *v2.UpdateTenantRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v2.NewTenantServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s tenantService) GetOperation(ctx context.Context, request *v1.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	client := v1.NewOperationServiceClient(con)
	op, err := client.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, client)
}

func (s tenantService) ListOperations(ctx context.Context, request *v1.ListOperationsRequest, opts ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, TenantServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewOperationServiceClient(con).List(ctx, request, opts...)
}
