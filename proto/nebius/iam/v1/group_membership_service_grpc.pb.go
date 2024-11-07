// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: nebius/iam/v1/group_membership_service.proto

package v1

import (
	context "context"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GroupMembershipService_Create_FullMethodName       = "/nebius.iam.v1.GroupMembershipService/Create"
	GroupMembershipService_Get_FullMethodName          = "/nebius.iam.v1.GroupMembershipService/Get"
	GroupMembershipService_Delete_FullMethodName       = "/nebius.iam.v1.GroupMembershipService/Delete"
	GroupMembershipService_ListMembers_FullMethodName  = "/nebius.iam.v1.GroupMembershipService/ListMembers"
	GroupMembershipService_ListMemberOf_FullMethodName = "/nebius.iam.v1.GroupMembershipService/ListMemberOf"
)

// GroupMembershipServiceClient is the client API for GroupMembershipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupMembershipServiceClient interface {
	Create(ctx context.Context, in *CreateGroupMembershipRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Get(ctx context.Context, in *GetGroupMembershipRequest, opts ...grpc.CallOption) (*GroupMembership, error)
	Delete(ctx context.Context, in *DeleteGroupMembershipRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	ListMembers(ctx context.Context, in *ListGroupMembershipsRequest, opts ...grpc.CallOption) (*ListGroupMembershipsResponse, error)
	ListMemberOf(ctx context.Context, in *ListMemberOfRequest, opts ...grpc.CallOption) (*ListMemberOfResponse, error)
}

type groupMembershipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupMembershipServiceClient(cc grpc.ClientConnInterface) GroupMembershipServiceClient {
	return &groupMembershipServiceClient{cc}
}

func (c *groupMembershipServiceClient) Create(ctx context.Context, in *CreateGroupMembershipRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupMembershipService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMembershipServiceClient) Get(ctx context.Context, in *GetGroupMembershipRequest, opts ...grpc.CallOption) (*GroupMembership, error) {
	out := new(GroupMembership)
	err := c.cc.Invoke(ctx, GroupMembershipService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMembershipServiceClient) Delete(ctx context.Context, in *DeleteGroupMembershipRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, GroupMembershipService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMembershipServiceClient) ListMembers(ctx context.Context, in *ListGroupMembershipsRequest, opts ...grpc.CallOption) (*ListGroupMembershipsResponse, error) {
	out := new(ListGroupMembershipsResponse)
	err := c.cc.Invoke(ctx, GroupMembershipService_ListMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMembershipServiceClient) ListMemberOf(ctx context.Context, in *ListMemberOfRequest, opts ...grpc.CallOption) (*ListMemberOfResponse, error) {
	out := new(ListMemberOfResponse)
	err := c.cc.Invoke(ctx, GroupMembershipService_ListMemberOf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupMembershipServiceServer is the server API for GroupMembershipService service.
// All implementations should embed UnimplementedGroupMembershipServiceServer
// for forward compatibility
type GroupMembershipServiceServer interface {
	Create(context.Context, *CreateGroupMembershipRequest) (*v1.Operation, error)
	Get(context.Context, *GetGroupMembershipRequest) (*GroupMembership, error)
	Delete(context.Context, *DeleteGroupMembershipRequest) (*v1.Operation, error)
	ListMembers(context.Context, *ListGroupMembershipsRequest) (*ListGroupMembershipsResponse, error)
	ListMemberOf(context.Context, *ListMemberOfRequest) (*ListMemberOfResponse, error)
}

// UnimplementedGroupMembershipServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGroupMembershipServiceServer struct {
}

func (UnimplementedGroupMembershipServiceServer) Create(context.Context, *CreateGroupMembershipRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupMembershipServiceServer) Get(context.Context, *GetGroupMembershipRequest) (*GroupMembership, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGroupMembershipServiceServer) Delete(context.Context, *DeleteGroupMembershipRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupMembershipServiceServer) ListMembers(context.Context, *ListGroupMembershipsRequest) (*ListGroupMembershipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedGroupMembershipServiceServer) ListMemberOf(context.Context, *ListMemberOfRequest) (*ListMemberOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemberOf not implemented")
}

// UnsafeGroupMembershipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupMembershipServiceServer will
// result in compilation errors.
type UnsafeGroupMembershipServiceServer interface {
	mustEmbedUnimplementedGroupMembershipServiceServer()
}

func RegisterGroupMembershipServiceServer(s grpc.ServiceRegistrar, srv GroupMembershipServiceServer) {
	s.RegisterService(&GroupMembershipService_ServiceDesc, srv)
}

func _GroupMembershipService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMembershipServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMembershipService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMembershipServiceServer).Create(ctx, req.(*CreateGroupMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMembershipService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMembershipServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMembershipService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMembershipServiceServer).Get(ctx, req.(*GetGroupMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMembershipService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupMembershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMembershipServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMembershipService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMembershipServiceServer).Delete(ctx, req.(*DeleteGroupMembershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMembershipService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupMembershipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMembershipServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMembershipService_ListMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMembershipServiceServer).ListMembers(ctx, req.(*ListGroupMembershipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMembershipService_ListMemberOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMemberOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMembershipServiceServer).ListMemberOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMembershipService_ListMemberOf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMembershipServiceServer).ListMemberOf(ctx, req.(*ListMemberOfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupMembershipService_ServiceDesc is the grpc.ServiceDesc for GroupMembershipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupMembershipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.GroupMembershipService",
	HandlerType: (*GroupMembershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GroupMembershipService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GroupMembershipService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupMembershipService_Delete_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _GroupMembershipService_ListMembers_Handler,
		},
		{
			MethodName: "ListMemberOf",
			Handler:    _GroupMembershipService_ListMemberOf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/group_membership_service.proto",
}
