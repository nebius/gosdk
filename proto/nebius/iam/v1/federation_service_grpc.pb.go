// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: nebius/iam/v1/federation_service.proto

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
	FederationService_Create_FullMethodName    = "/nebius.iam.v1.FederationService/Create"
	FederationService_Get_FullMethodName       = "/nebius.iam.v1.FederationService/Get"
	FederationService_GetByName_FullMethodName = "/nebius.iam.v1.FederationService/GetByName"
	FederationService_List_FullMethodName      = "/nebius.iam.v1.FederationService/List"
	FederationService_Update_FullMethodName    = "/nebius.iam.v1.FederationService/Update"
	FederationService_Delete_FullMethodName    = "/nebius.iam.v1.FederationService/Delete"
)

// FederationServiceClient is the client API for FederationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FederationServiceClient interface {
	Create(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Get(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*Federation, error)
	GetByName(ctx context.Context, in *v1.GetByNameRequest, opts ...grpc.CallOption) (*Federation, error)
	List(ctx context.Context, in *ListFederationsRequest, opts ...grpc.CallOption) (*ListFederationsResponse, error)
	Update(ctx context.Context, in *UpdateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type federationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFederationServiceClient(cc grpc.ClientConnInterface) FederationServiceClient {
	return &federationServiceClient{cc}
}

func (c *federationServiceClient) Create(ctx context.Context, in *CreateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Get(ctx context.Context, in *GetFederationRequest, opts ...grpc.CallOption) (*Federation, error) {
	out := new(Federation)
	err := c.cc.Invoke(ctx, FederationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) GetByName(ctx context.Context, in *v1.GetByNameRequest, opts ...grpc.CallOption) (*Federation, error) {
	out := new(Federation)
	err := c.cc.Invoke(ctx, FederationService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) List(ctx context.Context, in *ListFederationsRequest, opts ...grpc.CallOption) (*ListFederationsResponse, error) {
	out := new(ListFederationsResponse)
	err := c.cc.Invoke(ctx, FederationService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Update(ctx context.Context, in *UpdateFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federationServiceClient) Delete(ctx context.Context, in *DeleteFederationRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederationServiceServer is the server API for FederationService service.
// All implementations should embed UnimplementedFederationServiceServer
// for forward compatibility
type FederationServiceServer interface {
	Create(context.Context, *CreateFederationRequest) (*v1.Operation, error)
	Get(context.Context, *GetFederationRequest) (*Federation, error)
	GetByName(context.Context, *v1.GetByNameRequest) (*Federation, error)
	List(context.Context, *ListFederationsRequest) (*ListFederationsResponse, error)
	Update(context.Context, *UpdateFederationRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteFederationRequest) (*v1.Operation, error)
}

// UnimplementedFederationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFederationServiceServer struct {
}

func (UnimplementedFederationServiceServer) Create(context.Context, *CreateFederationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFederationServiceServer) Get(context.Context, *GetFederationRequest) (*Federation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFederationServiceServer) GetByName(context.Context, *v1.GetByNameRequest) (*Federation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedFederationServiceServer) List(context.Context, *ListFederationsRequest) (*ListFederationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFederationServiceServer) Update(context.Context, *UpdateFederationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFederationServiceServer) Delete(context.Context, *DeleteFederationRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeFederationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederationServiceServer will
// result in compilation errors.
type UnsafeFederationServiceServer interface {
	mustEmbedUnimplementedFederationServiceServer()
}

func RegisterFederationServiceServer(s grpc.ServiceRegistrar, srv FederationServiceServer) {
	s.RegisterService(&FederationService_ServiceDesc, srv)
}

func _FederationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Create(ctx, req.(*CreateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Get(ctx, req.(*GetFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).GetByName(ctx, req.(*v1.GetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).List(ctx, req.(*ListFederationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Update(ctx, req.(*UpdateFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederationServiceServer).Delete(ctx, req.(*DeleteFederationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FederationService_ServiceDesc is the grpc.ServiceDesc for FederationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FederationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.FederationService",
	HandlerType: (*FederationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FederationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FederationService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _FederationService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FederationService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FederationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FederationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/federation_service.proto",
}
