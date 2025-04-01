// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/mk8s/v1alpha1/progress_data.proto

package v1alpha1

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

type ProgressData struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Problems encountered during the operation execution.
	Problems      []*Problem `protobuf:"bytes,1,rep,name=problems,proto3" json:"problems,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProgressData) Reset() {
	*x = ProgressData{}
	mi := &file_nebius_mk8s_v1alpha1_progress_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProgressData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressData) ProtoMessage() {}

func (x *ProgressData) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_progress_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressData.ProtoReflect.Descriptor instead.
func (*ProgressData) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescGZIP(), []int{0}
}

func (x *ProgressData) GetProblems() []*Problem {
	if x != nil {
		return x.Problems
	}
	return nil
}

type Problem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Stage on which the problem occurred.
	Stage string `protobuf:"bytes,1,opt,name=stage,proto3" json:"stage,omitempty"`
	// Message describing the problem.
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Problem) Reset() {
	*x = Problem{}
	mi := &file_nebius_mk8s_v1alpha1_progress_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_progress_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescGZIP(), []int{1}
}

func (x *Problem) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *Problem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_nebius_mk8s_v1alpha1_progress_data_proto protoreflect.FileDescriptor

var file_nebius_mk8s_v1alpha1_progress_data_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x22, 0x49, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x22, 0x39, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x66, 0x0a, 0x1b, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x11, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescOnce sync.Once
	file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescData []byte
)

func file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescGZIP() []byte {
	file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescOnce.Do(func() {
		file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_mk8s_v1alpha1_progress_data_proto_rawDesc), len(file_nebius_mk8s_v1alpha1_progress_data_proto_rawDesc)))
	})
	return file_nebius_mk8s_v1alpha1_progress_data_proto_rawDescData
}

var file_nebius_mk8s_v1alpha1_progress_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nebius_mk8s_v1alpha1_progress_data_proto_goTypes = []any{
	(*ProgressData)(nil), // 0: nebius.mk8s.v1alpha1.ProgressData
	(*Problem)(nil),      // 1: nebius.mk8s.v1alpha1.Problem
}
var file_nebius_mk8s_v1alpha1_progress_data_proto_depIdxs = []int32{
	1, // 0: nebius.mk8s.v1alpha1.ProgressData.problems:type_name -> nebius.mk8s.v1alpha1.Problem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_mk8s_v1alpha1_progress_data_proto_init() }
func file_nebius_mk8s_v1alpha1_progress_data_proto_init() {
	if File_nebius_mk8s_v1alpha1_progress_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_mk8s_v1alpha1_progress_data_proto_rawDesc), len(file_nebius_mk8s_v1alpha1_progress_data_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_mk8s_v1alpha1_progress_data_proto_goTypes,
		DependencyIndexes: file_nebius_mk8s_v1alpha1_progress_data_proto_depIdxs,
		MessageInfos:      file_nebius_mk8s_v1alpha1_progress_data_proto_msgTypes,
	}.Build()
	File_nebius_mk8s_v1alpha1_progress_data_proto = out.File
	file_nebius_mk8s_v1alpha1_progress_data_proto_goTypes = nil
	file_nebius_mk8s_v1alpha1_progress_data_proto_depIdxs = nil
}
