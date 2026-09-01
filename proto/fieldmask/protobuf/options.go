package protobuf

import (
	"slices"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/proto/nebius"
)

type maskOptions struct {
	immutables bool
}

// Option configures mask conversion. Defaults depend on the converter:
// reset-mask converters default to [NoImmutables], while
// [KnownFieldsFromDescriptor] and [KnownFieldsFromMessage] default to
// [WithImmutables].
type Option func(*maskOptions)

// NoImmutables excludes immutable fields from converted masks.
// This is the default for reset-mask converters, but not for known-fields
// converters.
func NoImmutables() Option {
	return func(options *maskOptions) {
		options.immutables = false
	}
}

// WithImmutables includes immutable fields in converted masks.
// This is the default for known-fields converters, but not for reset-mask
// converters.
func WithImmutables() Option {
	return func(options *maskOptions) {
		options.immutables = true
	}
}

func decodeOptions(options ...Option) maskOptions {
	return decodeOptionsWithDefaults(maskOptions{}, options...)
}

func decodeOptionsWithDefaults(defaults maskOptions, options ...Option) maskOptions {
	ret := defaults
	for _, option := range options {
		option(&ret)
	}
	return ret
}

func shouldSkipImmutable(
	field protoreflect.FieldDescriptor,
	fieldIsSet bool,
	options *maskOptions,
) bool {
	if options.immutables {
		return false
	}
	behaviors, ok := proto.GetExtension(
		field.Options(), nebius.E_FieldBehavior,
	).([]nebius.FieldBehavior)
	if ok && slices.Contains(behaviors, nebius.FieldBehavior_IMMUTABLE) {
		return true
	}
	oneof := field.ContainingOneof()
	if oneof == nil || oneof.IsSynthetic() {
		return false
	}
	behaviors, ok = proto.GetExtension(
		oneof.Options(), nebius.E_OneofBehavior,
	).([]nebius.FieldBehavior)
	return !fieldIsSet && ok && slices.Contains(behaviors, nebius.FieldBehavior_IMMUTABLE)
}
