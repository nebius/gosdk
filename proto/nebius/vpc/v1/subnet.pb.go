// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/vpc/v1/subnet.proto

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

// Enumeration of possible states of the subnet.
type SubnetStatus_State int32

const (
	SubnetStatus_STATE_UNSPECIFIED SubnetStatus_State = 0 // Default state, unspecified.
	SubnetStatus_CREATING          SubnetStatus_State = 1 // Subnet is being created.
	SubnetStatus_READY             SubnetStatus_State = 2 // Subnet is ready for use.
	SubnetStatus_DELETING          SubnetStatus_State = 3 // Subnet is being deleted.
)

// Enum value maps for SubnetStatus_State.
var (
	SubnetStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
	}
	SubnetStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
	}
)

func (x SubnetStatus_State) Enum() *SubnetStatus_State {
	p := new(SubnetStatus_State)
	*p = x
	return p
}

func (x SubnetStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubnetStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1_subnet_proto_enumTypes[0].Descriptor()
}

func (SubnetStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1_subnet_proto_enumTypes[0]
}

func (x SubnetStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubnetStatus_State.Descriptor instead.
func (SubnetStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{6, 0}
}

// Defines a Subnet, a segment of a network used for more granular control and management.
// Subnet uses pools to organize address space.
type Subnet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata for the subnet resource.
	// `metadata.parent_id` represents IAM container
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification of the subnet.
	Spec *SubnetSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status of the subnet.
	Status        *SubnetStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Subnet) Reset() {
	*x = Subnet{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subnet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subnet) ProtoMessage() {}

func (x *Subnet) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subnet.ProtoReflect.Descriptor instead.
func (*Subnet) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{0}
}

func (x *Subnet) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Subnet) GetSpec() *SubnetSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Subnet) GetStatus() *SubnetStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type SubnetSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Network ID.
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Pools for private ipv4 addresses.
	// Default is 'use_network_pools = true'
	Ipv4PrivatePools *IPv4PrivateSubnetPools `protobuf:"bytes,2,opt,name=ipv4_private_pools,json=ipv4PrivatePools,proto3" json:"ipv4_private_pools,omitempty"`
	// Pools for public ipv4 addresses.
	// Default is 'use_network_pools = true'
	Ipv4PublicPools *IPv4PublicSubnetPools `protobuf:"bytes,3,opt,name=ipv4_public_pools,json=ipv4PublicPools,proto3" json:"ipv4_public_pools,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SubnetSpec) Reset() {
	*x = SubnetSpec{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubnetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetSpec) ProtoMessage() {}

func (x *SubnetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetSpec.ProtoReflect.Descriptor instead.
func (*SubnetSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{1}
}

func (x *SubnetSpec) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *SubnetSpec) GetIpv4PrivatePools() *IPv4PrivateSubnetPools {
	if x != nil {
		return x.Ipv4PrivatePools
	}
	return nil
}

func (x *SubnetSpec) GetIpv4PublicPools() *IPv4PublicSubnetPools {
	if x != nil {
		return x.Ipv4PublicPools
	}
	return nil
}

type IPv4PrivateSubnetPools struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Pools for private ipv4 allocations in subnet
	// Must be empty if 'use_network_pools = true'
	Pools []*SubnetPool `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// Allow using of private ipv4 pools which are specified in network
	// Must be false if 'pools' is not empty
	UseNetworkPools bool `protobuf:"varint,2,opt,name=use_network_pools,json=useNetworkPools,proto3" json:"use_network_pools,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *IPv4PrivateSubnetPools) Reset() {
	*x = IPv4PrivateSubnetPools{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPv4PrivateSubnetPools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4PrivateSubnetPools) ProtoMessage() {}

func (x *IPv4PrivateSubnetPools) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4PrivateSubnetPools.ProtoReflect.Descriptor instead.
func (*IPv4PrivateSubnetPools) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{2}
}

func (x *IPv4PrivateSubnetPools) GetPools() []*SubnetPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *IPv4PrivateSubnetPools) GetUseNetworkPools() bool {
	if x != nil {
		return x.UseNetworkPools
	}
	return false
}

type IPv4PublicSubnetPools struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Pools for public ipv4 allocations in subnet
	// Must be empty if 'use_network_pools = true'
	Pools []*SubnetPool `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
	// Allow using of public ipv4 pools which are specified in network
	// Must be false if 'pools' is not empty
	UseNetworkPools bool `protobuf:"varint,2,opt,name=use_network_pools,json=useNetworkPools,proto3" json:"use_network_pools,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *IPv4PublicSubnetPools) Reset() {
	*x = IPv4PublicSubnetPools{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPv4PublicSubnetPools) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4PublicSubnetPools) ProtoMessage() {}

func (x *IPv4PublicSubnetPools) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4PublicSubnetPools.ProtoReflect.Descriptor instead.
func (*IPv4PublicSubnetPools) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{3}
}

func (x *IPv4PublicSubnetPools) GetPools() []*SubnetPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *IPv4PublicSubnetPools) GetUseNetworkPools() bool {
	if x != nil {
		return x.UseNetworkPools
	}
	return false
}

type SubnetPool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cidrs         []*SubnetCidr          `protobuf:"bytes,2,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubnetPool) Reset() {
	*x = SubnetPool{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubnetPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetPool) ProtoMessage() {}

func (x *SubnetPool) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetPool.ProtoReflect.Descriptor instead.
func (*SubnetPool) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{4}
}

func (x *SubnetPool) GetCidrs() []*SubnetCidr {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

type SubnetCidr struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CIDR block.
	// May be a prefix length (such as /24) or a CIDR-formatted string (such as 10.1.2.0/24).
	Cidr string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	// State of the Cidr.
	// Default state is AVAILABLE
	State AddressBlockState `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.vpc.v1.AddressBlockState" json:"state,omitempty"`
	// Maximum mask length for allocation from this cidr
	// Default max_mask_length is 32 for IPv4 and 128 for IPv6
	MaxMaskLength int64 `protobuf:"varint,3,opt,name=max_mask_length,json=maxMaskLength,proto3" json:"max_mask_length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubnetCidr) Reset() {
	*x = SubnetCidr{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubnetCidr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetCidr) ProtoMessage() {}

func (x *SubnetCidr) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetCidr.ProtoReflect.Descriptor instead.
func (*SubnetCidr) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{5}
}

func (x *SubnetCidr) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *SubnetCidr) GetState() AddressBlockState {
	if x != nil {
		return x.State
	}
	return AddressBlockState_STATE_UNSPECIFIED
}

func (x *SubnetCidr) GetMaxMaskLength() int64 {
	if x != nil {
		return x.MaxMaskLength
	}
	return 0
}

type SubnetStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current state of the subnet.
	State SubnetStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.vpc.v1.SubnetStatus_State" json:"state,omitempty"`
	// CIDR blocks.
	Ipv4PrivateCidrs []string `protobuf:"bytes,2,rep,name=ipv4_private_cidrs,json=ipv4PrivateCidrs,proto3" json:"ipv4_private_cidrs,omitempty"`
	// CIDR blocks.
	Ipv4PublicCidrs []string `protobuf:"bytes,3,rep,name=ipv4_public_cidrs,json=ipv4PublicCidrs,proto3" json:"ipv4_public_cidrs,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *SubnetStatus) Reset() {
	*x = SubnetStatus{}
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubnetStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetStatus) ProtoMessage() {}

func (x *SubnetStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_subnet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetStatus.ProtoReflect.Descriptor instead.
func (*SubnetStatus) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_subnet_proto_rawDescGZIP(), []int{6}
}

func (x *SubnetStatus) GetState() SubnetStatus_State {
	if x != nil {
		return x.State
	}
	return SubnetStatus_STATE_UNSPECIFIED
}

func (x *SubnetStatus) GetIpv4PrivateCidrs() []string {
	if x != nil {
		return x.Ipv4PrivateCidrs
	}
	return nil
}

func (x *SubnetStatus) GetIpv4PublicCidrs() []string {
	if x != nil {
		return x.Ipv4PublicCidrs
	}
	return nil
}

var File_nebius_vpc_v1_subnet_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1_subnet_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01,
	0x0a, 0x06, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe8, 0x01, 0x0a,
	0x0a, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0a, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x49, 0x64, 0x12, 0x5a, 0x0a, 0x12, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x50, 0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x42, 0x05, 0xba, 0x4a, 0x02, 0x07, 0x06, 0x52, 0x10, 0x69, 0x70,
	0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x57,
	0x0a, 0x11, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x70, 0x6f,
	0x6f, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x42,
	0x05, 0xba, 0x4a, 0x02, 0x07, 0x06, 0x52, 0x0f, 0x69, 0x70, 0x76, 0x34, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x75, 0x0a, 0x16, 0x49, 0x50, 0x76, 0x34, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x12, 0x2f, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x70, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x75,
	0x73, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x74,
	0x0a, 0x15, 0x49, 0x50, 0x76, 0x34, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x5f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50,
	0x6f, 0x6f, 0x6c, 0x73, 0x22, 0x3d, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x2f, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x43, 0x69, 0x64, 0x72, 0x52, 0x05, 0x63, 0x69,
	0x64, 0x72, 0x73, 0x22, 0xa2, 0x03, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x43, 0x69,
	0x64, 0x72, 0x12, 0x9d, 0x02, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x88, 0x02, 0xba, 0x48, 0x84, 0x02, 0xba, 0x01, 0xac, 0x01, 0x0a, 0x11, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x12,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61,
	0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x49, 0x50, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2c, 0x20, 0x43, 0x49, 0x44, 0x52, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x61, 0x73, 0x6b, 0x1a,
	0x67, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20, 0x7c, 0x7c, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x2f, 0x28,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x5d,
	0x7c, 0x31, 0x5b, 0x30, 0x2d, 0x32, 0x5d, 0x5b, 0x30, 0x2d, 0x38, 0x5d, 0x29, 0x24, 0x27, 0x29,
	0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49, 0x70, 0x28, 0x29, 0x20,
	0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49, 0x70, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x28, 0x74, 0x72, 0x75, 0x65, 0x29, 0xba, 0x01, 0x4e, 0x0a, 0x0f, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x69, 0x70, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x20, 0x69, 0x73, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2c, 0x20, 0x77, 0x68,
	0x69, 0x63, 0x68, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x20, 0x49, 0x50, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x0a, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x63, 0x69,
	0x64, 0x72, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x07, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x36, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xba, 0x48, 0x07, 0x22, 0x05,
	0x18, 0x80, 0x01, 0x28, 0x00, 0xba, 0x4a, 0x01, 0x07, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x4d, 0x61,
	0x73, 0x6b, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x69, 0x70, 0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x69, 0x64, 0x72, 0x73,
	0x12, 0x2a, 0x0a, 0x11, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x70, 0x76,
	0x34, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x69, 0x64, 0x72, 0x73, 0x22, 0x45, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x03, 0x42, 0x52, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_vpc_v1_subnet_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1_subnet_proto_rawDescData = file_nebius_vpc_v1_subnet_proto_rawDesc
)

func file_nebius_vpc_v1_subnet_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1_subnet_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1_subnet_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_vpc_v1_subnet_proto_rawDescData)
	})
	return file_nebius_vpc_v1_subnet_proto_rawDescData
}

var file_nebius_vpc_v1_subnet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_vpc_v1_subnet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nebius_vpc_v1_subnet_proto_goTypes = []any{
	(SubnetStatus_State)(0),        // 0: nebius.vpc.v1.SubnetStatus.State
	(*Subnet)(nil),                 // 1: nebius.vpc.v1.Subnet
	(*SubnetSpec)(nil),             // 2: nebius.vpc.v1.SubnetSpec
	(*IPv4PrivateSubnetPools)(nil), // 3: nebius.vpc.v1.IPv4PrivateSubnetPools
	(*IPv4PublicSubnetPools)(nil),  // 4: nebius.vpc.v1.IPv4PublicSubnetPools
	(*SubnetPool)(nil),             // 5: nebius.vpc.v1.SubnetPool
	(*SubnetCidr)(nil),             // 6: nebius.vpc.v1.SubnetCidr
	(*SubnetStatus)(nil),           // 7: nebius.vpc.v1.SubnetStatus
	(*v1.ResourceMetadata)(nil),    // 8: nebius.common.v1.ResourceMetadata
	(AddressBlockState)(0),         // 9: nebius.vpc.v1.AddressBlockState
}
var file_nebius_vpc_v1_subnet_proto_depIdxs = []int32{
	8,  // 0: nebius.vpc.v1.Subnet.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2,  // 1: nebius.vpc.v1.Subnet.spec:type_name -> nebius.vpc.v1.SubnetSpec
	7,  // 2: nebius.vpc.v1.Subnet.status:type_name -> nebius.vpc.v1.SubnetStatus
	3,  // 3: nebius.vpc.v1.SubnetSpec.ipv4_private_pools:type_name -> nebius.vpc.v1.IPv4PrivateSubnetPools
	4,  // 4: nebius.vpc.v1.SubnetSpec.ipv4_public_pools:type_name -> nebius.vpc.v1.IPv4PublicSubnetPools
	5,  // 5: nebius.vpc.v1.IPv4PrivateSubnetPools.pools:type_name -> nebius.vpc.v1.SubnetPool
	5,  // 6: nebius.vpc.v1.IPv4PublicSubnetPools.pools:type_name -> nebius.vpc.v1.SubnetPool
	6,  // 7: nebius.vpc.v1.SubnetPool.cidrs:type_name -> nebius.vpc.v1.SubnetCidr
	9,  // 8: nebius.vpc.v1.SubnetCidr.state:type_name -> nebius.vpc.v1.AddressBlockState
	0,  // 9: nebius.vpc.v1.SubnetStatus.state:type_name -> nebius.vpc.v1.SubnetStatus.State
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1_subnet_proto_init() }
func file_nebius_vpc_v1_subnet_proto_init() {
	if File_nebius_vpc_v1_subnet_proto != nil {
		return
	}
	file_nebius_vpc_v1_pool_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_vpc_v1_subnet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_vpc_v1_subnet_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1_subnet_proto_depIdxs,
		EnumInfos:         file_nebius_vpc_v1_subnet_proto_enumTypes,
		MessageInfos:      file_nebius_vpc_v1_subnet_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1_subnet_proto = out.File
	file_nebius_vpc_v1_subnet_proto_rawDesc = nil
	file_nebius_vpc_v1_subnet_proto_goTypes = nil
	file_nebius_vpc_v1_subnet_proto_depIdxs = nil
}
