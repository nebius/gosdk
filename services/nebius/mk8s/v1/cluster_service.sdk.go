package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/mk8s/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Cluster() ClusterService {
	return NewClusterService(s.sdk)
}

const ClusterServiceID conn.ServiceID = "nebius.mk8s.v1.ClusterService"

type ClusterService interface {
	Get(context.Context, *v1.GetClusterRequest, ...grpc.CallOption) (*v1.Cluster, error)
	GetByName(context.Context, *v11.GetByNameRequest, ...grpc.CallOption) (*v1.Cluster, error)
	List(context.Context, *v1.ListClustersRequest, ...grpc.CallOption) (*v1.ListClustersResponse, error)
	Filter(context.Context, *v1.ListClustersRequest, ...grpc.CallOption) iter.Seq2[*v1.Cluster, error]
	Create(context.Context, *v1.CreateClusterRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateClusterRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteClusterRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type clusterService struct {
	sdk iface.SDK
}

func NewClusterService(sdk iface.SDK) ClusterService {
	return clusterService{
		sdk: sdk,
	}
}

func (s clusterService) Get(ctx context.Context, request *v1.GetClusterRequest, opts ...grpc.CallOption) (*v1.Cluster, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewClusterServiceClient(con).Get(ctx, request, opts...)
}

func (s clusterService) GetByName(ctx context.Context, request *v11.GetByNameRequest, opts ...grpc.CallOption) (*v1.Cluster, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewClusterServiceClient(con).GetByName(ctx, request, opts...)
}

func (s clusterService) List(ctx context.Context, request *v1.ListClustersRequest, opts ...grpc.CallOption) (*v1.ListClustersResponse, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewClusterServiceClient(con).List(ctx, request, opts...)
}

func (s clusterService) Filter(ctx context.Context, request *v1.ListClustersRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Cluster, error] {
	req := proto.Clone(request).(*v1.ListClustersRequest)
	return func(yield func(*v1.Cluster, error) bool) {
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

func (s clusterService) Create(ctx context.Context, request *v1.CreateClusterRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewClusterServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s clusterService) Update(ctx context.Context, request *v1.UpdateClusterRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewClusterServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s clusterService) Delete(ctx context.Context, request *v1.DeleteClusterRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewClusterServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s clusterService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
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

func (s clusterService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
