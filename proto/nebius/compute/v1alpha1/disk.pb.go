// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// nebius/compute/v1alpha1/disk.proto is a deprecated file.

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
type DiskSpec_DiskType int32

const (
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskSpec_UNSPECIFIED DiskSpec_DiskType = 0
	// the list of available types will be clarified later, it is not final version
	//
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskSpec_NETWORK_SSD DiskSpec_DiskType = 1
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskSpec_NETWORK_HDD DiskSpec_DiskType = 2
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskSpec_NETWORK_SSD_NON_REPLICATED DiskSpec_DiskType = 3
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskSpec_NETWORK_SSD_IO_M3 DiskSpec_DiskType = 4
)

// Enum value maps for DiskSpec_DiskType.
var (
	DiskSpec_DiskType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "NETWORK_SSD",
		2: "NETWORK_HDD",
		3: "NETWORK_SSD_NON_REPLICATED",
		4: "NETWORK_SSD_IO_M3",
	}
	DiskSpec_DiskType_value = map[string]int32{
		"UNSPECIFIED":                0,
		"NETWORK_SSD":                1,
		"NETWORK_HDD":                2,
		"NETWORK_SSD_NON_REPLICATED": 3,
		"NETWORK_SSD_IO_M3":          4,
	}
)

func (x DiskSpec_DiskType) Enum() *DiskSpec_DiskType {
	p := new(DiskSpec_DiskType)
	*p = x
	return p
}

func (x DiskSpec_DiskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiskSpec_DiskType) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_compute_v1alpha1_disk_proto_enumTypes[0].Descriptor()
}

func (DiskSpec_DiskType) Type() protoreflect.EnumType {
	return &file_nebius_compute_v1alpha1_disk_proto_enumTypes[0]
}

func (x DiskSpec_DiskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiskSpec_DiskType.Descriptor instead.
func (DiskSpec_DiskType) EnumDescriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP(), []int{1, 0}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
type DiskStatus_State int32

const (
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskStatus_UNSPECIFIED DiskStatus_State = 0
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskStatus_CREATING DiskStatus_State = 1
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskStatus_READY DiskStatus_State = 2
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskStatus_UPDATING DiskStatus_State = 3
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskStatus_DELETING DiskStatus_State = 4
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	DiskStatus_ERROR DiskStatus_State = 5
)

// Enum value maps for DiskStatus_State.
var (
	DiskStatus_State_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "UPDATING",
		4: "DELETING",
		5: "ERROR",
	}
	DiskStatus_State_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CREATING":    1,
		"READY":       2,
		"UPDATING":    3,
		"DELETING":    4,
		"ERROR":       5,
	}
)

func (x DiskStatus_State) Enum() *DiskStatus_State {
	p := new(DiskStatus_State)
	*p = x
	return p
}

func (x DiskStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiskStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_compute_v1alpha1_disk_proto_enumTypes[1].Descriptor()
}

func (DiskStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_compute_v1alpha1_disk_proto_enumTypes[1]
}

func (x DiskStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiskStatus_State.Descriptor instead.
func (DiskStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP(), []int{3, 0}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
type Disk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	Spec *DiskSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	Status *DiskStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Disk) Reset() {
	*x = Disk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Disk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Disk) ProtoMessage() {}

func (x *Disk) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Disk.ProtoReflect.Descriptor instead.
func (*Disk) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP(), []int{0}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *Disk) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *Disk) GetSpec() *DiskSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *Disk) GetStatus() *DiskStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
type DiskSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Size:
	//
	//	*DiskSpec_SizeBytes
	//	*DiskSpec_SizeKibibytes
	//	*DiskSpec_SizeMebibytes
	//	*DiskSpec_SizeGibibytes
	Size isDiskSpec_Size `protobuf_oneof:"size"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	BlockSizeBytes int64 `protobuf:"varint,5,opt,name=block_size_bytes,json=blockSizeBytes,proto3" json:"block_size_bytes,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	Type DiskSpec_DiskType `protobuf:"varint,6,opt,name=type,proto3,enum=nebius.compute.v1alpha1.DiskSpec_DiskType" json:"type,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	PlacementPolicy *DiskPlacementPolicy `protobuf:"bytes,7,opt,name=placement_policy,json=placementPolicy,proto3" json:"placement_policy,omitempty"`
	// Types that are assignable to Source:
	//
	//	*DiskSpec_SourceImageId
	//	*DiskSpec_SourceImageFamily
	Source isDiskSpec_Source `protobuf_oneof:"source"`
}

func (x *DiskSpec) Reset() {
	*x = DiskSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskSpec) ProtoMessage() {}

func (x *DiskSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskSpec.ProtoReflect.Descriptor instead.
func (*DiskSpec) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP(), []int{1}
}

func (m *DiskSpec) GetSize() isDiskSpec_Size {
	if m != nil {
		return m.Size
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetSizeBytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeBytes); ok {
		return x.SizeBytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetSizeKibibytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeKibibytes); ok {
		return x.SizeKibibytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetSizeMebibytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeMebibytes); ok {
		return x.SizeMebibytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetSizeGibibytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeGibibytes); ok {
		return x.SizeGibibytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetBlockSizeBytes() int64 {
	if x != nil {
		return x.BlockSizeBytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetType() DiskSpec_DiskType {
	if x != nil {
		return x.Type
	}
	return DiskSpec_UNSPECIFIED
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetPlacementPolicy() *DiskPlacementPolicy {
	if x != nil {
		return x.PlacementPolicy
	}
	return nil
}

func (m *DiskSpec) GetSource() isDiskSpec_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetSourceImageId() string {
	if x, ok := x.GetSource().(*DiskSpec_SourceImageId); ok {
		return x.SourceImageId
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskSpec) GetSourceImageFamily() string {
	if x, ok := x.GetSource().(*DiskSpec_SourceImageFamily); ok {
		return x.SourceImageFamily
	}
	return ""
}

type isDiskSpec_Size interface {
	isDiskSpec_Size()
}

type DiskSpec_SizeBytes struct {
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SizeBytes int64 `protobuf:"varint,1,opt,name=size_bytes,json=sizeBytes,proto3,oneof"`
}

type DiskSpec_SizeKibibytes struct {
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SizeKibibytes int64 `protobuf:"varint,2,opt,name=size_kibibytes,json=sizeKibibytes,proto3,oneof"`
}

type DiskSpec_SizeMebibytes struct {
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SizeMebibytes int64 `protobuf:"varint,3,opt,name=size_mebibytes,json=sizeMebibytes,proto3,oneof"`
}

type DiskSpec_SizeGibibytes struct {
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SizeGibibytes int64 `protobuf:"varint,4,opt,name=size_gibibytes,json=sizeGibibytes,proto3,oneof"`
}

func (*DiskSpec_SizeBytes) isDiskSpec_Size() {}

func (*DiskSpec_SizeKibibytes) isDiskSpec_Size() {}

func (*DiskSpec_SizeMebibytes) isDiskSpec_Size() {}

func (*DiskSpec_SizeGibibytes) isDiskSpec_Size() {}

type isDiskSpec_Source interface {
	isDiskSpec_Source()
}

type DiskSpec_SourceImageId struct {
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SourceImageId string `protobuf:"bytes,8,opt,name=source_image_id,json=sourceImageId,proto3,oneof"`
}

type DiskSpec_SourceImageFamily struct {
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SourceImageFamily string `protobuf:"bytes,9,opt,name=source_image_family,json=sourceImageFamily,proto3,oneof"`
}

func (*DiskSpec_SourceImageId) isDiskSpec_Source() {}

func (*DiskSpec_SourceImageFamily) isDiskSpec_Source() {}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
type DiskPlacementPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	PlacementGroupId string `protobuf:"bytes,1,opt,name=placement_group_id,json=placementGroupId,proto3" json:"placement_group_id,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	PlacementGroupPartition int64 `protobuf:"varint,2,opt,name=placement_group_partition,json=placementGroupPartition,proto3" json:"placement_group_partition,omitempty"`
}

func (x *DiskPlacementPolicy) Reset() {
	*x = DiskPlacementPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskPlacementPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskPlacementPolicy) ProtoMessage() {}

func (x *DiskPlacementPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskPlacementPolicy.ProtoReflect.Descriptor instead.
func (*DiskPlacementPolicy) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP(), []int{2}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskPlacementPolicy) GetPlacementGroupId() string {
	if x != nil {
		return x.PlacementGroupId
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskPlacementPolicy) GetPlacementGroupPartition() int64 {
	if x != nil {
		return x.PlacementGroupPartition
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
type DiskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	State DiskStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.compute.v1alpha1.DiskStatus_State" json:"state,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	StateDescription string `protobuf:"bytes,2,opt,name=state_description,json=stateDescription,proto3" json:"state_description,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	ReadWriteAttachment string `protobuf:"bytes,3,opt,name=read_write_attachment,json=readWriteAttachment,proto3" json:"read_write_attachment,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	ReadOnlyAttachments []string `protobuf:"bytes,4,rep,name=read_only_attachments,json=readOnlyAttachments,proto3" json:"read_only_attachments,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SourceImageId string `protobuf:"bytes,5,opt,name=source_image_id,json=sourceImageId,proto3" json:"source_image_id,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	SizeBytes int64 `protobuf:"varint,6,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// Indicates whether there is an ongoing operation
	//
	// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
	Reconciling bool `protobuf:"varint,7,opt,name=reconciling,proto3" json:"reconciling,omitempty"`
}

func (x *DiskStatus) Reset() {
	*x = DiskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskStatus) ProtoMessage() {}

func (x *DiskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_disk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskStatus.ProtoReflect.Descriptor instead.
func (*DiskStatus) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP(), []int{3}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetState() DiskStatus_State {
	if x != nil {
		return x.State
	}
	return DiskStatus_UNSPECIFIED
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetStateDescription() string {
	if x != nil {
		return x.StateDescription
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetReadWriteAttachment() string {
	if x != nil {
		return x.ReadWriteAttachment
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetReadOnlyAttachments() []string {
	if x != nil {
		return x.ReadOnlyAttachments
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetSourceImageId() string {
	if x != nil {
		return x.SourceImageId
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/disk.proto is marked as deprecated.
func (x *DiskStatus) GetReconciling() bool {
	if x != nil {
		return x.Reconciling
	}
	return false
}

var File_nebius_compute_v1alpha1_disk_proto protoreflect.FileDescriptor

var file_nebius_compute_v1alpha1_disk_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x3e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x90, 0x05, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x25, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x48, 0x00, 0x52, 0x09, 0x73, 0x69, 0x7a,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6b,
	0x69, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04,
	0xba, 0x4a, 0x01, 0x02, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x4b, 0x69, 0x62, 0x69,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x65,
	0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xba,
	0x4a, 0x01, 0x02, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x62, 0x69, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x69, 0x62,
	0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xba, 0x4a,
	0x01, 0x02, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x47, 0x69, 0x62, 0x69, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x04, 0xba,
	0x4a, 0x01, 0x02, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b,
	0x53, 0x70, 0x65, 0x63, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x5d, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x0f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e,
	0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x48, 0x01, 0x52,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01,
	0x02, 0x48, 0x01, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x74, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53,
	0x53, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f,
	0x48, 0x44, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x53, 0x53, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x53, 0x53, 0x44, 0x5f, 0x49, 0x4f, 0x5f, 0x4d, 0x33, 0x10, 0x04, 0x42, 0x0d, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x7f, 0x0a, 0x13, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa5, 0x03, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x22, 0x58, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0x67,
	0x0a, 0x1e, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xb8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_compute_v1alpha1_disk_proto_rawDescOnce sync.Once
	file_nebius_compute_v1alpha1_disk_proto_rawDescData = file_nebius_compute_v1alpha1_disk_proto_rawDesc
)

func file_nebius_compute_v1alpha1_disk_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1alpha1_disk_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1alpha1_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_compute_v1alpha1_disk_proto_rawDescData)
	})
	return file_nebius_compute_v1alpha1_disk_proto_rawDescData
}

var file_nebius_compute_v1alpha1_disk_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nebius_compute_v1alpha1_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_compute_v1alpha1_disk_proto_goTypes = []any{
	(DiskSpec_DiskType)(0),      // 0: nebius.compute.v1alpha1.DiskSpec.DiskType
	(DiskStatus_State)(0),       // 1: nebius.compute.v1alpha1.DiskStatus.State
	(*Disk)(nil),                // 2: nebius.compute.v1alpha1.Disk
	(*DiskSpec)(nil),            // 3: nebius.compute.v1alpha1.DiskSpec
	(*DiskPlacementPolicy)(nil), // 4: nebius.compute.v1alpha1.DiskPlacementPolicy
	(*DiskStatus)(nil),          // 5: nebius.compute.v1alpha1.DiskStatus
	(*v1.ResourceMetadata)(nil), // 6: nebius.common.v1.ResourceMetadata
}
var file_nebius_compute_v1alpha1_disk_proto_depIdxs = []int32{
	6, // 0: nebius.compute.v1alpha1.Disk.metadata:type_name -> nebius.common.v1.ResourceMetadata
	3, // 1: nebius.compute.v1alpha1.Disk.spec:type_name -> nebius.compute.v1alpha1.DiskSpec
	5, // 2: nebius.compute.v1alpha1.Disk.status:type_name -> nebius.compute.v1alpha1.DiskStatus
	0, // 3: nebius.compute.v1alpha1.DiskSpec.type:type_name -> nebius.compute.v1alpha1.DiskSpec.DiskType
	4, // 4: nebius.compute.v1alpha1.DiskSpec.placement_policy:type_name -> nebius.compute.v1alpha1.DiskPlacementPolicy
	1, // 5: nebius.compute.v1alpha1.DiskStatus.state:type_name -> nebius.compute.v1alpha1.DiskStatus.State
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1alpha1_disk_proto_init() }
func file_nebius_compute_v1alpha1_disk_proto_init() {
	if File_nebius_compute_v1alpha1_disk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nebius_compute_v1alpha1_disk_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Disk); i {
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
		file_nebius_compute_v1alpha1_disk_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DiskSpec); i {
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
		file_nebius_compute_v1alpha1_disk_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DiskPlacementPolicy); i {
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
		file_nebius_compute_v1alpha1_disk_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DiskStatus); i {
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
	file_nebius_compute_v1alpha1_disk_proto_msgTypes[1].OneofWrappers = []any{
		(*DiskSpec_SizeBytes)(nil),
		(*DiskSpec_SizeKibibytes)(nil),
		(*DiskSpec_SizeMebibytes)(nil),
		(*DiskSpec_SizeGibibytes)(nil),
		(*DiskSpec_SourceImageId)(nil),
		(*DiskSpec_SourceImageFamily)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_compute_v1alpha1_disk_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_compute_v1alpha1_disk_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1alpha1_disk_proto_depIdxs,
		EnumInfos:         file_nebius_compute_v1alpha1_disk_proto_enumTypes,
		MessageInfos:      file_nebius_compute_v1alpha1_disk_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1alpha1_disk_proto = out.File
	file_nebius_compute_v1alpha1_disk_proto_rawDesc = nil
	file_nebius_compute_v1alpha1_disk_proto_goTypes = nil
	file_nebius_compute_v1alpha1_disk_proto_depIdxs = nil
}
