// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/vpc/v1alpha1/scope.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Type of scope.
type ScopeSpec_Type int32

const (
	ScopeSpec_SCOPE_TYPE_UNSPECIFIED ScopeSpec_Type = 0 // Default, unspecified scope type.
	ScopeSpec_PUBLIC                 ScopeSpec_Type = 1 // Public scope.
	ScopeSpec_PRIVATE                ScopeSpec_Type = 2 // Private scope.
)

// Enum value maps for ScopeSpec_Type.
var (
	ScopeSpec_Type_name = map[int32]string{
		0: "SCOPE_TYPE_UNSPECIFIED",
		1: "PUBLIC",
		2: "PRIVATE",
	}
	ScopeSpec_Type_value = map[string]int32{
		"SCOPE_TYPE_UNSPECIFIED": 0,
		"PUBLIC":                 1,
		"PRIVATE":                2,
	}
)

func (x ScopeSpec_Type) Enum() *ScopeSpec_Type {
	p := new(ScopeSpec_Type)
	*p = x
	return p
}

func (x ScopeSpec_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScopeSpec_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1alpha1_scope_proto_enumTypes[0].Descriptor()
}

func (ScopeSpec_Type) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1alpha1_scope_proto_enumTypes[0]
}

func (x ScopeSpec_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScopeSpec_Type.Descriptor instead.
func (ScopeSpec_Type) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_scope_proto_rawDescGZIP(), []int{1, 0}
}

// Possible states of the Scope.
type ScopeStatus_State int32

const (
	ScopeStatus_STATE_UNSPECIFIED ScopeStatus_State = 0 // Default, unspecified state.
	ScopeStatus_CREATING          ScopeStatus_State = 1 // Scope is being created.
	ScopeStatus_READY             ScopeStatus_State = 2 // Scope is ready for use.
	ScopeStatus_DELETING          ScopeStatus_State = 3 // Scope is being deleted.
)

// Enum value maps for ScopeStatus_State.
var (
	ScopeStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "DELETING",
	}
	ScopeStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"READY":             2,
		"DELETING":          3,
	}
)

func (x ScopeStatus_State) Enum() *ScopeStatus_State {
	p := new(ScopeStatus_State)
	*p = x
	return p
}

func (x ScopeStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScopeStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_vpc_v1alpha1_scope_proto_enumTypes[1].Descriptor()
}

func (ScopeStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_vpc_v1alpha1_scope_proto_enumTypes[1]
}

func (x ScopeStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScopeStatus_State.Descriptor instead.
func (ScopeStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_scope_proto_rawDescGZIP(), []int{2, 0}
}

type Scope struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Metadata associated with the Scope.
	// `metadata.parent_id` represents the parent IAM container.
	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specification of the Scope.
	Spec *ScopeSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Status information for the Scope.
	Status        *ScopeStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Scope) Reset() {
	*x = Scope{}
	mi := &file_nebius_vpc_v1alpha1_scope_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_scope_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_scope_proto_rawDescGZIP(), []int{0}
}

func (x *Scope) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Scope) GetSpec() *ScopeSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Scope) GetStatus() *ScopeStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ScopeSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Type of the Scope (Private or Public).
	Type          ScopeSpec_Type `protobuf:"varint,1,opt,name=type,proto3,enum=nebius.vpc.v1alpha1.ScopeSpec_Type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScopeSpec) Reset() {
	*x = ScopeSpec{}
	mi := &file_nebius_vpc_v1alpha1_scope_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScopeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeSpec) ProtoMessage() {}

func (x *ScopeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_scope_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeSpec.ProtoReflect.Descriptor instead.
func (*ScopeSpec) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_scope_proto_rawDescGZIP(), []int{1}
}

func (x *ScopeSpec) GetType() ScopeSpec_Type {
	if x != nil {
		return x.Type
	}
	return ScopeSpec_SCOPE_TYPE_UNSPECIFIED
}

type ScopeStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current state of the Scope.
	State         ScopeStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.vpc.v1alpha1.ScopeStatus_State" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScopeStatus) Reset() {
	*x = ScopeStatus{}
	mi := &file_nebius_vpc_v1alpha1_scope_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScopeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopeStatus) ProtoMessage() {}

func (x *ScopeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_vpc_v1alpha1_scope_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopeStatus.ProtoReflect.Descriptor instead.
func (*ScopeStatus) Descriptor() ([]byte, []int) {
	return file_nebius_vpc_v1alpha1_scope_proto_rawDescGZIP(), []int{2}
}

func (x *ScopeStatus) GetState() ScopeStatus_State {
	if x != nil {
		return x.State
	}
	return ScopeStatus_STATE_UNSPECIFIED
}

var File_nebius_vpc_v1alpha1_scope_proto protoreflect.FileDescriptor

var file_nebius_vpc_v1alpha1_scope_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3e,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x89, 0x01, 0x0a,
	0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3b, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x02, 0x22, 0x92, 0x01, 0x0a, 0x0b, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x42, 0x5d, 0x0a,
	0x1a, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_vpc_v1alpha1_scope_proto_rawDescOnce sync.Once
	file_nebius_vpc_v1alpha1_scope_proto_rawDescData []byte
)

func file_nebius_vpc_v1alpha1_scope_proto_rawDescGZIP() []byte {
	file_nebius_vpc_v1alpha1_scope_proto_rawDescOnce.Do(func() {
		file_nebius_vpc_v1alpha1_scope_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1alpha1_scope_proto_rawDesc), len(file_nebius_vpc_v1alpha1_scope_proto_rawDesc)))
	})
	return file_nebius_vpc_v1alpha1_scope_proto_rawDescData
}

var file_nebius_vpc_v1alpha1_scope_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nebius_vpc_v1alpha1_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_vpc_v1alpha1_scope_proto_goTypes = []any{
	(ScopeSpec_Type)(0),         // 0: nebius.vpc.v1alpha1.ScopeSpec.Type
	(ScopeStatus_State)(0),      // 1: nebius.vpc.v1alpha1.ScopeStatus.State
	(*Scope)(nil),               // 2: nebius.vpc.v1alpha1.Scope
	(*ScopeSpec)(nil),           // 3: nebius.vpc.v1alpha1.ScopeSpec
	(*ScopeStatus)(nil),         // 4: nebius.vpc.v1alpha1.ScopeStatus
	(*v1.ResourceMetadata)(nil), // 5: nebius.common.v1.ResourceMetadata
}
var file_nebius_vpc_v1alpha1_scope_proto_depIdxs = []int32{
	5, // 0: nebius.vpc.v1alpha1.Scope.metadata:type_name -> nebius.common.v1.ResourceMetadata
	3, // 1: nebius.vpc.v1alpha1.Scope.spec:type_name -> nebius.vpc.v1alpha1.ScopeSpec
	4, // 2: nebius.vpc.v1alpha1.Scope.status:type_name -> nebius.vpc.v1alpha1.ScopeStatus
	0, // 3: nebius.vpc.v1alpha1.ScopeSpec.type:type_name -> nebius.vpc.v1alpha1.ScopeSpec.Type
	1, // 4: nebius.vpc.v1alpha1.ScopeStatus.state:type_name -> nebius.vpc.v1alpha1.ScopeStatus.State
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_vpc_v1alpha1_scope_proto_init() }
func file_nebius_vpc_v1alpha1_scope_proto_init() {
	if File_nebius_vpc_v1alpha1_scope_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_vpc_v1alpha1_scope_proto_rawDesc), len(file_nebius_vpc_v1alpha1_scope_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_vpc_v1alpha1_scope_proto_goTypes,
		DependencyIndexes: file_nebius_vpc_v1alpha1_scope_proto_depIdxs,
		EnumInfos:         file_nebius_vpc_v1alpha1_scope_proto_enumTypes,
		MessageInfos:      file_nebius_vpc_v1alpha1_scope_proto_msgTypes,
	}.Build()
	File_nebius_vpc_v1alpha1_scope_proto = out.File
	file_nebius_vpc_v1alpha1_scope_proto_goTypes = nil
	file_nebius_vpc_v1alpha1_scope_proto_depIdxs = nil
}
