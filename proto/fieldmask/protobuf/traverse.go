package protobuf

import (
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
)

// TraverseFunc is called on each selected value while traversing a message by
// a field mask.
//
// Parameters:
//   - val: current selected value.
//   - desc: field descriptor of the current value. For container elements this
//     is the container field descriptor (list or map field).
//   - path: full path to current value from traversal root. For message fields
//     path looks like `test_struct.test_uint32`; for list/map elements it
//     includes element key/index as path part (e.g. `test_repeated.0`,
//     `test_intmap.7`).
//   - innerMask: nil for leaf mask elements, otherwise a sub-mask used for
//     descending deeper.
//     For example, for mask `test_struct.test_uint32`, callback for
//     `test_struct` receives non-nil innerMask, while callback for
//     `test_struct.test_uint32` receives nil.
//
// Return:
//   - bool: true to continue traversal, false to stop.
type TraverseFunc func(
	val protoreflect.Value,
	desc protoreflect.FieldDescriptor,
	path mask.FieldPath,
	innerMask *mask.Mask,
) bool

type traverseNode struct {
	val                protoreflect.Value
	desc               protoreflect.FieldDescriptor
	path               mask.FieldPath
	innerMask          *mask.Mask
	recursion          int
	isContainerElement bool
}

func getTraverseSubMask(
	fieldMask *mask.Mask,
	key mask.FieldKey,
) (*mask.Mask, error) {
	if fieldMask == nil || fieldMask.IsEmpty() {
		return mask.New(), nil
	}
	return fieldMask.GetSubMask(key)
}

func toTraverseInnerMask(innerMask *mask.Mask) *mask.Mask {
	if innerMask == nil || innerMask.IsEmpty() {
		return nil
	}
	return innerMask
}

func collectTraverseMessageChildren(
	message protoreflect.Message,
	fieldMask *mask.Mask,
	recursion int,
	path mask.FieldPath,
) ([]traverseNode, error) {
	if recursion >= recursionTooDeep {
		return nil, mask.ErrRecursionTooDeep
	}
	recursion++
	msgDesc := message.Descriptor()
	ret := make([]traverseNode, 0, msgDesc.Fields().Len())

	for i := range msgDesc.Fields().Len() {
		fieldDesc := msgDesc.Fields().Get(i)
		if !message.Has(fieldDesc) {
			continue
		}
		fieldName := string(fieldDesc.Name())
		fieldPath := path.Join(mask.FieldKey(fieldName))
		innerMask, err := getTraverseSubMask(
			fieldMask, mask.FieldKey(fieldDesc.Name()),
		)
		if err != nil {
			return nil, fmt.Errorf(
				"%s: couldn't get full field mask: %w", fieldPath.String(), err,
			)
		}
		if innerMask == nil && fieldDesc.ContainingOneof() != nil &&
			fieldMask != nil && !fieldMask.IsEmpty() {
			oneOfName := fieldDesc.ContainingOneof().Name()
			oneOfPath := path.Join(mask.FieldKey(oneOfName))
			oneOfMask, err := fieldMask.GetSubMask(mask.FieldKey(oneOfName))
			if err != nil {
				return nil, fmt.Errorf(
					"%s: couldn't get full field mask: %w", oneOfPath.String(), err,
				)
			}
			if oneOfMask != nil {
				// OneOf group names select a concrete field as a leaf.
				innerMask = mask.New()
			}
		}
		if innerMask == nil {
			continue
		}
		ret = append(ret, traverseNode{
			val:       message.Get(fieldDesc),
			desc:      fieldDesc,
			path:      fieldPath,
			innerMask: innerMask,
			recursion: recursion,
		})
	}

	return ret, nil
}

func expandTraverseListElement(node traverseNode) ([]traverseNode, error) {
	if node.desc.Message() == nil {
		return nil, fmt.Errorf(
			"%s: %w", node.path.String(), ErrBasicTypeInTheMiddle,
		)
	}
	return collectTraverseMessageChildren(
		node.val.Message(),
		node.innerMask,
		node.recursion,
		node.path,
	)
}

func expandTraverseListField(node traverseNode) ([]traverseNode, error) {
	list := node.val.List()
	ret := make([]traverseNode, 0, list.Len())
	for i := range list.Len() {
		innerMask, err := getTraverseSubMask(
			node.innerMask, mask.FieldKey(strconv.Itoa(i)),
		)
		if err != nil {
			return nil, fmt.Errorf(
				"%s[%d]: couldn't get full field mask: %w",
				node.path.String(), i, err,
			)
		}
		if innerMask == nil {
			continue
		}
		childPath := node.path.Join(mask.FieldKey(strconv.Itoa(i)))
		ret = append(ret, traverseNode{
			val:                list.Get(i),
			desc:               node.desc,
			path:               childPath,
			innerMask:          innerMask,
			recursion:          node.recursion,
			isContainerElement: true,
		})
	}
	return ret, nil
}

func expandTraverseMapElement(node traverseNode) ([]traverseNode, error) {
	if node.desc.MapValue().Message() == nil {
		return nil, fmt.Errorf(
			"%s: %w", node.path.String(), ErrBasicTypeInTheMiddle,
		)
	}
	return collectTraverseMessageChildren(
		node.val.Message(),
		node.innerMask,
		node.recursion,
		node.path,
	)
}

func expandTraverseMapField(node traverseNode) ([]traverseNode, error) {
	fMap := node.val.Map()
	ret := make([]traverseNode, 0, fMap.Len())
	var innerErr error
	fMap.Range(func(mk protoreflect.MapKey, v protoreflect.Value) bool {
		fk, err := mapKeyToFieldKey(mk)
		if err != nil {
			innerErr = fmt.Errorf(
				"%s[%s]: failed to convert map key to field key: %w",
				node.path.String(), mk, err,
			)
			return false
		}
		innerMask, err := getTraverseSubMask(node.innerMask, fk)
		if err != nil {
			innerErr = fmt.Errorf(
				"%s[%s]: couldn't get full field mask: %w",
				node.path.String(), mk, err,
			)
			return false
		}
		if innerMask == nil {
			return true
		}
		childPath := node.path.Join(fk)
		ret = append(ret, traverseNode{
			val:                v,
			desc:               node.desc,
			path:               childPath,
			innerMask:          innerMask,
			recursion:          node.recursion,
			isContainerElement: true,
		})
		return true
	})
	if innerErr != nil {
		return nil, innerErr
	}
	return ret, nil
}

func expandTraverseNode(node traverseNode) ([]traverseNode, error) {
	if node.innerMask == nil || node.innerMask.IsEmpty() {
		return nil, nil
	}
	if node.desc.IsList() {
		if node.isContainerElement {
			return expandTraverseListElement(node)
		}
		return expandTraverseListField(node)
	}
	if node.desc.IsMap() {
		if node.isContainerElement {
			return expandTraverseMapElement(node)
		}
		return expandTraverseMapField(node)
	}
	if node.desc.Message() != nil {
		return collectTraverseMessageChildren(
			node.val.Message(),
			node.innerMask,
			node.recursion,
			node.path,
		)
	}
	return nil, fmt.Errorf("%s: %w", node.path.String(), ErrBasicTypeInTheMiddle)
}

func traverseDFNode(node traverseNode, cb TraverseFunc) (bool, error) {
	if !cb(
		node.val,
		node.desc,
		node.path.Copy(),
		toTraverseInnerMask(node.innerMask),
	) {
		return false, nil
	}
	children, err := expandTraverseNode(node)
	if err != nil {
		return false, err
	}
	for _, child := range children {
		ok, err := traverseDFNode(child, cb)
		if err != nil || !ok {
			return ok, err
		}
	}
	return true, nil
}

func traverseRoot(
	message proto.Message,
	fieldMask *mask.Mask,
) ([]traverseNode, error) {
	if message == nil {
		return nil, errors.New("received nil")
	}
	return collectTraverseMessageChildren(
		message.ProtoReflect(), fieldMask, 0, mask.FieldPath{},
	)
}

// TraverseDF traverses message values selected by fieldMask in depth-first
// order (parent, then children).
//
// Parameters:
//   - message: source message to traverse.
//   - fieldMask: select mask describing which fields/branches to visit.
//     nil or empty mask means "visit all set fields".
//   - cb: callback called for every visited selected value.
//
// Example:
//
//	msg := &testdata.TestStructs{
//		TestStruct: &testdata.TestSimple{TestUint32: 42},
//	}
//	fm := mask.ParseMust("test_struct.test_uint32")
//	ok, err := TraverseDF(msg, fm, func(
//		val protoreflect.Value,
//		desc protoreflect.FieldDescriptor,
//		path mask.FieldPath,
//		innerMask *mask.Mask,
//	) bool {
//		fmt.Println(path.String(), desc.Name(), innerMask != nil)
//		return true
//	})
//	// Output order:
//	// test_struct test_struct true
//	// test_struct.test_uint32 test_uint32 false
//
// It returns:
//   - bool: true if traversal passed fully, false if callback requested stop.
//   - error: traversal error.
func TraverseDF(
	message proto.Message,
	fieldMask *mask.Mask,
	cb TraverseFunc,
) (bool, error) {
	if cb == nil {
		return false, errors.New("received nil callback")
	}
	root, err := traverseRoot(message, fieldMask)
	if err != nil {
		return false, err
	}
	for _, node := range root {
		ok, err := traverseDFNode(node, cb)
		if err != nil || !ok {
			return ok, err
		}
	}
	return true, nil
}

// TraverseBF traverses message values selected by fieldMask in breadth-first
// order (all current level values first, then next level).
//
// Parameters:
//   - message: source message to traverse.
//   - fieldMask: select mask describing which fields/branches to visit.
//     nil or empty mask means "visit all set fields".
//   - cb: callback called for every visited selected value.
//
// Example:
//
//	msg := &testdata.TestStructs{
//		TestStruct: &testdata.TestSimple{TestUint32: 42},
//		TestRepeated: []*testdata.TestSimple{{TestBool: true}},
//	}
//	fm := mask.ParseMust("test_struct.test_uint32,test_repeated.0.test_bool")
//	ok, err := TraverseBF(msg, fm, func(
//		_ protoreflect.Value,
//		_ protoreflect.FieldDescriptor,
//		path mask.FieldPath,
//		_ *mask.Mask,
//	) bool {
//		fmt.Println(path.String())
//		return true
//	})
//	// One possible order:
//	// test_struct
//	// test_repeated
//	// test_struct.test_uint32
//	// test_repeated.0
//	// test_repeated.0.test_bool
//
// It returns:
//   - bool: true if traversal passed fully, false if callback requested stop.
//   - error: traversal error.
func TraverseBF(
	message proto.Message,
	fieldMask *mask.Mask,
	cb TraverseFunc,
) (bool, error) {
	if cb == nil {
		return false, errors.New("received nil callback")
	}
	queue, err := traverseRoot(message, fieldMask)
	if err != nil {
		return false, err
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !cb(
			node.val,
			node.desc,
			node.path.Copy(),
			toTraverseInnerMask(node.innerMask),
		) {
			return false, nil
		}
		children, err := expandTraverseNode(node)
		if err != nil {
			return false, err
		}
		queue = append(queue, children...)
	}
	return true, nil
}
