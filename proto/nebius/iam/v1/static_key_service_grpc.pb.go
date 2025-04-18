// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/iam/v1/static_key_service.proto

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
	StaticKeyService_Issue_FullMethodName     = "/nebius.iam.v1.StaticKeyService/Issue"
	StaticKeyService_List_FullMethodName      = "/nebius.iam.v1.StaticKeyService/List"
	StaticKeyService_Get_FullMethodName       = "/nebius.iam.v1.StaticKeyService/Get"
	StaticKeyService_GetByName_FullMethodName = "/nebius.iam.v1.StaticKeyService/GetByName"
	StaticKeyService_Delete_FullMethodName    = "/nebius.iam.v1.StaticKeyService/Delete"
	StaticKeyService_Find_FullMethodName      = "/nebius.iam.v1.StaticKeyService/Find"
	StaticKeyService_Revoke_FullMethodName    = "/nebius.iam.v1.StaticKeyService/Revoke"
)

// StaticKeyServiceClient is the client API for StaticKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticKeyServiceClient interface {
	Issue(ctx context.Context, in *IssueStaticKeyRequest, opts ...grpc.CallOption) (*IssueStaticKeyResponse, error)
	List(ctx context.Context, in *ListStaticKeysRequest, opts ...grpc.CallOption) (*ListStaticKeysResponse, error)
	Get(ctx context.Context, in *GetStaticKeyRequest, opts ...grpc.CallOption) (*StaticKey, error)
	GetByName(ctx context.Context, in *GetStaticKeyByNameRequest, opts ...grpc.CallOption) (*StaticKey, error)
	Delete(ctx context.Context, in *DeleteStaticKeyRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Find(ctx context.Context, in *FindStaticKeyRequest, opts ...grpc.CallOption) (*FindStaticKeyResponse, error)
	Revoke(ctx context.Context, in *RevokeStaticKeyRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type staticKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticKeyServiceClient(cc grpc.ClientConnInterface) StaticKeyServiceClient {
	return &staticKeyServiceClient{cc}
}

func (c *staticKeyServiceClient) Issue(ctx context.Context, in *IssueStaticKeyRequest, opts ...grpc.CallOption) (*IssueStaticKeyResponse, error) {
	out := new(IssueStaticKeyResponse)
	err := c.cc.Invoke(ctx, StaticKeyService_Issue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticKeyServiceClient) List(ctx context.Context, in *ListStaticKeysRequest, opts ...grpc.CallOption) (*ListStaticKeysResponse, error) {
	out := new(ListStaticKeysResponse)
	err := c.cc.Invoke(ctx, StaticKeyService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticKeyServiceClient) Get(ctx context.Context, in *GetStaticKeyRequest, opts ...grpc.CallOption) (*StaticKey, error) {
	out := new(StaticKey)
	err := c.cc.Invoke(ctx, StaticKeyService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticKeyServiceClient) GetByName(ctx context.Context, in *GetStaticKeyByNameRequest, opts ...grpc.CallOption) (*StaticKey, error) {
	out := new(StaticKey)
	err := c.cc.Invoke(ctx, StaticKeyService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticKeyServiceClient) Delete(ctx context.Context, in *DeleteStaticKeyRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, StaticKeyService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticKeyServiceClient) Find(ctx context.Context, in *FindStaticKeyRequest, opts ...grpc.CallOption) (*FindStaticKeyResponse, error) {
	out := new(FindStaticKeyResponse)
	err := c.cc.Invoke(ctx, StaticKeyService_Find_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staticKeyServiceClient) Revoke(ctx context.Context, in *RevokeStaticKeyRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, StaticKeyService_Revoke_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticKeyServiceServer is the server API for StaticKeyService service.
// All implementations should embed UnimplementedStaticKeyServiceServer
// for forward compatibility
type StaticKeyServiceServer interface {
	Issue(context.Context, *IssueStaticKeyRequest) (*IssueStaticKeyResponse, error)
	List(context.Context, *ListStaticKeysRequest) (*ListStaticKeysResponse, error)
	Get(context.Context, *GetStaticKeyRequest) (*StaticKey, error)
	GetByName(context.Context, *GetStaticKeyByNameRequest) (*StaticKey, error)
	Delete(context.Context, *DeleteStaticKeyRequest) (*v1.Operation, error)
	Find(context.Context, *FindStaticKeyRequest) (*FindStaticKeyResponse, error)
	Revoke(context.Context, *RevokeStaticKeyRequest) (*v1.Operation, error)
}

// UnimplementedStaticKeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStaticKeyServiceServer struct {
}

func (UnimplementedStaticKeyServiceServer) Issue(context.Context, *IssueStaticKeyRequest) (*IssueStaticKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedStaticKeyServiceServer) List(context.Context, *ListStaticKeysRequest) (*ListStaticKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStaticKeyServiceServer) Get(context.Context, *GetStaticKeyRequest) (*StaticKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStaticKeyServiceServer) GetByName(context.Context, *GetStaticKeyByNameRequest) (*StaticKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedStaticKeyServiceServer) Delete(context.Context, *DeleteStaticKeyRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaticKeyServiceServer) Find(context.Context, *FindStaticKeyRequest) (*FindStaticKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedStaticKeyServiceServer) Revoke(context.Context, *RevokeStaticKeyRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}

// UnsafeStaticKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticKeyServiceServer will
// result in compilation errors.
type UnsafeStaticKeyServiceServer interface {
	mustEmbedUnimplementedStaticKeyServiceServer()
}

func RegisterStaticKeyServiceServer(s grpc.ServiceRegistrar, srv StaticKeyServiceServer) {
	s.RegisterService(&StaticKeyService_ServiceDesc, srv)
}

func _StaticKeyService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueStaticKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_Issue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).Issue(ctx, req.(*IssueStaticKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaticKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).List(ctx, req.(*ListStaticKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).Get(ctx, req.(*GetStaticKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticKeyService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaticKeyByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).GetByName(ctx, req.(*GetStaticKeyByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaticKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).Delete(ctx, req.(*DeleteStaticKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticKeyService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStaticKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_Find_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).Find(ctx, req.(*FindStaticKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaticKeyService_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeStaticKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticKeyServiceServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaticKeyService_Revoke_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticKeyServiceServer).Revoke(ctx, req.(*RevokeStaticKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaticKeyService_ServiceDesc is the grpc.ServiceDesc for StaticKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaticKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.StaticKeyService",
	HandlerType: (*StaticKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Issue",
			Handler:    _StaticKeyService_Issue_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StaticKeyService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StaticKeyService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _StaticKeyService_GetByName_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaticKeyService_Delete_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _StaticKeyService_Find_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _StaticKeyService_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/static_key_service.proto",
}
