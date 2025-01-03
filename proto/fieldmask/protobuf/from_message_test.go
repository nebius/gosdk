package protobuf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"

	"github.com/nebius/gosdk/proto/fieldmask/protobuf/testdata"
)

func TestResetMaskFromMessage(t *testing.T) {
	t.Parallel()
	someuint32full := uint32(42)
	someuint32null := uint32(0)
	veryRecursive := &testdata.RecursiveStruct{
		SomeString: "foo",
	}
	for range recursionTooDeep + 42 {
		veryRecursive = &testdata.RecursiveStruct{
			Recursive: veryRecursive,
		}
	}
	cases := []struct {
		Input proto.Message
		Mask  string
		Err   string
	}{
		{
			Input: &testdata.TestOneOf{},
			Mask:  "test_string,test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: nil,
			},
			Mask: "test_string,test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{},
			},
			Mask: "test_string,test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 0,
				},
			},
			Mask: "test_string,test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestUint32{
					TestUint32: 42,
				},
			},
			Mask: "test_string,test_struct",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{},
			},
			Mask: "test_string,test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "",
				},
			},
			Mask: "test_string,test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestString{
					TestString: "abc",
				},
			},
			Mask: "test_struct,test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{},
			},
			Mask: "test_string,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: nil,
				},
			},
			Mask: "test_string,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{},
				},
			},
			Mask: "test_string,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_uint32",
		},
		{
			Input: &testdata.TestOneOf{
				TestOneof: &testdata.TestOneOf_TestStruct{
					TestStruct: &testdata.TestSimple{
						TestInt32: 42,
					},
				},
			},
			Mask: "test_string,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_uint32",
		},
		{
			Input: &testdata.TestSimple{},
			Mask: "test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32," +
				"test_sint64,test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestSimple{
				TestInt32: 42,
			},
			Mask: "test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int64,test_sfixed32,test_sfixed64,test_sint32," +
				"test_sint64,test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestSimple{
				TestDouble:      8008.5,
				TestInt32:       42,
				TestFloat:       4.2,
				TestInt64:       80085,
				TestUint32:      0xF00D,
				TestUint64:      0x7457E_F00D,
				TestSint32:      0xD06_F00D,
				TestSint64:      0xE47_F00D,
				TestSfixed32:    7357,
				TestFixed64:     69,
				TestFixed32:     0xDEADBEEF,
				TestSfixed64:    -0xDEADBEEF,
				TestBool:        true,
				TestString:      "foo",
				TestBytes:       []byte("bar"),
				TestEnum:        testdata.ABC_ABC_B,
				TestAliasedEnum: testdata.AliasedABC_AABC_B2,
			},
			Mask: "",
		},
		{
			Input: &testdata.TestStructs{},
			Mask:  "test_intmap,test_repeated,test_stringmap,test_struct",
		},
		{
			Input: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{},
			},
			Mask: "test_intmap,test_repeated,test_stringmap,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				")",
		},
		{
			Input: &testdata.TestStructs{
				TestStruct: &testdata.TestSimple{
					TestDouble: 4.2,
				},
			},
			Mask: "test_intmap,test_repeated,test_stringmap,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				")",
		},
		{
			Input: &testdata.TestStructs{
				TestRepeated:  []*testdata.TestSimple{},
				TestStringmap: map[string]*testdata.TestSimple{},
			},
			Mask: "test_intmap,test_repeated,test_stringmap,test_struct",
		},
		{
			Input: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestDouble: 4.2,
					},
					{},
				},
			},
			Mask: "test_intmap,test_repeated.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_stringmap,test_struct",
		},
		{
			Input: &testdata.TestStructs{
				TestRepeated: []*testdata.TestSimple{
					{
						TestDouble: 4.2,
					},
					{
						TestDouble: 4.2,
					},
				},
			},
			Mask: "test_intmap,test_repeated.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_stringmap,test_struct",
		},
		{
			Input: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": {
						TestDouble: 4.2,
					},
					"b": {},
				},
			},
			Mask: "test_intmap,test_repeated,test_stringmap.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum," +
				"test_fixed32,test_fixed64,test_float,test_int32,test_int64,test_sfixed32," +
				"test_sfixed64,test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_struct",
		},
		{
			Input: &testdata.TestStructs{
				TestStringmap: map[string]*testdata.TestSimple{
					"a": {
						TestDouble: 4.2,
					},
					"b": {
						TestDouble: 4.2,
					},
				},
			},
			Mask: "test_intmap,test_repeated,test_stringmap.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_enum," +
				"test_fixed32,test_fixed64,test_float,test_int32,test_int64,test_sfixed32," +
				"test_sfixed64,test_sint32,test_sint64,test_string,test_uint32,test_uint64" +
				"),test_struct",
		},
		{
			Input: &testdata.TestRepeated{},
			Mask: "test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestRepeated{
				TestDouble: []float64{},
			},
			Mask: "test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestRepeated{
				TestDouble: []float64{
					4.2,
				},
			},
			Mask: "test_aliased_enum,test_bool,test_bytes,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestTypeMap{},
			Mask: "test_bool,test_fixed32,test_fixed64,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestTypeMap{
				TestUint32: map[uint32]string{},
			},
			Mask: "test_bool,test_fixed32,test_fixed64,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint32,test_uint64",
		},
		{
			Input: &testdata.TestTypeMap{
				TestUint32: map[uint32]string{
					42: "this is the answer",
				},
			},
			Mask: "test_bool,test_fixed32,test_fixed64,test_int32,test_int64,test_sfixed32,test_sfixed64," +
				"test_sint32,test_sint64,test_string,test_uint64",
		},
		{
			Input: &testdata.TestOptional{},
			Mask:  "test_single,test_struct",
		},
		{
			Input: &testdata.TestOptional{
				TestSingle: &someuint32null,
			},
			Mask: "test_single,test_struct",
		},
		{
			Input: &testdata.TestOptional{
				TestSingle: &someuint32full,
			},
			Mask: "test_struct",
		},
		{
			Input: &testdata.TestOptional{
				TestStruct: &testdata.TestSimple{},
			},
			Mask: "test_single,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64" +
				")",
		},
		{
			Input: &testdata.TestOptional{
				TestStruct: &testdata.TestSimple{
					TestDouble: 4.2,
				},
			},
			Mask: "test_single,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64" +
				")",
		},
		{
			Input: &testdata.TestRecursiveSingle{},
			Mask:  "field",
		},
		{
			Input: &testdata.TestRecursiveSingle{
				Field: &testdata.RecursiveStruct{},
			},
			Mask: "field.(recursive,some_string)",
		},
		{
			Input: &testdata.TestRecursiveSingle{
				Field: &testdata.RecursiveStruct{
					SomeString: "foo",
					Recursive: &testdata.RecursiveStruct{
						Recursive: &testdata.RecursiveStruct{
							SomeString: "bar",
							Recursive:  &testdata.RecursiveStruct{},
						},
					},
				},
			},
			Mask: "field.recursive.(recursive.recursive.(recursive,some_string),some_string)",
		},
		{
			Input: &testdata.TestRecursiveSingle{
				Field: proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
			},
			Err: "failed to collect modification mask: TestRecursiveSingle: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Input: &testdata.TestRecursiveRepeated{
				Field: []*testdata.RecursiveStruct{
					proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Err: "failed to collect modification mask: TestRecursiveRepeated[0]: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Input: &testdata.TestRecursiveMap{
				Field: map[string]*testdata.RecursiveStruct{
					"foo": proto.Clone(veryRecursive).(*testdata.RecursiveStruct),
				},
			},
			Err: "failed to collect modification mask: TestRecursiveMap[foo]: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Input: &testdata.TestX{
				Y: []*testdata.TestY{
					{
						Z: []*testdata.TestZ{
							{
								A: "a",
								B: "",
							},
						},
					},
					{
						Z: nil,
					},
				},
			},
			Mask: "y.*.z.*.b",
		},
		{
			Input: &testdata.TestA{
				B: map[string]*testdata.TestB{
					"1": {
						C: map[string]*testdata.TestC{
							"1": {
								X: "x",
								Y: "",
							},
						},
					},
					"2": {
						C: nil,
					},
				},
			},
			Mask: "b.*.c.*.y",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			ret, err := ResetMaskFromMessage(c.Input)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				str, err := ret.Marshal()
				assert.NoError(t, err)
				assert.Equal(t, c.Mask, str)
			}
		})
	}
	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		ret, err := ResetMaskFromMessage(nil)
		assert.NoError(t, err)
		assert.Nil(t, ret)
	})
}
