// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [NodeGroup] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *NodeGroup) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [NodeGroup].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *NodeGroup
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [NodeGroup], use the following code:
//
//	var original *NodeGroup
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*NodeGroup)
func (x *NodeGroup) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*NodeGroup) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperNodeGroup)(c))
}

// wrapperNodeGroup is used to return [NodeGroup] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperNodeGroup NodeGroup

func (w *wrapperNodeGroup) String() string {
	return (*NodeGroup)(w).String()
}

func (*wrapperNodeGroup) ProtoMessage() {}

func (w *wrapperNodeGroup) ProtoReflect() protoreflect.Message {
	return (*NodeGroup)(w).ProtoReflect()
}

// Sanitize mutates [NodeGroupSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *NodeGroupSpec) Sanitize() {
	if x == nil {
		return
	}
	x.Template.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [NodeGroupSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *NodeGroupSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [NodeGroupSpec], use the following code:
//
//	var original *NodeGroupSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*NodeGroupSpec)
func (x *NodeGroupSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*NodeGroupSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperNodeGroupSpec)(c))
}

// wrapperNodeGroupSpec is used to return [NodeGroupSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperNodeGroupSpec NodeGroupSpec

func (w *wrapperNodeGroupSpec) String() string {
	return (*NodeGroupSpec)(w).String()
}

func (*wrapperNodeGroupSpec) ProtoMessage() {}

func (w *wrapperNodeGroupSpec) ProtoReflect() protoreflect.Message {
	return (*NodeGroupSpec)(w).ProtoReflect()
}

// Sanitize mutates [NodeTemplate] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *NodeTemplate) Sanitize() {
	if x == nil {
		return
	}
	x.CloudInitUserData = "**HIDDEN**"
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [NodeTemplate].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *NodeTemplate
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [NodeTemplate], use the following code:
//
//	var original *NodeTemplate
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*NodeTemplate)
func (x *NodeTemplate) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*NodeTemplate) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperNodeTemplate)(c))
}

// wrapperNodeTemplate is used to return [NodeTemplate] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperNodeTemplate NodeTemplate

func (w *wrapperNodeTemplate) String() string {
	return (*NodeTemplate)(w).String()
}

func (*wrapperNodeTemplate) ProtoMessage() {}

func (w *wrapperNodeTemplate) ProtoReflect() protoreflect.Message {
	return (*NodeTemplate)(w).ProtoReflect()
}

// func (x *NodeMetadataTemplate) Sanitize()            // is not generated as no sensitive fields found
// func (x *NodeMetadataTemplate) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GpuSettings) Sanitize()            // is not generated as no sensitive fields found
// func (x *GpuSettings) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GpuClusterSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *GpuClusterSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *NetworkInterfaceTemplate) Sanitize()            // is not generated as no sensitive fields found
// func (x *NetworkInterfaceTemplate) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *PublicIPAddress) Sanitize()            // is not generated as no sensitive fields found
// func (x *PublicIPAddress) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *AttachedFilesystemSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *AttachedFilesystemSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ExistingFilesystem) Sanitize()            // is not generated as no sensitive fields found
// func (x *ExistingFilesystem) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *NodeGroupAutoscalingSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *NodeGroupAutoscalingSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *PreemptibleSpec) Sanitize()            // is not generated as no sensitive fields found
// func (x *PreemptibleSpec) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *NodeTaint) Sanitize()            // is not generated as no sensitive fields found
// func (x *NodeTaint) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *NodeGroupDeploymentStrategy) Sanitize()            // is not generated as no sensitive fields found
// func (x *NodeGroupDeploymentStrategy) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *PercentOrCount) Sanitize()            // is not generated as no sensitive fields found
// func (x *PercentOrCount) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *NodeGroupStatus) Sanitize()            // is not generated as no sensitive fields found
// func (x *NodeGroupStatus) LogValue() slog.Value // is not generated as no sensitive fields found
