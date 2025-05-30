// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [CreateJobRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *CreateJobRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [CreateJobRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *CreateJobRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [CreateJobRequest], use the following code:
//
//	var original *CreateJobRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*CreateJobRequest)
func (x *CreateJobRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*CreateJobRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperCreateJobRequest)(c))
}

// wrapperCreateJobRequest is used to return [CreateJobRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperCreateJobRequest CreateJobRequest

func (w *wrapperCreateJobRequest) String() string {
	return (*CreateJobRequest)(w).String()
}

func (*wrapperCreateJobRequest) ProtoMessage() {}

func (w *wrapperCreateJobRequest) ProtoReflect() protoreflect.Message {
	return (*CreateJobRequest)(w).ProtoReflect()
}

// func (x *CancelJobRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *CancelJobRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [ListJobsResponse] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ListJobsResponse) Sanitize() {
	if x == nil {
		return
	}
	for _, y := range x.Items {
		y.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ListJobsResponse].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ListJobsResponse
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ListJobsResponse], use the following code:
//
//	var original *ListJobsResponse
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ListJobsResponse)
func (x *ListJobsResponse) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ListJobsResponse) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperListJobsResponse)(c))
}

// wrapperListJobsResponse is used to return [ListJobsResponse] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperListJobsResponse ListJobsResponse

func (w *wrapperListJobsResponse) String() string {
	return (*ListJobsResponse)(w).String()
}

func (*wrapperListJobsResponse) ProtoMessage() {}

func (w *wrapperListJobsResponse) ProtoReflect() protoreflect.Message {
	return (*ListJobsResponse)(w).ProtoReflect()
}
