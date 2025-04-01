// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: nebius/msp/postgresql/v1alpha1/backup.proto

package v1alpha1

import (
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

type Backup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the PostgreSQL cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,2,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Creation start timestamp.
	CreationStart *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=creation_start,json=creationStart,proto3" json:"creation_start,omitempty"`
	// Creation finish timestamp.
	CreationFinish *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=creation_finish,json=creationFinish,proto3" json:"creation_finish,omitempty"`
	// Name of the PostgreSQL cluster that the backup was created for.
	SourceClusterName string `protobuf:"bytes,5,opt,name=source_cluster_name,json=sourceClusterName,proto3" json:"source_cluster_name,omitempty"`
	// Is PostgreSQL cluster that the backup was created for visible.
	SourceClusterVisible bool `protobuf:"varint,6,opt,name=source_cluster_visible,json=sourceClusterVisible,proto3" json:"source_cluster_visible,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Backup) Reset() {
	*x = Backup{}
	mi := &file_nebius_msp_postgresql_v1alpha1_backup_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Backup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Backup) ProtoMessage() {}

func (x *Backup) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_postgresql_v1alpha1_backup_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Backup.ProtoReflect.Descriptor instead.
func (*Backup) Descriptor() ([]byte, []int) {
	return file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescGZIP(), []int{0}
}

func (x *Backup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Backup) GetSourceClusterId() string {
	if x != nil {
		return x.SourceClusterId
	}
	return ""
}

func (x *Backup) GetCreationStart() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationStart
	}
	return nil
}

func (x *Backup) GetCreationFinish() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationFinish
	}
	return nil
}

func (x *Backup) GetSourceClusterName() string {
	if x != nil {
		return x.SourceClusterName
	}
	return ""
}

func (x *Backup) GetSourceClusterVisible() bool {
	if x != nil {
		return x.SourceClusterVisible
	}
	return false
}

var File_nebius_msp_postgresql_v1alpha1_backup_proto protoreflect.FileDescriptor

var file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2,
	0x02, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x2e, 0x0a,
	0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x16, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x6c, 0x65, 0x42, 0x74, 0x0a, 0x25, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67,
	0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x71, 0x6c,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescOnce sync.Once
	file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescData []byte
)

func file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescGZIP() []byte {
	file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescOnce.Do(func() {
		file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDesc), len(file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDesc)))
	})
	return file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDescData
}

var file_nebius_msp_postgresql_v1alpha1_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nebius_msp_postgresql_v1alpha1_backup_proto_goTypes = []any{
	(*Backup)(nil),                // 0: nebius.msp.postgresql.v1alpha1.Backup
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_nebius_msp_postgresql_v1alpha1_backup_proto_depIdxs = []int32{
	1, // 0: nebius.msp.postgresql.v1alpha1.Backup.creation_start:type_name -> google.protobuf.Timestamp
	1, // 1: nebius.msp.postgresql.v1alpha1.Backup.creation_finish:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nebius_msp_postgresql_v1alpha1_backup_proto_init() }
func file_nebius_msp_postgresql_v1alpha1_backup_proto_init() {
	if File_nebius_msp_postgresql_v1alpha1_backup_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDesc), len(file_nebius_msp_postgresql_v1alpha1_backup_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_postgresql_v1alpha1_backup_proto_goTypes,
		DependencyIndexes: file_nebius_msp_postgresql_v1alpha1_backup_proto_depIdxs,
		MessageInfos:      file_nebius_msp_postgresql_v1alpha1_backup_proto_msgTypes,
	}.Build()
	File_nebius_msp_postgresql_v1alpha1_backup_proto = out.File
	file_nebius_msp_postgresql_v1alpha1_backup_proto_goTypes = nil
	file_nebius_msp_postgresql_v1alpha1_backup_proto_depIdxs = nil
}
