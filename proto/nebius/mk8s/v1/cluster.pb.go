// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/mk8s/v1/cluster.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
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

type ClusterStatus_State int32

const (
	ClusterStatus_STATE_UNSPECIFIED ClusterStatus_State = 0
	ClusterStatus_PROVISIONING      ClusterStatus_State = 1
	ClusterStatus_RUNNING           ClusterStatus_State = 2
	ClusterStatus_DELETING          ClusterStatus_State = 3
)

// Enum value maps for ClusterStatus_State.
var (
	ClusterStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PROVISIONING",
		2: "RUNNING",
		3: "DELETING",
	}
	ClusterStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PROVISIONING":      1,
		"RUNNING":           2,
		"DELETING":          3,
	}
)

func (x ClusterStatus_State) Enum() *ClusterStatus_State {
	p := new(ClusterStatus_State)
	*p = x
	return p
}

func (x ClusterStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_mk8s_v1_cluster_proto_enumTypes[0].Descriptor()
}

func (ClusterStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_mk8s_v1_cluster_proto_enumTypes[0]
}

func (x ClusterStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterStatus_State.Descriptor instead.
func (ClusterStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{6, 0}
}

type Cluster struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *ClusterSpec           `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *ClusterStatus         `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[0]
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
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{0}
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

type ClusterSpec struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	ControlPlane *ControlPlaneSpec      `protobuf:"bytes,2,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	// Defines kubernetes network configuration, like IP allocation.
	KubeNetwork   *KubeNetworkSpec `protobuf:"bytes,3,opt,name=kube_network,json=kubeNetwork,proto3" json:"kube_network,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterSpec) Reset() {
	*x = ClusterSpec{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterSpec) ProtoMessage() {}

func (x *ClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[1]
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
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterSpec) GetControlPlane() *ControlPlaneSpec {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

func (x *ClusterSpec) GetKubeNetwork() *KubeNetworkSpec {
	if x != nil {
		return x.KubeNetwork
	}
	return nil
}

type ControlPlaneSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Desired Kubernetes version of the cluster. For now only acceptable format is
	// `<major>.<minor>` like "1.31". Option for patch version update will be added later.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Nebius VPC Subnet ID where control plane instances will be located.
	// Also will be default NodeGroup subnet.
	SubnetId  string                     `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	Endpoints *ControlPlaneEndpointsSpec `protobuf:"bytes,3,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Number of instances in etcd cluster.
	// 3 by default.
	// Control plane with `etcd_cluster_size: 3` called "Highly Available" ("HA"), because it's Kubernetes API
	// will be available despite a failure of one control plane instance.
	EtcdClusterSize int64 `protobuf:"varint,4,opt,name=etcd_cluster_size,json=etcdClusterSize,proto3" json:"etcd_cluster_size,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ControlPlaneSpec) Reset() {
	*x = ControlPlaneSpec{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlPlaneSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlaneSpec) ProtoMessage() {}

func (x *ControlPlaneSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlaneSpec.ProtoReflect.Descriptor instead.
func (*ControlPlaneSpec) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ControlPlaneSpec) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ControlPlaneSpec) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *ControlPlaneSpec) GetEndpoints() *ControlPlaneEndpointsSpec {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ControlPlaneSpec) GetEtcdClusterSize() int64 {
	if x != nil {
		return x.EtcdClusterSize
	}
	return 0
}

type ControlPlaneEndpointsSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Specification of public endpoint for control plane.
	// Set value to empty, to enable it.
	PublicEndpoint *PublicEndpointSpec `protobuf:"bytes,1,opt,name=public_endpoint,json=publicEndpoint,proto3" json:"public_endpoint,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ControlPlaneEndpointsSpec) Reset() {
	*x = ControlPlaneEndpointsSpec{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlPlaneEndpointsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlaneEndpointsSpec) ProtoMessage() {}

func (x *ControlPlaneEndpointsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlaneEndpointsSpec.ProtoReflect.Descriptor instead.
func (*ControlPlaneEndpointsSpec) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ControlPlaneEndpointsSpec) GetPublicEndpoint() *PublicEndpointSpec {
	if x != nil {
		return x.PublicEndpoint
	}
	return nil
}

type PublicEndpointSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicEndpointSpec) Reset() {
	*x = PublicEndpointSpec{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicEndpointSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicEndpointSpec) ProtoMessage() {}

func (x *PublicEndpointSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicEndpointSpec.ProtoReflect.Descriptor instead.
func (*PublicEndpointSpec) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{4}
}

type KubeNetworkSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CIDR blocks for Service ClusterIP allocation.
	// For now, only one value is supported.
	// Must be a valid CIDR block or prefix length.
	// In case of prefix length, certain CIDR is auto allocated.
	// Specified CIDR blocks will be reserved in Cluster.spec.control_plane.subnet_id to prevent address duplication.
	// Allowed prefix length is from "/12" to "/28".
	// Empty value treated as ["/16"].
	ServiceCidrs  []string `protobuf:"bytes,1,rep,name=service_cidrs,json=serviceCidrs,proto3" json:"service_cidrs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubeNetworkSpec) Reset() {
	*x = KubeNetworkSpec{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubeNetworkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeNetworkSpec) ProtoMessage() {}

func (x *KubeNetworkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeNetworkSpec.ProtoReflect.Descriptor instead.
func (*KubeNetworkSpec) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *KubeNetworkSpec) GetServiceCidrs() []string {
	if x != nil {
		return x.ServiceCidrs
	}
	return nil
}

type ClusterStatus struct {
	state        protoimpl.MessageState `protogen:"open.v1"`
	State        ClusterStatus_State    `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.mk8s.v1.ClusterStatus_State" json:"state,omitempty"`
	ControlPlane *ControlPlaneStatus    `protobuf:"bytes,2,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	// Show that changes are in flight
	Reconciling   bool `protobuf:"varint,100,opt,name=reconciling,proto3" json:"reconciling,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterStatus) Reset() {
	*x = ClusterStatus{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatus) ProtoMessage() {}

func (x *ClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[6]
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
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *ClusterStatus) GetState() ClusterStatus_State {
	if x != nil {
		return x.State
	}
	return ClusterStatus_STATE_UNSPECIFIED
}

func (x *ClusterStatus) GetControlPlane() *ControlPlaneStatus {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

func (x *ClusterStatus) GetReconciling() bool {
	if x != nil {
		return x.Reconciling
	}
	return false
}

type ControlPlaneStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Actual Kubernetes and configuration version.
	// Version have format `<major>.<minor>.<patch>-nebius-cp.<infra_version>` like "1.30.0-nebius-cp.3".
	// Where <major>.<minor>.<patch> is Kubernetes version and <infra_version> is version of control plane infrastructure and configuration,
	// which update may include bug fixes, security updates and new features of components running on control plane, like CCM or Cluster Autoscaler.
	Version   string                       `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Endpoints *ControlPlaneStatusEndpoints `protobuf:"bytes,2,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Number of instances in etcd cluster.
	EtcdClusterSize int64                   `protobuf:"varint,3,opt,name=etcd_cluster_size,json=etcdClusterSize,proto3" json:"etcd_cluster_size,omitempty"`
	Auth            *ControlPlaneStatusAuth `protobuf:"bytes,100,opt,name=auth,proto3" json:"auth,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ControlPlaneStatus) Reset() {
	*x = ControlPlaneStatus{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlPlaneStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlaneStatus) ProtoMessage() {}

func (x *ControlPlaneStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlaneStatus.ProtoReflect.Descriptor instead.
func (*ControlPlaneStatus) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{7}
}

func (x *ControlPlaneStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ControlPlaneStatus) GetEndpoints() *ControlPlaneStatusEndpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ControlPlaneStatus) GetEtcdClusterSize() int64 {
	if x != nil {
		return x.EtcdClusterSize
	}
	return 0
}

func (x *ControlPlaneStatus) GetAuth() *ControlPlaneStatusAuth {
	if x != nil {
		return x.Auth
	}
	return nil
}

// Endpoints of Kubernetes control plane. Kubernetes API can be accessed at `https://endpoint/`.
type ControlPlaneStatusEndpoints struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// DNS name or IP address accessible from the Internet.
	PublicEndpoint string `protobuf:"bytes,1,opt,name=public_endpoint,json=publicEndpoint,proto3" json:"public_endpoint,omitempty"`
	// DNS name or IP address accessible from the user VPC.
	PrivateEndpoint string `protobuf:"bytes,2,opt,name=private_endpoint,json=privateEndpoint,proto3" json:"private_endpoint,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ControlPlaneStatusEndpoints) Reset() {
	*x = ControlPlaneStatusEndpoints{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlPlaneStatusEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlaneStatusEndpoints) ProtoMessage() {}

func (x *ControlPlaneStatusEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlaneStatusEndpoints.ProtoReflect.Descriptor instead.
func (*ControlPlaneStatusEndpoints) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{8}
}

func (x *ControlPlaneStatusEndpoints) GetPublicEndpoint() string {
	if x != nil {
		return x.PublicEndpoint
	}
	return ""
}

func (x *ControlPlaneStatusEndpoints) GetPrivateEndpoint() string {
	if x != nil {
		return x.PrivateEndpoint
	}
	return ""
}

type ControlPlaneStatusAuth struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// PEM with the Cluster Certificate Authority, that must be used for TLS connection to Kubernetes API.
	ClusterCaCertificate string `protobuf:"bytes,1,opt,name=cluster_ca_certificate,json=clusterCaCertificate,proto3" json:"cluster_ca_certificate,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ControlPlaneStatusAuth) Reset() {
	*x = ControlPlaneStatusAuth{}
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ControlPlaneStatusAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPlaneStatusAuth) ProtoMessage() {}

func (x *ControlPlaneStatusAuth) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_mk8s_v1_cluster_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPlaneStatusAuth.ProtoReflect.Descriptor instead.
func (*ControlPlaneStatusAuth) Descriptor() ([]byte, []int) {
	return file_nebius_mk8s_v1_cluster_proto_rawDescGZIP(), []int{9}
}

func (x *ControlPlaneStatusAuth) GetClusterCaCertificate() string {
	if x != nil {
		return x.ClusterCaCertificate
	}
	return ""
}

var File_nebius_mk8s_v1_cluster_proto protoreflect.FileDescriptor

var file_nebius_mk8s_v1_cluster_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4d, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x6b, 0x75, 0x62,
	0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0xea, 0x01,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x2c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x12, 0xba, 0x48, 0x0f, 0x72, 0x0d, 0x32, 0x0b, 0x7c, 0x5e, 0x5c, 0x64,
	0x5c, 0x2e, 0x5c, 0x64, 0x5c, 0x64, 0x24, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x52,
	0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x09, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x36, 0x0a, 0x11, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xba,
	0x48, 0x07, 0x22, 0x05, 0x32, 0x03, 0x00, 0x01, 0x03, 0x52, 0x0f, 0x65, 0x74, 0x63, 0x64, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6e, 0x0a, 0x19, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x51, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x06, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x22, 0xd1, 0x01, 0x0a, 0x0f, 0x4b, 0x75, 0x62, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x70, 0x65, 0x63, 0x12, 0xbd, 0x01, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x97, 0x01, 0xba,
	0x48, 0x8f, 0x01, 0xba, 0x01, 0x86, 0x01, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x12, 0x3f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x43, 0x49, 0x44, 0x52,
	0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x20, 0x6f, 0x72, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x20, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x22, 0x2f, 0x31,
	0x32, 0x22, 0x20, 0x74, 0x6f, 0x20, 0x22, 0x2f, 0x32, 0x38, 0x22, 0x1a, 0x30, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x61, 0x6c, 0x6c, 0x28, 0x78, 0x2c, 0x20, 0x78, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x2e, 0x2a, 0x29, 0x2f, 0x28, 0x31, 0x5b, 0x32, 0x2d, 0x39,
	0x5d, 0x7c, 0x32, 0x5b, 0x30, 0x2d, 0x38, 0x5d, 0x29, 0x24, 0x27, 0x29, 0x29, 0x92, 0x01, 0x02,
	0x10, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x69, 0x64, 0x72, 0x73, 0x22, 0x82, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d,
	0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x47, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x4b, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x22, 0xe1, 0x01, 0x0a, 0x12, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x74, 0x63, 0x64, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x65, 0x74, 0x63, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x3a, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x71, 0x0a,
	0x1b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x22, 0x4e, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x75, 0x74, 0x68, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x42, 0x55, 0x0a, 0x15, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x6d, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x6d, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_mk8s_v1_cluster_proto_rawDescOnce sync.Once
	file_nebius_mk8s_v1_cluster_proto_rawDescData []byte
)

func file_nebius_mk8s_v1_cluster_proto_rawDescGZIP() []byte {
	file_nebius_mk8s_v1_cluster_proto_rawDescOnce.Do(func() {
		file_nebius_mk8s_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_mk8s_v1_cluster_proto_rawDesc), len(file_nebius_mk8s_v1_cluster_proto_rawDesc)))
	})
	return file_nebius_mk8s_v1_cluster_proto_rawDescData
}

var file_nebius_mk8s_v1_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_mk8s_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_nebius_mk8s_v1_cluster_proto_goTypes = []any{
	(ClusterStatus_State)(0),            // 0: nebius.mk8s.v1.ClusterStatus.State
	(*Cluster)(nil),                     // 1: nebius.mk8s.v1.Cluster
	(*ClusterSpec)(nil),                 // 2: nebius.mk8s.v1.ClusterSpec
	(*ControlPlaneSpec)(nil),            // 3: nebius.mk8s.v1.ControlPlaneSpec
	(*ControlPlaneEndpointsSpec)(nil),   // 4: nebius.mk8s.v1.ControlPlaneEndpointsSpec
	(*PublicEndpointSpec)(nil),          // 5: nebius.mk8s.v1.PublicEndpointSpec
	(*KubeNetworkSpec)(nil),             // 6: nebius.mk8s.v1.KubeNetworkSpec
	(*ClusterStatus)(nil),               // 7: nebius.mk8s.v1.ClusterStatus
	(*ControlPlaneStatus)(nil),          // 8: nebius.mk8s.v1.ControlPlaneStatus
	(*ControlPlaneStatusEndpoints)(nil), // 9: nebius.mk8s.v1.ControlPlaneStatusEndpoints
	(*ControlPlaneStatusAuth)(nil),      // 10: nebius.mk8s.v1.ControlPlaneStatusAuth
	(*v1.ResourceMetadata)(nil),         // 11: nebius.common.v1.ResourceMetadata
}
var file_nebius_mk8s_v1_cluster_proto_depIdxs = []int32{
	11, // 0: nebius.mk8s.v1.Cluster.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2,  // 1: nebius.mk8s.v1.Cluster.spec:type_name -> nebius.mk8s.v1.ClusterSpec
	7,  // 2: nebius.mk8s.v1.Cluster.status:type_name -> nebius.mk8s.v1.ClusterStatus
	3,  // 3: nebius.mk8s.v1.ClusterSpec.control_plane:type_name -> nebius.mk8s.v1.ControlPlaneSpec
	6,  // 4: nebius.mk8s.v1.ClusterSpec.kube_network:type_name -> nebius.mk8s.v1.KubeNetworkSpec
	4,  // 5: nebius.mk8s.v1.ControlPlaneSpec.endpoints:type_name -> nebius.mk8s.v1.ControlPlaneEndpointsSpec
	5,  // 6: nebius.mk8s.v1.ControlPlaneEndpointsSpec.public_endpoint:type_name -> nebius.mk8s.v1.PublicEndpointSpec
	0,  // 7: nebius.mk8s.v1.ClusterStatus.state:type_name -> nebius.mk8s.v1.ClusterStatus.State
	8,  // 8: nebius.mk8s.v1.ClusterStatus.control_plane:type_name -> nebius.mk8s.v1.ControlPlaneStatus
	9,  // 9: nebius.mk8s.v1.ControlPlaneStatus.endpoints:type_name -> nebius.mk8s.v1.ControlPlaneStatusEndpoints
	10, // 10: nebius.mk8s.v1.ControlPlaneStatus.auth:type_name -> nebius.mk8s.v1.ControlPlaneStatusAuth
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_nebius_mk8s_v1_cluster_proto_init() }
func file_nebius_mk8s_v1_cluster_proto_init() {
	if File_nebius_mk8s_v1_cluster_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_mk8s_v1_cluster_proto_rawDesc), len(file_nebius_mk8s_v1_cluster_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_mk8s_v1_cluster_proto_goTypes,
		DependencyIndexes: file_nebius_mk8s_v1_cluster_proto_depIdxs,
		EnumInfos:         file_nebius_mk8s_v1_cluster_proto_enumTypes,
		MessageInfos:      file_nebius_mk8s_v1_cluster_proto_msgTypes,
	}.Build()
	File_nebius_mk8s_v1_cluster_proto = out.File
	file_nebius_mk8s_v1_cluster_proto_goTypes = nil
	file_nebius_mk8s_v1_cluster_proto_depIdxs = nil
}
