// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/tenant_user_account.proto

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

type TenantUserAccountStatus_State int32

const (
	TenantUserAccountStatus_STATE_UNSPECIFIED TenantUserAccountStatus_State = 0
	TenantUserAccountStatus_ACTIVE            TenantUserAccountStatus_State = 1
	TenantUserAccountStatus_INACTIVE          TenantUserAccountStatus_State = 2
	TenantUserAccountStatus_BLOCKED           TenantUserAccountStatus_State = 3
)

// Enum value maps for TenantUserAccountStatus_State.
var (
	TenantUserAccountStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "INACTIVE",
		3: "BLOCKED",
	}
	TenantUserAccountStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"INACTIVE":          2,
		"BLOCKED":           3,
	}
)

func (x TenantUserAccountStatus_State) Enum() *TenantUserAccountStatus_State {
	p := new(TenantUserAccountStatus_State)
	*p = x
	return p
}

func (x TenantUserAccountStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TenantUserAccountStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_tenant_user_account_proto_enumTypes[0].Descriptor()
}

func (TenantUserAccountStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_tenant_user_account_proto_enumTypes[0]
}

func (x TenantUserAccountStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TenantUserAccountStatus_State.Descriptor instead.
func (TenantUserAccountStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{5, 0}
}

type TenantUserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ResourceMetadata     `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *TenantUserAccountSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *TenantUserAccountStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TenantUserAccount) Reset() {
	*x = TenantUserAccount{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantUserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUserAccount) ProtoMessage() {}

func (x *TenantUserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUserAccount.ProtoReflect.Descriptor instead.
func (*TenantUserAccount) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{0}
}

func (x *TenantUserAccount) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TenantUserAccount) GetSpec() *TenantUserAccountSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TenantUserAccount) GetStatus() *TenantUserAccountStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type TenantUserAccountWithAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantUserAccount *TenantUserAccount `protobuf:"bytes,1,opt,name=tenant_user_account,json=tenantUserAccount,proto3" json:"tenant_user_account,omitempty"`
	// Types that are assignable to AttributesOptional:
	//
	//	*TenantUserAccountWithAttributes_Attributes
	//	*TenantUserAccountWithAttributes_Error
	AttributesOptional isTenantUserAccountWithAttributes_AttributesOptional `protobuf_oneof:"attributesOptional"`
}

func (x *TenantUserAccountWithAttributes) Reset() {
	*x = TenantUserAccountWithAttributes{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantUserAccountWithAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUserAccountWithAttributes) ProtoMessage() {}

func (x *TenantUserAccountWithAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUserAccountWithAttributes.ProtoReflect.Descriptor instead.
func (*TenantUserAccountWithAttributes) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{1}
}

func (x *TenantUserAccountWithAttributes) GetTenantUserAccount() *TenantUserAccount {
	if x != nil {
		return x.TenantUserAccount
	}
	return nil
}

func (m *TenantUserAccountWithAttributes) GetAttributesOptional() isTenantUserAccountWithAttributes_AttributesOptional {
	if m != nil {
		return m.AttributesOptional
	}
	return nil
}

func (x *TenantUserAccountWithAttributes) GetAttributes() *UserAttributes {
	if x, ok := x.GetAttributesOptional().(*TenantUserAccountWithAttributes_Attributes); ok {
		return x.Attributes
	}
	return nil
}

func (x *TenantUserAccountWithAttributes) GetError() *Error {
	if x, ok := x.GetAttributesOptional().(*TenantUserAccountWithAttributes_Error); ok {
		return x.Error
	}
	return nil
}

type isTenantUserAccountWithAttributes_AttributesOptional interface {
	isTenantUserAccountWithAttributes_AttributesOptional()
}

type TenantUserAccountWithAttributes_Attributes struct {
	Attributes *UserAttributes `protobuf:"bytes,2,opt,name=attributes,proto3,oneof"`
}

type TenantUserAccountWithAttributes_Error struct {
	// in a case of issues of getting attributes from pds service, we can still return some data from cpl
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*TenantUserAccountWithAttributes_Attributes) isTenantUserAccountWithAttributes_AttributesOptional() {
}

func (*TenantUserAccountWithAttributes_Error) isTenantUserAccountWithAttributes_AttributesOptional() {
}

type UserAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sub               *string `protobuf:"bytes,20,opt,name=sub,proto3,oneof" json:"sub,omitempty"`
	Name              *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	GivenName         *string `protobuf:"bytes,2,opt,name=given_name,json=givenName,proto3,oneof" json:"given_name,omitempty"`
	FamilyName        *string `protobuf:"bytes,3,opt,name=family_name,json=familyName,proto3,oneof" json:"family_name,omitempty"`
	PreferredUsername *string `protobuf:"bytes,4,opt,name=preferred_username,json=preferredUsername,proto3,oneof" json:"preferred_username,omitempty"`
	Picture           *string `protobuf:"bytes,5,opt,name=picture,proto3,oneof" json:"picture,omitempty"`
	Email             *string `protobuf:"bytes,6,opt,name=email,proto3,oneof" json:"email,omitempty"`
	// Deprecated: Marked as deprecated in nebius/iam/v1/tenant_user_account.proto.
	EmailVerified *bool `protobuf:"varint,7,opt,name=email_verified,json=emailVerified,proto3,oneof" json:"email_verified,omitempty"`
	// Deprecated: Marked as deprecated in nebius/iam/v1/tenant_user_account.proto.
	Zoneinfo    *string `protobuf:"bytes,8,opt,name=zoneinfo,proto3,oneof" json:"zoneinfo,omitempty"`
	Locale      *string `protobuf:"bytes,9,opt,name=locale,proto3,oneof" json:"locale,omitempty"`
	PhoneNumber *string `protobuf:"bytes,10,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	// Deprecated: Marked as deprecated in nebius/iam/v1/tenant_user_account.proto.
	PhoneNumberVerified *bool `protobuf:"varint,11,opt,name=phone_number_verified,json=phoneNumberVerified,proto3,oneof" json:"phone_number_verified,omitempty"`
}

func (x *UserAttributes) Reset() {
	*x = UserAttributes{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAttributes) ProtoMessage() {}

func (x *UserAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAttributes.ProtoReflect.Descriptor instead.
func (*UserAttributes) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{2}
}

func (x *UserAttributes) GetSub() string {
	if x != nil && x.Sub != nil {
		return *x.Sub
	}
	return ""
}

func (x *UserAttributes) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UserAttributes) GetGivenName() string {
	if x != nil && x.GivenName != nil {
		return *x.GivenName
	}
	return ""
}

func (x *UserAttributes) GetFamilyName() string {
	if x != nil && x.FamilyName != nil {
		return *x.FamilyName
	}
	return ""
}

func (x *UserAttributes) GetPreferredUsername() string {
	if x != nil && x.PreferredUsername != nil {
		return *x.PreferredUsername
	}
	return ""
}

func (x *UserAttributes) GetPicture() string {
	if x != nil && x.Picture != nil {
		return *x.Picture
	}
	return ""
}

func (x *UserAttributes) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

// Deprecated: Marked as deprecated in nebius/iam/v1/tenant_user_account.proto.
func (x *UserAttributes) GetEmailVerified() bool {
	if x != nil && x.EmailVerified != nil {
		return *x.EmailVerified
	}
	return false
}

// Deprecated: Marked as deprecated in nebius/iam/v1/tenant_user_account.proto.
func (x *UserAttributes) GetZoneinfo() string {
	if x != nil && x.Zoneinfo != nil {
		return *x.Zoneinfo
	}
	return ""
}

func (x *UserAttributes) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *UserAttributes) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

// Deprecated: Marked as deprecated in nebius/iam/v1/tenant_user_account.proto.
func (x *UserAttributes) GetPhoneNumberVerified() bool {
	if x != nil && x.PhoneNumberVerified != nil {
		return *x.PhoneNumberVerified
	}
	return false
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TenantUserAccountSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisibleAttributes *TenantUserAccountSpec_VisibleAttributes `protobuf:"bytes,1,opt,name=visible_attributes,json=visibleAttributes,proto3" json:"visible_attributes,omitempty"`
}

func (x *TenantUserAccountSpec) Reset() {
	*x = TenantUserAccountSpec{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantUserAccountSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUserAccountSpec) ProtoMessage() {}

func (x *TenantUserAccountSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUserAccountSpec.ProtoReflect.Descriptor instead.
func (*TenantUserAccountSpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{4}
}

func (x *TenantUserAccountSpec) GetVisibleAttributes() *TenantUserAccountSpec_VisibleAttributes {
	if x != nil {
		return x.VisibleAttributes
	}
	return nil
}

type TenantUserAccountStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State TenantUserAccountStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.iam.v1.TenantUserAccountStatus_State" json:"state,omitempty"`
	// if a tenant user account is created during invitation it gets a reference to the invitation resource
	// once invitation is accepted it looses this reference (and internally gets a reference to their global federated user account)
	InvitationId string `protobuf:"bytes,2,opt,name=invitation_id,json=invitationId,proto3" json:"invitation_id,omitempty"`
	// currently can only accept the values: custom, unknown, google, github.
	FederationId string `protobuf:"bytes,3,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
}

func (x *TenantUserAccountStatus) Reset() {
	*x = TenantUserAccountStatus{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantUserAccountStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUserAccountStatus) ProtoMessage() {}

func (x *TenantUserAccountStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUserAccountStatus.ProtoReflect.Descriptor instead.
func (*TenantUserAccountStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{5}
}

func (x *TenantUserAccountStatus) GetState() TenantUserAccountStatus_State {
	if x != nil {
		return x.State
	}
	return TenantUserAccountStatus_STATE_UNSPECIFIED
}

func (x *TenantUserAccountStatus) GetInvitationId() string {
	if x != nil {
		return x.InvitationId
	}
	return ""
}

func (x *TenantUserAccountStatus) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

// when a global user account is projected to a specific tenant
// they can give consent for that tenant's owner to view specific personal data
// by listing explicitly visible PDS attributes
// complete list of PDS attributes is described in ../../pds/inner/v1alpha1/iam_identifier.proto
type TenantUserAccountSpec_VisibleAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attribute []string `protobuf:"bytes,1,rep,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *TenantUserAccountSpec_VisibleAttributes) Reset() {
	*x = TenantUserAccountSpec_VisibleAttributes{}
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantUserAccountSpec_VisibleAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantUserAccountSpec_VisibleAttributes) ProtoMessage() {}

func (x *TenantUserAccountSpec_VisibleAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantUserAccountSpec_VisibleAttributes.ProtoReflect.Descriptor instead.
func (*TenantUserAccountSpec_VisibleAttributes) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP(), []int{4, 0}
}

func (x *TenantUserAccountSpec_VisibleAttributes) GetAttribute() []string {
	if x != nil {
		return x.Attribute
	}
	return nil
}

var File_nebius_iam_v1_tenant_user_account_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_tenant_user_account_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe9, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x40,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x22, 0xf8, 0x01, 0x0a,
	0x1f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x50, 0x0a, 0x13, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x11, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x14, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x89, 0x06, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x03, 0x73, 0x75,
	0x62, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x00, 0x52, 0x03,
	0x73, 0x75, 0x62, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x02, 0x52,
	0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x03, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x04, 0x52, 0x11, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x22, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x05, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x75, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x5a, 0xba, 0x48, 0x54, 0xba, 0x01, 0x51, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a,
	0x1c, 0x74, 0x68, 0x69, 0x73, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x20, 0x7c, 0x7c, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x28, 0x29, 0xc0, 0x4a, 0x01,
	0x48, 0x06, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x0e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x48, 0x07, 0x52, 0x0d, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x08,
	0x7a, 0x6f, 0x6e, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05,
	0xc0, 0x4a, 0x01, 0x18, 0x01, 0x48, 0x08, 0x52, 0x08, 0x7a, 0x6f, 0x6e, 0x65, 0x69, 0x6e, 0x66,
	0x6f, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x48, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a,
	0x01, 0x48, 0x0a, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x15, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x48, 0x0b, 0x52, 0x13, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x75, 0x62, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x65, 0x0a, 0x12, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x11, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x36, 0x0a, 0x11, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x09,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x03, 0xc0, 0x4a, 0x01, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22,
	0xee, 0x01, 0x0a, 0x17, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03,
	0x42, 0x5d, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_tenant_user_account_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_tenant_user_account_proto_rawDescData = file_nebius_iam_v1_tenant_user_account_proto_rawDesc
)

func file_nebius_iam_v1_tenant_user_account_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_tenant_user_account_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_tenant_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_tenant_user_account_proto_rawDescData)
	})
	return file_nebius_iam_v1_tenant_user_account_proto_rawDescData
}

var file_nebius_iam_v1_tenant_user_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_tenant_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_nebius_iam_v1_tenant_user_account_proto_goTypes = []any{
	(TenantUserAccountStatus_State)(0),              // 0: nebius.iam.v1.TenantUserAccountStatus.State
	(*TenantUserAccount)(nil),                       // 1: nebius.iam.v1.TenantUserAccount
	(*TenantUserAccountWithAttributes)(nil),         // 2: nebius.iam.v1.TenantUserAccountWithAttributes
	(*UserAttributes)(nil),                          // 3: nebius.iam.v1.UserAttributes
	(*Error)(nil),                                   // 4: nebius.iam.v1.Error
	(*TenantUserAccountSpec)(nil),                   // 5: nebius.iam.v1.TenantUserAccountSpec
	(*TenantUserAccountStatus)(nil),                 // 6: nebius.iam.v1.TenantUserAccountStatus
	(*TenantUserAccountSpec_VisibleAttributes)(nil), // 7: nebius.iam.v1.TenantUserAccountSpec.VisibleAttributes
	(*v1.ResourceMetadata)(nil),                     // 8: nebius.common.v1.ResourceMetadata
}
var file_nebius_iam_v1_tenant_user_account_proto_depIdxs = []int32{
	8, // 0: nebius.iam.v1.TenantUserAccount.metadata:type_name -> nebius.common.v1.ResourceMetadata
	5, // 1: nebius.iam.v1.TenantUserAccount.spec:type_name -> nebius.iam.v1.TenantUserAccountSpec
	6, // 2: nebius.iam.v1.TenantUserAccount.status:type_name -> nebius.iam.v1.TenantUserAccountStatus
	1, // 3: nebius.iam.v1.TenantUserAccountWithAttributes.tenant_user_account:type_name -> nebius.iam.v1.TenantUserAccount
	3, // 4: nebius.iam.v1.TenantUserAccountWithAttributes.attributes:type_name -> nebius.iam.v1.UserAttributes
	4, // 5: nebius.iam.v1.TenantUserAccountWithAttributes.error:type_name -> nebius.iam.v1.Error
	7, // 6: nebius.iam.v1.TenantUserAccountSpec.visible_attributes:type_name -> nebius.iam.v1.TenantUserAccountSpec.VisibleAttributes
	0, // 7: nebius.iam.v1.TenantUserAccountStatus.state:type_name -> nebius.iam.v1.TenantUserAccountStatus.State
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_tenant_user_account_proto_init() }
func file_nebius_iam_v1_tenant_user_account_proto_init() {
	if File_nebius_iam_v1_tenant_user_account_proto != nil {
		return
	}
	file_nebius_iam_v1_tenant_user_account_proto_msgTypes[1].OneofWrappers = []any{
		(*TenantUserAccountWithAttributes_Attributes)(nil),
		(*TenantUserAccountWithAttributes_Error)(nil),
	}
	file_nebius_iam_v1_tenant_user_account_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_tenant_user_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_tenant_user_account_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_tenant_user_account_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_tenant_user_account_proto_enumTypes,
		MessageInfos:      file_nebius_iam_v1_tenant_user_account_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_tenant_user_account_proto = out.File
	file_nebius_iam_v1_tenant_user_account_proto_rawDesc = nil
	file_nebius_iam_v1_tenant_user_account_proto_goTypes = nil
	file_nebius_iam_v1_tenant_user_account_proto_depIdxs = nil
}
