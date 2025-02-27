// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/iam/v1/auth_public_key_service.proto

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

type CreateAuthPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *AuthPublicKeySpec     `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAuthPublicKeyRequest) Reset() {
	*x = CreateAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthPublicKeyRequest) ProtoMessage() {}

func (x *CreateAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAuthPublicKeyRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateAuthPublicKeyRequest) GetSpec() *AuthPublicKeySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type GetAuthPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAuthPublicKeyRequest) Reset() {
	*x = GetAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthPublicKeyRequest) ProtoMessage() {}

func (x *GetAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*GetAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthPublicKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListAuthPublicKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the container ID.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	// Default value: 10
	PageSize *int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter to narrow down the results based on specific criteria.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAuthPublicKeyRequest) Reset() {
	*x = ListAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthPublicKeyRequest) ProtoMessage() {}

func (x *ListAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*ListAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListAuthPublicKeyRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListAuthPublicKeyRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListAuthPublicKeyRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAuthPublicKeyRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListAuthPublicKeyByAccountRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the parent account ID.
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
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

func (x *ListAuthPublicKeyByAccountRequest) Reset() {
	*x = ListAuthPublicKeyByAccountRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthPublicKeyByAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthPublicKeyByAccountRequest) ProtoMessage() {}

func (x *ListAuthPublicKeyByAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthPublicKeyByAccountRequest.ProtoReflect.Descriptor instead.
func (*ListAuthPublicKeyByAccountRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAuthPublicKeyByAccountRequest) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *ListAuthPublicKeyByAccountRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAuthPublicKeyByAccountRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAuthPublicKeyByAccountRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type UpdateAuthPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *AuthPublicKeySpec     `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAuthPublicKeyRequest) Reset() {
	*x = UpdateAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthPublicKeyRequest) ProtoMessage() {}

func (x *UpdateAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAuthPublicKeyRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateAuthPublicKeyRequest) GetSpec() *AuthPublicKeySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type ActivateAuthPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActivateAuthPublicKeyRequest) Reset() {
	*x = ActivateAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivateAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateAuthPublicKeyRequest) ProtoMessage() {}

func (x *ActivateAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*ActivateAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{5}
}

func (x *ActivateAuthPublicKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeactivateAuthPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeactivateAuthPublicKeyRequest) Reset() {
	*x = DeactivateAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeactivateAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateAuthPublicKeyRequest) ProtoMessage() {}

func (x *DeactivateAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*DeactivateAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeactivateAuthPublicKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAuthPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAuthPublicKeyRequest) Reset() {
	*x = DeleteAuthPublicKeyRequest{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAuthPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAuthPublicKeyRequest) ProtoMessage() {}

func (x *DeleteAuthPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAuthPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteAuthPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAuthPublicKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListAuthPublicKeyResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of auth public keys returned in the response. The field should be named as `items` for consistency.
	Items []*AuthPublicKey `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAuthPublicKeyResponse) Reset() {
	*x = ListAuthPublicKeyResponse{}
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthPublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthPublicKeyResponse) ProtoMessage() {}

func (x *ListAuthPublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthPublicKeyResponse.ProtoReflect.Descriptor instead.
func (*ListAuthPublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListAuthPublicKeyResponse) GetItems() []*AuthPublicKey {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListAuthPublicKeyResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_iam_v1_auth_public_key_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_auth_public_key_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22,
	0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x21,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x92, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x2e, 0x0a, 0x1c,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x1e,
	0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c,
	0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x19,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a,
	0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xdd, 0x05, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x59, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x0a, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0xba, 0x4a, 0x07, 0x63, 0x70,
	0x6c, 0x2e, 0x69, 0x61, 0x6d, 0x42, 0x60, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_auth_public_key_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_auth_public_key_service_proto_rawDescData = file_nebius_iam_v1_auth_public_key_service_proto_rawDesc
)

func file_nebius_iam_v1_auth_public_key_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_auth_public_key_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_auth_public_key_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_auth_public_key_service_proto_rawDescData)
	})
	return file_nebius_iam_v1_auth_public_key_service_proto_rawDescData
}

var file_nebius_iam_v1_auth_public_key_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_nebius_iam_v1_auth_public_key_service_proto_goTypes = []any{
	(*CreateAuthPublicKeyRequest)(nil),        // 0: nebius.iam.v1.CreateAuthPublicKeyRequest
	(*GetAuthPublicKeyRequest)(nil),           // 1: nebius.iam.v1.GetAuthPublicKeyRequest
	(*ListAuthPublicKeyRequest)(nil),          // 2: nebius.iam.v1.ListAuthPublicKeyRequest
	(*ListAuthPublicKeyByAccountRequest)(nil), // 3: nebius.iam.v1.ListAuthPublicKeyByAccountRequest
	(*UpdateAuthPublicKeyRequest)(nil),        // 4: nebius.iam.v1.UpdateAuthPublicKeyRequest
	(*ActivateAuthPublicKeyRequest)(nil),      // 5: nebius.iam.v1.ActivateAuthPublicKeyRequest
	(*DeactivateAuthPublicKeyRequest)(nil),    // 6: nebius.iam.v1.DeactivateAuthPublicKeyRequest
	(*DeleteAuthPublicKeyRequest)(nil),        // 7: nebius.iam.v1.DeleteAuthPublicKeyRequest
	(*ListAuthPublicKeyResponse)(nil),         // 8: nebius.iam.v1.ListAuthPublicKeyResponse
	(*v1.ResourceMetadata)(nil),               // 9: nebius.common.v1.ResourceMetadata
	(*AuthPublicKeySpec)(nil),                 // 10: nebius.iam.v1.AuthPublicKeySpec
	(*Account)(nil),                           // 11: nebius.iam.v1.Account
	(*AuthPublicKey)(nil),                     // 12: nebius.iam.v1.AuthPublicKey
	(*v1.Operation)(nil),                      // 13: nebius.common.v1.Operation
}
var file_nebius_iam_v1_auth_public_key_service_proto_depIdxs = []int32{
	9,  // 0: nebius.iam.v1.CreateAuthPublicKeyRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	10, // 1: nebius.iam.v1.CreateAuthPublicKeyRequest.spec:type_name -> nebius.iam.v1.AuthPublicKeySpec
	11, // 2: nebius.iam.v1.ListAuthPublicKeyByAccountRequest.account:type_name -> nebius.iam.v1.Account
	9,  // 3: nebius.iam.v1.UpdateAuthPublicKeyRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	10, // 4: nebius.iam.v1.UpdateAuthPublicKeyRequest.spec:type_name -> nebius.iam.v1.AuthPublicKeySpec
	12, // 5: nebius.iam.v1.ListAuthPublicKeyResponse.items:type_name -> nebius.iam.v1.AuthPublicKey
	0,  // 6: nebius.iam.v1.AuthPublicKeyService.Create:input_type -> nebius.iam.v1.CreateAuthPublicKeyRequest
	1,  // 7: nebius.iam.v1.AuthPublicKeyService.Get:input_type -> nebius.iam.v1.GetAuthPublicKeyRequest
	2,  // 8: nebius.iam.v1.AuthPublicKeyService.List:input_type -> nebius.iam.v1.ListAuthPublicKeyRequest
	3,  // 9: nebius.iam.v1.AuthPublicKeyService.ListByAccount:input_type -> nebius.iam.v1.ListAuthPublicKeyByAccountRequest
	4,  // 10: nebius.iam.v1.AuthPublicKeyService.Update:input_type -> nebius.iam.v1.UpdateAuthPublicKeyRequest
	5,  // 11: nebius.iam.v1.AuthPublicKeyService.Activate:input_type -> nebius.iam.v1.ActivateAuthPublicKeyRequest
	6,  // 12: nebius.iam.v1.AuthPublicKeyService.Deactivate:input_type -> nebius.iam.v1.DeactivateAuthPublicKeyRequest
	7,  // 13: nebius.iam.v1.AuthPublicKeyService.Delete:input_type -> nebius.iam.v1.DeleteAuthPublicKeyRequest
	13, // 14: nebius.iam.v1.AuthPublicKeyService.Create:output_type -> nebius.common.v1.Operation
	12, // 15: nebius.iam.v1.AuthPublicKeyService.Get:output_type -> nebius.iam.v1.AuthPublicKey
	8,  // 16: nebius.iam.v1.AuthPublicKeyService.List:output_type -> nebius.iam.v1.ListAuthPublicKeyResponse
	8,  // 17: nebius.iam.v1.AuthPublicKeyService.ListByAccount:output_type -> nebius.iam.v1.ListAuthPublicKeyResponse
	13, // 18: nebius.iam.v1.AuthPublicKeyService.Update:output_type -> nebius.common.v1.Operation
	13, // 19: nebius.iam.v1.AuthPublicKeyService.Activate:output_type -> nebius.common.v1.Operation
	13, // 20: nebius.iam.v1.AuthPublicKeyService.Deactivate:output_type -> nebius.common.v1.Operation
	13, // 21: nebius.iam.v1.AuthPublicKeyService.Delete:output_type -> nebius.common.v1.Operation
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_auth_public_key_service_proto_init() }
func file_nebius_iam_v1_auth_public_key_service_proto_init() {
	if File_nebius_iam_v1_auth_public_key_service_proto != nil {
		return
	}
	file_nebius_iam_v1_auth_public_key_proto_init()
	file_nebius_iam_v1_access_proto_init()
	file_nebius_iam_v1_auth_public_key_service_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_auth_public_key_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_auth_public_key_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_auth_public_key_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_auth_public_key_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_auth_public_key_service_proto = out.File
	file_nebius_iam_v1_auth_public_key_service_proto_rawDesc = nil
	file_nebius_iam_v1_auth_public_key_service_proto_goTypes = nil
	file_nebius_iam_v1_auth_public_key_service_proto_depIdxs = nil
}
