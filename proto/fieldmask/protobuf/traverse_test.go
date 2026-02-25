package protobuf

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/fieldmask/protobuf/testdata"
)

func TestTraverseOrder(t *testing.T) {
	t.Parallel()
	msg := &testdata.TestStructs{
		TestStruct: &testdata.TestSimple{
			TestUint32: 42,
		},
		TestRepeated: []*testdata.TestSimple{
			{
				TestBool: true,
			},
		},
	}
	fm := mask.ParseMust("test_struct.test_uint32,test_repeated.0.test_bool")
	toEvent := func(
		_ protoreflect.Value,
		desc protoreflect.FieldDescriptor,
		path mask.FieldPath,
		innerMask *mask.Mask,
	) string {
		kind := "leaf"
		if innerMask != nil {
			kind = "inner"
		}
		return fmt.Sprintf("%s|%s|%s", desc.Name(), path.String(), kind)
	}

	dfOrder := make([]string, 0, 5)
	ok, err := TraverseDF(msg, fm, func(
		val protoreflect.Value,
		desc protoreflect.FieldDescriptor,
		path mask.FieldPath,
		innerMask *mask.Mask,
	) bool {
		dfOrder = append(dfOrder, toEvent(val, desc, path, innerMask))
		return true
	})
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"test_struct|test_struct|inner",
		"test_uint32|test_struct.test_uint32|leaf",
		"test_repeated|test_repeated|inner",
		"test_repeated|test_repeated.0|inner",
		"test_bool|test_repeated.0.test_bool|leaf",
	}, dfOrder)

	bfOrder := make([]string, 0, 5)
	ok, err = TraverseBF(msg, fm, func(
		val protoreflect.Value,
		desc protoreflect.FieldDescriptor,
		path mask.FieldPath,
		innerMask *mask.Mask,
	) bool {
		bfOrder = append(bfOrder, toEvent(val, desc, path, innerMask))
		return true
	})
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"test_struct|test_struct|inner",
		"test_repeated|test_repeated|inner",
		"test_uint32|test_struct.test_uint32|leaf",
		"test_repeated|test_repeated.0|inner",
		"test_bool|test_repeated.0.test_bool|leaf",
	}, bfOrder)
}

func TestTraverseFieldPath(t *testing.T) {
	t.Parallel()
	msg := &testdata.TestStructs{
		TestRepeated: []*testdata.TestSimple{
			{
				TestUint32: 1,
			},
		},
		TestIntmap: map[int32]*testdata.TestSimple{
			7: {
				TestUint32: 2,
			},
		},
	}
	fm := mask.ParseMust("test_repeated.0.test_uint32,test_intmap.7.test_uint32")

	var listSeen, mapSeen bool
	ok, err := TraverseDF(msg, fm, func(
		_ protoreflect.Value,
		desc protoreflect.FieldDescriptor,
		path mask.FieldPath,
		innerMask *mask.Mask,
	) bool {
		switch {
		case desc.Name() == "test_repeated" && len(path) == 2:
			assert.NotNil(t, innerMask)
			assert.True(t, path.Equal(mask.NewFieldPath("test_repeated", "0")))
			listSeen = true
		case desc.Name() == "test_intmap" && len(path) == 2:
			assert.NotNil(t, innerMask)
			assert.True(t, path.Equal(mask.NewFieldPath("test_intmap", "7")))
			mapSeen = true
		}
		return true
	})
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.True(t, listSeen)
	assert.True(t, mapSeen)
}

func TestTraversePartial(t *testing.T) {
	t.Parallel()

	type traverseEntry struct {
		Name string
		Fn   func(proto.Message, *mask.Mask, TraverseFunc) (bool, error)
	}
	traversers := []traverseEntry{
		{
			Name: "DF",
			Fn:   TraverseDF,
		},
		{
			Name: "BF",
			Fn:   TraverseBF,
		},
	}

	cases := []struct {
		Name     string
		Message  proto.Message
		Mask     *mask.Mask
		Expected []string
	}{
		{
			Name: "skip_unmasked_message_fields",
			Message: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{
					TestUint32: 42,
					TestString: "extra",
				},
				TestRepeated: []*testdata.TestSimple{
					{
						TestBool: true,
					},
				},
			},
			Mask: mask.ParseMust("test_struct.test_uint32"),
			Expected: []string{
				"test_struct",
				"test_struct.test_uint32",
			},
		},
		{
			Name: "inner_mask_fields_missing_in_message",
			Message: &testdata.TestStructs{
				TestStruct:   &testdata.TestSimple{},
				TestRepeated: []*testdata.TestSimple{{TestBool: true}},
				TestIntmap: map[int32]*testdata.TestSimple{
					7: {
						TestUint32: 2,
					},
				},
			},
			Mask: mask.ParseMust(
				"test_struct.test_uint32,test_repeated.2.test_bool,test_intmap.9.test_uint32",
			),
			Expected: []string{
				"test_struct",
				"test_repeated",
				"test_intmap",
			},
		},
	}

	for _, tc := range cases {
		for _, tr := range traversers {
			t.Run(tc.Name+"_"+tr.Name, func(t *testing.T) {
				t.Parallel()
				visited := make([]string, 0)
				ok, err := tr.Fn(tc.Message, tc.Mask, func(
					_ protoreflect.Value,
					_ protoreflect.FieldDescriptor,
					path mask.FieldPath,
					_ *mask.Mask,
				) bool {
					visited = append(visited, path.String())
					return true
				})
				assert.True(t, ok)
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, visited)
			})
		}
	}
}

func TestTraverseEarlyStop(t *testing.T) {
	t.Parallel()
	msg := &testdata.TestStructs{
		TestStruct: &testdata.TestSimple{
			TestUint32: 42,
		},
		TestRepeated: []*testdata.TestSimple{
			{TestBool: true},
		},
	}
	fm := mask.ParseMust("test_struct.test_uint32,test_repeated.0.test_bool")

	cases := []struct {
		Name string
		Fn   func(protoMessage proto.Message, fieldMask *mask.Mask, cb TraverseFunc) (bool, error)
	}{
		{
			Name: "DF",
			Fn:   TraverseDF,
		},
		{
			Name: "BF",
			Fn:   TraverseBF,
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			calls := 0
			ok, err := tc.Fn(msg, fm, func(
				_ protoreflect.Value,
				_ protoreflect.FieldDescriptor,
				_ mask.FieldPath,
				_ *mask.Mask,
			) bool {
				calls++
				return false
			})
			assert.NoError(t, err)
			assert.False(t, ok)
			assert.Equal(t, 1, calls)
		})
	}
}

func TestTraverseErrors(t *testing.T) {
	t.Parallel()

	veryRecursive := &testdata.RecursiveStruct{}
	veryMask := &mask.Mask{}
	for range recursionTooDeep + 42 {
		veryRecursive = &testdata.RecursiveStruct{
			Recursive: veryRecursive,
		}
		veryMask = &mask.Mask{
			FieldParts: map[mask.FieldKey]*mask.Mask{
				"recursive": veryMask,
			},
		}
	}

	cases := []struct {
		Name    string
		Fn      func(message proto.Message, fieldMask *mask.Mask, cb TraverseFunc) (bool, error)
		Message proto.Message
		Mask    *mask.Mask
		ErrIs   error
		ErrLike string
	}{
		{
			Name:    "DF recursion",
			Fn:      TraverseDF,
			Message: veryRecursive,
			Mask:    veryMask,
			ErrIs:   mask.ErrRecursionTooDeep,
		},
		{
			Name:    "BF recursion",
			Fn:      TraverseBF,
			Message: veryRecursive,
			Mask:    veryMask,
			ErrIs:   mask.ErrRecursionTooDeep,
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			ok, err := tc.Fn(tc.Message, tc.Mask, func(
				_ protoreflect.Value,
				_ protoreflect.FieldDescriptor,
				_ mask.FieldPath,
				_ *mask.Mask,
			) bool {
				return true
			})
			assert.False(t, ok)
			assert.ErrorIs(t, err, tc.ErrIs)
			if tc.ErrLike != "" {
				assert.ErrorContains(t, err, tc.ErrLike)
			}
		})
	}

	scalarMsg := &testdata.TestSimple{TestInt32: 7}
	scalarMask := mask.ParseMust("test_int32.foo")
	dfsOK, dfsErr := TraverseDF(scalarMsg, scalarMask, func(
		_ protoreflect.Value,
		_ protoreflect.FieldDescriptor,
		_ mask.FieldPath,
		_ *mask.Mask,
	) bool {
		return true
	})
	assert.False(t, dfsOK)
	assert.True(t, errors.Is(dfsErr, ErrBasicTypeInTheMiddle))
	assert.ErrorContains(t, dfsErr, "test_int32")

	bfsOK, bfsErr := TraverseBF(scalarMsg, scalarMask, func(
		_ protoreflect.Value,
		_ protoreflect.FieldDescriptor,
		_ mask.FieldPath,
		_ *mask.Mask,
	) bool {
		return true
	})
	assert.False(t, bfsOK)
	assert.True(t, errors.Is(bfsErr, ErrBasicTypeInTheMiddle))
	assert.ErrorContains(t, bfsErr, "test_int32")
}

func TestTraverseNilInput(t *testing.T) {
	t.Parallel()

	msg := &testdata.TestSimple{}
	fm := mask.New()

	ok, err := TraverseDF(nil, fm, func(
		_ protoreflect.Value,
		_ protoreflect.FieldDescriptor,
		_ mask.FieldPath,
		_ *mask.Mask,
	) bool {
		return true
	})
	assert.False(t, ok)
	assert.EqualError(t, err, "received nil")

	ok, err = TraverseBF(msg, fm, nil)
	assert.False(t, ok)
	assert.EqualError(t, err, "received nil callback")

	ok, err = TraverseBF(msg, fm, func(
		_ protoreflect.Value,
		_ protoreflect.FieldDescriptor,
		_ mask.FieldPath,
		_ *mask.Mask,
	) bool {
		return true
	})
	assert.True(t, ok)
	assert.NoError(t, err)
}
