// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/iam/v1/tenant_user_account_service.proto

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
	TenantUserAccountService_Get_FullMethodName     = "/nebius.iam.v1.TenantUserAccountService/Get"
	TenantUserAccountService_List_FullMethodName    = "/nebius.iam.v1.TenantUserAccountService/List"
	TenantUserAccountService_Block_FullMethodName   = "/nebius.iam.v1.TenantUserAccountService/Block"
	TenantUserAccountService_Unblock_FullMethodName = "/nebius.iam.v1.TenantUserAccountService/Unblock"
)

// TenantUserAccountServiceClient is the client API for TenantUserAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantUserAccountServiceClient interface {
	Get(ctx context.Context, in *GetTenantUserAccountRequest, opts ...grpc.CallOption) (*TenantUserAccount, error)
	List(ctx context.Context, in *ListTenantUserAccountsRequest, opts ...grpc.CallOption) (*ListTenantUserAccountsResponse, error)
	Block(ctx context.Context, in *BlockTenantUserAccountRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Unblock(ctx context.Context, in *UnblockTenantUserAccountRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type tenantUserAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantUserAccountServiceClient(cc grpc.ClientConnInterface) TenantUserAccountServiceClient {
	return &tenantUserAccountServiceClient{cc}
}

func (c *tenantUserAccountServiceClient) Get(ctx context.Context, in *GetTenantUserAccountRequest, opts ...grpc.CallOption) (*TenantUserAccount, error) {
	out := new(TenantUserAccount)
	err := c.cc.Invoke(ctx, TenantUserAccountService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantUserAccountServiceClient) List(ctx context.Context, in *ListTenantUserAccountsRequest, opts ...grpc.CallOption) (*ListTenantUserAccountsResponse, error) {
	out := new(ListTenantUserAccountsResponse)
	err := c.cc.Invoke(ctx, TenantUserAccountService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantUserAccountServiceClient) Block(ctx context.Context, in *BlockTenantUserAccountRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, TenantUserAccountService_Block_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantUserAccountServiceClient) Unblock(ctx context.Context, in *UnblockTenantUserAccountRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, TenantUserAccountService_Unblock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantUserAccountServiceServer is the server API for TenantUserAccountService service.
// All implementations should embed UnimplementedTenantUserAccountServiceServer
// for forward compatibility
type TenantUserAccountServiceServer interface {
	Get(context.Context, *GetTenantUserAccountRequest) (*TenantUserAccount, error)
	List(context.Context, *ListTenantUserAccountsRequest) (*ListTenantUserAccountsResponse, error)
	Block(context.Context, *BlockTenantUserAccountRequest) (*v1.Operation, error)
	Unblock(context.Context, *UnblockTenantUserAccountRequest) (*v1.Operation, error)
}

// UnimplementedTenantUserAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantUserAccountServiceServer struct {
}

func (UnimplementedTenantUserAccountServiceServer) Get(context.Context, *GetTenantUserAccountRequest) (*TenantUserAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTenantUserAccountServiceServer) List(context.Context, *ListTenantUserAccountsRequest) (*ListTenantUserAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTenantUserAccountServiceServer) Block(context.Context, *BlockTenantUserAccountRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (UnimplementedTenantUserAccountServiceServer) Unblock(context.Context, *UnblockTenantUserAccountRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unblock not implemented")
}

// UnsafeTenantUserAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantUserAccountServiceServer will
// result in compilation errors.
type UnsafeTenantUserAccountServiceServer interface {
	mustEmbedUnimplementedTenantUserAccountServiceServer()
}

func RegisterTenantUserAccountServiceServer(s grpc.ServiceRegistrar, srv TenantUserAccountServiceServer) {
	s.RegisterService(&TenantUserAccountService_ServiceDesc, srv)
}

func _TenantUserAccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantUserAccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantUserAccountService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantUserAccountServiceServer).Get(ctx, req.(*GetTenantUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantUserAccountService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantUserAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantUserAccountServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantUserAccountService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantUserAccountServiceServer).List(ctx, req.(*ListTenantUserAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantUserAccountService_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockTenantUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantUserAccountServiceServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantUserAccountService_Block_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantUserAccountServiceServer).Block(ctx, req.(*BlockTenantUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantUserAccountService_Unblock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnblockTenantUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantUserAccountServiceServer).Unblock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantUserAccountService_Unblock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantUserAccountServiceServer).Unblock(ctx, req.(*UnblockTenantUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantUserAccountService_ServiceDesc is the grpc.ServiceDesc for TenantUserAccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantUserAccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.TenantUserAccountService",
	HandlerType: (*TenantUserAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TenantUserAccountService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TenantUserAccountService_List_Handler,
		},
		{
			MethodName: "Block",
			Handler:    _TenantUserAccountService_Block_Handler,
		},
		{
			MethodName: "Unblock",
			Handler:    _TenantUserAccountService_Unblock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/tenant_user_account_service.proto",
}
