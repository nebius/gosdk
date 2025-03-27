// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// nebius/compute/v1alpha1/image_service.proto is a deprecated file.

package v1alpha1

import (
	context "context"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ImageService_Get_FullMethodName                    = "/nebius.compute.v1alpha1.ImageService/Get"
	ImageService_GetByName_FullMethodName              = "/nebius.compute.v1alpha1.ImageService/GetByName"
	ImageService_GetLatestByFamily_FullMethodName      = "/nebius.compute.v1alpha1.ImageService/GetLatestByFamily"
	ImageService_List_FullMethodName                   = "/nebius.compute.v1alpha1.ImageService/List"
	ImageService_ListOperationsByParent_FullMethodName = "/nebius.compute.v1alpha1.ImageService/ListOperationsByParent"
)

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type ImageServiceClient interface {
	Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error)
	GetByName(ctx context.Context, in *v1.GetByNameRequest, opts ...grpc.CallOption) (*Image, error)
	GetLatestByFamily(ctx context.Context, in *GetImageLatestByFamilyRequest, opts ...grpc.CallOption) (*Image, error)
	List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	ListOperationsByParent(ctx context.Context, in *v1alpha1.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1alpha1.ListOperationsResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, ImageService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) GetByName(ctx context.Context, in *v1.GetByNameRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, ImageService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) GetLatestByFamily(ctx context.Context, in *GetImageLatestByFamilyRequest, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := c.cc.Invoke(ctx, ImageService_GetLatestByFamily_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, ImageService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ListOperationsByParent(ctx context.Context, in *v1alpha1.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1alpha1.ListOperationsResponse, error) {
	out := new(v1alpha1.ListOperationsResponse)
	err := c.cc.Invoke(ctx, ImageService_ListOperationsByParent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations should embed UnimplementedImageServiceServer
// for forward compatibility
//
// Deprecated: Do not use.
type ImageServiceServer interface {
	Get(context.Context, *GetImageRequest) (*Image, error)
	GetByName(context.Context, *v1.GetByNameRequest) (*Image, error)
	GetLatestByFamily(context.Context, *GetImageLatestByFamilyRequest) (*Image, error)
	List(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	ListOperationsByParent(context.Context, *v1alpha1.ListOperationsByParentRequest) (*v1alpha1.ListOperationsResponse, error)
}

// UnimplementedImageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) Get(context.Context, *GetImageRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedImageServiceServer) GetByName(context.Context, *v1.GetByNameRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedImageServiceServer) GetLatestByFamily(context.Context, *GetImageLatestByFamilyRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestByFamily not implemented")
}
func (UnimplementedImageServiceServer) List(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedImageServiceServer) ListOperationsByParent(context.Context, *v1alpha1.ListOperationsByParentRequest) (*v1alpha1.ListOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperationsByParent not implemented")
}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

// Deprecated: Do not use.
func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Get(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetByName(ctx, req.(*v1.GetByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_GetLatestByFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageLatestByFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).GetLatestByFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_GetLatestByFamily_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).GetLatestByFamily(ctx, req.(*GetImageLatestByFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).List(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ListOperationsByParent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha1.ListOperationsByParentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ListOperationsByParent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_ListOperationsByParent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ListOperationsByParent(ctx, req.(*v1alpha1.ListOperationsByParentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.compute.v1alpha1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ImageService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _ImageService_GetByName_Handler,
		},
		{
			MethodName: "GetLatestByFamily",
			Handler:    _ImageService_GetLatestByFamily_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ImageService_List_Handler,
		},
		{
			MethodName: "ListOperationsByParent",
			Handler:    _ImageService_ListOperationsByParent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/compute/v1alpha1/image_service.proto",
}
