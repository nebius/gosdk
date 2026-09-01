package protobuf

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
)

const recursionTooDeep = 1000

func knownFieldFromDescRecursive(
	fieldDesc protoreflect.FieldDescriptor,
	recursion int,
	options *maskOptions,
) (*mask.Mask, error) {
	// OneOfs are not covered separately as they can be unset by unsetting
	// their inner fields
	var innerDesc protoreflect.MessageDescriptor
	collection := true
	switch {
	case fieldDesc.IsList() && fieldDesc.Message() != nil:
		innerDesc = fieldDesc.Message()
	case fieldDesc.IsMap() && fieldDesc.MapValue().Message() != nil:
		innerDesc = fieldDesc.MapValue().Message()
	case fieldDesc.Message() != nil:
		innerDesc = fieldDesc.Message()
		collection = false
	default:
		return mask.New(), nil
	}
	innerKnown, err := knownFieldsFromDescRecursive(innerDesc, recursion, options)
	if err != nil {
		return nil, err
	}
	if !collection {
		return innerKnown, nil
	}
	fieldMask := mask.New()
	fieldMask.Any = innerKnown
	return fieldMask, nil
}

func knownFieldsFromDescRecursive(
	desc protoreflect.MessageDescriptor,
	recursion int,
	options *maskOptions,
) (*mask.Mask, error) {
	if recursion >= recursionTooDeep {
		return nil, mask.ErrRecursionTooDeep
	}
	recursion++
	if desc == nil {
		return nil, nil
	}
	ret := mask.New()
	for i := range desc.Fields().Len() {
		fieldDesc := desc.Fields().Get(i)
		if shouldSkipImmutable(fieldDesc, false, options) {
			continue
		}
		fieldMask, err := knownFieldFromDescRecursive(fieldDesc, recursion, options)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", desc.Name(), err)
		}
		ret.FieldParts[mask.FieldKey(fieldDesc.Name())] = fieldMask
	}
	return ret, nil
}

// KnownFieldsFromDescriptor builds a set of known fields represented by
// a [mask.Mask] based on the given [protoreflect.MessageDescriptor].
//
// The returned [mask.Mask] represents the full set of known fields defined by
// the message descriptor. Each field in the [mask.Mask] corresponds to a field
// that is known and defined within the descriptor.
//
// Parameters:
//   - desc: The [protoreflect.MessageDescriptor] to extract known fields from.
//   - options: Conversion options. Unlike reset-mask converters, this function
//     defaults to [WithImmutables]. Pass [NoImmutables] to exclude them.
//
// Returns:
//   - *[mask.Mask]: A [mask.Mask] representing the full set of known fields
//     defined by the message descriptor.
//   - error: An error if there was any issue collecting the known fields.
func KnownFieldsFromDescriptor(
	desc protoreflect.MessageDescriptor,
	options ...Option,
) (*mask.Mask, error) {
	decodedOptions := decodeOptionsWithDefaults(maskOptions{immutables: true}, options...)
	ret, err := knownFieldsFromDescRecursive(desc, 0, &decodedOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to collect modification mask: %w", err)
	}

	return ret, nil
}

// KnownFieldsFromMessage builds a set of known fields represented by a
// [mask.Mask] based on the provided [proto.Message].
//
// The returned [mask.Mask] represents the full set of known fields defined by
// the message. Each field in the [mask.Mask] corresponds to a field that is
// known and defined within the message descriptor.
//
// This function is just a wrap around [KnownFieldsFromDescriptor]:
//
//	KnownFieldsFromDescriptor(msg.ProtoReflect().Descriptor(), options...)
//
// Parameters:
//   - msg: The [proto.Message] to extract known fields from.
//   - options: Conversion options. Unlike reset-mask converters, this function
//     defaults to [WithImmutables]. Pass [NoImmutables] to exclude them.
//
// Returns:
//   - *[mask.Mask]: A [mask.Mask] representing the full set of known fields
//     defined by the message.
//   - error: An error if there was any issue collecting the known fields.
func KnownFieldsFromMessage(msg proto.Message, options ...Option) (*mask.Mask, error) {
	return KnownFieldsFromDescriptor(msg.ProtoReflect().Descriptor(), options...)
}
