// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/iam/v1/profile_service.proto

package v1

import (
	_ "github.com/nebius/gosdk/proto/nebius"
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

type GetProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_profile_service_proto_rawDescGZIP(), []int{0}
}

type GetProfileResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Profile:
	//
	//	*GetProfileResponse_UserProfile
	//	*GetProfileResponse_ServiceAccountProfile
	//	*GetProfileResponse_AnonymousProfile
	Profile       isGetProfileResponse_Profile `protobuf_oneof:"profile"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_profile_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetProfile() isGetProfileResponse_Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *GetProfileResponse) GetUserProfile() *UserProfile {
	if x != nil {
		if x, ok := x.Profile.(*GetProfileResponse_UserProfile); ok {
			return x.UserProfile
		}
	}
	return nil
}

func (x *GetProfileResponse) GetServiceAccountProfile() *ServiceAccountProfile {
	if x != nil {
		if x, ok := x.Profile.(*GetProfileResponse_ServiceAccountProfile); ok {
			return x.ServiceAccountProfile
		}
	}
	return nil
}

func (x *GetProfileResponse) GetAnonymousProfile() *AnonymousAccount {
	if x != nil {
		if x, ok := x.Profile.(*GetProfileResponse_AnonymousProfile); ok {
			return x.AnonymousProfile
		}
	}
	return nil
}

type isGetProfileResponse_Profile interface {
	isGetProfileResponse_Profile()
}

type GetProfileResponse_UserProfile struct {
	UserProfile *UserProfile `protobuf:"bytes,1,opt,name=user_profile,json=userProfile,proto3,oneof"`
}

type GetProfileResponse_ServiceAccountProfile struct {
	ServiceAccountProfile *ServiceAccountProfile `protobuf:"bytes,2,opt,name=service_account_profile,json=serviceAccountProfile,proto3,oneof"`
}

type GetProfileResponse_AnonymousProfile struct {
	AnonymousProfile *AnonymousAccount `protobuf:"bytes,3,opt,name=anonymous_profile,json=anonymousProfile,proto3,oneof"`
}

func (*GetProfileResponse_UserProfile) isGetProfileResponse_Profile() {}

func (*GetProfileResponse_ServiceAccountProfile) isGetProfileResponse_Profile() {}

func (*GetProfileResponse_AnonymousProfile) isGetProfileResponse_Profile() {}

type UserProfile struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FederationInfo *UserAccountExternalId `protobuf:"bytes,2,opt,name=federation_info,json=federationInfo,proto3" json:"federation_info,omitempty"`
	// Types that are valid to be assigned to AttributesOptional:
	//
	//	*UserProfile_Attributes
	//	*UserProfile_RetrievingError
	AttributesOptional isUserProfile_AttributesOptional `protobuf_oneof:"attributes_optional"`
	Tenants            []*UserTenantInfo                `protobuf:"bytes,5,rep,name=tenants,proto3" json:"tenants,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_profile_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetFederationInfo() *UserAccountExternalId {
	if x != nil {
		return x.FederationInfo
	}
	return nil
}

func (x *UserProfile) GetAttributesOptional() isUserProfile_AttributesOptional {
	if x != nil {
		return x.AttributesOptional
	}
	return nil
}

func (x *UserProfile) GetAttributes() *UserAttributes {
	if x != nil {
		if x, ok := x.AttributesOptional.(*UserProfile_Attributes); ok {
			return x.Attributes
		}
	}
	return nil
}

func (x *UserProfile) GetRetrievingError() *Error {
	if x != nil {
		if x, ok := x.AttributesOptional.(*UserProfile_RetrievingError); ok {
			return x.RetrievingError
		}
	}
	return nil
}

func (x *UserProfile) GetTenants() []*UserTenantInfo {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type isUserProfile_AttributesOptional interface {
	isUserProfile_AttributesOptional()
}

type UserProfile_Attributes struct {
	Attributes *UserAttributes `protobuf:"bytes,3,opt,name=attributes,proto3,oneof"`
}

type UserProfile_RetrievingError struct {
	RetrievingError *Error `protobuf:"bytes,4,opt,name=retrieving_error,json=retrievingError,proto3,oneof"`
}

func (*UserProfile_Attributes) isUserProfile_AttributesOptional() {}

func (*UserProfile_RetrievingError) isUserProfile_AttributesOptional() {}

type UserTenantInfo struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	TenantId            string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	TenantUserAccountId string                 `protobuf:"bytes,2,opt,name=tenant_user_account_id,json=tenantUserAccountId,proto3" json:"tenant_user_account_id,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UserTenantInfo) Reset() {
	*x = UserTenantInfo{}
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserTenantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTenantInfo) ProtoMessage() {}

func (x *UserTenantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTenantInfo.ProtoReflect.Descriptor instead.
func (*UserTenantInfo) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_profile_service_proto_rawDescGZIP(), []int{3}
}

func (x *UserTenantInfo) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *UserTenantInfo) GetTenantUserAccountId() string {
	if x != nil {
		return x.TenantUserAccountId
	}
	return ""
}

type ServiceAccountProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *ServiceAccount        `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceAccountProfile) Reset() {
	*x = ServiceAccountProfile{}
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceAccountProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountProfile) ProtoMessage() {}

func (x *ServiceAccountProfile) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountProfile.ProtoReflect.Descriptor instead.
func (*ServiceAccountProfile) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_profile_service_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceAccountProfile) GetInfo() *ServiceAccount {
	if x != nil {
		return x.Info
	}
	return nil
}

type AnonymousAccount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnonymousAccount) Reset() {
	*x = AnonymousAccount{}
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnonymousAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnonymousAccount) ProtoMessage() {}

func (x *AnonymousAccount) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_profile_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnonymousAccount.ProtoReflect.Descriptor instead.
func (*AnonymousAccount) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_profile_service_proto_rawDescGZIP(), []int{5}
}

var File_nebius_iam_v1_profile_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_profile_service_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x90, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x5e, 0x0a, 0x17, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x48, 0x00, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x10, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x52, 0x0e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x10, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x62, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4a, 0x0a,
	0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x41, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x68, 0x0a,
	0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0a, 0xba, 0x4a, 0x07,
	0x63, 0x70, 0x6c, 0x2e, 0x69, 0x61, 0x6d, 0x42, 0x5a, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v1_profile_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_profile_service_proto_rawDescData []byte
)

func file_nebius_iam_v1_profile_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_profile_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_profile_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_profile_service_proto_rawDesc), len(file_nebius_iam_v1_profile_service_proto_rawDesc)))
	})
	return file_nebius_iam_v1_profile_service_proto_rawDescData
}

var file_nebius_iam_v1_profile_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nebius_iam_v1_profile_service_proto_goTypes = []any{
	(*GetProfileRequest)(nil),     // 0: nebius.iam.v1.GetProfileRequest
	(*GetProfileResponse)(nil),    // 1: nebius.iam.v1.GetProfileResponse
	(*UserProfile)(nil),           // 2: nebius.iam.v1.UserProfile
	(*UserTenantInfo)(nil),        // 3: nebius.iam.v1.UserTenantInfo
	(*ServiceAccountProfile)(nil), // 4: nebius.iam.v1.ServiceAccountProfile
	(*AnonymousAccount)(nil),      // 5: nebius.iam.v1.AnonymousAccount
	(*UserAccountExternalId)(nil), // 6: nebius.iam.v1.UserAccountExternalId
	(*UserAttributes)(nil),        // 7: nebius.iam.v1.UserAttributes
	(*Error)(nil),                 // 8: nebius.iam.v1.Error
	(*ServiceAccount)(nil),        // 9: nebius.iam.v1.ServiceAccount
}
var file_nebius_iam_v1_profile_service_proto_depIdxs = []int32{
	2, // 0: nebius.iam.v1.GetProfileResponse.user_profile:type_name -> nebius.iam.v1.UserProfile
	4, // 1: nebius.iam.v1.GetProfileResponse.service_account_profile:type_name -> nebius.iam.v1.ServiceAccountProfile
	5, // 2: nebius.iam.v1.GetProfileResponse.anonymous_profile:type_name -> nebius.iam.v1.AnonymousAccount
	6, // 3: nebius.iam.v1.UserProfile.federation_info:type_name -> nebius.iam.v1.UserAccountExternalId
	7, // 4: nebius.iam.v1.UserProfile.attributes:type_name -> nebius.iam.v1.UserAttributes
	8, // 5: nebius.iam.v1.UserProfile.retrieving_error:type_name -> nebius.iam.v1.Error
	3, // 6: nebius.iam.v1.UserProfile.tenants:type_name -> nebius.iam.v1.UserTenantInfo
	9, // 7: nebius.iam.v1.ServiceAccountProfile.info:type_name -> nebius.iam.v1.ServiceAccount
	0, // 8: nebius.iam.v1.ProfileService.Get:input_type -> nebius.iam.v1.GetProfileRequest
	1, // 9: nebius.iam.v1.ProfileService.Get:output_type -> nebius.iam.v1.GetProfileResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_profile_service_proto_init() }
func file_nebius_iam_v1_profile_service_proto_init() {
	if File_nebius_iam_v1_profile_service_proto != nil {
		return
	}
	file_nebius_iam_v1_tenant_user_account_proto_init()
	file_nebius_iam_v1_service_account_proto_init()
	file_nebius_iam_v1_user_account_proto_init()
	file_nebius_iam_v1_profile_service_proto_msgTypes[1].OneofWrappers = []any{
		(*GetProfileResponse_UserProfile)(nil),
		(*GetProfileResponse_ServiceAccountProfile)(nil),
		(*GetProfileResponse_AnonymousProfile)(nil),
	}
	file_nebius_iam_v1_profile_service_proto_msgTypes[2].OneofWrappers = []any{
		(*UserProfile_Attributes)(nil),
		(*UserProfile_RetrievingError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_profile_service_proto_rawDesc), len(file_nebius_iam_v1_profile_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_profile_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_profile_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_profile_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_profile_service_proto = out.File
	file_nebius_iam_v1_profile_service_proto_goTypes = nil
	file_nebius_iam_v1_profile_service_proto_depIdxs = nil
}
