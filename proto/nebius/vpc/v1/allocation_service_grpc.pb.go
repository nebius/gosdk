// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/vpc/v1/allocation_service.proto

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
	AllocationService_Get_FullMethodName        = "/nebius.vpc.v1.AllocationService/Get"
	AllocationService_GetByName_FullMethodName  = "/nebius.vpc.v1.AllocationService/GetByName"
	AllocationService_List_FullMethodName       = "/nebius.vpc.v1.AllocationService/List"
	AllocationService_ListByPool_FullMethodName = "/nebius.vpc.v1.AllocationService/ListByPool"
	AllocationService_Create_FullMethodName     = "/nebius.vpc.v1.AllocationService/Create"
	AllocationService_Update_FullMethodName     = "/nebius.vpc.v1.AllocationService/Update"
	AllocationService_Delete_FullMethodName     = "/nebius.vpc.v1.AllocationService/Delete"
)

// AllocationServiceClient is the client API for AllocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AllocationServiceClient interface {
	Get(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*Allocation, error)
	GetByName(ctx context.Context, in *GetAllocationByNameRequest, opts ...grpc.CallOption) (*Allocation, error)
	List(ctx context.Context, in *ListAllocationsRequest, opts ...grpc.CallOption) (*ListAllocationsResponse, error)
	ListByPool(ctx context.Context, in *ListAllocationsByPoolRequest, opts ...grpc.CallOption) (*ListAllocationsResponse, error)
	Create(ctx context.Context, in *CreateAllocationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateAllocationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteAllocationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type allocationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAllocationServiceClient(cc grpc.ClientConnInterface) AllocationServiceClient {
	return &allocationServiceClient{cc}
}

func (c *allocationServiceClient) Get(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*Allocation, error) {
	out := new(Allocation)
	err := c.cc.Invoke(ctx, AllocationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationServiceClient) GetByName(ctx context.Context, in *GetAllocationByNameRequest, opts ...grpc.CallOption) (*Allocation, error) {
	out := new(Allocation)
	err := c.cc.Invoke(ctx, AllocationService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationServiceClient) List(ctx context.Context, in *ListAllocationsRequest, opts ...grpc.CallOption) (*ListAllocationsResponse, error) {
	out := new(ListAllocationsResponse)
	err := c.cc.Invoke(ctx, AllocationService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationServiceClient) ListByPool(ctx context.Context, in *ListAllocationsByPoolRequest, opts ...grpc.CallOption) (*ListAllocationsResponse, error) {
	out := new(ListAllocationsResponse)
	err := c.cc.Invoke(ctx, AllocationService_ListByPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationServiceClient) Create(ctx context.Context, in *CreateAllocationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, AllocationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationServiceClient) Update(ctx context.Context, in *UpdateAllocationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, AllocationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocationServiceClient) Delete(ctx context.Context, in *DeleteAllocationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, AllocationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllocationServiceServer is the server API for AllocationService service.
// All implementations should embed UnimplementedAllocationServiceServer
// for forward compatibility
type AllocationServiceServer interface {
	Get(context.Context, *GetAllocationRequest) (*Allocation, error)
	GetByName(context.Context, *GetAllocationByNameRequest) (*Allocation, error)
	List(context.Context, *ListAllocationsRequest) (*ListAllocationsResponse, error)
	ListByPool(context.Context, *ListAllocationsByPoolRequest) (*ListAllocationsResponse, error)
	Create(context.Context, *CreateAllocationRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateAllocationRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteAllocationRequest) (*v1.Operation, error)
}

// UnimplementedAllocationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAllocationServiceServer struct {
}

func (UnimplementedAllocationServiceServer) Get(context.Context, *GetAllocationRequest) (*Allocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAllocationServiceServer) GetByName(context.Context, *GetAllocationByNameRequest) (*Allocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedAllocationServiceServer) List(context.Context, *ListAllocationsRequest) (*ListAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAllocationServiceServer) ListByPool(context.Context, *ListAllocationsByPoolRequest) (*ListAllocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByPool not implemented")
}
func (UnimplementedAllocationServiceServer) Create(context.Context, *CreateAllocationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAllocationServiceServer) Update(context.Context, *UpdateAllocationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAllocationServiceServer) Delete(context.Context, *DeleteAllocationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeAllocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AllocationServiceServer will
// result in compilation errors.
type UnsafeAllocationServiceServer interface {
	mustEmbedUnimplementedAllocationServiceServer()
}

func RegisterAllocationServiceServer(s grpc.ServiceRegistrar, srv AllocationServiceServer) {
	s.RegisterService(&AllocationService_ServiceDesc, srv)
}

func _AllocationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).Get(ctx, req.(*GetAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllocationByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).GetByName(ctx, req.(*GetAllocationByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).List(ctx, req.(*ListAllocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationService_ListByPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllocationsByPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).ListByPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_ListByPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).ListByPool(ctx, req.(*ListAllocationsByPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).Create(ctx, req.(*CreateAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).Update(ctx, req.(*UpdateAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocationServiceServer).Delete(ctx, req.(*DeleteAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AllocationService_ServiceDesc is the grpc.ServiceDesc for AllocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AllocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.vpc.v1.AllocationService",
	HandlerType: (*AllocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AllocationService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _AllocationService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AllocationService_List_Handler,
		},
		{
			MethodName: "ListByPool",
			Handler:    _AllocationService_ListByPool_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AllocationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AllocationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AllocationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/vpc/v1/allocation_service.proto",
}
