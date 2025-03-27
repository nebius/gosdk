// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/iam/v1/static_key_service.proto

package v1

import (
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
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

type GetStaticKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // id of the static key
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStaticKeyRequest) Reset() {
	*x = GetStaticKeyRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticKeyRequest) ProtoMessage() {}

func (x *GetStaticKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticKeyRequest.ProtoReflect.Descriptor instead.
func (*GetStaticKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetStaticKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetStaticKeyByNameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"` // id of the parent container (service account)
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                         // name of the static key
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStaticKeyByNameRequest) Reset() {
	*x = GetStaticKeyByNameRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStaticKeyByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStaticKeyByNameRequest) ProtoMessage() {}

func (x *GetStaticKeyByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStaticKeyByNameRequest.ProtoReflect.Descriptor instead.
func (*GetStaticKeyByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetStaticKeyByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetStaticKeyByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteStaticKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // id of the static key to delete
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStaticKeyRequest) Reset() {
	*x = DeleteStaticKeyRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStaticKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStaticKeyRequest) ProtoMessage() {}

func (x *DeleteStaticKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStaticKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteStaticKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteStaticKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListStaticKeysRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the container ID.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	// Default value: 10
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter to narrow down the results based on specific criteria.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStaticKeysRequest) Reset() {
	*x = ListStaticKeysRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStaticKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaticKeysRequest) ProtoMessage() {}

func (x *ListStaticKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaticKeysRequest.ProtoReflect.Descriptor instead.
func (*ListStaticKeysRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListStaticKeysRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListStaticKeysRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListStaticKeysRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListStaticKeysRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListStaticKeysResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of static keys returned in the response. The field should be named as `items` for consistency.
	Items []*StaticKey `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStaticKeysResponse) Reset() {
	*x = ListStaticKeysResponse{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStaticKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStaticKeysResponse) ProtoMessage() {}

func (x *ListStaticKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStaticKeysResponse.ProtoReflect.Descriptor instead.
func (*ListStaticKeysResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListStaticKeysResponse) GetItems() []*StaticKey {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListStaticKeysResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type IssueStaticKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *StaticKeySpec         `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueStaticKeyRequest) Reset() {
	*x = IssueStaticKeyRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueStaticKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStaticKeyRequest) ProtoMessage() {}

func (x *IssueStaticKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStaticKeyRequest.ProtoReflect.Descriptor instead.
func (*IssueStaticKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{5}
}

func (x *IssueStaticKeyRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IssueStaticKeyRequest) GetSpec() *StaticKeySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type IssueStaticKeyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Operation     *v1.Operation          `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IssueStaticKeyResponse) Reset() {
	*x = IssueStaticKeyResponse{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueStaticKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueStaticKeyResponse) ProtoMessage() {}

func (x *IssueStaticKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueStaticKeyResponse.ProtoReflect.Descriptor instead.
func (*IssueStaticKeyResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{6}
}

func (x *IssueStaticKeyResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *IssueStaticKeyResponse) GetOperation() *v1.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type FindStaticKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the method accepts a static key token with and without signature as an input
	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindStaticKeyRequest) Reset() {
	*x = FindStaticKeyRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindStaticKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStaticKeyRequest) ProtoMessage() {}

func (x *FindStaticKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStaticKeyRequest.ProtoReflect.Descriptor instead.
func (*FindStaticKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{7}
}

func (x *FindStaticKeyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type FindStaticKeyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StaticKey     *StaticKey             `protobuf:"bytes,1,opt,name=static_key,json=staticKey,proto3" json:"static_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindStaticKeyResponse) Reset() {
	*x = FindStaticKeyResponse{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindStaticKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStaticKeyResponse) ProtoMessage() {}

func (x *FindStaticKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStaticKeyResponse.ProtoReflect.Descriptor instead.
func (*FindStaticKeyResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{8}
}

func (x *FindStaticKeyResponse) GetStaticKey() *StaticKey {
	if x != nil {
		return x.StaticKey
	}
	return nil
}

type RevokeStaticKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the method accepts a static key token with and without signature as an input
	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeStaticKeyRequest) Reset() {
	*x = RevokeStaticKeyRequest{}
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeStaticKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeStaticKeyRequest) ProtoMessage() {}

func (x *RevokeStaticKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_static_key_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeStaticKeyRequest.ProtoReflect.Descriptor instead.
func (*RevokeStaticKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_static_key_service_proto_rawDescGZIP(), []int{9}
}

func (x *RevokeStaticKeyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_nebius_iam_v1_static_key_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_static_key_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x70,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x89, 0x01, 0x0a, 0x15, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x6e, 0x0a, 0x16,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x39, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x14,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x50, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x22, 0x33, 0x0a, 0x16, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xce, 0x04, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x05, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x22, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x4f, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x4c, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x04, 0x46, 0x69,
	0x6e, 0x64, 0x12, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a,
	0x06, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0xba, 0x4a, 0x07,
	0x63, 0x70, 0x6c, 0x2e, 0x69, 0x61, 0x6d, 0x42, 0x5c, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x15, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_static_key_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_static_key_service_proto_rawDescData = file_nebius_iam_v1_static_key_service_proto_rawDesc
)

func file_nebius_iam_v1_static_key_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_static_key_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_static_key_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_static_key_service_proto_rawDescData)
	})
	return file_nebius_iam_v1_static_key_service_proto_rawDescData
}

var file_nebius_iam_v1_static_key_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_nebius_iam_v1_static_key_service_proto_goTypes = []any{
	(*GetStaticKeyRequest)(nil),       // 0: nebius.iam.v1.GetStaticKeyRequest
	(*GetStaticKeyByNameRequest)(nil), // 1: nebius.iam.v1.GetStaticKeyByNameRequest
	(*DeleteStaticKeyRequest)(nil),    // 2: nebius.iam.v1.DeleteStaticKeyRequest
	(*ListStaticKeysRequest)(nil),     // 3: nebius.iam.v1.ListStaticKeysRequest
	(*ListStaticKeysResponse)(nil),    // 4: nebius.iam.v1.ListStaticKeysResponse
	(*IssueStaticKeyRequest)(nil),     // 5: nebius.iam.v1.IssueStaticKeyRequest
	(*IssueStaticKeyResponse)(nil),    // 6: nebius.iam.v1.IssueStaticKeyResponse
	(*FindStaticKeyRequest)(nil),      // 7: nebius.iam.v1.FindStaticKeyRequest
	(*FindStaticKeyResponse)(nil),     // 8: nebius.iam.v1.FindStaticKeyResponse
	(*RevokeStaticKeyRequest)(nil),    // 9: nebius.iam.v1.RevokeStaticKeyRequest
	(*StaticKey)(nil),                 // 10: nebius.iam.v1.StaticKey
	(*v1.ResourceMetadata)(nil),       // 11: nebius.common.v1.ResourceMetadata
	(*StaticKeySpec)(nil),             // 12: nebius.iam.v1.StaticKeySpec
	(*v1.Operation)(nil),              // 13: nebius.common.v1.Operation
}
var file_nebius_iam_v1_static_key_service_proto_depIdxs = []int32{
	10, // 0: nebius.iam.v1.ListStaticKeysResponse.items:type_name -> nebius.iam.v1.StaticKey
	11, // 1: nebius.iam.v1.IssueStaticKeyRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	12, // 2: nebius.iam.v1.IssueStaticKeyRequest.spec:type_name -> nebius.iam.v1.StaticKeySpec
	13, // 3: nebius.iam.v1.IssueStaticKeyResponse.operation:type_name -> nebius.common.v1.Operation
	10, // 4: nebius.iam.v1.FindStaticKeyResponse.static_key:type_name -> nebius.iam.v1.StaticKey
	5,  // 5: nebius.iam.v1.StaticKeyService.Issue:input_type -> nebius.iam.v1.IssueStaticKeyRequest
	3,  // 6: nebius.iam.v1.StaticKeyService.List:input_type -> nebius.iam.v1.ListStaticKeysRequest
	0,  // 7: nebius.iam.v1.StaticKeyService.Get:input_type -> nebius.iam.v1.GetStaticKeyRequest
	1,  // 8: nebius.iam.v1.StaticKeyService.GetByName:input_type -> nebius.iam.v1.GetStaticKeyByNameRequest
	2,  // 9: nebius.iam.v1.StaticKeyService.Delete:input_type -> nebius.iam.v1.DeleteStaticKeyRequest
	7,  // 10: nebius.iam.v1.StaticKeyService.Find:input_type -> nebius.iam.v1.FindStaticKeyRequest
	9,  // 11: nebius.iam.v1.StaticKeyService.Revoke:input_type -> nebius.iam.v1.RevokeStaticKeyRequest
	6,  // 12: nebius.iam.v1.StaticKeyService.Issue:output_type -> nebius.iam.v1.IssueStaticKeyResponse
	4,  // 13: nebius.iam.v1.StaticKeyService.List:output_type -> nebius.iam.v1.ListStaticKeysResponse
	10, // 14: nebius.iam.v1.StaticKeyService.Get:output_type -> nebius.iam.v1.StaticKey
	10, // 15: nebius.iam.v1.StaticKeyService.GetByName:output_type -> nebius.iam.v1.StaticKey
	13, // 16: nebius.iam.v1.StaticKeyService.Delete:output_type -> nebius.common.v1.Operation
	8,  // 17: nebius.iam.v1.StaticKeyService.Find:output_type -> nebius.iam.v1.FindStaticKeyResponse
	13, // 18: nebius.iam.v1.StaticKeyService.Revoke:output_type -> nebius.common.v1.Operation
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_static_key_service_proto_init() }
func file_nebius_iam_v1_static_key_service_proto_init() {
	if File_nebius_iam_v1_static_key_service_proto != nil {
		return
	}
	file_nebius_iam_v1_static_key_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_static_key_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_static_key_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_static_key_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_static_key_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_static_key_service_proto = out.File
	file_nebius_iam_v1_static_key_service_proto_rawDesc = nil
	file_nebius_iam_v1_static_key_service_proto_goTypes = nil
	file_nebius_iam_v1_static_key_service_proto_depIdxs = nil
}
