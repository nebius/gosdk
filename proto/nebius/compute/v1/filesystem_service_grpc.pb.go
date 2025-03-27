// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/compute/v1/filesystem_service.proto

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
	FilesystemService_Get_FullMethodName                    = "/nebius.compute.v1.FilesystemService/Get"
	FilesystemService_GetByName_FullMethodName              = "/nebius.compute.v1.FilesystemService/GetByName"
	FilesystemService_List_FullMethodName                   = "/nebius.compute.v1.FilesystemService/List"
	FilesystemService_Create_FullMethodName                 = "/nebius.compute.v1.FilesystemService/Create"
	FilesystemService_Update_FullMethodName                 = "/nebius.compute.v1.FilesystemService/Update"
	FilesystemService_Delete_FullMethodName                 = "/nebius.compute.v1.FilesystemService/Delete"
	FilesystemService_ListOperationsByParent_FullMethodName = "/nebius.compute.v1.FilesystemService/ListOperationsByParent"
)

// FilesystemServiceClient is the client API for FilesystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesystemServiceClient interface {
	Get(ctx context.Context, in *GetFilesystemRequest, opts ...grpc.CallOption) (*Filesystem, error)
	GetByName(ctx context.Context, in *v1.GetByNameRequest, opts ...grpc.CallOption) (*Filesystem, error)
	List(ctx context.Context, in *ListFilesystemsRequest, opts ...grpc.CallOption) (*ListFilesystemsResponse, error)
	Create(ctx context.Context, in *CreateFilesystemRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateFilesystemRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteFilesystemRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	ListOperationsByParent(ctx context.Context, in *ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1.ListOperationsResponse, error)
}

type filesystemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesystemServiceClient(cc grpc.ClientConnInterface) FilesystemServiceClient {
	return &filesystemServiceClient{cc}
}

func (c *filesystemServiceClient) Get(ctx context.Context, in *GetFilesystemRequest, opts ...grpc.CallOption) (*Filesystem, error) {
	out := new(Filesystem)
	err := c.cc.Invoke(ctx, FilesystemService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) GetByName(ctx context.Context, in *v1.GetByNameRequest, opts ...grpc.CallOption) (*Filesystem, error) {
	out := new(Filesystem)
	err := c.cc.Invoke(ctx, FilesystemService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) List(ctx context.Context, in *ListFilesystemsRequest, opts ...grpc.CallOption) (*ListFilesystemsResponse, error) {
	out := new(ListFilesystemsResponse)
	err := c.cc.Invoke(ctx, FilesystemService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) Create(ctx context.Context, in *CreateFilesystemRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FilesystemService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) Update(ctx context.Context, in *UpdateFilesystemRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FilesystemService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) Delete(ctx context.Context, in *DeleteFilesystemRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FilesystemService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemServiceClient) ListOperationsByParent(ctx context.Context, in *ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	out := new(v1.ListOperationsResponse)
	err := c.cc.Invoke(ctx, FilesystemService_ListOperationsByParent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesystemServiceServer is the server API for FilesystemService service.
// All implementations should embed UnimplementedFilesystemServiceServer
// for forward compatibility
type FilesystemServiceServer interface {
	Get(context.Context, *GetFilesystemRequest) (*Filesystem, error)
	GetByName(context.Context, *v1.GetByNameRequest) (*Filesystem, error)
	List(context.Context, *ListFilesystemsRequest) (*ListFilesystemsResponse, error)
	Create(context.Context, *CreateFilesystemRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateFilesystemRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteFilesystemRequest) (*v1.Operation, error)
	ListOperationsByParent(context.Context, *ListOperationsByParentRequest) (*v1.ListOperationsResponse, error)
}

// UnimplementedFilesystemServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFilesystemServiceServer struct {
}

func (UnimplementedFilesystemServiceServer) Get(context.Context, *GetFilesystemRequest) (*Filesystem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFilesystemServiceServer) GetByName(context.Context, *v1.GetByNameRequest) (*Filesystem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedFilesystemServiceServer) List(context.Context, *ListFilesystemsRequest) (*ListFilesystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFilesystemServiceServer) Create(context.Context, *CreateFilesystemRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFilesystemServiceServer) Update(context.Context, *UpdateFilesystemRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFilesystemServiceServer) Delete(context.Context, *DeleteFilesystemRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFilesystemServiceServer) ListOperationsByParent(context.Context, *ListOperationsByParentRequest) (*v1.ListOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperationsByParent not implemented")
}

// UnsafeFilesystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesystemServiceServer will
// result in compilation errors.
type UnsafeFilesystemServiceServer interface {
	mustEmbedUnimplementedFilesystemServiceServer()
}

func RegisterFilesystemServiceServer(s grpc.ServiceRegistrar, srv FilesystemServiceServer) {
	s.RegisterService(&FilesystemService_ServiceDesc, srv)
}

func _FilesystemService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Get(ctx, req.(*GetFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).GetByName(ctx, req.(*v1.GetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).List(ctx, req.(*ListFilesystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Create(ctx, req.(*CreateFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Update(ctx, req.(*UpdateFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).Delete(ctx, req.(*DeleteFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesystemService_ListOperationsByParent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsByParentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServiceServer).ListOperationsByParent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FilesystemService_ListOperationsByParent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServiceServer).ListOperationsByParent(ctx, req.(*ListOperationsByParentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FilesystemService_ServiceDesc is the grpc.ServiceDesc for FilesystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilesystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.compute.v1.FilesystemService",
	HandlerType: (*FilesystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FilesystemService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _FilesystemService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FilesystemService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FilesystemService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FilesystemService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FilesystemService_Delete_Handler,
		},
		{
			MethodName: "ListOperationsByParent",
			Handler:    _FilesystemService_ListOperationsByParent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/compute/v1/filesystem_service.proto",
}
