// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/msp/mlflow/v1alpha1/cluster_service.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
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

type GetClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the cluster to retrieve.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetClusterByNameRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of IAM container to get cluster from.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Name of the cluster.
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetClusterByNameRequest) Reset() {
	*x = GetClusterByNameRequest{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterByNameRequest) ProtoMessage() {}

func (x *GetClusterByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterByNameRequest.ProtoReflect.Descriptor instead.
func (*GetClusterByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetClusterByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetClusterByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListClustersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of IAM container to list clusters from.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response. Default value is 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersRequest.ProtoReflect.Descriptor instead.
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListClustersRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListClustersRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListClustersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListClustersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of clusters.
	Items []*Cluster `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListClustersResponse) Reset() {
	*x = ListClustersResponse{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersResponse) ProtoMessage() {}

func (x *ListClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersResponse.ProtoReflect.Descriptor instead.
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListClustersResponse) GetItems() []*Cluster {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListClustersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata associated with the new cluster. Must include parent_id in which we create the cluster.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification for the new cluster.
	Spec          *ClusterSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateClusterRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateClusterRequest) GetSpec() *ClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the cluster to delete.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_nebius_msp_mlflow_v1alpha1_cluster_service_proto protoreflect.FileDescriptor

var file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x6c, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d,
	0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x76, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d,
	0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd3, 0x02,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x3a, 0xad, 0x01, 0xba, 0x48, 0xa9, 0x01, 0x1a, 0x5a, 0x0a, 0x12, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x12, 0x26, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x29, 0x1a, 0x4b, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f,
	0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x68, 0x61, 0x73, 0x28,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x29, 0x22, 0x2e, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x8a, 0x04, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2d, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x65, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d,
	0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x30, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x0d, 0xba, 0x4a, 0x0a, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x6d, 0x73, 0x70,
	0x42, 0x74, 0x0a, 0x21, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescOnce sync.Once
	file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescData = file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDesc
)

func file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescGZIP() []byte {
	file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescOnce.Do(func() {
		file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescData)
	})
	return file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDescData
}

var file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_goTypes = []any{
	(*GetClusterRequest)(nil),       // 0: nebius.msp.mlflow.v1alpha1.GetClusterRequest
	(*GetClusterByNameRequest)(nil), // 1: nebius.msp.mlflow.v1alpha1.GetClusterByNameRequest
	(*ListClustersRequest)(nil),     // 2: nebius.msp.mlflow.v1alpha1.ListClustersRequest
	(*ListClustersResponse)(nil),    // 3: nebius.msp.mlflow.v1alpha1.ListClustersResponse
	(*CreateClusterRequest)(nil),    // 4: nebius.msp.mlflow.v1alpha1.CreateClusterRequest
	(*DeleteClusterRequest)(nil),    // 5: nebius.msp.mlflow.v1alpha1.DeleteClusterRequest
	(*Cluster)(nil),                 // 6: nebius.msp.mlflow.v1alpha1.Cluster
	(*v1.ResourceMetadata)(nil),     // 7: nebius.common.v1.ResourceMetadata
	(*ClusterSpec)(nil),             // 8: nebius.msp.mlflow.v1alpha1.ClusterSpec
	(*v1alpha1.Operation)(nil),      // 9: nebius.common.v1alpha1.Operation
}
var file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_depIdxs = []int32{
	6, // 0: nebius.msp.mlflow.v1alpha1.ListClustersResponse.items:type_name -> nebius.msp.mlflow.v1alpha1.Cluster
	7, // 1: nebius.msp.mlflow.v1alpha1.CreateClusterRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	8, // 2: nebius.msp.mlflow.v1alpha1.CreateClusterRequest.spec:type_name -> nebius.msp.mlflow.v1alpha1.ClusterSpec
	0, // 3: nebius.msp.mlflow.v1alpha1.ClusterService.Get:input_type -> nebius.msp.mlflow.v1alpha1.GetClusterRequest
	1, // 4: nebius.msp.mlflow.v1alpha1.ClusterService.GetByName:input_type -> nebius.msp.mlflow.v1alpha1.GetClusterByNameRequest
	2, // 5: nebius.msp.mlflow.v1alpha1.ClusterService.List:input_type -> nebius.msp.mlflow.v1alpha1.ListClustersRequest
	4, // 6: nebius.msp.mlflow.v1alpha1.ClusterService.Create:input_type -> nebius.msp.mlflow.v1alpha1.CreateClusterRequest
	5, // 7: nebius.msp.mlflow.v1alpha1.ClusterService.Delete:input_type -> nebius.msp.mlflow.v1alpha1.DeleteClusterRequest
	6, // 8: nebius.msp.mlflow.v1alpha1.ClusterService.Get:output_type -> nebius.msp.mlflow.v1alpha1.Cluster
	6, // 9: nebius.msp.mlflow.v1alpha1.ClusterService.GetByName:output_type -> nebius.msp.mlflow.v1alpha1.Cluster
	3, // 10: nebius.msp.mlflow.v1alpha1.ClusterService.List:output_type -> nebius.msp.mlflow.v1alpha1.ListClustersResponse
	9, // 11: nebius.msp.mlflow.v1alpha1.ClusterService.Create:output_type -> nebius.common.v1alpha1.Operation
	9, // 12: nebius.msp.mlflow.v1alpha1.ClusterService.Delete:output_type -> nebius.common.v1alpha1.Operation
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_init() }
func file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_init() {
	if File_nebius_msp_mlflow_v1alpha1_cluster_service_proto != nil {
		return
	}
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_goTypes,
		DependencyIndexes: file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_depIdxs,
		MessageInfos:      file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_msgTypes,
	}.Build()
	File_nebius_msp_mlflow_v1alpha1_cluster_service_proto = out.File
	file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_rawDesc = nil
	file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_goTypes = nil
	file_nebius_msp_mlflow_v1alpha1_cluster_service_proto_depIdxs = nil
}
