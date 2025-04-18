// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: nebius/storage/v1/bucket_service.proto

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
	BucketService_Get_FullMethodName       = "/nebius.storage.v1.BucketService/Get"
	BucketService_GetByName_FullMethodName = "/nebius.storage.v1.BucketService/GetByName"
	BucketService_List_FullMethodName      = "/nebius.storage.v1.BucketService/List"
	BucketService_Create_FullMethodName    = "/nebius.storage.v1.BucketService/Create"
	BucketService_Update_FullMethodName    = "/nebius.storage.v1.BucketService/Update"
	BucketService_Delete_FullMethodName    = "/nebius.storage.v1.BucketService/Delete"
	BucketService_Purge_FullMethodName     = "/nebius.storage.v1.BucketService/Purge"
	BucketService_Undelete_FullMethodName  = "/nebius.storage.v1.BucketService/Undelete"
)

// BucketServiceClient is the client API for BucketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BucketServiceClient interface {
	Get(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
	GetByName(ctx context.Context, in *GetBucketByNameRequest, opts ...grpc.CallOption) (*Bucket, error)
	List(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	Create(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Update(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	Delete(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Purge instantly deletes the bucket in ScheduledForDeletion state.
	// It can be used only for buckets in ScheduledForDeletion state.
	// If you want to delete Active bucket instantly, use Delete with zero ttl.
	Purge(ctx context.Context, in *PurgeBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error)
	// Undelete recovers the bucket from ScheduledForDeletion state to Active.
	Undelete(ctx context.Context, in *UndeleteBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error)
}

type bucketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBucketServiceClient(cc grpc.ClientConnInterface) BucketServiceClient {
	return &bucketServiceClient{cc}
}

func (c *bucketServiceClient) Get(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, BucketService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) GetByName(ctx context.Context, in *GetBucketByNameRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, BucketService_GetByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) List(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, BucketService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) Create(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, BucketService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) Update(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, BucketService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) Delete(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, BucketService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) Purge(ctx context.Context, in *PurgeBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, BucketService_Purge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) Undelete(ctx context.Context, in *UndeleteBucketRequest, opts ...grpc.CallOption) (*v1.Operation, error) {
	out := new(v1.Operation)
	err := c.cc.Invoke(ctx, BucketService_Undelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BucketServiceServer is the server API for BucketService service.
// All implementations should embed UnimplementedBucketServiceServer
// for forward compatibility
type BucketServiceServer interface {
	Get(context.Context, *GetBucketRequest) (*Bucket, error)
	GetByName(context.Context, *GetBucketByNameRequest) (*Bucket, error)
	List(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error)
	Create(context.Context, *CreateBucketRequest) (*v1.Operation, error)
	Update(context.Context, *UpdateBucketRequest) (*v1.Operation, error)
	Delete(context.Context, *DeleteBucketRequest) (*v1.Operation, error)
	// Purge instantly deletes the bucket in ScheduledForDeletion state.
	// It can be used only for buckets in ScheduledForDeletion state.
	// If you want to delete Active bucket instantly, use Delete with zero ttl.
	Purge(context.Context, *PurgeBucketRequest) (*v1.Operation, error)
	// Undelete recovers the bucket from ScheduledForDeletion state to Active.
	Undelete(context.Context, *UndeleteBucketRequest) (*v1.Operation, error)
}

// UnimplementedBucketServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBucketServiceServer struct {
}

func (UnimplementedBucketServiceServer) Get(context.Context, *GetBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBucketServiceServer) GetByName(context.Context, *GetBucketByNameRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedBucketServiceServer) List(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBucketServiceServer) Create(context.Context, *CreateBucketRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBucketServiceServer) Update(context.Context, *UpdateBucketRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBucketServiceServer) Delete(context.Context, *DeleteBucketRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBucketServiceServer) Purge(context.Context, *PurgeBucketRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purge not implemented")
}
func (UnimplementedBucketServiceServer) Undelete(context.Context, *UndeleteBucketRequest) (*v1.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Undelete not implemented")
}

// UnsafeBucketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BucketServiceServer will
// result in compilation errors.
type UnsafeBucketServiceServer interface {
	mustEmbedUnimplementedBucketServiceServer()
}

func RegisterBucketServiceServer(s grpc.ServiceRegistrar, srv BucketServiceServer) {
	s.RegisterService(&BucketService_ServiceDesc, srv)
}

func _BucketService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Get(ctx, req.(*GetBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).GetByName(ctx, req.(*GetBucketByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).List(ctx, req.(*ListBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Create(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Update(ctx, req.(*UpdateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Delete(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_Purge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Purge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_Purge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Purge(ctx, req.(*PurgeBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_Undelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).Undelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BucketService_Undelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).Undelete(ctx, req.(*UndeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BucketService_ServiceDesc is the grpc.ServiceDesc for BucketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BucketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nebius.storage.v1.BucketService",
	HandlerType: (*BucketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BucketService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _BucketService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BucketService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BucketService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BucketService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BucketService_Delete_Handler,
		},
		{
			MethodName: "Purge",
			Handler:    _BucketService_Purge_Handler,
		},
		{
			MethodName: "Undelete",
			Handler:    _BucketService_Undelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebius/storage/v1/bucket_service.proto",
}
