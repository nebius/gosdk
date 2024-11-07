// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// func (x *GetClusterRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetClusterRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GetClusterByNameRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetClusterByNameRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ListClustersRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListClustersRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [ListClustersResponse] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ListClustersResponse) Sanitize() {
	if x == nil {
		return
	}
	for _, y := range x.Items {
		y.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ListClustersResponse].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ListClustersResponse
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ListClustersResponse], use the following code:
//
//	var original *ListClustersResponse
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ListClustersResponse)
func (x *ListClustersResponse) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ListClustersResponse) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperListClustersResponse)(c))
}

// wrapperListClustersResponse is used to return [ListClustersResponse] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperListClustersResponse ListClustersResponse

func (w *wrapperListClustersResponse) String() string {
	return (*ListClustersResponse)(w).String()
}

func (*wrapperListClustersResponse) ProtoMessage() {}

func (w *wrapperListClustersResponse) ProtoReflect() protoreflect.Message {
	return (*ListClustersResponse)(w).ProtoReflect()
}

// Sanitize mutates [CreateClusterRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *CreateClusterRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [CreateClusterRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *CreateClusterRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [CreateClusterRequest], use the following code:
//
//	var original *CreateClusterRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*CreateClusterRequest)
func (x *CreateClusterRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*CreateClusterRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperCreateClusterRequest)(c))
}

// wrapperCreateClusterRequest is used to return [CreateClusterRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperCreateClusterRequest CreateClusterRequest

func (w *wrapperCreateClusterRequest) String() string {
	return (*CreateClusterRequest)(w).String()
}

func (*wrapperCreateClusterRequest) ProtoMessage() {}

func (w *wrapperCreateClusterRequest) ProtoReflect() protoreflect.Message {
	return (*CreateClusterRequest)(w).ProtoReflect()
}

// Sanitize mutates [UpdateClusterRequest] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *UpdateClusterRequest) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [UpdateClusterRequest].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *UpdateClusterRequest
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [UpdateClusterRequest], use the following code:
//
//	var original *UpdateClusterRequest
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*UpdateClusterRequest)
func (x *UpdateClusterRequest) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*UpdateClusterRequest) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperUpdateClusterRequest)(c))
}

// wrapperUpdateClusterRequest is used to return [UpdateClusterRequest] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperUpdateClusterRequest UpdateClusterRequest

func (w *wrapperUpdateClusterRequest) String() string {
	return (*UpdateClusterRequest)(w).String()
}

func (*wrapperUpdateClusterRequest) ProtoMessage() {}

func (w *wrapperUpdateClusterRequest) ProtoReflect() protoreflect.Message {
	return (*UpdateClusterRequest)(w).ProtoReflect()
}

// func (x *DeleteClusterRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *DeleteClusterRequest) LogValue() slog.Value // is not generated as no sensitive fields found
