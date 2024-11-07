// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// Sanitize mutates [TenantUserAccount] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *TenantUserAccount) Sanitize() {
	if x == nil {
		return
	}
	x.Spec.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [TenantUserAccount].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *TenantUserAccount
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [TenantUserAccount], use the following code:
//
//	var original *TenantUserAccount
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*TenantUserAccount)
func (x *TenantUserAccount) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*TenantUserAccount) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperTenantUserAccount)(c))
}

// wrapperTenantUserAccount is used to return [TenantUserAccount] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperTenantUserAccount TenantUserAccount

func (w *wrapperTenantUserAccount) String() string {
	return (*TenantUserAccount)(w).String()
}

func (*wrapperTenantUserAccount) ProtoMessage() {}

func (w *wrapperTenantUserAccount) ProtoReflect() protoreflect.Message {
	return (*TenantUserAccount)(w).ProtoReflect()
}

// Sanitize mutates [TenantUserAccountWithAttributes] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *TenantUserAccountWithAttributes) Sanitize() {
	if x == nil {
		return
	}
	x.TenantUserAccount.Sanitize()
	if o, ok := x.AttributesOptional.(*TenantUserAccountWithAttributes_Attributes); ok && o != nil {
		o.Attributes.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [TenantUserAccountWithAttributes].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *TenantUserAccountWithAttributes
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [TenantUserAccountWithAttributes], use the following code:
//
//	var original *TenantUserAccountWithAttributes
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*TenantUserAccountWithAttributes)
func (x *TenantUserAccountWithAttributes) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*TenantUserAccountWithAttributes) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperTenantUserAccountWithAttributes)(c))
}

// wrapperTenantUserAccountWithAttributes is used to return [TenantUserAccountWithAttributes] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperTenantUserAccountWithAttributes TenantUserAccountWithAttributes

func (w *wrapperTenantUserAccountWithAttributes) String() string {
	return (*TenantUserAccountWithAttributes)(w).String()
}

func (*wrapperTenantUserAccountWithAttributes) ProtoMessage() {}

func (w *wrapperTenantUserAccountWithAttributes) ProtoReflect() protoreflect.Message {
	return (*TenantUserAccountWithAttributes)(w).ProtoReflect()
}

// Sanitize mutates [UserAttributes] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *UserAttributes) Sanitize() {
	if x == nil {
		return
	}
	x.Sub = proto.String("***")
	x.Name = proto.String("***")
	x.GivenName = proto.String("***")
	x.FamilyName = proto.String("***")
	x.PreferredUsername = proto.String("***")
	x.Picture = proto.String("***")
	x.Email = proto.String("***")
	x.Zoneinfo = proto.String("***")
	x.Locale = proto.String("***")
	x.PhoneNumber = proto.String("***")
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [UserAttributes].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *UserAttributes
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [UserAttributes], use the following code:
//
//	var original *UserAttributes
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*UserAttributes)
func (x *UserAttributes) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*UserAttributes) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperUserAttributes)(c))
}

// wrapperUserAttributes is used to return [UserAttributes] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperUserAttributes UserAttributes

func (w *wrapperUserAttributes) String() string {
	return (*UserAttributes)(w).String()
}

func (*wrapperUserAttributes) ProtoMessage() {}

func (w *wrapperUserAttributes) ProtoReflect() protoreflect.Message {
	return (*UserAttributes)(w).ProtoReflect()
}

// func (x *Error) Sanitize()            // is not generated as no sensitive fields found
// func (x *Error) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [TenantUserAccountSpec] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *TenantUserAccountSpec) Sanitize() {
	if x == nil {
		return
	}
	x.VisibleAttributes.Sanitize()
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [TenantUserAccountSpec].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *TenantUserAccountSpec
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [TenantUserAccountSpec], use the following code:
//
//	var original *TenantUserAccountSpec
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*TenantUserAccountSpec)
func (x *TenantUserAccountSpec) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*TenantUserAccountSpec) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperTenantUserAccountSpec)(c))
}

// wrapperTenantUserAccountSpec is used to return [TenantUserAccountSpec] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperTenantUserAccountSpec TenantUserAccountSpec

func (w *wrapperTenantUserAccountSpec) String() string {
	return (*TenantUserAccountSpec)(w).String()
}

func (*wrapperTenantUserAccountSpec) ProtoMessage() {}

func (w *wrapperTenantUserAccountSpec) ProtoReflect() protoreflect.Message {
	return (*TenantUserAccountSpec)(w).ProtoReflect()
}

// Sanitize mutates [TenantUserAccountSpec_VisibleAttributes] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *TenantUserAccountSpec_VisibleAttributes) Sanitize() {
	if x == nil {
		return
	}
	x.Attribute = []string{"***"}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [TenantUserAccountSpec_VisibleAttributes].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *TenantUserAccountSpec_VisibleAttributes
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [TenantUserAccountSpec_VisibleAttributes], use the following code:
//
//	var original *TenantUserAccountSpec_VisibleAttributes
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*TenantUserAccountSpec_VisibleAttributes)
func (x *TenantUserAccountSpec_VisibleAttributes) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*TenantUserAccountSpec_VisibleAttributes) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperTenantUserAccountSpec_VisibleAttributes)(c))
}

// wrapperTenantUserAccountSpec_VisibleAttributes is used to return [TenantUserAccountSpec_VisibleAttributes] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperTenantUserAccountSpec_VisibleAttributes TenantUserAccountSpec_VisibleAttributes

func (w *wrapperTenantUserAccountSpec_VisibleAttributes) String() string {
	return (*TenantUserAccountSpec_VisibleAttributes)(w).String()
}

func (*wrapperTenantUserAccountSpec_VisibleAttributes) ProtoMessage() {}

func (w *wrapperTenantUserAccountSpec_VisibleAttributes) ProtoReflect() protoreflect.Message {
	return (*TenantUserAccountSpec_VisibleAttributes)(w).ProtoReflect()
}

// func (x *TenantUserAccountStatus) Sanitize()            // is not generated as no sensitive fields found
// func (x *TenantUserAccountStatus) LogValue() slog.Value // is not generated as no sensitive fields found
