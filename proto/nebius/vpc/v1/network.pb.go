// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/vpc/v1/network.proto

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

// Enumeration of possible states of the network.
type NetworkStatus_State int32

const (
	NetworkStatus_STATE_UNSPECIFIED NetworkStatus_State = 0 // Default state, unspecified.
	NetworkStatus_CREATING          NetworkStatus_State = 1 // Network is being created.
	NetworkStatus_READY             NetworkStatus_State = 2 // Network is ready for use.
	NetworkStatus_DELETING          NetworkStatus_State = 3 // Network is being deleted.
)

// Enum value maps for NetworkStatus_State.
var (
	NetworkStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
	}
	NetworkStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
	}
)

func (x NetworkStatus_State) Enum() *NetworkStatus_State {
	p := new(NetworkStatus_State)
	*p = x
	return p
}

func (x NetworkStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1_network_proto_enumTypes[0].Descriptor()
}

func (NetworkStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1_network_proto_enumTypes[0]
}

func (x NetworkStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkStatus_State.Descriptor instead.
func (NetworkStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{5, 0}
}

// Defines a Network, which serves as a virtual representation of a traditional LAN
// within a cloud environment.
// Networks facilitate communication between subnets.
type Network struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata for the network resource.
	// `metadata.parent_id` represents IAM container
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification of the network.
	Spec *NetworkSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status of the network.
	Status        *NetworkStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Network) Reset() {
	*x = Network{}
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Network) GetSpec() *NetworkSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Network) GetStatus() *NetworkStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type NetworkSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Pools for private ipv4 addresses.
	// Default private pools will be created if not specified.
	Ipv4PrivatePools *IPv4PrivateNetworkPools `protobuf:"bytes,1,opt,name=ipv4_private_pools,json=ipv4PrivatePools,proto3" json:"ipv4_private_pools,omitempty"`
	// Pools for public ipv4 addresses.
	// Default public pool will be used if not specified.
	Ipv4PublicPools *IPv4PublicNetworkPools `protobuf:"bytes,2,opt,name=ipv4_public_pools,json=ipv4PublicPools,proto3" json:"ipv4_public_pools,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NetworkSpec) Reset() {
	*x = NetworkSpec{}
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkSpec) ProtoMessage() {}

func (x *NetworkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkSpec.ProtoReflect.Descriptor instead.
func (*NetworkSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkSpec) GetIpv4PrivatePools() *IPv4PrivateNetworkPools {
	if x != nil {
		return x.Ipv4PrivatePools
	}
	return nil
}

func (x *NetworkSpec) GetIpv4PublicPools() *IPv4PublicNetworkPools {
	if x != nil {
		return x.Ipv4PublicPools
	}
	return nil
}

type IPv4PrivateNetworkPools struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pools         []*NetworkPool         `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPv4PrivateNetworkPools) Reset() {
	*x = IPv4PrivateNetworkPools{}
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPv4PrivateNetworkPools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4PrivateNetworkPools) ProtoMessage() {}

func (x *IPv4PrivateNetworkPools) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4PrivateNetworkPools.ProtoReflect.Descriptor instead.
func (*IPv4PrivateNetworkPools) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{2}
}

func (x *IPv4PrivateNetworkPools) GetPools() []*NetworkPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

type IPv4PublicNetworkPools struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pools         []*NetworkPool         `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPv4PublicNetworkPools) Reset() {
	*x = IPv4PublicNetworkPools{}
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPv4PublicNetworkPools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4PublicNetworkPools) ProtoMessage() {}

func (x *IPv4PublicNetworkPools) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4PublicNetworkPools.ProtoReflect.Descriptor instead.
func (*IPv4PublicNetworkPools) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{3}
}

func (x *IPv4PublicNetworkPools) GetPools() []*NetworkPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

type NetworkPool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkPool) Reset() {
	*x = NetworkPool{}
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkPool) ProtoMessage() {}

func (x *NetworkPool) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkPool.ProtoReflect.Descriptor instead.
func (*NetworkPool) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkPool) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NetworkStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current state of the network.
	State         NetworkStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.vpc.v1.NetworkStatus_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkStatus) Reset() {
	*x = NetworkStatus{}
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkStatus) ProtoMessage() {}

func (x *NetworkStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_network_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkStatus.ProtoReflect.Descriptor instead.
func (*NetworkStatus) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_network_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkStatus) GetState() NetworkStatus_State {
	if x != nil {
		return x.State
	}
	return NetworkStatus_STATE_UNSPECIFIED
}

var File_nebius_vpc_v1_network_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1_network_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5a, 0x0a, 0x12, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x07,
	0x52, 0x10, 0x69, 0x70, 0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50,
	0x76, 0x34, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x07, 0x52, 0x0f, 0x69, 0x70, 0x76, 0x34,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x4b, 0x0a, 0x17, 0x49,
	0x50, 0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6f,
	0x6c, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x4a, 0x0a, 0x16, 0x49, 0x50, 0x76, 0x34,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x70,
	0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x0d,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x53,
	0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e,
	0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_vpc_v1_network_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1_network_proto_rawDescData = file_nebius_vpc_v1_network_proto_rawDesc
)

func file_nebius_vpc_v1_network_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1_network_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_vpc_v1_network_proto_rawDescData)
	})
	return file_nebius_vpc_v1_network_proto_rawDescData
}

var file_nebius_vpc_v1_network_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_vpc_v1_network_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nebius_vpc_v1_network_proto_goTypes = []any{
	(NetworkStatus_State)(0),        // 0: nebius.vpc.v1.NetworkStatus.State
	(*Network)(nil),                 // 1: nebius.vpc.v1.Network
	(*NetworkSpec)(nil),             // 2: nebius.vpc.v1.NetworkSpec
	(*IPv4PrivateNetworkPools)(nil), // 3: nebius.vpc.v1.IPv4PrivateNetworkPools
	(*IPv4PublicNetworkPools)(nil),  // 4: nebius.vpc.v1.IPv4PublicNetworkPools
	(*NetworkPool)(nil),             // 5: nebius.vpc.v1.NetworkPool
	(*NetworkStatus)(nil),           // 6: nebius.vpc.v1.NetworkStatus
	(*v1.ResourceMetadata)(nil),     // 7: nebius.common.v1.ResourceMetadata
}
var file_nebius_vpc_v1_network_proto_depIdxs = []int32{
	7, // 0: nebius.vpc.v1.Network.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2, // 1: nebius.vpc.v1.Network.spec:type_name -> nebius.vpc.v1.NetworkSpec
	6, // 2: nebius.vpc.v1.Network.status:type_name -> nebius.vpc.v1.NetworkStatus
	3, // 3: nebius.vpc.v1.NetworkSpec.ipv4_private_pools:type_name -> nebius.vpc.v1.IPv4PrivateNetworkPools
	4, // 4: nebius.vpc.v1.NetworkSpec.ipv4_public_pools:type_name -> nebius.vpc.v1.IPv4PublicNetworkPools
	5, // 5: nebius.vpc.v1.IPv4PrivateNetworkPools.pools:type_name -> nebius.vpc.v1.NetworkPool
	5, // 6: nebius.vpc.v1.IPv4PublicNetworkPools.pools:type_name -> nebius.vpc.v1.NetworkPool
	0, // 7: nebius.vpc.v1.NetworkStatus.state:type_name -> nebius.vpc.v1.NetworkStatus.State
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1_network_proto_init() }
func file_nebius_vpc_v1_network_proto_init() {
	if File_nebius_vpc_v1_network_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_vpc_v1_network_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_vpc_v1_network_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1_network_proto_depIdxs,
		EnumInfos:         file_nebius_vpc_v1_network_proto_enumTypes,
		MessageInfos:      file_nebius_vpc_v1_network_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1_network_proto = out.File
	file_nebius_vpc_v1_network_proto_rawDesc = nil
	file_nebius_vpc_v1_network_proto_goTypes = nil
	file_nebius_vpc_v1_network_proto_depIdxs = nil
}
