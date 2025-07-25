// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/storage/v1/bucket.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type BucketStatus_State int32

const (
	BucketStatus_STATE_UNSPECIFIED BucketStatus_State = 0
	// Bucket is under creation and cannot be used yet.
	BucketStatus_CREATING BucketStatus_State = 1
	// Bucket is active and ready for usage.
	BucketStatus_ACTIVE BucketStatus_State = 2
	// Bucket is being updated.
	// It can be used, but some settings are being modified and you can observe their inconsistency.
	BucketStatus_UPDATING BucketStatus_State = 3
	// Bucket is scheduled for deletion.
	// It cannot be used in s3 api anymore.
	BucketStatus_SCHEDULED_FOR_DELETION BucketStatus_State = 4
)

// Enum value maps for BucketStatus_State.
var (
	BucketStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "UPDATING",
		4: "SCHEDULED_FOR_DELETION",
	}
	BucketStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED":      0,
		"CREATING":               1,
		"ACTIVE":                 2,
		"UPDATING":               3,
		"SCHEDULED_FOR_DELETION": 4,
	}
)

func (x BucketStatus_State) Enum() *BucketStatus_State {
	p := new(BucketStatus_State)
	*p = x
	return p
}

func (x BucketStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BucketStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_storage_v1_bucket_proto_enumTypes[0].Descriptor()
}

func (BucketStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_storage_v1_bucket_proto_enumTypes[0]
}

func (x BucketStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BucketStatus_State.Descriptor instead.
func (BucketStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_proto_rawDescGZIP(), []int{2, 0}
}

type BucketStatus_SuspensionState int32

const (
	BucketStatus_SUSPENSION_STATE_UNSPECIFIED BucketStatus_SuspensionState = 0
	BucketStatus_NOT_SUSPENDED                BucketStatus_SuspensionState = 1
	BucketStatus_SUSPENDED                    BucketStatus_SuspensionState = 2
)

// Enum value maps for BucketStatus_SuspensionState.
var (
	BucketStatus_SuspensionState_name = map[int32]string{
		0: "SUSPENSION_STATE_UNSPECIFIED",
		1: "NOT_SUSPENDED",
		2: "SUSPENDED",
	}
	BucketStatus_SuspensionState_value = map[string]int32{
		"SUSPENSION_STATE_UNSPECIFIED": 0,
		"NOT_SUSPENDED":                1,
		"SUSPENDED":                    2,
	}
)

func (x BucketStatus_SuspensionState) Enum() *BucketStatus_SuspensionState {
	p := new(BucketStatus_SuspensionState)
	*p = x
	return p
}

func (x BucketStatus_SuspensionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BucketStatus_SuspensionState) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_storage_v1_bucket_proto_enumTypes[1].Descriptor()
}

func (BucketStatus_SuspensionState) Type() protoreflect.EnumType {
	return &file_nebius_storage_v1_bucket_proto_enumTypes[1]
}

func (x BucketStatus_SuspensionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BucketStatus_SuspensionState.Descriptor instead.
func (BucketStatus_SuspensionState) EnumDescriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_proto_rawDescGZIP(), []int{2, 1}
}

type Bucket struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *BucketSpec            `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *BucketStatus          `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bucket) Reset() {
	*x = Bucket{}
	mi := &file_nebius_storage_v1_bucket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bucket) ProtoMessage() {}

func (x *Bucket) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_bucket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bucket.ProtoReflect.Descriptor instead.
func (*Bucket) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_proto_rawDescGZIP(), []int{0}
}

func (x *Bucket) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Bucket) GetSpec() *BucketSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Bucket) GetStatus() *BucketStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type BucketSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Supports transitions:
	//   - disabled -> enabled
	//   - disabled -> suspended
	//   - enabled <-> suspended
	VersioningPolicy VersioningPolicy `protobuf:"varint,2,opt,name=versioning_policy,json=versioningPolicy,proto3,enum=nebius.storage.v1.VersioningPolicy" json:"versioning_policy,omitempty"`
	// Maximum bucket size.
	// Zero means unlimited.
	// Actual limit can be lower if customer doesn't have enough quota.
	// Real bucket size can go a little higher if customer writes too fast.
	MaxSizeBytes           int64                   `protobuf:"varint,4,opt,name=max_size_bytes,json=maxSizeBytes,proto3" json:"max_size_bytes,omitempty"`
	LifecycleConfiguration *LifecycleConfiguration `protobuf:"bytes,5,opt,name=lifecycle_configuration,json=lifecycleConfiguration,proto3" json:"lifecycle_configuration,omitempty"`
	// Storage class to use by default for uploads to the bucket. It may be overridden by `x-amz-storage-class` header.
	// If not set - STANDARD is used as a default storage class.
	DefaultStorageClass StorageClass `protobuf:"varint,9,opt,name=default_storage_class,json=defaultStorageClass,proto3,enum=nebius.storage.v1.StorageClass" json:"default_storage_class,omitempty"`
	// Storage class to override any other storage class of uploading objects. It overrides the storage class regardless
	// of how the original storage class was specified - either the default storage class
	// or the one provided via the `x-amz-storage-class` header.
	OverrideStorageClass StorageClass `protobuf:"varint,10,opt,name=override_storage_class,json=overrideStorageClass,proto3,enum=nebius.storage.v1.StorageClass" json:"override_storage_class,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *BucketSpec) Reset() {
	*x = BucketSpec{}
	mi := &file_nebius_storage_v1_bucket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BucketSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketSpec) ProtoMessage() {}

func (x *BucketSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_bucket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketSpec.ProtoReflect.Descriptor instead.
func (*BucketSpec) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_proto_rawDescGZIP(), []int{1}
}

func (x *BucketSpec) GetVersioningPolicy() VersioningPolicy {
	if x != nil {
		return x.VersioningPolicy
	}
	return VersioningPolicy_VERSIONING_POLICY_UNSPECIFIED
}

func (x *BucketSpec) GetMaxSizeBytes() int64 {
	if x != nil {
		return x.MaxSizeBytes
	}
	return 0
}

func (x *BucketSpec) GetLifecycleConfiguration() *LifecycleConfiguration {
	if x != nil {
		return x.LifecycleConfiguration
	}
	return nil
}

func (x *BucketSpec) GetDefaultStorageClass() StorageClass {
	if x != nil {
		return x.DefaultStorageClass
	}
	return StorageClass_STORAGE_CLASS_UNSPECIFIED
}

func (x *BucketSpec) GetOverrideStorageClass() StorageClass {
	if x != nil {
		return x.OverrideStorageClass
	}
	return StorageClass_STORAGE_CLASS_UNSPECIFIED
}

type BucketStatus struct {
	state           protoimpl.MessageState       `protogen:"open.v1"`
	Counters        []*BucketCounters            `protobuf:"bytes,1,rep,name=counters,proto3" json:"counters,omitempty"`
	State           BucketStatus_State           `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.storage.v1.BucketStatus_State" json:"state,omitempty"`
	SuspensionState BucketStatus_SuspensionState `protobuf:"varint,3,opt,name=suspension_state,json=suspensionState,proto3,enum=nebius.storage.v1.BucketStatus_SuspensionState" json:"suspension_state,omitempty"`
	// The time when the bucket was deleted (or scheduled for deletion).
	// It resets to null if the bucket is undeleted.
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	// The time when the bucket will be automatically purged in case it was soft-deleted.
	PurgeAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=purge_at,json=purgeAt,proto3" json:"purge_at,omitempty"`
	// The domain of the endpoint where the bucket can be accessed. It omits the scheme (HTTPS) and the port (443)
	// and contains only the FQDN address.
	DomainName string `protobuf:"bytes,6,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	// The name of the region where the bucket is located for use with S3 clients, i.e. "eu-west1".
	Region        string `protobuf:"bytes,8,opt,name=region,proto3" json:"region,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BucketStatus) Reset() {
	*x = BucketStatus{}
	mi := &file_nebius_storage_v1_bucket_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BucketStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketStatus) ProtoMessage() {}

func (x *BucketStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_bucket_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketStatus.ProtoReflect.Descriptor instead.
func (*BucketStatus) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_proto_rawDescGZIP(), []int{2}
}

func (x *BucketStatus) GetCounters() []*BucketCounters {
	if x != nil {
		return x.Counters
	}
	return nil
}

func (x *BucketStatus) GetState() BucketStatus_State {
	if x != nil {
		return x.State
	}
	return BucketStatus_STATE_UNSPECIFIED
}

func (x *BucketStatus) GetSuspensionState() BucketStatus_SuspensionState {
	if x != nil {
		return x.SuspensionState
	}
	return BucketStatus_SUSPENSION_STATE_UNSPECIFIED
}

func (x *BucketStatus) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *BucketStatus) GetPurgeAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PurgeAt
	}
	return nil
}

func (x *BucketStatus) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *BucketStatus) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_nebius_storage_v1_bucket_proto protoreflect.FileDescriptor

var file_nebius_storage_v1_bucket_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x3a, 0x04, 0xba, 0x4a, 0x01, 0x03, 0x22, 0x9a, 0x03, 0x0a, 0x0a, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x56, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x07, 0x52, 0x10,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x53, 0x69, 0x7a,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x17, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79,
	0x63, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x16, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x15, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x13, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x55, 0x0a, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x14, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x22, 0xcc, 0x04, 0x0a, 0x0c, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x5a, 0x0a, 0x10, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53,
	0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f,
	0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x75,
	0x72, 0x67, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x70, 0x75, 0x72, 0x67, 0x65, 0x41,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x5f,
	0x46, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x22, 0x55,
	0x0a, 0x0f, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45,
	0x4e, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e,
	0x44, 0x45, 0x44, 0x10, 0x02, 0x42, 0x5a, 0x0a, 0x18, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_storage_v1_bucket_proto_rawDescOnce sync.Once
	file_nebius_storage_v1_bucket_proto_rawDescData []byte
)

func file_nebius_storage_v1_bucket_proto_rawDescGZIP() []byte {
	file_nebius_storage_v1_bucket_proto_rawDescOnce.Do(func() {
		file_nebius_storage_v1_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_storage_v1_bucket_proto_rawDesc), len(file_nebius_storage_v1_bucket_proto_rawDesc)))
	})
	return file_nebius_storage_v1_bucket_proto_rawDescData
}

var file_nebius_storage_v1_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nebius_storage_v1_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_storage_v1_bucket_proto_goTypes = []any{
	(BucketStatus_State)(0),           // 0: nebius.storage.v1.BucketStatus.State
	(BucketStatus_SuspensionState)(0), // 1: nebius.storage.v1.BucketStatus.SuspensionState
	(*Bucket)(nil),                    // 2: nebius.storage.v1.Bucket
	(*BucketSpec)(nil),                // 3: nebius.storage.v1.BucketSpec
	(*BucketStatus)(nil),              // 4: nebius.storage.v1.BucketStatus
	(*v1.ResourceMetadata)(nil),       // 5: nebius.common.v1.ResourceMetadata
	(VersioningPolicy)(0),             // 6: nebius.storage.v1.VersioningPolicy
	(*LifecycleConfiguration)(nil),    // 7: nebius.storage.v1.LifecycleConfiguration
	(StorageClass)(0),                 // 8: nebius.storage.v1.StorageClass
	(*BucketCounters)(nil),            // 9: nebius.storage.v1.BucketCounters
	(*timestamppb.Timestamp)(nil),     // 10: google.protobuf.Timestamp
}
var file_nebius_storage_v1_bucket_proto_depIdxs = []int32{
	5,  // 0: nebius.storage.v1.Bucket.metadata:type_name -> nebius.common.v1.ResourceMetadata
	3,  // 1: nebius.storage.v1.Bucket.spec:type_name -> nebius.storage.v1.BucketSpec
	4,  // 2: nebius.storage.v1.Bucket.status:type_name -> nebius.storage.v1.BucketStatus
	6,  // 3: nebius.storage.v1.BucketSpec.versioning_policy:type_name -> nebius.storage.v1.VersioningPolicy
	7,  // 4: nebius.storage.v1.BucketSpec.lifecycle_configuration:type_name -> nebius.storage.v1.LifecycleConfiguration
	8,  // 5: nebius.storage.v1.BucketSpec.default_storage_class:type_name -> nebius.storage.v1.StorageClass
	8,  // 6: nebius.storage.v1.BucketSpec.override_storage_class:type_name -> nebius.storage.v1.StorageClass
	9,  // 7: nebius.storage.v1.BucketStatus.counters:type_name -> nebius.storage.v1.BucketCounters
	0,  // 8: nebius.storage.v1.BucketStatus.state:type_name -> nebius.storage.v1.BucketStatus.State
	1,  // 9: nebius.storage.v1.BucketStatus.suspension_state:type_name -> nebius.storage.v1.BucketStatus.SuspensionState
	10, // 10: nebius.storage.v1.BucketStatus.deleted_at:type_name -> google.protobuf.Timestamp
	10, // 11: nebius.storage.v1.BucketStatus.purge_at:type_name -> google.protobuf.Timestamp
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_nebius_storage_v1_bucket_proto_init() }
func file_nebius_storage_v1_bucket_proto_init() {
	if File_nebius_storage_v1_bucket_proto != nil {
		return
	}
	file_nebius_storage_v1_base_proto_init()
	file_nebius_storage_v1_bucket_counters_proto_init()
	file_nebius_storage_v1_lifecycle_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_storage_v1_bucket_proto_rawDesc), len(file_nebius_storage_v1_bucket_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_storage_v1_bucket_proto_goTypes,
		DependencyIndexes: file_nebius_storage_v1_bucket_proto_depIdxs,
		EnumInfos:         file_nebius_storage_v1_bucket_proto_enumTypes,
		MessageInfos:      file_nebius_storage_v1_bucket_proto_msgTypes,
	}.Build()
	File_nebius_storage_v1_bucket_proto = out.File
	file_nebius_storage_v1_bucket_proto_goTypes = nil
	file_nebius_storage_v1_bucket_proto_depIdxs = nil
}
