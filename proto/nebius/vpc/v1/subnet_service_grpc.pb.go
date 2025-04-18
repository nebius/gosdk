// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/vpc/v1/subnet_service.proto

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
	SubnetService_Get_FullMethodName           = "/nebius.vpc.v1.SubnetService/Get"
	SubnetService_GetByName_FullMethodName     = "/nebius.vpc.v1.SubnetService/GetByName"
	SubnetService_List_FullMethodName          = "/nebius.vpc.v1.SubnetService/List"
	SubnetService_ListByNetwork_FullMethodName = "/nebius.vpc.v1.SubnetService/ListByNetwork"
	SubnetService_Create_FullMethodName        = "/nebius.vpc.v1.SubnetService/Create"
	SubnetService_Update_FullMethodName        = "/nebius.vpc.v1.SubnetService/Update"
	SubnetService_Delete_FullMethodName        = "/nebius.vpc.v1.SubnetService/Delete"
)

// SubnetServiceClient is the client API for SubnetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubnetServiceClient interface {
	Get(ctx context.Context, in *GetSubnetRequest, opts ...grpc.CallOption) (*Subnet, error)
	GetByName(ctx context.Context, in *GetSubnetByNameRequest, opts ...grpc.CallOption) (*Subnet, error)
	List(ctx context.Context, in *ListSubnetsRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error)
	ListByNetwork(ctx context.Context, in *ListSubnetsByNetworkRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error)
	Create(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateSubnetRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteSubnetRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type subnetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubnetServiceClient(cc grpc.ClientConnInterface) SubnetServiceClient {
	return &subnetServiceClient{cc}
}

func (c *subnetServiceClient) Get(ctx context.Context, in *GetSubnetRequest, opts ...grpc.CallOption) (*Subnet, error) {
	out := new(Subnet)
	err := c.cc.Invoke(ctx, SubnetService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetServiceClient) GetByName(ctx context.Context, in *GetSubnetByNameRequest, opts ...grpc.CallOption) (*Subnet, error) {
	out := new(Subnet)
	err := c.cc.Invoke(ctx, SubnetService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetServiceClient) List(ctx context.Context, in *ListSubnetsRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error) {
	out := new(ListSubnetsResponse)
	err := c.cc.Invoke(ctx, SubnetService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetServiceClient) ListByNetwork(ctx context.Context, in *ListSubnetsByNetworkRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error) {
	out := new(ListSubnetsResponse)
	err := c.cc.Invoke(ctx, SubnetService_ListByNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetServiceClient) Create(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, SubnetService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetServiceClient) Update(ctx context.Context, in *UpdateSubnetRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, SubnetService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subnetServiceClient) Delete(ctx context.Context, in *DeleteSubnetRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, SubnetService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubnetServiceServer is the server API for SubnetService service.
// All implementations should embed UnimplementedSubnetServiceServer
// for forward compatibility
type SubnetServiceServer interface {
	Get(context.Context, *GetSubnetRequest) (*Subnet, error)
	GetByName(context.Context, *GetSubnetByNameRequest) (*Subnet, error)
	List(context.Context, *ListSubnetsRequest) (*ListSubnetsResponse, error)
	ListByNetwork(context.Context, *ListSubnetsByNetworkRequest) (*ListSubnetsResponse, error)
	Create(context.Context, *CreateSubnetRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateSubnetRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteSubnetRequest) (*v1.Operation, error)
}

// UnimplementedSubnetServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSubnetServiceServer struct {
}

func (UnimplementedSubnetServiceServer) Get(context.Context, *GetSubnetRequest) (*Subnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSubnetServiceServer) GetByName(context.Context, *GetSubnetByNameRequest) (*Subnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedSubnetServiceServer) List(context.Context, *ListSubnetsRequest) (*ListSubnetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSubnetServiceServer) ListByNetwork(context.Context, *ListSubnetsByNetworkRequest) (*ListSubnetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByNetwork not implemented")
}
func (UnimplementedSubnetServiceServer) Create(context.Context, *CreateSubnetRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSubnetServiceServer) Update(context.Context, *UpdateSubnetRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSubnetServiceServer) Delete(context.Context, *DeleteSubnetRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeSubnetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubnetServiceServer will
// result in compilation errors.
type UnsafeSubnetServiceServer interface {
	mustEmbedUnimplementedSubnetServiceServer()
}

func RegisterSubnetServiceServer(s grpc.ServiceRegistrar, srv SubnetServiceServer) {
	s.RegisterService(&SubnetService_ServiceDesc, srv)
}

func _SubnetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).Get(ctx, req.(*GetSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubnetService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubnetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).GetByName(ctx, req.(*GetSubnetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubnetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubnetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).List(ctx, req.(*ListSubnetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubnetService_ListByNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubnetsByNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).ListByNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_ListByNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).ListByNetwork(ctx, req.(*ListSubnetsByNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubnetService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).Create(ctx, req.(*CreateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubnetService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).Update(ctx, req.(*UpdateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubnetService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubnetService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetServiceServer).Delete(ctx, req.(*DeleteSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubnetService_ServiceDesc is the grpc.ServiceDesc for SubnetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubnetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.vpc.v1.SubnetService",
	HandlerType: (*SubnetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SubnetService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _SubnetService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SubnetService_List_Handler,
		},
		{
			MethodName: "ListByNetwork",
			Handler:    _SubnetService_ListByNetwork_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SubnetService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SubnetService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SubnetService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/vpc/v1/subnet_service.proto",
}
