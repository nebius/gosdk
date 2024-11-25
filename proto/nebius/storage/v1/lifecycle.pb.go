// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/storage/v1/lifecycle.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LifecycleRule_Status int32

const (
	LifecycleRule_STATUS_UNSPECIFIED LifecycleRule_Status = 0
	LifecycleRule_ENABLED            LifecycleRule_Status = 1
	LifecycleRule_DISABLED           LifecycleRule_Status = 2
)

// Enum value maps for LifecycleRule_Status.
var (
	LifecycleRule_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ENABLED",
		2: "DISABLED",
	}
	LifecycleRule_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ENABLED":            1,
		"DISABLED":           2,
	}
)

func (x LifecycleRule_Status) Enum() *LifecycleRule_Status {
	p := new(LifecycleRule_Status)
	*p = x
	return p
}

func (x LifecycleRule_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LifecycleRule_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_storage_v1_lifecycle_proto_enumTypes[0].Descriptor()
}

func (LifecycleRule_Status) Type() protoreflect.EnumType {
	return &file_nebius_storage_v1_lifecycle_proto_enumTypes[0]
}

func (x LifecycleRule_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LifecycleRule_Status.Descriptor instead.
func (LifecycleRule_Status) EnumDescriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{1, 0}
}

// The lifecycle configuration consists of one or more rules.
// An Lifecycle configuration can have up to 1,000 rules.
// Each rule consists of the following:
//   - A filter identifying a subset of objects to which the rule applies.
//     The filter can be based on a key name prefix, object size, or any combination of these.
//   - A status indicating whether the rule is currently active.
//   - One or more lifecycle expiration actions that you want to be performed on the objects
//     identified by the filter. If the state of your bucket is versioning-enabled or versioning-suspended
//     (bucket.spec.versioning_policy equals to ENABLED or SUSPENDED) you can have many versions of the same
//     object (one current version and zero or more noncurrent versions). The system provides predefined actions
//     that you can specify for current and noncurrent object versions.
type LifecycleConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*LifecycleRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *LifecycleConfiguration) Reset() {
	*x = LifecycleConfiguration{}
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LifecycleConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleConfiguration) ProtoMessage() {}

func (x *LifecycleConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleConfiguration.ProtoReflect.Descriptor instead.
func (*LifecycleConfiguration) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{0}
}

func (x *LifecycleConfiguration) GetRules() []*LifecycleRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type LifecycleRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the rule per configuration.
	// The value cannot be longer than 255 characters.
	Id     string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status LifecycleRule_Status `protobuf:"varint,2,opt,name=status,proto3,enum=nebius.storage.v1.LifecycleRule_Status" json:"status,omitempty"`
	// The Filter is used to identify objects that a Lifecycle Rule applies to.
	// The Lifecycle Rule will apply to any object matching all of the predicates
	// configured inside (using logical AND).
	Filter *LifecycleFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Specifies the expiration for the lifecycle of the object in the form of date, days and,
	// whether the object has a delete marker.
	Expiration *LifecycleExpiration `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// Specifies when noncurrent object versions expire.
	// It works only on a bucket that has versioning enabled (or suspended).
	NoncurrentVersionExpiration *LifecycleNoncurrentVersionExpiration `protobuf:"bytes,5,opt,name=noncurrent_version_expiration,json=noncurrentVersionExpiration,proto3" json:"noncurrent_version_expiration,omitempty"`
	// Specifies the days since the initiation of an incomplete multipart upload that
	// the system will wait before permanently removing all parts of the upload.
	AbortIncompleteMultipartUpload *LifecycleAbortIncompleteMultipartUpload `protobuf:"bytes,6,opt,name=abort_incomplete_multipart_upload,json=abortIncompleteMultipartUpload,proto3" json:"abort_incomplete_multipart_upload,omitempty"`
}

func (x *LifecycleRule) Reset() {
	*x = LifecycleRule{}
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LifecycleRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleRule) ProtoMessage() {}

func (x *LifecycleRule) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleRule.ProtoReflect.Descriptor instead.
func (*LifecycleRule) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{1}
}

func (x *LifecycleRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LifecycleRule) GetStatus() LifecycleRule_Status {
	if x != nil {
		return x.Status
	}
	return LifecycleRule_STATUS_UNSPECIFIED
}

func (x *LifecycleRule) GetFilter() *LifecycleFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *LifecycleRule) GetExpiration() *LifecycleExpiration {
	if x != nil {
		return x.Expiration
	}
	return nil
}

func (x *LifecycleRule) GetNoncurrentVersionExpiration() *LifecycleNoncurrentVersionExpiration {
	if x != nil {
		return x.NoncurrentVersionExpiration
	}
	return nil
}

func (x *LifecycleRule) GetAbortIncompleteMultipartUpload() *LifecycleAbortIncompleteMultipartUpload {
	if x != nil {
		return x.AbortIncompleteMultipartUpload
	}
	return nil
}

type LifecycleFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prefix identifying one or more objects to which the rule applies.
	// If prefix is empty, the rule applies to all objects in the bucket.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Minimum object size to which the rule applies.
	ObjectSizeGreaterThanBytes int64 `protobuf:"varint,2,opt,name=object_size_greater_than_bytes,json=objectSizeGreaterThanBytes,proto3" json:"object_size_greater_than_bytes,omitempty"`
	// Maximum object size to which the rule applies.
	ObjectSizeLessThanBytes int64 `protobuf:"varint,3,opt,name=object_size_less_than_bytes,json=objectSizeLessThanBytes,proto3" json:"object_size_less_than_bytes,omitempty"`
}

func (x *LifecycleFilter) Reset() {
	*x = LifecycleFilter{}
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LifecycleFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleFilter) ProtoMessage() {}

func (x *LifecycleFilter) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleFilter.ProtoReflect.Descriptor instead.
func (*LifecycleFilter) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{2}
}

func (x *LifecycleFilter) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *LifecycleFilter) GetObjectSizeGreaterThanBytes() int64 {
	if x != nil {
		return x.ObjectSizeGreaterThanBytes
	}
	return 0
}

func (x *LifecycleFilter) GetObjectSizeLessThanBytes() int64 {
	if x != nil {
		return x.ObjectSizeLessThanBytes
	}
	return 0
}

type LifecycleExpiration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ExpiredWith:
	//
	//	*LifecycleExpiration_Date
	//	*LifecycleExpiration_Days
	ExpiredWith isLifecycleExpiration_ExpiredWith `protobuf_oneof:"expired_with"`
	// Indicates whether the system will remove a "delete marker" with no noncurrent versions.
	// If set to true, the "delete marker" will be permanently removed.
	// If set to false the policy takes no action.
	// This cannot be specified with Days or Date in a LifecycleExpiration Policy.
	ExpiredObjectDeleteMarker bool `protobuf:"varint,3,opt,name=expired_object_delete_marker,json=expiredObjectDeleteMarker,proto3" json:"expired_object_delete_marker,omitempty"`
}

func (x *LifecycleExpiration) Reset() {
	*x = LifecycleExpiration{}
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LifecycleExpiration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleExpiration) ProtoMessage() {}

func (x *LifecycleExpiration) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleExpiration.ProtoReflect.Descriptor instead.
func (*LifecycleExpiration) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{3}
}

func (m *LifecycleExpiration) GetExpiredWith() isLifecycleExpiration_ExpiredWith {
	if m != nil {
		return m.ExpiredWith
	}
	return nil
}

func (x *LifecycleExpiration) GetDate() *timestamppb.Timestamp {
	if x, ok := x.GetExpiredWith().(*LifecycleExpiration_Date); ok {
		return x.Date
	}
	return nil
}

func (x *LifecycleExpiration) GetDays() int32 {
	if x, ok := x.GetExpiredWith().(*LifecycleExpiration_Days); ok {
		return x.Days
	}
	return 0
}

func (x *LifecycleExpiration) GetExpiredObjectDeleteMarker() bool {
	if x != nil {
		return x.ExpiredObjectDeleteMarker
	}
	return false
}

type isLifecycleExpiration_ExpiredWith interface {
	isLifecycleExpiration_ExpiredWith()
}

type LifecycleExpiration_Date struct {
	// Indicates at what date the object will be deleted. The time is always midnight UTC.
	Date *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3,oneof"`
}

type LifecycleExpiration_Days struct {
	// Indicates the lifetime, in days, of the objects that are subject to the rule.
	// The value must be a non-zero positive integer.
	Days int32 `protobuf:"varint,2,opt,name=days,proto3,oneof"`
}

func (*LifecycleExpiration_Date) isLifecycleExpiration_ExpiredWith() {}

func (*LifecycleExpiration_Days) isLifecycleExpiration_ExpiredWith() {}

type LifecycleNoncurrentVersionExpiration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies how many noncurrent versions the system will retain.
	NewerNoncurrentVersions *int32 `protobuf:"varint,1,opt,name=newer_noncurrent_versions,json=newerNoncurrentVersions,proto3,oneof" json:"newer_noncurrent_versions,omitempty"`
	// Specifies the number of days an object is noncurrent before the system will expire it.
	NoncurrentDays int32 `protobuf:"varint,2,opt,name=noncurrent_days,json=noncurrentDays,proto3" json:"noncurrent_days,omitempty"`
}

func (x *LifecycleNoncurrentVersionExpiration) Reset() {
	*x = LifecycleNoncurrentVersionExpiration{}
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LifecycleNoncurrentVersionExpiration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleNoncurrentVersionExpiration) ProtoMessage() {}

func (x *LifecycleNoncurrentVersionExpiration) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleNoncurrentVersionExpiration.ProtoReflect.Descriptor instead.
func (*LifecycleNoncurrentVersionExpiration) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{4}
}

func (x *LifecycleNoncurrentVersionExpiration) GetNewerNoncurrentVersions() int32 {
	if x != nil && x.NewerNoncurrentVersions != nil {
		return *x.NewerNoncurrentVersions
	}
	return 0
}

func (x *LifecycleNoncurrentVersionExpiration) GetNoncurrentDays() int32 {
	if x != nil {
		return x.NoncurrentDays
	}
	return 0
}

type LifecycleAbortIncompleteMultipartUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the days since the initiation of an incomplete multipart upload that
	// the system will wait before permanently removing all parts of the upload.
	DaysAfterInitiation int32 `protobuf:"varint,1,opt,name=days_after_initiation,json=daysAfterInitiation,proto3" json:"days_after_initiation,omitempty"`
}

func (x *LifecycleAbortIncompleteMultipartUpload) Reset() {
	*x = LifecycleAbortIncompleteMultipartUpload{}
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LifecycleAbortIncompleteMultipartUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleAbortIncompleteMultipartUpload) ProtoMessage() {}

func (x *LifecycleAbortIncompleteMultipartUpload) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_lifecycle_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleAbortIncompleteMultipartUpload.ProtoReflect.Descriptor instead.
func (*LifecycleAbortIncompleteMultipartUpload) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_lifecycle_proto_rawDescGZIP(), []int{5}
}

func (x *LifecycleAbortIncompleteMultipartUpload) GetDaysAfterInitiation() int32 {
	if x != nil {
		return x.DaysAfterInitiation
	}
	return 0
}

var File_nebius_storage_v1_lifecycle_proto protoreflect.FileDescriptor

var file_nebius_storage_v1_lifecycle_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b,
	0x0a, 0x16, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x09, 0xba, 0x48, 0x06, 0x92, 0x01,
	0x03, 0x10, 0xe8, 0x07, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xbc, 0x04, 0x0a, 0x0d,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x07, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x46, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7b, 0x0a, 0x1d, 0x6e, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x6e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1b, 0x6e, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x85, 0x01, 0x0a, 0x21, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x41,
	0x62, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x1e, 0x61,
	0x62, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x3b, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x4c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x42, 0x0a, 0x1e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x68,
	0x61, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x72, 0x54, 0x68, 0x61, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x1b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x68, 0x61, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x17, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x54,
	0x68, 0x61, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x22, 0xae, 0x01, 0x0a, 0x24, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x19, 0x6e, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x17, 0x6e, 0x65, 0x77, 0x65, 0x72, 0x4e, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x6f,
	0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x79, 0x73, 0x42, 0x1c, 0x0a, 0x1a,
	0x5f, 0x6e, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5d, 0x0a, 0x27, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x64, 0x61, 0x79, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x5d, 0x0a, 0x18, 0x61, 0x69, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_storage_v1_lifecycle_proto_rawDescOnce sync.Once
	file_nebius_storage_v1_lifecycle_proto_rawDescData = file_nebius_storage_v1_lifecycle_proto_rawDesc
)

func file_nebius_storage_v1_lifecycle_proto_rawDescGZIP() []byte {
	file_nebius_storage_v1_lifecycle_proto_rawDescOnce.Do(func() {
		file_nebius_storage_v1_lifecycle_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_storage_v1_lifecycle_proto_rawDescData)
	})
	return file_nebius_storage_v1_lifecycle_proto_rawDescData
}

var file_nebius_storage_v1_lifecycle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_storage_v1_lifecycle_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nebius_storage_v1_lifecycle_proto_goTypes = []any{
	(LifecycleRule_Status)(0),                       // 0: nebius.storage.v1.LifecycleRule.Status
	(*LifecycleConfiguration)(nil),                  // 1: nebius.storage.v1.LifecycleConfiguration
	(*LifecycleRule)(nil),                           // 2: nebius.storage.v1.LifecycleRule
	(*LifecycleFilter)(nil),                         // 3: nebius.storage.v1.LifecycleFilter
	(*LifecycleExpiration)(nil),                     // 4: nebius.storage.v1.LifecycleExpiration
	(*LifecycleNoncurrentVersionExpiration)(nil),    // 5: nebius.storage.v1.LifecycleNoncurrentVersionExpiration
	(*LifecycleAbortIncompleteMultipartUpload)(nil), // 6: nebius.storage.v1.LifecycleAbortIncompleteMultipartUpload
	(*timestamppb.Timestamp)(nil),                   // 7: google.protobuf.Timestamp
}
var file_nebius_storage_v1_lifecycle_proto_depIdxs = []int32{
	2, // 0: nebius.storage.v1.LifecycleConfiguration.rules:type_name -> nebius.storage.v1.LifecycleRule
	0, // 1: nebius.storage.v1.LifecycleRule.status:type_name -> nebius.storage.v1.LifecycleRule.Status
	3, // 2: nebius.storage.v1.LifecycleRule.filter:type_name -> nebius.storage.v1.LifecycleFilter
	4, // 3: nebius.storage.v1.LifecycleRule.expiration:type_name -> nebius.storage.v1.LifecycleExpiration
	5, // 4: nebius.storage.v1.LifecycleRule.noncurrent_version_expiration:type_name -> nebius.storage.v1.LifecycleNoncurrentVersionExpiration
	6, // 5: nebius.storage.v1.LifecycleRule.abort_incomplete_multipart_upload:type_name -> nebius.storage.v1.LifecycleAbortIncompleteMultipartUpload
	7, // 6: nebius.storage.v1.LifecycleExpiration.date:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_nebius_storage_v1_lifecycle_proto_init() }
func file_nebius_storage_v1_lifecycle_proto_init() {
	if File_nebius_storage_v1_lifecycle_proto != nil {
		return
	}
	file_nebius_storage_v1_lifecycle_proto_msgTypes[3].OneofWrappers = []any{
		(*LifecycleExpiration_Date)(nil),
		(*LifecycleExpiration_Days)(nil),
	}
	file_nebius_storage_v1_lifecycle_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_storage_v1_lifecycle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_storage_v1_lifecycle_proto_goTypes,
		DependencyIndexes: file_nebius_storage_v1_lifecycle_proto_depIdxs,
		EnumInfos:         file_nebius_storage_v1_lifecycle_proto_enumTypes,
		MessageInfos:      file_nebius_storage_v1_lifecycle_proto_msgTypes,
	}.Build()
	File_nebius_storage_v1_lifecycle_proto = out.File
	file_nebius_storage_v1_lifecycle_proto_rawDesc = nil
	file_nebius_storage_v1_lifecycle_proto_goTypes = nil
	file_nebius_storage_v1_lifecycle_proto_depIdxs = nil
}
