// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/msp/spark/v1alpha1/cluster.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/msp/v1alpha1"
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

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *ClusterSpec         `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *ClusterStatus       `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Cluster) GetSpec() *ClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Cluster) GetStatus() *ClusterStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Cluster specification
type ClusterSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Description of the cluster.
	Description *string `protobuf:"bytes,1,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Limits for the cluster
	Limits *Limits `protobuf:"bytes,3,opt,name=limits,proto3" json:"limits,omitempty"`
	// Password for Spark History server and Sessions.
	Authorization *Password `protobuf:"bytes,4,opt,name=authorization,proto3" json:"authorization,omitempty"`
	// ID of the user service account for accessing
	// S3 buckets in the user project
	ServiceAccountId string `protobuf:"bytes,5,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	NetworkId        string `protobuf:"bytes,6,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (x *ClusterSpec) Reset() {
	*x = ClusterSpec{}
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterSpec) ProtoMessage() {}

func (x *ClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterSpec.ProtoReflect.Descriptor instead.
func (*ClusterSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterSpec) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ClusterSpec) GetLimits() *Limits {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *ClusterSpec) GetAuthorization() *Password {
	if x != nil {
		return x.Authorization
	}
	return nil
}

func (x *ClusterSpec) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *ClusterSpec) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

type ClusterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current phase (or stage) of the cluster.
	Phase v1alpha1.ClusterStatus_Phase `protobuf:"varint,1,opt,name=phase,proto3,enum=nebius.msp.v1alpha1.ClusterStatus_Phase" json:"phase,omitempty"`
	// State reflects substatus of the stage to define whether it's healthy or not.
	State v1alpha1.ClusterStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.msp.v1alpha1.ClusterStatus_State" json:"state,omitempty"`
	// History Server WebUI endpoint
	HistoryServerEndpoint *string `protobuf:"bytes,3,opt,name=history_server_endpoint,json=historyServerEndpoint,proto3,oneof" json:"history_server_endpoint,omitempty"`
}

func (x *ClusterStatus) Reset() {
	*x = ClusterStatus{}
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatus) ProtoMessage() {}

func (x *ClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatus.ProtoReflect.Descriptor instead.
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterStatus) GetPhase() v1alpha1.ClusterStatus_Phase {
	if x != nil {
		return x.Phase
	}
	return v1alpha1.ClusterStatus_Phase(0)
}

func (x *ClusterStatus) GetState() v1alpha1.ClusterStatus_State {
	if x != nil {
		return x.State
	}
	return v1alpha1.ClusterStatus_State(0)
}

func (x *ClusterStatus) GetHistoryServerEndpoint() string {
	if x != nil && x.HistoryServerEndpoint != nil {
		return *x.HistoryServerEndpoint
	}
	return ""
}

type Limits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu             int64 `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	MemoryGibibytes int64 `protobuf:"varint,2,opt,name=memory_gibibytes,json=memoryGibibytes,proto3" json:"memory_gibibytes,omitempty"`
}

func (x *Limits) Reset() {
	*x = Limits{}
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Limits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limits) ProtoMessage() {}

func (x *Limits) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limits.ProtoReflect.Descriptor instead.
func (*Limits) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Limits) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Limits) GetMemoryGibibytes() int64 {
	if x != nil {
		return x.MemoryGibibytes
	}
	return 0
}

type Password struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Password) Reset() {
	*x = Password{}
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Password) ProtoMessage() {}

func (x *Password) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Password.ProtoReflect.Descriptor instead.
func (*Password) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *Password) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_nebius_msp_spark_v1alpha1_cluster_proto protoreflect.FileDescriptor

var file_nebius_msp_spark_v1alpha1_cluster_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdd, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xc3, 0x02, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x04, 0x52, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x12, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xba, 0x4a,
	0x01, 0x02, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0xba, 0x4a, 0x01, 0x02, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4a,
	0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xe8, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x17, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x15, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x63, 0x0a, 0x06, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x63, 0x70,
	0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x22,
	0x05, 0x18, 0x80, 0x0a, 0x20, 0x00, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x38, 0x0a, 0x10, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x69, 0x62, 0x69, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x22, 0x05, 0x18,
	0x80, 0x50, 0x20, 0x00, 0x52, 0x0f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x47, 0x69, 0x62, 0x69,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0xaa, 0x06, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x9d, 0x06, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x80, 0x06, 0xba, 0x48, 0xf9, 0x05, 0xba, 0x01, 0x58, 0x0a,
	0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x30, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x74, 0x20, 0x6c, 0x65,
	0x61, 0x73, 0x74, 0x20, 0x38, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73,
	0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x2e, 0x1a, 0x0f, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69,
	0x73, 0x29, 0x20, 0x3e, 0x3d, 0x20, 0x38, 0xba, 0x01, 0x54, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x2e, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x2b, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x6d, 0x75,
	0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x20, 0x36, 0x34,
	0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x1a, 0x10, 0x73, 0x69,
	0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x29, 0x20, 0x3c, 0x3d, 0x20, 0x36, 0x34, 0xba, 0x01,
	0x6a, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x61, 0x73, 0x63, 0x69,
	0x69, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x12, 0x38, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x20, 0x41,
	0x53, 0x43, 0x49, 0x49, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x1a, 0x19, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2c, 0x20,
	0x27, 0x5e, 0x5b, 0x21, 0x2d, 0x7e, 0x5d, 0x2b, 0x24, 0x27, 0x29, 0xba, 0x01, 0x6a, 0x0a, 0x17,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x12, 0x38, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x20, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x2e, 0x1a, 0x15, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28,
	0x27, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x27, 0x29, 0xba, 0x01, 0x6a, 0x0a, 0x17, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x63, 0x61, 0x73, 0x65, 0x12, 0x38, 0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20,
	0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x20, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x1a, 0x15,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x41,
	0x2d, 0x5a, 0x5d, 0x27, 0x29, 0xba, 0x01, 0x5b, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x2d, 0x54,
	0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x64, 0x69, 0x67, 0x69, 0x74, 0x2e, 0x1a, 0x15, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x30, 0x2d, 0x39,
	0x5d, 0x27, 0x29, 0xba, 0x01, 0x9b, 0x01, 0x0a, 0x14, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x2e, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x5a, 0x54,
	0x68, 0x65, 0x20, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x20, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x3a, 0x20, 0x21, 0x40, 0x23, 0x24, 0x25, 0x5e, 0x26,
	0x2a, 0x28, 0x29, 0x5f, 0x2b, 0x2d, 0x3d, 0x5b, 0x5d, 0x7b, 0x7d, 0x3c, 0x3e, 0x3b, 0x3a, 0x27,
	0x22, 0x5c, 0x7c, 0x2c, 0x2e, 0x2f, 0x3f, 0x60, 0x7e, 0x1a, 0x27, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5b, 0x21, 0x2d, 0x2f, 0x5d, 0x7c, 0x5b,
	0x3a, 0x2d, 0x40, 0x5d, 0x7c, 0x5b, 0x5b, 0x2d, 0x60, 0x5d, 0x7c, 0x5b, 0x7b, 0x2d, 0x7e, 0x5d,
	0x27, 0x29, 0xc8, 0x01, 0x01, 0xc0, 0x4a, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x42, 0x6b, 0x0a, 0x20, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70,
	0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescOnce sync.Once
	file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescData = file_nebius_msp_spark_v1alpha1_cluster_proto_rawDesc
)

func file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescGZIP() []byte {
	file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescOnce.Do(func() {
		file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescData)
	})
	return file_nebius_msp_spark_v1alpha1_cluster_proto_rawDescData
}

var file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_msp_spark_v1alpha1_cluster_proto_goTypes = []any{
	(*Cluster)(nil),                   // 0: nebius.msp.spark.v1alpha1.Cluster
	(*ClusterSpec)(nil),               // 1: nebius.msp.spark.v1alpha1.ClusterSpec
	(*ClusterStatus)(nil),             // 2: nebius.msp.spark.v1alpha1.ClusterStatus
	(*Limits)(nil),                    // 3: nebius.msp.spark.v1alpha1.Limits
	(*Password)(nil),                  // 4: nebius.msp.spark.v1alpha1.Password
	(*v1.ResourceMetadata)(nil),       // 5: nebius.common.v1.ResourceMetadata
	(v1alpha1.ClusterStatus_Phase)(0), // 6: nebius.msp.v1alpha1.ClusterStatus.Phase
	(v1alpha1.ClusterStatus_State)(0), // 7: nebius.msp.v1alpha1.ClusterStatus.State
}
var file_nebius_msp_spark_v1alpha1_cluster_proto_depIdxs = []int32{
	5, // 0: nebius.msp.spark.v1alpha1.Cluster.metadata:type_name -> nebius.common.v1.ResourceMetadata
	1, // 1: nebius.msp.spark.v1alpha1.Cluster.spec:type_name -> nebius.msp.spark.v1alpha1.ClusterSpec
	2, // 2: nebius.msp.spark.v1alpha1.Cluster.status:type_name -> nebius.msp.spark.v1alpha1.ClusterStatus
	3, // 3: nebius.msp.spark.v1alpha1.ClusterSpec.limits:type_name -> nebius.msp.spark.v1alpha1.Limits
	4, // 4: nebius.msp.spark.v1alpha1.ClusterSpec.authorization:type_name -> nebius.msp.spark.v1alpha1.Password
	6, // 5: nebius.msp.spark.v1alpha1.ClusterStatus.phase:type_name -> nebius.msp.v1alpha1.ClusterStatus.Phase
	7, // 6: nebius.msp.spark.v1alpha1.ClusterStatus.state:type_name -> nebius.msp.v1alpha1.ClusterStatus.State
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_nebius_msp_spark_v1alpha1_cluster_proto_init() }
func file_nebius_msp_spark_v1alpha1_cluster_proto_init() {
	if File_nebius_msp_spark_v1alpha1_cluster_proto != nil {
		return
	}
	file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[1].OneofWrappers = []any{}
	file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_spark_v1alpha1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_spark_v1alpha1_cluster_proto_goTypes,
		DependencyIndexes: file_nebius_msp_spark_v1alpha1_cluster_proto_depIdxs,
		MessageInfos:      file_nebius_msp_spark_v1alpha1_cluster_proto_msgTypes,
	}.Build()
	File_nebius_msp_spark_v1alpha1_cluster_proto = out.File
	file_nebius_msp_spark_v1alpha1_cluster_proto_rawDesc = nil
	file_nebius_msp_spark_v1alpha1_cluster_proto_goTypes = nil
	file_nebius_msp_spark_v1alpha1_cluster_proto_depIdxs = nil
}
