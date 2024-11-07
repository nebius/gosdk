// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [CreateNodeGroupRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *CreateNodeGroupRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [CreateNodeGroupRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *CreateNodeGroupRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [CreateNodeGroupRequest], use the following code:
//
//	var original *CreateNodeGroupRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*CreateNodeGroupRequest)
func (x *CreateNodeGroupRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*CreateNodeGroupRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperCreateNodeGroupRequest)(c))
}

// wrapperCreateNodeGroupRequest is used to return [CreateNodeGroupRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperCreateNodeGroupRequest CreateNodeGroupRequest

func (w *wrapperCreateNodeGroupRequest) String() string {
	return (*CreateNodeGroupRequest)(w).String()
}

func (*wrapperCreateNodeGroupRequest) ProtoMessage() {}

func (w *wrapperCreateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	return (*CreateNodeGroupRequest)(w).ProtoReflect()
}

// func (x *GetNodeGroupRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetNodeGroupRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GetNodeGroupByNameRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetNodeGroupByNameRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ListNodeGroupsRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListNodeGroupsRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [ListNodeGroupsResponse] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ListNodeGroupsResponse) Sanitize() {
	if x == nil {
		return
	}
	for _, y := range x.Items {
		y.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ListNodeGroupsResponse].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ListNodeGroupsResponse
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ListNodeGroupsResponse], use the following code:
//
//	var original *ListNodeGroupsResponse
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ListNodeGroupsResponse)
func (x *ListNodeGroupsResponse) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ListNodeGroupsResponse) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperListNodeGroupsResponse)(c))
}

// wrapperListNodeGroupsResponse is used to return [ListNodeGroupsResponse] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperListNodeGroupsResponse ListNodeGroupsResponse

func (w *wrapperListNodeGroupsResponse) String() string {
	return (*ListNodeGroupsResponse)(w).String()
}

func (*wrapperListNodeGroupsResponse) ProtoMessage() {}

func (w *wrapperListNodeGroupsResponse) ProtoReflect() protoreflect.Message {
	return (*ListNodeGroupsResponse)(w).ProtoReflect()
}

// Sanitize mutates [UpdateNodeGroupRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *UpdateNodeGroupRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [UpdateNodeGroupRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *UpdateNodeGroupRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [UpdateNodeGroupRequest], use the following code:
//
//	var original *UpdateNodeGroupRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*UpdateNodeGroupRequest)
func (x *UpdateNodeGroupRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*UpdateNodeGroupRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperUpdateNodeGroupRequest)(c))
}

// wrapperUpdateNodeGroupRequest is used to return [UpdateNodeGroupRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperUpdateNodeGroupRequest UpdateNodeGroupRequest

func (w *wrapperUpdateNodeGroupRequest) String() string {
	return (*UpdateNodeGroupRequest)(w).String()
}

func (*wrapperUpdateNodeGroupRequest) ProtoMessage() {}

func (w *wrapperUpdateNodeGroupRequest) ProtoReflect() protoreflect.Message {
	return (*UpdateNodeGroupRequest)(w).ProtoReflect()
}

// func (x *DeleteNodeGroupRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *DeleteNodeGroupRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *UpgradeNodeGroupRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *UpgradeNodeGroupRequest) LogValue() slog.Value // is not generated as no sensitive fields found
