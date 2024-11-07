package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/registry/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Registry() RegistryService {
	return NewRegistryService(s.sdk)
}

const RegistryServiceID conn.ServiceID = "nebius.registry.v1.RegistryService"

type RegistryService interface {
	Get(context.Context, *v1.GetRegistryRequest, ...grpc.CallOption) (*v1.Registry, error)
	List(context.Context, *v1.ListRegistriesRequest, ...grpc.CallOption) (*v1.ListRegistriesResponse, error)
	Filter(context.Context, *v1.ListRegistriesRequest, ...grpc.CallOption) iter.Seq2[*v1.Registry, error]
	Create(context.Context, *v1.CreateRegistryRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateRegistryRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteRegistryRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type registryService struct {
	sdk iface.SDK
}

func NewRegistryService(sdk iface.SDK) RegistryService {
	return registryService{
		sdk: sdk,
	}
}

func (s registryService) Get(ctx context.Context, request *v1.GetRegistryRequest, opts ...grpc.CallOption) (*v1.Registry, error) {
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewRegistryServiceClient(con).Get(ctx, request, opts...)
}

func (s registryService) List(ctx context.Context, request *v1.ListRegistriesRequest, opts ...grpc.CallOption) (*v1.ListRegistriesResponse, error) {
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewRegistryServiceClient(con).List(ctx, request, opts...)
}

func (s registryService) Filter(ctx context.Context, request *v1.ListRegistriesRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Registry, error] {
	req := proto.Clone(request).(*v1.ListRegistriesRequest)
	return func(yield func(*v1.Registry, error) bool) {
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

func (s registryService) Create(ctx context.Context, request *v1.CreateRegistryRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewRegistryServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s registryService) Update(ctx context.Context, request *v1.UpdateRegistryRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewRegistryServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s registryService) Delete(ctx context.Context, request *v1.DeleteRegistryRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewRegistryServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s registryService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
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

func (s registryService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, RegistryServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
