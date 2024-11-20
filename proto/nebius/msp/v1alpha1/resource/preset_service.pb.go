// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/msp/v1alpha1/resource/preset_service.proto

package resource

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

type ListPresetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int64  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPresetsRequest) Reset() {
	*x = ListPresetsRequest{}
	mi := &file_nebius_msp_v1alpha1_resource_preset_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPresetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPresetsRequest) ProtoMessage() {}

func (x *ListPresetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPresetsRequest.ProtoReflect.Descriptor instead.
func (*ListPresetsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListPresetsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPresetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPresetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items         []*Preset `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string    `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPresetsResponse) Reset() {
	*x = ListPresetsResponse{}
	mi := &file_nebius_msp_v1alpha1_resource_preset_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPresetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPresetsResponse) ProtoMessage() {}

func (x *ListPresetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_v1alpha1_resource_preset_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPresetsResponse.ProtoReflect.Descriptor instead.
func (*ListPresetsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPresetsResponse) GetItems() []*Preset {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPresetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_msp_v1alpha1_resource_preset_service_proto protoreflect.FileDescriptor

var file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x1a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x7c, 0x0a, 0x0d, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x77, 0x0a, 0x23, 0x61, 0x69, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x12,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescOnce sync.Once
	file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescData = file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDesc
)

func file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescGZIP() []byte {
	file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescOnce.Do(func() {
		file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescData)
	})
	return file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDescData
}

var file_nebius_msp_v1alpha1_resource_preset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nebius_msp_v1alpha1_resource_preset_service_proto_goTypes = []any{
	(*ListPresetsRequest)(nil),  // 0: nebius.msp.v1alpha1.resource.ListPresetsRequest
	(*ListPresetsResponse)(nil), // 1: nebius.msp.v1alpha1.resource.ListPresetsResponse
	(*Preset)(nil),              // 2: nebius.msp.v1alpha1.resource.Preset
}
var file_nebius_msp_v1alpha1_resource_preset_service_proto_depIdxs = []int32{
	2, // 0: nebius.msp.v1alpha1.resource.ListPresetsResponse.items:type_name -> nebius.msp.v1alpha1.resource.Preset
	0, // 1: nebius.msp.v1alpha1.resource.PresetService.List:input_type -> nebius.msp.v1alpha1.resource.ListPresetsRequest
	1, // 2: nebius.msp.v1alpha1.resource.PresetService.List:output_type -> nebius.msp.v1alpha1.resource.ListPresetsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_msp_v1alpha1_resource_preset_service_proto_init() }
func file_nebius_msp_v1alpha1_resource_preset_service_proto_init() {
	if File_nebius_msp_v1alpha1_resource_preset_service_proto != nil {
		return
	}
	file_nebius_msp_v1alpha1_resource_preset_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_msp_v1alpha1_resource_preset_service_proto_goTypes,
		DependencyIndexes: file_nebius_msp_v1alpha1_resource_preset_service_proto_depIdxs,
		MessageInfos:      file_nebius_msp_v1alpha1_resource_preset_service_proto_msgTypes,
	}.Build()
	File_nebius_msp_v1alpha1_resource_preset_service_proto = out.File
	file_nebius_msp_v1alpha1_resource_preset_service_proto_rawDesc = nil
	file_nebius_msp_v1alpha1_resource_preset_service_proto_goTypes = nil
	file_nebius_msp_v1alpha1_resource_preset_service_proto_depIdxs = nil
}
