// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/compute/v1/image_service.proto

package v1

import (
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

type GetImageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetImageRequest) Reset() {
	*x = GetImageRequest{}
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageRequest) ProtoMessage() {}

func (x *GetImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageRequest.ProtoReflect.Descriptor instead.
func (*GetImageRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_image_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetImageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetImageLatestByFamilyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ImageFamily   string                 `protobuf:"bytes,1,opt,name=image_family,json=imageFamily,proto3" json:"image_family,omitempty"`
	ParentId      string                 `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"` // default 'project-{region}public-images'
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetImageLatestByFamilyRequest) Reset() {
	*x = GetImageLatestByFamilyRequest{}
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetImageLatestByFamilyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImageLatestByFamilyRequest) ProtoMessage() {}

func (x *GetImageLatestByFamilyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImageLatestByFamilyRequest.ProtoReflect.Descriptor instead.
func (*GetImageLatestByFamilyRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_image_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetImageLatestByFamilyRequest) GetImageFamily() string {
	if x != nil {
		return x.ImageFamily
	}
	return ""
}

func (x *GetImageLatestByFamilyRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type ListImagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter        string                 `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListImagesRequest) Reset() {
	*x = ListImagesRequest{}
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesRequest) ProtoMessage() {}

func (x *ListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesRequest.ProtoReflect.Descriptor instead.
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_image_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListImagesRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListImagesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListImagesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListImagesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListImagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Image               `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListImagesResponse) Reset() {
	*x = ListImagesResponse{}
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesResponse) ProtoMessage() {}

func (x *ListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_image_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesResponse.ProtoReflect.Descriptor instead.
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_image_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListImagesResponse) GetItems() []*Image {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListImagesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_compute_v1_image_service_proto protoreflect.FileDescriptor

var file_nebius_compute_v1_image_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x5f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x42, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x6c, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xca, 0x03, 0x0a, 0x0c, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x49,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x30,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x74, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x60, 0x0a, 0x18, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_compute_v1_image_service_proto_rawDescOnce sync.Once
	file_nebius_compute_v1_image_service_proto_rawDescData = file_nebius_compute_v1_image_service_proto_rawDesc
)

func file_nebius_compute_v1_image_service_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1_image_service_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1_image_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_compute_v1_image_service_proto_rawDescData)
	})
	return file_nebius_compute_v1_image_service_proto_rawDescData
}

var file_nebius_compute_v1_image_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_compute_v1_image_service_proto_goTypes = []any{
	(*GetImageRequest)(nil),               // 0: nebius.compute.v1.GetImageRequest
	(*GetImageLatestByFamilyRequest)(nil), // 1: nebius.compute.v1.GetImageLatestByFamilyRequest
	(*ListImagesRequest)(nil),             // 2: nebius.compute.v1.ListImagesRequest
	(*ListImagesResponse)(nil),            // 3: nebius.compute.v1.ListImagesResponse
	(*Image)(nil),                         // 4: nebius.compute.v1.Image
	(*v1.GetByNameRequest)(nil),           // 5: nebius.common.v1.GetByNameRequest
	(*ListOperationsByParentRequest)(nil), // 6: nebius.compute.v1.ListOperationsByParentRequest
	(*v1.ListOperationsResponse)(nil),     // 7: nebius.common.v1.ListOperationsResponse
}
var file_nebius_compute_v1_image_service_proto_depIdxs = []int32{
	4, // 0: nebius.compute.v1.ListImagesResponse.items:type_name -> nebius.compute.v1.Image
	0, // 1: nebius.compute.v1.ImageService.Get:input_type -> nebius.compute.v1.GetImageRequest
	5, // 2: nebius.compute.v1.ImageService.GetByName:input_type -> nebius.common.v1.GetByNameRequest
	1, // 3: nebius.compute.v1.ImageService.GetLatestByFamily:input_type -> nebius.compute.v1.GetImageLatestByFamilyRequest
	2, // 4: nebius.compute.v1.ImageService.List:input_type -> nebius.compute.v1.ListImagesRequest
	6, // 5: nebius.compute.v1.ImageService.ListOperationsByParent:input_type -> nebius.compute.v1.ListOperationsByParentRequest
	4, // 6: nebius.compute.v1.ImageService.Get:output_type -> nebius.compute.v1.Image
	4, // 7: nebius.compute.v1.ImageService.GetByName:output_type -> nebius.compute.v1.Image
	4, // 8: nebius.compute.v1.ImageService.GetLatestByFamily:output_type -> nebius.compute.v1.Image
	3, // 9: nebius.compute.v1.ImageService.List:output_type -> nebius.compute.v1.ListImagesResponse
	7, // 10: nebius.compute.v1.ImageService.ListOperationsByParent:output_type -> nebius.common.v1.ListOperationsResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1_image_service_proto_init() }
func file_nebius_compute_v1_image_service_proto_init() {
	if File_nebius_compute_v1_image_service_proto != nil {
		return
	}
	file_nebius_compute_v1_image_proto_init()
	file_nebius_compute_v1_operation_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_compute_v1_image_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_compute_v1_image_service_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1_image_service_proto_depIdxs,
		MessageInfos:      file_nebius_compute_v1_image_service_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1_image_service_proto = out.File
	file_nebius_compute_v1_image_service_proto_rawDesc = nil
	file_nebius_compute_v1_image_service_proto_goTypes = nil
	file_nebius_compute_v1_image_service_proto_depIdxs = nil
}
