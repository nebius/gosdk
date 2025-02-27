// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/registry/v1/artifact_service.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
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

type GetArtifactRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArtifactRequest) Reset() {
	*x = GetArtifactRequest{}
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArtifactRequest) ProtoMessage() {}

func (x *GetArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArtifactRequest.ProtoReflect.Descriptor instead.
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return file_nebius_registry_v1_artifact_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetArtifactRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListArtifactsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter        string                 `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListArtifactsRequest) Reset() {
	*x = ListArtifactsRequest{}
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListArtifactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsRequest) ProtoMessage() {}

func (x *ListArtifactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsRequest.ProtoReflect.Descriptor instead.
func (*ListArtifactsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_registry_v1_artifact_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListArtifactsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListArtifactsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListArtifactsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListArtifactsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListArtifactsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Artifact            `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListArtifactsResponse) Reset() {
	*x = ListArtifactsResponse{}
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListArtifactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtifactsResponse) ProtoMessage() {}

func (x *ListArtifactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtifactsResponse.ProtoReflect.Descriptor instead.
func (*ListArtifactsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_registry_v1_artifact_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListArtifactsResponse) GetItems() []*Artifact {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListArtifactsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DeleteArtifactRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteArtifactRequest) Reset() {
	*x = DeleteArtifactRequest{}
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteArtifactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtifactRequest) ProtoMessage() {}

func (x *DeleteArtifactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_registry_v1_artifact_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtifactRequest.ProtoReflect.Descriptor instead.
func (*DeleteArtifactRequest) Descriptor() ([]byte, []int) {
	return file_nebius_registry_v1_artifact_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteArtifactRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_nebius_registry_v1_artifact_service_proto protoreflect.FileDescriptor

var file_nebius_registry_v1_artifact_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x87, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x73, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2f,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x8d, 0x02, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x12, 0x5b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x65, 0x0a, 0x19, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_registry_v1_artifact_service_proto_rawDescOnce sync.Once
	file_nebius_registry_v1_artifact_service_proto_rawDescData = file_nebius_registry_v1_artifact_service_proto_rawDesc
)

func file_nebius_registry_v1_artifact_service_proto_rawDescGZIP() []byte {
	file_nebius_registry_v1_artifact_service_proto_rawDescOnce.Do(func() {
		file_nebius_registry_v1_artifact_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_registry_v1_artifact_service_proto_rawDescData)
	})
	return file_nebius_registry_v1_artifact_service_proto_rawDescData
}

var file_nebius_registry_v1_artifact_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_registry_v1_artifact_service_proto_goTypes = []any{
	(*GetArtifactRequest)(nil),    // 0: nebius.registry.v1.GetArtifactRequest
	(*ListArtifactsRequest)(nil),  // 1: nebius.registry.v1.ListArtifactsRequest
	(*ListArtifactsResponse)(nil), // 2: nebius.registry.v1.ListArtifactsResponse
	(*DeleteArtifactRequest)(nil), // 3: nebius.registry.v1.DeleteArtifactRequest
	(*Artifact)(nil),              // 4: nebius.registry.v1.Artifact
	(*v1.Operation)(nil),          // 5: nebius.common.v1.Operation
}
var file_nebius_registry_v1_artifact_service_proto_depIdxs = []int32{
	4, // 0: nebius.registry.v1.ListArtifactsResponse.items:type_name -> nebius.registry.v1.Artifact
	0, // 1: nebius.registry.v1.ArtifactService.Get:input_type -> nebius.registry.v1.GetArtifactRequest
	1, // 2: nebius.registry.v1.ArtifactService.List:input_type -> nebius.registry.v1.ListArtifactsRequest
	3, // 3: nebius.registry.v1.ArtifactService.Delete:input_type -> nebius.registry.v1.DeleteArtifactRequest
	4, // 4: nebius.registry.v1.ArtifactService.Get:output_type -> nebius.registry.v1.Artifact
	2, // 5: nebius.registry.v1.ArtifactService.List:output_type -> nebius.registry.v1.ListArtifactsResponse
	5, // 6: nebius.registry.v1.ArtifactService.Delete:output_type -> nebius.common.v1.Operation
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_registry_v1_artifact_service_proto_init() }
func file_nebius_registry_v1_artifact_service_proto_init() {
	if File_nebius_registry_v1_artifact_service_proto != nil {
		return
	}
	file_nebius_registry_v1_artifact_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_registry_v1_artifact_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_registry_v1_artifact_service_proto_goTypes,
		DependencyIndexes: file_nebius_registry_v1_artifact_service_proto_depIdxs,
		MessageInfos:      file_nebius_registry_v1_artifact_service_proto_msgTypes,
	}.Build()
	File_nebius_registry_v1_artifact_service_proto = out.File
	file_nebius_registry_v1_artifact_service_proto_rawDesc = nil
	file_nebius_registry_v1_artifact_service_proto_goTypes = nil
	file_nebius_registry_v1_artifact_service_proto_depIdxs = nil
}
