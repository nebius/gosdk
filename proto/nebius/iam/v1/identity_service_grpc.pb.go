// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: nebius/iam/v1/identity_service.proto

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
	IdentityService_ExchangeToken_FullMethodName = "/nebius.iam.v1.IdentityService/ExchangeToken"
)

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type IdentityServiceClient interface {
	ExchangeToken(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) ExchangeToken(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, IdentityService_ExchangeToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations should embed UnimplementedIdentityServiceServer
// for forward compatibility
//
// Deprecated: Do not use.
type IdentityServiceServer interface {
	ExchangeToken(context.Context, *ExchangeTokenRequest) (*CreateTokenResponse, error)
}

// UnimplementedIdentityServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) ExchangeToken(context.Context, *ExchangeTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeToken not implemented")
}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

// Deprecated: Do not use.
func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_ExchangeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).ExchangeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityService_ExchangeToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).ExchangeToken(ctx, req.(*ExchangeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExchangeToken",
			Handler:    _IdentityService_ExchangeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/identity_service.proto",
}
