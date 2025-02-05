// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/group_membership.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GroupMemberKind_Kind int32

const (
	GroupMemberKind_KIND_UNSPECIFIED             GroupMemberKind_Kind = 0
	GroupMemberKind_ORDINARY_TENANT_USER_ACCOUNT GroupMemberKind_Kind = 1
	GroupMemberKind_INVITED_TENANT_USER_ACCOUNT  GroupMemberKind_Kind = 2
	GroupMemberKind_SERVICE_ACCOUNT              GroupMemberKind_Kind = 3
)

// Enum value maps for GroupMemberKind_Kind.
var (
	GroupMemberKind_Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "ORDINARY_TENANT_USER_ACCOUNT",
		2: "INVITED_TENANT_USER_ACCOUNT",
		3: "SERVICE_ACCOUNT",
	}
	GroupMemberKind_Kind_value = map[string]int32{
		"KIND_UNSPECIFIED":             0,
		"ORDINARY_TENANT_USER_ACCOUNT": 1,
		"INVITED_TENANT_USER_ACCOUNT":  2,
		"SERVICE_ACCOUNT":              3,
	}
)

func (x GroupMemberKind_Kind) Enum() *GroupMemberKind_Kind {
	p := new(GroupMemberKind_Kind)
	*p = x
	return p
}

func (x GroupMemberKind_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupMemberKind_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_group_membership_proto_enumTypes[0].Descriptor()
}

func (GroupMemberKind_Kind) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_group_membership_proto_enumTypes[0]
}

func (x GroupMemberKind_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupMemberKind_Kind.Descriptor instead.
func (GroupMemberKind_Kind) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_proto_rawDescGZIP(), []int{3, 0}
}

type GroupMembership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *GroupMembershipSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *GroupMembershipStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"` // Dummy field for compliance with terraform resource generator.
	RevokeAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=revoke_at,json=revokeAt,proto3" json:"revoke_at,omitempty"`
}

func (x *GroupMembership) Reset() {
	*x = GroupMembership{}
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMembership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembership) ProtoMessage() {}

func (x *GroupMembership) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembership.ProtoReflect.Descriptor instead.
func (*GroupMembership) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMembership) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GroupMembership) GetSpec() *GroupMembershipSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GroupMembership) GetStatus() *GroupMembershipStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GroupMembership) GetRevokeAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RevokeAt
	}
	return nil
}

type GroupMembershipSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Member of the group. Can be tenant user account id or service account id.
	MemberId string `protobuf:"bytes,1,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
}

func (x *GroupMembershipSpec) Reset() {
	*x = GroupMembershipSpec{}
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMembershipSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembershipSpec) ProtoMessage() {}

func (x *GroupMembershipSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembershipSpec.ProtoReflect.Descriptor instead.
func (*GroupMembershipSpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_proto_rawDescGZIP(), []int{1}
}

func (x *GroupMembershipSpec) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type GroupMembershipStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupMembershipStatus) Reset() {
	*x = GroupMembershipStatus{}
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMembershipStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembershipStatus) ProtoMessage() {}

func (x *GroupMembershipStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembershipStatus.ProtoReflect.Descriptor instead.
func (*GroupMembershipStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_proto_rawDescGZIP(), []int{2}
}

type GroupMemberKind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind GroupMemberKind_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=nebius.iam.v1.GroupMemberKind_Kind" json:"kind,omitempty"`
}

func (x *GroupMemberKind) Reset() {
	*x = GroupMemberKind{}
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMemberKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMemberKind) ProtoMessage() {}

func (x *GroupMemberKind) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMemberKind.ProtoReflect.Descriptor instead.
func (*GroupMemberKind) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_proto_rawDescGZIP(), []int{3}
}

func (x *GroupMemberKind) GetKind() GroupMemberKind_Kind {
	if x != nil {
		return x.Kind
	}
	return GroupMemberKind_KIND_UNSPECIFIED
}

// see also nebius/iam/v1/tenant_user_account.proto/TenantUserAccountWithAttributes
type GroupMembershipWithAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupMembership *GroupMembership `protobuf:"bytes,1,opt,name=group_membership,json=groupMembership,proto3" json:"group_membership,omitempty"`
	GroupMemberKind *GroupMemberKind `protobuf:"bytes,11,opt,name=group_member_kind,json=groupMemberKind,proto3" json:"group_member_kind,omitempty"`
	// Types that are assignable to AttributesOptional:
	//
	//	*GroupMembershipWithAttributes_UserAttributes
	//	*GroupMembershipWithAttributes_ServiceAccountAttributes
	//	*GroupMembershipWithAttributes_Error
	AttributesOptional isGroupMembershipWithAttributes_AttributesOptional `protobuf_oneof:"attributesOptional"`
}

func (x *GroupMembershipWithAttributes) Reset() {
	*x = GroupMembershipWithAttributes{}
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMembershipWithAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMembershipWithAttributes) ProtoMessage() {}

func (x *GroupMembershipWithAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMembershipWithAttributes.ProtoReflect.Descriptor instead.
func (*GroupMembershipWithAttributes) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_proto_rawDescGZIP(), []int{4}
}

func (x *GroupMembershipWithAttributes) GetGroupMembership() *GroupMembership {
	if x != nil {
		return x.GroupMembership
	}
	return nil
}

func (x *GroupMembershipWithAttributes) GetGroupMemberKind() *GroupMemberKind {
	if x != nil {
		return x.GroupMemberKind
	}
	return nil
}

func (m *GroupMembershipWithAttributes) GetAttributesOptional() isGroupMembershipWithAttributes_AttributesOptional {
	if m != nil {
		return m.AttributesOptional
	}
	return nil
}

func (x *GroupMembershipWithAttributes) GetUserAttributes() *UserAttributes {
	if x, ok := x.GetAttributesOptional().(*GroupMembershipWithAttributes_UserAttributes); ok {
		return x.UserAttributes
	}
	return nil
}

func (x *GroupMembershipWithAttributes) GetServiceAccountAttributes() *ServiceAccountAttributes {
	if x, ok := x.GetAttributesOptional().(*GroupMembershipWithAttributes_ServiceAccountAttributes); ok {
		return x.ServiceAccountAttributes
	}
	return nil
}

func (x *GroupMembershipWithAttributes) GetError() *Error {
	if x, ok := x.GetAttributesOptional().(*GroupMembershipWithAttributes_Error); ok {
		return x.Error
	}
	return nil
}

type isGroupMembershipWithAttributes_AttributesOptional interface {
	isGroupMembershipWithAttributes_AttributesOptional()
}

type GroupMembershipWithAttributes_UserAttributes struct {
	// filled with known data for members corresponding to tenant user accounts and provisional tenant user accounts (invitees)
	// left unset for service accounts
	UserAttributes *UserAttributes `protobuf:"bytes,2,opt,name=user_attributes,json=userAttributes,proto3,oneof"`
}

type GroupMembershipWithAttributes_ServiceAccountAttributes struct {
	// filled with known data for members corresponding to service accounts
	// left unset for any kind of tenant user accounts
	ServiceAccountAttributes *ServiceAccountAttributes `protobuf:"bytes,3,opt,name=service_account_attributes,json=serviceAccountAttributes,proto3,oneof"`
}

type GroupMembershipWithAttributes_Error struct {
	// in a case of issues of getting attributes from pds service, we can still return some data from cpl
	Error *Error `protobuf:"bytes,99,opt,name=error,proto3,oneof"`
}

func (*GroupMembershipWithAttributes_UserAttributes) isGroupMembershipWithAttributes_AttributesOptional() {
}

func (*GroupMembershipWithAttributes_ServiceAccountAttributes) isGroupMembershipWithAttributes_AttributesOptional() {
}

func (*GroupMembershipWithAttributes_Error) isGroupMembershipWithAttributes_AttributesOptional() {}

var File_nebius_iam_v1_group_membership_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_group_membership_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a,
	0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04,
	0xba, 0x4a, 0x01, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x09,
	0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xba, 0x4a, 0x01,
	0x05, 0x52, 0x08, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x74, 0x3a, 0x04, 0xba, 0x4a, 0x01,
	0x02, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01,
	0x02, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x74, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x20, 0x0a, 0x1c, 0x4f, 0x52, 0x44, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f, 0x54, 0x45, 0x4e, 0x41,
	0x4e, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x54, 0x45, 0x4e,
	0x41, 0x4e, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x03, 0x22, 0xad, 0x03, 0x0a, 0x1d, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x57, 0x69, 0x74, 0x68, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x10, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x12, 0x4a, 0x0a, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52,
	0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x48, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x1a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x14, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x5b, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_group_membership_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_group_membership_proto_rawDescData = file_nebius_iam_v1_group_membership_proto_rawDesc
)

func file_nebius_iam_v1_group_membership_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_group_membership_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_group_membership_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_group_membership_proto_rawDescData)
	})
	return file_nebius_iam_v1_group_membership_proto_rawDescData
}

var file_nebius_iam_v1_group_membership_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_group_membership_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_iam_v1_group_membership_proto_goTypes = []any{
	(GroupMemberKind_Kind)(0),             // 0: nebius.iam.v1.GroupMemberKind.Kind
	(*GroupMembership)(nil),               // 1: nebius.iam.v1.GroupMembership
	(*GroupMembershipSpec)(nil),           // 2: nebius.iam.v1.GroupMembershipSpec
	(*GroupMembershipStatus)(nil),         // 3: nebius.iam.v1.GroupMembershipStatus
	(*GroupMemberKind)(nil),               // 4: nebius.iam.v1.GroupMemberKind
	(*GroupMembershipWithAttributes)(nil), // 5: nebius.iam.v1.GroupMembershipWithAttributes
	(*v1.ResourceMetadata)(nil),           // 6: nebius.common.v1.ResourceMetadata
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
	(*UserAttributes)(nil),                // 8: nebius.iam.v1.UserAttributes
	(*ServiceAccountAttributes)(nil),      // 9: nebius.iam.v1.ServiceAccountAttributes
	(*Error)(nil),                         // 10: nebius.iam.v1.Error
}
var file_nebius_iam_v1_group_membership_proto_depIdxs = []int32{
	6,  // 0: nebius.iam.v1.GroupMembership.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2,  // 1: nebius.iam.v1.GroupMembership.spec:type_name -> nebius.iam.v1.GroupMembershipSpec
	3,  // 2: nebius.iam.v1.GroupMembership.status:type_name -> nebius.iam.v1.GroupMembershipStatus
	7,  // 3: nebius.iam.v1.GroupMembership.revoke_at:type_name -> google.protobuf.Timestamp
	0,  // 4: nebius.iam.v1.GroupMemberKind.kind:type_name -> nebius.iam.v1.GroupMemberKind.Kind
	1,  // 5: nebius.iam.v1.GroupMembershipWithAttributes.group_membership:type_name -> nebius.iam.v1.GroupMembership
	4,  // 6: nebius.iam.v1.GroupMembershipWithAttributes.group_member_kind:type_name -> nebius.iam.v1.GroupMemberKind
	8,  // 7: nebius.iam.v1.GroupMembershipWithAttributes.user_attributes:type_name -> nebius.iam.v1.UserAttributes
	9,  // 8: nebius.iam.v1.GroupMembershipWithAttributes.service_account_attributes:type_name -> nebius.iam.v1.ServiceAccountAttributes
	10, // 9: nebius.iam.v1.GroupMembershipWithAttributes.error:type_name -> nebius.iam.v1.Error
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_group_membership_proto_init() }
func file_nebius_iam_v1_group_membership_proto_init() {
	if File_nebius_iam_v1_group_membership_proto != nil {
		return
	}
	file_nebius_iam_v1_service_account_proto_init()
	file_nebius_iam_v1_tenant_user_account_proto_init()
	file_nebius_iam_v1_group_membership_proto_msgTypes[4].OneofWrappers = []any{
		(*GroupMembershipWithAttributes_UserAttributes)(nil),
		(*GroupMembershipWithAttributes_ServiceAccountAttributes)(nil),
		(*GroupMembershipWithAttributes_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_group_membership_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_group_membership_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_group_membership_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_group_membership_proto_enumTypes,
		MessageInfos:      file_nebius_iam_v1_group_membership_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_group_membership_proto = out.File
	file_nebius_iam_v1_group_membership_proto_rawDesc = nil
	file_nebius_iam_v1_group_membership_proto_goTypes = nil
	file_nebius_iam_v1_group_membership_proto_depIdxs = nil
}
