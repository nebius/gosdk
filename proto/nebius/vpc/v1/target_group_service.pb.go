// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/vpc/v1/target_group_service.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
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

type GetTargetGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTargetGroupRequest) Reset() {
	*x = GetTargetGroupRequest{}
	mi := &file_nebius_vpc_v1_target_group_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTargetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetGroupRequest) ProtoMessage() {}

func (x *GetTargetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_target_group_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetTargetGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_target_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTargetGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateTargetGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *TargetGroupSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTargetGroupRequest) Reset() {
	*x = UpdateTargetGroupRequest{}
	mi := &file_nebius_vpc_v1_target_group_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTargetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTargetGroupRequest) ProtoMessage() {}

func (x *UpdateTargetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_target_group_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTargetGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateTargetGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_target_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTargetGroupRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateTargetGroupRequest) GetSpec() *TargetGroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

var File_nebius_vpc_v1_target_group_service_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1_target_group_service_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd8, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x87, 0x02, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xc6, 0x01, 0xba, 0x48, 0xc2,
	0x01, 0xba, 0x01, 0xbe, 0x01, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x73, 0x27, 0x6e, 0x61, 0x6d, 0x65, 0x27, 0x20, 0x6d, 0x75, 0x73,
	0x74, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61, 0x20, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x20, 0x6f, 0x72, 0x20, 0x64, 0x69, 0x67, 0x69, 0x74, 0x2c, 0x20,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x20, 0x27, 0x2d, 0x27, 0x2c, 0x20, 0x27, 0x5f, 0x27, 0x2c, 0x20,
	0x27, 0x2e, 0x27, 0x2c, 0x20, 0x27, 0x2f, 0x27, 0x2c, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x68, 0x61,
	0x76, 0x65, 0x20, 0x61, 0x20, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x62, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x20, 0x32, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x32, 0x35, 0x35, 0x20, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x1a, 0x38, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x2d, 0x5f, 0x2e, 0x2f, 0x61,
	0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x35, 0x34, 0x7d,
	0x24, 0x27, 0x29, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x32, 0xb5, 0x01, 0x0a, 0x12, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x4e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x06, 0xba, 0x4a, 0x03, 0x76, 0x70, 0x63, 0x42, 0x5e, 0x0a, 0x14, 0x61, 0x69, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x42, 0x17, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_nebius_vpc_v1_target_group_service_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1_target_group_service_proto_rawDescData []byte
)

func file_nebius_vpc_v1_target_group_service_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1_target_group_service_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1_target_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1_target_group_service_proto_rawDesc), len(file_nebius_vpc_v1_target_group_service_proto_rawDesc)))
	})
	return file_nebius_vpc_v1_target_group_service_proto_rawDescData
}

var file_nebius_vpc_v1_target_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nebius_vpc_v1_target_group_service_proto_goTypes = []any{
	(*GetTargetGroupRequest)(nil),    // 0: nebius.vpc.v1.GetTargetGroupRequest
	(*UpdateTargetGroupRequest)(nil), // 1: nebius.vpc.v1.UpdateTargetGroupRequest
	(*v1.ResourceMetadata)(nil),      // 2: nebius.common.v1.ResourceMetadata
	(*TargetGroupSpec)(nil),          // 3: nebius.vpc.v1.TargetGroupSpec
	(*TargetGroup)(nil),              // 4: nebius.vpc.v1.TargetGroup
	(*v1.Operation)(nil),             // 5: nebius.common.v1.Operation
}
var file_nebius_vpc_v1_target_group_service_proto_depIdxs = []int32{
	2, // 0: nebius.vpc.v1.UpdateTargetGroupRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	3, // 1: nebius.vpc.v1.UpdateTargetGroupRequest.spec:type_name -> nebius.vpc.v1.TargetGroupSpec
	0, // 2: nebius.vpc.v1.TargetGroupService.Get:input_type -> nebius.vpc.v1.GetTargetGroupRequest
	1, // 3: nebius.vpc.v1.TargetGroupService.Update:input_type -> nebius.vpc.v1.UpdateTargetGroupRequest
	4, // 4: nebius.vpc.v1.TargetGroupService.Get:output_type -> nebius.vpc.v1.TargetGroup
	5, // 5: nebius.vpc.v1.TargetGroupService.Update:output_type -> nebius.common.v1.Operation
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1_target_group_service_proto_init() }
func file_nebius_vpc_v1_target_group_service_proto_init() {
	if File_nebius_vpc_v1_target_group_service_proto != nil {
		return
	}
	file_nebius_vpc_v1_target_group_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1_target_group_service_proto_rawDesc), len(file_nebius_vpc_v1_target_group_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_vpc_v1_target_group_service_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1_target_group_service_proto_depIdxs,
		MessageInfos:      file_nebius_vpc_v1_target_group_service_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1_target_group_service_proto = out.File
	file_nebius_vpc_v1_target_group_service_proto_goTypes = nil
	file_nebius_vpc_v1_target_group_service_proto_depIdxs = nil
}
