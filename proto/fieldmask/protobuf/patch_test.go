package protobuf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/fieldmask/protobuf/testdata"
)

func TestPatchWithResetMask(t *testing.T) {
	t.Parallel()
	infMask := mask.New()
	infMask.Any = infMask
	veryRecursive := &testdata.RecursiveStruct{
		SomeString: "foo",
	}
	veryRecursive2 := &testdata.RecursiveStruct{
		SomeString: "bar",
	}
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
	}
	cases := []struct {
		Data   proto.Message
		Patch  proto.Message
		Mask   *mask.Mask
		Err    string
		Result proto.Message
	}{
		{
			Data: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 35,
			},
			Patch: &testdata.TestSimple{
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Mask: nil,
		},
		{
			Data: &testdata.TestSimple{
				TestInt32: 42,
			},
			Patch: &testdata.TestSimple{
				TestInt32: 0,
			},
			Result: &testdata.TestSimple{
				TestInt32: 42,
			},
			Mask: nil,
		},
		{
			Data: &testdata.TestSimple{
				TestInt32: 42,
			},
			Patch: &testdata.TestSimple{
				TestInt32: 0,
			},
			Result: &testdata.TestSimple{
				TestInt32: 0,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 35,
			},
			Patch: &testdata.TestSimple{
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Result: &testdata.TestSimple{
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 35,
			},
			Patch: &testdata.TestSimple{
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_uint32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 35,
			},
			Patch: &testdata.TestSimple{
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Result: &testdata.TestSimple{
				TestInt32:  42,
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_uint64": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestSimple{
				TestInt32:  42,
				TestUint32: 35,
			},
			Patch: &testdata.TestSimple{
				TestUint64: 7357,
				TestUint32: 80085,
			},
			Err: "test_uint32: couldn't get full field mask: failed to copy key mask: " +
				strings.Repeat("*: ", recursionTooDeep) + "recursion too deep",
			Mask: &mask.Mask{
				Any: infMask,
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_uint32": infMask,
				},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Patch: &testdata.TestRepeated{},
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Mask: mask.New(),
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Patch: &testdata.TestRepeated{
				TestInt32: []int32{},
			},
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Mask: mask.New(),
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Patch:  &testdata.TestRepeated{},
			Result: &testdata.TestRepeated{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Patch: &testdata.TestRepeated{
				TestInt32: []int32{},
			},
			Result: &testdata.TestRepeated{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2},
			},
			Patch: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2},
			},
			Patch: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2},
			},
			Patch: &testdata.TestRepeated{
				TestInt32: []int32{1},
			},
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2},
			},
			Patch: &testdata.TestRepeated{
				TestInt32: []int32{1},
			},
			Result: &testdata.TestRepeated{
				TestInt32: []int32{1},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 6,
					"d": 4,
					"e": 5,
				},
			},
			Result: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 6,
					"d": 4,
					"e": 5,
				},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 6,
					"d": 4,
					"e": 5,
				},
			},
			Result: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 6,
					"d": 4,
					"e": 5,
				},
			},
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch: &testdata.TestStringMap{
				TestInt32: map[string]int32{},
			},
			Result: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch: &testdata.TestStringMap{},
			Result: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch: &testdata.TestStringMap{
				TestInt32: map[string]int32{},
			},
			Result: &testdata.TestStringMap{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch:  &testdata.TestStringMap{},
			Result: &testdata.TestStringMap{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int32": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestInt32:   123,
						TestInt64:   456,
						TestUint32:  69,
						TestUint64:  43,
						TestSint32:  80085,
						TestSint64:  0x_BAD_F00D,
						TestFixed32: 987,
						TestFixed64: 654,
					},
					{
						TestInt32:   124,
						TestInt64:   457,
						TestUint32:  690,
						TestUint64:  42,
						TestSint32:  80086,
						TestSint64:  0x_BAD_F00D,
						TestFixed32: 988,
						TestFixed64: 655,
					},
					{ // should be removed
						TestInt32: 0x_BAD_F00D,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"b": {
						TestInt32:   125,
						TestInt64:   458,
						TestUint32:  691,
						TestUint64:  44,
						TestSint32:  80087,
						TestSint64:  0x_BAD_F00D,
						TestFixed32: 989,
						TestFixed64: 656,
					},
					"1": {
						TestInt32:   126,
						TestInt64:   459,
						TestUint32:  692,
						TestUint64:  45,
						TestSint32:  80088,
						TestSint64:  0x_BAD_F00D,
						TestFixed32: 980,
						TestFixed64: 657,
					},
					"c": { // should be removed
						TestInt32: 0x_BAD_F00D,
					},
				},
			},
			Patch: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						// TestInt32: , // left as is
						TestInt64: 789, // changed
						// TestUint32 // removed (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						// TestFixed32 // left (tsm.*)
						// TestFixed64 // removed in b
					},
					{
						// TestInt32: , // left as is
						TestInt64: 780, // changed
						// TestUint32 // removed (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						// TestFixed32 // left (tsm.*)
						// TestFixed64 // removed in b
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"b": {
						// TestInt32: , // left as is
						TestInt64: 781, // changed
						// TestUint32 // left (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						// TestFixed32 // removed (tsm.*)
						// TestFixed64 // removed in b
					},
					"1": {
						// TestInt32: , // left as is
						TestInt64: 782, // changed
						// TestUint32 // left (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						// TestFixed32 // removed (tsm.*)
						// TestFixed64 // removed in b
					},
					"d": { // should be added
						TestInt32: 0x_D06_F00D,
					},
				},
			},
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestInt32: 123, // left as is
						TestInt64: 789, // changed
						// TestUint32 // removed (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						TestSint32: 80085, // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						TestFixed32: 987, // left (tsm.*)
						TestFixed64: 654, // removed in b
					},
					{
						TestInt32: 124, // left as is
						TestInt64: 780, // changed
						// TestUint32 // removed (test_repeated.*)
						TestUint64: 42, // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						TestFixed32: 988, // left (tsm.*)
						TestFixed64: 655, // removed in b
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"b": {
						TestInt32:  125,   // left as is
						TestInt64:  781,   // changed
						TestUint32: 691,   // left (test_repeated.*)
						TestUint64: 44,    // left in 1 removed in 0
						TestSint32: 80087, // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						// TestFixed32 // removed (tsm.*)
						// TestFixed64 // removed in b
					},
					"1": {
						TestInt32:  126, // left as is
						TestInt64:  782, // changed
						TestUint32: 692, // left (test_repeated.*)
						TestUint64: 45,  // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
						// TestFixed32 // removed (tsm.*)
						TestFixed64: 657, // removed in b
					},
					"d": { // should be added
						TestInt32: 0x_D06_F00D,
					},
				},
			},
			Mask: &mask.Mask{
				Any: &mask.Mask{
					Any: &mask.Mask{
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"test_sint64": mask.New(),
						},
					},
					FieldParts: map[mask.FieldKey]*mask.Mask{
						"1": {
							FieldParts: map[mask.FieldKey]*mask.Mask{
								"test_sint32": mask.New(),
							},
						},
					},
				},
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_repeated": {
						Any: &mask.Mask{
							FieldParts: map[mask.FieldKey]*mask.Mask{
								"test_uint32": mask.New(),
							},
						},
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"0": {
								FieldParts: map[mask.FieldKey]*mask.Mask{
									"test_uint64": mask.New(),
								},
							},
						},
					},
					"test_stringmap": {
						Any: &mask.Mask{
							FieldParts: map[mask.FieldKey]*mask.Mask{
								"test_fixed32": mask.New(),
							},
						},
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"b": {
								FieldParts: map[mask.FieldKey]*mask.Mask{
									"test_fixed64": mask.New(),
								},
							},
						},
					},
				},
			},
		},
		{
			Data: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestInt32:  69,
						TestInt64:  456,
						TestUint32: 123,
						TestUint64: 43,
						TestSint32: 80085,
						TestSint64: 0x_BAD_F00D,
					},
					{
						TestInt32:  0x_D0660,
						TestInt64:  457,
						TestUint32: 690,
						TestUint64: 42,
						TestSint32: 80086,
						TestSint64: 0x_BAD_F00D,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{},
			},
			Patch: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						// TestInt32: , // left as is
						TestInt64: 7357, // changed
						// TestUint32 // removed (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
					},
					{
						// TestInt32: , // left as is
						TestInt64: 0x7E57, // changed
						// TestUint32 // removed (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
					},
					{
						TestSint64: 0x_600D_F00D,
					},
				},
				TestStringmap: map[string]*testdata.TestSimple{},
			},
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestInt32: 69,   // left as is
						TestInt64: 7357, // changed
						// TestUint32 // removed (test_repeated.*)
						// TestUint64 // left in 1 removed in 0
						TestSint32: 80085, // left in 0 removed in 1
						// TestSint64 // removed (*.*)
					},
					{
						TestInt32: 0x_D0660, // left as is
						TestInt64: 0x7E57,   // changed
						// TestUint32 // removed (test_repeated.*)
						TestUint64: 42, // left in 1 removed in 0
						// TestSint32 // left in 0 removed in 1
						// TestSint64 // removed (*.*)
					},
					{
						TestSint64: 0x_600D_F00D,
					},
				},
			},
			Mask: &mask.Mask{
				Any: &mask.Mask{
					Any: &mask.Mask{
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"test_sint64": mask.New(),
						},
					},
					FieldParts: map[mask.FieldKey]*mask.Mask{
						"1": {
							FieldParts: map[mask.FieldKey]*mask.Mask{
								"test_sint32": mask.New(),
							},
						},
					},
				},
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_repeated": {
						Any: &mask.Mask{
							FieldParts: map[mask.FieldKey]*mask.Mask{
								"test_uint32": mask.New(),
							},
						},
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"0": {
								FieldParts: map[mask.FieldKey]*mask.Mask{
									"test_uint64": mask.New(),
								},
							},
						},
					},
				},
			},
		},
		{
			Data:  proto.Clone(veryRecursive),
			Patch: proto.Clone(veryRecursive2),
			Err:   strings.Repeat("recursive: ", recursionTooDeep) + "recursion too deep",
		},
		{
			Data: &testdata.TestSimple{
				TestInt32: 42,
				TestInt64: 0x_BAD_F00D,
			},
			Patch: &testdata.TestOptional{
				TestSingle: proto.Uint32(42),
			},
			Result: &testdata.TestSimple{
				TestInt32: 42,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_int64": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Patch: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Err: "test_double: type mismatch: data is double, patch is message",
		},
		{
			Data: &testdata.TestRepeated{
				TestInt32: []int32{1, 2, 3},
			},
			Patch: &testdata.TestSimple{
				TestInt32: 42,
			},
			Err: "test_int32: data field is repeated while patch field is not",
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			Patch: &testdata.NotAStringMap{
				TestInt32: &testdata.TestOptional{},
			},
			Err: "test_int32: data field is map while patch field is not",
		},
		{
			Data: &testdata.TestStructs{},
			Patch: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{TestInt32: 42},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"a": {TestInt64: 69},
				},
				TestStruct: &testdata.TestSimple{TestUint32: 0x_F00D},
			},
			Result: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{TestInt32: 42},
				},
				TestStringmap: map[string]*testdata.TestSimple{
					"a": {TestInt64: 69},
				},
				TestStruct: &testdata.TestSimple{TestUint32: 0x_F00D},
			},
		},
		{
			Data: &testdata.TestOptional{
				TestSingle: proto.Uint32(35),
			},
			Patch: &testdata.TestOptional{
				TestSingle: proto.Uint32(42),
			},
			Result: &testdata.TestOptional{
				TestSingle: proto.Uint32(42),
			},
		},
		{
			Data: &testdata.TestOptional{
				TestSingle: proto.Uint32(35),
			},
			Patch: &testdata.TestOptional{
				TestSingle: proto.Uint32(0),
			},
			Result: &testdata.TestOptional{
				TestSingle: proto.Uint32(0),
			},
		},
		{
			Data: &testdata.TestOptional{
				TestSingle: proto.Uint32(35),
			},
			Patch: &testdata.TestOptional{
				TestSingle: proto.Uint32(0),
			},
			Result: &testdata.TestOptional{
				TestSingle: proto.Uint32(0),
			},
		},
		{
			Data: &testdata.TestOptional{
				TestSingle: proto.Uint32(35),
			},
			Patch: &testdata.TestOptional{
				TestSingle: nil,
			},
			Result: &testdata.TestOptional{
				TestSingle: proto.Uint32(35),
			},
		},
		{
			Data: &testdata.TestOptional{
				TestSingle: proto.Uint32(35),
			},
			Patch: &testdata.TestOptional{
				TestSingle: nil,
			},
			Result: &testdata.TestOptional{
				TestSingle: nil,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_single": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestOptional{
				TestStruct: nil,
			},
			Patch: &testdata.TestOptional{
				TestStruct: &testdata.TestSimple{
					TestInt32: 42,
				},
			},
			Result: &testdata.TestOptional{
				TestStruct: &testdata.TestSimple{
					TestInt32: 42,
				},
			},
		},
		{
			Data: &testdata.TestOptional{
				TestStruct: &testdata.TestSimple{
					TestInt32: 42,
				},
			},
			Patch: &testdata.TestOptional{
				TestStruct: nil,
			},
			Result: &testdata.TestOptional{
				TestStruct: nil,
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_struct": mask.New(),
				},
			},
		},
		{
			Data: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Patch: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive2).(*testdata.RecursiveStruct),
				},
			},
			Err: "field[0]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Data: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"a": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Patch: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"a": proto.Clone(veryRecursive2).(*testdata.RecursiveStruct),
				},
			},
			Err: "field[a]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Data: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{},
			},
			Patch: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive2).(*testdata.RecursiveStruct),
				},
			},
			Err: "field[0]: " + strings.Repeat("recursive: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Data: &testdata.TestStringMap{
				TestInt32: map[string]int32{
					"a": 1,
				},
			},
			Patch: &testdata.TestTypeMap{
				TestInt32: map[int32]string{
					1: "a",
				},
			},
			Err: "test_int32: map key type mismatch: data has string, patch has int32",
		},
		{
			Data: &testdata.TestTypeMap{
				TestInt32: map[int32]string{
					1: "a",
				},
			},
			Patch: &testdata.WrongTypeMap{
				TestInt32: map[int32]int64{
					1: 42,
				},
			},
			Err: "test_int32: map value type mismatch: data has string, patch has int64",
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Patch: &testdata.TestOneOf{},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Patch: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 0,
				},
			},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 0,
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Patch:  &testdata.TestOneOf{},
			Result: &testdata.TestOneOf{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_string": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Patch:  &testdata.TestOneOf{},
			Result: &testdata.TestOneOf{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_oneof": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Patch: &testdata.TestOneOf{},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_uint32": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Patch: &testdata.TestOneOf{},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_uint32": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Patch: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{},
				},
			},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_oneof": {
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"test_int32": {},
						},
					},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Patch: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{},
				},
			},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{},
				},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_struct": {
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"test_int32": {},
						},
					},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Patch:  &testdata.TestOneOf{},
			Result: &testdata.TestOneOf{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_oneof": {
						FieldParts: map[mask.FieldKey]*mask.Mask{
							"test_int32": {},
						},
					},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Patch:  &testdata.TestAnotherOneOf{},
			Result: &testdata.TestOneOf{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_oneof": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Patch: &testdata.TestAnotherOneOf{},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_another_oneof": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 42,
				},
			},
			Patch:  &testdata.TestAnotherOneOf{},
			Result: &testdata.TestOneOf{},
			Mask: &mask.Mask{
				FieldParts: map[mask.FieldKey]*mask.Mask{
					"test_another_oneof": {},
				},
			},
		},
		{
			Data: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 38,
				},
			},
			Patch: &testdata.TestAnotherOneOf{
				TestAnotherOneof: &testdata.TestAnotherOneOf_TestUint32{
					TestUint32: 42,
				},
			},
			Result: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 42,
				},
			},
		},
		{
			Data: &testdata.TestOutputOnly{
				I: 42,
				S: "foo",
				Msg: &testdata.TestSimple{
					TestInt64: 0x_F00D,
				},
				M: map[string]int32{"bar": 7357},
				R: []int32{1, -2, 3},
			},
			Mask:  mask.ParseMust("*.*.*"),
			Patch: &testdata.TestOutputOnly{},
			Result: &testdata.TestOutputOnly{
				I: 42,
				S: "foo",
				Msg: &testdata.TestSimple{
					TestInt64: 0x_F00D,
				},
				M: map[string]int32{"bar": 7357},
				R: []int32{1, -2, 3},
			},
		},
		{
			Data: &testdata.TestImmutable{
				I: 42,
				S: "foo",
				Msg: &testdata.TestSimple{
					TestInt64: 0x_F00D,
				},
				M: map[string]int32{"bar": 7357},
				R: []int32{1, -2, 3},
			},
			Mask:   mask.ParseMust("*.*.*"),
			Patch:  &testdata.TestImmutable{},
			Result: &testdata.TestImmutable{},
		},
		{
			Data: &testdata.TestOutputOnly{
				I: 42,
				S: "foo",
				Msg: &testdata.TestSimple{
					TestInt64: 0x_F00D,
				},
				M: map[string]int32{"bar": 7357},
				R: []int32{1, -2, 3},
			},
			Mask: mask.ParseMust("*.*.*"),
			Patch: &testdata.TestOutputOnly{
				I: 37,
				S: "nope",
				M: map[string]int32{
					"baz": 0x_BAD,
				},
				R: []int32{-1},
				Msg: &testdata.TestSimple{
					TestInt32: 0x_BAD_F00D,
				},
			},
			Result: &testdata.TestOutputOnly{
				I: 42,
				S: "foo",
				Msg: &testdata.TestSimple{
					TestInt64: 0x_F00D,
				},
				M: map[string]int32{"bar": 7357},
				R: []int32{1, -2, 3},
			},
		},
		{
			Data: &testdata.TestImmutable{
				I: 42,
				S: "foo",
				Msg: &testdata.TestSimple{
					TestInt64: 0x_F00D,
				},
				M: map[string]int32{"bar": 7357},
				R: []int32{1, -2, 3},
			},
			Mask: mask.ParseMust("*.*.*"),
			Patch: &testdata.TestImmutable{
				I: 37,
				S: "nope",
				M: map[string]int32{
					"baz": 0x_BAD,
				},
				R: []int32{-1},
				Msg: &testdata.TestSimple{
					TestInt32: 0x_BAD_F00D,
				},
			},
			Result: &testdata.TestImmutable{
				I: 37,
				S: "nope",
				M: map[string]int32{
					"baz": 0x_BAD,
				},
				R: []int32{-1},
				Msg: &testdata.TestSimple{
					TestInt32: 0x_BAD_F00D,
				},
			},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			err := PatchWithResetMask(c.Data, c.Patch, c.Mask)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err, protojson.Format(c.Data))
			} else {
				assert.NoError(t, err)
				assert.EqualExportedValues(t, c.Result, c.Data)
			}
		})
	}
}
