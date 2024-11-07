// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/mk8s/v1/instance_template.proto

package v1

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

type DiskSpec_DiskType int32

const (
	DiskSpec_UNSPECIFIED DiskSpec_DiskType = 0
	// the list of available types will be clarified later, it is not final version
	DiskSpec_NETWORK_SSD                DiskSpec_DiskType = 1
	DiskSpec_NETWORK_HDD                DiskSpec_DiskType = 2
	DiskSpec_NETWORK_SSD_IO_M3          DiskSpec_DiskType = 3
	DiskSpec_NETWORK_SSD_NON_REPLICATED DiskSpec_DiskType = 4
)

// Enum value maps for DiskSpec_DiskType.
var (
	DiskSpec_DiskType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "NETWORK_SSD",
		2: "NETWORK_HDD",
		3: "NETWORK_SSD_IO_M3",
		4: "NETWORK_SSD_NON_REPLICATED",
	}
	DiskSpec_DiskType_value = map[string]int32{
		"UNSPECIFIED":                0,
		"NETWORK_SSD":                1,
		"NETWORK_HDD":                2,
		"NETWORK_SSD_IO_M3":          3,
		"NETWORK_SSD_NON_REPLICATED": 4,
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
	return file_nebius_mk8s_v1_instance_template_proto_enumTypes[0].Descriptor()
}

func (DiskSpec_DiskType) Type() protoreflect.EnumType {
	return &file_nebius_mk8s_v1_instance_template_proto_enumTypes[0]
}

func (x DiskSpec_DiskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiskSpec_DiskType.Descriptor instead.
func (DiskSpec_DiskType) EnumDescriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_instance_template_proto_rawDescGZIP(), []int{0, 0}
}

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
	Size           isDiskSpec_Size   `protobuf_oneof:"size"`
	BlockSizeBytes int64             `protobuf:"varint,5,opt,name=block_size_bytes,json=blockSizeBytes,proto3" json:"block_size_bytes,omitempty"`
	Type           DiskSpec_DiskType `protobuf:"varint,6,opt,name=type,proto3,enum=nebius.mk8s.v1.DiskSpec_DiskType" json:"type,omitempty"`
}

func (x *DiskSpec) Reset() {
	*x = DiskSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_mk8s_v1_instance_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskSpec) ProtoMessage() {}

func (x *DiskSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_instance_template_proto_msgTypes[0]
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
	return file_nebius_mk8s_v1_instance_template_proto_rawDescGZIP(), []int{0}
}

func (m *DiskSpec) GetSize() isDiskSpec_Size {
	if m != nil {
		return m.Size
	}
	return nil
}

func (x *DiskSpec) GetSizeBytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeBytes); ok {
		return x.SizeBytes
	}
	return 0
}

func (x *DiskSpec) GetSizeKibibytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeKibibytes); ok {
		return x.SizeKibibytes
	}
	return 0
}

func (x *DiskSpec) GetSizeMebibytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeMebibytes); ok {
		return x.SizeMebibytes
	}
	return 0
}

func (x *DiskSpec) GetSizeGibibytes() int64 {
	if x, ok := x.GetSize().(*DiskSpec_SizeGibibytes); ok {
		return x.SizeGibibytes
	}
	return 0
}

func (x *DiskSpec) GetBlockSizeBytes() int64 {
	if x != nil {
		return x.BlockSizeBytes
	}
	return 0
}

func (x *DiskSpec) GetType() DiskSpec_DiskType {
	if x != nil {
		return x.Type
	}
	return DiskSpec_UNSPECIFIED
}

type isDiskSpec_Size interface {
	isDiskSpec_Size()
}

type DiskSpec_SizeBytes struct {
	SizeBytes int64 `protobuf:"varint,1,opt,name=size_bytes,json=sizeBytes,proto3,oneof"`
}

type DiskSpec_SizeKibibytes struct {
	SizeKibibytes int64 `protobuf:"varint,2,opt,name=size_kibibytes,json=sizeKibibytes,proto3,oneof"`
}

type DiskSpec_SizeMebibytes struct {
	SizeMebibytes int64 `protobuf:"varint,3,opt,name=size_mebibytes,json=sizeMebibytes,proto3,oneof"`
}

type DiskSpec_SizeGibibytes struct {
	SizeGibibytes int64 `protobuf:"varint,4,opt,name=size_gibibytes,json=sizeGibibytes,proto3,oneof"`
}

func (*DiskSpec_SizeBytes) isDiskSpec_Size() {}

func (*DiskSpec_SizeKibibytes) isDiskSpec_Size() {}

func (*DiskSpec_SizeMebibytes) isDiskSpec_Size() {}

func (*DiskSpec_SizeGibibytes) isDiskSpec_Size() {}

type ResourcesSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	// Types that are assignable to Size:
	//
	//	*ResourcesSpec_Preset
	Size isResourcesSpec_Size `protobuf_oneof:"size"`
}

func (x *ResourcesSpec) Reset() {
	*x = ResourcesSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_mk8s_v1_instance_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcesSpec) ProtoMessage() {}

func (x *ResourcesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_instance_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcesSpec.ProtoReflect.Descriptor instead.
func (*ResourcesSpec) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_instance_template_proto_rawDescGZIP(), []int{1}
}

func (x *ResourcesSpec) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (m *ResourcesSpec) GetSize() isResourcesSpec_Size {
	if m != nil {
		return m.Size
	}
	return nil
}

func (x *ResourcesSpec) GetPreset() string {
	if x, ok := x.GetSize().(*ResourcesSpec_Preset); ok {
		return x.Preset
	}
	return ""
}

type isResourcesSpec_Size interface {
	isResourcesSpec_Size()
}

type ResourcesSpec_Preset struct {
	Preset string `protobuf:"bytes,2,opt,name=preset,proto3,oneof"`
}

func (*ResourcesSpec_Preset) isResourcesSpec_Size() {}

var File_nebius_mk8s_v1_instance_template_proto protoreflect.FileDescriptor

var file_nebius_mk8s_v1_instance_template_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x03, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0a,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a,
	0x0e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6b, 0x69, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x4b, 0x69, 0x62,
	0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d,
	0x65, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x0e, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x67, 0x69, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x7a, 0x65, 0x47,
	0x69, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x07, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x74, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53, 0x53, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x48, 0x44, 0x44, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x53, 0x53, 0x44, 0x5f, 0x49, 0x4f,
	0x5f, 0x4d, 0x33, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x53, 0x53, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x04, 0x42, 0x0d, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x05, 0xba,
	0x48, 0x02, 0x08, 0x01, 0x22, 0x5c, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x42, 0x0d, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02,
	0x08, 0x01, 0x42, 0x5e, 0x0a, 0x15, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_mk8s_v1_instance_template_proto_rawDescOnce sync.Once
	file_nebius_mk8s_v1_instance_template_proto_rawDescData = file_nebius_mk8s_v1_instance_template_proto_rawDesc
)

func file_nebius_mk8s_v1_instance_template_proto_rawDescGZIP() []byte {
	file_nebius_mk8s_v1_instance_template_proto_rawDescOnce.Do(func() {
		file_nebius_mk8s_v1_instance_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_mk8s_v1_instance_template_proto_rawDescData)
	})
	return file_nebius_mk8s_v1_instance_template_proto_rawDescData
}

var file_nebius_mk8s_v1_instance_template_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_mk8s_v1_instance_template_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nebius_mk8s_v1_instance_template_proto_goTypes = []any{
	(DiskSpec_DiskType)(0), // 0: nebius.mk8s.v1.DiskSpec.DiskType
	(*DiskSpec)(nil),       // 1: nebius.mk8s.v1.DiskSpec
	(*ResourcesSpec)(nil),  // 2: nebius.mk8s.v1.ResourcesSpec
}
var file_nebius_mk8s_v1_instance_template_proto_depIdxs = []int32{
	0, // 0: nebius.mk8s.v1.DiskSpec.type:type_name -> nebius.mk8s.v1.DiskSpec.DiskType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_mk8s_v1_instance_template_proto_init() }
func file_nebius_mk8s_v1_instance_template_proto_init() {
	if File_nebius_mk8s_v1_instance_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nebius_mk8s_v1_instance_template_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_nebius_mk8s_v1_instance_template_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResourcesSpec); i {
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
	file_nebius_mk8s_v1_instance_template_proto_msgTypes[0].OneofWrappers = []any{
		(*DiskSpec_SizeBytes)(nil),
		(*DiskSpec_SizeKibibytes)(nil),
		(*DiskSpec_SizeMebibytes)(nil),
		(*DiskSpec_SizeGibibytes)(nil),
	}
	file_nebius_mk8s_v1_instance_template_proto_msgTypes[1].OneofWrappers = []any{
		(*ResourcesSpec_Preset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_mk8s_v1_instance_template_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_mk8s_v1_instance_template_proto_goTypes,
		DependencyIndexes: file_nebius_mk8s_v1_instance_template_proto_depIdxs,
		EnumInfos:         file_nebius_mk8s_v1_instance_template_proto_enumTypes,
		MessageInfos:      file_nebius_mk8s_v1_instance_template_proto_msgTypes,
	}.Build()
	File_nebius_mk8s_v1_instance_template_proto = out.File
	file_nebius_mk8s_v1_instance_template_proto_rawDesc = nil
	file_nebius_mk8s_v1_instance_template_proto_goTypes = nil
	file_nebius_mk8s_v1_instance_template_proto_depIdxs = nil
}
