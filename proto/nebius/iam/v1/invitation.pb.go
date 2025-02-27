// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v4.25.1
// source: nebius/iam/v1/invitation.proto

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

type InvitationStatus_State int32

const (
	InvitationStatus_UNSPECIFIED InvitationStatus_State = 0
	InvitationStatus_CREATING    InvitationStatus_State = 4 // contacts data is not stored in pds yet. probably will GC it later
	InvitationStatus_CREATED     InvitationStatus_State = 5 // notification is not sent yet
	InvitationStatus_PENDING     InvitationStatus_State = 1 // notification is sent, we are waiting for the user to approve the notification
	InvitationStatus_EXPIRED     InvitationStatus_State = 2 // notification is expired, accept is no longer possible
	InvitationStatus_ACCEPTED    InvitationStatus_State = 3 // notification is accepted
)

// Enum value maps for InvitationStatus_State.
var (
	InvitationStatus_State_name = map[int32]string{
		0: "UNSPECIFIED",
		4: "CREATING",
		5: "CREATED",
		1: "PENDING",
		2: "EXPIRED",
		3: "ACCEPTED",
	}
	InvitationStatus_State_value = map[string]int32{
		"UNSPECIFIED": 0,
		"CREATING":    4,
		"CREATED":     5,
		"PENDING":     1,
		"EXPIRED":     2,
		"ACCEPTED":    3,
	}
)

func (x InvitationStatus_State) Enum() *InvitationStatus_State {
	p := new(InvitationStatus_State)
	*p = x
	return p
}

func (x InvitationStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvitationStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_invitation_proto_enumTypes[0].Descriptor()
}

func (InvitationStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_invitation_proto_enumTypes[0]
}

func (x InvitationStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvitationStatus_State.Descriptor instead.
func (InvitationStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_invitation_proto_rawDescGZIP(), []int{2, 0}
}

type Invitation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *InvitationSpec        `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *InvitationStatus      `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	mi := &file_nebius_iam_v1_invitation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_invitation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_invitation_proto_rawDescGZIP(), []int{0}
}

func (x *Invitation) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Invitation) GetSpec() *InvitationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Invitation) GetStatus() *InvitationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type InvitationSpec struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	Description string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are valid to be assigned to Contact:
	//
	//	*InvitationSpec_Email
	Contact       isInvitationSpec_Contact `protobuf_oneof:"contact"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvitationSpec) Reset() {
	*x = InvitationSpec{}
	mi := &file_nebius_iam_v1_invitation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvitationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationSpec) ProtoMessage() {}

func (x *InvitationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_invitation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationSpec.ProtoReflect.Descriptor instead.
func (*InvitationSpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_invitation_proto_rawDescGZIP(), []int{1}
}

func (x *InvitationSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InvitationSpec) GetContact() isInvitationSpec_Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *InvitationSpec) GetEmail() string {
	if x != nil {
		if x, ok := x.Contact.(*InvitationSpec_Email); ok {
			return x.Email
		}
	}
	return ""
}

type isInvitationSpec_Contact interface {
	isInvitationSpec_Contact()
}

type InvitationSpec_Email struct {
	Email string `protobuf:"bytes,11,opt,name=email,proto3,oneof"`
}

func (*InvitationSpec_Email) isInvitationSpec_Contact() {}

type InvitationStatus struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	TenantUserAccountId string                 `protobuf:"bytes,1,opt,name=tenant_user_account_id,json=tenantUserAccountId,proto3" json:"tenant_user_account_id,omitempty"`
	ExpiresAt           *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	State               InvitationStatus_State `protobuf:"varint,3,opt,name=state,proto3,enum=nebius.iam.v1.InvitationStatus_State" json:"state,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *InvitationStatus) Reset() {
	*x = InvitationStatus{}
	mi := &file_nebius_iam_v1_invitation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvitationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationStatus) ProtoMessage() {}

func (x *InvitationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_invitation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationStatus.ProtoReflect.Descriptor instead.
func (*InvitationStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_invitation_proto_rawDescGZIP(), []int{2}
}

func (x *InvitationStatus) GetTenantUserAccountId() string {
	if x != nil {
		return x.TenantUserAccountId
	}
	return ""
}

func (x *InvitationStatus) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *InvitationStatus) GetState() InvitationStatus_State {
	if x != nil {
		return x.State
	}
	return InvitationStatus_UNSPECIFIED
}

var File_nebius_iam_v1_invitation_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_invitation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x39, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x22,
	0x5a, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x9c, 0x02, 0x0a, 0x10,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x33, 0x0a, 0x16, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74,
	0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x56, 0x0a, 0x14, 0x61, 0x69,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x42, 0x0f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_invitation_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_invitation_proto_rawDescData = file_nebius_iam_v1_invitation_proto_rawDesc
)

func file_nebius_iam_v1_invitation_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_invitation_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_invitation_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_invitation_proto_rawDescData)
	})
	return file_nebius_iam_v1_invitation_proto_rawDescData
}

var file_nebius_iam_v1_invitation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_invitation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_iam_v1_invitation_proto_goTypes = []any{
	(InvitationStatus_State)(0),   // 0: nebius.iam.v1.InvitationStatus.State
	(*Invitation)(nil),            // 1: nebius.iam.v1.Invitation
	(*InvitationSpec)(nil),        // 2: nebius.iam.v1.InvitationSpec
	(*InvitationStatus)(nil),      // 3: nebius.iam.v1.InvitationStatus
	(*v1.ResourceMetadata)(nil),   // 4: nebius.common.v1.ResourceMetadata
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_nebius_iam_v1_invitation_proto_depIdxs = []int32{
	4, // 0: nebius.iam.v1.Invitation.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2, // 1: nebius.iam.v1.Invitation.spec:type_name -> nebius.iam.v1.InvitationSpec
	3, // 2: nebius.iam.v1.Invitation.status:type_name -> nebius.iam.v1.InvitationStatus
	5, // 3: nebius.iam.v1.InvitationStatus.expires_at:type_name -> google.protobuf.Timestamp
	0, // 4: nebius.iam.v1.InvitationStatus.state:type_name -> nebius.iam.v1.InvitationStatus.State
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_invitation_proto_init() }
func file_nebius_iam_v1_invitation_proto_init() {
	if File_nebius_iam_v1_invitation_proto != nil {
		return
	}
	file_nebius_iam_v1_invitation_proto_msgTypes[1].OneofWrappers = []any{
		(*InvitationSpec_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_invitation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_invitation_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_invitation_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_invitation_proto_enumTypes,
		MessageInfos:      file_nebius_iam_v1_invitation_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_invitation_proto = out.File
	file_nebius_iam_v1_invitation_proto_rawDesc = nil
	file_nebius_iam_v1_invitation_proto_goTypes = nil
	file_nebius_iam_v1_invitation_proto_depIdxs = nil
}
