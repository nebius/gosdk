package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/compute/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) GpuCluster() GpuClusterService {
	return NewGpuClusterService(s.sdk)
}

const GpuClusterServiceID conn.ServiceID = "nebius.compute.v1.GpuClusterService"

type GpuClusterService interface {
	Get(context.Context, *v1.GetGpuClusterRequest, ...grpc.CallOption) (*v1.GpuCluster, error)
	GetByName(context.Context, *v11.GetByNameRequest, ...grpc.CallOption) (*v1.GpuCluster, error)
	List(context.Context, *v1.ListGpuClustersRequest, ...grpc.CallOption) (*v1.ListGpuClustersResponse, error)
	Filter(context.Context, *v1.ListGpuClustersRequest, ...grpc.CallOption) iter.Seq2[*v1.GpuCluster, error]
	Create(context.Context, *v1.CreateGpuClusterRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateGpuClusterRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteGpuClusterRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperationsByParent(context.Context, *v1.ListOperationsByParentRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type gpuClusterService struct {
	sdk iface.SDK
}

func NewGpuClusterService(sdk iface.SDK) GpuClusterService {
	return gpuClusterService{
		sdk: sdk,
	}
}

func (s gpuClusterService) Get(ctx context.Context, request *v1.GetGpuClusterRequest, opts ...grpc.CallOption) (*v1.GpuCluster, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGpuClusterServiceClient(con).Get(ctx, request, opts...)
}

func (s gpuClusterService) GetByName(ctx context.Context, request *v11.GetByNameRequest, opts ...grpc.CallOption) (*v1.GpuCluster, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGpuClusterServiceClient(con).GetByName(ctx, request, opts...)
}

func (s gpuClusterService) List(ctx context.Context, request *v1.ListGpuClustersRequest, opts ...grpc.CallOption) (*v1.ListGpuClustersResponse, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGpuClusterServiceClient(con).List(ctx, request, opts...)
}

func (s gpuClusterService) Filter(ctx context.Context, request *v1.ListGpuClustersRequest, opts ...grpc.CallOption) iter.Seq2[*v1.GpuCluster, error] {
	req := proto.Clone(request).(*v1.ListGpuClustersRequest)
	return func(yield func(*v1.GpuCluster, error) bool) {
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

func (s gpuClusterService) Create(ctx context.Context, request *v1.CreateGpuClusterRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewGpuClusterServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s gpuClusterService) Update(ctx context.Context, request *v1.UpdateGpuClusterRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewGpuClusterServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s gpuClusterService) Delete(ctx context.Context, request *v1.DeleteGpuClusterRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewGpuClusterServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s gpuClusterService) ListOperationsByParent(ctx context.Context, request *v1.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGpuClusterServiceClient(con).ListOperationsByParent(ctx, request, opts...)
}

func (s gpuClusterService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
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

func (s gpuClusterService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, GpuClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
