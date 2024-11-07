package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[GroupMembershipServiceID] = "cpl.iam"
}

func (s Services) GroupMembership() GroupMembershipService {
	return NewGroupMembershipService(s.sdk)
}

const GroupMembershipServiceID conn.ServiceID = "nebius.iam.v1.GroupMembershipService"

type GroupMembershipService interface {
	Create(context.Context, *v1.CreateGroupMembershipRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetGroupMembershipRequest, ...grpc.CallOption) (*v1.GroupMembership, error)
	Delete(context.Context, *v1.DeleteGroupMembershipRequest, ...grpc.CallOption) (operations.Operation, error)
	ListMembers(context.Context, *v1.ListGroupMembershipsRequest, ...grpc.CallOption) (*v1.ListGroupMembershipsResponse, error)
	ListMemberOf(context.Context, *v1.ListMemberOfRequest, ...grpc.CallOption) (*v1.ListMemberOfResponse, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type groupMembershipService struct {
	sdk iface.SDK
}

func NewGroupMembershipService(sdk iface.SDK) GroupMembershipService {
	return groupMembershipService{
		sdk: sdk,
	}
}

func (s groupMembershipService) Create(ctx context.Context, request *v1.CreateGroupMembershipRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewGroupMembershipServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s groupMembershipService) Get(ctx context.Context, request *v1.GetGroupMembershipRequest, opts ...grpc.CallOption) (*v1.GroupMembership, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupMembershipServiceClient(con).Get(ctx, request, opts...)
}

func (s groupMembershipService) Delete(ctx context.Context, request *v1.DeleteGroupMembershipRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewGroupMembershipServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s groupMembershipService) ListMembers(ctx context.Context, request *v1.ListGroupMembershipsRequest, opts ...grpc.CallOption) (*v1.ListGroupMembershipsResponse, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupMembershipServiceClient(con).ListMembers(ctx, request, opts...)
}

func (s groupMembershipService) ListMemberOf(ctx context.Context, request *v1.ListMemberOfRequest, opts ...grpc.CallOption) (*v1.ListMemberOfResponse, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupMembershipServiceClient(con).ListMemberOf(ctx, request, opts...)
}

func (s groupMembershipService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
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

func (s groupMembershipService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, GroupMembershipServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
