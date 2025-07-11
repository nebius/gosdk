// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/iam/v1/access_permit_service.proto

package v1

import (
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccessPermitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *AccessPermitSpec      `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccessPermitRequest) Reset() {
	*x = CreateAccessPermitRequest{}
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccessPermitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessPermitRequest) ProtoMessage() {}

func (x *CreateAccessPermitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessPermitRequest.ProtoReflect.Descriptor instead.
func (*CreateAccessPermitRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_access_permit_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccessPermitRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateAccessPermitRequest) GetSpec() *AccessPermitSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type ListAccessPermitRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the container ID.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter to narrow down the results based on specific criteria.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAccessPermitRequest) Reset() {
	*x = ListAccessPermitRequest{}
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccessPermitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessPermitRequest) ProtoMessage() {}

func (x *ListAccessPermitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessPermitRequest.ProtoReflect.Descriptor instead.
func (*ListAccessPermitRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_access_permit_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListAccessPermitRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListAccessPermitRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAccessPermitRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAccessPermitRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type DeleteAccessPermitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAccessPermitRequest) Reset() {
	*x = DeleteAccessPermitRequest{}
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccessPermitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessPermitRequest) ProtoMessage() {}

func (x *DeleteAccessPermitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessPermitRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccessPermitRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_access_permit_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAccessPermitRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAccessPermitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccessPermitRequest) Reset() {
	*x = GetAccessPermitRequest{}
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccessPermitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessPermitRequest) ProtoMessage() {}

func (x *GetAccessPermitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessPermitRequest.ProtoReflect.Descriptor instead.
func (*GetAccessPermitRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_access_permit_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccessPermitRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListAccessPermitResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of access bindings returned in the response. The field should be named as `items` for consistency.
	Items []*AccessPermit `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAccessPermitResponse) Reset() {
	*x = ListAccessPermitResponse{}
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccessPermitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessPermitResponse) ProtoMessage() {}

func (x *ListAccessPermitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_access_permit_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessPermitResponse.ProtoReflect.Descriptor instead.
func (*ListAccessPermitResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_access_permit_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListAccessPermitResponse) GetItems() []*AccessPermit {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListAccessPermitResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_iam_v1_access_permit_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_access_permit_service_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x8a, 0x01,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x75, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe7, 0x02, 0x0a, 0x13, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x57, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x1a, 0x0a, 0xba, 0x4a, 0x07, 0x63, 0x70, 0x6c, 0x2e, 0x69,
	0x61, 0x6d, 0x42, 0x5f, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x18, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v1_access_permit_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_access_permit_service_proto_rawDescData []byte
)

func file_nebius_iam_v1_access_permit_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_access_permit_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_access_permit_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_access_permit_service_proto_rawDesc), len(file_nebius_iam_v1_access_permit_service_proto_rawDesc)))
	})
	return file_nebius_iam_v1_access_permit_service_proto_rawDescData
}

var file_nebius_iam_v1_access_permit_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_iam_v1_access_permit_service_proto_goTypes = []any{
	(*CreateAccessPermitRequest)(nil), // 0: nebius.iam.v1.CreateAccessPermitRequest
	(*ListAccessPermitRequest)(nil),   // 1: nebius.iam.v1.ListAccessPermitRequest
	(*DeleteAccessPermitRequest)(nil), // 2: nebius.iam.v1.DeleteAccessPermitRequest
	(*GetAccessPermitRequest)(nil),    // 3: nebius.iam.v1.GetAccessPermitRequest
	(*ListAccessPermitResponse)(nil),  // 4: nebius.iam.v1.ListAccessPermitResponse
	(*v1.ResourceMetadata)(nil),       // 5: nebius.common.v1.ResourceMetadata
	(*AccessPermitSpec)(nil),          // 6: nebius.iam.v1.AccessPermitSpec
	(*AccessPermit)(nil),              // 7: nebius.iam.v1.AccessPermit
	(*v1.Operation)(nil),              // 8: nebius.common.v1.Operation
}
var file_nebius_iam_v1_access_permit_service_proto_depIdxs = []int32{
	5, // 0: nebius.iam.v1.CreateAccessPermitRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	6, // 1: nebius.iam.v1.CreateAccessPermitRequest.spec:type_name -> nebius.iam.v1.AccessPermitSpec
	7, // 2: nebius.iam.v1.ListAccessPermitResponse.items:type_name -> nebius.iam.v1.AccessPermit
	0, // 3: nebius.iam.v1.AccessPermitService.Create:input_type -> nebius.iam.v1.CreateAccessPermitRequest
	1, // 4: nebius.iam.v1.AccessPermitService.List:input_type -> nebius.iam.v1.ListAccessPermitRequest
	2, // 5: nebius.iam.v1.AccessPermitService.Delete:input_type -> nebius.iam.v1.DeleteAccessPermitRequest
	3, // 6: nebius.iam.v1.AccessPermitService.Get:input_type -> nebius.iam.v1.GetAccessPermitRequest
	8, // 7: nebius.iam.v1.AccessPermitService.Create:output_type -> nebius.common.v1.Operation
	4, // 8: nebius.iam.v1.AccessPermitService.List:output_type -> nebius.iam.v1.ListAccessPermitResponse
	8, // 9: nebius.iam.v1.AccessPermitService.Delete:output_type -> nebius.common.v1.Operation
	7, // 10: nebius.iam.v1.AccessPermitService.Get:output_type -> nebius.iam.v1.AccessPermit
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_access_permit_service_proto_init() }
func file_nebius_iam_v1_access_permit_service_proto_init() {
	if File_nebius_iam_v1_access_permit_service_proto != nil {
		return
	}
	file_nebius_iam_v1_access_permit_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_access_permit_service_proto_rawDesc), len(file_nebius_iam_v1_access_permit_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_access_permit_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_access_permit_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_access_permit_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_access_permit_service_proto = out.File
	file_nebius_iam_v1_access_permit_service_proto_goTypes = nil
	file_nebius_iam_v1_access_permit_service_proto_depIdxs = nil
}
