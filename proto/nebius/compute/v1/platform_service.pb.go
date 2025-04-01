// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/compute/v1/platform_service.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type ListPlatformsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageSize      int64                  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	ParentId      string                 `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPlatformsRequest) Reset() {
	*x = ListPlatformsRequest{}
	mi := &file_nebius_compute_v1_platform_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPlatformsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlatformsRequest) ProtoMessage() {}

func (x *ListPlatformsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_platform_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlatformsRequest.ProtoReflect.Descriptor instead.
func (*ListPlatformsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_platform_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListPlatformsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPlatformsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListPlatformsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type ListPlatformsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Platform            `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPlatformsResponse) Reset() {
	*x = ListPlatformsResponse{}
	mi := &file_nebius_compute_v1_platform_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPlatformsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlatformsResponse) ProtoMessage() {}

func (x *ListPlatformsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_platform_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlatformsResponse.ProtoReflect.Descriptor instead.
func (*ListPlatformsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_platform_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPlatformsResponse) GetItems() []*Platform {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPlatformsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_compute_v1_platform_service_proto protoreflect.FileDescriptor

var file_nebius_compute_v1_platform_service_proto_rawDesc = string([]byte{
	0x0a, 0x28, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xba, 0x01, 0x0a, 0x0f, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x59, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x63, 0x0a, 0x18, 0x61, 0x69, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x14, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67,
	0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_compute_v1_platform_service_proto_rawDescOnce sync.Once
	file_nebius_compute_v1_platform_service_proto_rawDescData []byte
)

func file_nebius_compute_v1_platform_service_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1_platform_service_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1_platform_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_compute_v1_platform_service_proto_rawDesc), len(file_nebius_compute_v1_platform_service_proto_rawDesc)))
	})
	return file_nebius_compute_v1_platform_service_proto_rawDescData
}

var file_nebius_compute_v1_platform_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nebius_compute_v1_platform_service_proto_goTypes = []any{
	(*ListPlatformsRequest)(nil),  // 0: nebius.compute.v1.ListPlatformsRequest
	(*ListPlatformsResponse)(nil), // 1: nebius.compute.v1.ListPlatformsResponse
	(*Platform)(nil),              // 2: nebius.compute.v1.Platform
	(*v1.GetByNameRequest)(nil),   // 3: nebius.common.v1.GetByNameRequest
}
var file_nebius_compute_v1_platform_service_proto_depIdxs = []int32{
	2, // 0: nebius.compute.v1.ListPlatformsResponse.items:type_name -> nebius.compute.v1.Platform
	3, // 1: nebius.compute.v1.PlatformService.GetByName:input_type -> nebius.common.v1.GetByNameRequest
	0, // 2: nebius.compute.v1.PlatformService.List:input_type -> nebius.compute.v1.ListPlatformsRequest
	2, // 3: nebius.compute.v1.PlatformService.GetByName:output_type -> nebius.compute.v1.Platform
	1, // 4: nebius.compute.v1.PlatformService.List:output_type -> nebius.compute.v1.ListPlatformsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1_platform_service_proto_init() }
func file_nebius_compute_v1_platform_service_proto_init() {
	if File_nebius_compute_v1_platform_service_proto != nil {
		return
	}
	file_nebius_compute_v1_platform_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_compute_v1_platform_service_proto_rawDesc), len(file_nebius_compute_v1_platform_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_compute_v1_platform_service_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1_platform_service_proto_depIdxs,
		MessageInfos:      file_nebius_compute_v1_platform_service_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1_platform_service_proto = out.File
	file_nebius_compute_v1_platform_service_proto_goTypes = nil
	file_nebius_compute_v1_platform_service_proto_depIdxs = nil
}
