package protobuf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/fieldmask/protobuf/testdata"
)

func TestFilterWithSelectMask(t *testing.T) {
	t.Parallel()
	veryRecursive := &testdata.RecursiveStruct{}
	veryMask := &mask.Mask{}
	for _ = range recursionTooDeep + 42 {
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
		Message proto.Message
		Mask    *mask.Mask
		Reduce  bool
		Err     string
		Result  proto.Message
	}{
		{
			Message: &testdata.TestSimple{},
			Result:  &testdata.TestSimple{},
		},
		{
			Message: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
		},
		{
			Message: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
			Mask: mask.ParseMust(""),
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
		},
		{
			Message: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
			Mask: mask.ParseMust("*"),
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
		},
		{
			Message: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
			Mask: mask.ParseMust("test_int32"),
			Result: &testdata.TestSimple{
				TestInt32: 42,
			},
		},
		{
			Message: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
			Mask: mask.ParseMust("test_int32,test_uint32"),
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 0x_D06_F00D,
			},
		},
		{
			Message: &testdata.TestSimple{
				TestInt32:  42,
				TestInt64:  80085,
				TestUint32: 0x_D06_F00D,
			},
			Mask: mask.ParseMust("test_int32,test_uint32,test_uint64"),
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 0x_D06_F00D,
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask: mask.ParseMust("test_int32"),
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask: mask.ParseMust("test_int32.*,test_int64"),
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask:   mask.ParseMust("test_int32.*,test_int64"),
			Reduce: true,
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4, 0},
			},
			Mask:   mask.ParseMust("test_int32.*,test_int64"),
			Reduce: true,
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4, 0},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask: mask.ParseMust("test_int32.1"),
			Result: &testdata.TestRepeated{
				TestInt32: []int32{0, 2, 0, 0},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask: mask.ParseMust("test_int32.(1,3)"),
			Result: &testdata.TestRepeated{
				TestInt32: []int32{0, 2, 0, 4},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask:   mask.ParseMust("test_int32.1"),
			Reduce: true,
			Result: &testdata.TestRepeated{
				TestInt32: []int32{2},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4},
			},
			Mask:   mask.ParseMust("test_int32.(1,3)"),
			Reduce: true,
			Result: &testdata.TestRepeated{
				TestInt32: []int32{2, 4},
			},
		},
		{
			Message: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3, 4, 0},
			},
			Mask:   mask.ParseMust("test_int32.(1,4)"),
			Reduce: true,
			Result: &testdata.TestRepeated{
				TestInt32: []int32{2, 0},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
			},
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
			},
			Mask: mask.ParseMust("*.(meat,1)"),
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
			},
			Mask: mask.ParseMust("test_stringmap"),
			Result: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					"empty": {},
				},
			},
			Mask:   mask.ParseMust("*.(meat,1)"),
			Reduce: true,
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					"empty": {},
				},
			},
			Mask:   mask.ParseMust("*.(meat,1,empty,3)"),
			Reduce: true,
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"empty": {},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					"empty": {},
				},
			},
			Mask:   mask.ParseMust("*.(meat,1,empty,3,nonexistent,42)"),
			Reduce: true,
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"empty": {},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					"empty": {},
				},
			},
			Mask:   mask.ParseMust("*.(meat,1).test_string"),
			Reduce: true,
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "meat",
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"meat": {
						TestString: "meat",
					},
				},
			},
		},
		{
			Message: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
						TestUint32: 42,
					},
					{
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					{
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
						TestUint32: 42,
					},
					"meat": {
						TestString: "meat",
						TestUint32: 0x_D06_F00D,
					},
					"test": {
						TestString: "test",
						TestUint32: 0x_7E57,
					},
					"empty": {},
				},
			},
			Mask:   mask.ParseMust("*.*.test_string"),
			Reduce: true,
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestString: "answer",
					},
					{
						TestString: "meat",
					},
					{
						TestString: "test",
					},
					{},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"answer": {
						TestString: "answer",
					},
					"meat": {
						TestString: "meat",
					},
					"test": {
						TestString: "test",
					},
					"empty": {},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask: mask.ParseMust("test_struct"),
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask: mask.ParseMust("test_oneof"),
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask: mask.ParseMust("test_oneof.test_string"),
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask: mask.ParseMust("test_struct.test_string"),
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask: mask.ParseMust("*.test_string"),
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask: mask.ParseMust("*"),
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestString: "answer",
						TestInt32:  42,
					},
				},
			},
			Mask:   mask.ParseMust("test_uint32"),
			Result: &testdata.TestOneOf{},
		},
		{
			Message: proto.Clone(veryRecursive),
			Mask:    veryMask,
			Err:     strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Mask: &mask.Mask{
				Any: &mask.Mask{
					Any: veryMask,
				},
			},
			Err: "field[0]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Mask: &mask.Mask{
				Any: &mask.Mask{
					Any: veryMask,
					FieldParts: map[mask.FieldKey]*mask.Mask{
						"0": veryMask,
					},
				},
			},
			Err: "field[0]: couldn't get full field mask: failed to copy key mask: " +
				strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Mask: &mask.Mask{
				Any: &mask.Mask{
					Any: veryMask,
				},
			},
			Err: "field[foo]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Message: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Mask: &mask.Mask{
				Any: &mask.Mask{
					Any: veryMask,
					FieldParts: map[mask.FieldKey]*mask.Mask{
						"foo": veryMask,
					},
				},
			},
			Err: "field[foo]: couldn't get full field mask: failed to copy key mask: " +
				strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
		},
		{
			Message: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{TestString: "foo"},
			},
			Mask: &mask.Mask{
				Any: veryMask,
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_string": veryMask,
				},
			},
			Err: "test_string: couldn't get full field mask: failed to copy key mask: " +
				strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d FilterWithSelectMaskAndReduceLists", i), func(t *testing.T) {
			t.Parallel()
			msg := proto.Clone(c.Message)
			err := FilterWithSelectMaskAndReduceLists(
				msg, c.Mask, c.Reduce,
			)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				assert.EqualExportedValues(t, c.Result, msg)
			}
		})
		if !c.Reduce {
			t.Run(fmt.Sprintf("case %d FilterWithSelectMask", i), func(t *testing.T) {
				t.Parallel()
				msg := proto.Clone(c.Message)
				err := FilterWithSelectMask(
					msg, c.Mask,
				)
				if c.Err != "" {
					assert.EqualError(t, err, c.Err)
				} else {
					assert.NoError(t, err)
					assert.EqualExportedValues(t, c.Result, msg)
				}
			})
		}
	}
}
