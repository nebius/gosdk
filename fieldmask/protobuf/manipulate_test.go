package protobuf

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/nebius/gosdk/fieldmask/mask"
	"github.com/nebius/gosdk/fieldmask/protobuf/testdata"
)

func TestGetAtFieldPath(t *testing.T) {
	t.Parallel()
	veryRecursive := &testdata.RecursiveStruct{
		SomeString: "foo",
	}
	veryPath := mask.FieldPath{}
	for _ = range recursionTooDeep + 42 {
		veryRecursive = &testdata.RecursiveStruct{
			Recursive: veryRecursive,
		}
		veryPath = append(veryPath, "recursive")
	}
	cases := []struct {
		Message    proto.Message
		Path       mask.FieldPath
		Val        protoreflect.Value
		NoDesc     bool
		Err        string
		IsNotFound bool
	}{
		{
			Message: &testdata.TestOneOf{},
			Path:    mask.FieldPath{"test_oneof"},
			Val:     protoreflect.ValueOf(nil),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{},
			},
			Path: mask.FieldPath{"test_oneof"},
			Val:  protoreflect.ValueOfUint32(0),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 42,
				},
			},
			Path: mask.FieldPath{"test_oneof"},
			Val:  protoreflect.ValueOfUint32(42),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 42,
				},
			},
			Path: mask.FieldPath{"test_oneof", "abc"},
			Val:  protoreflect.ValueOfUint32(42),
			Err:  "test_oneof: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{},
			},
			Path: mask.FieldPath{"test_oneof"},
			Val:  protoreflect.ValueOf((*testdata.TestSimple)(nil).ProtoReflect()),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{},
			},
			Path: mask.FieldPath{"test_oneof", "abc"},
			Val:  protoreflect.ValueOf((*testdata.TestSimple)(nil).ProtoReflect()),
			Err:  "test_oneof: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_oneof"},
			Val: protoreflect.ValueOf((&testdata.TestSimple{
				TestInt32: 42,
			}).ProtoReflect()),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_oneof", "test_struct"},
			Val: protoreflect.ValueOf((&testdata.TestSimple{
				TestInt32: 42,
			}).ProtoReflect()),
			Err: "test_oneof: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestStructs{},
			Path:    mask.FieldPath{"test_struct"},
			Val:     protoreflect.ValueOf((*testdata.TestSimple)(nil).ProtoReflect()),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_oneof", "test_int32"},
			Val: protoreflect.ValueOf((&testdata.TestSimple{
				TestInt32: 42,
			}).ProtoReflect()),
			Err: "test_oneof: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestSimple{
				TestInt32: 42,
			},
			Path: mask.FieldPath{"test_int32"},
			Val:  protoreflect.ValueOf(int32(42)),
		},
		{
			Message: &testdata.TestSimple{
				TestInt32: 42,
			},
			Path: mask.FieldPath{"test_uint32"},
			Val:  protoreflect.ValueOf(uint32(0)),
		},
		{
			Message:    &testdata.TestSimple{},
			Path:       mask.FieldPath{"nonexistent"},
			Err:        "nonexistent: field name not found",
			IsNotFound: true,
		},
		{
			Message:    &testdata.TestSimple{},
			Path:       mask.FieldPath{"nonexistent", "foo"},
			Err:        "nonexistent: field name not found",
			IsNotFound: true,
		},
		{
			Message: &testdata.TestSimple{
				TestInt32: 42,
			},
			Path: mask.FieldPath{"test_int32", "foo"},
			Val:  protoreflect.ValueOf(int32(42)),
			Err:  "test_int32: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestSimple{
				TestInt32: 42,
			},
			Path: mask.FieldPath{},
			Val: protoreflect.ValueOf((&testdata.TestSimple{
				TestInt32: 42,
			}).ProtoReflect()),
			NoDesc: true,
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{42},
			},
			Path: mask.FieldPath{"test_int32", "0"},
			Val:  protoreflect.ValueOf(int32(42)),
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{42},
			},
			Path: mask.FieldPath{"test_uint32", "0"},
			Err:  "test_uint32[0]: out of list bounds",
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{42},
			},
			Path: mask.FieldPath{"test_int32", "1"},
			Err:  "test_int32[1]: out of list bounds",
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{42},
			},
			Path: mask.FieldPath{"test_int32", "abracadabra"},
			Err:  "test_int32.abracadabra: failed to convert to list index",
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{42},
			},
			Path: mask.FieldPath{"test_int32", "0", "foo"},
			Val:  protoreflect.ValueOf(int32(42)),
			Err:  "test_int32[0]: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestTypeMap{
				TestInt32: map[int32]string{
					42: "answer",
				},
			},
			Path: mask.FieldPath{"test_int32", "42"},
			Val:  protoreflect.ValueOf("answer"),
		},
		{
			Message: &testdata.TestTypeMap{
				TestInt32: map[int32]string{
					42: "answer",
				},
			},
			Path:       mask.FieldPath{"test_int32", "23"},
			Err:        "test_int32[23]: field name not found",
			IsNotFound: true,
		},
		{
			Message: &testdata.TestTypeMap{
				TestInt32: map[int32]string{
					42: "answer",
				},
			},
			Path: mask.FieldPath{"test_int32", "abracadabra"},
			Err:  "test_int32.abracadabra: failed to convert to map key: failed to convert to int32 map key: strconv.ParseInt: parsing \"abracadabra\": invalid syntax",
		},
		{
			Message: &testdata.TestTypeMap{
				TestInt32: map[int32]string{
					42: "answer",
				},
			},
			Path: mask.FieldPath{"test_int32", "42", "foo"},
			Val:  protoreflect.ValueOf("answer"),
			Err:  "test_int32[42]: found basic type in the middle of the path, can't descend deeper",
		},
		{
			Message: &testdata.TestStructs{},
			Path:    mask.FieldPath{"test_struct", "test_int32"},
			Val:     protoreflect.ValueOf(int32(0)),
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestInt32: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_repeated", "0", "test_int32"},
			Val:  protoreflect.ValueOf(int32(42)),
		},
		{
			Message: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"foo": {
						TestInt32: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_stringmap", "foo", "test_int32"},
			Val:  protoreflect.ValueOf(int32(42)),
		},
		{
			Message: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"foo": {
						TestInt32: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_stringmap", "foo"},
			Val: protoreflect.ValueOf((&testdata.TestSimple{
				TestInt32: 42,
			}).ProtoReflect()),
		},
		{
			Message: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{
					TestInt32: 42,
				},
			},
			Path: mask.FieldPath{"test_struct", "test_int32"},
			Val:  protoreflect.ValueOf(int32(42)),
		},
		{
			Message: proto.Clone(veryRecursive),
			Path:    veryPath,
			Err:     strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Path: mask.FieldPath{"field", "0"}.Join(veryPath...),
			Err:  "field[0]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Path: mask.FieldPath{"field", "foo"}.Join(veryPath...),
			Err:  "field[foo]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Path: mask.FieldPath{"field", "foo", "recursive", "recursive", "recursive", "recursive", "some_string"},
			Val:  protoreflect.ValueOf(""),
		},
		{
			Message: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Path:       mask.FieldPath{"field", "foo", "recursive", "recursive", "recursive", "recursive", "nonexistent"},
			Err:        "field[foo]: recursive: recursive: recursive: recursive: nonexistent: field name not found",
			IsNotFound: true,
		},
		{
			Message: &testdata.TestOneOf{},
			Path:    mask.FieldPath{"test_uint32"},
			Val:     protoreflect.ValueOf(nil),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Path: mask.FieldPath{"test_uint32"},
			Val:  protoreflect.ValueOf(nil),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Path: mask.FieldPath{"test_uint32"},
			Val:  protoreflect.ValueOf(nil),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Path: mask.FieldPath{"test_uint32"},
			Val:  protoreflect.ValueOf(nil),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{},
			},
			Path: mask.FieldPath{"test_uint32"},
			Val:  protoreflect.ValueOf(uint32(0)),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
			},
			Path: mask.FieldPath{"test_uint32"},
			Val:  protoreflect.ValueOf(uint32(42)),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
			},
			Path: mask.FieldPath{"test_struct", "test_int64"},
			Val:  protoreflect.ValueOf(int64(0)),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
			},
			Path: mask.FieldPath{"test_struct", "test_int64"},
			Val:  protoreflect.ValueOf(int64(0)),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
			},
			Path: mask.FieldPath{"test_struct", "test_int64"},
			Val:  protoreflect.ValueOf(int64(0)),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt64: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_struct", "test_int64"},
			Val:  protoreflect.ValueOf(int64(42)),
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt64: 42,
					},
				},
			},
			Path: mask.FieldPath{"test_string"},
			Val:  protoreflect.ValueOf(nil),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			a, merr1 := protojson.Marshal(c.Message)

			val, desc, err := GetAtFieldPath(c.Message, c.Path)

			b, merr2 := protojson.Marshal(c.Message)
			if merr1 != nil {
				assert.Equal(t, merr1.Error(), merr2.Error())
			} else {
				assert.JSONEq(t, string(a), string(b))
			}

			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
				assert.Equal(t, c.IsNotFound, errors.Is(err, ErrNotFound))
			} else {
				assert.NoError(t, err)
			}
			if c.Val.IsValid() {
				assert.True(t, val.IsValid(), "expecting valid value", c.Val.Interface())
				assert.Equal(t, c.Val.Interface(), val.Interface())
				if !c.NoDesc {
					assert.NotNil(t, desc)
				}
			} else {
				assert.False(t, val.IsValid(), "unexpected received val", val.Interface())
			}
		})
	}
}

func TestReplaceAtFieldPath(t *testing.T) {
	t.Parallel()
	t.Run("just call", func(t *testing.T) {
		t.Parallel()
		veryRecursive := &testdata.RecursiveStruct{
			SomeString: "foo",
		}
		veryRecursive2 := &testdata.RecursiveStruct{
			SomeString: "foo",
		}
		veryPath := mask.FieldPath{}
		for i := range recursionTooDeep + 42 {
			veryRecursive = &testdata.RecursiveStruct{
				Recursive: veryRecursive,
			}
			veryRecursive2 = &testdata.RecursiveStruct{
				Recursive: veryRecursive2,
			}
			if i == recursionTooDeep+42-3 {
				veryRecursive2.SomeString = "surprise"
			}
			veryPath = append(veryPath, "recursive")
		}
		cases := []struct {
			Message proto.Message
			Result  proto.Message
			Path    mask.FieldPath
			Val     interface{}
			Err     string
		}{
			{
				Message: &testdata.TestOptional{},
				Result:  &testdata.TestOptional{},
				Path:    nil,
				Val:     &testdata.TestOptional{},
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
				},
				Result: &testdata.TestSimple{
					TestInt32: 42,
				},
				Path: mask.FieldPath{"test_int32"},
				Val:  int32(42),
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
				},
				Result: &testdata.TestSimple{},
				Path:   mask.FieldPath{},
				Val:    nil,
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
				},
				Result: &testdata.TestSimple{
					TestInt32: 42,
				},
				Path: mask.FieldPath{"test_int32"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
				},
				Path: mask.FieldPath{"test_int32"},
				Val:  time.Hour,
				Err:  "failed to convert to protoreflect.Value: panic: invalid type: time.Duration",
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
				},
				Path: mask.FieldPath{"test_int32"},
				Val:  "abracadabra",
				Err:  "test_int32: failed to set value: panic: type mismatch: cannot convert string to int",
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 42,
				},
				Result: &testdata.TestSimple{
					TestInt32:  42,
					TestUint32: 42,
				},
				Path: mask.FieldPath{"test_uint32"},
				Val:  uint32(42),
			},
			{
				Message: &testdata.TestSimple{},
				Val:     uint32(42),
				Path:    mask.FieldPath{"nonexistent"},
				Err:     "nonexistent: field name not found",
			},
			{
				Message: &testdata.TestSimple{},
				Path:    mask.FieldPath{"nonexistent", "foo"},
				Err:     "nonexistent: field name not found",
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 42,
				},
				Path: mask.FieldPath{"test_int32", "foo"},
				Val:  int32(42),
				Err:  "test_int32: found basic type in the middle of the path, can't descend deeper",
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
					TestInt64: 0x7457E_F00D,
				},
				Path: mask.FieldPath{},
				Val: protoreflect.ValueOf((&testdata.TestSimple{
					TestUint32: 80085,
					TestInt32:  42,
				}).ProtoReflect()),
				Result: &testdata.TestSimple{
					TestInt32:  42,
					TestUint32: 80085,
				},
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
					TestInt64: 0x7457E_F00D,
				},
				Path: mask.FieldPath{},
				Val: (&testdata.TestSimple{
					TestUint32: 80085,
					TestInt32:  42,
				}).ProtoReflect(),
				Result: &testdata.TestSimple{
					TestInt32:  42,
					TestUint32: 80085,
				},
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
					TestInt64: 0x7457E_F00D,
				},
				Path: mask.FieldPath{},
				Val: &testdata.TestSimple{
					TestUint32: 80085,
					TestInt32:  42,
				},
				Result: &testdata.TestSimple{
					TestInt32:  42,
					TestUint32: 80085,
				},
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
					TestInt64: 0x7457E_F00D,
				},
				Path: mask.FieldPath{},
				Val:  &testdata.TestOptional{},
				Err:  "can't replace message nebius.testdata.TestSimple with message nebius.testdata.TestOptional",
			},
			{
				Message: &testdata.TestSimple{
					TestInt32: 23,
					TestInt64: 0x7457E_F00D,
				},
				Path: mask.FieldPath{},
				Val:  "str",
				Err:  "can't replace message nebius.testdata.TestSimple with value of type string",
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{32},
				},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{42},
				},
				Path: mask.FieldPath{"test_int32", "0"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{32},
				},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{32, 42},
				},
				Path: mask.FieldPath{"test_int32", "1"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{32},
				},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{},
				},
				Path: mask.FieldPath{"test_int32", "0"},
				Val:  nil,
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{1, 2, 3, 4},
				},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{2, 3, 4},
				},
				Path: mask.FieldPath{"test_int32", "0"},
				Val:  nil,
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{1, 2, 3, 4},
				},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{1, 3, 4},
				},
				Path: mask.FieldPath{"test_int32", "1"},
				Val:  nil,
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{1, 2, 3, 4},
				},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{1, 2, 3},
				},
				Path: mask.FieldPath{"test_int32", "3"},
				Val:  nil,
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{32},
				},
				Path: mask.FieldPath{"test_int32", "2"},
				Val:  protoreflect.ValueOf(int32(42)),
				Err:  "test_int32[2]: out of list bounds",
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{32},
				},
				Result: &testdata.TestRepeated{
					TestInt32:  []int32{32},
					TestUint32: []uint32{42},
				},
				Path: mask.FieldPath{"test_uint32", "0"},
				Val:  protoreflect.ValueOf(uint32(42)),
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{32},
				},
				Err:  "test_uint32[2]: out of list bounds",
				Path: mask.FieldPath{"test_uint32", "2"},
				Val:  protoreflect.ValueOf(uint32(42)),
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{22},
				},
				Path: mask.FieldPath{"test_int32", "abracadabra"},
				Val:  protoreflect.ValueOf(uint32(42)),
				Err:  "test_int32.abracadabra: failed to convert to list index",
			},
			{
				Message: &testdata.TestRepeated{
					TestInt32: []int32{42},
				},
				Path: mask.FieldPath{"test_int32", "0", "foo"},
				Val:  protoreflect.ValueOf(int32(42)),
				Err:  "test_int32[0]: found basic type in the middle of the path, can't descend deeper",
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "question",
					},
				},
				Result: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_int32", "42"},
				Val:  protoreflect.ValueOf("answer"),
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "question",
					},
				},
				Result: &testdata.TestTypeMap{
					TestInt32: map[int32]string{},
				},
				Path: mask.FieldPath{"test_int32", "42"},
				Val:  nil,
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42:      "answer",
						0x_F00D: "eat me",
					},
				},
				Result: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_int32", "61453"},
				Val:  nil,
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Result: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_int32", "61453"},
				Val:  nil,
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						7357: "test",
					},
				},
				Result: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						7357: "test",
						42:   "answer",
					},
				},
				Val:  protoreflect.ValueOf("answer"),
				Path: mask.FieldPath{"test_int32", "42"},
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_int32", "abracadabra"},
				Val:  protoreflect.ValueOf("foo"),
				Err:  "test_int32.abracadabra: failed to convert to map key: failed to convert to int32 map key: strconv.ParseInt: parsing \"abracadabra\": invalid syntax",
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						23: "foo",
					},
				},
				Result: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						23: "foo",
					},
					TestUint32: map[uint32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_uint32", "42"},
				Val:  protoreflect.ValueOf("answer"),
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_int32", "42", "foo"},
				Val:  protoreflect.ValueOf("answer"),
				Err:  "test_int32[42]: found basic type in the middle of the path, can't descend deeper",
			},
			{
				Message: &testdata.TestTypeMap{
					TestInt32: map[int32]string{
						42: "answer",
					},
				},
				Path: mask.FieldPath{"test_int32", "42"},
				Val:  int32(0x_BAD_F00D),
				Err:  "test_int32[42]: failed to set value: panic: interface conversion: interface {} is int32, not string",
			},
			{
				Message: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 32,
						},
					},
				},
				Result: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 42,
						},
					},
				},
				Path: mask.FieldPath{"test_repeated", "0", "test_int32"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 32,
						},
					},
				},
				Result: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 42,
						},
					},
				},
				Path: mask.FieldPath{"test_repeated", "0"},
				Val: &testdata.TestSimple{
					TestInt32: 42,
				},
			},
			{
				Message: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 32,
						},
					},
				},
				Result: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 32,
						},
						{
							TestInt32: 42,
						},
					},
				},
				Path: mask.FieldPath{"test_repeated", "1", "test_int32"},
				Val:  int32(42),
			},
			{
				Message: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 32,
						},
					},
				},
				Path: mask.FieldPath{"test_repeated", "1"},
				Val:  &testdata.TestOptional{},
				Err:  "test_repeated[1]: failed to set value: panic: invalid type: got *testdata.TestOptional, want *testdata.TestSimple",
			},
			{
				Message: &testdata.TestStructs{
					TestStringmap: map[string]*testdata.TestSimple{
						"foo": {
							TestInt32: 32,
						},
					},
				},
				Result: &testdata.TestStructs{
					TestStringmap: map[string]*testdata.TestSimple{
						"foo": {
							TestInt32: 42,
						},
					},
				},
				Path: mask.FieldPath{"test_stringmap", "foo", "test_int32"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: &testdata.TestStructs{
					TestStringmap: map[string]*testdata.TestSimple{
						"foo": {
							TestInt32: 32,
						},
					},
				},
				Result: &testdata.TestStructs{
					TestStringmap: map[string]*testdata.TestSimple{
						"foo": {
							TestInt32: 32,
						},
						"bar": {
							TestInt32: 42,
						},
					},
				},
				Path: mask.FieldPath{"test_stringmap", "bar", "test_int32"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32: 32,
					},
				},
				Result: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
				Path: mask.FieldPath{"test_struct", "test_int32"},
				Val:  protoreflect.ValueOf(int32(42)),
			},
			{
				Message: proto.Clone(veryRecursive),
				Path:    veryPath,
				Val:     protoreflect.ValueOf(int32(42)),
				Err:     strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
			},
			{
				Message: &testdata.TestRecursiveRepeated{
					Field: []*testdata.RecursiveStruct{
						proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
					},
				},
				Path: mask.FieldPath{"field", "0"}.Join(veryPath...),
				Val:  protoreflect.ValueOf(int32(42)),
				Err:  "field[0]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
			},
			{
				Message: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
					},
				},
				Path: mask.FieldPath{"field", "foo"}.Join(veryPath...),
				Val:  protoreflect.ValueOf(int32(42)),
				Err:  "field[foo]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
			},
			{
				Message: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
					},
				},
				Result: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": proto.Clone(veryRecursive2).(*testdata.RecursiveStruct),
					},
				},
				Path: mask.FieldPath{"field", "foo", "recursive", "recursive", "some_string"},
				Val:  protoreflect.ValueOf("surprise"),
			},
			{
				Message: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{},
				},
				Result: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": {
							Recursive: &testdata.RecursiveStruct{
								Recursive: &testdata.RecursiveStruct{
									SomeString: "yay",
								},
							},
						},
					},
				},
				Path: mask.FieldPath{"field", "foo", "recursive", "recursive", "some_string"},
				Val:  "yay",
			},
			{
				Message: &testdata.TestRecursiveMap{},
				Result: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": {
							Recursive: &testdata.RecursiveStruct{
								Recursive: &testdata.RecursiveStruct{
									SomeString: "yay",
								},
							},
						},
					},
				},
				Path: mask.FieldPath{"field", "foo", "recursive", "recursive", "some_string"},
				Val:  "yay",
			},
			{
				Message: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": {
							Recursive: &testdata.RecursiveStruct{
								Recursive: &testdata.RecursiveStruct{
									SomeString: "yay",
								},
							},
						},
					},
				},
				Result: &testdata.TestRecursiveMap{},
				Path:   mask.FieldPath{"field"},
				Val:    nil,
			},
			{
				Message: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": {
							Recursive: &testdata.RecursiveStruct{
								Recursive: &testdata.RecursiveStruct{
									SomeString: "yay",
								},
							},
						},
					},
				},
				Result: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{},
				},
				Path: mask.FieldPath{"field", "foo"},
				Val:  nil,
			},
			{
				Message: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": {
							Recursive: &testdata.RecursiveStruct{
								Recursive: &testdata.RecursiveStruct{
									SomeString: "yay",
								},
							},
						},
					},
				},
				Result: &testdata.TestRecursiveMap{
					Field: map[string]*testdata.RecursiveStruct{
						"foo": {},
					},
				},
				Path: mask.FieldPath{"field", "foo", "recursive"},
				Val:  nil,
			},
			{
				Message: &testdata.TestOneOf{},
				Result:  &testdata.TestOneOf{},
				Path:    mask.FieldPath{"test_oneof"},
				Err:     "test_oneof: field name not found",
			},
			{
				Message: &testdata.TestOneOf{},
				Result: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Path: mask.FieldPath{"test_uint32"},
				Val:  uint32(42),
			},
			{
				Message: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestString{TestString: "abc"},
				},
				Result: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Path: mask.FieldPath{"test_uint32"},
				Val:  uint32(42),
			},
			{
				Message: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestStruct{
						TestStruct: &testdata.TestSimple{
							TestInt32: 35,
						},
					},
				},
				Result: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Path: mask.FieldPath{"test_uint32"},
				Val:  uint32(42),
			},
			{
				Message: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 35},
				},
				Result: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestStruct{
						TestStruct: &testdata.TestSimple{
							TestInt32: 42,
						},
					},
				},
				Path: mask.FieldPath{"test_struct", "test_int32"},
				Val:  int32(42),
			},
			{
				Message: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Result: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Path: mask.FieldPath{"test_struct"},
				Val:  nil,
			},
			{
				Message: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Result: &testdata.TestOneOf{},
				Path:   mask.FieldPath{"test_uint32"},
				Val:    nil,
			},
			{
				Message: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 42},
				},
				Result: &testdata.TestOneOf{
					TestOneof: &testdata.TestOneOf_TestUint32{TestUint32: 0},
				},
				Path: mask.FieldPath{"test_uint32"},
				Val:  uint32(0),
			},
		}
		for i, c := range cases {
			t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
				t.Parallel()
				err := ReplaceAtFieldPath(c.Message, c.Path, c.Val)
				if c.Err != "" {
					assert.EqualError(t, err, c.Err)
				} else {
					assert.NoError(t, err)
					assert.EqualExportedValues(t, c.Result, c.Message)
				}
			})
		}
	})
	t.Run("with get", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			Source   proto.Message
			Dest     proto.Message
			Result   proto.Message
			Path     mask.FieldPath
			DestPath mask.FieldPath
		}{
			{
				Source: &testdata.TestSimple{},
				Dest:   &testdata.TestSimple{},
				Path:   mask.FieldPath{"test_int32"},
				Result: &testdata.TestSimple{},
			},
			{
				Source: &testdata.TestSimple{},
				Dest: &testdata.TestSimple{
					TestInt32: 42,
				},
				Path:   mask.FieldPath{"test_int32"},
				Result: &testdata.TestSimple{},
			},
			{
				Source: &testdata.TestSimple{
					TestInt32: 42,
				},
				Dest: &testdata.TestSimple{},
				Path: mask.FieldPath{"test_int32"},
				Result: &testdata.TestSimple{
					TestInt32: 42,
				},
			},
			{
				Source: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
						TestInt64: 0x_600D_F00D,
					},
				},
				Dest: &testdata.TestStructs{},
				Path: mask.FieldPath{"test_struct"},
				Result: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
						TestInt64: 0x_600D_F00D,
					},
				},
			},
			{
				Source: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
						TestInt64: 0x_600D_F00D,
					},
				},
				Dest: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32:  35,
						TestUint64: 0x_BAD_F00D,
					},
				},
				Path: mask.FieldPath{"test_struct"},
				Result: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
						TestInt64: 0x_600D_F00D,
					},
				},
			},
			{
				Source: &testdata.TestStructs{},
				Dest: &testdata.TestStructs{
					TestStruct: &testdata.TestSimple{
						TestInt32:  35,
						TestUint64: 0x_BAD_F00D,
					},
				},
				Path:   mask.FieldPath{"test_struct"},
				Result: &testdata.TestStructs{},
			},
			{
				Source: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 42,
							TestInt64: 0x_600D_F00D,
						},
					},
				},
				Dest: &testdata.TestStructs{},
				Path: mask.FieldPath{"test_repeated"},
				Result: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 42,
							TestInt64: 0x_600D_F00D,
						},
					},
				},
			},
			{
				Source: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 42,
							TestInt64: 0x_600D_F00D,
						},
						nil,
					},
				},
				Dest: &testdata.TestStructs{},
				Path: mask.FieldPath{"test_repeated"},
				Result: &testdata.TestStructs{
					TestRepeated: []*testdata.TestSimple{
						{
							TestInt32: 42,
							TestInt64: 0x_600D_F00D,
						},
					},
				},
			},
			{
				Source: &testdata.TestStructs{
					TestIntmap: map[int32]*testdata.TestSimple{
						42: {
							TestInt32: 42,
							TestInt64: 0x_600D_F00D,
						},
						23: nil,
					},
				},
				Dest: &testdata.TestStructs{},
				Path: mask.FieldPath{"test_intmap"},
				Result: &testdata.TestStructs{
					TestIntmap: map[int32]*testdata.TestSimple{
						42: {
							TestInt32: 42,
							TestInt64: 0x_600D_F00D,
						},
					},
				},
			},
			{
				Source: &testdata.TestRepeated{
					TestInt32: []int32{1, 2, 3},
				},
				Dest: &testdata.TestRepeated{},
				Path: mask.FieldPath{"test_int32"},
				Result: &testdata.TestRepeated{
					TestInt32: []int32{1, 2, 3},
				},
			},
			{
				Source: &testdata.TestStringMap{
					TestInt32: map[string]int32{"1": 1, "z": 2, "": 3},
				},
				Dest: &testdata.TestStringMap{},
				Path: mask.FieldPath{"test_int32"},
				Result: &testdata.TestStringMap{
					TestInt32: map[string]int32{"1": 1, "z": 2, "": 3},
				},
			},
		}
		for i, c := range cases {
			t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
				t.Parallel()
				if c.DestPath == nil {
					c.DestPath = c.Path
				}
				val, _, err := GetAtFieldPath(c.Source, c.Path)
				assert.NoError(t, err)
				err = ReplaceAtFieldPath(c.Dest, c.DestPath, val)
				assert.NoError(t, err)
				assert.EqualExportedValues(t, c.Result, c.Dest)
			})
		}
	})
}
