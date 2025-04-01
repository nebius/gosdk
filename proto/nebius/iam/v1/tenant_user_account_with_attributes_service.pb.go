// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/iam/v1/tenant_user_account_with_attributes_service.proto

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

type GetTenantUserAccountWithAttributesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // tenant user account id like 'tenantuseraccount-{region}someuniquesuffix'
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantUserAccountWithAttributesRequest) Reset() {
	*x = GetTenantUserAccountWithAttributesRequest{}
	mi := &file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantUserAccountWithAttributesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantUserAccountWithAttributesRequest) ProtoMessage() {}

func (x *GetTenantUserAccountWithAttributesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantUserAccountWithAttributesRequest.ProtoReflect.Descriptor instead.
func (*GetTenantUserAccountWithAttributesRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantUserAccountWithAttributesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListTenantUserAccountsWithAttributesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the tenant ID like 'tenant-{region}someuniquesuffix'
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

func (x *ListTenantUserAccountsWithAttributesRequest) Reset() {
	*x = ListTenantUserAccountsWithAttributesRequest{}
	mi := &file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantUserAccountsWithAttributesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantUserAccountsWithAttributesRequest) ProtoMessage() {}

func (x *ListTenantUserAccountsWithAttributesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantUserAccountsWithAttributesRequest.ProtoReflect.Descriptor instead.
func (*ListTenantUserAccountsWithAttributesRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListTenantUserAccountsWithAttributesRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListTenantUserAccountsWithAttributesRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListTenantUserAccountsWithAttributesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTenantUserAccountsWithAttributesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListTenantUserAccountsWithAttributesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of user accounts returned in the response. The field should be named as `items` for consistency.
	Items []*TenantUserAccountWithAttributes `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantUserAccountsWithAttributesResponse) Reset() {
	*x = ListTenantUserAccountsWithAttributesResponse{}
	mi := &file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantUserAccountsWithAttributesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantUserAccountsWithAttributesResponse) ProtoMessage() {}

func (x *ListTenantUserAccountsWithAttributesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantUserAccountsWithAttributesResponse.ProtoReflect.Descriptor instead.
func (*ListTenantUserAccountsWithAttributesResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListTenantUserAccountsWithAttributesResponse) GetItems() []*TenantUserAccountWithAttributes {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTenantUserAccountsWithAttributesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_iam_v1_tenant_user_account_with_attributes_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDesc = string([]byte{
	0x0a, 0x3f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xb6, 0x01, 0x0a, 0x2b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xc0, 0x4a, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x2c, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xa6, 0x02, 0x0a, 0x26, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x38, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x7f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0a, 0xba, 0x4a, 0x07, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x61,
	0x6d, 0x42, 0x72, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70,
	0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x2b, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64,
	0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescData []byte
)

func file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDesc), len(file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDesc)))
	})
	return file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDescData
}

var file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_goTypes = []any{
	(*GetTenantUserAccountWithAttributesRequest)(nil),    // 0: nebius.iam.v1.GetTenantUserAccountWithAttributesRequest
	(*ListTenantUserAccountsWithAttributesRequest)(nil),  // 1: nebius.iam.v1.ListTenantUserAccountsWithAttributesRequest
	(*ListTenantUserAccountsWithAttributesResponse)(nil), // 2: nebius.iam.v1.ListTenantUserAccountsWithAttributesResponse
	(*TenantUserAccountWithAttributes)(nil),              // 3: nebius.iam.v1.TenantUserAccountWithAttributes
}
var file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_depIdxs = []int32{
	3, // 0: nebius.iam.v1.ListTenantUserAccountsWithAttributesResponse.items:type_name -> nebius.iam.v1.TenantUserAccountWithAttributes
	0, // 1: nebius.iam.v1.TenantUserAccountWithAttributesService.Get:input_type -> nebius.iam.v1.GetTenantUserAccountWithAttributesRequest
	1, // 2: nebius.iam.v1.TenantUserAccountWithAttributesService.List:input_type -> nebius.iam.v1.ListTenantUserAccountsWithAttributesRequest
	3, // 3: nebius.iam.v1.TenantUserAccountWithAttributesService.Get:output_type -> nebius.iam.v1.TenantUserAccountWithAttributes
	2, // 4: nebius.iam.v1.TenantUserAccountWithAttributesService.List:output_type -> nebius.iam.v1.ListTenantUserAccountsWithAttributesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_init() }
func file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_init() {
	if File_nebius_iam_v1_tenant_user_account_with_attributes_service_proto != nil {
		return
	}
	file_nebius_iam_v1_tenant_user_account_proto_init()
	file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDesc), len(file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_tenant_user_account_with_attributes_service_proto = out.File
	file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_goTypes = nil
	file_nebius_iam_v1_tenant_user_account_with_attributes_service_proto_depIdxs = nil
}
