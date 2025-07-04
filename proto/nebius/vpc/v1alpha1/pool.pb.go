// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/vpc/v1alpha1/pool.proto

package v1alpha1

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

type PoolCidrState int32

const (
	PoolCidrState_STATE_UNSPECIFIED PoolCidrState = 0 // Default, unspecified state.
	PoolCidrState_AVAILABLE         PoolCidrState = 1 // Allocation from range is available.
	PoolCidrState_DISABLED          PoolCidrState = 2 // New allocation would not be created.
)

// Enum value maps for PoolCidrState.
var (
	PoolCidrState_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "AVAILABLE",
		2: "DISABLED",
	}
	PoolCidrState_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"AVAILABLE":         1,
		"DISABLED":          2,
	}
)

func (x PoolCidrState) Enum() *PoolCidrState {
	p := new(PoolCidrState)
	*p = x
	return p
}

func (x PoolCidrState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PoolCidrState) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1alpha1_pool_proto_enumTypes[0].Descriptor()
}

func (PoolCidrState) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1alpha1_pool_proto_enumTypes[0]
}

func (x PoolCidrState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PoolCidrState.Descriptor instead.
func (PoolCidrState) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{0}
}

type IpVersion int32

const (
	IpVersion_IP_VERSION_UNSPECIFIED IpVersion = 0 // Default, unspecified IP version.
	IpVersion_IPV4                   IpVersion = 1 // IPv4 address.
	IpVersion_IPV6                   IpVersion = 2 // IPv6 address.
)

// Enum value maps for IpVersion.
var (
	IpVersion_name = map[int32]string{
		0: "IP_VERSION_UNSPECIFIED",
		1: "IPV4",
		2: "IPV6",
	}
	IpVersion_value = map[string]int32{
		"IP_VERSION_UNSPECIFIED": 0,
		"IPV4":                   1,
		"IPV6":                   2,
	}
)

func (x IpVersion) Enum() *IpVersion {
	p := new(IpVersion)
	*p = x
	return p
}

func (x IpVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IpVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1alpha1_pool_proto_enumTypes[1].Descriptor()
}

func (IpVersion) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1alpha1_pool_proto_enumTypes[1]
}

func (x IpVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IpVersion.Descriptor instead.
func (IpVersion) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{1}
}

// Possible states of the Pool.
type PoolStatus_State int32

const (
	PoolStatus_STATE_UNSPECIFIED PoolStatus_State = 0 // Default, unspecified state.
	PoolStatus_CREATING          PoolStatus_State = 1 // Pool is being created.
	PoolStatus_READY             PoolStatus_State = 2 // Pool is ready for use.
	PoolStatus_DELETING          PoolStatus_State = 3 // Pool is being deleted.
)

// Enum value maps for PoolStatus_State.
var (
	PoolStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
	}
	PoolStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
	}
)

func (x PoolStatus_State) Enum() *PoolStatus_State {
	p := new(PoolStatus_State)
	*p = x
	return p
}

func (x PoolStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PoolStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1alpha1_pool_proto_enumTypes[2].Descriptor()
}

func (PoolStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1alpha1_pool_proto_enumTypes[2]
}

func (x PoolStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PoolStatus_State.Descriptor instead.
func (PoolStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{3, 0}
}

type Pool struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata associated with the Pool.
	// `metadata.parent_id` represents the Project.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification of the Pool.
	Spec *PoolSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status information for the Pool.
	Status        *PoolStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pool) Reset() {
	*x = Pool{}
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

func (x *Pool) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{0}
}

func (x *Pool) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Pool) GetSpec() *PoolSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Pool) GetStatus() *PoolStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type PoolSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Source:
	//
	//	*PoolSpec_SourcePoolId
	//	*PoolSpec_SourceScopeId
	Source isPoolSpec_Source `protobuf_oneof:"source"`
	// IP version for the Pool.
	Version IpVersion `protobuf:"varint,3,opt,name=version,proto3,enum=nebius.vpc.v1alpha1.IpVersion" json:"version,omitempty"`
	// CIDR blocks.
	Cidrs         []*PoolCidr `protobuf:"bytes,4,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoolSpec) Reset() {
	*x = PoolSpec{}
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoolSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolSpec) ProtoMessage() {}

func (x *PoolSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolSpec.ProtoReflect.Descriptor instead.
func (*PoolSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{1}
}

func (x *PoolSpec) GetSource() isPoolSpec_Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *PoolSpec) GetSourcePoolId() string {
	if x != nil {
		if x, ok := x.Source.(*PoolSpec_SourcePoolId); ok {
			return x.SourcePoolId
		}
	}
	return ""
}

func (x *PoolSpec) GetSourceScopeId() string {
	if x != nil {
		if x, ok := x.Source.(*PoolSpec_SourceScopeId); ok {
			return x.SourceScopeId
		}
	}
	return ""
}

func (x *PoolSpec) GetVersion() IpVersion {
	if x != nil {
		return x.Version
	}
	return IpVersion_IP_VERSION_UNSPECIFIED
}

func (x *PoolSpec) GetCidrs() []*PoolCidr {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

type isPoolSpec_Source interface {
	isPoolSpec_Source()
}

type PoolSpec_SourcePoolId struct {
	// ID of source pool. Current pool will be created with the same scope.
	SourcePoolId string `protobuf:"bytes,1,opt,name=source_pool_id,json=sourcePoolId,proto3,oneof"`
}

type PoolSpec_SourceScopeId struct {
	// ID of the scope. Pool will be considered as top-level pool within scope.
	SourceScopeId string `protobuf:"bytes,2,opt,name=source_scope_id,json=sourceScopeId,proto3,oneof"`
}

func (*PoolSpec_SourcePoolId) isPoolSpec_Source() {}

func (*PoolSpec_SourceScopeId) isPoolSpec_Source() {}

type PoolCidr struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CIDR block.
	// May be a prefix length (such as /24) for non-top-level pools
	// or a CIDR-formatted string (such as 10.1.2.0/24).
	Cidr string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	// State of the Cidr.
	State PoolCidrState `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.vpc.v1alpha1.PoolCidrState" json:"state,omitempty"`
	// Maximum mask length for allocation from this IP pool including creation of sub-pools
	AllowedMask   int64 `protobuf:"varint,3,opt,name=allowed_mask,json=allowedMask,proto3" json:"allowed_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoolCidr) Reset() {
	*x = PoolCidr{}
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoolCidr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolCidr) ProtoMessage() {}

func (x *PoolCidr) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolCidr.ProtoReflect.Descriptor instead.
func (*PoolCidr) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{2}
}

func (x *PoolCidr) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *PoolCidr) GetState() PoolCidrState {
	if x != nil {
		return x.State
	}
	return PoolCidrState_STATE_UNSPECIFIED
}

func (x *PoolCidr) GetAllowedMask() int64 {
	if x != nil {
		return x.AllowedMask
	}
	return 0
}

type PoolStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current state of the Pool.
	State PoolStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.vpc.v1alpha1.PoolStatus_State" json:"state,omitempty"`
	// CIDR blocks.
	Cidrs []string `protobuf:"bytes,2,rep,name=cidrs,proto3" json:"cidrs,omitempty"`
	// ID of the scope
	ScopeId       string `protobuf:"bytes,3,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PoolStatus) Reset() {
	*x = PoolStatus{}
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoolStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolStatus) ProtoMessage() {}

func (x *PoolStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_pool_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoolStatus.ProtoReflect.Descriptor instead.
func (*PoolStatus) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP(), []int{3}
}

func (x *PoolStatus) GetState() PoolStatus_State {
	if x != nil {
		return x.State
	}
	return PoolStatus_STATE_UNSPECIFIED
}

func (x *PoolStatus) GetCidrs() []string {
	if x != nil {
		return x.Cidrs
	}
	return nil
}

func (x *PoolStatus) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

var File_nebius_vpc_v1alpha1_pool_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1alpha1_pool_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01,
	0x0a, 0x04, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70,
	0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x2c, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x48, 0x00, 0x52,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x48, 0x00, 0x52, 0x0d,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0a,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x69, 0x64,
	0x72, 0x52, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x42, 0x0f, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x8f, 0x03, 0x0a, 0x08, 0x50, 0x6f,
	0x6f, 0x6c, 0x43, 0x69, 0x64, 0x72, 0x12, 0x8e, 0x02, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xf9, 0x01, 0xba, 0x48, 0xf5, 0x01, 0xba, 0x01, 0x9d, 0x01,
	0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63,
	0x69, 0x64, 0x72, 0x12, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x49, 0x50, 0x20, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x2c, 0x20, 0x43, 0x49, 0x44, 0x52, 0x20, 0x6f, 0x72, 0x20, 0x6d,
	0x61, 0x73, 0x6b, 0x1a, 0x58, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20,
	0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28,
	0x27, 0x5e, 0x2f, 0x28, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x5b, 0x31, 0x2d, 0x39, 0x5d, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x7c, 0x31, 0x5b, 0x30, 0x2d, 0x32, 0x5d, 0x5b, 0x30, 0x2d, 0x38, 0x5d,
	0x29, 0x24, 0x27, 0x29, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x49,
	0x70, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x28, 0x74, 0x72, 0x75, 0x65, 0x29, 0xba, 0x01, 0x4e,
	0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x70, 0x5f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x69, 0x73, 0x20, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2c, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x49, 0x50, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x1a, 0x0a, 0x74, 0x68, 0x69, 0x73, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0xc8, 0x01,
	0x01, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6f,
	0x6c, 0x43, 0x69, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x22, 0x05, 0x18, 0x80, 0x01, 0x28, 0x00, 0x52, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0xc1, 0x01, 0x0a, 0x0a,
	0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x69, 0x64, 0x72, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x2a,
	0x43, 0x0a, 0x0d, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x69, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x2a, 0x3b, 0x0a, 0x09, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x50, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x36, 0x10,
	0x02, 0x42, 0x5c, 0x0a, 0x1a, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x09, 0x50, 0x6f, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_vpc_v1alpha1_pool_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1alpha1_pool_proto_rawDescData []byte
)

func file_nebius_vpc_v1alpha1_pool_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1alpha1_pool_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1alpha1_pool_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1alpha1_pool_proto_rawDesc), len(file_nebius_vpc_v1alpha1_pool_proto_rawDesc)))
	})
	return file_nebius_vpc_v1alpha1_pool_proto_rawDescData
}

var file_nebius_vpc_v1alpha1_pool_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_nebius_vpc_v1alpha1_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_vpc_v1alpha1_pool_proto_goTypes = []any{
	(PoolCidrState)(0),          // 0: nebius.vpc.v1alpha1.PoolCidrState
	(IpVersion)(0),              // 1: nebius.vpc.v1alpha1.IpVersion
	(PoolStatus_State)(0),       // 2: nebius.vpc.v1alpha1.PoolStatus.State
	(*Pool)(nil),                // 3: nebius.vpc.v1alpha1.Pool
	(*PoolSpec)(nil),            // 4: nebius.vpc.v1alpha1.PoolSpec
	(*PoolCidr)(nil),            // 5: nebius.vpc.v1alpha1.PoolCidr
	(*PoolStatus)(nil),          // 6: nebius.vpc.v1alpha1.PoolStatus
	(*v1.ResourceMetadata)(nil), // 7: nebius.common.v1.ResourceMetadata
}
var file_nebius_vpc_v1alpha1_pool_proto_depIdxs = []int32{
	7, // 0: nebius.vpc.v1alpha1.Pool.metadata:type_name -> nebius.common.v1.ResourceMetadata
	4, // 1: nebius.vpc.v1alpha1.Pool.spec:type_name -> nebius.vpc.v1alpha1.PoolSpec
	6, // 2: nebius.vpc.v1alpha1.Pool.status:type_name -> nebius.vpc.v1alpha1.PoolStatus
	1, // 3: nebius.vpc.v1alpha1.PoolSpec.version:type_name -> nebius.vpc.v1alpha1.IpVersion
	5, // 4: nebius.vpc.v1alpha1.PoolSpec.cidrs:type_name -> nebius.vpc.v1alpha1.PoolCidr
	0, // 5: nebius.vpc.v1alpha1.PoolCidr.state:type_name -> nebius.vpc.v1alpha1.PoolCidrState
	2, // 6: nebius.vpc.v1alpha1.PoolStatus.state:type_name -> nebius.vpc.v1alpha1.PoolStatus.State
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1alpha1_pool_proto_init() }
func file_nebius_vpc_v1alpha1_pool_proto_init() {
	if File_nebius_vpc_v1alpha1_pool_proto != nil {
		return
	}
	file_nebius_vpc_v1alpha1_pool_proto_msgTypes[1].OneofWrappers = []any{
		(*PoolSpec_SourcePoolId)(nil),
		(*PoolSpec_SourceScopeId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1alpha1_pool_proto_rawDesc), len(file_nebius_vpc_v1alpha1_pool_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_vpc_v1alpha1_pool_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1alpha1_pool_proto_depIdxs,
		EnumInfos:         file_nebius_vpc_v1alpha1_pool_proto_enumTypes,
		MessageInfos:      file_nebius_vpc_v1alpha1_pool_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1alpha1_pool_proto = out.File
	file_nebius_vpc_v1alpha1_pool_proto_goTypes = nil
	file_nebius_vpc_v1alpha1_pool_proto_depIdxs = nil
}
