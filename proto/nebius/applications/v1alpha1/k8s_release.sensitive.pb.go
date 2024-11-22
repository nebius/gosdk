// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [K8SRelease] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *K8SRelease) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [K8SRelease].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *K8SRelease
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [K8SRelease], use the following code:
//
//	var original *K8SRelease
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*K8SRelease)
func (x *K8SRelease) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*K8SRelease) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperK8SRelease)(c))
}

// wrapperK8SRelease is used to return [K8SRelease] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperK8SRelease K8SRelease

func (w *wrapperK8SRelease) String() string {
	return (*K8SRelease)(w).String()
}

func (*wrapperK8SRelease) ProtoMessage() {}

func (w *wrapperK8SRelease) ProtoReflect() protoreflect.Message {
	return (*K8SRelease)(w).ProtoReflect()
}

// Sanitize mutates [K8SReleaseSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *K8SReleaseSpec) Sanitize() {
	if x == nil {
		return
	}
	x.Values = "**HIDDEN**"
	x.Set = map[string]string{"**HIDDEN**": "**HIDDEN**"}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [K8SReleaseSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *K8SReleaseSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [K8SReleaseSpec], use the following code:
//
//	var original *K8SReleaseSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*K8SReleaseSpec)
func (x *K8SReleaseSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*K8SReleaseSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperK8SReleaseSpec)(c))
}

// wrapperK8SReleaseSpec is used to return [K8SReleaseSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperK8SReleaseSpec K8SReleaseSpec

func (w *wrapperK8SReleaseSpec) String() string {
	return (*K8SReleaseSpec)(w).String()
}

func (*wrapperK8SReleaseSpec) ProtoMessage() {}

func (w *wrapperK8SReleaseSpec) ProtoReflect() protoreflect.Message {
	return (*K8SReleaseSpec)(w).ProtoReflect()
}

// func (x *K8SReleaseStatus) Sanitize()            // is not generated as no sensitive fields found
// func (x *K8SReleaseStatus) LogValue() slog.Value // is not generated as no sensitive fields found
