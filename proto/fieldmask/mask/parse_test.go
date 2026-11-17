package mask

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Input  string
		Output string
		Err    string
	}{
		{
			Input: "*,",
			Err:   "unexpected end of mask",
		},
		{
			Input: "*.",
			Err:   "unexpected end of mask",
		},
		{
			Input: "*.(",
			Err:   "unclosed left brace at position 2 near \"*.\u20DE(\"",
		},
		{
			Input: " \t\r\n*.( \t\r\n",
			Err:   "unclosed left brace at position 6 near \" \\t\\r\\n*.\u20DE( \\t\\r\\n\"",
		},
		{
			Input: "abcdefghijklmonpqrst.(abcdefghijklmonpqrst",
			Err:   "unclosed left brace at position 21 near \"jklmonpqrst.\u20DE(abcdefghijklmo...\"",
		},
		{
			Input: "abcdefghijklmonpqrst)abcdefghijklmonpqrst",
			Err:   "unmatched right brace at position 20 near \"ijklmonpqrst\u20DE)abcdefghijklmo...\"",
		},
		{
			Input: "abcdefghijklmonpqrst..abcdefghijklmonpqrst",
			Err:   "unexpected token TokenDot(\".\" pos 21), expecting field or submask at position 21 near \"jklmonpqrst.\u20DE.abcdefghijklmo...\"",
		},
		{
			Input: "abcdefghijklmonpqrst abcdefghijklmonpqrst feafwafwawfadwadw",
			Err:   "unexpected token TokenPlainKey(\"abcdefghijklmonpqrst\" pos 21), expecting separator or closing brace at position 21 near \"jklmonpqrst \u20DEabcdefghijklmon...\"",
		},
		{
			Input: "#",
			Err:   "tokenizer failure: failed to tokenize: unexpected symbol at position 0 near \"\u20DE#\"",
		},
		{
			Input: "#1234567890abcdefghijklmnopqrst",
			Err:   "tokenizer failure: failed to tokenize: unexpected symbol at position 0 near \"\u20DE#1234567890abcdefghijklmnop...\"",
		},
		{
			Input: "\"1234567890abcdefghijklmnopqrst",
			Err:   "tokenizer failure: failed to tokenize: unterminated quoted string at position 0 near \"\u20DE\\\"1234567890abcdefghijklmnop...\"",
		},
		{
			Input:  "",
			Output: "",
		},
		{
			Input:  "()",
			Output: "",
		},
		{
			Input:  " \t\r\n( \t\r\n) \t\r\n",
			Output: "",
		},
		{
			Input:  " \t\r\n",
			Output: "",
		},
		{
			Input:  "*",
			Output: "*",
		},
		{
			Input:  " \t\r\n* \t\r\n",
			Output: "*",
		},
		{
			Input:  " \t\r\na \t\r\n",
			Output: "a",
		},
		{
			Input:  " \t\r\n( \t\r\na \t\r\n) \t\r\n",
			Output: "a",
		},
		{
			Input:  "*,*,*,*",
			Output: "*",
		},
		{
			Input:  "test",
			Output: "test",
		},
		{
			Input:  `"test"`,
			Output: `test`,
		},
		{
			Input:  `"test "`,
			Output: `"test "`,
		},
		{
			Input:  "a,a",
			Output: "a",
		},
		{
			Input:  "a.b",
			Output: "a.b",
		},
		{
			Input:  "a \t\r\n. \t\r\nb",
			Output: "a.b",
		},
		{
			Input:  "a \t\r\n, \t\r\na",
			Output: "a",
		},
		{
			Input:  "*,test",
			Output: "*,test",
		},
		{
			Input:  "* \t\r\n, \t\r\ntest",
			Output: "*,test",
		},
		{
			Input:  "test,*",
			Output: "*,test",
		},
		{
			Input:  "a.b,a.b",
			Output: "a.b",
		},
		{
			Input:  "a.*,a.*",
			Output: "a.*",
		},
		{
			Input:  "*.b,*.b",
			Output: "*.b",
		},
		{
			Input:  "a.b,a.c",
			Output: "a.(b,c)",
		},
		{
			Input:  "a.(b)",
			Output: "a.b",
		},
		{
			Input:  "a.(b,())",
			Output: "a.b",
		},
		{
			Input:  "a.((),b)",
			Output: "a.b",
		},
		{
			Input:  "a.((()))",
			Output: "a",
		},
		{
			Input:  "a.(b,c)",
			Output: "a.(b,c)",
		},
		{
			Input:  "*.(b,c)",
			Output: "*.(b,c)",
		},
		{
			Input:  "a.(b,c).(d,e)",
			Output: "a.(b.(d,e),c.(d,e))",
		},
		{
			Input:  "a.(*,c).(d,e)",
			Output: "a.(*.(d,e),c.(d,e))",
		},
		{
			Input:  "a.(((((((*,c))))))).(d,e)",
			Output: "a.(*.(d,e),c.(d,e))",
		},
		{
			Input:  "(((((((a))))))).(((((*)))),(((c)))).(((((((d,(((e))))))))))",
			Output: "a.(*.(d,e),c.(d,e))",
		},
		{
			Input:  "a.(b,c.(d,e))",
			Output: "a.(b,c.(d,e))",
		},
		{
			Input:  "a.(*,b,c)",
			Output: "a.(*,b,c)",
		},
		{
			Input:  "a.(*,b,c.(d,e))",
			Output: "a.(*,b,c.(d,e))",
		},
		{
			Input:  "*.*.(a,b,c.*,d.(e,f),g.(*.(h,i),j,k)),1,A,B,l,m,n.*,o.(p,q,w,x),r.(*.(s,t),u.*,v),z.(*.y,u.*,v)",
			Output: "*.*.(a,b,c.*,d.(e,f),g.(*.(h,i),j,k)),1,A,B,l,m,n.*,o.(p,q,w,x),r.(*.(s,t),u.*,v),z.(*.y,u.*,v)",
		},
		{
			Input:  "a.\"\\\",.() \\t\\r\\n\".b",
			Output: "a.\"\\\",.() \\t\\r\\n\".b",
		},
		// A branch of a group stays independent of the other branches.
		{
			Input:  "(a,b).c,a.c.d",
			Output: "a.c.d,b.c",
		},
		{
			Input:  "a.c.d,(a,b).c",
			Output: "a.c.d,b.c",
		},
		// A key after a group extends the branches of the nested groups too.
		{
			Input:  "(a.(b,c),d).e",
			Output: "a.(b.e,c.e),d.e",
		},
	}
	for i, c := range cases {
		name := fmt.Sprintf("case %d ", i)
		if c.Err != "" {
			name += "fail"
		} else {
			name += "ok"
		}
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			mask, err := Parse(c.Input)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, mask)
				normalized, err := mask.Marshal()
				assert.NoError(t, err)
				assert.Equal(t, c.Output, normalized)
			}
		})
	}
}

func TestParseDepth(t *testing.T) {
	t.Parallel()
	path := func(depth int) string {
		return strings.Repeat("a.", depth-1) + "a"
	}
	cases := map[string]func(int) string{
		"keys": path,
		"wildcards": func(depth int) string {
			return strings.Repeat("*.", depth-1) + "*"
		},
		"comma resets active depth": func(depth int) string {
			return path(depth) + ",b.c"
		},
		"group restores starting depth": func(depth int) string {
			return "a.(" + path(depth-1) + ",b.c)"
		},
		"deep branch first": func(depth int) string {
			return "(" + path(depth-1) + ",b).c"
		},
		"deep branch last": func(depth int) string {
			return "(b," + path(depth-1) + ").c"
		},
		"nested group continuation": func(depth int) string {
			return "(a.(" + path(depth-2) + ",b),c).d"
		},
		"empty branch": func(depth int) string {
			return "(()," + path(depth-1) + ").b"
		},
		"equal branches": func(depth int) string {
			return path(depth-1) + ".(a,a)"
		},
		"braces do not add depth": func(depth int) string {
			return strings.Repeat("(", recursionTooDeep) + path(depth) +
				strings.Repeat(")", recursionTooDeep)
		},
	}
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			parsed, err := Parse(input(recursionTooDeep - 1))
			require.NoError(t, err)
			_, err = parsed.Copy()
			require.NoError(t, err)

			_, err = Parse(input(recursionTooDeep))
			require.ErrorIs(t, err, ErrRecursionTooDeep)
		})
	}
}

func TestParseSteps(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		Input  string
		Steps  int64 // Steps is the exact number of steps that Input needs.
		Output string
	}{
		"sequence of groups": {
			Input:  "(a,b).(c,d)",
			Steps:  12,
			Output: "a.(c,d),b.(c,d)",
		},
		// Each group ends in one mask, thus it needs four steps.
		"equal branches": {
			Input:  "a" + strings.Repeat(".(a,a)", 20),
			Steps:  81,
			Output: strings.Repeat("a.", 20) + "a",
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			mask, err := parse(c.Input, c.Steps)
			require.NoError(t, err)
			normalized, err := mask.Marshal()
			require.NoError(t, err)
			assert.Equal(t, c.Output, normalized)

			_, err = parse(c.Input, c.Steps-1)
			assert.ErrorIs(t, err, ErrTooComplex)
		})
	}
}

func TestParseTooComplex(t *testing.T) {
	t.Parallel()
	// The expanded result would contain 196607 masks.
	_, err := Parse(strings.Repeat("(a,b).", 16) + "c")
	require.ErrorIs(t, err, ErrTooComplex)
	var unmarshalErr *UnmarshalError
	require.ErrorAs(t, err, &unmarshalErr)
	// Parse stops in the eighth group, before it adds 128 masks for "b".
	assert.Equal(t, 45, unmarshalErr.Position)
	assert.EqualError(t, err, "cannot expand mask at position 45 near "+
		"\"b).(a,b).(a,\u20DEb).(a,b).(a,b)....\": mask is too complex")
}

// TestParseMarshaled checks that a mask from [Mask.Marshal] needs one step for
// each token at most. Thus, the step limit has no effect on such masks.
func TestParseMarshaled(t *testing.T) {
	t.Parallel()
	keys := strings.Split("abcdefghij", "")
	wide := New()
	for _, a := range keys {
		for _, b := range keys {
			for _, c := range keys {
				path := NewFieldPath(FieldKey(a), FieldKey(b), FieldKey(c))
				require.NoError(t, wide.Merge(path.ToMask()))
			}
		}
	}
	// Each of the 100 nested groups contains ten fields.
	deep := New()
	for last, i := deep, 0; i < 100; i++ {
		for _, k := range keys {
			last.FieldParts[FieldKey(k)] = New()
		}
		last = last.FieldParts["j"]
	}
	for name, m := range map[string]*Mask{"wide": wide, "deep": deep} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			str, err := m.Marshal()
			require.NoError(t, err)
			// The number of tokens is not more than the length of str.
			parsed, err := parse(str, int64(len(str)))
			require.NoError(t, err)
			assert.True(t, m.Equal(parsed))
		})
	}
}

func TestParseYAML(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input    string
		expected string
		err      string
	}{
		"Simple YAML": {
			input: `
spec: some data
metadata: some data
`,
			expected: "Mask<metadata,spec>",
		},
		"Nested YAML": {
			input: `
spec:
  field1: value1
  field2: value2
  arr_field:
    - subfield: 1
    - another: 2
    - {}
metadata:
  name: object1
  labels:
    env: production
`,
			expected: "Mask<metadata.(labels.env,name),spec.(arr_field.(0.subfield,1.another,2),field1,field2)>",
		},
		"More sophisticated yaml": {
			input: `
spec:
  with space: value1
  has "quotes" inside: value2
  dots,commas.included:
    - subfield: 1
    - another: 2
    - {}
  hyphen-separated: yup
  🐈: everything for our furry friends
metadata:
  (braces[are]cool): object1
`,
			expected: `Mask<metadata."(braces[are]cool)",spec.("dots,commas.included".(0.subfield,1.another,2),"has \"quotes\" inside","hyphen-separated","with space","🐈")>`,
		},
		"not a yaml": {
			input: "!@#$%^&*()",
			err:   "yaml unmarshal: yaml: did not find expected whitespace or line break",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			msk, err := ParseYAML([]byte(tt.input))
			if tt.err != "" {
				require.EqualError(t, err, tt.err)
			} else {
				require.NoError(t, err)

				require.Equal(t, tt.expected, msk.String())
			}
		})
	}
}

func TestParseJSON(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		input    string
		expected string
		err      string
	}{
		"Simple JSON": {
			input: `{
				"spec":"text",
				"metadata":2
			}`,
			expected: "Mask<metadata,spec>",
		},
		"Nested JSON": {
			input: `{
				"spec":{
					"something":1,
					"other_thing":2
				},
				"metadata":{
					"name":"string",
					"id":123
				}
			}`,
			expected: "Mask<metadata.(id,name),spec.(other_thing,something)>",
		},
		"More sophisticated json": {
			input: `{
				"spec": {
					"with space": 123,
					"has \"quotes\" inside": 456,
					"dots,commas.included": [
						{"subfield":1},
						{"another":2},
						{}
					],
					"hyphen-separated": "yup",
					"🐈": "everything for our furry friends"
				},
				"metadata": {
					"(braces[are]cool)": "object1"
				}
			}`,
			expected: `Mask<metadata."(braces[are]cool)",spec.("dots,commas.included".(0.subfield,1.another,2),"has \"quotes\" inside","hyphen-separated","with space","🐈")>`,
		},
		"not a json": {
			input: "}",
			err:   "json unmarshal: invalid character '}' looking for beginning of value",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			msk, err := ParseJSON([]byte(tt.input))
			if tt.err != "" {
				require.EqualError(t, err, tt.err)
			} else {
				require.NoError(t, err)

				require.Equal(t, tt.expected, msk.String())
			}
		})
	}
}
