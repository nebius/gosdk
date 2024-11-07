package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	alphaops "github.com/nebius/gosdk/operations/alphaops"
	v1alpha11 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/mk8s/v1alpha1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) NodeGroup() NodeGroupService {
	return NewNodeGroupService(s.sdk)
}

const NodeGroupServiceID conn.ServiceID = "nebius.mk8s.v1alpha1.NodeGroupService"

type NodeGroupService interface {
	Get(context.Context, *v1alpha1.GetNodeGroupRequest, ...grpc.CallOption) (*v1alpha1.NodeGroup, error)
	GetByName(context.Context, *v1alpha1.GetNodeGroupByNameRequest, ...grpc.CallOption) (*v1alpha1.NodeGroup, error)
	List(context.Context, *v1alpha1.ListNodeGroupsRequest, ...grpc.CallOption) (*v1alpha1.ListNodeGroupsResponse, error)
	Filter(context.Context, *v1alpha1.ListNodeGroupsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.NodeGroup, error]
	Create(context.Context, *v1alpha1.CreateNodeGroupRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Update(context.Context, *v1alpha1.UpdateNodeGroupRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Delete(context.Context, *v1alpha1.DeleteNodeGroupRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Upgrade(context.Context, *v1alpha1.UpgradeNodeGroupRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	GetOperation(context.Context, *v1alpha11.GetOperationRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	ListOperations(context.Context, *v1alpha11.ListOperationsRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
}

type nodeGroupService struct {
	sdk iface.SDK
}

func NewNodeGroupService(sdk iface.SDK) NodeGroupService {
	return nodeGroupService{
		sdk: sdk,
	}
}

func (s nodeGroupService) Get(ctx context.Context, request *v1alpha1.GetNodeGroupRequest, opts ...grpc.CallOption) (*v1alpha1.NodeGroup, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewNodeGroupServiceClient(con).Get(ctx, request, opts...)
}

func (s nodeGroupService) GetByName(ctx context.Context, request *v1alpha1.GetNodeGroupByNameRequest, opts ...grpc.CallOption) (*v1alpha1.NodeGroup, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewNodeGroupServiceClient(con).GetByName(ctx, request, opts...)
}

func (s nodeGroupService) List(ctx context.Context, request *v1alpha1.ListNodeGroupsRequest, opts ...grpc.CallOption) (*v1alpha1.ListNodeGroupsResponse, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewNodeGroupServiceClient(con).List(ctx, request, opts...)
}

func (s nodeGroupService) Filter(ctx context.Context, request *v1alpha1.ListNodeGroupsRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.NodeGroup, error] {
	req := proto.Clone(request).(*v1alpha1.ListNodeGroupsRequest)
	return func(yield func(*v1alpha1.NodeGroup, error) bool) {
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

func (s nodeGroupService) Create(ctx context.Context, request *v1alpha1.CreateNodeGroupRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewNodeGroupServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s nodeGroupService) Update(ctx context.Context, request *v1alpha1.UpdateNodeGroupRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewNodeGroupServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s nodeGroupService) Delete(ctx context.Context, request *v1alpha1.DeleteNodeGroupRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewNodeGroupServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s nodeGroupService) Upgrade(ctx context.Context, request *v1alpha1.UpgradeNodeGroupRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewNodeGroupServiceClient(con).Upgrade(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s nodeGroupService) GetOperation(ctx context.Context, request *v1alpha11.GetOperationRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
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

func (s nodeGroupService) ListOperations(ctx context.Context, request *v1alpha11.ListOperationsRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
