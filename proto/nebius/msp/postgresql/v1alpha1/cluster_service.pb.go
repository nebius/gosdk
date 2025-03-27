// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/msp/postgresql/v1alpha1/cluster_service.proto

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
	// ID of the PostgreSQL Cluster resource to return.
	// To get the cluster ID use a [ClusterService.List] request.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[0]
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
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListClustersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of IAM container to list clusters from.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[1]
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
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{1}
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
	// List of PostgreSQL Cluster resources.
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListClustersRequest.page_size], use the [next_page_token] as the value
	// for the [ListClustersRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListClustersResponse) Reset() {
	*x = ListClustersResponse{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersResponse) ProtoMessage() {}

func (x *ListClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[2]
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
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListClustersResponse) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
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
	// Metadata associated with the new cluster.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification for the new cluster.
	Spec          *ClusterSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[3]
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
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{3}
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

type RestoreClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata associated with the new cluster.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification for the new cluster.
	Spec *ClusterSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// ID of the backup to restore from.
	BackupId string `protobuf:"bytes,3,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	// ID of the PostgreSQL cluster to restore from.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RestoreClusterRequest) Reset() {
	*x = RestoreClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestoreClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreClusterRequest) ProtoMessage() {}

func (x *RestoreClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreClusterRequest.ProtoReflect.Descriptor instead.
func (*RestoreClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{4}
}

func (x *RestoreClusterRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RestoreClusterRequest) GetSpec() *ClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *RestoreClusterRequest) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

func (x *RestoreClusterRequest) GetSourceClusterId() string {
	if x != nil {
		return x.SourceClusterId
	}
	return ""
}

type DeleteClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster to delete.
	// To get the PostgreSQL cluster ID, use a [ClusterService.List] request.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteClusterRequest) Reset() {
	*x = DeleteClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClusterRequest) ProtoMessage() {}

func (x *DeleteClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[5]
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
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata associated with the cluster. Must include id of the cluster we are going to update.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Updated specification for the cluster.
	Spec          *ClusterSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateClusterRequest) Reset() {
	*x = UpdateClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClusterRequest) ProtoMessage() {}

func (x *UpdateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClusterRequest.ProtoReflect.Descriptor instead.
func (*UpdateClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateClusterRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateClusterRequest) GetSpec() *ClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type StopClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL Cluster resource to pause.
	// To get the cluster ID use a [ClusterService.List] request.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StopClusterRequest) Reset() {
	*x = StopClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StopClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopClusterRequest) ProtoMessage() {}

func (x *StopClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopClusterRequest.ProtoReflect.Descriptor instead.
func (*StopClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{7}
}

func (x *StopClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StartClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL Cluster resource to resume.
	// To get the cluster ID use a [ClusterService.List] request.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartClusterRequest) Reset() {
	*x = StartClusterRequest{}
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartClusterRequest) ProtoMessage() {}

func (x *StartClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartClusterRequest.ProtoReflect.Descriptor instead.
func (*StartClusterRequest) Descriptor() ([]byte, []int) {
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP(), []int{8}
}

func (x *StartClusterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_nebius_msp_postgresql_v1alpha1_cluster_service_proto protoreflect.FileDescriptor

var file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x76, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x83, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xa7, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x47, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x81, 0x02, 0x0a, 0x15, 0x52, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71,
	0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x11, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x2c, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x70, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x32, 0xa1, 0x07, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x31,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x34, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x61, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x63, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x32, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x33, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0xba, 0x4a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x71, 0x6c, 0x2e, 0x6d, 0x73, 0x70, 0x42, 0x7c, 0x0a, 0x25, 0x61, 0x69, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x13, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d,
	0x73, 0x70, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescOnce sync.Once
	file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescData = file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDesc
)

func file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescGZIP() []byte {
	file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescOnce.Do(func() {
		file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescData)
	})
	return file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDescData
}

var file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_goTypes = []any{
	(*GetClusterRequest)(nil),     // 0: nebius.msp.postgresql.v1alpha1.GetClusterRequest
	(*ListClustersRequest)(nil),   // 1: nebius.msp.postgresql.v1alpha1.ListClustersRequest
	(*ListClustersResponse)(nil),  // 2: nebius.msp.postgresql.v1alpha1.ListClustersResponse
	(*CreateClusterRequest)(nil),  // 3: nebius.msp.postgresql.v1alpha1.CreateClusterRequest
	(*RestoreClusterRequest)(nil), // 4: nebius.msp.postgresql.v1alpha1.RestoreClusterRequest
	(*DeleteClusterRequest)(nil),  // 5: nebius.msp.postgresql.v1alpha1.DeleteClusterRequest
	(*UpdateClusterRequest)(nil),  // 6: nebius.msp.postgresql.v1alpha1.UpdateClusterRequest
	(*StopClusterRequest)(nil),    // 7: nebius.msp.postgresql.v1alpha1.StopClusterRequest
	(*StartClusterRequest)(nil),   // 8: nebius.msp.postgresql.v1alpha1.StartClusterRequest
	(*Cluster)(nil),               // 9: nebius.msp.postgresql.v1alpha1.Cluster
	(*v1.ResourceMetadata)(nil),   // 10: nebius.common.v1.ResourceMetadata
	(*ClusterSpec)(nil),           // 11: nebius.msp.postgresql.v1alpha1.ClusterSpec
	(*v1.GetByNameRequest)(nil),   // 12: nebius.common.v1.GetByNameRequest
	(*v1alpha1.Operation)(nil),    // 13: nebius.common.v1alpha1.Operation
}
var file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_depIdxs = []int32{
	9,  // 0: nebius.msp.postgresql.v1alpha1.ListClustersResponse.clusters:type_name -> nebius.msp.postgresql.v1alpha1.Cluster
	10, // 1: nebius.msp.postgresql.v1alpha1.CreateClusterRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	11, // 2: nebius.msp.postgresql.v1alpha1.CreateClusterRequest.spec:type_name -> nebius.msp.postgresql.v1alpha1.ClusterSpec
	10, // 3: nebius.msp.postgresql.v1alpha1.RestoreClusterRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	11, // 4: nebius.msp.postgresql.v1alpha1.RestoreClusterRequest.spec:type_name -> nebius.msp.postgresql.v1alpha1.ClusterSpec
	10, // 5: nebius.msp.postgresql.v1alpha1.UpdateClusterRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	11, // 6: nebius.msp.postgresql.v1alpha1.UpdateClusterRequest.spec:type_name -> nebius.msp.postgresql.v1alpha1.ClusterSpec
	0,  // 7: nebius.msp.postgresql.v1alpha1.ClusterService.Get:input_type -> nebius.msp.postgresql.v1alpha1.GetClusterRequest
	12, // 8: nebius.msp.postgresql.v1alpha1.ClusterService.GetByName:input_type -> nebius.common.v1.GetByNameRequest
	1,  // 9: nebius.msp.postgresql.v1alpha1.ClusterService.List:input_type -> nebius.msp.postgresql.v1alpha1.ListClustersRequest
	3,  // 10: nebius.msp.postgresql.v1alpha1.ClusterService.Create:input_type -> nebius.msp.postgresql.v1alpha1.CreateClusterRequest
	5,  // 11: nebius.msp.postgresql.v1alpha1.ClusterService.Delete:input_type -> nebius.msp.postgresql.v1alpha1.DeleteClusterRequest
	6,  // 12: nebius.msp.postgresql.v1alpha1.ClusterService.Update:input_type -> nebius.msp.postgresql.v1alpha1.UpdateClusterRequest
	4,  // 13: nebius.msp.postgresql.v1alpha1.ClusterService.Restore:input_type -> nebius.msp.postgresql.v1alpha1.RestoreClusterRequest
	7,  // 14: nebius.msp.postgresql.v1alpha1.ClusterService.Stop:input_type -> nebius.msp.postgresql.v1alpha1.StopClusterRequest
	8,  // 15: nebius.msp.postgresql.v1alpha1.ClusterService.Start:input_type -> nebius.msp.postgresql.v1alpha1.StartClusterRequest
	9,  // 16: nebius.msp.postgresql.v1alpha1.ClusterService.Get:output_type -> nebius.msp.postgresql.v1alpha1.Cluster
	9,  // 17: nebius.msp.postgresql.v1alpha1.ClusterService.GetByName:output_type -> nebius.msp.postgresql.v1alpha1.Cluster
	2,  // 18: nebius.msp.postgresql.v1alpha1.ClusterService.List:output_type -> nebius.msp.postgresql.v1alpha1.ListClustersResponse
	13, // 19: nebius.msp.postgresql.v1alpha1.ClusterService.Create:output_type -> nebius.common.v1alpha1.Operation
	13, // 20: nebius.msp.postgresql.v1alpha1.ClusterService.Delete:output_type -> nebius.common.v1alpha1.Operation
	13, // 21: nebius.msp.postgresql.v1alpha1.ClusterService.Update:output_type -> nebius.common.v1alpha1.Operation
	13, // 22: nebius.msp.postgresql.v1alpha1.ClusterService.Restore:output_type -> nebius.common.v1alpha1.Operation
	13, // 23: nebius.msp.postgresql.v1alpha1.ClusterService.Stop:output_type -> nebius.common.v1alpha1.Operation
	13, // 24: nebius.msp.postgresql.v1alpha1.ClusterService.Start:output_type -> nebius.common.v1alpha1.Operation
	16, // [16:25] is the sub-list for method output_type
	7,  // [7:16] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_init() }
func file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_init() {
	if File_nebius_msp_postgresql_v1alpha1_cluster_service_proto != nil {
		return
	}
	file_nebius_msp_postgresql_v1alpha1_cluster_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_goTypes,
		DependencyIndexes: file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_depIdxs,
		MessageInfos:      file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_msgTypes,
	}.Build()
	File_nebius_msp_postgresql_v1alpha1_cluster_service_proto = out.File
	file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_rawDesc = nil
	file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_goTypes = nil
	file_nebius_msp_postgresql_v1alpha1_cluster_service_proto_depIdxs = nil
}
