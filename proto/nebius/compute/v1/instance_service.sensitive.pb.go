// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// func (x *GetInstanceRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetInstanceRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ListInstancesRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListInstancesRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [CreateInstanceRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *CreateInstanceRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [CreateInstanceRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *CreateInstanceRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [CreateInstanceRequest], use the following code:
//
//	var original *CreateInstanceRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*CreateInstanceRequest)
func (x *CreateInstanceRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*CreateInstanceRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperCreateInstanceRequest)(c))
}

// wrapperCreateInstanceRequest is used to return [CreateInstanceRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperCreateInstanceRequest CreateInstanceRequest

func (w *wrapperCreateInstanceRequest) String() string {
	return (*CreateInstanceRequest)(w).String()
}

func (*wrapperCreateInstanceRequest) ProtoMessage() {}

func (w *wrapperCreateInstanceRequest) ProtoReflect() protoreflect.Message {
	return (*CreateInstanceRequest)(w).ProtoReflect()
}

// Sanitize mutates [UpdateInstanceRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *UpdateInstanceRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [UpdateInstanceRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *UpdateInstanceRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [UpdateInstanceRequest], use the following code:
//
//	var original *UpdateInstanceRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*UpdateInstanceRequest)
func (x *UpdateInstanceRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*UpdateInstanceRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperUpdateInstanceRequest)(c))
}

// wrapperUpdateInstanceRequest is used to return [UpdateInstanceRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperUpdateInstanceRequest UpdateInstanceRequest

func (w *wrapperUpdateInstanceRequest) String() string {
	return (*UpdateInstanceRequest)(w).String()
}

func (*wrapperUpdateInstanceRequest) ProtoMessage() {}

func (w *wrapperUpdateInstanceRequest) ProtoReflect() protoreflect.Message {
	return (*UpdateInstanceRequest)(w).ProtoReflect()
}

// func (x *DeleteInstanceRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *DeleteInstanceRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [ListInstancesResponse] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ListInstancesResponse) Sanitize() {
	if x == nil {
		return
	}
	for _, y := range x.Items {
		y.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ListInstancesResponse].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ListInstancesResponse
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ListInstancesResponse], use the following code:
//
//	var original *ListInstancesResponse
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ListInstancesResponse)
func (x *ListInstancesResponse) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ListInstancesResponse) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperListInstancesResponse)(c))
}

// wrapperListInstancesResponse is used to return [ListInstancesResponse] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperListInstancesResponse ListInstancesResponse

func (w *wrapperListInstancesResponse) String() string {
	return (*ListInstancesResponse)(w).String()
}

func (*wrapperListInstancesResponse) ProtoMessage() {}

func (w *wrapperListInstancesResponse) ProtoReflect() protoreflect.Message {
	return (*ListInstancesResponse)(w).ProtoReflect()
}

// func (x *StartInstanceRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *StartInstanceRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *StopInstanceRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *StopInstanceRequest) LogValue() slog.Value // is not generated as no sensitive fields found
