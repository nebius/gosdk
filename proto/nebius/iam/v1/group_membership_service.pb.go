// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/group_membership_service.proto

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

type CreateGroupMembershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata         *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec             *GroupMembershipSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	RevokeAfterHours int64                `protobuf:"varint,4,opt,name=revoke_after_hours,json=revokeAfterHours,proto3" json:"revoke_after_hours,omitempty"`
}

func (x *CreateGroupMembershipRequest) Reset() {
	*x = CreateGroupMembershipRequest{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupMembershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupMembershipRequest) ProtoMessage() {}

func (x *CreateGroupMembershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupMembershipRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupMembershipRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupMembershipRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateGroupMembershipRequest) GetSpec() *GroupMembershipSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *CreateGroupMembershipRequest) GetRevokeAfterHours() int64 {
	if x != nil {
		return x.RevokeAfterHours
	}
	return 0
}

type DeleteGroupMembershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroupMembershipRequest) Reset() {
	*x = DeleteGroupMembershipRequest{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGroupMembershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupMembershipRequest) ProtoMessage() {}

func (x *DeleteGroupMembershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupMembershipRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupMembershipRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteGroupMembershipRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGroupMembershipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroupMembershipRequest) Reset() {
	*x = GetGroupMembershipRequest{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupMembershipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMembershipRequest) ProtoMessage() {}

func (x *GetGroupMembershipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMembershipRequest.ProtoReflect.Descriptor instead.
func (*GetGroupMembershipRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupMembershipRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListGroupMembershipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Default value: 10
	PageSize  *int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListGroupMembershipsRequest) Reset() {
	*x = ListGroupMembershipsRequest{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupMembershipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupMembershipsRequest) ProtoMessage() {}

func (x *ListGroupMembershipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupMembershipsRequest.ProtoReflect.Descriptor instead.
func (*ListGroupMembershipsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListGroupMembershipsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListGroupMembershipsRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListGroupMembershipsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListGroupMembershipsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListMemberOfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested subject id. Can be tenant user account id or service account id.
	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// Default value: 10
	PageSize  *int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter    string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListMemberOfRequest) Reset() {
	*x = ListMemberOfRequest{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMemberOfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemberOfRequest) ProtoMessage() {}

func (x *ListMemberOfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemberOfRequest.ProtoReflect.Descriptor instead.
func (*ListMemberOfRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListMemberOfRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *ListMemberOfRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListMemberOfRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListMemberOfRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListGroupMembershipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Members of the group. Can be tenant user account ids or service account ids.
	Memberships   []*GroupMembership `protobuf:"bytes,1,rep,name=memberships,proto3" json:"memberships,omitempty"`
	NextPageToken string             `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListGroupMembershipsResponse) Reset() {
	*x = ListGroupMembershipsResponse{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupMembershipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupMembershipsResponse) ProtoMessage() {}

func (x *ListGroupMembershipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupMembershipsResponse.ProtoReflect.Descriptor instead.
func (*ListGroupMembershipsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListGroupMembershipsResponse) GetMemberships() []*GroupMembership {
	if x != nil {
		return x.Memberships
	}
	return nil
}

func (x *ListGroupMembershipsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListGroupMembershipsWithAttributesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Members of the group with their attributes if supported by member type.
	// Can be tenant user accounts (regular or provisional/invited) or service accounts.
	// Regular tenant user accounts may have a full set of attributes, provisional may have contacts only,
	// service accounts
	Memberships   []*GroupMembershipWithAttributes `protobuf:"bytes,1,rep,name=memberships,proto3" json:"memberships,omitempty"`
	NextPageToken string                           `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListGroupMembershipsWithAttributesResponse) Reset() {
	*x = ListGroupMembershipsWithAttributesResponse{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGroupMembershipsWithAttributesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupMembershipsWithAttributesResponse) ProtoMessage() {}

func (x *ListGroupMembershipsWithAttributesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupMembershipsWithAttributesResponse.ProtoReflect.Descriptor instead.
func (*ListGroupMembershipsWithAttributesResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListGroupMembershipsWithAttributesResponse) GetMemberships() []*GroupMembershipWithAttributes {
	if x != nil {
		return x.Memberships
	}
	return nil
}

func (x *ListGroupMembershipsWithAttributesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListMemberOfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Groups that requested entity is a member of
	Items         []*Group `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListMemberOfResponse) Reset() {
	*x = ListMemberOfResponse{}
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMemberOfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemberOfResponse) ProtoMessage() {}

func (x *ListMemberOfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_group_membership_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemberOfResponse.ProtoReflect.Descriptor instead.
func (*ListMemberOfResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListMemberOfResponse) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListMemberOfResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_iam_v1_group_membership_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_group_membership_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42,
	0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x36, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x76,
	0x6f, 0x6b, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x36, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x33, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0xa3, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4f,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0xa4, 0x01, 0x0a, 0x2a, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x57, 0x69, 0x74, 0x68, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x57, 0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe3, 0x04, 0x0a, 0x16, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x52, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x28, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x12, 0x52, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x66, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x82, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2a,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x12, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0a,
	0xba, 0x4a, 0x07, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x61, 0x6d, 0x42, 0x62, 0x0a, 0x14, 0x61, 0x69,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x42, 0x1b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_group_membership_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_group_membership_service_proto_rawDescData = file_nebius_iam_v1_group_membership_service_proto_rawDesc
)

func file_nebius_iam_v1_group_membership_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_group_membership_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_group_membership_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_group_membership_service_proto_rawDescData)
	})
	return file_nebius_iam_v1_group_membership_service_proto_rawDescData
}

var file_nebius_iam_v1_group_membership_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_nebius_iam_v1_group_membership_service_proto_goTypes = []any{
	(*CreateGroupMembershipRequest)(nil),               // 0: nebius.iam.v1.CreateGroupMembershipRequest
	(*DeleteGroupMembershipRequest)(nil),               // 1: nebius.iam.v1.DeleteGroupMembershipRequest
	(*GetGroupMembershipRequest)(nil),                  // 2: nebius.iam.v1.GetGroupMembershipRequest
	(*ListGroupMembershipsRequest)(nil),                // 3: nebius.iam.v1.ListGroupMembershipsRequest
	(*ListMemberOfRequest)(nil),                        // 4: nebius.iam.v1.ListMemberOfRequest
	(*ListGroupMembershipsResponse)(nil),               // 5: nebius.iam.v1.ListGroupMembershipsResponse
	(*ListGroupMembershipsWithAttributesResponse)(nil), // 6: nebius.iam.v1.ListGroupMembershipsWithAttributesResponse
	(*ListMemberOfResponse)(nil),                       // 7: nebius.iam.v1.ListMemberOfResponse
	(*v1.ResourceMetadata)(nil),                        // 8: nebius.common.v1.ResourceMetadata
	(*GroupMembershipSpec)(nil),                        // 9: nebius.iam.v1.GroupMembershipSpec
	(*GroupMembership)(nil),                            // 10: nebius.iam.v1.GroupMembership
	(*GroupMembershipWithAttributes)(nil),              // 11: nebius.iam.v1.GroupMembershipWithAttributes
	(*Group)(nil),                                      // 12: nebius.iam.v1.Group
	(*v1.Operation)(nil),                               // 13: nebius.common.v1.Operation
}
var file_nebius_iam_v1_group_membership_service_proto_depIdxs = []int32{
	8,  // 0: nebius.iam.v1.CreateGroupMembershipRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	9,  // 1: nebius.iam.v1.CreateGroupMembershipRequest.spec:type_name -> nebius.iam.v1.GroupMembershipSpec
	10, // 2: nebius.iam.v1.ListGroupMembershipsResponse.memberships:type_name -> nebius.iam.v1.GroupMembership
	11, // 3: nebius.iam.v1.ListGroupMembershipsWithAttributesResponse.memberships:type_name -> nebius.iam.v1.GroupMembershipWithAttributes
	12, // 4: nebius.iam.v1.ListMemberOfResponse.items:type_name -> nebius.iam.v1.Group
	0,  // 5: nebius.iam.v1.GroupMembershipService.Create:input_type -> nebius.iam.v1.CreateGroupMembershipRequest
	2,  // 6: nebius.iam.v1.GroupMembershipService.Get:input_type -> nebius.iam.v1.GetGroupMembershipRequest
	1,  // 7: nebius.iam.v1.GroupMembershipService.Delete:input_type -> nebius.iam.v1.DeleteGroupMembershipRequest
	3,  // 8: nebius.iam.v1.GroupMembershipService.ListMembers:input_type -> nebius.iam.v1.ListGroupMembershipsRequest
	3,  // 9: nebius.iam.v1.GroupMembershipService.ListMembersWithAttributes:input_type -> nebius.iam.v1.ListGroupMembershipsRequest
	4,  // 10: nebius.iam.v1.GroupMembershipService.ListMemberOf:input_type -> nebius.iam.v1.ListMemberOfRequest
	13, // 11: nebius.iam.v1.GroupMembershipService.Create:output_type -> nebius.common.v1.Operation
	10, // 12: nebius.iam.v1.GroupMembershipService.Get:output_type -> nebius.iam.v1.GroupMembership
	13, // 13: nebius.iam.v1.GroupMembershipService.Delete:output_type -> nebius.common.v1.Operation
	5,  // 14: nebius.iam.v1.GroupMembershipService.ListMembers:output_type -> nebius.iam.v1.ListGroupMembershipsResponse
	6,  // 15: nebius.iam.v1.GroupMembershipService.ListMembersWithAttributes:output_type -> nebius.iam.v1.ListGroupMembershipsWithAttributesResponse
	7,  // 16: nebius.iam.v1.GroupMembershipService.ListMemberOf:output_type -> nebius.iam.v1.ListMemberOfResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_group_membership_service_proto_init() }
func file_nebius_iam_v1_group_membership_service_proto_init() {
	if File_nebius_iam_v1_group_membership_service_proto != nil {
		return
	}
	file_nebius_iam_v1_group_membership_proto_init()
	file_nebius_iam_v1_group_proto_init()
	file_nebius_iam_v1_group_membership_service_proto_msgTypes[3].OneofWrappers = []any{}
	file_nebius_iam_v1_group_membership_service_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_group_membership_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_group_membership_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_group_membership_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_group_membership_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_group_membership_service_proto = out.File
	file_nebius_iam_v1_group_membership_service_proto_rawDesc = nil
	file_nebius_iam_v1_group_membership_service_proto_goTypes = nil
	file_nebius_iam_v1_group_membership_service_proto_depIdxs = nil
}
