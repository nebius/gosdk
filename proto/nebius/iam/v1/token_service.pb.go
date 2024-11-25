// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/iam/v1/token_service.proto

package v1

import (
	_ "github.com/nebius/gosdk/proto/nebius"
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

// https://www.rfc-editor.org/rfc/rfc8693.html
type ExchangeTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrantType          string   `protobuf:"bytes,1,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`                              // required - urn:ietf:params:oauth:grant-type:token-exchange
	RequestedTokenType string   `protobuf:"bytes,2,opt,name=requested_token_type,json=requestedTokenType,proto3" json:"requested_token_type,omitempty"` // optional type of requested token, default is urn:ietf:params:oauth:token-type:access_token
	SubjectToken       string   `protobuf:"bytes,3,opt,name=subject_token,json=subjectToken,proto3" json:"subject_token,omitempty"`                     // required - could be self signed JWT token
	SubjectTokenType   string   `protobuf:"bytes,4,opt,name=subject_token_type,json=subjectTokenType,proto3" json:"subject_token_type,omitempty"`       // required, in case of jwt - urn:ietf:params:oauth:token-type:jwt
	Scopes             []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`                                                     // optional (scopes of the token)
	Audience           string   `protobuf:"bytes,6,opt,name=audience,proto3" json:"audience,omitempty"`                                                 // optional, name of the oauth client id on which this token will be used
	ActorToken         string   `protobuf:"bytes,7,opt,name=actor_token,json=actorToken,proto3" json:"actor_token,omitempty"`                           // optional, subject token for impersonation/delegation (who want to impersonate/delegate) in subject_token.
	ActorTokenType     string   `protobuf:"bytes,8,opt,name=actor_token_type,json=actorTokenType,proto3" json:"actor_token_type,omitempty"`             // optional, token type for the impersonation/delegation (who want to impersonate/delegate). Usually it's urn:ietf:params:oauth:token-type:access_token
	Resource           []string `protobuf:"bytes,9,rep,name=resource,proto3" json:"resource,omitempty"`                                                 // optional, list of resources approved to use by token, if applicable
}

func (x *ExchangeTokenRequest) Reset() {
	*x = ExchangeTokenRequest{}
	mi := &file_nebius_iam_v1_token_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExchangeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeTokenRequest) ProtoMessage() {}

func (x *ExchangeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_token_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeTokenRequest.ProtoReflect.Descriptor instead.
func (*ExchangeTokenRequest) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_token_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeTokenRequest) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *ExchangeTokenRequest) GetRequestedTokenType() string {
	if x != nil {
		return x.RequestedTokenType
	}
	return ""
}

func (x *ExchangeTokenRequest) GetSubjectToken() string {
	if x != nil {
		return x.SubjectToken
	}
	return ""
}

func (x *ExchangeTokenRequest) GetSubjectTokenType() string {
	if x != nil {
		return x.SubjectTokenType
	}
	return ""
}

func (x *ExchangeTokenRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *ExchangeTokenRequest) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *ExchangeTokenRequest) GetActorToken() string {
	if x != nil {
		return x.ActorToken
	}
	return ""
}

func (x *ExchangeTokenRequest) GetActorTokenType() string {
	if x != nil {
		return x.ActorTokenType
	}
	return ""
}

func (x *ExchangeTokenRequest) GetResource() []string {
	if x != nil {
		return x.Resource
	}
	return nil
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken     string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`               // required
	IssuedTokenType string   `protobuf:"bytes,2,opt,name=issued_token_type,json=issuedTokenType,proto3" json:"issued_token_type,omitempty"` // required
	TokenType       string   `protobuf:"bytes,3,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`                     // required - Bearer
	ExpiresIn       int64    `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	Scopes          []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	mi := &file_nebius_iam_v1_token_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_iam_v1_token_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_nebius_iam_v1_token_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CreateTokenResponse) GetIssuedTokenType() string {
	if x != nil {
		return x.IssuedTokenType
	}
	return ""
}

func (x *CreateTokenResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *CreateTokenResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *CreateTokenResponse) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

var File_nebius_iam_v1_token_service_proto protoreflect.FileDescriptor

var file_nebius_iam_v1_token_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x18, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a,
	0x14, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc8,
	0x4a, 0x01, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc8, 0x4a, 0x01, 0x52, 0x0a, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xbf,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xc8, 0x4a,
	0x01, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a,
	0x0a, 0x11, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x42, 0x58, 0x0a, 0x14, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75,
	0x62, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_nebius_iam_v1_token_service_proto_rawDescOnce sync.Once
	file_nebius_iam_v1_token_service_proto_rawDescData = file_nebius_iam_v1_token_service_proto_rawDesc
)

func file_nebius_iam_v1_token_service_proto_rawDescGZIP() []byte {
	file_nebius_iam_v1_token_service_proto_rawDescOnce.Do(func() {
		file_nebius_iam_v1_token_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_iam_v1_token_service_proto_rawDescData)
	})
	return file_nebius_iam_v1_token_service_proto_rawDescData
}

var file_nebius_iam_v1_token_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nebius_iam_v1_token_service_proto_goTypes = []any{
	(*ExchangeTokenRequest)(nil), // 0: nebius.iam.v1.ExchangeTokenRequest
	(*CreateTokenResponse)(nil),  // 1: nebius.iam.v1.CreateTokenResponse
}
var file_nebius_iam_v1_token_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nebius_iam_v1_token_service_proto_init() }
func file_nebius_iam_v1_token_service_proto_init() {
	if File_nebius_iam_v1_token_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_iam_v1_token_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_iam_v1_token_service_proto_goTypes,
		DependencyIndexes: file_nebius_iam_v1_token_service_proto_depIdxs,
		MessageInfos:      file_nebius_iam_v1_token_service_proto_msgTypes,
	}.Build()
	File_nebius_iam_v1_token_service_proto = out.File
	file_nebius_iam_v1_token_service_proto_rawDesc = nil
	file_nebius_iam_v1_token_service_proto_goTypes = nil
	file_nebius_iam_v1_token_service_proto_depIdxs = nil
}
