// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/auth_public_key.proto

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

type AuthPublicKeyStatus_State int32

const (
	AuthPublicKeyStatus_STATE_UNSPECIFIED AuthPublicKeyStatus_State = 0
	AuthPublicKeyStatus_ACTIVE            AuthPublicKeyStatus_State = 1
	AuthPublicKeyStatus_INACTIVE          AuthPublicKeyStatus_State = 2
	AuthPublicKeyStatus_EXPIRED           AuthPublicKeyStatus_State = 3
	AuthPublicKeyStatus_DELETING          AuthPublicKeyStatus_State = 4
	AuthPublicKeyStatus_DELETED           AuthPublicKeyStatus_State = 5
)

// Enum value maps for AuthPublicKeyStatus_State.
var (
	AuthPublicKeyStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "INACTIVE",
		3: "EXPIRED",
		4: "DELETING",
		5: "DELETED",
	}
	AuthPublicKeyStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"INACTIVE":          2,
		"EXPIRED":           3,
		"DELETING":          4,
		"DELETED":           5,
	}
)

func (x AuthPublicKeyStatus_State) Enum() *AuthPublicKeyStatus_State {
	p := new(AuthPublicKeyStatus_State)
	*p = x
	return p
}

func (x AuthPublicKeyStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthPublicKeyStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_auth_public_key_proto_enumTypes[0].Descriptor()
}

func (AuthPublicKeyStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_auth_public_key_proto_enumTypes[0]
}

func (x AuthPublicKeyStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthPublicKeyStatus_State.Descriptor instead.
func (AuthPublicKeyStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_proto_rawDescGZIP(), []int{2, 0}
}

type AuthPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *AuthPublicKeySpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *AuthPublicKeyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AuthPublicKey) Reset() {
	*x = AuthPublicKey{}
	mi := &file_nebius_iam_v1_auth_public_key_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthPublicKey) ProtoMessage() {}

func (x *AuthPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthPublicKey.ProtoReflect.Descriptor instead.
func (*AuthPublicKey) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_proto_rawDescGZIP(), []int{0}
}

func (x *AuthPublicKey) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AuthPublicKey) GetSpec() *AuthPublicKeySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AuthPublicKey) GetStatus() *AuthPublicKeyStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type AuthPublicKeySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     *Account               `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ExpiresAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Data        string                 `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthPublicKeySpec) Reset() {
	*x = AuthPublicKeySpec{}
	mi := &file_nebius_iam_v1_auth_public_key_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthPublicKeySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthPublicKeySpec) ProtoMessage() {}

func (x *AuthPublicKeySpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthPublicKeySpec.ProtoReflect.Descriptor instead.
func (*AuthPublicKeySpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_proto_rawDescGZIP(), []int{1}
}

func (x *AuthPublicKeySpec) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AuthPublicKeySpec) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *AuthPublicKeySpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AuthPublicKeySpec) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type AuthPublicKeyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State       AuthPublicKeyStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.iam.v1.AuthPublicKeyStatus_State" json:"state,omitempty"`
	Fingerprint string                    `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Algorithm   string                    `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	KeySize     int32                     `protobuf:"varint,4,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
}

func (x *AuthPublicKeyStatus) Reset() {
	*x = AuthPublicKeyStatus{}
	mi := &file_nebius_iam_v1_auth_public_key_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthPublicKeyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthPublicKeyStatus) ProtoMessage() {}

func (x *AuthPublicKeyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_auth_public_key_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthPublicKeyStatus.ProtoReflect.Descriptor instead.
func (*AuthPublicKeyStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_auth_public_key_proto_rawDescGZIP(), []int{2}
}

func (x *AuthPublicKeyStatus) GetState() AuthPublicKeyStatus_State {
	if x != nil {
		return x.State
	}
	return AuthPublicKeyStatus_STATE_UNSPECIFIED
}

func (x *AuthPublicKeyStatus) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *AuthPublicKeyStatus) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *AuthPublicKeyStatus) GetKeySize() int32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

var File_nebius_iam_v1_auth_public_key_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_auth_public_key_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x46, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53,
	0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x3a, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x22, 0xc8, 0x01, 0x0a, 0x11, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x36, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x92, 0x02, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x19, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x60, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x42, 0x59, 0x0a, 0x14, 0x61, 0x69, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x41, 0x75, 0x74, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_auth_public_key_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_auth_public_key_proto_rawDescData = file_nebius_iam_v1_auth_public_key_proto_rawDesc
)

func file_nebius_iam_v1_auth_public_key_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_auth_public_key_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_auth_public_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_auth_public_key_proto_rawDescData)
	})
	return file_nebius_iam_v1_auth_public_key_proto_rawDescData
}

var file_nebius_iam_v1_auth_public_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_auth_public_key_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_iam_v1_auth_public_key_proto_goTypes = []any{
	(AuthPublicKeyStatus_State)(0), // 0: nebius.iam.v1.AuthPublicKeyStatus.State
	(*AuthPublicKey)(nil),          // 1: nebius.iam.v1.AuthPublicKey
	(*AuthPublicKeySpec)(nil),      // 2: nebius.iam.v1.AuthPublicKeySpec
	(*AuthPublicKeyStatus)(nil),    // 3: nebius.iam.v1.AuthPublicKeyStatus
	(*v1.ResourceMetadata)(nil),    // 4: nebius.common.v1.ResourceMetadata
	(*Account)(nil),                // 5: nebius.iam.v1.Account
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_nebius_iam_v1_auth_public_key_proto_depIdxs = []int32{
	4, // 0: nebius.iam.v1.AuthPublicKey.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2, // 1: nebius.iam.v1.AuthPublicKey.spec:type_name -> nebius.iam.v1.AuthPublicKeySpec
	3, // 2: nebius.iam.v1.AuthPublicKey.status:type_name -> nebius.iam.v1.AuthPublicKeyStatus
	5, // 3: nebius.iam.v1.AuthPublicKeySpec.account:type_name -> nebius.iam.v1.Account
	6, // 4: nebius.iam.v1.AuthPublicKeySpec.expires_at:type_name -> google.protobuf.Timestamp
	0, // 5: nebius.iam.v1.AuthPublicKeyStatus.state:type_name -> nebius.iam.v1.AuthPublicKeyStatus.State
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_auth_public_key_proto_init() }
func file_nebius_iam_v1_auth_public_key_proto_init() {
	if File_nebius_iam_v1_auth_public_key_proto != nil {
		return
	}
	file_nebius_iam_v1_access_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_auth_public_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_auth_public_key_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_auth_public_key_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_auth_public_key_proto_enumTypes,
		MessageInfos:      file_nebius_iam_v1_auth_public_key_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_auth_public_key_proto = out.File
	file_nebius_iam_v1_auth_public_key_proto_rawDesc = nil
	file_nebius_iam_v1_auth_public_key_proto_goTypes = nil
	file_nebius_iam_v1_auth_public_key_proto_depIdxs = nil
}
