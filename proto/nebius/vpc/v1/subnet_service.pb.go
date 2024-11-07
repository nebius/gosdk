// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/vpc/v1/subnet_service.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
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

type GetSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSubnetRequest) Reset() {
	*x = GetSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubnetRequest) ProtoMessage() {}

func (x *GetSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubnetRequest.ProtoReflect.Descriptor instead.
func (*GetSubnetRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSubnetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSubnetByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSubnetByNameRequest) Reset() {
	*x = GetSubnetByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubnetByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubnetByNameRequest) ProtoMessage() {}

func (x *GetSubnetByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubnetByNameRequest.ProtoReflect.Descriptor instead.
func (*GetSubnetByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSubnetByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetSubnetByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListSubnetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId  string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PageSize  int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSubnetsRequest) Reset() {
	*x = ListSubnetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubnetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubnetsRequest) ProtoMessage() {}

func (x *ListSubnetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubnetsRequest.ProtoReflect.Descriptor instead.
func (*ListSubnetsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSubnetsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListSubnetsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSubnetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListSubnetsByNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	PageSize  int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSubnetsByNetworkRequest) Reset() {
	*x = ListSubnetsByNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubnetsByNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubnetsByNetworkRequest) ProtoMessage() {}

func (x *ListSubnetsByNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubnetsByNetworkRequest.ProtoReflect.Descriptor instead.
func (*ListSubnetsByNetworkRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListSubnetsByNetworkRequest) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *ListSubnetsByNetworkRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSubnetsByNetworkRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListSubnetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items         []*Subnet `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string    `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSubnetsResponse) Reset() {
	*x = ListSubnetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubnetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubnetsResponse) ProtoMessage() {}

func (x *ListSubnetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubnetsResponse.ProtoReflect.Descriptor instead.
func (*ListSubnetsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListSubnetsResponse) GetItems() []*Subnet {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListSubnetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_vpc_v1_subnet_service_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1_subnet_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x59, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x75, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0xdb, 0x02, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12,
	0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x2a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x42, 0x79, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x10, 0x9a, 0xb5, 0x18, 0x0c, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x64, 0x42, 0x59, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_vpc_v1_subnet_service_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1_subnet_service_proto_rawDescData = file_nebius_vpc_v1_subnet_service_proto_rawDesc
)

func file_nebius_vpc_v1_subnet_service_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1_subnet_service_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1_subnet_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_vpc_v1_subnet_service_proto_rawDescData)
	})
	return file_nebius_vpc_v1_subnet_service_proto_rawDescData
}

var file_nebius_vpc_v1_subnet_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_vpc_v1_subnet_service_proto_goTypes = []any{
	(*GetSubnetRequest)(nil),            // 0: nebius.vpc.v1.GetSubnetRequest
	(*GetSubnetByNameRequest)(nil),      // 1: nebius.vpc.v1.GetSubnetByNameRequest
	(*ListSubnetsRequest)(nil),          // 2: nebius.vpc.v1.ListSubnetsRequest
	(*ListSubnetsByNetworkRequest)(nil), // 3: nebius.vpc.v1.ListSubnetsByNetworkRequest
	(*ListSubnetsResponse)(nil),         // 4: nebius.vpc.v1.ListSubnetsResponse
	(*Subnet)(nil),                      // 5: nebius.vpc.v1.Subnet
}
var file_nebius_vpc_v1_subnet_service_proto_depIdxs = []int32{
	5, // 0: nebius.vpc.v1.ListSubnetsResponse.items:type_name -> nebius.vpc.v1.Subnet
	0, // 1: nebius.vpc.v1.SubnetService.Get:input_type -> nebius.vpc.v1.GetSubnetRequest
	1, // 2: nebius.vpc.v1.SubnetService.GetByName:input_type -> nebius.vpc.v1.GetSubnetByNameRequest
	2, // 3: nebius.vpc.v1.SubnetService.List:input_type -> nebius.vpc.v1.ListSubnetsRequest
	3, // 4: nebius.vpc.v1.SubnetService.ListByNetwork:input_type -> nebius.vpc.v1.ListSubnetsByNetworkRequest
	5, // 5: nebius.vpc.v1.SubnetService.Get:output_type -> nebius.vpc.v1.Subnet
	5, // 6: nebius.vpc.v1.SubnetService.GetByName:output_type -> nebius.vpc.v1.Subnet
	4, // 7: nebius.vpc.v1.SubnetService.List:output_type -> nebius.vpc.v1.ListSubnetsResponse
	4, // 8: nebius.vpc.v1.SubnetService.ListByNetwork:output_type -> nebius.vpc.v1.ListSubnetsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1_subnet_service_proto_init() }
func file_nebius_vpc_v1_subnet_service_proto_init() {
	if File_nebius_vpc_v1_subnet_service_proto != nil {
		return
	}
	file_nebius_vpc_v1_subnet_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nebius_vpc_v1_subnet_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubnetRequest); i {
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
		file_nebius_vpc_v1_subnet_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubnetByNameRequest); i {
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
		file_nebius_vpc_v1_subnet_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubnetsRequest); i {
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
		file_nebius_vpc_v1_subnet_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubnetsByNetworkRequest); i {
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
		file_nebius_vpc_v1_subnet_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubnetsResponse); i {
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
			RawDescriptor: file_nebius_vpc_v1_subnet_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_vpc_v1_subnet_service_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1_subnet_service_proto_depIdxs,
		MessageInfos:      file_nebius_vpc_v1_subnet_service_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1_subnet_service_proto = out.File
	file_nebius_vpc_v1_subnet_service_proto_rawDesc = nil
	file_nebius_vpc_v1_subnet_service_proto_goTypes = nil
	file_nebius_vpc_v1_subnet_service_proto_depIdxs = nil
}
