// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/iam/v2/tenant_service.proto

package v2

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

type GetTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantRequest) Reset() {
	*x = GetTenantRequest{}
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantRequest) ProtoMessage() {}

func (x *GetTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantRequest.ProtoReflect.Descriptor instead.
func (*GetTenantRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v2_tenant_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetTenantRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTenantByNameRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Only empty value is allowed. Field is needed for compatibility.
	// parent_id parameter for tenants doesn't make real sense, because tenants are top-level objects.
	ParentId      string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTenantByNameRequest) Reset() {
	*x = GetTenantByNameRequest{}
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTenantByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantByNameRequest) ProtoMessage() {}

func (x *GetTenantByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantByNameRequest.ProtoReflect.Descriptor instead.
func (*GetTenantByNameRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v2_tenant_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTenantByNameRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *GetTenantByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateTenantRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Metadata *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *TenantSpec            `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// Tenant name prefix. A few random characters will be added to this prefix.
	NamePrefix    string `protobuf:"bytes,3,opt,name=name_prefix,json=namePrefix,proto3" json:"name_prefix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTenantRequest) Reset() {
	*x = UpdateTenantRequest{}
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenantRequest) ProtoMessage() {}

func (x *UpdateTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenantRequest.ProtoReflect.Descriptor instead.
func (*UpdateTenantRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v2_tenant_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTenantRequest) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateTenantRequest) GetSpec() *TenantSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *UpdateTenantRequest) GetNamePrefix() string {
	if x != nil {
		return x.NamePrefix
	}
	return ""
}

type ListTenantsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Specifies the maximum number of items to return in the response.
	// Default value: 10
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token for pagination, allowing the retrieval of the next set of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter to narrow down the results based on specific criteria.
	Filter        string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantsRequest) Reset() {
	*x = ListTenantsRequest{}
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsRequest) ProtoMessage() {}

func (x *ListTenantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsRequest.ProtoReflect.Descriptor instead.
func (*ListTenantsRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v2_tenant_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListTenantsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTenantsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTenantsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListTenantsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Items []*Tenant              `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	// Token for pagination, indicating the next set of results can be retrieved using this token.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTenantsResponse) Reset() {
	*x = ListTenantsResponse{}
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTenantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTenantsResponse) ProtoMessage() {}

func (x *ListTenantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v2_tenant_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTenantsResponse.ProtoReflect.Descriptor instead.
func (*ListTenantsResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v2_tenant_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListTenantsResponse) GetItems() []*Tenant {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListTenantsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_nebius_iam_v2_tenant_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v2_tenant_service_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xa5, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61,
	0x6d, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x68, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x6a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xbf,
	0x02, 0x0a, 0x0d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12,
	0x49, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x32, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x21, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0a, 0xba, 0x4a, 0x07, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x61, 0x6d,
	0x42, 0x59, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v2_tenant_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v2_tenant_service_proto_rawDescData []byte
)

func file_nebius_iam_v2_tenant_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v2_tenant_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v2_tenant_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v2_tenant_service_proto_rawDesc), len(file_nebius_iam_v2_tenant_service_proto_rawDesc)))
	})
	return file_nebius_iam_v2_tenant_service_proto_rawDescData
}

var file_nebius_iam_v2_tenant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_iam_v2_tenant_service_proto_goTypes = []any{
	(*GetTenantRequest)(nil),       // 0: nebius.iam.v2.GetTenantRequest
	(*GetTenantByNameRequest)(nil), // 1: nebius.iam.v2.GetTenantByNameRequest
	(*UpdateTenantRequest)(nil),    // 2: nebius.iam.v2.UpdateTenantRequest
	(*ListTenantsRequest)(nil),     // 3: nebius.iam.v2.ListTenantsRequest
	(*ListTenantsResponse)(nil),    // 4: nebius.iam.v2.ListTenantsResponse
	(*v1.ResourceMetadata)(nil),    // 5: nebius.common.v1.ResourceMetadata
	(*TenantSpec)(nil),             // 6: nebius.iam.v2.TenantSpec
	(*Tenant)(nil),                 // 7: nebius.iam.v2.Tenant
	(*v1.Operation)(nil),           // 8: nebius.common.v1.Operation
}
var file_nebius_iam_v2_tenant_service_proto_depIdxs = []int32{
	5, // 0: nebius.iam.v2.UpdateTenantRequest.metadata:type_name -> nebius.common.v1.ResourceMetadata
	6, // 1: nebius.iam.v2.UpdateTenantRequest.spec:type_name -> nebius.iam.v2.TenantSpec
	7, // 2: nebius.iam.v2.ListTenantsResponse.items:type_name -> nebius.iam.v2.Tenant
	0, // 3: nebius.iam.v2.TenantService.Get:input_type -> nebius.iam.v2.GetTenantRequest
	1, // 4: nebius.iam.v2.TenantService.GetByName:input_type -> nebius.iam.v2.GetTenantByNameRequest
	3, // 5: nebius.iam.v2.TenantService.List:input_type -> nebius.iam.v2.ListTenantsRequest
	2, // 6: nebius.iam.v2.TenantService.Update:input_type -> nebius.iam.v2.UpdateTenantRequest
	7, // 7: nebius.iam.v2.TenantService.Get:output_type -> nebius.iam.v2.Tenant
	7, // 8: nebius.iam.v2.TenantService.GetByName:output_type -> nebius.iam.v2.Tenant
	4, // 9: nebius.iam.v2.TenantService.List:output_type -> nebius.iam.v2.ListTenantsResponse
	8, // 10: nebius.iam.v2.TenantService.Update:output_type -> nebius.common.v1.Operation
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nebius_iam_v2_tenant_service_proto_init() }
func file_nebius_iam_v2_tenant_service_proto_init() {
	if File_nebius_iam_v2_tenant_service_proto != nil {
		return
	}
	file_nebius_iam_v2_tenant_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v2_tenant_service_proto_rawDesc), len(file_nebius_iam_v2_tenant_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nebius_iam_v2_tenant_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v2_tenant_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v2_tenant_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v2_tenant_service_proto = out.File
	file_nebius_iam_v2_tenant_service_proto_goTypes = nil
	file_nebius_iam_v2_tenant_service_proto_depIdxs = nil
}
