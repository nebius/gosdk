// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	alphaops "github.com/nebius/gosdk/operations/alphaops"
	grpcheader "github.com/nebius/gosdk/proto/fieldmask/grpcheader"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha11 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/msp/postgresql/v1alpha1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[ClusterServiceID] = "postgresql.msp"
}

func (s Services) Cluster() ClusterService {
	return NewClusterService(s.sdk)
}

const ClusterServiceID conn.ServiceID = "nebius.msp.postgresql.v1alpha1.ClusterService"

type ClusterService interface {
	Get(context.Context, *v1alpha1.GetClusterRequest, ...grpc.CallOption) (*v1alpha1.Cluster, error)
	GetByName(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v1alpha1.Cluster, error)
	GetForBackup(context.Context, *v1alpha1.GetClusterForBackupRequest, ...grpc.CallOption) (*v1alpha1.Cluster, error)
	List(context.Context, *v1alpha1.ListClustersRequest, ...grpc.CallOption) (*v1alpha1.ListClustersResponse, error)
	Filter(context.Context, *v1alpha1.ListClustersRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Cluster, error]
	Create(context.Context, *v1alpha1.CreateClusterRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Delete(context.Context, *v1alpha1.DeleteClusterRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Update(context.Context, *v1alpha1.UpdateClusterRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Restore(context.Context, *v1alpha1.RestoreClusterRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Stop(context.Context, *v1alpha1.StopClusterRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Start(context.Context, *v1alpha1.StartClusterRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	GetOperation(context.Context, *v1alpha11.GetOperationRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	ListOperations(context.Context, *v1alpha11.ListOperationsRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
}

type clusterService struct {
	sdk iface.SDK
}

func NewClusterService(sdk iface.SDK) ClusterService {
	return clusterService{
		sdk: sdk,
	}
}

func (s clusterService) Get(ctx context.Context, request *v1alpha1.GetClusterRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewClusterServiceClient(con).Get(ctx, request, opts...)
}

func (s clusterService) GetByName(ctx context.Context, request *v1.GetByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewClusterServiceClient(con).GetByName(ctx, request, opts...)
}

func (s clusterService) GetForBackup(ctx context.Context, request *v1alpha1.GetClusterForBackupRequest, opts ...grpc.CallOption) (*v1alpha1.Cluster, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewClusterServiceClient(con).GetForBackup(ctx, request, opts...)
}

func (s clusterService) List(ctx context.Context, request *v1alpha1.ListClustersRequest, opts ...grpc.CallOption) (*v1alpha1.ListClustersResponse, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewClusterServiceClient(con).List(ctx, request, opts...)
}

func (s clusterService) Filter(ctx context.Context, request *v1alpha1.ListClustersRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Cluster, error] {
	req := proto.Clone(request).(*v1alpha1.ListClustersRequest)
	return func(yield func(*v1alpha1.Cluster, error) bool) {
		for {
			res, err := s.List(ctx, req, opts...)
			if err != nil {
				yield(nil, err)
				return
			}

			for _, item := range res.GetClusters() {
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

func (s clusterService) Create(ctx context.Context, request *v1alpha1.CreateClusterRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewClusterServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s clusterService) Delete(ctx context.Context, request *v1alpha1.DeleteClusterRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewClusterServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s clusterService) Update(ctx context.Context, request *v1alpha1.UpdateClusterRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
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
	op, err := v1alpha1.NewClusterServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s clusterService) Restore(ctx context.Context, request *v1alpha1.RestoreClusterRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewClusterServiceClient(con).Restore(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s clusterService) Stop(ctx context.Context, request *v1alpha1.StopClusterRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewClusterServiceClient(con).Stop(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s clusterService) Start(ctx context.Context, request *v1alpha1.StartClusterRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewClusterServiceClient(con).Start(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s clusterService) GetOperation(ctx context.Context, request *v1alpha11.GetOperationRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	client := v1alpha11.NewOperationServiceClient(con)
	op, err := client.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, client)
}

func (s clusterService) ListOperations(ctx context.Context, request *v1alpha11.ListOperationsRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ClusterServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
