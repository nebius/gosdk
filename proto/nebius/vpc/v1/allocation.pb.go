// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/vpc/v1/allocation.proto

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

// Enumeration of possible states of the Allocation.
type AllocationStatus_State int32

const (
	AllocationStatus_STATE_UNSPECIFIED AllocationStatus_State = 0 // Default state, unspecified.
	AllocationStatus_CREATING          AllocationStatus_State = 1 // Allocation is being created.
	AllocationStatus_ALLOCATED         AllocationStatus_State = 2 // Allocation is ready for use.
	AllocationStatus_ASSIGNED          AllocationStatus_State = 3 // Allocation is used.
	AllocationStatus_DELETING          AllocationStatus_State = 4 // Allocation is being deleted.
)

// Enum value maps for AllocationStatus_State.
var (
	AllocationStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "ALLOCATED",
		3: "ASSIGNED",
		4: "DELETING",
	}
	AllocationStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"ALLOCATED":         2,
		"ASSIGNED":          3,
		"DELETING":          4,
	}
)

func (x AllocationStatus_State) Enum() *AllocationStatus_State {
	p := new(AllocationStatus_State)
	*p = x
	return p
}

func (x AllocationStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AllocationStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1_allocation_proto_enumTypes[0].Descriptor()
}

func (AllocationStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1_allocation_proto_enumTypes[0]
}

func (x AllocationStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AllocationStatus_State.Descriptor instead.
func (AllocationStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{4, 0}
}

type Allocation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata for the Allocation.
	// `metadata.parent_id` represents IAM Container.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specifications for the allocation, detailing its name and IP configuration.
	Spec *AllocationSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Contains the current status of the allocation, indicating its state and any additional details.
	Status        *AllocationStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Allocation) Reset() {
	*x = Allocation{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Allocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Allocation) ProtoMessage() {}

func (x *Allocation) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Allocation.ProtoReflect.Descriptor instead.
func (*Allocation) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{0}
}

func (x *Allocation) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Allocation) GetSpec() *AllocationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Allocation) GetStatus() *AllocationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type AllocationSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Holds the IP specifications for the allocation, including the type of IP (IPv4 or IPv6) and its corresponding configuration.
	//
	// Types that are valid to be assigned to IpSpec:
	//
	//	*AllocationSpec_Ipv4Private
	//	*AllocationSpec_Ipv4Public
	IpSpec        isAllocationSpec_IpSpec `protobuf_oneof:"ip_spec"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocationSpec) Reset() {
	*x = AllocationSpec{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationSpec) ProtoMessage() {}

func (x *AllocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationSpec.ProtoReflect.Descriptor instead.
func (*AllocationSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{1}
}

func (x *AllocationSpec) GetIpSpec() isAllocationSpec_IpSpec {
	if x != nil {
		return x.IpSpec
	}
	return nil
}

func (x *AllocationSpec) GetIpv4Private() *IPv4PrivateAllocationSpec {
	if x != nil {
		if x, ok := x.IpSpec.(*AllocationSpec_Ipv4Private); ok {
			return x.Ipv4Private
		}
	}
	return nil
}

func (x *AllocationSpec) GetIpv4Public() *IPv4PublicAllocationSpec {
	if x != nil {
		if x, ok := x.IpSpec.(*AllocationSpec_Ipv4Public); ok {
			return x.Ipv4Public
		}
	}
	return nil
}

type isAllocationSpec_IpSpec interface {
	isAllocationSpec_IpSpec()
}

type AllocationSpec_Ipv4Private struct {
	Ipv4Private *IPv4PrivateAllocationSpec `protobuf:"bytes,1,opt,name=ipv4_private,json=ipv4Private,proto3,oneof"`
}

type AllocationSpec_Ipv4Public struct {
	Ipv4Public *IPv4PublicAllocationSpec `protobuf:"bytes,2,opt,name=ipv4_public,json=ipv4Public,proto3,oneof"`
}

func (*AllocationSpec_Ipv4Private) isAllocationSpec_IpSpec() {}

func (*AllocationSpec_Ipv4Public) isAllocationSpec_IpSpec() {}

type IPv4PrivateAllocationSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CIDR block for IPv4 Allocation.
	// May be a single IP address (such as 10.2.3.4),
	// a prefix length (such as /24) or a CIDR-formatted string (such as 10.1.2.0/24).
	// Random address (/32) from pool would be allocated if field is omitted.
	Cidr string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	// Types that are valid to be assigned to Pool:
	//
	//	*IPv4PrivateAllocationSpec_SubnetId
	//	*IPv4PrivateAllocationSpec_PoolId
	Pool          isIPv4PrivateAllocationSpec_Pool `protobuf_oneof:"pool"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPv4PrivateAllocationSpec) Reset() {
	*x = IPv4PrivateAllocationSpec{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPv4PrivateAllocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4PrivateAllocationSpec) ProtoMessage() {}

func (x *IPv4PrivateAllocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4PrivateAllocationSpec.ProtoReflect.Descriptor instead.
func (*IPv4PrivateAllocationSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{2}
}

func (x *IPv4PrivateAllocationSpec) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *IPv4PrivateAllocationSpec) GetPool() isIPv4PrivateAllocationSpec_Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

func (x *IPv4PrivateAllocationSpec) GetSubnetId() string {
	if x != nil {
		if x, ok := x.Pool.(*IPv4PrivateAllocationSpec_SubnetId); ok {
			return x.SubnetId
		}
	}
	return ""
}

func (x *IPv4PrivateAllocationSpec) GetPoolId() string {
	if x != nil {
		if x, ok := x.Pool.(*IPv4PrivateAllocationSpec_PoolId); ok {
			return x.PoolId
		}
	}
	return ""
}

type isIPv4PrivateAllocationSpec_Pool interface {
	isIPv4PrivateAllocationSpec_Pool()
}

type IPv4PrivateAllocationSpec_SubnetId struct {
	// Subnet ID.
	// Required same subnet to use allocation in subnet-resources (e.g. Network Interface)
	SubnetId string `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3,oneof"`
}

type IPv4PrivateAllocationSpec_PoolId struct {
	// Pool for the IPv4 private allocation.
	PoolId string `protobuf:"bytes,3,opt,name=pool_id,json=poolId,proto3,oneof"`
}

func (*IPv4PrivateAllocationSpec_SubnetId) isIPv4PrivateAllocationSpec_Pool() {}

func (*IPv4PrivateAllocationSpec_PoolId) isIPv4PrivateAllocationSpec_Pool() {}

type IPv4PublicAllocationSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CIDR block for IPv4 Allocation.
	// May be a single IP address (such as 10.2.3.4),
	// a prefix length (such as /32) or a CIDR-formatted string (such as 10.1.2.0/32).
	// Random address (/32) from pool would be allocated if field is omitted.
	Cidr string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	// Types that are valid to be assigned to Pool:
	//
	//	*IPv4PublicAllocationSpec_SubnetId
	//	*IPv4PublicAllocationSpec_PoolId
	Pool          isIPv4PublicAllocationSpec_Pool `protobuf_oneof:"pool"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IPv4PublicAllocationSpec) Reset() {
	*x = IPv4PublicAllocationSpec{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IPv4PublicAllocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4PublicAllocationSpec) ProtoMessage() {}

func (x *IPv4PublicAllocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4PublicAllocationSpec.ProtoReflect.Descriptor instead.
func (*IPv4PublicAllocationSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{3}
}

func (x *IPv4PublicAllocationSpec) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *IPv4PublicAllocationSpec) GetPool() isIPv4PublicAllocationSpec_Pool {
	if x != nil {
		return x.Pool
	}
	return nil
}

func (x *IPv4PublicAllocationSpec) GetSubnetId() string {
	if x != nil {
		if x, ok := x.Pool.(*IPv4PublicAllocationSpec_SubnetId); ok {
			return x.SubnetId
		}
	}
	return ""
}

func (x *IPv4PublicAllocationSpec) GetPoolId() string {
	if x != nil {
		if x, ok := x.Pool.(*IPv4PublicAllocationSpec_PoolId); ok {
			return x.PoolId
		}
	}
	return ""
}

type isIPv4PublicAllocationSpec_Pool interface {
	isIPv4PublicAllocationSpec_Pool()
}

type IPv4PublicAllocationSpec_SubnetId struct {
	// Subnet ID.
	// Required same subnet to use allocation in subnet-resources (e.g. Network Interface)
	SubnetId string `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3,oneof"`
}

type IPv4PublicAllocationSpec_PoolId struct {
	// Pool for the IPv4 public allocation.
	PoolId string `protobuf:"bytes,3,opt,name=pool_id,json=poolId,proto3,oneof"`
}

func (*IPv4PublicAllocationSpec_SubnetId) isIPv4PublicAllocationSpec_Pool() {}

func (*IPv4PublicAllocationSpec_PoolId) isIPv4PublicAllocationSpec_Pool() {}

type AllocationStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field represents the current state of the allocation.
	State AllocationStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.vpc.v1.AllocationStatus_State" json:"state,omitempty"`
	// Detailed information about the allocation status,
	// including the allocated CIDR, pool ID and IP version.
	Details *AllocationDetails `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	// Information about the assignment associated with the allocation,
	// such as network interface or load balancer assignment.
	Assignment *Assignment `protobuf:"bytes,3,opt,name=assignment,proto3" json:"assignment,omitempty"`
	// If false - Lifecycle of allocation depends on resource that using it.
	Static        bool `protobuf:"varint,4,opt,name=static,proto3" json:"static,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocationStatus) Reset() {
	*x = AllocationStatus{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationStatus) ProtoMessage() {}

func (x *AllocationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationStatus.ProtoReflect.Descriptor instead.
func (*AllocationStatus) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{4}
}

func (x *AllocationStatus) GetState() AllocationStatus_State {
	if x != nil {
		return x.State
	}
	return AllocationStatus_STATE_UNSPECIFIED
}

func (x *AllocationStatus) GetDetails() *AllocationDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *AllocationStatus) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

func (x *AllocationStatus) GetStatic() bool {
	if x != nil {
		return x.Static
	}
	return false
}

type AllocationDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The actual CIDR block that has been allocated.
	AllocatedCidr string `protobuf:"bytes,1,opt,name=allocated_cidr,json=allocatedCidr,proto3" json:"allocated_cidr,omitempty"`
	// ID of the pool from which this allocation was made.
	PoolId string `protobuf:"bytes,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// The IP version of this allocation (IPv4 or IPv6).
	Version       IpVersion `protobuf:"varint,4,opt,name=version,proto3,enum=nebius.vpc.v1.IpVersion" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AllocationDetails) Reset() {
	*x = AllocationDetails{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AllocationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationDetails) ProtoMessage() {}

func (x *AllocationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationDetails.ProtoReflect.Descriptor instead.
func (*AllocationDetails) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{5}
}

func (x *AllocationDetails) GetAllocatedCidr() string {
	if x != nil {
		return x.AllocatedCidr
	}
	return ""
}

func (x *AllocationDetails) GetPoolId() string {
	if x != nil {
		return x.PoolId
	}
	return ""
}

func (x *AllocationDetails) GetVersion() IpVersion {
	if x != nil {
		return x.Version
	}
	return IpVersion_IP_VERSION_UNSPECIFIED
}

type Assignment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field specifies the type of assignment associated with the allocation,
	// which could be a network interface or load balancer assignment.
	//
	// Types that are valid to be assigned to Type:
	//
	//	*Assignment_NetworkInterface
	//	*Assignment_LoadBalancer
	Type          isAssignment_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Assignment) Reset() {
	*x = Assignment{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Assignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assignment) ProtoMessage() {}

func (x *Assignment) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assignment.ProtoReflect.Descriptor instead.
func (*Assignment) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{6}
}

func (x *Assignment) GetType() isAssignment_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Assignment) GetNetworkInterface() *NetworkInterfaceAssignment {
	if x != nil {
		if x, ok := x.Type.(*Assignment_NetworkInterface); ok {
			return x.NetworkInterface
		}
	}
	return nil
}

func (x *Assignment) GetLoadBalancer() *LoadBalancerAssignment {
	if x != nil {
		if x, ok := x.Type.(*Assignment_LoadBalancer); ok {
			return x.LoadBalancer
		}
	}
	return nil
}

type isAssignment_Type interface {
	isAssignment_Type()
}

type Assignment_NetworkInterface struct {
	NetworkInterface *NetworkInterfaceAssignment `protobuf:"bytes,1,opt,name=network_interface,json=networkInterface,proto3,oneof"`
}

type Assignment_LoadBalancer struct {
	LoadBalancer *LoadBalancerAssignment `protobuf:"bytes,2,opt,name=load_balancer,json=loadBalancer,proto3,oneof"`
}

func (*Assignment_NetworkInterface) isAssignment_Type() {}

func (*Assignment_LoadBalancer) isAssignment_Type() {}

type NetworkInterfaceAssignment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Compute instance network interface belongs to.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Network interface name
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetworkInterfaceAssignment) Reset() {
	*x = NetworkInterfaceAssignment{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkInterfaceAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterfaceAssignment) ProtoMessage() {}

func (x *NetworkInterfaceAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterfaceAssignment.ProtoReflect.Descriptor instead.
func (*NetworkInterfaceAssignment) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{7}
}

func (x *NetworkInterfaceAssignment) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *NetworkInterfaceAssignment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LoadBalancerAssignment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Load Balancer.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoadBalancerAssignment) Reset() {
	*x = LoadBalancerAssignment{}
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoadBalancerAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerAssignment) ProtoMessage() {}

func (x *LoadBalancerAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1_allocation_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerAssignment.ProtoReflect.Descriptor instead.
func (*LoadBalancerAssignment) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1_allocation_proto_rawDescGZIP(), []int{8}
}

func (x *LoadBalancerAssignment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_nebius_vpc_v1_allocation_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1_allocation_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x31, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc9, 0x01, 0x0a,
	0x0e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x53, 0x0a, 0x0c, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x42,
	0x04, 0xba, 0x4a, 0x01, 0x06, 0x48, 0x00, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x34, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x06, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x70, 0x76, 0x34,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x10, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0xc3, 0x02, 0x0a, 0x19, 0x49, 0x50, 0x76,
	0x34, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0xcc, 0x01, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb7, 0x01, 0xba, 0x48, 0xaf, 0x01, 0xba, 0x01, 0xab, 0x01,
	0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63,
	0x69, 0x64, 0x72, 0x12, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x49, 0x50, 0x20, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x2c, 0x20, 0x43, 0x49, 0x44, 0x52, 0x20, 0x6f, 0x72, 0x20, 0x6d,
	0x61, 0x73, 0x6b, 0x1a, 0x66, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20,
	0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28,
	0x27, 0x5e, 0x2f, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x5b, 0x31, 0x2d, 0x32, 0x5d, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x33, 0x5b, 0x30, 0x2d, 0x32, 0x5d, 0x29, 0x24, 0x27, 0x29, 0x20,
	0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49, 0x70, 0x28, 0x34, 0x29, 0x20,
	0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49, 0x70, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x28, 0x34, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x29, 0xba, 0x4a, 0x01, 0x02, 0x52,
	0x04, 0x63, 0x69, 0x64, 0x72, 0x12, 0x23, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x48, 0x00,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x07, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01,
	0x02, 0x48, 0x00, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x42, 0x11, 0x0a, 0x04, 0x70,
	0x6f, 0x6f, 0x6c, 0x12, 0x09, 0xba, 0x48, 0x02, 0x08, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x22, 0xc2,
	0x02, 0x0a, 0x18, 0x49, 0x50, 0x76, 0x34, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0xcc, 0x01, 0x0a, 0x04,
	0x63, 0x69, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xb7, 0x01, 0xba, 0x48, 0xaf,
	0x01, 0xba, 0x01, 0xab, 0x01, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x12, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20,
	0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20,
	0x49, 0x50, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2c, 0x20, 0x43, 0x49, 0x44, 0x52,
	0x20, 0x6f, 0x72, 0x20, 0x6d, 0x61, 0x73, 0x6b, 0x1a, 0x66, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3d,
	0x3d, 0x20, 0x27, 0x27, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x2f, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x5b,
	0x31, 0x2d, 0x32, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x33, 0x5b, 0x30, 0x2d, 0x32, 0x5d,
	0x29, 0x24, 0x27, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49,
	0x70, 0x28, 0x34, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49,
	0x70, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x28, 0x34, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x29,
	0xba, 0x4a, 0x01, 0x02, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x12, 0x23, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba,
	0x4a, 0x01, 0x02, 0x48, 0x00, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x48, 0x00, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64,
	0x42, 0x11, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x12, 0x09, 0xba, 0x48, 0x02, 0x08, 0x01, 0xba,
	0x4a, 0x01, 0x02, 0x22, 0xb7, 0x02, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x63, 0x22, 0x57, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x22, 0x87, 0x01,
	0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x43, 0x69, 0x64, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6f,
	0x6c, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x01, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x4c, 0x0a, 0x0d, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x0c, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x42, 0x06,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x4c, 0x6f, 0x61,
	0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x56, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x41, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_nebius_vpc_v1_allocation_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1_allocation_proto_rawDescData []byte
)

func file_nebius_vpc_v1_allocation_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1_allocation_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1_allocation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1_allocation_proto_rawDesc), len(file_nebius_vpc_v1_allocation_proto_rawDesc)))
	})
	return file_nebius_vpc_v1_allocation_proto_rawDescData
}

var file_nebius_vpc_v1_allocation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_vpc_v1_allocation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_nebius_vpc_v1_allocation_proto_goTypes = []any{
	(AllocationStatus_State)(0),        // 0: nebius.vpc.v1.AllocationStatus.State
	(*Allocation)(nil),                 // 1: nebius.vpc.v1.Allocation
	(*AllocationSpec)(nil),             // 2: nebius.vpc.v1.AllocationSpec
	(*IPv4PrivateAllocationSpec)(nil),  // 3: nebius.vpc.v1.IPv4PrivateAllocationSpec
	(*IPv4PublicAllocationSpec)(nil),   // 4: nebius.vpc.v1.IPv4PublicAllocationSpec
	(*AllocationStatus)(nil),           // 5: nebius.vpc.v1.AllocationStatus
	(*AllocationDetails)(nil),          // 6: nebius.vpc.v1.AllocationDetails
	(*Assignment)(nil),                 // 7: nebius.vpc.v1.Assignment
	(*NetworkInterfaceAssignment)(nil), // 8: nebius.vpc.v1.NetworkInterfaceAssignment
	(*LoadBalancerAssignment)(nil),     // 9: nebius.vpc.v1.LoadBalancerAssignment
	(*v1.ResourceMetadata)(nil),        // 10: nebius.common.v1.ResourceMetadata
	(IpVersion)(0),                     // 11: nebius.vpc.v1.IpVersion
}
var file_nebius_vpc_v1_allocation_proto_depIdxs = []int32{
	10, // 0: nebius.vpc.v1.Allocation.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2,  // 1: nebius.vpc.v1.Allocation.spec:type_name -> nebius.vpc.v1.AllocationSpec
	5,  // 2: nebius.vpc.v1.Allocation.status:type_name -> nebius.vpc.v1.AllocationStatus
	3,  // 3: nebius.vpc.v1.AllocationSpec.ipv4_private:type_name -> nebius.vpc.v1.IPv4PrivateAllocationSpec
	4,  // 4: nebius.vpc.v1.AllocationSpec.ipv4_public:type_name -> nebius.vpc.v1.IPv4PublicAllocationSpec
	0,  // 5: nebius.vpc.v1.AllocationStatus.state:type_name -> nebius.vpc.v1.AllocationStatus.State
	6,  // 6: nebius.vpc.v1.AllocationStatus.details:type_name -> nebius.vpc.v1.AllocationDetails
	7,  // 7: nebius.vpc.v1.AllocationStatus.assignment:type_name -> nebius.vpc.v1.Assignment
	11, // 8: nebius.vpc.v1.AllocationDetails.version:type_name -> nebius.vpc.v1.IpVersion
	8,  // 9: nebius.vpc.v1.Assignment.network_interface:type_name -> nebius.vpc.v1.NetworkInterfaceAssignment
	9,  // 10: nebius.vpc.v1.Assignment.load_balancer:type_name -> nebius.vpc.v1.LoadBalancerAssignment
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1_allocation_proto_init() }
func file_nebius_vpc_v1_allocation_proto_init() {
	if File_nebius_vpc_v1_allocation_proto != nil {
		return
	}
	file_nebius_vpc_v1_pool_proto_init()
	file_nebius_vpc_v1_allocation_proto_msgTypes[1].OneofWrappers = []any{
		(*AllocationSpec_Ipv4Private)(nil),
		(*AllocationSpec_Ipv4Public)(nil),
	}
	file_nebius_vpc_v1_allocation_proto_msgTypes[2].OneofWrappers = []any{
		(*IPv4PrivateAllocationSpec_SubnetId)(nil),
		(*IPv4PrivateAllocationSpec_PoolId)(nil),
	}
	file_nebius_vpc_v1_allocation_proto_msgTypes[3].OneofWrappers = []any{
		(*IPv4PublicAllocationSpec_SubnetId)(nil),
		(*IPv4PublicAllocationSpec_PoolId)(nil),
	}
	file_nebius_vpc_v1_allocation_proto_msgTypes[6].OneofWrappers = []any{
		(*Assignment_NetworkInterface)(nil),
		(*Assignment_LoadBalancer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1_allocation_proto_rawDesc), len(file_nebius_vpc_v1_allocation_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_vpc_v1_allocation_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1_allocation_proto_depIdxs,
		EnumInfos:         file_nebius_vpc_v1_allocation_proto_enumTypes,
		MessageInfos:      file_nebius_vpc_v1_allocation_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1_allocation_proto = out.File
	file_nebius_vpc_v1_allocation_proto_goTypes = nil
	file_nebius_vpc_v1_allocation_proto_depIdxs = nil
}
