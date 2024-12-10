package protobuf

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
)

var ErrBasicTypeInTheMiddle = errors.New("found basic type in the middle of the path, can't descend deeper")
var ErrNotFound = errors.New("field name not found")
var ErrOutOfBounds = errors.New("out of list bounds")

func removeListElement(l protoreflect.List, pos int) {
	rest := []protoreflect.Value{}
	for i := pos + 1; i < l.Len(); i++ {
		rest = append(rest, l.Get(i))
	}
	if l.Len() > pos {
		l.Truncate(pos)
	}
	for _, el := range rest {
		l.Append(el)
	}
}

func recursiveReplaceAtFieldPath( //nolint:funlen,gocognit,gocyclo,cyclop // TODO: split
	data protoreflect.Message,
	path mask.FieldPath,
	newVal protoreflect.Value,
	recursion int,
) error {
	if recursion >= recursionTooDeep {
		return mask.ErrRecursionTooDeep
	}
	recursion++

	if len(path) == 0 {
		if newVal.Interface() == nil {
			desc := data.Descriptor()
			fieldDescs := desc.Fields()
			for i := range fieldDescs.Len() {
				fd := fieldDescs.Get(i)
				data.Clear(fd)
			}
			return nil
		}
		if msg, ok := newVal.Interface().(protoreflect.Message); ok {
			if msg.Descriptor().FullName() != data.Descriptor().FullName() {
				return fmt.Errorf(
					"can't replace message %s with message %s",
					data.Descriptor().FullName(), msg.Descriptor().FullName(),
				)
			}
			desc := data.Descriptor()
			fieldDescs := desc.Fields()
			for i := range fieldDescs.Len() {
				fd := fieldDescs.Get(i)
				if !msg.Has(fd) {
					if data.Has(fd) {
						data.Clear(fd)
					}
				} else {
					data.Set(fd, msg.Get(fd))
				}
			}
			return nil
		}
		return fmt.Errorf(
			"can't replace message %s with value of type %T",
			data.Descriptor().FullName(), newVal.Interface(),
		)
	}
	fieldKey, rest := path[0], path[1:]

	desc := data.Descriptor()
	fieldDesc := desc.Fields().ByName(protoreflect.Name(fieldKey))
	if fieldDesc == nil {
		return fmt.Errorf("%s: %w", fieldKey, ErrNotFound)
	}
	if len(rest) == 0 {
		// all nil types
		if newVal.Interface() == nil {
			data.Clear(fieldDesc)
			return nil
		}
		switch o := newVal.Interface().(type) {
		case protoreflect.Map:
			if !fieldDesc.IsMap() {
				return fmt.Errorf(
					"%s: can't replace %s with map", fieldKey, fieldDesc.Kind(),
				)
			}
			fMapVal := data.NewField(fieldDesc)
			fMap := fMapVal.Map()
			var innerErr error
			o.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
				defer func() {
					if r := recover(); r != nil {
						innerErr = fmt.Errorf(
							"%s[%s]: failed to set map value: panic: %s",
							fieldKey, mk, r,
						)
					}
				}()
				if v.IsValid() {
					fMap.Set(mk, v)
				}
				return true
			})
			if innerErr != nil {
				return innerErr
			}
			data.Set(fieldDesc, fMapVal)
			return nil
		case protoreflect.List:
			if !fieldDesc.IsList() {
				return fmt.Errorf(
					"%s: can't replace %s with list",
					fieldKey, fieldDesc.Kind(),
				)
			}
			fListVal := data.NewField(fieldDesc)
			fList := fListVal.List()
			for i := range o.Len() {
				err := func() (err error) {
					defer func() {
						if r := recover(); r != nil {
							err = fmt.Errorf("panic: %s", r)
						}
					}()
					el := o.Get(i)
					if el.IsValid() {
						fList.Append(o.Get(i))
					}
					return nil
				}()
				if err != nil {
					return fmt.Errorf(
						"%s: failed to set list value: %w", fieldKey, err,
					)
				}
			}
			data.Set(fieldDesc, fListVal)
			return nil
		}
		err := func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic: %s", r)
				}
			}()
			data.Set(fieldDesc, newVal)
			return nil
		}()
		if err != nil {
			return fmt.Errorf("%s: failed to set value: %w", fieldKey, err)
		}
		return nil
	}
	if fieldDesc.IsList() {
		list := data.Mutable(fieldDesc).List()
		listNum, restList := rest[0], rest[1:]
		i, err := strconv.Atoi(string(listNum))
		if err != nil {
			return fmt.Errorf(
				"%s.%s: failed to convert to list index", fieldKey, listNum,
			)
		}
		if i < 0 || i > list.Len() {
			return fmt.Errorf("%s[%d]: %w", fieldKey, i, ErrOutOfBounds)
		}
		if len(restList) == 0 {
			if newVal.Interface() == nil {
				removeListElement(list, i)
				return nil
			}
			err := func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("panic: %s", r)
					}
				}()
				if list.Len() == i {
					list.Append(newVal)
				} else {
					list.Set(i, newVal)
				}
				return nil
			}()
			if err != nil {
				return fmt.Errorf(
					"%s[%d]: failed to set value: %w", fieldKey, i, err,
				)
			}
			return nil
		}
		if fieldDesc.Kind() != protoreflect.MessageKind {
			return fmt.Errorf("%s[%d]: %w", fieldKey, i, ErrBasicTypeInTheMiddle)
		}
		var listMsg protoreflect.Message
		if i == list.Len() {
			listMsg = list.NewElement().Message()
		} else {
			listMsg = list.Get(i).Message()
		}
		if err := recursiveReplaceAtFieldPath(
			listMsg, restList, newVal, recursion,
		); err != nil {
			return fmt.Errorf("%s[%d]: %w", fieldKey, i, err)
		}
		if i == list.Len() {
			list.Append(protoreflect.ValueOf(listMsg))
		} else {
			list.Set(i, protoreflect.ValueOf(listMsg))
		}
		return nil
	}
	if fieldDesc.IsMap() {
		fMap := data.Mutable(fieldDesc).Map()
		mapKey, restMap := rest[0], rest[1:]
		mk, err := mapKeyFromFieldKey(mapKey, fieldDesc.MapKey().Kind())
		if err != nil {
			return fmt.Errorf(
				"%s.%s: failed to convert to map key: %w",
				fieldKey, mapKey, err,
			)
		}
		if len(restMap) == 0 {
			if newVal.Interface() == nil {
				fMap.Clear(mk)
				return nil
			}
			err := func() (err error) {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("panic: %s", r)
					}
				}()
				fMap.Set(mk, newVal)
				return nil
			}()
			if err != nil {
				return fmt.Errorf(
					"%s[%s]: failed to set value: %w",
					fieldKey, mk, err,
				)
			}
			return nil
		}
		if fieldDesc.MapValue().Kind() != protoreflect.MessageKind {
			return fmt.Errorf("%s[%s]: %w", fieldKey, mk, ErrBasicTypeInTheMiddle)
		}
		var valMsg protoreflect.Message
		if fMap.Has(mk) {
			valMsg = fMap.Get(mk).Message()
		} else {
			valMsg = fMap.NewValue().Message()
		}
		if err := recursiveReplaceAtFieldPath(
			valMsg, restMap, newVal, recursion,
		); err != nil {
			return fmt.Errorf("%s[%s]: %w", fieldKey, mk, err)
		}
		fMap.Set(mk, protoreflect.ValueOf(valMsg))
		return nil
	}
	if fieldDesc.Kind() != protoreflect.MessageKind {
		return fmt.Errorf("%s: %w", fieldKey, ErrBasicTypeInTheMiddle)
	}
	valMsg := data.Mutable(fieldDesc).Message()
	if err := recursiveReplaceAtFieldPath(
		valMsg, rest, newVal, recursion,
	); err != nil {
		return fmt.Errorf("%s: %w", fieldKey, err)
	}

	return nil
}

func mapKeyFromFieldKey(
	from mask.FieldKey,
	kind protoreflect.Kind,
) (protoreflect.MapKey, error) {
	switch kind { //nolint:exhaustive // don't need to check all kinds
	case protoreflect.BoolKind:
		ret, err := strconv.ParseBool(string(from))
		if err != nil {
			return protoreflect.MapKey{}, fmt.Errorf(
				"failed to convert to bool map key: %w", err,
			)
		}
		return protoreflect.ValueOf(ret).MapKey(), nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Sfixed32Kind:
		ret, err := strconv.ParseInt(string(from), 10, 32)
		if err != nil {
			return protoreflect.MapKey{}, fmt.Errorf(
				"failed to convert to int32 map key: %w", err,
			)
		}
		return protoreflect.ValueOf(int32(ret)).MapKey(), nil
	case protoreflect.Int64Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed64Kind:
		ret, err := strconv.ParseInt(string(from), 10, 64)
		if err != nil {
			return protoreflect.MapKey{}, fmt.Errorf(
				"failed to convert to int64 map key: %w", err,
			)
		}
		return protoreflect.ValueOf(ret).MapKey(), nil
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		ret, err := strconv.ParseUint(string(from), 10, 32)
		if err != nil {
			return protoreflect.MapKey{}, fmt.Errorf(
				"failed to convert to uint32 map key: %w", err,
			)
		}
		return protoreflect.ValueOf(uint32(ret)).MapKey(), nil
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		ret, err := strconv.ParseUint(string(from), 10, 64)
		if err != nil {
			return protoreflect.MapKey{}, fmt.Errorf(
				"failed to convert to uint64 map key: %w", err,
			)
		}
		return protoreflect.ValueOf(ret).MapKey(), nil
	case protoreflect.StringKind:
		return protoreflect.ValueOf(string(from)).MapKey(), nil
	default:
		return protoreflect.MapKey{}, fmt.Errorf(
			"unsupported map key kind %s", kind,
		)
	}
}

type cloneMap struct {
	Source protoreflect.Map
}

func newCloneMap(m protoreflect.Map) protoreflect.Map {
	return cloneMap{
		Source: m,
	}
}

func (c cloneMap) Clear(mk protoreflect.MapKey) {
	c.Source.Clear(mk)
}
func (c cloneMap) Get(mk protoreflect.MapKey) protoreflect.Value {
	return cloneValue(c.Source.Get(mk))
}
func (c cloneMap) Has(mk protoreflect.MapKey) bool {
	return c.Source.Has(mk)
}
func (c cloneMap) IsValid() bool {
	return c.Source.IsValid()
}
func (c cloneMap) Len() int {
	return c.Source.Len()
}
func (c cloneMap) Mutable(mk protoreflect.MapKey) protoreflect.Value {
	return cloneValue(c.Source.Get(mk))
}
func (c cloneMap) NewValue() protoreflect.Value {
	return c.Source.NewValue()
}
func (c cloneMap) Range(
	f func(mk protoreflect.MapKey, v protoreflect.Value) bool,
) {
	c.Source.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
		vClone := cloneValue(v)
		return f(mk, vClone)
	})
}
func (c cloneMap) Set(mk protoreflect.MapKey, v protoreflect.Value) {
	c.Source.Set(mk, v)
}

type cloneList struct {
	Source protoreflect.List
}

func newCloneList(l protoreflect.List) protoreflect.List {
	return cloneList{
		Source: l,
	}
}

func (c cloneList) Append(v protoreflect.Value) {
	c.Source.Append(v)
}
func (c cloneList) AppendMutable() protoreflect.Value {
	return c.Source.AppendMutable()
}
func (c cloneList) Get(i int) protoreflect.Value {
	return cloneValue(c.Source.Get(i))
}
func (c cloneList) IsValid() bool {
	return c.Source.IsValid()
}
func (c cloneList) Len() int {
	return c.Source.Len()
}
func (c cloneList) NewElement() protoreflect.Value {
	return c.Source.NewElement()
}
func (c cloneList) Set(i int, v protoreflect.Value) {
	c.Source.Set(i, v)
}
func (c cloneList) Truncate(i int) {
	c.Source.Truncate(i)
}

func cloneValue(val protoreflect.Value) protoreflect.Value {
	switch o := val.Interface().(type) {
	case protoreflect.Message:
		if o == nil || o.Interface() == nil {
			return protoreflect.ValueOf(nil)
		}
		if reflect.ValueOf(o.Interface()).IsNil() {
			return protoreflect.ValueOf(nil)
		}
		return protoreflect.ValueOfMessage(
			proto.Clone(o.Interface()).ProtoReflect(),
		)
	case protoreflect.Map:
		return protoreflect.ValueOfMap(newCloneMap(o))
	case protoreflect.List:
		return protoreflect.ValueOfList(newCloneList(o))
	}
	return protoreflect.ValueOf(val.Interface())
}

func convertToValue(val interface{}) (_ protoreflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %s", r)
		}
	}()
	switch valOfType := val.(type) {
	case protoreflect.Value:
		return cloneValue(valOfType), nil
	case proto.Message:
		if valOfType == nil {
			return protoreflect.ValueOf(nil), nil
		}
		return protoreflect.ValueOfMessage(
			proto.Clone(valOfType).ProtoReflect(),
		), nil
	default:
		return protoreflect.ValueOf(val), nil
	}
}

// ReplaceAtFieldPath traverses a [proto.Message] based on the
// provided [mask.FieldPath] and replaces or clears the value at the end of the
// path.
//
// The ReplaceAtFieldPath function traverses the specified
// [proto.Message] based on the given [mask.FieldPath] to locate a
// specific field. It replaces the value of the field at the end of the path
// with the provided `newVal`, regardless of whether the field is a message,
// list, map, or a scalar. If `nil` is provided as `newVal`, it clears the field
// (sets it to its zero value) or removes the list/map value at the specified
// position.
//
// The value type has to be the same as the target field type, strictly.
//
// If the [mask.FieldPath] is empty, the function performs a shallow replace of
// each message field in the [proto.Message] with the ones from the
// newVal.
//
// Map keys and list indexes are treated as separate FieldKeys during traversal.
// List index will be replaced or added if it's i+1 index. Removal will shrink
// the list one element down, removal of i+1 element takes no effect, as well as
// removal of unset element of a map.
//
// If a basic field (scalar) is encountered along the path, the function fails
// with an error wrapped around [ErrBasicTypeInTheMiddle], possibly multiple times
// if the basic field is not at the beginning of the path.
//
// Parameters:
//   - data [proto.Message]: The message to traverse and modify.
//   - path [mask.FieldPath]: The path indicating the path to traverse within
//     the `data`.
//   - newVal: The new value to replace the existing value at the end of the
//     [mask.FieldPath]. Pass `nil` to clear the field or remove a list/map
//     value.
//
// Returns:
//   - error: An error encountered during traversal or replacement.
func ReplaceAtFieldPath(
	data proto.Message,
	path mask.FieldPath,
	newVal interface{},
) error {
	valValue, err := convertToValue(newVal)
	if err != nil {
		return fmt.Errorf("failed to convert to protoreflect.Value: %w", err)
	}
	return recursiveReplaceAtFieldPath(data.ProtoReflect(), path, valValue, 0)
}

func recursiveGetAtFieldPath( //nolint:funlen,gocognit // TODO: split?
	data protoreflect.Message,
	path mask.FieldPath,
	recursion int,
) (protoreflect.Value, protoreflect.FieldDescriptor, error) {
	if recursion >= recursionTooDeep {
		return protoreflect.Value{}, nil, mask.ErrRecursionTooDeep
	}
	recursion++

	if len(path) == 0 {
		return protoreflect.ValueOf(data), nil, nil
	}
	fieldKey, rest := path[0], path[1:]
	wasOneOf := false

	desc := data.Descriptor()
	fieldDesc := desc.Fields().ByName(protoreflect.Name(fieldKey))
	if fieldDesc == nil {
		if oof := desc.Oneofs().ByName(
			protoreflect.Name(fieldKey),
		); oof != nil {
			wasOneOf = true
			fieldDesc = data.WhichOneof(oof)
			if fieldDesc == nil {
				fieldDesc = oof.Fields().ByNumber(0)
				return protoreflect.Value{}, fieldDesc, nil
			}
		} else {
			return protoreflect.Value{}, nil, fmt.Errorf(
				"%s: %w", fieldKey, ErrNotFound,
			)
		}
	}
	if len(rest) == 0 {
		if fieldDesc.HasPresence() && !data.Has(fieldDesc) &&
			fieldDesc.Kind() != protoreflect.MessageKind {
			return protoreflect.ValueOf(nil), fieldDesc, nil
		}
		return data.Get(fieldDesc), fieldDesc, nil
	}
	if fieldDesc.IsList() {
		list := data.Get(fieldDesc).List()
		listNum, restList := rest[0], rest[1:]
		i, err := strconv.Atoi(string(listNum))
		if err != nil {
			return protoreflect.Value{}, nil, fmt.Errorf(
				"%s.%s: failed to convert to list index", fieldKey, listNum,
			)
		}
		if i < 0 || i >= list.Len() {
			return protoreflect.Value{}, nil, fmt.Errorf(
				"%s[%d]: %w", fieldKey, i, ErrOutOfBounds,
			)
		}
		if len(restList) == 0 {
			return list.Get(i), fieldDesc, nil
		}
		if fieldDesc.Kind() != protoreflect.MessageKind {
			return list.Get(i), fieldDesc, fmt.Errorf(
				"%s[%d]: %w", fieldKey, i, ErrBasicTypeInTheMiddle,
			)
		}
		listMsg := list.Get(i).Message()
		rv, rd, err := recursiveGetAtFieldPath(listMsg, restList, recursion)
		if err != nil {
			return rv, rd, fmt.Errorf("%s[%d]: %w", fieldKey, i, err)
		}
		return rv, rd, nil
	}
	if fieldDesc.IsMap() {
		fMap := data.Get(fieldDesc).Map()
		mapKey, restMap := rest[0], rest[1:]
		mk, err := mapKeyFromFieldKey(mapKey, fieldDesc.MapKey().Kind())
		if err != nil {
			return protoreflect.Value{}, nil, fmt.Errorf(
				"%s.%s: failed to convert to map key: %w",
				fieldKey, mapKey, err,
			)
		}
		if len(restMap) == 0 {
			ret := fMap.Get(mk)
			if ret.IsValid() {
				return ret, fieldDesc.MapValue(), nil
			}
			return ret, fieldDesc.MapValue(), fmt.Errorf(
				"%s[%s]: %w", fieldKey, mapKey, ErrNotFound,
			)
		}
		if fieldDesc.MapValue().Kind() != protoreflect.MessageKind {
			return fMap.Get(mk), fieldDesc.MapValue(), fmt.Errorf(
				"%s[%s]: %w", fieldKey, mk, ErrBasicTypeInTheMiddle,
			)
		}
		valMsg := fMap.Get(mk).Message()
		rv, rd, err := recursiveGetAtFieldPath(valMsg, restMap, recursion)
		if err != nil {
			return rv, rd, fmt.Errorf("%s[%s]: %w", fieldKey, mk, err)
		}
		return rv, rd, nil
	}
	if fieldDesc.Kind() != protoreflect.MessageKind || wasOneOf {
		return data.Get(fieldDesc), fieldDesc, fmt.Errorf(
			"%s: %w", fieldKey, ErrBasicTypeInTheMiddle,
		)
	}
	valMsg := data.Get(fieldDesc).Message()
	rv, rd, err := recursiveGetAtFieldPath(valMsg, rest, recursion)
	if err != nil {
		return rv, rd, fmt.Errorf("%s: %w", fieldKey, err)
	}
	return rv, rd, nil
}

// GetAtFieldPath traverses a [proto.Message] based on the
// provided [mask.FieldPath] to retrieve a value and its field descriptor.
//
// The GetAtFieldPath function traverses the specified [proto.Message] based on
// the given [mask.FieldPath] to locate a specific field. It returns the
// [protoreflect.Value] of the field and its associated
// [protoreflect.FieldDescriptor] if found. If the [mask.FieldPath] is empty,
// the function returns the passed [proto.Message] itself (wrapped into
// [protoreflect.Value]) and a `nil` instead of [protoreflect.FieldDescriptor],
// as it's not considered a field.
//
// It will return [protoreflect.ValueOf](nil) if the field can be tested for
// presence (eg optional or oneof), and is actually not present.
//
// Map keys and list indexes are treated as separate FieldKeys during traversal.
//
// If a basic field (scalar) is encountered along the path, the function fails
// with an error wrapped around [ErrBasicTypeInTheMiddle], possibly multiple times
// if the basic field is not at the beginning of the path.
//
// Other errors will be accompanied by the value that is not
// [protoreflect.Value.IsValid]().
//
// Note: [protoreflect.ValueOf](nil) for fields with explicit presence will also
// show negative [protoreflect.Value.IsValid]().
//
// Parameters:
//   - data [proto.Message]: The message to traverse.
//   - path [mask.FieldPath]: The path indicating the path to traverse within
//     the [proto.Message].
//
// Returns:
//   - [protoreflect.Value]: The value of the field located at the end of the
//     [mask.FieldPath].
//   - [protoreflect.FieldDescriptor]: The field descriptor corresponding to the
//     located field.
//   - error: An error encountered during traversal.
func GetAtFieldPath(
	data proto.Message,
	path mask.FieldPath,
) (protoreflect.Value, protoreflect.FieldDescriptor, error) {
	return recursiveGetAtFieldPath(data.ProtoReflect(), path, 0)
}
