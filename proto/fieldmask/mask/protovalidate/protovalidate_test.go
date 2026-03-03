package protovalidate

import (
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldPathFromProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   *validate.FieldPath
		want    []string
		wantErr string
	}{
		{
			name:  "nil path",
			input: nil,
			want:  nil,
		},
		{
			name:  "empty path",
			input: &validate.FieldPath{},
			want:  []string{},
		},
		{
			name: "simple named path",
			input: &validate.FieldPath{Elements: []*validate.FieldPathElement{
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("metadata")
					return e
				}(),
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("labels")
					return e
				}(),
			}},
			want: []string{"metadata", "labels"},
		},
		{
			name: "all subscript kinds",
			input: &validate.FieldPath{Elements: []*validate.FieldPathElement{
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("items")
					e.SetIndex(3)
					return e
				}(),
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("bool_map")
					e.SetBoolKey(true)
					return e
				}(),
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("int_map")
					e.SetIntKey(-7)
					return e
				}(),
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("uint_map")
					e.SetUintKey(9)
					return e
				}(),
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("string_map")
					e.SetStringKey("hello")
					return e
				}(),
			}},
			want: []string{"items", "3", "bool_map", "true", "int_map", "-7", "uint_map", "9", "string_map", "hello"},
		},
		{
			name: "fallback to field number",
			input: &validate.FieldPath{Elements: []*validate.FieldPathElement{
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldNumber(42)
					return e
				}(),
			}},
			want: []string{"42"},
		},
		{
			name: "prefer field name over number",
			input: &validate.FieldPath{Elements: []*validate.FieldPathElement{
				func() *validate.FieldPathElement {
					e := &validate.FieldPathElement{}
					e.SetFieldName("name")
					e.SetFieldNumber(17)
					return e
				}(),
			}},
			want: []string{"name"},
		},
		{
			name: "nil element error",
			input: &validate.FieldPath{Elements: []*validate.FieldPathElement{
				nil,
			}},
			wantErr: "field path element 0 is nil",
		},
		{
			name: "missing name and number error",
			input: &validate.FieldPath{Elements: []*validate.FieldPathElement{
				{},
			}},
			wantErr: "field path element 0 has no field name or field number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := FieldPathFromProto(tt.input)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Nil(t, got)
				return
			}

			require.NoError(t, err)
			if tt.want == nil {
				assert.Nil(t, got)
				return
			}

			require.NotNil(t, got)
			gotStrings := make([]string, len(*got))
			for i, key := range *got {
				gotStrings[i] = string(key)
			}
			assert.Equal(t, tt.want, gotStrings)
		})
	}
}
