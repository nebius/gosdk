package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Allocation() AllocationService {
	return NewAllocationService(s.sdk)
}

const AllocationServiceID conn.ServiceID = "nebius.vpc.v1.AllocationService"

type AllocationService interface {
	Get(context.Context, *v1.GetAllocationRequest, ...grpc.CallOption) (*v1.Allocation, error)
	GetByName(context.Context, *v1.GetAllocationByNameRequest, ...grpc.CallOption) (*v1.Allocation, error)
	List(context.Context, *v1.ListAllocationsRequest, ...grpc.CallOption) (*v1.ListAllocationsResponse, error)
	Filter(context.Context, *v1.ListAllocationsRequest, ...grpc.CallOption) iter.Seq2[*v1.Allocation, error]
	Create(context.Context, *v1.CreateAllocationRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateAllocationRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteAllocationRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type allocationService struct {
	sdk iface.SDK
}

func NewAllocationService(sdk iface.SDK) AllocationService {
	return allocationService{
		sdk: sdk,
	}
}

func (s allocationService) Get(ctx context.Context, request *v1.GetAllocationRequest, opts ...grpc.CallOption) (*v1.Allocation, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAllocationServiceClient(con).Get(ctx, request, opts...)
}

func (s allocationService) GetByName(ctx context.Context, request *v1.GetAllocationByNameRequest, opts ...grpc.CallOption) (*v1.Allocation, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAllocationServiceClient(con).GetByName(ctx, request, opts...)
}

func (s allocationService) List(ctx context.Context, request *v1.ListAllocationsRequest, opts ...grpc.CallOption) (*v1.ListAllocationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAllocationServiceClient(con).List(ctx, request, opts...)
}

func (s allocationService) Filter(ctx context.Context, request *v1.ListAllocationsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Allocation, error] {
	req := proto.Clone(request).(*v1.ListAllocationsRequest)
	return func(yield func(*v1.Allocation, error) bool) {
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

func (s allocationService) Create(ctx context.Context, request *v1.CreateAllocationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAllocationServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s allocationService) Update(ctx context.Context, request *v1.UpdateAllocationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAllocationServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s allocationService) Delete(ctx context.Context, request *v1.DeleteAllocationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAllocationServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s allocationService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
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

func (s allocationService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, AllocationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
