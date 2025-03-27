// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.3
// source: nebius/iam/v1/tenant_user_account_service.proto

package v1

import (
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

type GetTenantUserAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // tenant user account id like 'tenantuseraccount-{region}someuniquesuffix'
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantUserAccountRequest) Reset() {
	*x = GetTenantUserAccountRequest{}
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantUserAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantUserAccountRequest) ProtoMessage() {}

func (x *GetTenantUserAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantUserAccountRequest.ProtoReflect.Descriptor instead.
func (*GetTenantUserAccountRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantUserAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTenantUserAccountsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the tenant ID like 'tenant-someuniqueprefix'
	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Specifies the maximum number of items to return in the response.
	// Default value: 10
	PageSize *int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantUserAccountsRequest) Reset() {
	*x = ListTenantUserAccountsRequest{}
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantUserAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantUserAccountsRequest) ProtoMessage() {}

func (x *ListTenantUserAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantUserAccountsRequest.ProtoReflect.Descriptor instead.
func (*ListTenantUserAccountsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListTenantUserAccountsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListTenantUserAccountsRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListTenantUserAccountsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTenantUserAccountsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListTenantUserAccountsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of service accounts returned in the response. The field should be named as `items` for consistency.
	Items []*TenantUserAccount `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantUserAccountsResponse) Reset() {
	*x = ListTenantUserAccountsResponse{}
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantUserAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantUserAccountsResponse) ProtoMessage() {}

func (x *ListTenantUserAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantUserAccountsResponse.ProtoReflect.Descriptor instead.
func (*ListTenantUserAccountsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTenantUserAccountsResponse) GetItems() []*TenantUserAccount {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTenantUserAccountsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type BlockTenantUserAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlockTenantUserAccountRequest) Reset() {
	*x = BlockTenantUserAccountRequest{}
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlockTenantUserAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockTenantUserAccountRequest) ProtoMessage() {}

func (x *BlockTenantUserAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockTenantUserAccountRequest.ProtoReflect.Descriptor instead.
func (*BlockTenantUserAccountRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *BlockTenantUserAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnblockTenantUserAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnblockTenantUserAccountRequest) Reset() {
	*x = UnblockTenantUserAccountRequest{}
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnblockTenantUserAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockTenantUserAccountRequest) ProtoMessage() {}

func (x *UnblockTenantUserAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockTenantUserAccountRequest.ProtoReflect.Descriptor instead.
func (*UnblockTenantUserAccountRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_service_proto_rawDescGZIP(), []int{4}
}

func (x *UnblockTenantUserAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_nebius_iam_v1_tenant_user_account_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_tenant_user_account_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc0, 0x4a, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x2f, 0x0a, 0x1d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x1f, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8c, 0x03, 0x0a, 0x18, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x63, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x56, 0x0a, 0x07, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0xba, 0x4a, 0x07, 0x63, 0x70,
	0x6c, 0x2e, 0x69, 0x61, 0x6d, 0x42, 0x64, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x1d, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_tenant_user_account_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_tenant_user_account_service_proto_rawDescData = file_nebius_iam_v1_tenant_user_account_service_proto_rawDesc
)

func file_nebius_iam_v1_tenant_user_account_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_tenant_user_account_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_tenant_user_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_tenant_user_account_service_proto_rawDescData)
	})
	return file_nebius_iam_v1_tenant_user_account_service_proto_rawDescData
}

var file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_iam_v1_tenant_user_account_service_proto_goTypes = []any{
	(*GetTenantUserAccountRequest)(nil),     // 0: nebius.iam.v1.GetTenantUserAccountRequest
	(*ListTenantUserAccountsRequest)(nil),   // 1: nebius.iam.v1.ListTenantUserAccountsRequest
	(*ListTenantUserAccountsResponse)(nil),  // 2: nebius.iam.v1.ListTenantUserAccountsResponse
	(*BlockTenantUserAccountRequest)(nil),   // 3: nebius.iam.v1.BlockTenantUserAccountRequest
	(*UnblockTenantUserAccountRequest)(nil), // 4: nebius.iam.v1.UnblockTenantUserAccountRequest
	(*TenantUserAccount)(nil),               // 5: nebius.iam.v1.TenantUserAccount
	(*v1.Operation)(nil),                    // 6: nebius.common.v1.Operation
}
var file_nebius_iam_v1_tenant_user_account_service_proto_depIdxs = []int32{
	5, // 0: nebius.iam.v1.ListTenantUserAccountsResponse.items:type_name -> nebius.iam.v1.TenantUserAccount
	0, // 1: nebius.iam.v1.TenantUserAccountService.Get:input_type -> nebius.iam.v1.GetTenantUserAccountRequest
	1, // 2: nebius.iam.v1.TenantUserAccountService.List:input_type -> nebius.iam.v1.ListTenantUserAccountsRequest
	3, // 3: nebius.iam.v1.TenantUserAccountService.Block:input_type -> nebius.iam.v1.BlockTenantUserAccountRequest
	4, // 4: nebius.iam.v1.TenantUserAccountService.Unblock:input_type -> nebius.iam.v1.UnblockTenantUserAccountRequest
	5, // 5: nebius.iam.v1.TenantUserAccountService.Get:output_type -> nebius.iam.v1.TenantUserAccount
	2, // 6: nebius.iam.v1.TenantUserAccountService.List:output_type -> nebius.iam.v1.ListTenantUserAccountsResponse
	6, // 7: nebius.iam.v1.TenantUserAccountService.Block:output_type -> nebius.common.v1.Operation
	6, // 8: nebius.iam.v1.TenantUserAccountService.Unblock:output_type -> nebius.common.v1.Operation
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_tenant_user_account_service_proto_init() }
func file_nebius_iam_v1_tenant_user_account_service_proto_init() {
	if File_nebius_iam_v1_tenant_user_account_service_proto != nil {
		return
	}
	file_nebius_iam_v1_tenant_user_account_proto_init()
	file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_tenant_user_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_tenant_user_account_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_tenant_user_account_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_tenant_user_account_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_tenant_user_account_service_proto = out.File
	file_nebius_iam_v1_tenant_user_account_service_proto_rawDesc = nil
	file_nebius_iam_v1_tenant_user_account_service_proto_goTypes = nil
	file_nebius_iam_v1_tenant_user_account_service_proto_depIdxs = nil
}
