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

func (s Services) NodeGroup() NodeGroupService {
	return NewNodeGroupService(s.sdk)
}

const NodeGroupServiceID conn.ServiceID = "nebius.mk8s.v1.NodeGroupService"

type NodeGroupService interface {
	Get(context.Context, *v1.GetNodeGroupRequest, ...grpc.CallOption) (*v1.NodeGroup, error)
	GetByName(context.Context, *v11.GetByNameRequest, ...grpc.CallOption) (*v1.NodeGroup, error)
	List(context.Context, *v1.ListNodeGroupsRequest, ...grpc.CallOption) (*v1.ListNodeGroupsResponse, error)
	Filter(context.Context, *v1.ListNodeGroupsRequest, ...grpc.CallOption) iter.Seq2[*v1.NodeGroup, error]
	Create(context.Context, *v1.CreateNodeGroupRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateNodeGroupRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteNodeGroupRequest, ...grpc.CallOption) (operations.Operation, error)
	Upgrade(context.Context, *v1.UpgradeNodeGroupRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type nodeGroupService struct {
	sdk iface.SDK
}

func NewNodeGroupService(sdk iface.SDK) NodeGroupService {
	return nodeGroupService{
		sdk: sdk,
	}
}

func (s nodeGroupService) Get(ctx context.Context, request *v1.GetNodeGroupRequest, opts ...grpc.CallOption) (*v1.NodeGroup, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewNodeGroupServiceClient(con).Get(ctx, request, opts...)
}

func (s nodeGroupService) GetByName(ctx context.Context, request *v11.GetByNameRequest, opts ...grpc.CallOption) (*v1.NodeGroup, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewNodeGroupServiceClient(con).GetByName(ctx, request, opts...)
}

func (s nodeGroupService) List(ctx context.Context, request *v1.ListNodeGroupsRequest, opts ...grpc.CallOption) (*v1.ListNodeGroupsResponse, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewNodeGroupServiceClient(con).List(ctx, request, opts...)
}

func (s nodeGroupService) Filter(ctx context.Context, request *v1.ListNodeGroupsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.NodeGroup, error] {
	req := proto.Clone(request).(*v1.ListNodeGroupsRequest)
	return func(yield func(*v1.NodeGroup, error) bool) {
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

func (s nodeGroupService) Create(ctx context.Context, request *v1.CreateNodeGroupRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewNodeGroupServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s nodeGroupService) Update(ctx context.Context, request *v1.UpdateNodeGroupRequest, opts ...grpc.CallOption) (operations.Operation, error) {
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
	op, err := v1.NewNodeGroupServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s nodeGroupService) Delete(ctx context.Context, request *v1.DeleteNodeGroupRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewNodeGroupServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s nodeGroupService) Upgrade(ctx context.Context, request *v1.UpgradeNodeGroupRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewNodeGroupServiceClient(con).Upgrade(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s nodeGroupService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
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

func (s nodeGroupService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, NodeGroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
