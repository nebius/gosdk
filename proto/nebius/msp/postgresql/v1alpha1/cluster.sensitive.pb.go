// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [Cluster] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *Cluster) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [Cluster].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *Cluster
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [Cluster], use the following code:
//
//	var original *Cluster
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*Cluster)
func (x *Cluster) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*Cluster) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperCluster)(c))
}

// wrapperCluster is used to return [Cluster] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperCluster Cluster

func (w *wrapperCluster) String() string {
	return (*Cluster)(w).String()
}

func (*wrapperCluster) ProtoMessage() {}

func (w *wrapperCluster) ProtoReflect() protoreflect.Message {
	return (*Cluster)(w).ProtoReflect()
}

// func (x *ConnectionPoolerConfig) Sanitize()            // is not generated as no sensitive fields found
// func (x *ConnectionPoolerConfig) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [ClusterSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ClusterSpec) Sanitize() {
	if x == nil {
		return
	}
	x.Bootstrap.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ClusterSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ClusterSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ClusterSpec], use the following code:
//
//	var original *ClusterSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ClusterSpec)
func (x *ClusterSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ClusterSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperClusterSpec)(c))
}

// wrapperClusterSpec is used to return [ClusterSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperClusterSpec ClusterSpec

func (w *wrapperClusterSpec) String() string {
	return (*ClusterSpec)(w).String()
}

func (*wrapperClusterSpec) ProtoMessage() {}

func (w *wrapperClusterSpec) ProtoReflect() protoreflect.Message {
	return (*ClusterSpec)(w).ProtoReflect()
}

// func (x *ClusterStatus) Sanitize()            // is not generated as no sensitive fields found
// func (x *ClusterStatus) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *EndpointsSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *EndpointsSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *Endpoints) Sanitize()            // is not generated as no sensitive fields found
// func (x *Endpoints) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ConfigSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *ConfigSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [BootstrapSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *BootstrapSpec) Sanitize() {
	if x == nil {
		return
	}
	x.UserPassword = "**HIDDEN**"
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [BootstrapSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *BootstrapSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [BootstrapSpec], use the following code:
//
//	var original *BootstrapSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*BootstrapSpec)
func (x *BootstrapSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*BootstrapSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperBootstrapSpec)(c))
}

// wrapperBootstrapSpec is used to return [BootstrapSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperBootstrapSpec BootstrapSpec

func (w *wrapperBootstrapSpec) String() string {
	return (*BootstrapSpec)(w).String()
}

func (*wrapperBootstrapSpec) ProtoMessage() {}

func (w *wrapperBootstrapSpec) ProtoReflect() protoreflect.Message {
	return (*BootstrapSpec)(w).ProtoReflect()
}

// func (x *BackupSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *BackupSpec) LogValue() slog.Value // is not generated as no sensitive fields found
