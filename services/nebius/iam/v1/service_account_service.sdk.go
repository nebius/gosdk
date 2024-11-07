package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[ServiceAccountServiceID] = "cpl.iam"
}

func (s Services) ServiceAccount() ServiceAccountService {
	return NewServiceAccountService(s.sdk)
}

const ServiceAccountServiceID conn.ServiceID = "nebius.iam.v1.ServiceAccountService"

type ServiceAccountService interface {
	Create(context.Context, *v1.CreateServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetServiceAccountRequest, ...grpc.CallOption) (*v1.ServiceAccount, error)
	GetByName(context.Context, *v1.GetServiceAccountByNameRequest, ...grpc.CallOption) (*v1.ServiceAccount, error)
	List(context.Context, *v1.ListServiceAccountRequest, ...grpc.CallOption) (*v1.ListServiceAccountResponse, error)
	Filter(context.Context, *v1.ListServiceAccountRequest, ...grpc.CallOption) iter.Seq2[*v1.ServiceAccount, error]
	Update(context.Context, *v1.UpdateServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type serviceAccountService struct {
	sdk iface.SDK
}

func NewServiceAccountService(sdk iface.SDK) ServiceAccountService {
	return serviceAccountService{
		sdk: sdk,
	}
}

func (s serviceAccountService) Create(ctx context.Context, request *v1.CreateServiceAccountRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewServiceAccountServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s serviceAccountService) Get(ctx context.Context, request *v1.GetServiceAccountRequest, opts ...grpc.CallOption) (*v1.ServiceAccount, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewServiceAccountServiceClient(con).Get(ctx, request, opts...)
}

func (s serviceAccountService) GetByName(ctx context.Context, request *v1.GetServiceAccountByNameRequest, opts ...grpc.CallOption) (*v1.ServiceAccount, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewServiceAccountServiceClient(con).GetByName(ctx, request, opts...)
}

func (s serviceAccountService) List(ctx context.Context, request *v1.ListServiceAccountRequest, opts ...grpc.CallOption) (*v1.ListServiceAccountResponse, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewServiceAccountServiceClient(con).List(ctx, request, opts...)
}

func (s serviceAccountService) Filter(ctx context.Context, request *v1.ListServiceAccountRequest, opts ...grpc.CallOption) iter.Seq2[*v1.ServiceAccount, error] {
	req := proto.Clone(request).(*v1.ListServiceAccountRequest)
	return func(yield func(*v1.ServiceAccount, error) bool) {
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

func (s serviceAccountService) Update(ctx context.Context, request *v1.UpdateServiceAccountRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewServiceAccountServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s serviceAccountService) Delete(ctx context.Context, request *v1.DeleteServiceAccountRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewServiceAccountServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s serviceAccountService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
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

func (s serviceAccountService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ServiceAccountServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
