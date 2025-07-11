// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// nebius/common/v1alpha1/operation.proto is a deprecated file.

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
type Operation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the operation.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Human readable description of the operation. 0-256 characters long.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Creation timestamp.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the user or service account who initiated the operation.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	CreatedBy string `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// The time when the operation finished.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// The request that generated this operation.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	Request *anypb.Any `protobuf:"bytes,6,opt,name=request,proto3" json:"request,omitempty"`
	// The request headers that are essential for the request that generated the operation.
	// For instance, `x-resetmask`. Without these headers the request might have been processed
	// differently if repeated.
	// All the header names *must* be converted to lower case.
	// Validator is based on:
	// https://httpwg.org/specs/rfc9110.html#considerations.for.new.field.names
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	RequestHeaders map[string]*OperationRequestHeader `protobuf:"bytes,11,rep,name=request_headers,json=requestHeaders,proto3" json:"request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ID of the resource that this operation creates, updates, deletes or otherwise changes.
	//
	// If the operation affects multiple resources or does not affect any API resources at all
	// (e.g. a routine maintenance operation visible to the user), the [resource_id] must be empty.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	ResourceId string `protobuf:"bytes,7,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// Snapshot of the resource at the moment this operation started.
	//   - [resource.spec] and [resource.metadata] reflect the desired resource state at the moment
	//     this operation started.
	//     E.g., in an Update operation it will be the *updated* resource spec and metadata,
	//     in a Create operation it will be the spec and metadata *of the resource being created*,
	//     and so on.
	//   - [resource.status] reflects the status of the resource at the moment this operation started.
	//     This is a snapshot, call the <Resource>Service/Get to get current status of the resource.
	//
	// The [resource] field MUST never be updated *after* this operation has started.
	//
	// In a Delete operation, an operation affecting multiple resources or an operation that doesn't
	// affect any API resources at all (e.g. a routine maintenance operation visible to the user),
	// the [resource] inside MUST be a [google.protobuf.Empty].
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	Resource *anypb.Any `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	// Additional information about the progress of an operation, e.g., a progress percentage.
	// MAY be absent while the operation is running, MUST be absent after the operation has completed.
	//
	// Format of message inside [progress_data] is service-dependent and MUST be documented by the
	// service, IF it is used.
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	ProgressData *anypb.Any `protobuf:"bytes,9,opt,name=progress_data,json=progressData,proto3" json:"progress_data,omitempty"`
	// The status of this operation. Set when this operation is completed.
	// See https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto.
	//
	// [status.code] is https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto:
	// - If [status.code] == OK, the operation has completed successfully.
	// - If [status.code] != OK, the operation has failed or has been cancelled.
	//   - [status.message] will contain a user-readable and actionable error message.
	//   - [status.details] will contain additional diagnostic information in the form of
	//     [ServiceError] from ../error/v1alpha1/error.proto
	//   - [status.code] must belong to an Operation-compatible subset of GRPC codes:
	//     OK, CANCELLED, PERMISSION_DENIED, RESOURCE_EXHAUSTED, FAILED_PRECONDITION, ABORTED, INTERNAL
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	Status        *status.Status `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Operation) Reset() {
	*x = Operation{}
	mi := &file_nebius_common_v1alpha1_operation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_common_v1alpha1_operation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_nebius_common_v1alpha1_operation_proto_rawDescGZIP(), []int{0}
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetRequest() *anypb.Any {
	if x != nil {
		return x.Request
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetRequestHeaders() map[string]*OperationRequestHeader {
	if x != nil {
		return x.RequestHeaders
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetResource() *anypb.Any {
	if x != nil {
		return x.Resource
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetProgressData() *anypb.Any {
	if x != nil {
		return x.ProgressData
	}
	return nil
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *Operation) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// Request header is a container for all the values of a particular header of a request
// as there is no such thing as map<string, repeated string>
//
// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
type OperationRequestHeader struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The values of a particular header from a request
	//
	// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
	Values        []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperationRequestHeader) Reset() {
	*x = OperationRequestHeader{}
	mi := &file_nebius_common_v1alpha1_operation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperationRequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationRequestHeader) ProtoMessage() {}

func (x *OperationRequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_common_v1alpha1_operation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationRequestHeader.ProtoReflect.Descriptor instead.
func (*OperationRequestHeader) Descriptor() ([]byte, []int) {
	return file_nebius_common_v1alpha1_operation_proto_rawDescGZIP(), []int{0, 0}
}

// Deprecated: The entire proto file nebius/common/v1alpha1/operation.proto is marked as deprecated.
func (x *OperationRequestHeader) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_nebius_common_v1alpha1_operation_proto protoreflect.FileDescriptor

var file_nebius_common_v1alpha1_operation_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x05, 0x0a,
	0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x7c, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x1c, 0xba, 0x48, 0x19, 0x9a, 0x01, 0x16, 0x22, 0x14, 0x72, 0x12, 0x32,
	0x10, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x5c, 0x2e, 0x5d, 0x2a,
	0x24, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x06, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x1a, 0x28, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x73, 0x0a, 0x13, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x46, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x02, 0x18, 0x01, 0x42, 0x6a, 0x0a, 0x1d, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xb8, 0x01, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_common_v1alpha1_operation_proto_rawDescOnce sync.Once
	file_nebius_common_v1alpha1_operation_proto_rawDescData []byte
)

func file_nebius_common_v1alpha1_operation_proto_rawDescGZIP() []byte {
	file_nebius_common_v1alpha1_operation_proto_rawDescOnce.Do(func() {
		file_nebius_common_v1alpha1_operation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_common_v1alpha1_operation_proto_rawDesc), len(file_nebius_common_v1alpha1_operation_proto_rawDesc)))
	})
	return file_nebius_common_v1alpha1_operation_proto_rawDescData
}

var file_nebius_common_v1alpha1_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_common_v1alpha1_operation_proto_goTypes = []any{
	(*Operation)(nil),              // 0: nebius.common.v1alpha1.Operation
	(*OperationRequestHeader)(nil), // 1: nebius.common.v1alpha1.Operation.request_header
	nil,                            // 2: nebius.common.v1alpha1.Operation.RequestHeadersEntry
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*anypb.Any)(nil),              // 4: google.protobuf.Any
	(*status.Status)(nil),          // 5: google.rpc.Status
}
var file_nebius_common_v1alpha1_operation_proto_depIdxs = []int32{
	3, // 0: nebius.common.v1alpha1.Operation.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: nebius.common.v1alpha1.Operation.finished_at:type_name -> google.protobuf.Timestamp
	4, // 2: nebius.common.v1alpha1.Operation.request:type_name -> google.protobuf.Any
	2, // 3: nebius.common.v1alpha1.Operation.request_headers:type_name -> nebius.common.v1alpha1.Operation.RequestHeadersEntry
	4, // 4: nebius.common.v1alpha1.Operation.resource:type_name -> google.protobuf.Any
	4, // 5: nebius.common.v1alpha1.Operation.progress_data:type_name -> google.protobuf.Any
	5, // 6: nebius.common.v1alpha1.Operation.status:type_name -> google.rpc.Status
	1, // 7: nebius.common.v1alpha1.Operation.RequestHeadersEntry.value:type_name -> nebius.common.v1alpha1.Operation.request_header
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_nebius_common_v1alpha1_operation_proto_init() }
func file_nebius_common_v1alpha1_operation_proto_init() {
	if File_nebius_common_v1alpha1_operation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_common_v1alpha1_operation_proto_rawDesc), len(file_nebius_common_v1alpha1_operation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_common_v1alpha1_operation_proto_goTypes,
		DependencyIndexes: file_nebius_common_v1alpha1_operation_proto_depIdxs,
		MessageInfos:      file_nebius_common_v1alpha1_operation_proto_msgTypes,
	}.Build()
	File_nebius_common_v1alpha1_operation_proto = out.File
	file_nebius_common_v1alpha1_operation_proto_goTypes = nil
	file_nebius_common_v1alpha1_operation_proto_depIdxs = nil
}
