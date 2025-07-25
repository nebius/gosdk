// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [ResourceSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ResourceSpec) Sanitize() {
	if x == nil {
		return
	}
	if o, ok := x.ResourceSpec.(*ResourceSpec_ComputeInstanceSpec); ok && o != nil {
		o.ComputeInstanceSpec.Sanitize()
	}
	if o, ok := x.ResourceSpec.(*ResourceSpec_ComputeInstanceUpdateSpec); ok && o != nil {
		o.ComputeInstanceUpdateSpec.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ResourceSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ResourceSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ResourceSpec], use the following code:
//
//	var original *ResourceSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ResourceSpec)
func (x *ResourceSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ResourceSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperResourceSpec)(c))
}

// wrapperResourceSpec is used to return [ResourceSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperResourceSpec ResourceSpec

func (w *wrapperResourceSpec) String() string {
	return (*ResourceSpec)(w).String()
}

func (*wrapperResourceSpec) ProtoMessage() {}

func (w *wrapperResourceSpec) ProtoReflect() protoreflect.Message {
	return (*ResourceSpec)(w).ProtoReflect()
}

// func (x *ResourceGroupCost) Sanitize()            // is not generated as no sensitive fields found
// func (x *ResourceGroupCost) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GeneralTotalCost) Sanitize()            // is not generated as no sensitive fields found
// func (x *GeneralTotalCost) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GeneralResourceCost) Sanitize()            // is not generated as no sensitive fields found
// func (x *GeneralResourceCost) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *CostBreakdown) Sanitize()            // is not generated as no sensitive fields found
// func (x *CostBreakdown) LogValue() slog.Value // is not generated as no sensitive fields found
