package protobuf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"

	"github.com/nebius/gosdk/proto/fieldmask/protobuf/testdata"
)

func TestResetMaskFromModified(t *testing.T) {
	t.Parallel()
	zero1 := uint32(0)
	zero2 := uint32(0)
	life1 := uint32(42)
	life2 := uint32(42)
	other1 := uint32(69)

	veryRecursive := &testdata.RecursiveStruct{
		SomeString: "foo",
	}
	for _ = range recursionTooDeep + 42 {
		veryRecursive = &testdata.RecursiveStruct{
			Recursive: veryRecursive,
		}
	}

	cases := []struct {
		Initial  proto.Message
		Modified proto.Message
		Mask     string
		Err      string
	}{
		{
			Initial:  nil,
			Modified: &testdata.TestSimple{},
			Err:      "received nil",
		},
		{
			Initial:  nil,
			Modified: nil,
			Err:      "received nil",
		},
		{
			Initial:  &testdata.TestSimple{},
			Modified: nil,
			Err:      "received nil",
		},
		{
			Initial: &testdata.TestSimple{
				TestDouble: 4.2,
			},
			Modified: &testdata.TestSimple{},
			Mask:     "test_double",
		},
		{
			Initial: &testdata.TestSimple{},
			Modified: &testdata.TestSimple{
				TestDouble: 4.2,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestSimple{
				TestDouble: 5.3,
			},
			Modified: &testdata.TestSimple{
				TestDouble: 4.2,
			},
			Mask: "",
		},
		{
			Initial:  &testdata.TestStructs{},
			Modified: &testdata.TestSimple{},
			Err:      "recieved messages of different types: initial nebius.testdata.TestStructs, modified nebius.testdata.TestSimple",
		},
		{
			Initial: &testdata.TestOptional{
				TestStruct: &testdata.TestSimple{},
			},
			Modified: &testdata.TestOptional{},
			Mask:     "test_struct",
		},
		{
			Initial: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "test_struct",
		},
		{
			Initial: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{},
			},
			Modified: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{
					TestDouble: 4.2,
				},
			},
			Modified: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{},
			},
			Mask: "test_struct.test_double",
		},
		{
			Initial: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{
					TestDouble: 4.2,
				},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "test_struct",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "",
		},
		{
			Initial: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "test_repeated",
		},
		{
			Initial: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"foo": &testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "test_stringmap",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{
						TestDouble: 4.2,
					},
					&testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "test_repeated",
		},
		{
			Initial: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": &testdata.TestSimple{
						TestDouble: 4.2,
					},
					"b": &testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{},
			Mask:     "test_stringmap",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{
						TestDouble: 4.2,
					},
					&testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{},
			},
			Mask: "test_repeated",
		},
		{
			Initial: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": &testdata.TestSimple{
						TestDouble: 4.2,
					},
					"b": &testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{},
			},
			Mask: "test_stringmap",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{
						TestDouble: 4.2,
					},
					&testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{},
				},
			},
			Mask: "test_repeated.0.test_double",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{
						TestDouble: 4.2,
					},
				},
			},
			Modified: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{},
					&testdata.TestSimple{},
				},
			},
			Mask: "test_repeated.0.test_double",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{},
			},
			Modified: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{},
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{
						TestDouble: 4.2,
					},
					&testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					&testdata.TestSimple{
						TestDouble: 5.3,
					},
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": &testdata.TestSimple{},
					"b": &testdata.TestSimple{
						TestDouble: 4.2,
					},
					"c": &testdata.TestSimple{
						TestDouble: 4.2,
					},
					"d": &testdata.TestSimple{
						TestDouble: 4.2,
					},
				},
			},
			Modified: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": &testdata.TestSimple{},
					"b": &testdata.TestSimple{},
					"c": &testdata.TestSimple{
						TestDouble: 4.3,
					},
					"e": &testdata.TestSimple{
						TestDouble: 4.2,
					},
				},
			},
			Mask: "test_stringmap.b.test_double",
		},
		{
			Initial: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": &testdata.TestSimple{},
					"c": &testdata.TestSimple{
						TestDouble: 4.2,
					},
				},
			},
			Modified: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": &testdata.TestSimple{},
					"b": &testdata.TestSimple{},
					"c": &testdata.TestSimple{
						TestDouble: 4.3,
					},
					"e": &testdata.TestSimple{
						TestDouble: 4.2,
					},
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: nil,
			},
			Modified: &testdata.TestOptional{
				TestSingle: nil,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &zero1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: nil,
			},
			Mask: "test_single",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &zero1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &zero2,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: nil,
			},
			Mask: "test_single",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &life2,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &other1,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &zero1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &zero1,
			},
			Mask: "test_single",
		},
		{
			Initial: &testdata.TestOptional{
				TestSingle: nil,
			},
			Modified: &testdata.TestOptional{
				TestSingle: &life1,
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestRecursiveSingle{
				Field: &testdata.RecursiveStruct{
					SomeString: "foo",
					Recursive: &testdata.RecursiveStruct{
						Recursive: &testdata.RecursiveStruct{
							SomeString: "bar",
							Recursive: &testdata.RecursiveStruct{
								SomeString: "baz",
							},
						},
					},
				},
			},
			Modified: &testdata.TestRecursiveSingle{
				Field: &testdata.RecursiveStruct{
					Recursive: &testdata.RecursiveStruct{
						SomeString: "bax",
						Recursive: &testdata.RecursiveStruct{
							SomeString: "bae",
							Recursive:  &testdata.RecursiveStruct{},
						},
					},
				},
			},
			Mask: "field.(recursive.recursive.recursive.some_string,some_string)",
		},
		{
			Initial: &testdata.TestRecursiveSingle{
				Field: &testdata.RecursiveStruct{
					SomeString: "foo",
					Recursive: &testdata.RecursiveStruct{
						Recursive: &testdata.RecursiveStruct{
							SomeString: "bar",
							Recursive: &testdata.RecursiveStruct{
								SomeString: "baz",
							},
						},
					},
				},
			},
			Modified: &testdata.TestRecursiveSingle{
				Field: &testdata.RecursiveStruct{
					Recursive: &testdata.RecursiveStruct{
						SomeString: "bax",
						Recursive: &testdata.RecursiveStruct{
							SomeString: "bae",
						},
					},
				},
			},
			Mask: "field.(recursive.recursive.recursive,some_string)",
		},
		{
			Initial: &testdata.TestRecursiveSingle{
				Field: proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
			},
			Modified: &testdata.TestRecursiveSingle{
				Field: proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
			},
			Err: "failed to collect modification mask: TestRecursiveSingle: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Initial: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Modified: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Err: "failed to collect modification mask: TestRecursiveRepeated: field[0]: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Initial: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Modified: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Err: "failed to collect modification mask: TestRecursiveMap: field[foo]: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Initial: &testdata.TestRepeated{
				TestDouble: []float64{},
			},
			Modified: &testdata.TestRepeated{},
			Mask:     "",
		},
		{
			Initial: &testdata.TestRepeated{
				TestDouble: []float64{},
			},
			Modified: &testdata.TestRepeated{
				TestDouble: []float64{},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestRepeated{
				TestDouble: []float64{1, 2},
			},
			Modified: &testdata.TestRepeated{},
			Mask:     "test_double",
		},
		{
			Initial: &testdata.TestRepeated{
				TestDouble: []float64{1, 2},
			},
			Modified: &testdata.TestRepeated{
				TestDouble: []float64{},
			},
			Mask: "test_double",
		},
		{
			Initial: &testdata.TestRepeated{
				TestDouble: []float64{1, 2},
			},
			Modified: &testdata.TestRepeated{
				TestDouble: []float64{4},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{},
			},
			Modified: &testdata.TestStringMap{},
			Mask:     "",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{},
			},
			Modified: &testdata.TestStringMap{
				TestDouble: map[string]float64{},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
				},
			},
			Modified: &testdata.TestStringMap{
				TestDouble: map[string]float64{},
			},
			Mask: "test_double",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
				},
			},
			Modified: &testdata.TestStringMap{},
			Mask:     "test_double",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
				},
			},
			Modified: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
				},
			},
			Modified: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 5.2,
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
				},
			},
			Modified: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"b": 4.2,
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"a": 4.2,
					"b": 4.3,
				},
			},
			Modified: &testdata.TestStringMap{
				TestDouble: map[string]float64{
					"b": 4.2,
				},
			},
			Mask: "",
		},
		{
			Initial:  &testdata.TestOneOf{},
			Modified: &testdata.TestOneOf{},
			Mask:     "",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Modified: &testdata.TestOneOf{},
			Mask:     "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Modified: &testdata.TestOneOf{},
			Mask:     "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Modified: &testdata.TestOneOf{},
			Mask:     "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Modified: &testdata.TestOneOf{},
			Mask:     "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Mask: "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Mask: "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Modified: &testdata.TestOneOf{},
			Mask:     "test_string",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: nil,
				},
			},
			Mask: "",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestUint32: 42,
					},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: nil,
				},
			},
			Mask: "test_struct.test_uint32",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestUint32: 42,
					},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{},
				},
			},
			Mask: "test_struct.test_uint32",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestUint32: 23,
						TestInt64:  35,
					},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32:  69,
						TestUint32: 42,
					},
				},
			},
			Mask: "test_struct.test_int64",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestUint32: 23,
						TestInt64:  35,
					},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Mask: "test_struct",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestUint32: 23,
						TestInt64:  35,
					},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Mask: "test_struct",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestUint32: 23,
						TestInt64:  35,
					},
				},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Mask: "test_struct",
		},
		{
			Initial: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{},
			},
			Modified: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Mask: "test_struct",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			ret, err := ResetMaskFromModified(c.Initial, c.Modified)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				if ret != nil {
					str, err := ret.Marshal()
					assert.NoError(t, err)
					assert.Equal(t, c.Mask, str)
				}
			}
		})
	}
}
