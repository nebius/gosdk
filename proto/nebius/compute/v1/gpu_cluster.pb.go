// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/compute/v1/gpu_cluster.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
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

type GpuCluster struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *GpuClusterSpec        `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *GpuClusterStatus      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GpuCluster) Reset() {
	*x = GpuCluster{}
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GpuCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuCluster) ProtoMessage() {}

func (x *GpuCluster) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuCluster.ProtoReflect.Descriptor instead.
func (*GpuCluster) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *GpuCluster) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GpuCluster) GetSpec() *GpuClusterSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GpuCluster) GetStatus() *GpuClusterStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type GpuClusterSpec struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	InfinibandFabric string                 `protobuf:"bytes,1,opt,name=infiniband_fabric,json=infinibandFabric,proto3" json:"infiniband_fabric,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GpuClusterSpec) Reset() {
	*x = GpuClusterSpec{}
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GpuClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuClusterSpec) ProtoMessage() {}

func (x *GpuClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuClusterSpec.ProtoReflect.Descriptor instead.
func (*GpuClusterSpec) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *GpuClusterSpec) GetInfinibandFabric() string {
	if x != nil {
		return x.InfinibandFabric
	}
	return ""
}

type GpuClusterStatus struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Instances []string               `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	// Indicates whether there is an ongoing operation
	Reconciling bool `protobuf:"varint,2,opt,name=reconciling,proto3" json:"reconciling,omitempty"`
	// Types that are valid to be assigned to Topology:
	//
	//	*GpuClusterStatus_InfinibandTopologyPath
	Topology      isGpuClusterStatus_Topology `protobuf_oneof:"topology"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GpuClusterStatus) Reset() {
	*x = GpuClusterStatus{}
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GpuClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuClusterStatus) ProtoMessage() {}

func (x *GpuClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuClusterStatus.ProtoReflect.Descriptor instead.
func (*GpuClusterStatus) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *GpuClusterStatus) GetInstances() []string {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *GpuClusterStatus) GetReconciling() bool {
	if x != nil {
		return x.Reconciling
	}
	return false
}

func (x *GpuClusterStatus) GetTopology() isGpuClusterStatus_Topology {
	if x != nil {
		return x.Topology
	}
	return nil
}

func (x *GpuClusterStatus) GetInfinibandTopologyPath() *GpuClusterStatusInfinibandTopologyPath {
	if x != nil {
		if x, ok := x.Topology.(*GpuClusterStatus_InfinibandTopologyPath); ok {
			return x.InfinibandTopologyPath
		}
	}
	return nil
}

type isGpuClusterStatus_Topology interface {
	isGpuClusterStatus_Topology()
}

type GpuClusterStatus_InfinibandTopologyPath struct {
	InfinibandTopologyPath *GpuClusterStatusInfinibandTopologyPath `protobuf:"bytes,3,opt,name=infiniband_topology_path,json=infinibandTopologyPath,proto3,oneof"`
}

func (*GpuClusterStatus_InfinibandTopologyPath) isGpuClusterStatus_Topology() {}

type GpuClusterStatusInfinibandTopologyPath struct {
	state         protoimpl.MessageState                            `protogen:"open.v1"`
	Instances     []*GpuClusterStatusInfinibandTopologyPathInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GpuClusterStatusInfinibandTopologyPath) Reset() {
	*x = GpuClusterStatusInfinibandTopologyPath{}
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GpuClusterStatusInfinibandTopologyPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuClusterStatusInfinibandTopologyPath) ProtoMessage() {}

func (x *GpuClusterStatusInfinibandTopologyPath) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuClusterStatusInfinibandTopologyPath.ProtoReflect.Descriptor instead.
func (*GpuClusterStatusInfinibandTopologyPath) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *GpuClusterStatusInfinibandTopologyPath) GetInstances() []*GpuClusterStatusInfinibandTopologyPathInstance {
	if x != nil {
		return x.Instances
	}
	return nil
}

type GpuClusterStatusInfinibandTopologyPathInstance struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Path          []string               `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GpuClusterStatusInfinibandTopologyPathInstance) Reset() {
	*x = GpuClusterStatusInfinibandTopologyPathInstance{}
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GpuClusterStatusInfinibandTopologyPathInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuClusterStatusInfinibandTopologyPathInstance) ProtoMessage() {}

func (x *GpuClusterStatusInfinibandTopologyPathInstance) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_gpu_cluster_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuClusterStatusInfinibandTopologyPathInstance.ProtoReflect.Descriptor instead.
func (*GpuClusterStatusInfinibandTopologyPathInstance) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *GpuClusterStatusInfinibandTopologyPathInstance) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GpuClusterStatusInfinibandTopologyPathInstance) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_nebius_compute_v1_gpu_cluster_proto protoreflect.FileDescriptor

var file_nebius_compute_v1_gpu_cluster_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc0, 0x01, 0x0a, 0x0a, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x35, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x0e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x37, 0x0a, 0x11, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62,
	0x61, 0x6e, 0x64, 0x5f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x10, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x62, 0x61, 0x6e, 0x64, 0x46, 0x61, 0x62, 0x72, 0x69, 0x63, 0x22, 0xd5,
	0x01, 0x0a, 0x10, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x12, 0x75, 0x0a, 0x18, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62, 0x61, 0x6e,
	0x64, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69,
	0x62, 0x61, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x48, 0x00, 0x52, 0x16, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62, 0x61, 0x6e, 0x64, 0x54, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x50, 0x61, 0x74, 0x68, 0x42, 0x0a, 0x0a, 0x08, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x26, 0x47, 0x70, 0x75, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x62, 0x61, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x5f, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62,
	0x61, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x50, 0x61, 0x74, 0x68, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x22, 0x65, 0x0a, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x62, 0x61, 0x6e, 0x64,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x50, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x5e, 0x0a, 0x18, 0x61, 0x69, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_nebius_compute_v1_gpu_cluster_proto_rawDescOnce sync.Once
	file_nebius_compute_v1_gpu_cluster_proto_rawDescData = file_nebius_compute_v1_gpu_cluster_proto_rawDesc
)

func file_nebius_compute_v1_gpu_cluster_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1_gpu_cluster_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1_gpu_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_compute_v1_gpu_cluster_proto_rawDescData)
	})
	return file_nebius_compute_v1_gpu_cluster_proto_rawDescData
}

var file_nebius_compute_v1_gpu_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_compute_v1_gpu_cluster_proto_goTypes = []any{
	(*GpuCluster)(nil),                                     // 0: nebius.compute.v1.GpuCluster
	(*GpuClusterSpec)(nil),                                 // 1: nebius.compute.v1.GpuClusterSpec
	(*GpuClusterStatus)(nil),                               // 2: nebius.compute.v1.GpuClusterStatus
	(*GpuClusterStatusInfinibandTopologyPath)(nil),         // 3: nebius.compute.v1.GpuClusterStatusInfinibandTopologyPath
	(*GpuClusterStatusInfinibandTopologyPathInstance)(nil), // 4: nebius.compute.v1.GpuClusterStatusInfinibandTopologyPathInstance
	(*v1.ResourceMetadata)(nil),                            // 5: nebius.common.v1.ResourceMetadata
}
var file_nebius_compute_v1_gpu_cluster_proto_depIdxs = []int32{
	5, // 0: nebius.compute.v1.GpuCluster.metadata:type_name -> nebius.common.v1.ResourceMetadata
	1, // 1: nebius.compute.v1.GpuCluster.spec:type_name -> nebius.compute.v1.GpuClusterSpec
	2, // 2: nebius.compute.v1.GpuCluster.status:type_name -> nebius.compute.v1.GpuClusterStatus
	3, // 3: nebius.compute.v1.GpuClusterStatus.infiniband_topology_path:type_name -> nebius.compute.v1.GpuClusterStatusInfinibandTopologyPath
	4, // 4: nebius.compute.v1.GpuClusterStatusInfinibandTopologyPath.instances:type_name -> nebius.compute.v1.GpuClusterStatusInfinibandTopologyPathInstance
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1_gpu_cluster_proto_init() }
func file_nebius_compute_v1_gpu_cluster_proto_init() {
	if File_nebius_compute_v1_gpu_cluster_proto != nil {
		return
	}
	file_nebius_compute_v1_gpu_cluster_proto_msgTypes[2].OneofWrappers = []any{
		(*GpuClusterStatus_InfinibandTopologyPath)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_compute_v1_gpu_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_compute_v1_gpu_cluster_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1_gpu_cluster_proto_depIdxs,
		MessageInfos:      file_nebius_compute_v1_gpu_cluster_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1_gpu_cluster_proto = out.File
	file_nebius_compute_v1_gpu_cluster_proto_rawDesc = nil
	file_nebius_compute_v1_gpu_cluster_proto_goTypes = nil
	file_nebius_compute_v1_gpu_cluster_proto_depIdxs = nil
}
