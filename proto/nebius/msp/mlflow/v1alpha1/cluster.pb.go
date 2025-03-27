// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/msp/mlflow/v1alpha1/cluster.proto

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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *ClusterSpec           `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *MlflowClusterStatus   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[0]
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
	return file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescGZIP(), []int{0}
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

func (x *Cluster) GetStatus() *MlflowClusterStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Cluster specification
type ClusterSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Description of the cluster.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Either make cluster public accessible or accessible only via private VPC.
	PublicAccess bool `protobuf:"varint,2,opt,name=public_access,json=publicAccess,proto3" json:"public_access,omitempty"`
	// MLflow admin username.
	AdminUsername string `protobuf:"bytes,4,opt,name=admin_username,json=adminUsername,proto3" json:"admin_username,omitempty"`
	// MLflow admin password.
	AdminPassword string `protobuf:"bytes,5,opt,name=admin_password,json=adminPassword,proto3" json:"admin_password,omitempty"`
	// Id of the service account that will be used to access S3 bucket (and create one if not provided).
	ServiceAccountId string `protobuf:"bytes,6,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Name of the Nebius S3 bucket for MLflow artifacts. If not provided, will be created under the same parent.
	StorageBucketName string `protobuf:"bytes,7,opt,name=storage_bucket_name,json=storageBucketName,proto3" json:"storage_bucket_name,omitempty"`
	// ID of the vpc network.
	NetworkId     string `protobuf:"bytes,8,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterSpec) Reset() {
	*x = ClusterSpec{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterSpec) ProtoMessage() {}

func (x *ClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[1]
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
	return file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ClusterSpec) GetPublicAccess() bool {
	if x != nil {
		return x.PublicAccess
	}
	return false
}

func (x *ClusterSpec) GetAdminUsername() string {
	if x != nil {
		return x.AdminUsername
	}
	return ""
}

func (x *ClusterSpec) GetAdminPassword() string {
	if x != nil {
		return x.AdminPassword
	}
	return ""
}

func (x *ClusterSpec) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *ClusterSpec) GetStorageBucketName() string {
	if x != nil {
		return x.StorageBucketName
	}
	return ""
}

func (x *ClusterSpec) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

type MlflowClusterStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current phase of the cluster.
	Phase v1alpha1.ClusterStatus_Phase `protobuf:"varint,1,opt,name=phase,proto3,enum=nebius.msp.v1alpha1.ClusterStatus_Phase" json:"phase,omitempty"`
	// State reflects substatus of the phase to define whether it's healthy or not.
	State v1alpha1.ClusterStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.msp.v1alpha1.ClusterStatus_State" json:"state,omitempty"`
	// Tracking endpoint url.
	// Will be removed soon in favor of private_tracking_endpoint and public_tracking_endpoint.
	TrackingEndpoint string `protobuf:"bytes,3,opt,name=tracking_endpoint,json=trackingEndpoint,proto3" json:"tracking_endpoint,omitempty"`
	// Name of the Nebius S3 bucket for MLflow artifacts.
	EffectiveStorageBucketName string `protobuf:"bytes,4,opt,name=effective_storage_bucket_name,json=effectiveStorageBucketName,proto3" json:"effective_storage_bucket_name,omitempty"`
	// Count of experiments in the MLflow cluster
	ExperimentsCount uint32 `protobuf:"varint,5,opt,name=experiments_count,json=experimentsCount,proto3" json:"experiments_count,omitempty"`
	// MLflow version
	MlflowVersion string `protobuf:"bytes,6,opt,name=mlflow_version,json=mlflowVersion,proto3" json:"mlflow_version,omitempty"`
	// Public and private tracking endpoints
	TrackingEndpoints *Endpoints `protobuf:"bytes,7,opt,name=tracking_endpoints,json=trackingEndpoints,proto3" json:"tracking_endpoints,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *MlflowClusterStatus) Reset() {
	*x = MlflowClusterStatus{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MlflowClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MlflowClusterStatus) ProtoMessage() {}

func (x *MlflowClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MlflowClusterStatus.ProtoReflect.Descriptor instead.
func (*MlflowClusterStatus) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *MlflowClusterStatus) GetPhase() v1alpha1.ClusterStatus_Phase {
	if x != nil {
		return x.Phase
	}
	return v1alpha1.ClusterStatus_Phase(0)
}

func (x *MlflowClusterStatus) GetState() v1alpha1.ClusterStatus_State {
	if x != nil {
		return x.State
	}
	return v1alpha1.ClusterStatus_State(0)
}

func (x *MlflowClusterStatus) GetTrackingEndpoint() string {
	if x != nil {
		return x.TrackingEndpoint
	}
	return ""
}

func (x *MlflowClusterStatus) GetEffectiveStorageBucketName() string {
	if x != nil {
		return x.EffectiveStorageBucketName
	}
	return ""
}

func (x *MlflowClusterStatus) GetExperimentsCount() uint32 {
	if x != nil {
		return x.ExperimentsCount
	}
	return 0
}

func (x *MlflowClusterStatus) GetMlflowVersion() string {
	if x != nil {
		return x.MlflowVersion
	}
	return ""
}

func (x *MlflowClusterStatus) GetTrackingEndpoints() *Endpoints {
	if x != nil {
		return x.TrackingEndpoints
	}
	return nil
}

type Endpoints struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Private endpoint
	Private string `protobuf:"bytes,1,opt,name=private,proto3" json:"private,omitempty"`
	// Public endpoint
	Public        string `protobuf:"bytes,2,opt,name=public,proto3" json:"public,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Endpoints) Reset() {
	*x = Endpoints{}
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoints) ProtoMessage() {}

func (x *Endpoints) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoints.ProtoReflect.Descriptor instead.
func (*Endpoints) Descriptor() ([]byte, []int) {
	return file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Endpoints) GetPrivate() string {
	if x != nil {
		return x.Private
	}
	return ""
}

func (x *Endpoints) GetPublic() string {
	if x != nil {
		return x.Public
	}
	return ""
}

var File_nebius_msp_mlflow_v1alpha1_cluster_proto protoreflect.FileDescriptor

var file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x6c, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd0, 0x02, 0x0a, 0x0b, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x2d, 0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x0e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xba,
	0x4a, 0x01, 0x04, 0xc0, 0x4a, 0x01, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x34, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x09, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0xaf, 0x03, 0x0a,
	0x13, 0x4d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x05, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x41, 0x0a, 0x1d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x10, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6c, 0x66, 0x6c, 0x6f,
	0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x11, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3d,
	0x0a, 0x09, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x6d, 0x0a,
	0x21, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x6d, 0x6c, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x6d, 0x6c, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescOnce sync.Once
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescData = file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDesc
)

func file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescGZIP() []byte {
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescOnce.Do(func() {
		file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescData)
	})
	return file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDescData
}

var file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_msp_mlflow_v1alpha1_cluster_proto_goTypes = []any{
	(*Cluster)(nil),                   // 0: nebius.msp.mlflow.v1alpha1.Cluster
	(*ClusterSpec)(nil),               // 1: nebius.msp.mlflow.v1alpha1.ClusterSpec
	(*MlflowClusterStatus)(nil),       // 2: nebius.msp.mlflow.v1alpha1.MlflowClusterStatus
	(*Endpoints)(nil),                 // 3: nebius.msp.mlflow.v1alpha1.Endpoints
	(*v1.ResourceMetadata)(nil),       // 4: nebius.common.v1.ResourceMetadata
	(v1alpha1.ClusterStatus_Phase)(0), // 5: nebius.msp.v1alpha1.ClusterStatus.Phase
	(v1alpha1.ClusterStatus_State)(0), // 6: nebius.msp.v1alpha1.ClusterStatus.State
}
var file_nebius_msp_mlflow_v1alpha1_cluster_proto_depIdxs = []int32{
	4, // 0: nebius.msp.mlflow.v1alpha1.Cluster.metadata:type_name -> nebius.common.v1.ResourceMetadata
	1, // 1: nebius.msp.mlflow.v1alpha1.Cluster.spec:type_name -> nebius.msp.mlflow.v1alpha1.ClusterSpec
	2, // 2: nebius.msp.mlflow.v1alpha1.Cluster.status:type_name -> nebius.msp.mlflow.v1alpha1.MlflowClusterStatus
	5, // 3: nebius.msp.mlflow.v1alpha1.MlflowClusterStatus.phase:type_name -> nebius.msp.v1alpha1.ClusterStatus.Phase
	6, // 4: nebius.msp.mlflow.v1alpha1.MlflowClusterStatus.state:type_name -> nebius.msp.v1alpha1.ClusterStatus.State
	3, // 5: nebius.msp.mlflow.v1alpha1.MlflowClusterStatus.tracking_endpoints:type_name -> nebius.msp.mlflow.v1alpha1.Endpoints
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_msp_mlflow_v1alpha1_cluster_proto_init() }
func file_nebius_msp_mlflow_v1alpha1_cluster_proto_init() {
	if File_nebius_msp_mlflow_v1alpha1_cluster_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_mlflow_v1alpha1_cluster_proto_goTypes,
		DependencyIndexes: file_nebius_msp_mlflow_v1alpha1_cluster_proto_depIdxs,
		MessageInfos:      file_nebius_msp_mlflow_v1alpha1_cluster_proto_msgTypes,
	}.Build()
	File_nebius_msp_mlflow_v1alpha1_cluster_proto = out.File
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_rawDesc = nil
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_goTypes = nil
	file_nebius_msp_mlflow_v1alpha1_cluster_proto_depIdxs = nil
}
