package mask

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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