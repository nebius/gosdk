// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/storage/v1/bucket_counters.proto

package v1

import (
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

type CurrentBucketCounters struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	SimpleObjectsQuantity    int64                  `protobuf:"varint,1,opt,name=simple_objects_quantity,json=simpleObjectsQuantity,proto3" json:"simple_objects_quantity,omitempty"`
	SimpleObjectsSize        int64                  `protobuf:"varint,2,opt,name=simple_objects_size,json=simpleObjectsSize,proto3" json:"simple_objects_size,omitempty"`
	MultipartObjectsQuantity int64                  `protobuf:"varint,3,opt,name=multipart_objects_quantity,json=multipartObjectsQuantity,proto3" json:"multipart_objects_quantity,omitempty"`
	MultipartObjectsSize     int64                  `protobuf:"varint,4,opt,name=multipart_objects_size,json=multipartObjectsSize,proto3" json:"multipart_objects_size,omitempty"`
	MultipartUploadsQuantity int64                  `protobuf:"varint,5,opt,name=multipart_uploads_quantity,json=multipartUploadsQuantity,proto3" json:"multipart_uploads_quantity,omitempty"`
	InflightPartsQuantity    int64                  `protobuf:"varint,6,opt,name=inflight_parts_quantity,json=inflightPartsQuantity,proto3" json:"inflight_parts_quantity,omitempty"`
	InflightPartsSize        int64                  `protobuf:"varint,7,opt,name=inflight_parts_size,json=inflightPartsSize,proto3" json:"inflight_parts_size,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *CurrentBucketCounters) Reset() {
	*x = CurrentBucketCounters{}
	mi := &file_nebius_storage_v1_bucket_counters_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentBucketCounters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBucketCounters) ProtoMessage() {}

func (x *CurrentBucketCounters) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_bucket_counters_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBucketCounters.ProtoReflect.Descriptor instead.
func (*CurrentBucketCounters) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_counters_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentBucketCounters) GetSimpleObjectsQuantity() int64 {
	if x != nil {
		return x.SimpleObjectsQuantity
	}
	return 0
}

func (x *CurrentBucketCounters) GetSimpleObjectsSize() int64 {
	if x != nil {
		return x.SimpleObjectsSize
	}
	return 0
}

func (x *CurrentBucketCounters) GetMultipartObjectsQuantity() int64 {
	if x != nil {
		return x.MultipartObjectsQuantity
	}
	return 0
}

func (x *CurrentBucketCounters) GetMultipartObjectsSize() int64 {
	if x != nil {
		return x.MultipartObjectsSize
	}
	return 0
}

func (x *CurrentBucketCounters) GetMultipartUploadsQuantity() int64 {
	if x != nil {
		return x.MultipartUploadsQuantity
	}
	return 0
}

func (x *CurrentBucketCounters) GetInflightPartsQuantity() int64 {
	if x != nil {
		return x.InflightPartsQuantity
	}
	return 0
}

func (x *CurrentBucketCounters) GetInflightPartsSize() int64 {
	if x != nil {
		return x.InflightPartsSize
	}
	return 0
}

// Counters for non-current object versions (for versioning buckets).
type NonCurrentBucketCounters struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	SimpleObjectsQuantity    int64                  `protobuf:"varint,1,opt,name=simple_objects_quantity,json=simpleObjectsQuantity,proto3" json:"simple_objects_quantity,omitempty"`
	SimpleObjectsSize        int64                  `protobuf:"varint,2,opt,name=simple_objects_size,json=simpleObjectsSize,proto3" json:"simple_objects_size,omitempty"`
	MultipartObjectsQuantity int64                  `protobuf:"varint,3,opt,name=multipart_objects_quantity,json=multipartObjectsQuantity,proto3" json:"multipart_objects_quantity,omitempty"`
	MultipartObjectsSize     int64                  `protobuf:"varint,4,opt,name=multipart_objects_size,json=multipartObjectsSize,proto3" json:"multipart_objects_size,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *NonCurrentBucketCounters) Reset() {
	*x = NonCurrentBucketCounters{}
	mi := &file_nebius_storage_v1_bucket_counters_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NonCurrentBucketCounters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonCurrentBucketCounters) ProtoMessage() {}

func (x *NonCurrentBucketCounters) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_bucket_counters_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonCurrentBucketCounters.ProtoReflect.Descriptor instead.
func (*NonCurrentBucketCounters) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_counters_proto_rawDescGZIP(), []int{1}
}

func (x *NonCurrentBucketCounters) GetSimpleObjectsQuantity() int64 {
	if x != nil {
		return x.SimpleObjectsQuantity
	}
	return 0
}

func (x *NonCurrentBucketCounters) GetSimpleObjectsSize() int64 {
	if x != nil {
		return x.SimpleObjectsSize
	}
	return 0
}

func (x *NonCurrentBucketCounters) GetMultipartObjectsQuantity() int64 {
	if x != nil {
		return x.MultipartObjectsQuantity
	}
	return 0
}

func (x *NonCurrentBucketCounters) GetMultipartObjectsSize() int64 {
	if x != nil {
		return x.MultipartObjectsSize
	}
	return 0
}

type BucketCounters struct {
	state              protoimpl.MessageState    `protogen:"open.v1"`
	StorageClass       StorageClass              `protobuf:"varint,1,opt,name=storage_class,json=storageClass,proto3,enum=nebius.storage.v1.StorageClass" json:"storage_class,omitempty"`
	Counters           *CurrentBucketCounters    `protobuf:"bytes,2,opt,name=counters,proto3" json:"counters,omitempty"`
	NonCurrentCounters *NonCurrentBucketCounters `protobuf:"bytes,3,opt,name=non_current_counters,json=nonCurrentCounters,proto3" json:"non_current_counters,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *BucketCounters) Reset() {
	*x = BucketCounters{}
	mi := &file_nebius_storage_v1_bucket_counters_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BucketCounters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketCounters) ProtoMessage() {}

func (x *BucketCounters) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_storage_v1_bucket_counters_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketCounters.ProtoReflect.Descriptor instead.
func (*BucketCounters) Descriptor() ([]byte, []int) {
	return file_nebius_storage_v1_bucket_counters_proto_rawDescGZIP(), []int{2}
}

func (x *BucketCounters) GetStorageClass() StorageClass {
	if x != nil {
		return x.StorageClass
	}
	return StorageClass_STORAGE_CLASS_UNSPECIFIED
}

func (x *BucketCounters) GetCounters() *CurrentBucketCounters {
	if x != nil {
		return x.Counters
	}
	return nil
}

func (x *BucketCounters) GetNonCurrentCounters() *NonCurrentBucketCounters {
	if x != nil {
		return x.NonCurrentCounters
	}
	return nil
}

var File_nebius_storage_v1_bucket_counters_proto protoreflect.FileDescriptor

var file_nebius_storage_v1_bucket_counters_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x6e, 0x65,
	0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x15, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x13,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3c, 0x0a, 0x1a,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x61, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x3c, 0x0a, 0x1a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x36,
	0x0a, 0x17, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73,
	0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x15, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x72, 0x74, 0x73, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x69, 0x6e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x18, 0x4e, 0x6f, 0x6e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x18, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x61, 0x72, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0xfb, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5d,
	0x0a, 0x14, 0x6e, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x12, 0x6e, 0x6f, 0x6e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x42, 0x62, 0x0a,
	0x18, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_storage_v1_bucket_counters_proto_rawDescOnce sync.Once
	file_nebius_storage_v1_bucket_counters_proto_rawDescData []byte
)

func file_nebius_storage_v1_bucket_counters_proto_rawDescGZIP() []byte {
	file_nebius_storage_v1_bucket_counters_proto_rawDescOnce.Do(func() {
		file_nebius_storage_v1_bucket_counters_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_storage_v1_bucket_counters_proto_rawDesc), len(file_nebius_storage_v1_bucket_counters_proto_rawDesc)))
	})
	return file_nebius_storage_v1_bucket_counters_proto_rawDescData
}

var file_nebius_storage_v1_bucket_counters_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nebius_storage_v1_bucket_counters_proto_goTypes = []any{
	(*CurrentBucketCounters)(nil),    // 0: nebius.storage.v1.CurrentBucketCounters
	(*NonCurrentBucketCounters)(nil), // 1: nebius.storage.v1.NonCurrentBucketCounters
	(*BucketCounters)(nil),           // 2: nebius.storage.v1.BucketCounters
	(StorageClass)(0),                // 3: nebius.storage.v1.StorageClass
}
var file_nebius_storage_v1_bucket_counters_proto_depIdxs = []int32{
	3, // 0: nebius.storage.v1.BucketCounters.storage_class:type_name -> nebius.storage.v1.StorageClass
	0, // 1: nebius.storage.v1.BucketCounters.counters:type_name -> nebius.storage.v1.CurrentBucketCounters
	1, // 2: nebius.storage.v1.BucketCounters.non_current_counters:type_name -> nebius.storage.v1.NonCurrentBucketCounters
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_nebius_storage_v1_bucket_counters_proto_init() }
func file_nebius_storage_v1_bucket_counters_proto_init() {
	if File_nebius_storage_v1_bucket_counters_proto != nil {
		return
	}
	file_nebius_storage_v1_base_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_storage_v1_bucket_counters_proto_rawDesc), len(file_nebius_storage_v1_bucket_counters_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_storage_v1_bucket_counters_proto_goTypes,
		DependencyIndexes: file_nebius_storage_v1_bucket_counters_proto_depIdxs,
		MessageInfos:      file_nebius_storage_v1_bucket_counters_proto_msgTypes,
	}.Build()
	File_nebius_storage_v1_bucket_counters_proto = out.File
	file_nebius_storage_v1_bucket_counters_proto_goTypes = nil
	file_nebius_storage_v1_bucket_counters_proto_depIdxs = nil
}
