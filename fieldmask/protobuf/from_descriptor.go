package protobuf

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/fieldmask/mask"
)

const recursionTooDeep = 1000

func knownFieldsFromDescRecursive(
	desc protoreflect.MessageDescriptor,
	recursion int,
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
		fieldMask := mask.New()
		// OneOfs are not covered separately as they can be unset by unsetting
		// their inner fields
		switch {
		case fieldDesc.IsList() && fieldDesc.Message() != nil:
			innerKnown, err := knownFieldsFromDescRecursive(
				fieldDesc.Message(), recursion,
			)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", desc.Name(), err)
			}
			fieldMask.Any = innerKnown
		case fieldDesc.IsMap() && fieldDesc.MapValue().Message() != nil:
			innerKnown, err := knownFieldsFromDescRecursive(
				fieldDesc.MapValue().Message(), recursion,
			)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", desc.Name(), err)
			}
			fieldMask.Any = innerKnown
		case fieldDesc.Message() != nil:
			innerKnown, err := knownFieldsFromDescRecursive(
				fieldDesc.Message(), recursion,
			)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", desc.Name(), err)
			}
			fieldMask = innerKnown
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
//
// Returns:
//   - *[mask.Mask]: A [mask.Mask] representing the full set of known fields
//     defined by the message descriptor.
//   - error: An error if there was any issue collecting the known fields.
func KnownFieldsFromDescriptor(
	desc protoreflect.MessageDescriptor,
) (*mask.Mask, error) {
	ret, err := knownFieldsFromDescRecursive(desc, 0)
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
//	KnownFieldsFromDescriptor(msg.ProtoReflect().Descriptor())
//
// Parameters:
//   - msg: The [proto.Message] to extract known fields from.
//
// Returns:
//   - *[mask.Mask]: A [mask.Mask] representing the full set of known fields
//     defined by the message.
//   - error: An error if there was any issue collecting the known fields.
func KnownFieldsFromMessage(msg proto.Message) (*mask.Mask, error) {
	return KnownFieldsFromDescriptor(msg.ProtoReflect().Descriptor())
}
