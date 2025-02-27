// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// nebius/common/v1alpha1/operation_service.proto is a deprecated file.

package v1alpha1

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

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
type GetOperationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Operation ID.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOperationRequest) Reset() {
	*x = GetOperationRequest{}
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationRequest) ProtoMessage() {}

func (x *GetOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationRequest.ProtoReflect.Descriptor instead.
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return file_nebius_common_v1alpha1_operation_service_proto_rawDescGZIP(), []int{0}
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *GetOperationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
type ListOperationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Resource to list operations for.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Page size. [1...1000]. Optional, if not specified, a reasonable default will be chosen by the service.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Listing continuation token. Empty to start listing from the first page.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Filter expression for the listing results. Optional.
	// Filter expression format: TBD.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperationsRequest) Reset() {
	*x = ListOperationsRequest{}
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationsRequest) ProtoMessage() {}

func (x *ListOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationsRequest.ProtoReflect.Descriptor instead.
func (*ListOperationsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_common_v1alpha1_operation_service_proto_rawDescGZIP(), []int{1}
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
type ListOperationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of operations on this result page.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	Operations []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// Listing continuation token for the next page of results.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperationsResponse) Reset() {
	*x = ListOperationsResponse{}
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationsResponse) ProtoMessage() {}

func (x *ListOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationsResponse.ProtoReflect.Descriptor instead.
func (*ListOperationsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_common_v1alpha1_operation_service_proto_rawDescGZIP(), []int{2}
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsResponse) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
type ListOperationsByParentRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the parent to list operations for resource type at.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Page size. [1...1000]. Optional, if not specified, a reasonable default will be chosen by the service.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Listing continuation token. Empty to start listing from the first page.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Filter expression for the listing results. Optional.
	// Filter expression format: TBD.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperationsByParentRequest) Reset() {
	*x = ListOperationsByParentRequest{}
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationsByParentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationsByParentRequest) ProtoMessage() {}

func (x *ListOperationsByParentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_common_v1alpha1_operation_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationsByParentRequest.ProtoReflect.Descriptor instead.
func (*ListOperationsByParentRequest) Descriptor() ([]byte, []int) {
	return file_nebius_common_v1alpha1_operation_service_proto_rawDescGZIP(), []int{3}
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsByParentRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsByParentRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsByParentRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation_service.proto is marked as deprecated.
func (x *ListOperationsByParentRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

var File_nebius_common_v1alpha1_operation_service_proto protoreflect.FileDescriptor

var file_nebius_common_v1alpha1_operation_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x02, 0x18, 0x01, 0x22,
	0x9c, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x3a, 0x02, 0x18, 0x01, 0x32, 0xe8,
	0x01, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x78, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x11, 0x9a, 0xb5, 0x18, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x1a, 0x03, 0x88, 0x02, 0x01, 0x42, 0x71, 0x0a, 0x1d, 0x61, 0x69, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x15, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xb8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_common_v1alpha1_operation_service_proto_rawDescOnce sync.Once
	file_nebius_common_v1alpha1_operation_service_proto_rawDescData = file_nebius_common_v1alpha1_operation_service_proto_rawDesc
)

func file_nebius_common_v1alpha1_operation_service_proto_rawDescGZIP() []byte {
	file_nebius_common_v1alpha1_operation_service_proto_rawDescOnce.Do(func() {
		file_nebius_common_v1alpha1_operation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_common_v1alpha1_operation_service_proto_rawDescData)
	})
	return file_nebius_common_v1alpha1_operation_service_proto_rawDescData
}

var file_nebius_common_v1alpha1_operation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_common_v1alpha1_operation_service_proto_goTypes = []any{
	(*GetOperationRequest)(nil),           // 0: nebius.common.v1alpha1.GetOperationRequest
	(*ListOperationsRequest)(nil),         // 1: nebius.common.v1alpha1.ListOperationsRequest
	(*ListOperationsResponse)(nil),        // 2: nebius.common.v1alpha1.ListOperationsResponse
	(*ListOperationsByParentRequest)(nil), // 3: nebius.common.v1alpha1.ListOperationsByParentRequest
	(*Operation)(nil),                     // 4: nebius.common.v1alpha1.Operation
}
var file_nebius_common_v1alpha1_operation_service_proto_depIdxs = []int32{
	4, // 0: nebius.common.v1alpha1.ListOperationsResponse.operations:type_name -> nebius.common.v1alpha1.Operation
	0, // 1: nebius.common.v1alpha1.OperationService.Get:input_type -> nebius.common.v1alpha1.GetOperationRequest
	1, // 2: nebius.common.v1alpha1.OperationService.List:input_type -> nebius.common.v1alpha1.ListOperationsRequest
	4, // 3: nebius.common.v1alpha1.OperationService.Get:output_type -> nebius.common.v1alpha1.Operation
	2, // 4: nebius.common.v1alpha1.OperationService.List:output_type -> nebius.common.v1alpha1.ListOperationsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_common_v1alpha1_operation_service_proto_init() }
func file_nebius_common_v1alpha1_operation_service_proto_init() {
	if File_nebius_common_v1alpha1_operation_service_proto != nil {
		return
	}
	file_nebius_common_v1alpha1_operation_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_common_v1alpha1_operation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_common_v1alpha1_operation_service_proto_goTypes,
		DependencyIndexes: file_nebius_common_v1alpha1_operation_service_proto_depIdxs,
		MessageInfos:      file_nebius_common_v1alpha1_operation_service_proto_msgTypes,
	}.Build()
	File_nebius_common_v1alpha1_operation_service_proto = out.File
	file_nebius_common_v1alpha1_operation_service_proto_rawDesc = nil
	file_nebius_common_v1alpha1_operation_service_proto_goTypes = nil
	file_nebius_common_v1alpha1_operation_service_proto_depIdxs = nil
}
