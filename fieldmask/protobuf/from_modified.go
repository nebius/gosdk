package protobuf

import (
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/fieldmask/mask"
)

func rmFrom2ListsRecursive(
	initList, modList protoreflect.List,
	desc protoreflect.FieldDescriptor,
	recursion int,
) (*mask.Mask, error) {
	if initList == nil || initList.Len() == 0 {
		return nil, nil
	}
	ret := mask.New()
	if initList.Len() > 0 && modList.Len() == 0 {
		return ret, nil
	}
	if desc.Message() == nil {
		return nil, nil
	}
	l := initList.Len()
	if l > modList.Len() {
		l = modList.Len()
	}
	for i := range l {
		initEl := initList.Get(i)
		modEl := modList.Get(i)
		resetMask, err := rmFrom2MessagesRecursive(
			initEl.Message(), modEl.Message(), recursion,
		)
		if err != nil {
			return nil, fmt.Errorf("%s[%d]: %w", desc.Name(), i, err)
		}
		if resetMask != nil && !resetMask.IsEmpty() {
			ret.FieldParts[mask.FieldKey(strconv.Itoa(i))] = resetMask
		}
	}
	if ret.IsEmpty() {
		return nil, nil
	}
	return ret, nil
}

func mapKeyToFieldKey(from protoreflect.MapKey) (mask.FieldKey, error) {
	val := from.Interface()
	switch typedVal := val.(type) {
	case bool:
		return mask.FieldKey(strconv.FormatBool(typedVal)), nil
	case int32:
		return mask.FieldKey(strconv.FormatInt(int64(typedVal), 10)), nil
	case int64:
		return mask.FieldKey(strconv.FormatInt(typedVal, 10)), nil
	case uint32:
		return mask.FieldKey(strconv.FormatUint(uint64(typedVal), 10)), nil
	case uint64:
		return mask.FieldKey(strconv.FormatUint(typedVal, 10)), nil
	case string:
		return mask.FieldKey(typedVal), nil
	default:
		return "", fmt.Errorf("unsupported map key type %T", typedVal)
	}
}

func rmFrom2MapsRecursive(
	initMap, modMap protoreflect.Map,
	desc protoreflect.FieldDescriptor,
	recursion int,
) (*mask.Mask, error) {
	if initMap == nil || initMap.Len() == 0 {
		return nil, nil
	}
	ret := mask.New()
	if initMap.Len() > 0 && modMap.Len() == 0 {
		return ret, nil
	}
	if desc.MapValue().Message() == nil {
		return nil, nil
	}
	var innerErr error
	initMap.Range(
		func(mk protoreflect.MapKey, initVal protoreflect.Value) bool {
			if modVal := modMap.Get(mk); modVal.IsValid() {
				innerMask, err := rmFrom2MessagesRecursive(
					initVal.Message(), modVal.Message(), recursion,
				)
				if err != nil {
					innerErr = fmt.Errorf(
						"%s[%s]: %w", desc.Name(), mk.String(), err,
					)
					return false
				}
				if innerMask != nil && !innerMask.IsEmpty() {
					key, err := mapKeyToFieldKey(mk)
					if err != nil {
						innerErr = fmt.Errorf(
							"%s[%s]: %w", desc.Name(), mk.String(), err,
						)
						return false
					}
					ret.FieldParts[key] = innerMask
				}
			}
			return true
		},
	)
	if innerErr != nil {
		return nil, innerErr
	}
	if ret.IsEmpty() {
		return nil, nil
	}
	return ret, nil
}

func rmFrom2MessagesRecursive( //nolint:gocognit // TODO: simplify?
	initMsg, modMsg protoreflect.Message,
	recursion int,
) (*mask.Mask, error) {
	if recursion >= recursionTooDeep {
		return nil, mask.ErrRecursionTooDeep
	}
	recursion++
	if initMsg == nil {
		return nil, nil
	}
	if modMsg == nil {
		return mask.New(), nil
	}
	initDesc := initMsg.Descriptor()
	ret := mask.New()
	for i := range initDesc.Fields().Len() {
		fieldDesc := initDesc.Fields().Get(i)
		if !initMsg.Has(fieldDesc) {
			continue
		}
		if !modMsg.Has(fieldDesc) {
			ret.FieldParts[mask.FieldKey(fieldDesc.Name())] = mask.New()
			continue
		}
		initField := initMsg.Get(fieldDesc)
		modField := modMsg.Get(fieldDesc)
		// OneOfs are not covered separately as they can be unset by unsetting
		// their inner fields
		if fieldDesc.IsList() {
			innerMask, err := rmFrom2ListsRecursive(
				initField.List(), modField.List(), fieldDesc, recursion,
			)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", initDesc.Name(), err)
			}
			if innerMask != nil {
				ret.FieldParts[mask.FieldKey(fieldDesc.Name())] = innerMask
			}
			continue
		}
		if fieldDesc.IsMap() {
			innerMask, err := rmFrom2MapsRecursive(
				initField.Map(), modField.Map(), fieldDesc, recursion,
			)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", initDesc.Name(), err)
			}
			if innerMask != nil {
				ret.FieldParts[mask.FieldKey(fieldDesc.Name())] = innerMask
			}
			continue
		}
		if fieldDesc.Message() != nil {
			innerMask, err := rmFrom2MessagesRecursive(
				initField.Message(), modField.Message(), recursion,
			)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", initDesc.Name(), err)
			}
			if innerMask != nil && !innerMask.IsEmpty() {
				ret.FieldParts[mask.FieldKey(fieldDesc.Name())] = innerMask
			}
			continue
		}
		if modField.Equal(fieldDesc.Default()) && !initField.Equal(modField) {
			ret.FieldParts[mask.FieldKey(fieldDesc.Name())] = mask.New()
		}
	}
	return ret, nil
}

// ResetMaskFromModified creates a reset [mask.Mask] of fields that have been
// removed from the object according to the [ResetMask specification] between
// two Protobuf messages: initial and modified. It traverses the Protobuf
// messages using protoreflect to compare their fields, identifying and
// capturing fields that have been removed in the modified message.
//
// Parameters:
//   - initial: The original [proto.Message] before modification.
//   - modified: The modified [proto.Message] after applying changes.
//
// Returns:
//   - *[mask.Mask]: A [mask.Mask] representing fields that have been removed
//     from the initial message based on the modification specified by the
//     [ResetMask specification] between initial and modified.
//   - error: An error if there was any issue collecting the modification mask.
//
// [ResetMask specification]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func ResetMaskFromModified(
	initial, modified proto.Message,
) (*mask.Mask, error) {
	if initial == nil || modified == nil {
		return nil, errors.New("received nil")
	}
	initMsg := initial.ProtoReflect()
	modMsg := modified.ProtoReflect()

	// Check if the messages are of the same type
	if initMsg.Descriptor().FullName() != modMsg.Descriptor().FullName() {
		return nil, fmt.Errorf(
			"recieved messages of different types: initial %s, modified %s",
			initMsg.Descriptor().FullName(), modMsg.Descriptor().FullName(),
		)
	}

	// Collect modification mask by comparing initial and modified messages
	ret, err := rmFrom2MessagesRecursive(initMsg, modMsg, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to collect modification mask: %w", err)
	}

	return ret, nil
}
