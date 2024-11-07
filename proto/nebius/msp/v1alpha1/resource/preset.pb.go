// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/msp/v1alpha1/resource/preset.proto

package resource

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Preset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *PresetSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Preset) Reset() {
	*x = Preset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Preset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Preset) ProtoMessage() {}

func (x *Preset) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Preset.ProtoReflect.Descriptor instead.
func (*Preset) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_proto_rawDescGZIP(), []int{0}
}

func (x *Preset) GetSpec() *PresetSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type PresetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flavor *FlavorSpec `protobuf:"bytes,1,opt,name=flavor,proto3" json:"flavor,omitempty"`
	Hosts  *Host       `protobuf:"bytes,2,opt,name=hosts,proto3" json:"hosts,omitempty"`
	Disk   *Disk       `protobuf:"bytes,3,opt,name=disk,proto3" json:"disk,omitempty"`
	Role   string      `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *PresetSpec) Reset() {
	*x = PresetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresetSpec) ProtoMessage() {}

func (x *PresetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresetSpec.ProtoReflect.Descriptor instead.
func (*PresetSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_proto_rawDescGZIP(), []int{1}
}

func (x *PresetSpec) GetFlavor() *FlavorSpec {
	if x != nil {
		return x.Flavor
	}
	return nil
}

func (x *PresetSpec) GetHosts() *Host {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *PresetSpec) GetDisk() *Disk {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *PresetSpec) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type FlavorSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu    *CpuSpec    `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory *MemorySpec `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
}

func (x *FlavorSpec) Reset() {
	*x = FlavorSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlavorSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlavorSpec) ProtoMessage() {}

func (x *FlavorSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlavorSpec.ProtoReflect.Descriptor instead.
func (*FlavorSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_proto_rawDescGZIP(), []int{2}
}

func (x *FlavorSpec) GetCpu() *CpuSpec {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *FlavorSpec) GetMemory() *MemorySpec {
	if x != nil {
		return x.Memory
	}
	return nil
}

type CpuSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count      int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Generation int32 `protobuf:"varint,2,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *CpuSpec) Reset() {
	*x = CpuSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuSpec) ProtoMessage() {}

func (x *CpuSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuSpec.ProtoReflect.Descriptor instead.
func (*CpuSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_proto_rawDescGZIP(), []int{3}
}

func (x *CpuSpec) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CpuSpec) GetGeneration() int32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

type MemorySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LimitGibibytes int64 `protobuf:"varint,1,opt,name=limit_gibibytes,json=limitGibibytes,proto3" json:"limit_gibibytes,omitempty"`
}

func (x *MemorySpec) Reset() {
	*x = MemorySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemorySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemorySpec) ProtoMessage() {}

func (x *MemorySpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemorySpec.ProtoReflect.Descriptor instead.
func (*MemorySpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_proto_rawDescGZIP(), []int{4}
}

func (x *MemorySpec) GetLimitGibibytes() int64 {
	if x != nil {
		return x.LimitGibibytes
	}
	return 0
}

var File_nebius_msp_v1alpha1_resource_preset_proto protoreflect.FileDescriptor

var file_nebius_msp_v1alpha1_resource_preset_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d,
	0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x44, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x22, 0xe4, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x48, 0x0a, 0x06, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x05,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x1a,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0a, 0x46,
	0x6c, 0x61, 0x76, 0x6f, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3f, 0x0a, 0x03, 0x63, 0x70, 0x75,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x70, 0x75, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x48, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x1c, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x2f, 0x0a, 0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x67, 0x69, 0x62,
	0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x47, 0x69, 0x62, 0x69, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x42, 0x70, 0x0a, 0x23, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0b, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_v1alpha1_resource_preset_proto_rawDescOnce sync.Once
	file_nebius_msp_v1alpha1_resource_preset_proto_rawDescData = file_nebius_msp_v1alpha1_resource_preset_proto_rawDesc
)

func file_nebius_msp_v1alpha1_resource_preset_proto_rawDescGZIP() []byte {
	file_nebius_msp_v1alpha1_resource_preset_proto_rawDescOnce.Do(func() {
		file_nebius_msp_v1alpha1_resource_preset_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_v1alpha1_resource_preset_proto_rawDescData)
	})
	return file_nebius_msp_v1alpha1_resource_preset_proto_rawDescData
}

var file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_msp_v1alpha1_resource_preset_proto_goTypes = []any{
	(*Preset)(nil),     // 0: nebius.msp.v1alpha1.resource.Preset
	(*PresetSpec)(nil), // 1: nebius.msp.v1alpha1.resource.PresetSpec
	(*FlavorSpec)(nil), // 2: nebius.msp.v1alpha1.resource.FlavorSpec
	(*CpuSpec)(nil),    // 3: nebius.msp.v1alpha1.resource.CpuSpec
	(*MemorySpec)(nil), // 4: nebius.msp.v1alpha1.resource.MemorySpec
	(*Host)(nil),       // 5: nebius.msp.v1alpha1.resource.Host
	(*Disk)(nil),       // 6: nebius.msp.v1alpha1.resource.Disk
}
var file_nebius_msp_v1alpha1_resource_preset_proto_depIdxs = []int32{
	1, // 0: nebius.msp.v1alpha1.resource.Preset.spec:type_name -> nebius.msp.v1alpha1.resource.PresetSpec
	2, // 1: nebius.msp.v1alpha1.resource.PresetSpec.flavor:type_name -> nebius.msp.v1alpha1.resource.FlavorSpec
	5, // 2: nebius.msp.v1alpha1.resource.PresetSpec.hosts:type_name -> nebius.msp.v1alpha1.resource.Host
	6, // 3: nebius.msp.v1alpha1.resource.PresetSpec.disk:type_name -> nebius.msp.v1alpha1.resource.Disk
	3, // 4: nebius.msp.v1alpha1.resource.FlavorSpec.cpu:type_name -> nebius.msp.v1alpha1.resource.CpuSpec
	4, // 5: nebius.msp.v1alpha1.resource.FlavorSpec.memory:type_name -> nebius.msp.v1alpha1.resource.MemorySpec
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_msp_v1alpha1_resource_preset_proto_init() }
func file_nebius_msp_v1alpha1_resource_preset_proto_init() {
	if File_nebius_msp_v1alpha1_resource_preset_proto != nil {
		return
	}
	file_nebius_msp_v1alpha1_resource_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Preset); i {
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
		file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PresetSpec); i {
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
		file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FlavorSpec); i {
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
		file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CpuSpec); i {
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
		file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MemorySpec); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_v1alpha1_resource_preset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_v1alpha1_resource_preset_proto_goTypes,
		DependencyIndexes: file_nebius_msp_v1alpha1_resource_preset_proto_depIdxs,
		MessageInfos:      file_nebius_msp_v1alpha1_resource_preset_proto_msgTypes,
	}.Build()
	File_nebius_msp_v1alpha1_resource_preset_proto = out.File
	file_nebius_msp_v1alpha1_resource_preset_proto_rawDesc = nil
	file_nebius_msp_v1alpha1_resource_preset_proto_goTypes = nil
	file_nebius_msp_v1alpha1_resource_preset_proto_depIdxs = nil
}
