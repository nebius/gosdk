package protobuf

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/fieldmask/mask"
)

func recursiveFilterListWithSelectMask( //nolint:gocognit
	list protoreflect.List,
	listDesc protoreflect.FieldDescriptor,
	fieldMask *mask.Mask,
	reduceLists bool,
	recursion int,
) error {
	if fieldMask.IsEmpty() { // empty = full for X-FilterMap
		return nil
	}
	listName := listDesc.Name()
	var ret []protoreflect.Value
	if reduceLists {
		ret = make([]protoreflect.Value, 0, list.Len())
	}
	for j := range list.Len() {
		val := list.Get(j)
		innerMask, err := fieldMask.GetSubMask(
			mask.FieldKey(strconv.Itoa(j)),
		)
		if err != nil {
			return fmt.Errorf(
				"%s[%d]: couldn't get full field mask: %w",
				listName, j, err,
			)
		}
		if innerMask == nil {
			if !reduceLists {
				list.Set(j, list.NewElement())
			}
			continue
		}
		if listDesc.Message() != nil {
			if err := recursiveFilterWithSelectMask(
				val.Message(), innerMask, reduceLists, recursion,
			); err != nil {
				return fmt.Errorf("%s[%d]: %w", listName, j, err)
			}
		}
		if reduceLists {
			ret = append(ret, val)
		} else {
			list.Set(j, val)
		}
	}
	if reduceLists {
		list.Truncate(len(ret))
		for i, v := range ret {
			list.Set(i, v)
		}
	}

	return nil
}

func recursiveFilterMapWithSelectMask(
	fMap protoreflect.Map,
	fMapDesc protoreflect.FieldDescriptor,
	fieldMask *mask.Mask,
	reduceLists bool,
	recursion int,
) error {
	if fieldMask.IsEmpty() { // empty = full for X-FilterMap
		return nil
	}
	mapName := fMapDesc.Name()

	isMsg := fMapDesc.MapValue().Message() != nil

	var innerErr error
	fMap.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
		fk, err := mapKeyToFieldKey(mk)
		if err != nil {
			innerErr = fmt.Errorf(
				"%s[%s]: failed to convert map key to field key: %w",
				mapName, mk, err,
			)
			return false
		}
		innerMask, err := fieldMask.GetSubMask(fk)
		if err != nil {
			innerErr = fmt.Errorf(
				"%s[%s]: couldn't get full field mask: %w",
				mapName, mk, err,
			)
			return false
		}
		if innerMask == nil {
			fMap.Clear(mk)
		} else if isMsg {
			if err := recursiveFilterWithSelectMask(
				v.Message(), innerMask, reduceLists, recursion,
			); err != nil {
				innerErr = fmt.Errorf("%s[%s]: %w", mapName, mk, err)
				return false
			}
			fMap.Set(mk, v)
		}
		return true
	})
	return innerErr
}

func recursiveFilterWithSelectMask( //nolint:gocognit
	message protoreflect.Message,
	fieldMask *mask.Mask,
	reduceLists bool,
	recursion int,
) error {
	if recursion >= recursionTooDeep {
		return mask.ErrRecursionTooDeep
	}
	recursion++
	msgDesc := message.Descriptor()

	if fieldMask == nil || fieldMask.IsEmpty() { // empty = full for X-FilterMap
		return nil
	}
	for i := range msgDesc.Fields().Len() {
		fieldDesc := msgDesc.Fields().Get(i)
		name := fieldDesc.Name()

		if !message.Has(fieldDesc) {
			continue
		}

		innerMask, err := fieldMask.GetSubMask(mask.FieldKey(name))
		if err != nil {
			return fmt.Errorf("%s: couldn't get full field mask: %w", name, err)
		}
		if innerMask == nil && fieldDesc.ContainingOneof() != nil {
			oname := fieldDesc.ContainingOneof().Name()
			oom, err := fieldMask.GetSubMask(mask.FieldKey(oname))
			if err != nil {
				return fmt.Errorf(
					"%s: couldn't get full field mask: %w", oname, err,
				)
			}
			if oom != nil {
				// grouping mask is always treated as empty
				innerMask = mask.New()
			}
		}
		if innerMask == nil {
			message.Clear(fieldDesc)
			continue
		}
		if fieldDesc.IsList() {
			list := message.Mutable(fieldDesc).List()
			if err := recursiveFilterListWithSelectMask(
				list, fieldDesc, innerMask, reduceLists, recursion,
			); err != nil {
				return err
			}
			continue
		}
		if fieldDesc.IsMap() {
			vMap := message.Mutable(fieldDesc).Map()
			if err := recursiveFilterMapWithSelectMask(
				vMap, fieldDesc, innerMask, reduceLists, recursion,
			); err != nil {
				return err
			}
			continue
		}
		if fieldDesc.Message() != nil {
			vMessage := message.Get(fieldDesc).Message()
			if err := recursiveFilterWithSelectMask(
				vMessage, innerMask, reduceLists, recursion,
			); err != nil {
				return fmt.Errorf("%s: %w", name, err)
			}
			message.Set(fieldDesc, protoreflect.ValueOf(vMessage))
			continue
		}
	}

	return nil
}

// FilterWithSelectMask recursively filters a target message based on the
// provided [mask.Mask], clearing all fields not present in the [mask.Mask]
// according to `X-SelectMask` behavior by [Design].
//
// If a field in the [mask.Mask] is empty (or equivalent to a wildcard mask), it
// signifies that the entire field and its children should be retained. If a
// [mask.Mask] specifies specific children for a field, only those specified
// children will be retained, and all other children will be removed.
//
// For repeated fields (lists) that have not-wildcard selector, this function
// will return default values for unselected fields. For proper filtering use
// [FilterWithSelectMaskAndReduceLists]. The filtering is disabled by default as
// this will preserve element indexes as well as be compatible to services that
// lack `X-SelectMask` support.
//
// Example:
//
//	msg := &MyMessage{...} // Initialize your proto message
//	fm := mask.ParseMust("a.b,c,d.*.e")
//
//	// Filter the message using the fieldMask
//	err := FilterWithSelectMask(msg, fm)
//	if err != nil {
//	    log.Fatalf("Error filtering message: %v", err)
//	}
//
// In the above example:
//   - Fields 'a', 'c' and 'd' will be retained.
//   - For field 'a', only its child 'b' will be retained, and all other
//     children of 'a' will be removed.
//   - Field 'c' will be retained as is, without any modifications.
//   - For field 'd', only the field 'e' of its children (if any) will be
//     retained, and all other fields of children of 'd' will be removed.
//
// Example with list:
//
//	msg := &SomeMessage{
//	   Field: []*OtherMessage{
//	     {SomeInnerFields: "for_field #0"}, // Index 0
//	     {SomeInnerFields: "for_field #1"}, // Index 1
//	     {SomeInnerFields: "for_field #2"}, // Index 2
//	     {SomeInnerFields: "for_field #3"}, // Index 3
//	   },
//	}
//	fm := mask.ParseMust("field.(1,3)")
//
//	FilterWithSelectMask(msg, fm)
//	msg == &SomeMessage{
//	  Field: []*OtherMessage{
//	    {}, // Index 0
//	    {SomeInnerFields: "for_field #1"}, // Index 1
//	    {}, // Index 2
//	    {SomeInnerFields: "for_field #3"}, // Index 3
//	  },
//	}
//
// Parameters:
//   - message [proto.Message]: The target message to be filtered.
//   - fieldMask *[mask.Mask]: The Mask specifying which fields to update or
//     reset.
//
// Returns:
//   - error: An error, if any, encountered during the filtering process.
//
// [Design]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func FilterWithSelectMask(
	message proto.Message,
	fieldMask *mask.Mask,
) error {
	return recursiveFilterWithSelectMask(
		message.ProtoReflect(), fieldMask, false, 0,
	)
}

// FilterWithSelectMaskAndReduceLists recursively filters a target message based
// on the provided [mask.Mask], clearing all fields not present in the
// [mask.Mask] according to `X-SelectMask` behavior by [Design].
//
// If a field in the [mask.Mask] is empty (or equivalent to a wildcard mask), it
// signifies that the entire field and its children should be retained. If a
// [mask.Mask] specifies specific children for a field, only those specified
// children will be retained, and all other children will be removed.
//
// For repeated fields (lists):
//
//   - When reduceLists is false:
//     The behavior is to retain the structure of the list while clearing
//     non-matching elements. Empty placeholders will be inserted for indices
//     that are explicitly excluded from the
//     [mask.Mask].
//
//   - When reduceLists is true:
//     The behavior is to compact the list by removing non-matching elements
//     entirely.
//
// Example:
//
//	msg := &MyMessage{...} // Initialize your proto message
//	fm := mask.ParseMust("a.b,c,d.*.e")
//
//	// Filter the message using the fieldMask
//	err := FilterWithSelectMaskAndReduceLists(msg, fm, true)
//	if err != nil {
//	    log.Fatalf("Error filtering message: %v", err)
//	}
//
// In the above example:
//   - Fields 'a', 'c' and 'd' will be retained.
//   - For field 'a', only its child 'b' will be retained, and all other
//     children of 'a' will be removed.
//   - Field 'c' will be retained as is, without any modifications.
//   - For field 'd', only the field 'e' of its children (if any) will be
//     retained, and all other fields of children of 'd' will be removed.
//
// Example with list:
//
//		msg := &SomeMessage{
//		  Field: []*OtherMessage{
//		    {SomeInnerFields: "for_field #0"}, // Index 0
//		    {SomeInnerFields: "for_field #1"}, // Index 1
//		    {SomeInnerFields: "for_field #2"}, // Index 2
//		    {SomeInnerFields: "for_field #3"}, // Index 3
//		    {SomeInnerFields: "for_field #4"}, // Index 4
//		    {SomeInnerFields: "for_field #5"}, // Index 5
//		    {SomeInnerFields: "for_field #6"}, // Index 6
//		    {SomeInnerFields: "for_field #7"}, // Index 7
//		  },
//		}
//	 fm := mask.ParseMust("field.(1,5)")
//	 // Without reduceLists
//	 FilterWithSelectMaskAndReduceLists(msg, fm, false)
//	 msg == &SomeMessage{
//		  Field: []*OtherMessage{
//		    {}, // Index 0
//		    {SomeInnerFields: "for_field #1"}, // Index 1
//		    {}, // Index 2
//		    {}, // Index 3
//		    {}, // Index 4
//		    {SomeInnerFields: "for_field #5"}, // Index 5
//		    {}, // Index 6
//		    {}, // Index 7
//		  },
//		}
//	 // With reduceLists
//	 FilterWithSelectMaskAndReduceLists(msg, fm, true)
//	 msg == &SomeMessage{
//		  Field: []*OtherMessage{
//		    {SomeInnerFields: "for_field #1"},
//		    {SomeInnerFields: "for_field #5"},
//		  },
//		}
//
// Parameters:
//   - message [proto.Message]: The target message to be filtered.
//   - fieldMask *[mask.Mask]: The Mask specifying which fields to update or
//     reset.
//   - reduceLists (bool): Whether to remove elements from lists or just set to
//     zeros.
//
// Returns:
//   - error: An error, if any, encountered during the filtering process.
//
// [Design]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func FilterWithSelectMaskAndReduceLists(
	message proto.Message,
	fieldMask *mask.Mask,
	reduceLists bool,
) error {
	return recursiveFilterWithSelectMask(
		message.ProtoReflect(), fieldMask, reduceLists, 0,
	)
}
