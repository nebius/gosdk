// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/iam/v1/federation_certificate.proto

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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FederationCertificateStatus_State int32

const (
	FederationCertificateStatus_STATE_UNSPECIFIED FederationCertificateStatus_State = 0
	FederationCertificateStatus_ACTIVE            FederationCertificateStatus_State = 1
	FederationCertificateStatus_EXPIRED           FederationCertificateStatus_State = 2
)

// Enum value maps for FederationCertificateStatus_State.
var (
	FederationCertificateStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "EXPIRED",
	}
	FederationCertificateStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"EXPIRED":           2,
	}
)

func (x FederationCertificateStatus_State) Enum() *FederationCertificateStatus_State {
	p := new(FederationCertificateStatus_State)
	*p = x
	return p
}

func (x FederationCertificateStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FederationCertificateStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_iam_v1_federation_certificate_proto_enumTypes[0].Descriptor()
}

func (FederationCertificateStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_iam_v1_federation_certificate_proto_enumTypes[0]
}

func (x FederationCertificateStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FederationCertificateStatus_State.Descriptor instead.
func (FederationCertificateStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_proto_rawDescGZIP(), []int{2, 0}
}

type FederationCertificate struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata         `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *FederationCertificateSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *FederationCertificateStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FederationCertificate) Reset() {
	*x = FederationCertificate{}
	mi := &file_nebius_iam_v1_federation_certificate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FederationCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationCertificate) ProtoMessage() {}

func (x *FederationCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationCertificate.ProtoReflect.Descriptor instead.
func (*FederationCertificate) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *FederationCertificate) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *FederationCertificate) GetSpec() *FederationCertificateSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FederationCertificate) GetStatus() *FederationCertificateStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type FederationCertificateSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Description   string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Data          string                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FederationCertificateSpec) Reset() {
	*x = FederationCertificateSpec{}
	mi := &file_nebius_iam_v1_federation_certificate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FederationCertificateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationCertificateSpec) ProtoMessage() {}

func (x *FederationCertificateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationCertificateSpec.ProtoReflect.Descriptor instead.
func (*FederationCertificateSpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *FederationCertificateSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FederationCertificateSpec) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type FederationCertificateStatus struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	State         FederationCertificateStatus_State `protobuf:"varint,1,opt,name=state,proto3,enum=nebius.iam.v1.FederationCertificateStatus_State" json:"state,omitempty"`
	Algorithm     string                            `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	KeySize       int64                             `protobuf:"varint,4,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	NotBefore     *timestamppb.Timestamp            `protobuf:"bytes,5,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	NotAfter      *timestamppb.Timestamp            `protobuf:"bytes,6,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FederationCertificateStatus) Reset() {
	*x = FederationCertificateStatus{}
	mi := &file_nebius_iam_v1_federation_certificate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FederationCertificateStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationCertificateStatus) ProtoMessage() {}

func (x *FederationCertificateStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_certificate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationCertificateStatus.ProtoReflect.Descriptor instead.
func (*FederationCertificateStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *FederationCertificateStatus) GetState() FederationCertificateStatus_State {
	if x != nil {
		return x.State
	}
	return FederationCertificateStatus_STATE_UNSPECIFIED
}

func (x *FederationCertificateStatus) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *FederationCertificateStatus) GetKeySize() int64 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

func (x *FederationCertificateStatus) GetNotBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *FederationCertificateStatus) GetNotAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.NotAfter
	}
	return nil
}

var File_nebius_iam_v1_federation_certificate_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_federation_certificate_proto_rawDesc = string([]byte{
	0x0a, 0x2a, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x15, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x46,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x48, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x22, 0x57, 0x0a, 0x19,
	0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcb, 0x02, 0x0a, 0x1b, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6b,
	0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x02, 0x42, 0x61, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_iam_v1_federation_certificate_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_federation_certificate_proto_rawDescData []byte
)

func file_nebius_iam_v1_federation_certificate_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_federation_certificate_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_federation_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_federation_certificate_proto_rawDesc), len(file_nebius_iam_v1_federation_certificate_proto_rawDesc)))
	})
	return file_nebius_iam_v1_federation_certificate_proto_rawDescData
}

var file_nebius_iam_v1_federation_certificate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_iam_v1_federation_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_iam_v1_federation_certificate_proto_goTypes = []any{
	(FederationCertificateStatus_State)(0), // 0: nebius.iam.v1.FederationCertificateStatus.State
	(*FederationCertificate)(nil),          // 1: nebius.iam.v1.FederationCertificate
	(*FederationCertificateSpec)(nil),      // 2: nebius.iam.v1.FederationCertificateSpec
	(*FederationCertificateStatus)(nil),    // 3: nebius.iam.v1.FederationCertificateStatus
	(*v1.ResourceMetadata)(nil),            // 4: nebius.common.v1.ResourceMetadata
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
}
var file_nebius_iam_v1_federation_certificate_proto_depIdxs = []int32{
	4, // 0: nebius.iam.v1.FederationCertificate.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2, // 1: nebius.iam.v1.FederationCertificate.spec:type_name -> nebius.iam.v1.FederationCertificateSpec
	3, // 2: nebius.iam.v1.FederationCertificate.status:type_name -> nebius.iam.v1.FederationCertificateStatus
	0, // 3: nebius.iam.v1.FederationCertificateStatus.state:type_name -> nebius.iam.v1.FederationCertificateStatus.State
	5, // 4: nebius.iam.v1.FederationCertificateStatus.not_before:type_name -> google.protobuf.Timestamp
	5, // 5: nebius.iam.v1.FederationCertificateStatus.not_after:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_federation_certificate_proto_init() }
func file_nebius_iam_v1_federation_certificate_proto_init() {
	if File_nebius_iam_v1_federation_certificate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_iam_v1_federation_certificate_proto_rawDesc), len(file_nebius_iam_v1_federation_certificate_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_federation_certificate_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_federation_certificate_proto_depIdxs,
		EnumInfos:         file_nebius_iam_v1_federation_certificate_proto_enumTypes,
		MessageInfos:      file_nebius_iam_v1_federation_certificate_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_federation_certificate_proto = out.File
	file_nebius_iam_v1_federation_certificate_proto_goTypes = nil
	file_nebius_iam_v1_federation_certificate_proto_depIdxs = nil
}
