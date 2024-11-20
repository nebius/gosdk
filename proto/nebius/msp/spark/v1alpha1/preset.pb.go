// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/msp/spark/v1alpha1/preset.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	resource "github.com/nebius/gosdk/proto/nebius/msp/v1alpha1/resource"
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

type DriverTemplateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disk      *resource.DiskSpec      `protobuf:"bytes,3,opt,name=disk,proto3" json:"disk,omitempty"`
	Resources *resource.ResourcesSpec `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *DriverTemplateSpec) Reset() {
	*x = DriverTemplateSpec{}
	mi := &file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DriverTemplateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverTemplateSpec) ProtoMessage() {}

func (x *DriverTemplateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverTemplateSpec.ProtoReflect.Descriptor instead.
func (*DriverTemplateSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_preset_proto_rawDescGZIP(), []int{0}
}

func (x *DriverTemplateSpec) GetDisk() *resource.DiskSpec {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *DriverTemplateSpec) GetResources() *resource.ResourcesSpec {
	if x != nil {
		return x.Resources
	}
	return nil
}

type DynamicAllocationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min int64 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max int64 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *DynamicAllocationSpec) Reset() {
	*x = DynamicAllocationSpec{}
	mi := &file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DynamicAllocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicAllocationSpec) ProtoMessage() {}

func (x *DynamicAllocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicAllocationSpec.ProtoReflect.Descriptor instead.
func (*DynamicAllocationSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_preset_proto_rawDescGZIP(), []int{1}
}

func (x *DynamicAllocationSpec) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *DynamicAllocationSpec) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type ExecutorTemplateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disk      *resource.DiskSpec      `protobuf:"bytes,2,opt,name=disk,proto3" json:"disk,omitempty"`
	Resources *resource.ResourcesSpec `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	// Types that are assignable to HostsSpec:
	//
	//	*ExecutorTemplateSpec_Hosts
	//	*ExecutorTemplateSpec_HostsDynamicAllocation
	HostsSpec isExecutorTemplateSpec_HostsSpec `protobuf_oneof:"hosts_spec"`
}

func (x *ExecutorTemplateSpec) Reset() {
	*x = ExecutorTemplateSpec{}
	mi := &file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecutorTemplateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutorTemplateSpec) ProtoMessage() {}

func (x *ExecutorTemplateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutorTemplateSpec.ProtoReflect.Descriptor instead.
func (*ExecutorTemplateSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_preset_proto_rawDescGZIP(), []int{2}
}

func (x *ExecutorTemplateSpec) GetDisk() *resource.DiskSpec {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *ExecutorTemplateSpec) GetResources() *resource.ResourcesSpec {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (m *ExecutorTemplateSpec) GetHostsSpec() isExecutorTemplateSpec_HostsSpec {
	if m != nil {
		return m.HostsSpec
	}
	return nil
}

func (x *ExecutorTemplateSpec) GetHosts() *resource.HostSpec {
	if x, ok := x.GetHostsSpec().(*ExecutorTemplateSpec_Hosts); ok {
		return x.Hosts
	}
	return nil
}

func (x *ExecutorTemplateSpec) GetHostsDynamicAllocation() *DynamicAllocationSpec {
	if x, ok := x.GetHostsSpec().(*ExecutorTemplateSpec_HostsDynamicAllocation); ok {
		return x.HostsDynamicAllocation
	}
	return nil
}

type isExecutorTemplateSpec_HostsSpec interface {
	isExecutorTemplateSpec_HostsSpec()
}

type ExecutorTemplateSpec_Hosts struct {
	Hosts *resource.HostSpec `protobuf:"bytes,101,opt,name=hosts,proto3,oneof"`
}

type ExecutorTemplateSpec_HostsDynamicAllocation struct {
	HostsDynamicAllocation *DynamicAllocationSpec `protobuf:"bytes,102,opt,name=hosts_dynamic_allocation,json=hostsDynamicAllocation,proto3,oneof"`
}

func (*ExecutorTemplateSpec_Hosts) isExecutorTemplateSpec_HostsSpec() {}

func (*ExecutorTemplateSpec_HostsDynamicAllocation) isExecutorTemplateSpec_HostsSpec() {}

var File_nebius_msp_spark_v1alpha1_preset_proto protoreflect.FileDescriptor

var file_nebius_msp_spark_v1alpha1_preset_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01,
	0x0a, 0x12, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x49, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x03,
	0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78,
	0x22, 0xe8, 0x02, 0x0a, 0x14, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x04, 0x64, 0x69, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x49, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x48,
	0x00, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x6c, 0x0a, 0x18, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x16,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x13, 0x0a, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x42, 0x6a, 0x0a, 0x20, 0x61,
	0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x0b, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_spark_v1alpha1_preset_proto_rawDescOnce sync.Once
	file_nebius_msp_spark_v1alpha1_preset_proto_rawDescData = file_nebius_msp_spark_v1alpha1_preset_proto_rawDesc
)

func file_nebius_msp_spark_v1alpha1_preset_proto_rawDescGZIP() []byte {
	file_nebius_msp_spark_v1alpha1_preset_proto_rawDescOnce.Do(func() {
		file_nebius_msp_spark_v1alpha1_preset_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_spark_v1alpha1_preset_proto_rawDescData)
	})
	return file_nebius_msp_spark_v1alpha1_preset_proto_rawDescData
}

var file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_msp_spark_v1alpha1_preset_proto_goTypes = []any{
	(*DriverTemplateSpec)(nil),     // 0: nebius.msp.spark.v1alpha1.DriverTemplateSpec
	(*DynamicAllocationSpec)(nil),  // 1: nebius.msp.spark.v1alpha1.DynamicAllocationSpec
	(*ExecutorTemplateSpec)(nil),   // 2: nebius.msp.spark.v1alpha1.ExecutorTemplateSpec
	(*resource.DiskSpec)(nil),      // 3: nebius.msp.v1alpha1.resource.DiskSpec
	(*resource.ResourcesSpec)(nil), // 4: nebius.msp.v1alpha1.resource.ResourcesSpec
	(*resource.HostSpec)(nil),      // 5: nebius.msp.v1alpha1.resource.HostSpec
}
var file_nebius_msp_spark_v1alpha1_preset_proto_depIdxs = []int32{
	3, // 0: nebius.msp.spark.v1alpha1.DriverTemplateSpec.disk:type_name -> nebius.msp.v1alpha1.resource.DiskSpec
	4, // 1: nebius.msp.spark.v1alpha1.DriverTemplateSpec.resources:type_name -> nebius.msp.v1alpha1.resource.ResourcesSpec
	3, // 2: nebius.msp.spark.v1alpha1.ExecutorTemplateSpec.disk:type_name -> nebius.msp.v1alpha1.resource.DiskSpec
	4, // 3: nebius.msp.spark.v1alpha1.ExecutorTemplateSpec.resources:type_name -> nebius.msp.v1alpha1.resource.ResourcesSpec
	5, // 4: nebius.msp.spark.v1alpha1.ExecutorTemplateSpec.hosts:type_name -> nebius.msp.v1alpha1.resource.HostSpec
	1, // 5: nebius.msp.spark.v1alpha1.ExecutorTemplateSpec.hosts_dynamic_allocation:type_name -> nebius.msp.spark.v1alpha1.DynamicAllocationSpec
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_msp_spark_v1alpha1_preset_proto_init() }
func file_nebius_msp_spark_v1alpha1_preset_proto_init() {
	if File_nebius_msp_spark_v1alpha1_preset_proto != nil {
		return
	}
	file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes[2].OneofWrappers = []any{
		(*ExecutorTemplateSpec_Hosts)(nil),
		(*ExecutorTemplateSpec_HostsDynamicAllocation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_spark_v1alpha1_preset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_spark_v1alpha1_preset_proto_goTypes,
		DependencyIndexes: file_nebius_msp_spark_v1alpha1_preset_proto_depIdxs,
		MessageInfos:      file_nebius_msp_spark_v1alpha1_preset_proto_msgTypes,
	}.Build()
	File_nebius_msp_spark_v1alpha1_preset_proto = out.File
	file_nebius_msp_spark_v1alpha1_preset_proto_rawDesc = nil
	file_nebius_msp_spark_v1alpha1_preset_proto_goTypes = nil
	file_nebius_msp_spark_v1alpha1_preset_proto_depIdxs = nil
}
