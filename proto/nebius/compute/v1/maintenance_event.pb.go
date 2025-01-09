// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: nebius/compute/v1/maintenance_event.proto

package v1

import (
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

type MaintenanceEventStatus_State int32

const (
	MaintenanceEventStatus_STATE_UNSPECIFIED MaintenanceEventStatus_State = 0
	MaintenanceEventStatus_STATE_PENDING     MaintenanceEventStatus_State = 1
	MaintenanceEventStatus_STATE_IN_PROGRESS MaintenanceEventStatus_State = 2
	MaintenanceEventStatus_STATE_COMPLETED   MaintenanceEventStatus_State = 3
	MaintenanceEventStatus_STATE_CANCELLED   MaintenanceEventStatus_State = 4
)

// Enum value maps for MaintenanceEventStatus_State.
var (
	MaintenanceEventStatus_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_PENDING",
		2: "STATE_IN_PROGRESS",
		3: "STATE_COMPLETED",
		4: "STATE_CANCELLED",
	}
	MaintenanceEventStatus_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_PENDING":     1,
		"STATE_IN_PROGRESS": 2,
		"STATE_COMPLETED":   3,
		"STATE_CANCELLED":   4,
	}
)

func (x MaintenanceEventStatus_State) Enum() *MaintenanceEventStatus_State {
	p := new(MaintenanceEventStatus_State)
	*p = x
	return p
}

func (x MaintenanceEventStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MaintenanceEventStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_compute_v1_maintenance_event_proto_enumTypes[0].Descriptor()
}

func (MaintenanceEventStatus_State) Type() protoreflect.EnumType {
	return &file_nebius_compute_v1_maintenance_event_proto_enumTypes[0]
}

func (x MaintenanceEventStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MaintenanceEventStatus_State.Descriptor instead.
func (MaintenanceEventStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_nebius_compute_v1_maintenance_event_proto_rawDescGZIP(), []int{0, 0}
}

type MaintenanceEventStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaintenanceId         string                       `protobuf:"bytes,1,opt,name=maintenance_id,json=maintenanceId,proto3" json:"maintenance_id,omitempty"`
	State                 MaintenanceEventStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.compute.v1.MaintenanceEventStatus_State" json:"state,omitempty"`
	OperationId           string                       `protobuf:"bytes,4,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`                                   // Operation ID of stopping or recovering operation
	CreatedAt             *timestamppb.Timestamp       `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                         // Time when the maintenance event is created
	FinishedAt            *timestamppb.Timestamp       `protobuf:"bytes,6,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`                                      // Time when the maintenance event is finished
	SlaDeadlineTs         *timestamppb.Timestamp       `protobuf:"bytes,7,opt,name=sla_deadline_ts,json=slaDeadlineTs,proto3" json:"sla_deadline_ts,omitempty"`                           // Time when the instance will be force stopped
	SupportCenterTicketId string                       `protobuf:"bytes,9,opt,name=support_center_ticket_id,json=supportCenterTicketId,proto3" json:"support_center_ticket_id,omitempty"` // Ticket key, can be transformed into url where support is talking with the client
}

func (x *MaintenanceEventStatus) Reset() {
	*x = MaintenanceEventStatus{}
	mi := &file_nebius_compute_v1_maintenance_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaintenanceEventStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceEventStatus) ProtoMessage() {}

func (x *MaintenanceEventStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_compute_v1_maintenance_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceEventStatus.ProtoReflect.Descriptor instead.
func (*MaintenanceEventStatus) Descriptor() ([]byte, []int) {
	return file_nebius_compute_v1_maintenance_event_proto_rawDescGZIP(), []int{0}
}

func (x *MaintenanceEventStatus) GetMaintenanceId() string {
	if x != nil {
		return x.MaintenanceId
	}
	return ""
}

func (x *MaintenanceEventStatus) GetState() MaintenanceEventStatus_State {
	if x != nil {
		return x.State
	}
	return MaintenanceEventStatus_STATE_UNSPECIFIED
}

func (x *MaintenanceEventStatus) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *MaintenanceEventStatus) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *MaintenanceEventStatus) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *MaintenanceEventStatus) GetSlaDeadlineTs() *timestamppb.Timestamp {
	if x != nil {
		return x.SlaDeadlineTs
	}
	return nil
}

func (x *MaintenanceEventStatus) GetSupportCenterTicketId() string {
	if x != nil {
		return x.SupportCenterTicketId
	}
	return ""
}

var File_nebius_compute_v1_maintenance_event_proto protoreflect.FileDescriptor

var file_nebius_compute_v1_maintenance_event_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x92, 0x04, 0x0a, 0x16, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x45, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2f, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x6c, 0x61, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x73, 0x6c, 0x61, 0x44, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x22, 0x72, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x04, 0x42, 0x64, 0x0a, 0x18, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x15, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_nebius_compute_v1_maintenance_event_proto_rawDescOnce sync.Once
	file_nebius_compute_v1_maintenance_event_proto_rawDescData = file_nebius_compute_v1_maintenance_event_proto_rawDesc
)

func file_nebius_compute_v1_maintenance_event_proto_rawDescGZIP() []byte {
	file_nebius_compute_v1_maintenance_event_proto_rawDescOnce.Do(func() {
		file_nebius_compute_v1_maintenance_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_nebius_compute_v1_maintenance_event_proto_rawDescData)
	})
	return file_nebius_compute_v1_maintenance_event_proto_rawDescData
}

var file_nebius_compute_v1_maintenance_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_compute_v1_maintenance_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nebius_compute_v1_maintenance_event_proto_goTypes = []any{
	(MaintenanceEventStatus_State)(0), // 0: nebius.compute.v1.MaintenanceEventStatus.State
	(*MaintenanceEventStatus)(nil),    // 1: nebius.compute.v1.MaintenanceEventStatus
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_nebius_compute_v1_maintenance_event_proto_depIdxs = []int32{
	0, // 0: nebius.compute.v1.MaintenanceEventStatus.state:type_name -> nebius.compute.v1.MaintenanceEventStatus.State
	2, // 1: nebius.compute.v1.MaintenanceEventStatus.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: nebius.compute.v1.MaintenanceEventStatus.finished_at:type_name -> google.protobuf.Timestamp
	2, // 3: nebius.compute.v1.MaintenanceEventStatus.sla_deadline_ts:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nebius_compute_v1_maintenance_event_proto_init() }
func file_nebius_compute_v1_maintenance_event_proto_init() {
	if File_nebius_compute_v1_maintenance_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nebius_compute_v1_maintenance_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_compute_v1_maintenance_event_proto_goTypes,
		DependencyIndexes: file_nebius_compute_v1_maintenance_event_proto_depIdxs,
		EnumInfos:         file_nebius_compute_v1_maintenance_event_proto_enumTypes,
		MessageInfos:      file_nebius_compute_v1_maintenance_event_proto_msgTypes,
	}.Build()
	File_nebius_compute_v1_maintenance_event_proto = out.File
	file_nebius_compute_v1_maintenance_event_proto_rawDesc = nil
	file_nebius_compute_v1_maintenance_event_proto_goTypes = nil
	file_nebius_compute_v1_maintenance_event_proto_depIdxs = nil
}
