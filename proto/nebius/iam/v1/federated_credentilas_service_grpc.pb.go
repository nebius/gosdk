// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/iam/v1/federated_credentilas_service.proto

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
	FederatedCredentialsService_Get_FullMethodName       = "/nebius.iam.v1.FederatedCredentialsService/Get"
	FederatedCredentialsService_GetByName_FullMethodName = "/nebius.iam.v1.FederatedCredentialsService/GetByName"
	FederatedCredentialsService_List_FullMethodName      = "/nebius.iam.v1.FederatedCredentialsService/List"
	FederatedCredentialsService_Create_FullMethodName    = "/nebius.iam.v1.FederatedCredentialsService/Create"
	FederatedCredentialsService_Update_FullMethodName    = "/nebius.iam.v1.FederatedCredentialsService/Update"
	FederatedCredentialsService_Delete_FullMethodName    = "/nebius.iam.v1.FederatedCredentialsService/Delete"
)

// FederatedCredentialsServiceClient is the client API for FederatedCredentialsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FederatedCredentialsServiceClient interface {
	Get(ctx context.Context, in *GetFederatedCredentialsRequest, opts ...grpc.CallOption) (*FederatedCredentials, error)
	GetByName(ctx context.Context, in *GetByNameFederatedCredentialsRequest, opts ...grpc.CallOption) (*FederatedCredentials, error)
	List(ctx context.Context, in *ListFederatedCredentialsRequest, opts ...grpc.CallOption) (*ListFederatedCredentialsResponse, error)
	Create(ctx context.Context, in *CreateFederatedCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateFederatedCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteFederatedCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type federatedCredentialsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFederatedCredentialsServiceClient(cc grpc.ClientConnInterface) FederatedCredentialsServiceClient {
	return &federatedCredentialsServiceClient{cc}
}

func (c *federatedCredentialsServiceClient) Get(ctx context.Context, in *GetFederatedCredentialsRequest, opts ...grpc.CallOption) (*FederatedCredentials, error) {
	out := new(FederatedCredentials)
	err := c.cc.Invoke(ctx, FederatedCredentialsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialsServiceClient) GetByName(ctx context.Context, in *GetByNameFederatedCredentialsRequest, opts ...grpc.CallOption) (*FederatedCredentials, error) {
	out := new(FederatedCredentials)
	err := c.cc.Invoke(ctx, FederatedCredentialsService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialsServiceClient) List(ctx context.Context, in *ListFederatedCredentialsRequest, opts ...grpc.CallOption) (*ListFederatedCredentialsResponse, error) {
	out := new(ListFederatedCredentialsResponse)
	err := c.cc.Invoke(ctx, FederatedCredentialsService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialsServiceClient) Create(ctx context.Context, in *CreateFederatedCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederatedCredentialsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialsServiceClient) Update(ctx context.Context, in *UpdateFederatedCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederatedCredentialsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialsServiceClient) Delete(ctx context.Context, in *DeleteFederatedCredentialsRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, FederatedCredentialsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederatedCredentialsServiceServer is the server API for FederatedCredentialsService service.
// All implementations should embed UnimplementedFederatedCredentialsServiceServer
// for forward compatibility
type FederatedCredentialsServiceServer interface {
	Get(context.Context, *GetFederatedCredentialsRequest) (*FederatedCredentials, error)
	GetByName(context.Context, *GetByNameFederatedCredentialsRequest) (*FederatedCredentials, error)
	List(context.Context, *ListFederatedCredentialsRequest) (*ListFederatedCredentialsResponse, error)
	Create(context.Context, *CreateFederatedCredentialsRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateFederatedCredentialsRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteFederatedCredentialsRequest) (*v1.Operation, error)
}

// UnimplementedFederatedCredentialsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFederatedCredentialsServiceServer struct {
}

func (UnimplementedFederatedCredentialsServiceServer) Get(context.Context, *GetFederatedCredentialsRequest) (*FederatedCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFederatedCredentialsServiceServer) GetByName(context.Context, *GetByNameFederatedCredentialsRequest) (*FederatedCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedFederatedCredentialsServiceServer) List(context.Context, *ListFederatedCredentialsRequest) (*ListFederatedCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFederatedCredentialsServiceServer) Create(context.Context, *CreateFederatedCredentialsRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFederatedCredentialsServiceServer) Update(context.Context, *UpdateFederatedCredentialsRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFederatedCredentialsServiceServer) Delete(context.Context, *DeleteFederatedCredentialsRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeFederatedCredentialsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederatedCredentialsServiceServer will
// result in compilation errors.
type UnsafeFederatedCredentialsServiceServer interface {
	mustEmbedUnimplementedFederatedCredentialsServiceServer()
}

func RegisterFederatedCredentialsServiceServer(s grpc.ServiceRegistrar, srv FederatedCredentialsServiceServer) {
	s.RegisterService(&FederatedCredentialsService_ServiceDesc, srv)
}

func _FederatedCredentialsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialsServiceServer).Get(ctx, req.(*GetFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialsService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByNameFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialsServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialsService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialsServiceServer).GetByName(ctx, req.(*GetByNameFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialsService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialsServiceServer).List(ctx, req.(*ListFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialsServiceServer).Create(ctx, req.(*CreateFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialsServiceServer).Update(ctx, req.(*UpdateFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialsServiceServer).Delete(ctx, req.(*DeleteFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FederatedCredentialsService_ServiceDesc is the grpc.ServiceDesc for FederatedCredentialsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FederatedCredentialsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.FederatedCredentialsService",
	HandlerType: (*FederatedCredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FederatedCredentialsService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _FederatedCredentialsService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FederatedCredentialsService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FederatedCredentialsService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FederatedCredentialsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FederatedCredentialsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/federated_credentilas_service.proto",
}
