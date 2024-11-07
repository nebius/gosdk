// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// func (x *GetProfileRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetProfileRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [GetProfileResponse] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *GetProfileResponse) Sanitize() {
	if x == nil {
		return
	}
	if o, ok := x.Profile.(*GetProfileResponse_UserProfile); ok && o != nil {
		o.UserProfile.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [GetProfileResponse].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *GetProfileResponse
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [GetProfileResponse], use the following code:
//
//	var original *GetProfileResponse
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*GetProfileResponse)
func (x *GetProfileResponse) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*GetProfileResponse) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperGetProfileResponse)(c))
}

// wrapperGetProfileResponse is used to return [GetProfileResponse] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperGetProfileResponse GetProfileResponse

func (w *wrapperGetProfileResponse) String() string {
	return (*GetProfileResponse)(w).String()
}

func (*wrapperGetProfileResponse) ProtoMessage() {}

func (w *wrapperGetProfileResponse) ProtoReflect() protoreflect.Message {
	return (*GetProfileResponse)(w).ProtoReflect()
}

// Sanitize mutates [UserProfile] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *UserProfile) Sanitize() {
	if x == nil {
		return
	}
	x.FederationInfo.Sanitize()
	if o, ok := x.AttributesOptional.(*UserProfile_Attributes); ok && o != nil {
		o.Attributes.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [UserProfile].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with `***`, other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *UserProfile
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [UserProfile], use the following code:
//
//	var original *UserProfile
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*UserProfile)
func (x *UserProfile) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*UserProfile) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperUserProfile)(c))
}

// wrapperUserProfile is used to return [UserProfile] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperUserProfile UserProfile

func (w *wrapperUserProfile) String() string {
	return (*UserProfile)(w).String()
}

func (*wrapperUserProfile) ProtoMessage() {}

func (w *wrapperUserProfile) ProtoReflect() protoreflect.Message {
	return (*UserProfile)(w).ProtoReflect()
}

// func (x *UserTenantInfo) Sanitize()            // is not generated as no sensitive fields found
// func (x *UserTenantInfo) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ServiceAccountProfile) Sanitize()            // is not generated as no sensitive fields found
// func (x *ServiceAccountProfile) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *AnonymousAccount) Sanitize()            // is not generated as no sensitive fields found
// func (x *AnonymousAccount) LogValue() slog.Value // is not generated as no sensitive fields found
