// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/msp/postgresql/v1alpha1/preset.proto

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

type TemplateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hosts     *resource.HostSpec      `protobuf:"bytes,2,opt,name=hosts,proto3" json:"hosts,omitempty"`
	Disk      *resource.DiskSpec      `protobuf:"bytes,3,opt,name=disk,proto3" json:"disk,omitempty"`
	Resources *resource.ResourcesSpec `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *TemplateSpec) Reset() {
	*x = TemplateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_msp_postgresql_v1alpha1_preset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateSpec) ProtoMessage() {}

func (x *TemplateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_preset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateSpec.ProtoReflect.Descriptor instead.
func (*TemplateSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateSpec) GetHosts() *resource.HostSpec {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *TemplateSpec) GetDisk() *resource.DiskSpec {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *TemplateSpec) GetResources() *resource.ResourcesSpec {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_nebius_msp_postgresql_v1alpha1_preset_proto protoreflect.FileDescriptor

var file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x44, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x42,
	0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x6b,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x64, 0x69,
	0x73, 0x6b, 0x12, 0x49, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x74, 0x0a,
	0x25, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescOnce sync.Once
	file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescData = file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDesc
)

func file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescGZIP() []byte {
	file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescOnce.Do(func() {
		file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescData)
	})
	return file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDescData
}

var file_nebius_msp_postgresql_v1alpha1_preset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nebius_msp_postgresql_v1alpha1_preset_proto_goTypes = []any{
	(*TemplateSpec)(nil),           // 0: nebius.msp.postgresql.v1alpha1.TemplateSpec
	(*resource.HostSpec)(nil),      // 1: nebius.msp.v1alpha1.resource.HostSpec
	(*resource.DiskSpec)(nil),      // 2: nebius.msp.v1alpha1.resource.DiskSpec
	(*resource.ResourcesSpec)(nil), // 3: nebius.msp.v1alpha1.resource.ResourcesSpec
}
var file_nebius_msp_postgresql_v1alpha1_preset_proto_depIdxs = []int32{
	1, // 0: nebius.msp.postgresql.v1alpha1.TemplateSpec.hosts:type_name -> nebius.msp.v1alpha1.resource.HostSpec
	2, // 1: nebius.msp.postgresql.v1alpha1.TemplateSpec.disk:type_name -> nebius.msp.v1alpha1.resource.DiskSpec
	3, // 2: nebius.msp.postgresql.v1alpha1.TemplateSpec.resources:type_name -> nebius.msp.v1alpha1.resource.ResourcesSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nebius_msp_postgresql_v1alpha1_preset_proto_init() }
func file_nebius_msp_postgresql_v1alpha1_preset_proto_init() {
	if File_nebius_msp_postgresql_v1alpha1_preset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nebius_msp_postgresql_v1alpha1_preset_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TemplateSpec); i {
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
			RawDescriptor: file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_postgresql_v1alpha1_preset_proto_goTypes,
		DependencyIndexes: file_nebius_msp_postgresql_v1alpha1_preset_proto_depIdxs,
		MessageInfos:      file_nebius_msp_postgresql_v1alpha1_preset_proto_msgTypes,
	}.Build()
	File_nebius_msp_postgresql_v1alpha1_preset_proto = out.File
	file_nebius_msp_postgresql_v1alpha1_preset_proto_rawDesc = nil
	file_nebius_msp_postgresql_v1alpha1_preset_proto_goTypes = nil
	file_nebius_msp_postgresql_v1alpha1_preset_proto_depIdxs = nil
}
