// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/compute/v1/gpu_cluster_service.proto

package v1

import (
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

type GetGpuClusterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGpuClusterRequest) Reset() {
	*x = GetGpuClusterRequest{}
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGpuClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGpuClusterRequest) ProtoMessage() {}

func (x *GetGpuClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGpuClusterRequest.ProtoReflect.Descriptor instead.
func (*GetGpuClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetGpuClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListGpuClustersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	PageSize      int64                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter        string                 `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGpuClustersRequest) Reset() {
	*x = ListGpuClustersRequest{}
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGpuClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpuClustersRequest) ProtoMessage() {}

func (x *ListGpuClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpuClustersRequest.ProtoReflect.Descriptor instead.
func (*ListGpuClustersRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListGpuClustersRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListGpuClustersRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGpuClustersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListGpuClustersRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type CreateGpuClusterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *GpuClusterSpec        `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateGpuClusterRequest) Reset() {
	*x = CreateGpuClusterRequest{}
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGpuClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGpuClusterRequest) ProtoMessage() {}

func (x *CreateGpuClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGpuClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateGpuClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGpuClusterRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateGpuClusterRequest) GetSpec() *GpuClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type UpdateGpuClusterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *GpuClusterSpec        `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateGpuClusterRequest) Reset() {
	*x = UpdateGpuClusterRequest{}
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateGpuClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGpuClusterRequest) ProtoMessage() {}

func (x *UpdateGpuClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGpuClusterRequest.ProtoReflect.Descriptor instead.
func (*UpdateGpuClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGpuClusterRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateGpuClusterRequest) GetSpec() *GpuClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteGpuClusterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteGpuClusterRequest) Reset() {
	*x = DeleteGpuClusterRequest{}
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGpuClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGpuClusterRequest) ProtoMessage() {}

func (x *DeleteGpuClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGpuClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteGpuClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteGpuClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListGpuClustersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*GpuCluster          `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGpuClustersResponse) Reset() {
	*x = ListGpuClustersResponse{}
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGpuClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGpuClustersResponse) ProtoMessage() {}

func (x *ListGpuClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGpuClustersResponse.ProtoReflect.Descriptor instead.
func (*ListGpuClustersResponse) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListGpuClustersResponse) GetItems() []*GpuCluster {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListGpuClustersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_compute_v1_gpu_cluster_service_proto protoreflect.FileDescriptor

var file_nebius_compute_v1_gpu_cluster_service_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70,
	0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x90, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x70, 0x75, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x90, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x35, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x76, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x80, 0x05, 0x0a, 0x11, 0x47,
	0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x5d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x70, 0x75, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x51, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x74, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x65, 0x0a,
	0x18, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x47, 0x70, 0x75, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_compute_v1_gpu_cluster_service_proto_rawDescOnce sync.Once
	file_nebius_compute_v1_gpu_cluster_service_proto_rawDescData []byte
)

func file_nebius_compute_v1_gpu_cluster_service_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1_gpu_cluster_service_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1_gpu_cluster_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_compute_v1_gpu_cluster_service_proto_rawDesc), len(file_nebius_compute_v1_gpu_cluster_service_proto_rawDesc)))
	})
	return file_nebius_compute_v1_gpu_cluster_service_proto_rawDescData
}

var file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nebius_compute_v1_gpu_cluster_service_proto_goTypes = []any{
	(*GetGpuClusterRequest)(nil),          // 0: nebius.compute.v1.GetGpuClusterRequest
	(*ListGpuClustersRequest)(nil),        // 1: nebius.compute.v1.ListGpuClustersRequest
	(*CreateGpuClusterRequest)(nil),       // 2: nebius.compute.v1.CreateGpuClusterRequest
	(*UpdateGpuClusterRequest)(nil),       // 3: nebius.compute.v1.UpdateGpuClusterRequest
	(*DeleteGpuClusterRequest)(nil),       // 4: nebius.compute.v1.DeleteGpuClusterRequest
	(*ListGpuClustersResponse)(nil),       // 5: nebius.compute.v1.ListGpuClustersResponse
	(*v1.ResourceMetadata)(nil),           // 6: nebius.common.v1.ResourceMetadata
	(*GpuClusterSpec)(nil),                // 7: nebius.compute.v1.GpuClusterSpec
	(*GpuCluster)(nil),                    // 8: nebius.compute.v1.GpuCluster
	(*v1.GetByNameRequest)(nil),           // 9: nebius.common.v1.GetByNameRequest
	(*ListOperationsByParentRequest)(nil), // 10: nebius.compute.v1.ListOperationsByParentRequest
	(*v1.Operation)(nil),                  // 11: nebius.common.v1.Operation
	(*v1.ListOperationsResponse)(nil),     // 12: nebius.common.v1.ListOperationsResponse
}
var file_nebius_compute_v1_gpu_cluster_service_proto_depIdxs = []int32{
	6,  // 0: nebius.compute.v1.CreateGpuClusterRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	7,  // 1: nebius.compute.v1.CreateGpuClusterRequest.spec:type_name -> nebius.compute.v1.GpuClusterSpec
	6,  // 2: nebius.compute.v1.UpdateGpuClusterRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	7,  // 3: nebius.compute.v1.UpdateGpuClusterRequest.spec:type_name -> nebius.compute.v1.GpuClusterSpec
	8,  // 4: nebius.compute.v1.ListGpuClustersResponse.items:type_name -> nebius.compute.v1.GpuCluster
	0,  // 5: nebius.compute.v1.GpuClusterService.Get:input_type -> nebius.compute.v1.GetGpuClusterRequest
	9,  // 6: nebius.compute.v1.GpuClusterService.GetByName:input_type -> nebius.common.v1.GetByNameRequest
	1,  // 7: nebius.compute.v1.GpuClusterService.List:input_type -> nebius.compute.v1.ListGpuClustersRequest
	2,  // 8: nebius.compute.v1.GpuClusterService.Create:input_type -> nebius.compute.v1.CreateGpuClusterRequest
	3,  // 9: nebius.compute.v1.GpuClusterService.Update:input_type -> nebius.compute.v1.UpdateGpuClusterRequest
	4,  // 10: nebius.compute.v1.GpuClusterService.Delete:input_type -> nebius.compute.v1.DeleteGpuClusterRequest
	10, // 11: nebius.compute.v1.GpuClusterService.ListOperationsByParent:input_type -> nebius.compute.v1.ListOperationsByParentRequest
	8,  // 12: nebius.compute.v1.GpuClusterService.Get:output_type -> nebius.compute.v1.GpuCluster
	8,  // 13: nebius.compute.v1.GpuClusterService.GetByName:output_type -> nebius.compute.v1.GpuCluster
	5,  // 14: nebius.compute.v1.GpuClusterService.List:output_type -> nebius.compute.v1.ListGpuClustersResponse
	11, // 15: nebius.compute.v1.GpuClusterService.Create:output_type -> nebius.common.v1.Operation
	11, // 16: nebius.compute.v1.GpuClusterService.Update:output_type -> nebius.common.v1.Operation
	11, // 17: nebius.compute.v1.GpuClusterService.Delete:output_type -> nebius.common.v1.Operation
	12, // 18: nebius.compute.v1.GpuClusterService.ListOperationsByParent:output_type -> nebius.common.v1.ListOperationsResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1_gpu_cluster_service_proto_init() }
func file_nebius_compute_v1_gpu_cluster_service_proto_init() {
	if File_nebius_compute_v1_gpu_cluster_service_proto != nil {
		return
	}
	file_nebius_compute_v1_gpu_cluster_proto_init()
	file_nebius_compute_v1_operation_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_compute_v1_gpu_cluster_service_proto_rawDesc), len(file_nebius_compute_v1_gpu_cluster_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_compute_v1_gpu_cluster_service_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1_gpu_cluster_service_proto_depIdxs,
		MessageInfos:      file_nebius_compute_v1_gpu_cluster_service_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1_gpu_cluster_service_proto = out.File
	file_nebius_compute_v1_gpu_cluster_service_proto_goTypes = nil
	file_nebius_compute_v1_gpu_cluster_service_proto_depIdxs = nil
}
