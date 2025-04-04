// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/iam/v1/token_exchange_service.proto

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
	TokenExchangeService_Exchange_FullMethodName = "/nebius.iam.v1.TokenExchangeService/Exchange"
)

// TokenExchangeServiceClient is the client API for TokenExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenExchangeServiceClient interface {
	Exchange(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
}

type tokenExchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenExchangeServiceClient(cc grpc.ClientConnInterface) TokenExchangeServiceClient {
	return &tokenExchangeServiceClient{cc}
}

func (c *tokenExchangeServiceClient) Exchange(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, TokenExchangeService_Exchange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenExchangeServiceServer is the server API for TokenExchangeService service.
// All implementations should embed UnimplementedTokenExchangeServiceServer
// for forward compatibility
type TokenExchangeServiceServer interface {
	Exchange(context.Context, *ExchangeTokenRequest) (*CreateTokenResponse, error)
}

// UnimplementedTokenExchangeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTokenExchangeServiceServer struct {
}

func (UnimplementedTokenExchangeServiceServer) Exchange(context.Context, *ExchangeTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exchange not implemented")
}

// UnsafeTokenExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenExchangeServiceServer will
// result in compilation errors.
type UnsafeTokenExchangeServiceServer interface {
	mustEmbedUnimplementedTokenExchangeServiceServer()
}

func RegisterTokenExchangeServiceServer(s grpc.ServiceRegistrar, srv TokenExchangeServiceServer) {
	s.RegisterService(&TokenExchangeService_ServiceDesc, srv)
}

func _TokenExchangeService_Exchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenExchangeServiceServer).Exchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenExchangeService_Exchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenExchangeServiceServer).Exchange(ctx, req.(*ExchangeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenExchangeService_ServiceDesc is the grpc.ServiceDesc for TokenExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.iam.v1.TokenExchangeService",
	HandlerType: (*TokenExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exchange",
			Handler:    _TokenExchangeService_Exchange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/iam/v1/token_exchange_service.proto",
}
