// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/iam/v1/federation_certificate_service.proto

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

type CreateFederationCertificateRequest struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *FederationCertificateSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateFederationCertificateRequest) Reset() {
	*x = CreateFederationCertificateRequest{}
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateFederationCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFederationCertificateRequest) ProtoMessage() {}

func (x *CreateFederationCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFederationCertificateRequest.ProtoReflect.Descriptor instead.
func (*CreateFederationCertificateRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFederationCertificateRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateFederationCertificateRequest) GetSpec() *FederationCertificateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type GetFederationCertificateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetFederationCertificateRequest) Reset() {
	*x = GetFederationCertificateRequest{}
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFederationCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFederationCertificateRequest) ProtoMessage() {}

func (x *GetFederationCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFederationCertificateRequest.ProtoReflect.Descriptor instead.
func (*GetFederationCertificateRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFederationCertificateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListFederationCertificateByFederationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the parent federation ID. Corresponds to the parent_id value.
	FederationId string `protobuf:"bytes,1,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	// Default value: 10
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFederationCertificateByFederationRequest) Reset() {
	*x = ListFederationCertificateByFederationRequest{}
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFederationCertificateByFederationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationCertificateByFederationRequest) ProtoMessage() {}

func (x *ListFederationCertificateByFederationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationCertificateByFederationRequest.ProtoReflect.Descriptor instead.
func (*ListFederationCertificateByFederationRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederationCertificateByFederationRequest) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

func (x *ListFederationCertificateByFederationRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFederationCertificateByFederationRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type UpdateFederationCertificateRequest struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata       `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *FederationCertificateSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateFederationCertificateRequest) Reset() {
	*x = UpdateFederationCertificateRequest{}
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFederationCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFederationCertificateRequest) ProtoMessage() {}

func (x *UpdateFederationCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFederationCertificateRequest.ProtoReflect.Descriptor instead.
func (*UpdateFederationCertificateRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFederationCertificateRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateFederationCertificateRequest) GetSpec() *FederationCertificateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteFederationCertificateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteFederationCertificateRequest) Reset() {
	*x = DeleteFederationCertificateRequest{}
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFederationCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFederationCertificateRequest) ProtoMessage() {}

func (x *DeleteFederationCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFederationCertificateRequest.ProtoReflect.Descriptor instead.
func (*DeleteFederationCertificateRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFederationCertificateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListFederationCertificateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of public keys returned in the response. The field should be named as `items` for consistency.
	Items []*FederationCertificate `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListFederationCertificateResponse) Reset() {
	*x = ListFederationCertificateResponse{}
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListFederationCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederationCertificateResponse) ProtoMessage() {}

func (x *ListFederationCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederationCertificateResponse.ProtoReflect.Descriptor instead.
func (*ListFederationCertificateResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListFederationCertificateResponse) GetItems() []*FederationCertificate {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListFederationCertificateResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_iam_v1_federation_certificate_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_federation_certificate_service_proto_rawDesc = string([]byte{
	0x0a, 0x32, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a,
	0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x22, 0x31, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x2c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x22, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x34, 0x0a, 0x22, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x87, 0x01, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x99, 0x04, 0x0a, 0x1c,
	0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x31, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x58, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0xba, 0x4a, 0x07,
	0x63, 0x70, 0x6c, 0x2e, 0x69, 0x61, 0x6d, 0x42, 0x68, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x21, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v1_federation_certificate_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_federation_certificate_service_proto_rawDescData []byte
)

func file_nebius_iam_v1_federation_certificate_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_federation_certificate_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_federation_certificate_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_federation_certificate_service_proto_rawDesc), len(file_nebius_iam_v1_federation_certificate_service_proto_rawDesc)))
	})
	return file_nebius_iam_v1_federation_certificate_service_proto_rawDescData
}

var file_nebius_iam_v1_federation_certificate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nebius_iam_v1_federation_certificate_service_proto_goTypes = []any{
	(*CreateFederationCertificateRequest)(nil),           // 0: nebius.iam.v1.CreateFederationCertificateRequest
	(*GetFederationCertificateRequest)(nil),              // 1: nebius.iam.v1.GetFederationCertificateRequest
	(*ListFederationCertificateByFederationRequest)(nil), // 2: nebius.iam.v1.ListFederationCertificateByFederationRequest
	(*UpdateFederationCertificateRequest)(nil),           // 3: nebius.iam.v1.UpdateFederationCertificateRequest
	(*DeleteFederationCertificateRequest)(nil),           // 4: nebius.iam.v1.DeleteFederationCertificateRequest
	(*ListFederationCertificateResponse)(nil),            // 5: nebius.iam.v1.ListFederationCertificateResponse
	(*v1.ResourceMetadata)(nil),                          // 6: nebius.common.v1.ResourceMetadata
	(*FederationCertificateSpec)(nil),                    // 7: nebius.iam.v1.FederationCertificateSpec
	(*FederationCertificate)(nil),                        // 8: nebius.iam.v1.FederationCertificate
	(*v1.Operation)(nil),                                 // 9: nebius.common.v1.Operation
}
var file_nebius_iam_v1_federation_certificate_service_proto_depIdxs = []int32{
	6,  // 0: nebius.iam.v1.CreateFederationCertificateRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	7,  // 1: nebius.iam.v1.CreateFederationCertificateRequest.spec:type_name -> nebius.iam.v1.FederationCertificateSpec
	6,  // 2: nebius.iam.v1.UpdateFederationCertificateRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	7,  // 3: nebius.iam.v1.UpdateFederationCertificateRequest.spec:type_name -> nebius.iam.v1.FederationCertificateSpec
	8,  // 4: nebius.iam.v1.ListFederationCertificateResponse.items:type_name -> nebius.iam.v1.FederationCertificate
	0,  // 5: nebius.iam.v1.FederationCertificateService.Create:input_type -> nebius.iam.v1.CreateFederationCertificateRequest
	1,  // 6: nebius.iam.v1.FederationCertificateService.Get:input_type -> nebius.iam.v1.GetFederationCertificateRequest
	2,  // 7: nebius.iam.v1.FederationCertificateService.ListByFederation:input_type -> nebius.iam.v1.ListFederationCertificateByFederationRequest
	3,  // 8: nebius.iam.v1.FederationCertificateService.Update:input_type -> nebius.iam.v1.UpdateFederationCertificateRequest
	4,  // 9: nebius.iam.v1.FederationCertificateService.Delete:input_type -> nebius.iam.v1.DeleteFederationCertificateRequest
	9,  // 10: nebius.iam.v1.FederationCertificateService.Create:output_type -> nebius.common.v1.Operation
	8,  // 11: nebius.iam.v1.FederationCertificateService.Get:output_type -> nebius.iam.v1.FederationCertificate
	5,  // 12: nebius.iam.v1.FederationCertificateService.ListByFederation:output_type -> nebius.iam.v1.ListFederationCertificateResponse
	9,  // 13: nebius.iam.v1.FederationCertificateService.Update:output_type -> nebius.common.v1.Operation
	9,  // 14: nebius.iam.v1.FederationCertificateService.Delete:output_type -> nebius.common.v1.Operation
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_federation_certificate_service_proto_init() }
func file_nebius_iam_v1_federation_certificate_service_proto_init() {
	if File_nebius_iam_v1_federation_certificate_service_proto != nil {
		return
	}
	file_nebius_iam_v1_federation_certificate_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_federation_certificate_service_proto_rawDesc), len(file_nebius_iam_v1_federation_certificate_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_federation_certificate_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_federation_certificate_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_federation_certificate_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_federation_certificate_service_proto = out.File
	file_nebius_iam_v1_federation_certificate_service_proto_goTypes = nil
	file_nebius_iam_v1_federation_certificate_service_proto_depIdxs = nil
}
