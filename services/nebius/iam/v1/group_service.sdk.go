package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[GroupServiceID] = "cpl.iam"
}

func (s Services) Group() GroupService {
	return NewGroupService(s.sdk)
}

const GroupServiceID conn.ServiceID = "nebius.iam.v1.GroupService"

type GroupService interface {
	Get(context.Context, *v1.GetGroupRequest, ...grpc.CallOption) (*v1.Group, error)
	GetByName(context.Context, *v1.GetGroupByNameRequest, ...grpc.CallOption) (*v1.Group, error)
	List(context.Context, *v1.ListGroupsRequest, ...grpc.CallOption) (*v1.ListGroupsResponse, error)
	Filter(context.Context, *v1.ListGroupsRequest, ...grpc.CallOption) iter.Seq2[*v1.Group, error]
}

type groupService struct {
	sdk iface.SDK
}

func NewGroupService(sdk iface.SDK) GroupService {
	return groupService{
		sdk: sdk,
	}
}

func (s groupService) Get(ctx context.Context, request *v1.GetGroupRequest, opts ...grpc.CallOption) (*v1.Group, error) {
	address, err := s.sdk.Resolve(ctx, GroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupServiceClient(con).Get(ctx, request, opts...)
}

func (s groupService) GetByName(ctx context.Context, request *v1.GetGroupByNameRequest, opts ...grpc.CallOption) (*v1.Group, error) {
	address, err := s.sdk.Resolve(ctx, GroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupServiceClient(con).GetByName(ctx, request, opts...)
}

func (s groupService) List(ctx context.Context, request *v1.ListGroupsRequest, opts ...grpc.CallOption) (*v1.ListGroupsResponse, error) {
	address, err := s.sdk.Resolve(ctx, GroupServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewGroupServiceClient(con).List(ctx, request, opts...)
}

func (s groupService) Filter(ctx context.Context, request *v1.ListGroupsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Group, error] {
	req := proto.Clone(request).(*v1.ListGroupsRequest)
	return func(yield func(*v1.Group, error) bool) {
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
