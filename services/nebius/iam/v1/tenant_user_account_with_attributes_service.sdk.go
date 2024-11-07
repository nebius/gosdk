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
	conn.ConventionResolverServiceIDToNameMap[TenantUserAccountWithAttributesServiceID] = "cpl.iam"
}

func (s Services) TenantUserAccountWithAttributes() TenantUserAccountWithAttributesService {
	return NewTenantUserAccountWithAttributesService(s.sdk)
}

const TenantUserAccountWithAttributesServiceID conn.ServiceID = "nebius.iam.v1.TenantUserAccountWithAttributesService"

type TenantUserAccountWithAttributesService interface {
	Get(context.Context, *v1.GetTenantUserAccountWithAttributesRequest, ...grpc.CallOption) (*v1.TenantUserAccountWithAttributes, error)
	List(context.Context, *v1.ListTenantUserAccountsWithAttributesRequest, ...grpc.CallOption) (*v1.ListTenantUserAccountsWithAttributesResponse, error)
	Filter(context.Context, *v1.ListTenantUserAccountsWithAttributesRequest, ...grpc.CallOption) iter.Seq2[*v1.TenantUserAccountWithAttributes, error]
}

type tenantUserAccountWithAttributesService struct {
	sdk iface.SDK
}

func NewTenantUserAccountWithAttributesService(sdk iface.SDK) TenantUserAccountWithAttributesService {
	return tenantUserAccountWithAttributesService{
		sdk: sdk,
	}
}

func (s tenantUserAccountWithAttributesService) Get(ctx context.Context, request *v1.GetTenantUserAccountWithAttributesRequest, opts ...grpc.CallOption) (*v1.TenantUserAccountWithAttributes, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountWithAttributesServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantUserAccountWithAttributesServiceClient(con).Get(ctx, request, opts...)
}

func (s tenantUserAccountWithAttributesService) List(ctx context.Context, request *v1.ListTenantUserAccountsWithAttributesRequest, opts ...grpc.CallOption) (*v1.ListTenantUserAccountsWithAttributesResponse, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountWithAttributesServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantUserAccountWithAttributesServiceClient(con).List(ctx, request, opts...)
}

func (s tenantUserAccountWithAttributesService) Filter(ctx context.Context, request *v1.ListTenantUserAccountsWithAttributesRequest, opts ...grpc.CallOption) iter.Seq2[*v1.TenantUserAccountWithAttributes, error] {
	req := proto.Clone(request).(*v1.ListTenantUserAccountsWithAttributesRequest)
	return func(yield func(*v1.TenantUserAccountWithAttributes, error) bool) {
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
