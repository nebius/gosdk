// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/federation.proto

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

type Federation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ResourceMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *FederationSpec      `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *FederationStatus    `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Federation) Reset() {
	*x = Federation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Federation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Federation) ProtoMessage() {}

func (x *Federation) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Federation.ProtoReflect.Descriptor instead.
func (*Federation) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_proto_rawDescGZIP(), []int{0}
}

func (x *Federation) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Federation) GetSpec() *FederationSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Federation) GetStatus() *FederationStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type FederationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccountAutoCreation bool `protobuf:"varint,1,opt,name=user_account_auto_creation,json=userAccountAutoCreation,proto3" json:"user_account_auto_creation,omitempty"`
	Active                  bool `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	// Types that are assignable to Settings:
	//
	//	*FederationSpec_SamlSettings
	Settings isFederationSpec_Settings `protobuf_oneof:"settings"`
}

func (x *FederationSpec) Reset() {
	*x = FederationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationSpec) ProtoMessage() {}

func (x *FederationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationSpec.ProtoReflect.Descriptor instead.
func (*FederationSpec) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_proto_rawDescGZIP(), []int{1}
}

func (x *FederationSpec) GetUserAccountAutoCreation() bool {
	if x != nil {
		return x.UserAccountAutoCreation
	}
	return false
}

func (x *FederationSpec) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (m *FederationSpec) GetSettings() isFederationSpec_Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (x *FederationSpec) GetSamlSettings() *SamlSettings {
	if x, ok := x.GetSettings().(*FederationSpec_SamlSettings); ok {
		return x.SamlSettings
	}
	return nil
}

type isFederationSpec_Settings interface {
	isFederationSpec_Settings()
}

type FederationSpec_SamlSettings struct {
	SamlSettings *SamlSettings `protobuf:"bytes,10,opt,name=saml_settings,json=samlSettings,proto3,oneof"`
}

func (*FederationSpec_SamlSettings) isFederationSpec_Settings() {}

type SamlSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpIssuer string `protobuf:"bytes,1,opt,name=idp_issuer,json=idpIssuer,proto3" json:"idp_issuer,omitempty"`
	SsoUrl    string `protobuf:"bytes,2,opt,name=sso_url,json=ssoUrl,proto3" json:"sso_url,omitempty"`
}

func (x *SamlSettings) Reset() {
	*x = SamlSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamlSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamlSettings) ProtoMessage() {}

func (x *SamlSettings) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamlSettings.ProtoReflect.Descriptor instead.
func (*SamlSettings) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_proto_rawDescGZIP(), []int{2}
}

func (x *SamlSettings) GetIdpIssuer() string {
	if x != nil {
		return x.IdpIssuer
	}
	return ""
}

func (x *SamlSettings) GetSsoUrl() string {
	if x != nil {
		return x.SsoUrl
	}
	return ""
}

type FederationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FederationStatus) Reset() {
	*x = FederationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nebius_iam_v1_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederationStatus) ProtoMessage() {}

func (x *FederationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederationStatus.ProtoReflect.Descriptor instead.
func (*FederationStatus) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_federation_proto_rawDescGZIP(), []int{3}
}

var File_nebius_iam_v1_federation_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_federation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0a, 0x46, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04, 0xba, 0x4a, 0x01, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x04, 0xba, 0x4a, 0x01, 0x02, 0x22, 0xc2,
	0x01, 0x0a, 0x0e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x3b, 0x0a, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x6c, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x6d, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x61,
	0x6d, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x11, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x04, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x61, 0x6d, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x70, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x73, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x73, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x12, 0x0a, 0x10, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x56, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f,
	0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nebius_iam_v1_federation_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_federation_proto_rawDescData = file_nebius_iam_v1_federation_proto_rawDesc
)

func file_nebius_iam_v1_federation_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_federation_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_federation_proto_rawDescData)
	})
	return file_nebius_iam_v1_federation_proto_rawDescData
}

var file_nebius_iam_v1_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nebius_iam_v1_federation_proto_goTypes = []any{
	(*Federation)(nil),          // 0: nebius.iam.v1.Federation
	(*FederationSpec)(nil),      // 1: nebius.iam.v1.FederationSpec
	(*SamlSettings)(nil),        // 2: nebius.iam.v1.SamlSettings
	(*FederationStatus)(nil),    // 3: nebius.iam.v1.FederationStatus
	(*v1.ResourceMetadata)(nil), // 4: nebius.common.v1.ResourceMetadata
}
var file_nebius_iam_v1_federation_proto_depIdxs = []int32{
	4, // 0: nebius.iam.v1.Federation.metadata:type_name -> nebius.common.v1.ResourceMetadata
	1, // 1: nebius.iam.v1.Federation.spec:type_name -> nebius.iam.v1.FederationSpec
	3, // 2: nebius.iam.v1.Federation.status:type_name -> nebius.iam.v1.FederationStatus
	2, // 3: nebius.iam.v1.FederationSpec.saml_settings:type_name -> nebius.iam.v1.SamlSettings
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_federation_proto_init() }
func file_nebius_iam_v1_federation_proto_init() {
	if File_nebius_iam_v1_federation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nebius_iam_v1_federation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Federation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_iam_v1_federation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FederationSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_iam_v1_federation_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SamlSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nebius_iam_v1_federation_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FederationStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_nebius_iam_v1_federation_proto_msgTypes[1].OneofWrappers = []any{
		(*FederationSpec_SamlSettings)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_federation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_federation_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_federation_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_federation_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_federation_proto = out.File
	file_nebius_iam_v1_federation_proto_rawDesc = nil
	file_nebius_iam_v1_federation_proto_goTypes = nil
	file_nebius_iam_v1_federation_proto_depIdxs = nil
}
