// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/vpc/v1alpha1/allocation_service.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
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

type GetAllocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllocationRequest) Reset() {
	*x = GetAllocationRequest{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllocationRequest) ProtoMessage() {}

func (x *GetAllocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllocationRequest.ProtoReflect.Descriptor instead.
func (*GetAllocationRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllocationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllocationByNameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllocationByNameRequest) Reset() {
	*x = GetAllocationByNameRequest{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllocationByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllocationByNameRequest) ProtoMessage() {}

func (x *GetAllocationByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllocationByNameRequest.ProtoReflect.Descriptor instead.
func (*GetAllocationByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllocationByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetAllocationByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListAllocationsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter        string                 `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllocationsRequest) Reset() {
	*x = ListAllocationsRequest{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllocationsRequest) ProtoMessage() {}

func (x *ListAllocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllocationsRequest.ProtoReflect.Descriptor instead.
func (*ListAllocationsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListAllocationsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListAllocationsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAllocationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAllocationsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListAllocationsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Allocation          `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllocationsResponse) Reset() {
	*x = ListAllocationsResponse{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllocationsResponse) ProtoMessage() {}

func (x *ListAllocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllocationsResponse.ProtoReflect.Descriptor instead.
func (*ListAllocationsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAllocationsResponse) GetItems() []*Allocation {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListAllocationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateAllocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *AllocationSpec        `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAllocationRequest) Reset() {
	*x = CreateAllocationRequest{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAllocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAllocationRequest) ProtoMessage() {}

func (x *CreateAllocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAllocationRequest.ProtoReflect.Descriptor instead.
func (*CreateAllocationRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAllocationRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateAllocationRequest) GetSpec() *AllocationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type UpdateAllocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *AllocationSpec        `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAllocationRequest) Reset() {
	*x = UpdateAllocationRequest{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAllocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAllocationRequest) ProtoMessage() {}

func (x *UpdateAllocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAllocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateAllocationRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAllocationRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateAllocationRequest) GetSpec() *AllocationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteAllocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAllocationRequest) Reset() {
	*x = DeleteAllocationRequest{}
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAllocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllocationRequest) ProtoMessage() {}

func (x *DeleteAllocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllocationRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllocationRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAllocationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_nebius_vpc_v1alpha1_allocation_service_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1alpha1_allocation_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x91,
	0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x78, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xe7, 0x02, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x8a, 0x02, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0xc9, 0x01, 0xba, 0x48, 0xc5, 0x01, 0xba, 0x01, 0xbe, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x73, 0x27, 0x6e, 0x61, 0x6d, 0x65,
	0x27, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x77, 0x69, 0x74,
	0x68, 0x20, 0x61, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x20, 0x6f, 0x72, 0x20, 0x64, 0x69,
	0x67, 0x69, 0x74, 0x2c, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x27, 0x2d, 0x27, 0x2c, 0x20,
	0x27, 0x5f, 0x27, 0x2c, 0x20, 0x27, 0x2e, 0x27, 0x2c, 0x20, 0x27, 0x2f, 0x27, 0x2c, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x61, 0x20, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x20, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x20, 0x32, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x32,
	0x35, 0x35, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x1a, 0x38,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b,
	0x2d, 0x5f, 0x2e, 0x2f, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31,
	0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24, 0x27, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xdc, 0x02, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x87, 0x02, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xc6, 0x01, 0xba, 0x48, 0xc2, 0x01,
	0xba, 0x01, 0xbe, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x73, 0x27, 0x6e, 0x61, 0x6d, 0x65, 0x27, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x6c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x20, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x67, 0x69, 0x74, 0x2c, 0x20, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x27, 0x2d, 0x27, 0x2c, 0x20, 0x27, 0x5f, 0x27, 0x2c, 0x20, 0x27,
	0x2e, 0x27, 0x2c, 0x20, 0x27, 0x2f, 0x27, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x61, 0x76,
	0x65, 0x20, 0x61, 0x20, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x62, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x20, 0x32, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x32, 0x35, 0x35, 0x20, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x1a, 0x38, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x5f, 0x2e, 0x2f, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x35, 0x34, 0x7d, 0x24,
	0x27, 0x29, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x31, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb9, 0x04, 0x0a, 0x11, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x61, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x0a, 0x1a, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x16, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescData = file_nebius_vpc_v1alpha1_allocation_service_proto_rawDesc
)

func file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescData)
	})
	return file_nebius_vpc_v1alpha1_allocation_service_proto_rawDescData
}

var file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nebius_vpc_v1alpha1_allocation_service_proto_goTypes = []any{
	(*GetAllocationRequest)(nil),       // 0: nebius.vpc.v1alpha1.GetAllocationRequest
	(*GetAllocationByNameRequest)(nil), // 1: nebius.vpc.v1alpha1.GetAllocationByNameRequest
	(*ListAllocationsRequest)(nil),     // 2: nebius.vpc.v1alpha1.ListAllocationsRequest
	(*ListAllocationsResponse)(nil),    // 3: nebius.vpc.v1alpha1.ListAllocationsResponse
	(*CreateAllocationRequest)(nil),    // 4: nebius.vpc.v1alpha1.CreateAllocationRequest
	(*UpdateAllocationRequest)(nil),    // 5: nebius.vpc.v1alpha1.UpdateAllocationRequest
	(*DeleteAllocationRequest)(nil),    // 6: nebius.vpc.v1alpha1.DeleteAllocationRequest
	(*Allocation)(nil),                 // 7: nebius.vpc.v1alpha1.Allocation
	(*v1.ResourceMetadata)(nil),        // 8: nebius.common.v1.ResourceMetadata
	(*AllocationSpec)(nil),             // 9: nebius.vpc.v1alpha1.AllocationSpec
	(*v1alpha1.Operation)(nil),         // 10: nebius.common.v1alpha1.Operation
}
var file_nebius_vpc_v1alpha1_allocation_service_proto_depIdxs = []int32{
	7,  // 0: nebius.vpc.v1alpha1.ListAllocationsResponse.items:type_name -> nebius.vpc.v1alpha1.Allocation
	8,  // 1: nebius.vpc.v1alpha1.CreateAllocationRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	9,  // 2: nebius.vpc.v1alpha1.CreateAllocationRequest.spec:type_name -> nebius.vpc.v1alpha1.AllocationSpec
	8,  // 3: nebius.vpc.v1alpha1.UpdateAllocationRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	9,  // 4: nebius.vpc.v1alpha1.UpdateAllocationRequest.spec:type_name -> nebius.vpc.v1alpha1.AllocationSpec
	0,  // 5: nebius.vpc.v1alpha1.AllocationService.Get:input_type -> nebius.vpc.v1alpha1.GetAllocationRequest
	1,  // 6: nebius.vpc.v1alpha1.AllocationService.GetByName:input_type -> nebius.vpc.v1alpha1.GetAllocationByNameRequest
	2,  // 7: nebius.vpc.v1alpha1.AllocationService.List:input_type -> nebius.vpc.v1alpha1.ListAllocationsRequest
	4,  // 8: nebius.vpc.v1alpha1.AllocationService.Create:input_type -> nebius.vpc.v1alpha1.CreateAllocationRequest
	5,  // 9: nebius.vpc.v1alpha1.AllocationService.Update:input_type -> nebius.vpc.v1alpha1.UpdateAllocationRequest
	6,  // 10: nebius.vpc.v1alpha1.AllocationService.Delete:input_type -> nebius.vpc.v1alpha1.DeleteAllocationRequest
	7,  // 11: nebius.vpc.v1alpha1.AllocationService.Get:output_type -> nebius.vpc.v1alpha1.Allocation
	7,  // 12: nebius.vpc.v1alpha1.AllocationService.GetByName:output_type -> nebius.vpc.v1alpha1.Allocation
	3,  // 13: nebius.vpc.v1alpha1.AllocationService.List:output_type -> nebius.vpc.v1alpha1.ListAllocationsResponse
	10, // 14: nebius.vpc.v1alpha1.AllocationService.Create:output_type -> nebius.common.v1alpha1.Operation
	10, // 15: nebius.vpc.v1alpha1.AllocationService.Update:output_type -> nebius.common.v1alpha1.Operation
	10, // 16: nebius.vpc.v1alpha1.AllocationService.Delete:output_type -> nebius.common.v1alpha1.Operation
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1alpha1_allocation_service_proto_init() }
func file_nebius_vpc_v1alpha1_allocation_service_proto_init() {
	if File_nebius_vpc_v1alpha1_allocation_service_proto != nil {
		return
	}
	file_nebius_vpc_v1alpha1_allocation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_vpc_v1alpha1_allocation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_vpc_v1alpha1_allocation_service_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1alpha1_allocation_service_proto_depIdxs,
		MessageInfos:      file_nebius_vpc_v1alpha1_allocation_service_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1alpha1_allocation_service_proto = out.File
	file_nebius_vpc_v1alpha1_allocation_service_proto_rawDesc = nil
	file_nebius_vpc_v1alpha1_allocation_service_proto_goTypes = nil
	file_nebius_vpc_v1alpha1_allocation_service_proto_depIdxs = nil
}
