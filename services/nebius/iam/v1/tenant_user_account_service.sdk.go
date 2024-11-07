package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[TenantUserAccountServiceID] = "cpl.iam"
}

func (s Services) TenantUserAccount() TenantUserAccountService {
	return NewTenantUserAccountService(s.sdk)
}

const TenantUserAccountServiceID conn.ServiceID = "nebius.iam.v1.TenantUserAccountService"

type TenantUserAccountService interface {
	Get(context.Context, *v1.GetTenantUserAccountRequest, ...grpc.CallOption) (*v1.TenantUserAccount, error)
	List(context.Context, *v1.ListTenantUserAccountsRequest, ...grpc.CallOption) (*v1.ListTenantUserAccountsResponse, error)
	Filter(context.Context, *v1.ListTenantUserAccountsRequest, ...grpc.CallOption) iter.Seq2[*v1.TenantUserAccount, error]
	Block(context.Context, *v1.BlockTenantUserAccountRequest, ...grpc.CallOption) (operations.Operation, error)
	Unblock(context.Context, *v1.UnblockTenantUserAccountRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type tenantUserAccountService struct {
	sdk iface.SDK
}

func NewTenantUserAccountService(sdk iface.SDK) TenantUserAccountService {
	return tenantUserAccountService{
		sdk: sdk,
	}
}

func (s tenantUserAccountService) Get(ctx context.Context, request *v1.GetTenantUserAccountRequest, opts ...grpc.CallOption) (*v1.TenantUserAccount, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantUserAccountServiceClient(con).Get(ctx, request, opts...)
}

func (s tenantUserAccountService) List(ctx context.Context, request *v1.ListTenantUserAccountsRequest, opts ...grpc.CallOption) (*v1.ListTenantUserAccountsResponse, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewTenantUserAccountServiceClient(con).List(ctx, request, opts...)
}

func (s tenantUserAccountService) Filter(ctx context.Context, request *v1.ListTenantUserAccountsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.TenantUserAccount, error] {
	req := proto.Clone(request).(*v1.ListTenantUserAccountsRequest)
	return func(yield func(*v1.TenantUserAccount, error) bool) {
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

func (s tenantUserAccountService) Block(ctx context.Context, request *v1.BlockTenantUserAccountRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewTenantUserAccountServiceClient(con).Block(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s tenantUserAccountService) Unblock(ctx context.Context, request *v1.UnblockTenantUserAccountRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewTenantUserAccountServiceClient(con).Unblock(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s tenantUserAccountService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	client := v11.NewOperationServiceClient(con)
	op, err := client.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, client)
}

func (s tenantUserAccountService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, TenantUserAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
