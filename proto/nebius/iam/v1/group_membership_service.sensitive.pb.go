// Code generated by protoc-gen-sensitive. DO NOT EDIT.

package v1

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	slog "log/slog"
)

// func (x *CreateGroupMembershipRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *CreateGroupMembershipRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *DeleteGroupMembershipRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *DeleteGroupMembershipRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *GetGroupMembershipRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *GetGroupMembershipRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ListGroupMembershipsRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListGroupMembershipsRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ListMemberOfRequest) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListMemberOfRequest) LogValue() slog.Value // is not generated as no sensitive fields found

// func (x *ListGroupMembershipsResponse) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListGroupMembershipsResponse) LogValue() slog.Value // is not generated as no sensitive fields found

// Sanitize mutates [ListGroupMembershipsWithAttributesResponse] to remove/mask all sensitive values.
// Sensitive fields are marked with [(nebius.sensitive) = true].
func (x *ListGroupMembershipsWithAttributesResponse) Sanitize() {
	if x == nil {
		return
	}
	for _, y := range x.Memberships {
		y.Sanitize()
	}
}

// LogValue implements [slog.LogValuer] interface. It returns sanitized copy of [ListGroupMembershipsWithAttributesResponse].
// Properly implemented [slog.Handler] must call LogValue, so sensitive values are not logged.
// Sensitive strings and bytes are masked with "**HIDDEN**", other sensitive fields are omitted.
//
// Returning value has kind [slog.KindAny]. To extract [proto.Message], use the following code:
//
//	var original *ListGroupMembershipsWithAttributesResponse
//	sanitized := original.LogValue().Any().(proto.Message)
//
// If you need to extract [ListGroupMembershipsWithAttributesResponse], use the following code:
//
//	var original *ListGroupMembershipsWithAttributesResponse
//	sanitized := original.LogValue().Any().(proto.Message).ProtoReflect().Interface().(*ListGroupMembershipsWithAttributesResponse)
func (x *ListGroupMembershipsWithAttributesResponse) LogValue() slog.Value {
	if x == nil {
		return slog.AnyValue(x)
	}
	c := proto.Clone(x).(*ListGroupMembershipsWithAttributesResponse) // TODO: generate static cloner without protoreflect
	c.Sanitize()
	return slog.AnyValue((*wrapperListGroupMembershipsWithAttributesResponse)(c))
}

// wrapperListGroupMembershipsWithAttributesResponse is used to return [ListGroupMembershipsWithAttributesResponse] not implementing [slog.LogValuer] to avoid recursion while resolving.
type wrapperListGroupMembershipsWithAttributesResponse ListGroupMembershipsWithAttributesResponse

func (w *wrapperListGroupMembershipsWithAttributesResponse) String() string {
	return (*ListGroupMembershipsWithAttributesResponse)(w).String()
}

func (*wrapperListGroupMembershipsWithAttributesResponse) ProtoMessage() {}

func (w *wrapperListGroupMembershipsWithAttributesResponse) ProtoReflect() protoreflect.Message {
	return (*ListGroupMembershipsWithAttributesResponse)(w).ProtoReflect()
}

// func (x *ListMemberOfResponse) Sanitize()            // is not generated as no sensitive fields found
// func (x *ListMemberOfResponse) LogValue() slog.Value // is not generated as no sensitive fields found
