// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [Job] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *Job) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [Job].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *Job
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [Job], use the following code:
//
//	var original *Job
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*Job)
func (x *Job) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*Job) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperJob)(c))
}

// wrapperJob is used to return [Job] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperJob Job

func (w *wrapperJob) String() string {
	return (*Job)(w).String()
}

func (*wrapperJob) ProtoMessage() {}

func (w *wrapperJob) ProtoReflect() protoreflect.Message {
	return (*Job)(w).ProtoReflect()
}

// Sanitize mutates [JobSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *JobSpec) Sanitize() {
	if x == nil {
		return
	}
	x.Container.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [JobSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *JobSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [JobSpec], use the following code:
//
//	var original *JobSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*JobSpec)
func (x *JobSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*JobSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperJobSpec)(c))
}

// wrapperJobSpec is used to return [JobSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperJobSpec JobSpec

func (w *wrapperJobSpec) String() string {
	return (*JobSpec)(w).String()
}

func (*wrapperJobSpec) ProtoMessage() {}

func (w *wrapperJobSpec) ProtoReflect() protoreflect.Message {
	return (*JobSpec)(w).ProtoReflect()
}

// Sanitize mutates [JobContainerSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *JobContainerSpec) Sanitize() {
	if x == nil {
		return
	}
	x.SensitiveEnvs = map[string]string{"**HIDDEN**": "**HIDDEN**"}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [JobContainerSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *JobContainerSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [JobContainerSpec], use the following code:
//
//	var original *JobContainerSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*JobContainerSpec)
func (x *JobContainerSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*JobContainerSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperJobContainerSpec)(c))
}

// wrapperJobContainerSpec is used to return [JobContainerSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperJobContainerSpec JobContainerSpec

func (w *wrapperJobContainerSpec) String() string {
	return (*JobContainerSpec)(w).String()
}

func (*wrapperJobContainerSpec) ProtoMessage() {}

func (w *wrapperJobContainerSpec) ProtoReflect() protoreflect.Message {
	return (*JobContainerSpec)(w).ProtoReflect()
}

// func (x *JobTemplateSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *JobTemplateSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *JobStatus) Sanitize()            // is not generated as no sensitive fields found
// func (x *JobStatus) LogValue() slog.Value // is not generated as no sensitive fields found
