// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// nebius/compute/v1alpha1/image.proto is a deprecated file.

package v1alpha1

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

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
type ImageStatus_State int32

const (
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageStatus_UNSPECIFIED ImageStatus_State = 0
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageStatus_CREATING ImageStatus_State = 1
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageStatus_READY ImageStatus_State = 2
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageStatus_UPDATING ImageStatus_State = 3
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageStatus_DELETING ImageStatus_State = 4
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageStatus_ERROR ImageStatus_State = 5
)

// Enum value maps for ImageStatus_State.
var (
	ImageStatus_State_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "UPDATING",
		4: "DELETING",
		5: "ERROR",
	}
	ImageStatus_State_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CREATING":    1,
		"READY":       2,
		"UPDATING":    3,
		"DELETING":    4,
		"ERROR":       5,
	}
)

func (x ImageStatus_State) Enum() *ImageStatus_State {
	p := new(ImageStatus_State)
	*p = x
	return p
}

func (x ImageStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_compute_v1alpha1_image_proto_enumTypes[0].Descriptor()
}

func (ImageStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_compute_v1alpha1_image_proto_enumTypes[0]
}

func (x ImageStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageStatus_State.Descriptor instead.
func (ImageStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_image_proto_rawDescGZIP(), []int{2, 0}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	Spec *ImageSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	Status *ImageStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_nebius_compute_v1alpha1_image_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_image_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_image_proto_rawDescGZIP(), []int{0}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *Image) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *Image) GetSpec() *ImageSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *Image) GetStatus() *ImageStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
type ImageSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	Description *string `protobuf:"bytes,1,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	ImageFamily string `protobuf:"bytes,2,opt,name=image_family,json=imageFamily,proto3" json:"image_family,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ImageSpec) Reset() {
	*x = ImageSpec{}
	mi := &file_nebius_compute_v1alpha1_image_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageSpec) ProtoMessage() {}

func (x *ImageSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_image_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageSpec.ProtoReflect.Descriptor instead.
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_image_proto_rawDescGZIP(), []int{1}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageSpec) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageSpec) GetImageFamily() string {
	if x != nil {
		return x.ImageFamily
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
type ImageStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	State ImageStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.compute.v1alpha1.ImageStatus_State" json:"state,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	StateDescription string `protobuf:"bytes,2,opt,name=state_description,json=stateDescription,proto3" json:"state_description,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	StorageSizeBytes int64 `protobuf:"varint,3,opt,name=storage_size_bytes,json=storageSizeBytes,proto3" json:"storage_size_bytes,omitempty"`
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	MinDiskSizeBytes int64 `protobuf:"varint,4,opt,name=min_disk_size_bytes,json=minDiskSizeBytes,proto3" json:"min_disk_size_bytes,omitempty"`
	// Indicates whether there is an ongoing operation
	//
	// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
	Reconciling bool `protobuf:"varint,5,opt,name=reconciling,proto3" json:"reconciling,omitempty"`
}

func (x *ImageStatus) Reset() {
	*x = ImageStatus{}
	mi := &file_nebius_compute_v1alpha1_image_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageStatus) ProtoMessage() {}

func (x *ImageStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1alpha1_image_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageStatus.ProtoReflect.Descriptor instead.
func (*ImageStatus) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1alpha1_image_proto_rawDescGZIP(), []int{2}
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageStatus) GetState() ImageStatus_State {
	if x != nil {
		return x.State
	}
	return ImageStatus_UNSPECIFIED
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageStatus) GetStateDescription() string {
	if x != nil {
		return x.StateDescription
	}
	return ""
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageStatus) GetStorageSizeBytes() int64 {
	if x != nil {
		return x.StorageSizeBytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageStatus) GetMinDiskSizeBytes() int64 {
	if x != nil {
		return x.MinDiskSizeBytes
	}
	return 0
}

// Deprecated: The entire proto file nebius/compute/v1alpha1/image.proto is marked as deprecated.
func (x *ImageStatus) GetReconciling() bool {
	if x != nil {
		return x.Reconciling
	}
	return false
}

var File_nebius_compute_v1alpha1_image_proto protoreflect.FileDescriptor

var file_nebius_compute_v1alpha1_image_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a,
	0x01, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02,
	0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1e, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xba, 0x4a, 0x01, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd5, 0x02,
	0x0a, 0x0b, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x69,
	0x6e, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x6b,
	0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x58, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x05, 0x42, 0x68, 0x0a, 0x1e, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xb8, 0x01, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_compute_v1alpha1_image_proto_rawDescOnce sync.Once
	file_nebius_compute_v1alpha1_image_proto_rawDescData = file_nebius_compute_v1alpha1_image_proto_rawDesc
)

func file_nebius_compute_v1alpha1_image_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1alpha1_image_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1alpha1_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_compute_v1alpha1_image_proto_rawDescData)
	})
	return file_nebius_compute_v1alpha1_image_proto_rawDescData
}

var file_nebius_compute_v1alpha1_image_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_compute_v1alpha1_image_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_compute_v1alpha1_image_proto_goTypes = []any{
	(ImageStatus_State)(0),      // 0: nebius.compute.v1alpha1.ImageStatus.State
	(*Image)(nil),               // 1: nebius.compute.v1alpha1.Image
	(*ImageSpec)(nil),           // 2: nebius.compute.v1alpha1.ImageSpec
	(*ImageStatus)(nil),         // 3: nebius.compute.v1alpha1.ImageStatus
	(*v1.ResourceMetadata)(nil), // 4: nebius.common.v1.ResourceMetadata
}
var file_nebius_compute_v1alpha1_image_proto_depIdxs = []int32{
	4, // 0: nebius.compute.v1alpha1.Image.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2, // 1: nebius.compute.v1alpha1.Image.spec:type_name -> nebius.compute.v1alpha1.ImageSpec
	3, // 2: nebius.compute.v1alpha1.Image.status:type_name -> nebius.compute.v1alpha1.ImageStatus
	0, // 3: nebius.compute.v1alpha1.ImageStatus.state:type_name -> nebius.compute.v1alpha1.ImageStatus.State
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1alpha1_image_proto_init() }
func file_nebius_compute_v1alpha1_image_proto_init() {
	if File_nebius_compute_v1alpha1_image_proto != nil {
		return
	}
	file_nebius_compute_v1alpha1_image_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_compute_v1alpha1_image_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_compute_v1alpha1_image_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1alpha1_image_proto_depIdxs,
		EnumInfos:         file_nebius_compute_v1alpha1_image_proto_enumTypes,
		MessageInfos:      file_nebius_compute_v1alpha1_image_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1alpha1_image_proto = out.File
	file_nebius_compute_v1alpha1_image_proto_rawDesc = nil
	file_nebius_compute_v1alpha1_image_proto_goTypes = nil
	file_nebius_compute_v1alpha1_image_proto_depIdxs = nil
}
