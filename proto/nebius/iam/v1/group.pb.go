// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/group.proto

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

type GroupStatus_State int32

const (
	GroupStatus_UNSPECIFIED GroupStatus_State = 0
	GroupStatus_ACTIVE      GroupStatus_State = 1
)

// Enum value maps for GroupStatus_State.
var (
	GroupStatus_State_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ACTIVE",
	}
	GroupStatus_State_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ACTIVE":      1,
	}
)

func (x GroupStatus_State) Enum() *GroupStatus_State {
	p := new(GroupStatus_State)
	*p = x
	return p
}

func (x GroupStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_group_proto_enumTypes[0].Descriptor()
}

func (GroupStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_group_proto_enumTypes[0]
}

func (x GroupStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupStatus_State.Descriptor instead.
func (GroupStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_proto_rawDescGZIP(), []int{2, 0}
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *GroupSpec           `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *GroupStatus         `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Group) GetSpec() *GroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Group) GetStatus() *GroupStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type GroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupSpec) Reset() {
	*x = GroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSpec) ProtoMessage() {}

func (x *GroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSpec.ProtoReflect.Descriptor instead.
func (*GroupSpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_proto_rawDescGZIP(), []int{1}
}

type GroupStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State        GroupStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.iam.v1.GroupStatus_State" json:"state,omitempty"`
	MembersCount int32             `protobuf:"varint,2,opt,name=members_count,json=membersCount,proto3" json:"members_count,omitempty"`
}

func (x *GroupStatus) Reset() {
	*x = GroupStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupStatus) ProtoMessage() {}

func (x *GroupStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupStatus.ProtoReflect.Descriptor instead.
func (*GroupStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *GroupStatus) GetState() GroupStatus_State {
	if x != nil {
		return x.State
	}
	return GroupStatus_UNSPECIFIED
}

func (x *GroupStatus) GetMembersCount() int32 {
	if x != nil {
		return x.MembersCount
	}
	return 0
}

var File_nebius_iam_v1_group_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_group_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x46, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65,
	0x63, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x24,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x42, 0x51, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_group_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_group_proto_rawDescData = file_nebius_iam_v1_group_proto_rawDesc
)

func file_nebius_iam_v1_group_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_group_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_group_proto_rawDescData)
	})
	return file_nebius_iam_v1_group_proto_rawDescData
}

var file_nebius_iam_v1_group_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_iam_v1_group_proto_goTypes = []any{
	(GroupStatus_State)(0),      // 0: nebius.iam.v1.GroupStatus.State
	(*Group)(nil),               // 1: nebius.iam.v1.Group
	(*GroupSpec)(nil),           // 2: nebius.iam.v1.GroupSpec
	(*GroupStatus)(nil),         // 3: nebius.iam.v1.GroupStatus
	(*v1.ResourceMetadata)(nil), // 4: nebius.common.v1.ResourceMetadata
}
var file_nebius_iam_v1_group_proto_depIdxs = []int32{
	4, // 0: nebius.iam.v1.Group.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2, // 1: nebius.iam.v1.Group.spec:type_name -> nebius.iam.v1.GroupSpec
	3, // 2: nebius.iam.v1.Group.status:type_name -> nebius.iam.v1.GroupStatus
	0, // 3: nebius.iam.v1.GroupStatus.state:type_name -> nebius.iam.v1.GroupStatus.State
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_group_proto_init() }
func file_nebius_iam_v1_group_proto_init() {
	if File_nebius_iam_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nebius_iam_v1_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Group); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_iam_v1_group_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GroupSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_iam_v1_group_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GroupStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_group_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_group_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_group_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_group_proto_enumTypes,
		MessageInfos:      file_nebius_iam_v1_group_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_group_proto = out.File
	file_nebius_iam_v1_group_proto_rawDesc = nil
	file_nebius_iam_v1_group_proto_goTypes = nil
	file_nebius_iam_v1_group_proto_depIdxs = nil
}
