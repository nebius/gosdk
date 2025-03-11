// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/storage/v1/base.proto

package v1

import (
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

type StorageClass int32

const (
	StorageClass_STORAGE_CLASS_UNSPECIFIED StorageClass = 0
	StorageClass_STANDARD                  StorageClass = 1
	StorageClass_ENHANCED_THROUGHPUT       StorageClass = 2
)

// Enum value maps for StorageClass.
var (
	StorageClass_name = map[int32]string{
		0: "STORAGE_CLASS_UNSPECIFIED",
		1: "STANDARD",
		2: "ENHANCED_THROUGHPUT",
	}
	StorageClass_value = map[string]int32{
		"STORAGE_CLASS_UNSPECIFIED": 0,
		"STANDARD":                  1,
		"ENHANCED_THROUGHPUT":       2,
	}
)

func (x StorageClass) Enum() *StorageClass {
	p := new(StorageClass)
	*p = x
	return p
}

func (x StorageClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageClass) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_storage_v1_base_proto_enumTypes[0].Descriptor()
}

func (StorageClass) Type() protoreflect.EnumType {
	return &file_nebius_storage_v1_base_proto_enumTypes[0]
}

func (x StorageClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageClass.Descriptor instead.
func (StorageClass) EnumDescriptor() ([]byte, []int) {
	return file_nebius_storage_v1_base_proto_rawDescGZIP(), []int{0}
}

type VersioningPolicy int32

const (
	VersioningPolicy_VERSIONING_POLICY_UNSPECIFIED VersioningPolicy = 0
	VersioningPolicy_DISABLED                      VersioningPolicy = 1
	VersioningPolicy_ENABLED                       VersioningPolicy = 2
	VersioningPolicy_SUSPENDED                     VersioningPolicy = 3
)

// Enum value maps for VersioningPolicy.
var (
	VersioningPolicy_name = map[int32]string{
		0: "VERSIONING_POLICY_UNSPECIFIED",
		1: "DISABLED",
		2: "ENABLED",
		3: "SUSPENDED",
	}
	VersioningPolicy_value = map[string]int32{
		"VERSIONING_POLICY_UNSPECIFIED": 0,
		"DISABLED":                      1,
		"ENABLED":                       2,
		"SUSPENDED":                     3,
	}
)

func (x VersioningPolicy) Enum() *VersioningPolicy {
	p := new(VersioningPolicy)
	*p = x
	return p
}

func (x VersioningPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VersioningPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_storage_v1_base_proto_enumTypes[1].Descriptor()
}

func (VersioningPolicy) Type() protoreflect.EnumType {
	return &file_nebius_storage_v1_base_proto_enumTypes[1]
}

func (x VersioningPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VersioningPolicy.Descriptor instead.
func (VersioningPolicy) EnumDescriptor() ([]byte, []int) {
	return file_nebius_storage_v1_base_proto_rawDescGZIP(), []int{1}
}

var File_nebius_storage_v1_base_proto protoreflect.FileDescriptor

var file_nebius_storage_v1_base_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2a, 0x54, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4c, 0x41,
	0x53, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x54, 0x48, 0x52, 0x4f, 0x55,
	0x47, 0x48, 0x50, 0x55, 0x54, 0x10, 0x02, 0x2a, 0x5f, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x0a, 0x1d, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53,
	0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x42, 0x58, 0x0a, 0x18, 0x61, 0x69, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_storage_v1_base_proto_rawDescOnce sync.Once
	file_nebius_storage_v1_base_proto_rawDescData = file_nebius_storage_v1_base_proto_rawDesc
)

func file_nebius_storage_v1_base_proto_rawDescGZIP() []byte {
	file_nebius_storage_v1_base_proto_rawDescOnce.Do(func() {
		file_nebius_storage_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_storage_v1_base_proto_rawDescData)
	})
	return file_nebius_storage_v1_base_proto_rawDescData
}

var file_nebius_storage_v1_base_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nebius_storage_v1_base_proto_goTypes = []any{
	(StorageClass)(0),     // 0: nebius.storage.v1.StorageClass
	(VersioningPolicy)(0), // 1: nebius.storage.v1.VersioningPolicy
}
var file_nebius_storage_v1_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nebius_storage_v1_base_proto_init() }
func file_nebius_storage_v1_base_proto_init() {
	if File_nebius_storage_v1_base_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_storage_v1_base_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_storage_v1_base_proto_goTypes,
		DependencyIndexes: file_nebius_storage_v1_base_proto_depIdxs,
		EnumInfos:         file_nebius_storage_v1_base_proto_enumTypes,
	}.Build()
	File_nebius_storage_v1_base_proto = out.File
	file_nebius_storage_v1_base_proto_rawDesc = nil
	file_nebius_storage_v1_base_proto_goTypes = nil
	file_nebius_storage_v1_base_proto_depIdxs = nil
}
