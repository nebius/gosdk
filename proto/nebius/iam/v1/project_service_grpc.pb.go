// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: nebius/iam/v1/project_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProjectService_Get_FullMethodName       = "/nebius.iam.v1.ProjectService/Get"
	ProjectService_GetByName_FullMethodName = "/nebius.iam.v1.ProjectService/GetByName"
	ProjectService_List_FullMethodName      = "/nebius.iam.v1.ProjectService/List"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Container, error)
	GetByName(ctx context.Context, in *GetProjectByNameRequest, opts ...grpc.CallOption) (*Container, error)
	List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Container, error) {
	out := new(Container)
	err := c.cc.Invoke(ctx, ProjectService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetByName(ctx context.Context, in *GetProjectByNameRequest, opts ...grpc.CallOption) (*Container, error) {
	out := new(Container)
	err := c.cc.Invoke(ctx, ProjectService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations should embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	Get(context.Context, *GetProjectRequest) (*Container, error)
	GetByName(context.Context, *GetProjectByNameRequest) (*Container, error)
	List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
}

// UnimplementedProjectServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) Get(context.Context, *GetProjectRequest) (*Container, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProjectServiceServer) GetByName(context.Context, *GetProjectByNameRequest) (*Container, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedProjectServiceServer) List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Get(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetByName(ctx, req.(*GetProjectByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).List(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProjectService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _ProjectService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProjectService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/project_service.proto",
}
