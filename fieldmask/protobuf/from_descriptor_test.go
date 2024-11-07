package protobuf

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"

	"github.com/nebius/gosdk/fieldmask/protobuf/testdata"
)

func TestKnownFieldsFromMessage(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Input proto.Message
		Err   string
		Mask  string
	}{
		{
			Input: (&testdata.TestSimple{}),
			Mask: "test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32," +
				"test_fixed64,test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32," +
				"test_sint64,test_string,test_uint32,test_uint64",
		},
		{
			Input: (&testdata.TestStructs{}),
			Mask: "test_intmap.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64" +
				"),test_repeated.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64" +
				"),test_stringmap.*.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64" +
				"),test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64" +
				")",
		},
		{
			Input: (&testdata.TestOptional{}),
			Mask: "test_single,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64)",
		},
		{
			Input: (&testdata.TestOneOf{}),
			Mask: "test_string,test_struct.(" +
				"test_aliased_enum,test_bool,test_bytes,test_double,test_enum,test_fixed32,test_fixed64," +
				"test_float,test_int32,test_int64,test_sfixed32,test_sfixed64,test_sint32,test_sint64," +
				"test_string,test_uint32,test_uint64),test_uint32",
		},
		{
			Input: (&testdata.TestWellKnown{}),
			Mask: "test_any.(type_url,value),test_repeated_any.*.(type_url,value),test_repeated_ts.*.(nanos,seconds)," +
				"test_ts.(nanos,seconds)",
		},
		{ // for now next three tests have errors, while I hope at some point they will change to noerror
			Input: (&testdata.TestRecursiveSingle{}),
			Err: "failed to collect modification mask: TestRecursiveSingle: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Input: (&testdata.TestRecursiveRepeated{}),
			Err: "failed to collect modification mask: TestRecursiveRepeated: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Input: (&testdata.TestRecursiveMap{}),
			Err: "failed to collect modification mask: TestRecursiveMap: " +
				strings.Repeat("RecursiveStruct: ", recursionTooDeep-1) + "recursion too deep",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d from message", i), func(t *testing.T) {
			t.Parallel()
			ret, err := KnownFieldsFromMessage(c.Input)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				str, err := ret.Marshal()
				assert.NoError(t, err)
				assert.Equal(t, c.Mask, str)
			}
		})
		t.Run(fmt.Sprintf("case %d from descriptor", i), func(t *testing.T) {
			t.Parallel()
			ret, err := KnownFieldsFromDescriptor(c.Input.ProtoReflect().Descriptor())
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
}
