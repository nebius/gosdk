// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/mk8s/v1alpha1/node_group_service.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateNodeGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *NodeGroupSpec         `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNodeGroupRequest) Reset() {
	*x = CreateNodeGroupRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeGroupRequest) ProtoMessage() {}

func (x *CreateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeGroupRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateNodeGroupRequest) GetSpec() *NodeGroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type GetNodeGroupRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ResourceVersion string                 `protobuf:"bytes,2,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetNodeGroupRequest) Reset() {
	*x = GetNodeGroupRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeGroupRequest) ProtoMessage() {}

func (x *GetNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*GetNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodeGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetNodeGroupRequest) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

type GetNodeGroupByNameRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParentId      string                 `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNodeGroupByNameRequest) Reset() {
	*x = GetNodeGroupByNameRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNodeGroupByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeGroupByNameRequest) ProtoMessage() {}

func (x *GetNodeGroupByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeGroupByNameRequest.ProtoReflect.Descriptor instead.
func (*GetNodeGroupByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetNodeGroupByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetNodeGroupByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListNodeGroupsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the parent Cluster.
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNodeGroupsRequest) Reset() {
	*x = ListNodeGroupsRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNodeGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeGroupsRequest) ProtoMessage() {}

func (x *ListNodeGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListNodeGroupsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListNodeGroupsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListNodeGroupsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListNodeGroupsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListNodeGroupsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*NodeGroup           `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNodeGroupsResponse) Reset() {
	*x = ListNodeGroupsResponse{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNodeGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeGroupsResponse) ProtoMessage() {}

func (x *ListNodeGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListNodeGroupsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListNodeGroupsResponse) GetItems() []*NodeGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListNodeGroupsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type UpdateNodeGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *NodeGroupSpec         `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateNodeGroupRequest) Reset() {
	*x = UpdateNodeGroupRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeGroupRequest) ProtoMessage() {}

func (x *UpdateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNodeGroupRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateNodeGroupRequest) GetSpec() *NodeGroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeleteNodeGroupRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteNodeGroupRequest) Reset() {
	*x = DeleteNodeGroupRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeGroupRequest) ProtoMessage() {}

func (x *DeleteNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteNodeGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpgradeNodeGroupRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Id    string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to UpgradeType:
	//
	//	*UpgradeNodeGroupRequest_LatestInfraVersion
	UpgradeType   isUpgradeNodeGroupRequest_UpgradeType `protobuf_oneof:"upgrade_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpgradeNodeGroupRequest) Reset() {
	*x = UpgradeNodeGroupRequest{}
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpgradeNodeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeNodeGroupRequest) ProtoMessage() {}

func (x *UpgradeNodeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeNodeGroupRequest.ProtoReflect.Descriptor instead.
func (*UpgradeNodeGroupRequest) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpgradeNodeGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpgradeNodeGroupRequest) GetUpgradeType() isUpgradeNodeGroupRequest_UpgradeType {
	if x != nil {
		return x.UpgradeType
	}
	return nil
}

func (x *UpgradeNodeGroupRequest) GetLatestInfraVersion() *emptypb.Empty {
	if x != nil {
		if x, ok := x.UpgradeType.(*UpgradeNodeGroupRequest_LatestInfraVersion); ok {
			return x.LatestInfraVersion
		}
	}
	return nil
}

type isUpgradeNodeGroupRequest_UpgradeType interface {
	isUpgradeNodeGroupRequest_UpgradeType()
}

type UpgradeNodeGroupRequest_LatestInfraVersion struct {
	// Upgrades to the latest infra version, which includes latest supported kubernetes patch version. Kubernetes minor version remain the same.
	LatestInfraVersion *emptypb.Empty `protobuf:"bytes,2,opt,name=latest_infra_version,json=latestInfraVersion,proto3,oneof"`
}

func (*UpgradeNodeGroupRequest_LatestInfraVersion) isUpgradeNodeGroupRequest_UpgradeType() {}

var File_nebius_mk8s_v1alpha1_node_group_service_proto protoreflect.FileDescriptor

var file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDesc = string([]byte{
	0x0a, 0x2d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x58, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x77, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x30, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x50, 0x0a, 0x14, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x04, 0xba, 0x4a, 0x01, 0x06, 0x48, 0x00, 0x52, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x15, 0x0a, 0x0c,
	0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x05, 0xba, 0x48,
	0x02, 0x08, 0x01, 0x32, 0x95, 0x05, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x5d, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x61, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2c, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b,
	0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x2d, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6a, 0x0a, 0x1b, 0x61,
	0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x6b, 0x38,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x15, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescOnce sync.Once
	file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescData []byte
)

func file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescGZIP() []byte {
	file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescOnce.Do(func() {
		file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDesc), len(file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDesc)))
	})
	return file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDescData
}

var file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_nebius_mk8s_v1alpha1_node_group_service_proto_goTypes = []any{
	(*CreateNodeGroupRequest)(nil),    // 0: nebius.mk8s.v1alpha1.CreateNodeGroupRequest
	(*GetNodeGroupRequest)(nil),       // 1: nebius.mk8s.v1alpha1.GetNodeGroupRequest
	(*GetNodeGroupByNameRequest)(nil), // 2: nebius.mk8s.v1alpha1.GetNodeGroupByNameRequest
	(*ListNodeGroupsRequest)(nil),     // 3: nebius.mk8s.v1alpha1.ListNodeGroupsRequest
	(*ListNodeGroupsResponse)(nil),    // 4: nebius.mk8s.v1alpha1.ListNodeGroupsResponse
	(*UpdateNodeGroupRequest)(nil),    // 5: nebius.mk8s.v1alpha1.UpdateNodeGroupRequest
	(*DeleteNodeGroupRequest)(nil),    // 6: nebius.mk8s.v1alpha1.DeleteNodeGroupRequest
	(*UpgradeNodeGroupRequest)(nil),   // 7: nebius.mk8s.v1alpha1.UpgradeNodeGroupRequest
	(*v1.ResourceMetadata)(nil),       // 8: nebius.common.v1.ResourceMetadata
	(*NodeGroupSpec)(nil),             // 9: nebius.mk8s.v1alpha1.NodeGroupSpec
	(*NodeGroup)(nil),                 // 10: nebius.mk8s.v1alpha1.NodeGroup
	(*emptypb.Empty)(nil),             // 11: google.protobuf.Empty
	(*v1alpha1.Operation)(nil),        // 12: nebius.common.v1alpha1.Operation
}
var file_nebius_mk8s_v1alpha1_node_group_service_proto_depIdxs = []int32{
	8,  // 0: nebius.mk8s.v1alpha1.CreateNodeGroupRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	9,  // 1: nebius.mk8s.v1alpha1.CreateNodeGroupRequest.spec:type_name -> nebius.mk8s.v1alpha1.NodeGroupSpec
	10, // 2: nebius.mk8s.v1alpha1.ListNodeGroupsResponse.items:type_name -> nebius.mk8s.v1alpha1.NodeGroup
	8,  // 3: nebius.mk8s.v1alpha1.UpdateNodeGroupRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	9,  // 4: nebius.mk8s.v1alpha1.UpdateNodeGroupRequest.spec:type_name -> nebius.mk8s.v1alpha1.NodeGroupSpec
	11, // 5: nebius.mk8s.v1alpha1.UpgradeNodeGroupRequest.latest_infra_version:type_name -> google.protobuf.Empty
	1,  // 6: nebius.mk8s.v1alpha1.NodeGroupService.Get:input_type -> nebius.mk8s.v1alpha1.GetNodeGroupRequest
	2,  // 7: nebius.mk8s.v1alpha1.NodeGroupService.GetByName:input_type -> nebius.mk8s.v1alpha1.GetNodeGroupByNameRequest
	3,  // 8: nebius.mk8s.v1alpha1.NodeGroupService.List:input_type -> nebius.mk8s.v1alpha1.ListNodeGroupsRequest
	0,  // 9: nebius.mk8s.v1alpha1.NodeGroupService.Create:input_type -> nebius.mk8s.v1alpha1.CreateNodeGroupRequest
	5,  // 10: nebius.mk8s.v1alpha1.NodeGroupService.Update:input_type -> nebius.mk8s.v1alpha1.UpdateNodeGroupRequest
	6,  // 11: nebius.mk8s.v1alpha1.NodeGroupService.Delete:input_type -> nebius.mk8s.v1alpha1.DeleteNodeGroupRequest
	7,  // 12: nebius.mk8s.v1alpha1.NodeGroupService.Upgrade:input_type -> nebius.mk8s.v1alpha1.UpgradeNodeGroupRequest
	10, // 13: nebius.mk8s.v1alpha1.NodeGroupService.Get:output_type -> nebius.mk8s.v1alpha1.NodeGroup
	10, // 14: nebius.mk8s.v1alpha1.NodeGroupService.GetByName:output_type -> nebius.mk8s.v1alpha1.NodeGroup
	4,  // 15: nebius.mk8s.v1alpha1.NodeGroupService.List:output_type -> nebius.mk8s.v1alpha1.ListNodeGroupsResponse
	12, // 16: nebius.mk8s.v1alpha1.NodeGroupService.Create:output_type -> nebius.common.v1alpha1.Operation
	12, // 17: nebius.mk8s.v1alpha1.NodeGroupService.Update:output_type -> nebius.common.v1alpha1.Operation
	12, // 18: nebius.mk8s.v1alpha1.NodeGroupService.Delete:output_type -> nebius.common.v1alpha1.Operation
	12, // 19: nebius.mk8s.v1alpha1.NodeGroupService.Upgrade:output_type -> nebius.common.v1alpha1.Operation
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_mk8s_v1alpha1_node_group_service_proto_init() }
func file_nebius_mk8s_v1alpha1_node_group_service_proto_init() {
	if File_nebius_mk8s_v1alpha1_node_group_service_proto != nil {
		return
	}
	file_nebius_mk8s_v1alpha1_node_group_proto_init()
	file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes[7].OneofWrappers = []any{
		(*UpgradeNodeGroupRequest_LatestInfraVersion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDesc), len(file_nebius_mk8s_v1alpha1_node_group_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_mk8s_v1alpha1_node_group_service_proto_goTypes,
		DependencyIndexes: file_nebius_mk8s_v1alpha1_node_group_service_proto_depIdxs,
		MessageInfos:      file_nebius_mk8s_v1alpha1_node_group_service_proto_msgTypes,
	}.Build()
	File_nebius_mk8s_v1alpha1_node_group_service_proto = out.File
	file_nebius_mk8s_v1alpha1_node_group_service_proto_goTypes = nil
	file_nebius_mk8s_v1alpha1_node_group_service_proto_depIdxs = nil
}
