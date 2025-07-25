// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/audit/v2/access_token.proto

package v2

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type AccessToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Token without signature.
	MaskedToken   string `protobuf:"bytes,2,opt,name=masked_token,json=maskedToken,proto3" json:"masked_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	mi := &file_nebius_audit_v2_access_token_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_audit_v2_access_token_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_nebius_audit_v2_access_token_proto_rawDescGZIP(), []int{0}
}

func (x *AccessToken) GetMaskedToken() string {
	if x != nil {
		return x.MaskedToken
	}
	return ""
}

var File_nebius_audit_v2_access_token_proto protoreflect.FileDescriptor

var file_nebius_audit_v2_access_token_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x29, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x0b, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x5b, 0x0a, 0x16,
	0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x42, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_nebius_audit_v2_access_token_proto_rawDescOnce sync.Once
	file_nebius_audit_v2_access_token_proto_rawDescData []byte
)

func file_nebius_audit_v2_access_token_proto_rawDescGZIP() []byte {
	file_nebius_audit_v2_access_token_proto_rawDescOnce.Do(func() {
		file_nebius_audit_v2_access_token_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_audit_v2_access_token_proto_rawDesc), len(file_nebius_audit_v2_access_token_proto_rawDesc)))
	})
	return file_nebius_audit_v2_access_token_proto_rawDescData
}

var file_nebius_audit_v2_access_token_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nebius_audit_v2_access_token_proto_goTypes = []any{
	(*AccessToken)(nil), // 0: nebius.audit.v2.AccessToken
}
var file_nebius_audit_v2_access_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nebius_audit_v2_access_token_proto_init() }
func file_nebius_audit_v2_access_token_proto_init() {
	if File_nebius_audit_v2_access_token_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_audit_v2_access_token_proto_rawDesc), len(file_nebius_audit_v2_access_token_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_audit_v2_access_token_proto_goTypes,
		DependencyIndexes: file_nebius_audit_v2_access_token_proto_depIdxs,
		MessageInfos:      file_nebius_audit_v2_access_token_proto_msgTypes,
	}.Build()
	File_nebius_audit_v2_access_token_proto = out.File
	file_nebius_audit_v2_access_token_proto_goTypes = nil
	file_nebius_audit_v2_access_token_proto_depIdxs = nil
}
