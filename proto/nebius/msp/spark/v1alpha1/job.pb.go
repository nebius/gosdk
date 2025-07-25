// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: nebius/msp/spark/v1alpha1/job.proto

package v1alpha1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nebius/gosdk/proto/nebius"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/msp/v1alpha1"
	resource "github.com/nebius/gosdk/proto/nebius/msp/v1alpha1/resource"
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

type JobResultCode int32

const (
	JobResultCode_JOB_RESULT_CODE_UNSPECIFIED JobResultCode = 0
	JobResultCode_SUCCEEDED                   JobResultCode = 1
	JobResultCode_ERROR                       JobResultCode = 2
)

// Enum value maps for JobResultCode.
var (
	JobResultCode_name = map[int32]string{
		0: "JOB_RESULT_CODE_UNSPECIFIED",
		1: "SUCCEEDED",
		2: "ERROR",
	}
	JobResultCode_value = map[string]int32{
		"JOB_RESULT_CODE_UNSPECIFIED": 0,
		"SUCCEEDED":                   1,
		"ERROR":                       2,
	}
)

func (x JobResultCode) Enum() *JobResultCode {
	p := new(JobResultCode)
	*p = x
	return p
}

func (x JobResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_nebius_msp_spark_v1alpha1_job_proto_enumTypes[0].Descriptor()
}

func (JobResultCode) Type() protoreflect.EnumType {
	return &file_nebius_msp_spark_v1alpha1_job_proto_enumTypes[0]
}

func (x JobResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobResultCode.Descriptor instead.
func (JobResultCode) EnumDescriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_job_proto_rawDescGZIP(), []int{0}
}

type Job struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Metadata      *v1.ResourceMetadata   `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec          *JobSpec               `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status        *JobStatus             `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Job) Reset() {
	*x = Job{}
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetMetadata() *v1.ResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Job) GetSpec() *JobSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Job) GetStatus() *JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Spark Job specification
type JobSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Description of the job.
	Description *string `protobuf:"bytes,1,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// S3 URI of main application file
	// Example: s3a://mybucket/myapp.py
	ApplicationFileUri string                `protobuf:"bytes,2,opt,name=application_file_uri,json=applicationFileUri,proto3" json:"application_file_uri,omitempty"`
	Driver             *DriverTemplateSpec   `protobuf:"bytes,3,opt,name=driver,proto3" json:"driver,omitempty"`
	Executor           *ExecutorTemplateSpec `protobuf:"bytes,5,opt,name=executor,proto3" json:"executor,omitempty"`
	SparkVersion       string                `protobuf:"bytes,6,opt,name=spark_version,json=sparkVersion,proto3" json:"spark_version,omitempty"`
	// Application args
	ApplicationArgs []string `protobuf:"bytes,101,rep,name=application_args,json=applicationArgs,proto3" json:"application_args,omitempty"`
	// S3 URIs of files to be placed in executor working directory
	FileUris []string `protobuf:"bytes,102,rep,name=file_uris,json=fileUris,proto3" json:"file_uris,omitempty"`
	// S3 URIs of Jars to be placed in classpaths of driver and executors for java applications
	JarUris []string `protobuf:"bytes,103,rep,name=jar_uris,json=jarUris,proto3" json:"jar_uris,omitempty"`
	// List of maven coordinates of jars to include on the driver and executor classpaths
	Packages []string `protobuf:"bytes,104,rep,name=packages,proto3" json:"packages,omitempty"`
	// Map of spark configuration parameters
	SparkConf map[string]string `protobuf:"bytes,105,rep,name=spark_conf,json=sparkConf,proto3" json:"spark_conf,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Runtime-specific job config
	//
	// Types that are valid to be assigned to RuntimeConfig:
	//
	//	*JobSpec_Python
	//	*JobSpec_Java
	RuntimeConfig isJobSpec_RuntimeConfig `protobuf_oneof:"runtime_config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobSpec) Reset() {
	*x = JobSpec{}
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSpec) ProtoMessage() {}

func (x *JobSpec) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSpec.ProtoReflect.Descriptor instead.
func (*JobSpec) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobSpec) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *JobSpec) GetApplicationFileUri() string {
	if x != nil {
		return x.ApplicationFileUri
	}
	return ""
}

func (x *JobSpec) GetDriver() *DriverTemplateSpec {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *JobSpec) GetExecutor() *ExecutorTemplateSpec {
	if x != nil {
		return x.Executor
	}
	return nil
}

func (x *JobSpec) GetSparkVersion() string {
	if x != nil {
		return x.SparkVersion
	}
	return ""
}

func (x *JobSpec) GetApplicationArgs() []string {
	if x != nil {
		return x.ApplicationArgs
	}
	return nil
}

func (x *JobSpec) GetFileUris() []string {
	if x != nil {
		return x.FileUris
	}
	return nil
}

func (x *JobSpec) GetJarUris() []string {
	if x != nil {
		return x.JarUris
	}
	return nil
}

func (x *JobSpec) GetPackages() []string {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *JobSpec) GetSparkConf() map[string]string {
	if x != nil {
		return x.SparkConf
	}
	return nil
}

func (x *JobSpec) GetRuntimeConfig() isJobSpec_RuntimeConfig {
	if x != nil {
		return x.RuntimeConfig
	}
	return nil
}

func (x *JobSpec) GetPython() *PythonConfig {
	if x != nil {
		if x, ok := x.RuntimeConfig.(*JobSpec_Python); ok {
			return x.Python
		}
	}
	return nil
}

func (x *JobSpec) GetJava() *JavaConfig {
	if x != nil {
		if x, ok := x.RuntimeConfig.(*JobSpec_Java); ok {
			return x.Java
		}
	}
	return nil
}

type isJobSpec_RuntimeConfig interface {
	isJobSpec_RuntimeConfig()
}

type JobSpec_Python struct {
	Python *PythonConfig `protobuf:"bytes,201,opt,name=python,proto3,oneof"`
}

type JobSpec_Java struct {
	Java *JavaConfig `protobuf:"bytes,202,opt,name=java,proto3,oneof"`
}

func (*JobSpec_Python) isJobSpec_RuntimeConfig() {}

func (*JobSpec_Java) isJobSpec_RuntimeConfig() {}

type JobResultDetails struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Result code
	Code          JobResultCode `protobuf:"varint,1,opt,name=code,proto3,enum=nebius.msp.spark.v1alpha1.JobResultCode" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobResultDetails) Reset() {
	*x = JobResultDetails{}
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobResultDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobResultDetails) ProtoMessage() {}

func (x *JobResultDetails) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobResultDetails.ProtoReflect.Descriptor instead.
func (*JobResultDetails) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_job_proto_rawDescGZIP(), []int{2}
}

func (x *JobResultDetails) GetCode() JobResultCode {
	if x != nil {
		return x.Code
	}
	return JobResultCode_JOB_RESULT_CODE_UNSPECIFIED
}

type JobStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current phase (or stage) of the cluster.
	Phase v1alpha1.ClusterStatus_Phase `protobuf:"varint,1,opt,name=phase,proto3,enum=nebius.msp.v1alpha1.ClusterStatus_Phase" json:"phase,omitempty"`
	// State reflects substatus of the stage to define whether it's healthy or not.
	State v1alpha1.ClusterStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=nebius.msp.v1alpha1.ClusterStatus_State" json:"state,omitempty"`
	// Job Driver Web UI endpoint
	DriverEndpoint *string `protobuf:"bytes,3,opt,name=driver_endpoint,json=driverEndpoint,proto3,oneof" json:"driver_endpoint,omitempty"`
	// Job driver resource preset details
	DriverPresetDetails *resource.PresetDetails `protobuf:"bytes,4,opt,name=driver_preset_details,json=driverPresetDetails,proto3" json:"driver_preset_details,omitempty"`
	// Job executor resource preset details
	ExecutorPresetDetails *resource.PresetDetails `protobuf:"bytes,5,opt,name=executor_preset_details,json=executorPresetDetails,proto3" json:"executor_preset_details,omitempty"`
	// Job execution result details
	ResultDetails *JobResultDetails `protobuf:"bytes,6,opt,name=result_details,json=resultDetails,proto3" json:"result_details,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JobStatus) Reset() {
	*x = JobStatus{}
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatus) ProtoMessage() {}

func (x *JobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatus.ProtoReflect.Descriptor instead.
func (*JobStatus) Descriptor() ([]byte, []int) {
	return file_nebius_msp_spark_v1alpha1_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobStatus) GetPhase() v1alpha1.ClusterStatus_Phase {
	if x != nil {
		return x.Phase
	}
	return v1alpha1.ClusterStatus_Phase(0)
}

func (x *JobStatus) GetState() v1alpha1.ClusterStatus_State {
	if x != nil {
		return x.State
	}
	return v1alpha1.ClusterStatus_State(0)
}

func (x *JobStatus) GetDriverEndpoint() string {
	if x != nil && x.DriverEndpoint != nil {
		return *x.DriverEndpoint
	}
	return ""
}

func (x *JobStatus) GetDriverPresetDetails() *resource.PresetDetails {
	if x != nil {
		return x.DriverPresetDetails
	}
	return nil
}

func (x *JobStatus) GetExecutorPresetDetails() *resource.PresetDetails {
	if x != nil {
		return x.ExecutorPresetDetails
	}
	return nil
}

func (x *JobStatus) GetResultDetails() *JobResultDetails {
	if x != nil {
		return x.ResultDetails
	}
	return nil
}

var File_nebius_msp_spark_v1alpha1_job_proto protoreflect.FileDescriptor

var file_nebius_msp_spark_v1alpha1_job_proto_rawDesc = string([]byte{
	0x0a, 0x23, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73,
	0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73,
	0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62,
	0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75,
	0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04,
	0xba, 0x4a, 0x01, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xee, 0x05, 0x0a,
	0x07, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x38, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x69, 0x12, 0x4d, 0x0a, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x0a,
	0x0d, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x65,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72,
	0x69, 0x73, 0x18, 0x66, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x69, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x69, 0x73, 0x18, 0x67,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x61, 0x72, 0x55, 0x72, 0x69, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x68, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x0a, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x69, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65,
	0x63, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x42, 0x0a, 0x06, 0x70,
	0x79, 0x74, 0x68, 0x6f, 0x6e, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x12,
	0x3c, 0x0a, 0x04, 0x6a, 0x61, 0x76, 0x61, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x61, 0x76, 0x61, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x61, 0x76, 0x61, 0x1a, 0x3c, 0x0a,
	0x0e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a,
	0x10, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x3c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0xe7, 0x03, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a,
	0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6e,
	0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a,
	0x0f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x5f, 0x0a, 0x15, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x62,
	0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x13, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x63, 0x0a, 0x17,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x15, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x52, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2a, 0x4a, 0x0a, 0x0d, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x4a, 0x4f,
	0x42, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x67, 0x0a, 0x20, 0x61, 0x69, 0x2e, 0x6e, 0x65, 0x62, 0x69,
	0x75, 0x73, 0x2e, 0x70, 0x75, 0x62, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x4a, 0x6f, 0x62, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x65, 0x62, 0x69, 0x75, 0x73, 0x2f, 0x6d, 0x73, 0x70, 0x2f,
	0x73, 0x70, 0x61, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_nebius_msp_spark_v1alpha1_job_proto_rawDescOnce sync.Once
	file_nebius_msp_spark_v1alpha1_job_proto_rawDescData []byte
)

func file_nebius_msp_spark_v1alpha1_job_proto_rawDescGZIP() []byte {
	file_nebius_msp_spark_v1alpha1_job_proto_rawDescOnce.Do(func() {
		file_nebius_msp_spark_v1alpha1_job_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_nebius_msp_spark_v1alpha1_job_proto_rawDesc), len(file_nebius_msp_spark_v1alpha1_job_proto_rawDesc)))
	})
	return file_nebius_msp_spark_v1alpha1_job_proto_rawDescData
}

var file_nebius_msp_spark_v1alpha1_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nebius_msp_spark_v1alpha1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nebius_msp_spark_v1alpha1_job_proto_goTypes = []any{
	(JobResultCode)(0),                // 0: nebius.msp.spark.v1alpha1.JobResultCode
	(*Job)(nil),                       // 1: nebius.msp.spark.v1alpha1.Job
	(*JobSpec)(nil),                   // 2: nebius.msp.spark.v1alpha1.JobSpec
	(*JobResultDetails)(nil),          // 3: nebius.msp.spark.v1alpha1.JobResultDetails
	(*JobStatus)(nil),                 // 4: nebius.msp.spark.v1alpha1.JobStatus
	nil,                               // 5: nebius.msp.spark.v1alpha1.JobSpec.SparkConfEntry
	(*v1.ResourceMetadata)(nil),       // 6: nebius.common.v1.ResourceMetadata
	(*DriverTemplateSpec)(nil),        // 7: nebius.msp.spark.v1alpha1.DriverTemplateSpec
	(*ExecutorTemplateSpec)(nil),      // 8: nebius.msp.spark.v1alpha1.ExecutorTemplateSpec
	(*PythonConfig)(nil),              // 9: nebius.msp.spark.v1alpha1.PythonConfig
	(*JavaConfig)(nil),                // 10: nebius.msp.spark.v1alpha1.JavaConfig
	(v1alpha1.ClusterStatus_Phase)(0), // 11: nebius.msp.v1alpha1.ClusterStatus.Phase
	(v1alpha1.ClusterStatus_State)(0), // 12: nebius.msp.v1alpha1.ClusterStatus.State
	(*resource.PresetDetails)(nil),    // 13: nebius.msp.v1alpha1.resource.PresetDetails
}
var file_nebius_msp_spark_v1alpha1_job_proto_depIdxs = []int32{
	6,  // 0: nebius.msp.spark.v1alpha1.Job.metadata:type_name -> nebius.common.v1.ResourceMetadata
	2,  // 1: nebius.msp.spark.v1alpha1.Job.spec:type_name -> nebius.msp.spark.v1alpha1.JobSpec
	4,  // 2: nebius.msp.spark.v1alpha1.Job.status:type_name -> nebius.msp.spark.v1alpha1.JobStatus
	7,  // 3: nebius.msp.spark.v1alpha1.JobSpec.driver:type_name -> nebius.msp.spark.v1alpha1.DriverTemplateSpec
	8,  // 4: nebius.msp.spark.v1alpha1.JobSpec.executor:type_name -> nebius.msp.spark.v1alpha1.ExecutorTemplateSpec
	5,  // 5: nebius.msp.spark.v1alpha1.JobSpec.spark_conf:type_name -> nebius.msp.spark.v1alpha1.JobSpec.SparkConfEntry
	9,  // 6: nebius.msp.spark.v1alpha1.JobSpec.python:type_name -> nebius.msp.spark.v1alpha1.PythonConfig
	10, // 7: nebius.msp.spark.v1alpha1.JobSpec.java:type_name -> nebius.msp.spark.v1alpha1.JavaConfig
	0,  // 8: nebius.msp.spark.v1alpha1.JobResultDetails.code:type_name -> nebius.msp.spark.v1alpha1.JobResultCode
	11, // 9: nebius.msp.spark.v1alpha1.JobStatus.phase:type_name -> nebius.msp.v1alpha1.ClusterStatus.Phase
	12, // 10: nebius.msp.spark.v1alpha1.JobStatus.state:type_name -> nebius.msp.v1alpha1.ClusterStatus.State
	13, // 11: nebius.msp.spark.v1alpha1.JobStatus.driver_preset_details:type_name -> nebius.msp.v1alpha1.resource.PresetDetails
	13, // 12: nebius.msp.spark.v1alpha1.JobStatus.executor_preset_details:type_name -> nebius.msp.v1alpha1.resource.PresetDetails
	3,  // 13: nebius.msp.spark.v1alpha1.JobStatus.result_details:type_name -> nebius.msp.spark.v1alpha1.JobResultDetails
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_nebius_msp_spark_v1alpha1_job_proto_init() }
func file_nebius_msp_spark_v1alpha1_job_proto_init() {
	if File_nebius_msp_spark_v1alpha1_job_proto != nil {
		return
	}
	file_nebius_msp_spark_v1alpha1_common_proto_init()
	file_nebius_msp_spark_v1alpha1_preset_proto_init()
	file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[1].OneofWrappers = []any{
		(*JobSpec_Python)(nil),
		(*JobSpec_Java)(nil),
	}
	file_nebius_msp_spark_v1alpha1_job_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_nebius_msp_spark_v1alpha1_job_proto_rawDesc), len(file_nebius_msp_spark_v1alpha1_job_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nebius_msp_spark_v1alpha1_job_proto_goTypes,
		DependencyIndexes: file_nebius_msp_spark_v1alpha1_job_proto_depIdxs,
		EnumInfos:         file_nebius_msp_spark_v1alpha1_job_proto_enumTypes,
		MessageInfos:      file_nebius_msp_spark_v1alpha1_job_proto_msgTypes,
	}.Build()
	File_nebius_msp_spark_v1alpha1_job_proto = out.File
	file_nebius_msp_spark_v1alpha1_job_proto_goTypes = nil
	file_nebius_msp_spark_v1alpha1_job_proto_depIdxs = nil
}
