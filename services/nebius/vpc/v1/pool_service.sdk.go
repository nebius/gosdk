// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	grpcheader "github.com/nebius/gosdk/proto/fieldmask/grpcheader"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[PoolServiceID] = "vpc"
}

func (s Services) Pool() PoolService {
	return NewPoolService(s.sdk)
}

const PoolServiceID conn.ServiceID = "nebius.vpc.v1.PoolService"

type PoolService interface {
	Get(context.Context, *v1.GetPoolRequest, ...grpc.CallOption) (*v1.Pool, error)
	GetByName(context.Context, *v1.GetPoolByNameRequest, ...grpc.CallOption) (*v1.Pool, error)
	List(context.Context, *v1.ListPoolsRequest, ...grpc.CallOption) (*v1.ListPoolsResponse, error)
	Filter(context.Context, *v1.ListPoolsRequest, ...grpc.CallOption) iter.Seq2[*v1.Pool, error]
	ListBySourcePool(context.Context, *v1.ListPoolsBySourcePoolRequest, ...grpc.CallOption) (*v1.ListPoolsResponse, error)
	Create(context.Context, *v1.CreatePoolRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdatePoolRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeletePoolRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type poolService struct {
	sdk iface.SDK
}

func NewPoolService(sdk iface.SDK) PoolService {
	return poolService{
		sdk: sdk,
	}
}

func (s poolService) Get(ctx context.Context, request *v1.GetPoolRequest, opts ...grpc.CallOption) (*v1.Pool, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewPoolServiceClient(con).Get(ctx, request, opts...)
}

func (s poolService) GetByName(ctx context.Context, request *v1.GetPoolByNameRequest, opts ...grpc.CallOption) (*v1.Pool, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewPoolServiceClient(con).GetByName(ctx, request, opts...)
}

func (s poolService) List(ctx context.Context, request *v1.ListPoolsRequest, opts ...grpc.CallOption) (*v1.ListPoolsResponse, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewPoolServiceClient(con).List(ctx, request, opts...)
}

func (s poolService) Filter(ctx context.Context, request *v1.ListPoolsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Pool, error] {
	req := proto.Clone(request).(*v1.ListPoolsRequest)
	return func(yield func(*v1.Pool, error) bool) {
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

func (s poolService) ListBySourcePool(ctx context.Context, request *v1.ListPoolsBySourcePoolRequest, opts ...grpc.CallOption) (*v1.ListPoolsResponse, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewPoolServiceClient(con).ListBySourcePool(ctx, request, opts...)
}

func (s poolService) Create(ctx context.Context, request *v1.CreatePoolRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewPoolServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s poolService) Update(ctx context.Context, request *v1.UpdatePoolRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewPoolServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s poolService) Delete(ctx context.Context, request *v1.DeletePoolRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewPoolServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s poolService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
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

func (s poolService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
