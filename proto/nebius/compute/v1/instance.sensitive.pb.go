// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [Instance] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *Instance) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [Instance].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *Instance
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [Instance], use the following code:
//
//	var original *Instance
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*Instance)
func (x *Instance) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*Instance) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperInstance)(c))
}

// wrapperInstance is used to return [Instance] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperInstance Instance

func (w *wrapperInstance) String() string {
	return (*Instance)(w).String()
}

func (*wrapperInstance) ProtoMessage() {}

func (w *wrapperInstance) ProtoReflect() protoreflect.Message {
	return (*Instance)(w).ProtoReflect()
}

// Sanitize mutates [InstanceSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *InstanceSpec) Sanitize() {
	if x == nil {
		return
	}
	x.CloudInitUserData = "**HIDDEN**"
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [InstanceSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *InstanceSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [InstanceSpec], use the following code:
//
//	var original *InstanceSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*InstanceSpec)
func (x *InstanceSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*InstanceSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperInstanceSpec)(c))
}

// wrapperInstanceSpec is used to return [InstanceSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperInstanceSpec InstanceSpec

func (w *wrapperInstanceSpec) String() string {
	return (*InstanceSpec)(w).String()
}

func (*wrapperInstanceSpec) ProtoMessage() {}

func (w *wrapperInstanceSpec) ProtoReflect() protoreflect.Message {
	return (*InstanceSpec)(w).ProtoReflect()
}

// func (x *ResourcesSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *ResourcesSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *InstanceGpuClusterSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *InstanceGpuClusterSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *AttachedDiskSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *AttachedDiskSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ExistingDisk) Sanitize()            // is not generated as no sensitive fields found
// func (x *ExistingDisk) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ExistingFilesystem) Sanitize()            // is not generated as no sensitive fields found
// func (x *ExistingFilesystem) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *AttachedFilesystemSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *AttachedFilesystemSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *InstanceStatus) Sanitize()            // is not generated as no sensitive fields found
// func (x *InstanceStatus) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *InstanceStatusInfinibandTopologyPath) Sanitize()            // is not generated as no sensitive fields found
// func (x *InstanceStatusInfinibandTopologyPath) LogValue() slog.Value // is not generated as no sensitive fields found
