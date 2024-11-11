package protobuf

import (
	"fmt"
	"slices"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/fieldmask/mask"
	"github.com/nebius/gosdk/proto/nebius"
)

func getAnnotation[T any](options proto.Message, xt protoreflect.ExtensionType) (_ T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"panic while getting extension %s from %s: %s",
				xt.TypeDescriptor().FullName(),
				options.ProtoReflect().Descriptor().FullName(),
				r,
			)
		}
	}()
	retInterface := proto.GetExtension(options, xt)
	ret, ok := retInterface.(T)
	if !ok {
		return ret, fmt.Errorf("expected %T, got %T", ret, retInterface)
	}
	return ret, nil
}

func preserveField(field protoreflect.FieldDescriptor) (bool, error) {
	behaviors, err := getAnnotation[[]nebius.FieldBehavior](field.Options(), nebius.E_FieldBehavior)
	if err != nil {
		return false, fmt.Errorf("get extension: %w", err)
	}
	return slices.Contains(behaviors, nebius.FieldBehavior_OUTPUT_ONLY), nil
}

func recursivePatchWithResetMask( //nolint:funlen,gocognit,gocyclo,cyclop // TODO: split
	data, patch protoreflect.Message,
	resetMask *mask.Mask,
	recursion int,
) error {
	if recursion >= recursionTooDeep {
		return mask.ErrRecursionTooDeep
	}
	recursion++

	dataDesc := data.Descriptor()
	patchDesc := patch.Descriptor()

	for i := range dataDesc.Fields().Len() {
		dataFieldDesc := dataDesc.Fields().Get(i)
		name := dataFieldDesc.Name()

		fieldMask, err := resetMask.GetSubMask(mask.FieldKey(name))
		if err != nil {
			return fmt.Errorf(
				"%s: couldn't get full field mask: %w", name, err,
			)
		}

		if fieldMask == nil && dataFieldDesc.ContainingOneof() != nil {
			ofm, errX := resetMask.GetSubMask(
				mask.FieldKey(dataFieldDesc.ContainingOneof().Name()),
			)
			if errX != nil {
				return fmt.Errorf(
					"%s: couldn't get full oneof mask: %w",
					dataFieldDesc.ContainingOneof().Name(), errX,
				)
			}
			if ofm != nil {
				// oneof group name can't denote inner fields
				fieldMask = mask.New()
			}
		}
		preserve, err := preserveField(dataFieldDesc)
		if err != nil {
			return fmt.Errorf(
				"%s: couldn't get field preservation options: %w", name, err,
			)
		}

		patchFieldDesc := patchDesc.Fields().ByName(name)
		if patchFieldDesc == nil {
			if fieldMask != nil && !preserve {
				data.Clear(dataFieldDesc)
			}
			continue
		}
		if fieldMask == nil &&
			patchFieldDesc.ContainingOneof() != nil {
			ofm, err := resetMask.GetSubMask(
				mask.FieldKey(patchFieldDesc.ContainingOneof().Name()),
			)
			if err != nil {
				return fmt.Errorf(
					"%s: couldn't get full oneof mask: %w",
					patchFieldDesc.ContainingOneof().Name(), err,
				)
			}
			if ofm != nil {
				// oneof group name can't denote inner fields
				fieldMask = mask.New()
			}
		}
		if dataFieldDesc.Kind() != patchFieldDesc.Kind() {
			return fmt.Errorf(
				"%s: type mismatch: data is %s, patch is %s",
				name, dataFieldDesc.Kind(), patchFieldDesc.Kind(),
			)
		}
		if !patch.Has(patchFieldDesc) {
			if fieldMask != nil && !preserve {
				data.Clear(dataFieldDesc)
			}
			continue
		}
		if dataFieldDesc.IsList() {
			if preserve {
				continue
			}
			if !patchFieldDesc.IsList() {
				return fmt.Errorf(
					"%s: data field is repeated while patch field is not", name,
				)
			}

			if !data.Has(dataFieldDesc) {
				data.Set(dataFieldDesc, data.NewField(dataFieldDesc))
			}
			dataList := data.Mutable(dataFieldDesc).List()
			patchList := patch.Get(patchFieldDesc).List()

			if dataFieldDesc.Kind() == protoreflect.MessageKind {
				// remove excess
				if dataList.Len() > patchList.Len() {
					dataList.Truncate(patchList.Len())
				}
				// patch intersection
				for j := range dataList.Len() {
					dataMsg := dataList.Get(j).Message()
					patchMsg := patchList.Get(j).Message()

					listElementMask, err := fieldMask.GetSubMask(
						mask.FieldKey(strconv.Itoa(j)),
					)
					if err != nil {
						return fmt.Errorf(
							"%s[%d]: couldn't get full field mask: %w",
							name, j, err,
						)
					}

					if err := recursivePatchWithResetMask(
						dataMsg, patchMsg, listElementMask, recursion,
					); err != nil {
						return fmt.Errorf("%s[%d]: %w", name, j, err)
					}
				}
				// add new
				for j := dataList.Len(); j < patchList.Len(); j++ {
					dataMsg := dataList.NewElement().Message()
					patchMsg := patchList.Get(j).Message()
					if err := recursivePatchWithResetMask(
						dataMsg, patchMsg, nil, recursion,
					); err != nil {
						return fmt.Errorf("%s[%d]: %w", name, j, err)
					}
					dataList.Append(protoreflect.ValueOf(dataMsg))
				}
			} else {
				dataList.Truncate(0)
				for j := range patchList.Len() {
					dataList.Append(patchList.Get(j))
				}
			}
			continue
		}
		if dataFieldDesc.IsMap() {
			if preserve {
				continue
			}
			if !patchFieldDesc.IsMap() {
				return fmt.Errorf(
					"%s: data field is map while patch field is not", name,
				)
			}
			if dataFieldDesc.MapKey().Kind() != patchFieldDesc.MapKey().Kind() {
				return fmt.Errorf(
					"%s: map key type mismatch: data has %s, patch has %s",
					name,
					dataFieldDesc.MapKey().Kind(),
					patchFieldDesc.MapKey().Kind(),
				)
			}
			if dataFieldDesc.MapValue().Kind() !=
				patchFieldDesc.MapValue().Kind() {
				return fmt.Errorf(
					"%s: map value type mismatch: data has %s, patch has %s",
					name,
					dataFieldDesc.MapValue().Kind(),
					patchFieldDesc.MapValue().Kind(),
				)
			}
			if !data.Has(dataFieldDesc) {
				data.Set(dataFieldDesc, data.NewField(dataFieldDesc))
			}
			dataMap := data.Mutable(dataFieldDesc).Map()
			patchMap := patch.Get(patchFieldDesc).Map()
			if dataFieldDesc.MapValue().Kind() == protoreflect.MessageKind {
				var innerErr error
				// patch or remove already existing values
				dataMap.Range(
					func(mk protoreflect.MapKey, v protoreflect.Value) bool {
						if patchMap.Has(mk) {
							dataMsg := v.Message()
							patchMsg := patchMap.Get(mk).Message()
							fk, err := mapKeyToFieldKey(mk)
							if err != nil {
								innerErr = fmt.Errorf(
									"%s[%s]: failed convert map key to field "+
										"key: %w",
									name, mk, err,
								)
								return false
							}

							elementMask, err := fieldMask.GetSubMask(fk)
							if err != nil {
								innerErr = fmt.Errorf(
									"%s[%s]: couldn't get full field mask:"+
										" %w",
									name, mk, err,
								)
								return false
							}

							if err := recursivePatchWithResetMask(
								dataMsg, patchMsg, elementMask, recursion,
							); err != nil {
								innerErr = fmt.Errorf(
									"%s[%s]: %w", name, mk, err,
								)
								return false
							}
						} else {
							dataMap.Clear(mk)
						}
						return true
					},
				)
				if innerErr != nil {
					return innerErr
				}
				// add newly created ones
				patchMap.Range(
					func(mk protoreflect.MapKey, _ protoreflect.Value) bool {
						if !dataMap.Has(mk) {
							dataMsg := dataMap.NewValue().Message()
							patchMsg := patchMap.Get(mk).Message()

							if err := recursivePatchWithResetMask(
								dataMsg, patchMsg, nil, recursion,
							); err != nil {
								innerErr = fmt.Errorf(
									"%s[%s]: %w", name, mk, err,
								)
								return false
							}
							dataMap.Set(mk, protoreflect.ValueOf(dataMsg))
						}
						return true
					},
				)
			} else {
				// just replace with a fully new map
				dataMap := data.NewField(dataFieldDesc).Map()
				patchMap.Range(
					func(mk protoreflect.MapKey, v protoreflect.Value) bool {
						dataMap.Set(mk, v)
						return true
					},
				)
				data.Set(dataFieldDesc, protoreflect.ValueOf(dataMap))
			}
			continue
		}
		if dataFieldDesc.Kind() == protoreflect.MessageKind {
			if preserve {
				continue
			}
			dataMsg := data.Mutable(dataFieldDesc).Message()
			patchMsg := patch.Get(patchFieldDesc).Message()
			if err := recursivePatchWithResetMask(
				dataMsg, patchMsg, fieldMask, recursion,
			); err != nil {
				return fmt.Errorf("%s: %w", name, err)
			}
			continue
		}
		patchVal := patch.Get(patchFieldDesc)
		if (patchVal.Equal(patchFieldDesc.Default()) &&
			fieldMask == nil &&
			!patchFieldDesc.HasPresence()) || preserve {
			continue
		}
		data.Set(dataFieldDesc, patchVal)
	}

	return nil
}

// PatchWithResetMask applies a patch to a target message using a [mask.Mask],
// allowing for selective updates and reset behavior as described in
// [Reset Mask Specification].
//
// The PatchWithResetMask function applies the changes specified in the patch
// [proto.Message] to the target [proto.Message], based on the provided
// [mask.Mask]. This function supports patching messages of different types,
// facilitating version conversion.
//
// The existing (and set) fields in the patch message should be of the same
// basic types as the corresponding fields in the target message, including map
// keys. If a field exists in both the target message and the patch message,
// the [mask.Mask] is ignored for that field.
//
// Fields (which are not messages, lists, or maps) absent in the patch message
// are left unchanged unless specified in the [mask.Mask]. The mask also handles
// list indexes and map keys as separate FieldKeys if the list or map holds
// messages.
//
// If a list or map field is absent in the patch message or empty, it will be
// removed only if specified in the Mask. If a list field in the patch message
// has fewer elements than the corresponding field in the target message, it is
// truncated. If it has more elements, new elements are added.
//
// For map fields, elements not present in the patch message will be removed
// from the target message if the map is not empty.
//
// Elements (submessages) present in both the patch and target maps or lists
// will be patched recursively.
//
// Parameters:
//   - data [proto.Message]: The target message to be patched.
//   - patch [proto.Message]: The patch message containing the changes.
//   - mask *[mask.Mask]: The reset mask specifying which fields to update or
//     reset.
//
// Returns:
//   - error: An error, if any, encountered during the patching process.
//
// [Reset Mask Specification]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func PatchWithResetMask(
	data, patch proto.Message,
	mask *mask.Mask,
) error {
	return recursivePatchWithResetMask(
		data.ProtoReflect(), patch.ProtoReflect(), mask, 0,
	)
}
