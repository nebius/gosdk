package protobuf

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/fieldmask/mask"
)

func rmFromMessageRecursive( //nolint:gocognit // TODO: simplify
	resetMask *mask.Mask,
	updMsg protoreflect.Message,
	recursion int,
) error {
	if recursion >= recursionTooDeep {
		return mask.ErrRecursionTooDeep
	}
	recursion++
	if updMsg == nil {
		return nil
	}
	desc := updMsg.Descriptor()
	for i := range desc.Fields().Len() {
		fieldDesc := desc.Fields().Get(i)
		fieldMask := resetMask.FieldParts[mask.FieldKey(fieldDesc.Name())]
		if fieldMask == nil {
			fieldMask = mask.New()
		}
		if !updMsg.Has(fieldDesc) {
			resetMask.FieldParts[mask.FieldKey(fieldDesc.Name())] = fieldMask
			continue
		}
		field := updMsg.Get(fieldDesc)

		// OneOfs are not covered separately as they can be unset by unsetting
		// their inner fields
		if fieldDesc.IsList() {
			list := field.List()
			if list.Len() == 0 {
				resetMask.FieldParts[mask.FieldKey(
					fieldDesc.Name(),
				)] = fieldMask
			} else if fieldDesc.Message() != nil {
				innerMask := fieldMask.Any
				if innerMask == nil {
					innerMask = mask.New()
					fieldMask.Any = innerMask
				}
				resetMask.FieldParts[mask.FieldKey(
					fieldDesc.Name(),
				)] = fieldMask
				for i := range list.Len() {
					err := rmFromMessageRecursive(
						innerMask, list.Get(i).Message(), recursion,
					)
					if err != nil {
						return fmt.Errorf("%s[%d]: %w", desc.Name(), i, err)
					}
				}
			}
			continue
		}
		if fieldDesc.IsMap() {
			fMap := field.Map()
			if fMap.Len() == 0 {
				resetMask.FieldParts[mask.FieldKey(
					fieldDesc.Name(),
				)] = fieldMask
			} else if fieldDesc.MapValue().Message() != nil {
				innerMask := fieldMask.Any
				if innerMask == nil {
					innerMask = mask.New()
					fieldMask.Any = innerMask
				}
				resetMask.FieldParts[mask.FieldKey(
					fieldDesc.Name(),
				)] = fieldMask
				var innerErr error
				fMap.Range(
					func(mk protoreflect.MapKey, v protoreflect.Value) bool {
						err := rmFromMessageRecursive(
							innerMask, v.Message(), recursion,
						)
						if err != nil {
							innerErr = fmt.Errorf(
								"%s[%s]: %w", desc.Name(), mk, err,
							)
							return false
						}
						return true
					},
				)
				if innerErr != nil {
					return innerErr
				}
			}
			continue
		}
		if fieldDesc.Message() != nil {
			err := rmFromMessageRecursive(fieldMask, field.Message(), recursion)
			if err != nil {
				return fmt.Errorf("%s: %w", desc.Name(), err)
			}
			resetMask.FieldParts[mask.FieldKey(fieldDesc.Name())] = fieldMask
			continue
		}
		if field.Equal(fieldDesc.Default()) {
			resetMask.FieldParts[mask.FieldKey(fieldDesc.Name())] = fieldMask
		}
	}
	return nil
}

// ResetMaskFromMessage creates a reset [mask.Mask] of fields that are default
// or unset in the given [proto.Message], representing the fields
// that would be reset when applying the message.
//
// The reset mask identifies fields that are default or unset in the message,
// indicating the fields that would be reset to their default values or removed
// when the message is applied.
//
// Parameters:
//   - update: The [proto.Message] to analyze and create the reset
//     mask from.
//
// Returns:
//   - *[mask.Mask]: A [mask.Mask] representing fields that are default or unset
//     in the message, indicating the fields that would be reset or removed when
//     applying the message.
//   - error: An error if there was any issue collecting the reset mask.
func ResetMaskFromMessage(
	update proto.Message,
) (*mask.Mask, error) {
	if update == nil {
		return nil, nil
	}
	updMsg := update.ProtoReflect()
	ret := mask.New()

	// Collect modification mask by comparing initial and modified messages
	err := rmFromMessageRecursive(ret, updMsg, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to collect modification mask: %w", err)
	}

	return ret, nil
}
