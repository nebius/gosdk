// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/vpc/v1alpha1/pool_service.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPoolRequest) Reset() {
	*x = GetPoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoolRequest) ProtoMessage() {}

func (x *GetPoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoolRequest.ProtoReflect.Descriptor instead.
func (*GetPoolRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPoolRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPoolByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPoolByNameRequest) Reset() {
	*x = GetPoolByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoolByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoolByNameRequest) ProtoMessage() {}

func (x *GetPoolByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoolByNameRequest.ProtoReflect.Descriptor instead.
func (*GetPoolByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPoolByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetPoolByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListPoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId  string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PageSize  int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListPoolsRequest) Reset() {
	*x = ListPoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoolsRequest) ProtoMessage() {}

func (x *ListPoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoolsRequest.ProtoReflect.Descriptor instead.
func (*ListPoolsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPoolsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListPoolsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPoolsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListPoolsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListPoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items         []*Pool `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPoolsResponse) Reset() {
	*x = ListPoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPoolsResponse) ProtoMessage() {}

func (x *ListPoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPoolsResponse.ProtoReflect.Descriptor instead.
func (*ListPoolsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListPoolsResponse) GetItems() []*Pool {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPoolsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_vpc_v1alpha1_pool_service_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1alpha1_pool_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8b, 0x01,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x6c, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xfe, 0x01, 0x0a, 0x0b, 0x50, 0x6f,
	0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c,
	0x12, 0x51, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x6f, 0x6f, 0x6c, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x63, 0x0a, 0x1a, 0x61, 0x69,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_vpc_v1alpha1_pool_service_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1alpha1_pool_service_proto_rawDescData = file_nebius_vpc_v1alpha1_pool_service_proto_rawDesc
)

func file_nebius_vpc_v1alpha1_pool_service_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1alpha1_pool_service_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1alpha1_pool_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_vpc_v1alpha1_pool_service_proto_rawDescData)
	})
	return file_nebius_vpc_v1alpha1_pool_service_proto_rawDescData
}

var file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_vpc_v1alpha1_pool_service_proto_goTypes = []any{
	(*GetPoolRequest)(nil),       // 0: nebius.vpc.v1alpha1.GetPoolRequest
	(*GetPoolByNameRequest)(nil), // 1: nebius.vpc.v1alpha1.GetPoolByNameRequest
	(*ListPoolsRequest)(nil),     // 2: nebius.vpc.v1alpha1.ListPoolsRequest
	(*ListPoolsResponse)(nil),    // 3: nebius.vpc.v1alpha1.ListPoolsResponse
	(*Pool)(nil),                 // 4: nebius.vpc.v1alpha1.Pool
}
var file_nebius_vpc_v1alpha1_pool_service_proto_depIdxs = []int32{
	4, // 0: nebius.vpc.v1alpha1.ListPoolsResponse.items:type_name -> nebius.vpc.v1alpha1.Pool
	0, // 1: nebius.vpc.v1alpha1.PoolService.Get:input_type -> nebius.vpc.v1alpha1.GetPoolRequest
	1, // 2: nebius.vpc.v1alpha1.PoolService.GetByName:input_type -> nebius.vpc.v1alpha1.GetPoolByNameRequest
	2, // 3: nebius.vpc.v1alpha1.PoolService.List:input_type -> nebius.vpc.v1alpha1.ListPoolsRequest
	4, // 4: nebius.vpc.v1alpha1.PoolService.Get:output_type -> nebius.vpc.v1alpha1.Pool
	4, // 5: nebius.vpc.v1alpha1.PoolService.GetByName:output_type -> nebius.vpc.v1alpha1.Pool
	3, // 6: nebius.vpc.v1alpha1.PoolService.List:output_type -> nebius.vpc.v1alpha1.ListPoolsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1alpha1_pool_service_proto_init() }
func file_nebius_vpc_v1alpha1_pool_service_proto_init() {
	if File_nebius_vpc_v1alpha1_pool_service_proto != nil {
		return
	}
	file_nebius_vpc_v1alpha1_pool_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPoolRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetPoolByNameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListPoolsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListPoolsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_vpc_v1alpha1_pool_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_vpc_v1alpha1_pool_service_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1alpha1_pool_service_proto_depIdxs,
		MessageInfos:      file_nebius_vpc_v1alpha1_pool_service_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1alpha1_pool_service_proto = out.File
	file_nebius_vpc_v1alpha1_pool_service_proto_rawDesc = nil
	file_nebius_vpc_v1alpha1_pool_service_proto_goTypes = nil
	file_nebius_vpc_v1alpha1_pool_service_proto_depIdxs = nil
}
