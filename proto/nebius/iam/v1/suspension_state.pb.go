// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/iam/v1/suspension_state.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SuspensionState int32

const (
	SuspensionState_SUSPENSION_STATE_UNSPECIFIED SuspensionState = 0
	SuspensionState_NONE                         SuspensionState = 1
	SuspensionState_SUSPENDING                   SuspensionState = 2
	SuspensionState_SUSPENDED                    SuspensionState = 3
	SuspensionState_RESUMING                     SuspensionState = 4
)

// Enum value maps for SuspensionState.
var (
	SuspensionState_name = map[int32]string{
		0: "SUSPENSION_STATE_UNSPECIFIED",
		1: "NONE",
		2: "SUSPENDING",
		3: "SUSPENDED",
		4: "RESUMING",
	}
	SuspensionState_value = map[string]int32{
		"SUSPENSION_STATE_UNSPECIFIED": 0,
		"NONE":                         1,
		"SUSPENDING":                   2,
		"SUSPENDED":                    3,
		"RESUMING":                     4,
	}
)

func (x SuspensionState) Enum() *SuspensionState {
	p := new(SuspensionState)
	*p = x
	return p
}

func (x SuspensionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SuspensionState) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_suspension_state_proto_enumTypes[0].Descriptor()
}

func (SuspensionState) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_suspension_state_proto_enumTypes[0]
}

func (x SuspensionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SuspensionState.Descriptor instead.
func (SuspensionState) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_suspension_state_proto_rawDescGZIP(), []int{0}
}

var File_nebius_iam_v1_suspension_state_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_suspension_state_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2a, 0x6a, 0x0a, 0x0f, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x55, 0x53, 0x50,
	0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x49, 0x4e, 0x47, 0x10,
	0x04, 0x42, 0x5b, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x53, 0x75, 0x73, 0x70, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v1_suspension_state_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_suspension_state_proto_rawDescData []byte
)

func file_nebius_iam_v1_suspension_state_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_suspension_state_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_suspension_state_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_suspension_state_proto_rawDesc), len(file_nebius_iam_v1_suspension_state_proto_rawDesc)))
	})
	return file_nebius_iam_v1_suspension_state_proto_rawDescData
}

var file_nebius_iam_v1_suspension_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_suspension_state_proto_goTypes = []any{
	(SuspensionState)(0), // 0: nebius.iam.v1.SuspensionState
}
var file_nebius_iam_v1_suspension_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_suspension_state_proto_init() }
func file_nebius_iam_v1_suspension_state_proto_init() {
	if File_nebius_iam_v1_suspension_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_suspension_state_proto_rawDesc), len(file_nebius_iam_v1_suspension_state_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_suspension_state_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_suspension_state_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_suspension_state_proto_enumTypes,
	}.Build()
	File_nebius_iam_v1_suspension_state_proto = out.File
	file_nebius_iam_v1_suspension_state_proto_goTypes = nil
	file_nebius_iam_v1_suspension_state_proto_depIdxs = nil
}
