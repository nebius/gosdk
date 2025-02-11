package mask

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMask_IsEmpty(t *testing.T) {
	t.Parallel()
	t.Run("new mask", func(t *testing.T) {
		t.Parallel()
		mask := New()
		assert.True(t, mask.IsEmpty())
	})
	t.Run("default mask", func(t *testing.T) {
		t.Parallel()
		mask := &Mask{}
		assert.True(t, mask.IsEmpty())
	})
	t.Run("mask with empty map", func(t *testing.T) {
		t.Parallel()
		mask := &Mask{
			FieldParts: map[FieldKey]*Mask{},
		}
		assert.True(t, mask.IsEmpty())
	})
	t.Run("has any", func(t *testing.T) {
		t.Parallel()
		mask := &Mask{
			Any: New(),
		}
		assert.False(t, mask.IsEmpty())
	})
	t.Run("has key", func(t *testing.T) {
		t.Parallel()
		mask := &Mask{
			FieldParts: map[FieldKey]*Mask{
				"test": New(),
			},
		}
		assert.False(t, mask.IsEmpty())
	})
}

func TestMask_Marshal(t *testing.T) {
	t.Parallel()
	infiniteMask := New()
	infiniteMask.Any = infiniteMask
	infiniteMaskError := strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep"
	cases := []struct {
		Mask   *Mask
		Err    string
		Result string
	}{
		{
			Mask:   New(),
			Result: "",
		},
		{
			Mask:   &Mask{},
			Result: "",
		},
		{
			Mask: infiniteMask,
			Err:  "*: " + infiniteMaskError,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": infiniteMask,
				},
			},
			Err: "test: " + infiniteMaskError,
		},
		{
			Mask: &Mask{
				Any: New(),
			},
			Result: "*",
		},
		{
			Mask: &Mask{
				Any: &Mask{},
			},
			Result: "*",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": New(),
				},
			},
			Result: "test",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": nil,
				},
			},
			Result: "",
		},
		{
			Mask: &Mask{
				Any: New(),
				FieldParts: map[FieldKey]*Mask{
					"test": New(),
				},
			},
			Result: "*,test",
		},
		{
			Mask: &Mask{
				Any: New(),
				FieldParts: map[FieldKey]*Mask{
					"test": New(),
					"nil":  nil,
				},
			},
			Result: "*,test",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": New(),
					"foo":  New(),
				},
			},
			Result: "foo,test",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": New(),
					"foo":  New(),
					"nil":  nil,
				},
			},
			Result: "foo,test",
		},
		{
			Mask: &Mask{
				Any: New(),
				FieldParts: map[FieldKey]*Mask{
					"test": New(),
					"foo":  New(),
				},
			},
			Result: "*,foo,test",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					Any: New(),
				},
			},
			Result: "*.*",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					Any: New(),
					FieldParts: map[FieldKey]*Mask{
						"test": New(),
					},
				},
			},
			Result: "*.(*,test)",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					Any: New(),
					FieldParts: map[FieldKey]*Mask{
						"test": {
							Any: New(),
						},
					},
				},
			},
			Result: "*.(*,test.*)",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": {
							Any: New(),
						},
					},
				},
			},
			Result: "*.test.*",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					Any: &Mask{
						Any: New(),
					},
				},
			},
			Result: "*.*.*",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {
						Any: New(),
					},
				},
			},
			Result: "test.*",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": New(),
					},
				},
			},
			Result: "*.test",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {
						Any: New(),
					},
				},
			},
			Result: "test.*",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
						},
					},
				},
			},
			Result: "test.inner",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
							"nil":   nil,
						},
					},
				},
			},
			Result: "test.inner",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {
						FieldParts: map[FieldKey]*Mask{
							"inner-with-hyphen": New(),
							"nil":               nil,
						},
					},
				},
			},
			Result: "test.\"inner-with-hyphen\"",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test.inner": {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
						},
					},
				},
			},
			Result: `"test.inner".inner`,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test,inner": {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
						},
					},
				},
			},
			Result: `"test,inner".inner`,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test(inner)": {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
						},
					},
				},
			},
			Result: `"test(inner)".inner`,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					`"test"`: {},
				},
			},
			Result: `"\"test\""`,
		},
		{
			Mask: &Mask{
				Any: &Mask{},
				FieldParts: map[FieldKey]*Mask{
					`"test"`: {},
				},
			},
			Result: `"\"test\"",*`,
		},
		{
			Mask: &Mask{
				Any: &Mask{},
				FieldParts: map[FieldKey]*Mask{
					`",.() ` + "\t\r\n": {},
				},
			},
			Result: "\"\\\",.() \\t\\r\\n\",*",
		},
		{
			Mask: &Mask{
				Any: &Mask{},
				FieldParts: map[FieldKey]*Mask{
					`"test"`: {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
						},
					},
				},
			},
			Result: `"\"test\"".inner,*`,
		},
		{
			Mask: &Mask{
				Any: &Mask{},
				FieldParts: map[FieldKey]*Mask{
					`"test"`: {
						FieldParts: map[FieldKey]*Mask{
							"inner": New(),
							"*":     New(),
						},
					},
				},
			},
			Result: `"\"test\"".("*",inner),*`,
		},
		{
			Mask: &Mask{
				Any: &Mask{
					Any: &Mask{
						FieldParts: map[FieldKey]*Mask{
							"a": {},
							"b": {},
							"c": {
								Any: &Mask{},
							},
							"d": {
								FieldParts: map[FieldKey]*Mask{
									"e":   {},
									"f":   {},
									"nil": nil,
								},
							},
							"g": {
								Any: &Mask{
									FieldParts: map[FieldKey]*Mask{
										"h": {},
										"i": {},
									},
								},
								FieldParts: map[FieldKey]*Mask{
									"j":   {},
									"nil": nil,
									"k":   {},
								},
							},
						},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"l": {},
					"m": {},
					"n": {
						Any: &Mask{},
					},
					"o": {
						FieldParts: map[FieldKey]*Mask{
							"p":   {},
							"q":   {},
							"nil": nil,
						},
					},
					"r": {
						Any: &Mask{
							FieldParts: map[FieldKey]*Mask{
								"s": {},
								"t": {},
							},
						},
						FieldParts: map[FieldKey]*Mask{
							"u": {
								Any: &Mask{},
							},
							"v": {},
						},
					},
				},
			},
			Result: "*.*.(a,b,c.*,d.(e,f),g.(*.(h,i),j,k)),l,m,n.*,o.(p,q),r.(*.(s,t),u.*,v)",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					Any: &Mask{
						FieldParts: map[FieldKey]*Mask{
							"a":   {},
							`"b"`: {},
							"nil": nil,
							`"c"`: {
								Any: &Mask{},
							},
							`"d"`: {
								FieldParts: map[FieldKey]*Mask{
									"e":   {},
									`"f"`: {},
								},
							},
							"g": {
								Any: &Mask{
									FieldParts: map[FieldKey]*Mask{
										`"h"`: {},
										"i":   {},
										"nil": nil,
									},
								},
								FieldParts: map[FieldKey]*Mask{
									`"j"`: {},
									"k":   {},
									"nil": nil,
								},
							},
						},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"l":   {},
					"m":   {},
					"nil": nil,
					"n": {
						Any: &Mask{},
					},
					`"o"`: {
						FieldParts: map[FieldKey]*Mask{
							`"p"`: {},
							"q":   {},
							"*":   {},
						},
					},
					"r": {
						Any: &Mask{
							FieldParts: map[FieldKey]*Mask{
								"s": {},
								"t": {},
								"*": {},
							},
						},
						FieldParts: map[FieldKey]*Mask{
							"u": {
								Any: &Mask{},
							},
							"v": {},
							"*": {},
						},
					},
				},
			},
			Result: `"\"o\"".("*","\"p\"",q),*.*.("\"b\"","\"c\"".*,"\"d\"".("\"f\"",e),a,g.("\"j\"",*.("\"h\"",i),k)),l,m,n.*,r.("*",*.("*",s,t),u.*,v)`,
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
			str, err := c.Mask.Marshal()
			bb, err2 := c.Mask.MarshalText()
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
				assert.EqualError(t, err2, c.Err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, err2)
				assert.Equal(t, c.Result, str)
				assert.Equal(t, c.Result, string(bb))
			}
		})
	}
}

func TestMask_UnmarshalText(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Mask      string
		NoStarter bool
		Err       string
		Result    *Mask
	}{
		{
			NoStarter: true,
			Err:       "mask not initialized",
		},
		{
			Mask: "(",
			Err:  "unclosed left brace at position 0 near \"⃞(\"",
		},
		{
			Mask:   "",
			Result: &Mask{},
		},
		{
			Mask: "a",
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"a": {},
				},
			},
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
			var m *Mask
			if !c.NoStarter {
				m = &Mask{}
			}
			err := m.UnmarshalText([]byte(c.Mask))
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				assert.True(t, c.Result.Equal(m), "not equal:", m, c.Result)
			}
		})
	}
}

func TestMask_Equal(t *testing.T) {
	t.Parallel()
	someMask := &Mask{
		Any: &Mask{},
		FieldParts: map[FieldKey]*Mask{
			"test": {},
		},
	}
	copyMask, err := someMask.Copy()
	assert.NoError(t, err)
	t.Run("equals", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			A *Mask
			B *Mask
		}{
			{
				A: New(),
				B: New(),
			},
			{
				A: &Mask{},
				B: New(),
			},
			{
				A: someMask,
				B: someMask,
			},
			{
				A: someMask,
				B: copyMask,
			},
			{
				A: someMask,
				B: copyMask,
			},
			{
				A: &Mask{
					Any: &Mask{},
				},
				B: &Mask{
					Any: New(),
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": New()},
				},
			},
			{
				A: &Mask{
					FieldParts: nil,
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{},
				},
			},
			{
				A: &Mask{
					Any: someMask,
				},
				B: &Mask{
					Any: someMask,
				},
			},
			{
				A: &Mask{
					Any: someMask,
				},
				B: &Mask{
					Any: copyMask,
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
					},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
					},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
					},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": copyMask,
					},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
						"foo":  nil,
					},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": copyMask,
					},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
						"foo":  nil,
					},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": copyMask,
						"bar":  nil,
					},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": nil,
					},
				},
			},
			{
				A: &Mask{
					FieldParts: nil,
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": nil,
					},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"foo": nil,
					},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": nil,
					},
				},
			},
		}
		for i, c := range cases {
			name := fmt.Sprintf("case %d ", i)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				assert.True(t, c.A.Equal(c.B))
				assert.True(t, c.B.Equal(c.A))
			})
		}
	})
	t.Run("not equals", func(t *testing.T) {
		t.Parallel()
		infiniteMask := New()
		infiniteMask.Any = infiniteMask
		cases := []struct {
			A *Mask
			B *Mask
		}{
			{
				A: infiniteMask,
				B: infiniteMask,
			},
			{
				A: &Mask{},
				B: infiniteMask,
			},
			{
				A: someMask,
				B: infiniteMask,
			},
			{
				A: someMask,
				B: &Mask{},
			},
			{
				A: someMask,
				B: New(),
			},
			{
				A: &Mask{
					Any: nil,
				},
				B: &Mask{
					Any: &Mask{},
				},
			},
			{
				A: &Mask{},
				B: &Mask{
					Any: &Mask{},
				},
			},
			{
				A: &Mask{
					FieldParts: nil,
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
			},
			{
				A: &Mask{},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": nil},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{"foo": someMask},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": {}},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{"foo": someMask},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
			},
			{
				A: &Mask{
					Any: someMask,
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
			},
			{
				A: &Mask{
					Any:        someMask,
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
			},
			{
				A: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
						"foo":  someMask,
					},
				},
				B: &Mask{
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
			},
			{
				A: &Mask{
					Any: &Mask{},
					FieldParts: map[FieldKey]*Mask{
						"test": someMask,
						"foo":  someMask,
					},
				},
				B: &Mask{
					Any:        &Mask{},
					FieldParts: map[FieldKey]*Mask{"test": someMask},
				},
			},
		}
		for i, c := range cases {
			name := fmt.Sprintf("case %d ", i)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				assert.False(t, c.A.Equal(c.B))
				assert.False(t, c.B.Equal(c.A))
			})
		}
	})
}

func TestMask_Copy(t *testing.T) {
	t.Parallel()
	t.Run("simple", func(t *testing.T) {
		infiniteMask := New()
		infiniteMask.Any = infiniteMask
		infiniteMaskError := strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep"
		cases := []struct {
			Mask   *Mask
			Err    string
			Result string
		}{
			{
				Mask:   New(),
				Result: "",
			},
			{
				Mask:   &Mask{},
				Result: "",
			},
			{
				Mask: infiniteMask,
				Err:  "*: " + infiniteMaskError,
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": infiniteMask,
					},
				},
				Err: "test: " + infiniteMaskError,
			},
			{
				Mask: &Mask{
					Any: New(),
				},
				Result: "*",
			},
			{
				Mask: &Mask{
					Any: &Mask{},
				},
				Result: "*",
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": New(),
					},
				},
				Result: "test",
			},
			{
				Mask: &Mask{
					Any: New(),
					FieldParts: map[FieldKey]*Mask{
						"test": New(),
					},
				},
				Result: "*,test",
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": New(),
						"foo":  New(),
					},
				},
				Result: "foo,test",
			},
			{
				Mask: &Mask{
					Any: New(),
					FieldParts: map[FieldKey]*Mask{
						"test": New(),
						"foo":  New(),
					},
				},
				Result: "*,foo,test",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						Any: New(),
					},
				},
				Result: "*.*",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						Any: New(),
						FieldParts: map[FieldKey]*Mask{
							"test": New(),
						},
					},
				},
				Result: "*.(*,test)",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						Any: New(),
						FieldParts: map[FieldKey]*Mask{
							"test": {
								Any: New(),
							},
						},
					},
				},
				Result: "*.(*,test.*)",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						FieldParts: map[FieldKey]*Mask{
							"test": {
								Any: New(),
							},
						},
					},
				},
				Result: "*.test.*",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						Any: &Mask{
							Any: New(),
						},
					},
				},
				Result: "*.*.*",
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": {
							Any: New(),
						},
					},
				},
				Result: "test.*",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						FieldParts: map[FieldKey]*Mask{
							"test": New(),
						},
					},
				},
				Result: "*.test",
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": {
							Any: New(),
						},
					},
				},
				Result: "test.*",
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test": {
							FieldParts: map[FieldKey]*Mask{
								"inner": New(),
							},
						},
					},
				},
				Result: "test.inner",
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test.inner": {
							FieldParts: map[FieldKey]*Mask{
								"inner": New(),
							},
						},
					},
				},
				Result: `"test.inner".inner`,
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test,inner": {
							FieldParts: map[FieldKey]*Mask{
								"inner": New(),
							},
						},
					},
				},
				Result: `"test,inner".inner`,
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"test(inner)": {
							FieldParts: map[FieldKey]*Mask{
								"inner": New(),
							},
						},
					},
				},
				Result: `"test(inner)".inner`,
			},
			{
				Mask: &Mask{
					FieldParts: map[FieldKey]*Mask{
						`"test"`: {},
					},
				},
				Result: `"\"test\""`,
			},
			{
				Mask: &Mask{
					Any: &Mask{},
					FieldParts: map[FieldKey]*Mask{
						`"test"`: {},
					},
				},
				Result: `"\"test\"",*`,
			},
			{
				Mask: &Mask{
					Any: &Mask{},
					FieldParts: map[FieldKey]*Mask{
						`"test"`: {
							FieldParts: map[FieldKey]*Mask{
								"inner": New(),
							},
						},
					},
				},
				Result: `"\"test\"".inner,*`,
			},
			{
				Mask: &Mask{
					Any: &Mask{},
					FieldParts: map[FieldKey]*Mask{
						`"test"`: {
							FieldParts: map[FieldKey]*Mask{
								"inner": New(),
								"*":     New(),
							},
						},
					},
				},
				Result: `"\"test\"".("*",inner),*`,
			},
			{
				Mask: &Mask{
					Any: &Mask{
						Any: &Mask{
							FieldParts: map[FieldKey]*Mask{
								"a": {},
								"b": {},
								"c": {
									Any: &Mask{},
								},
								"d": {
									FieldParts: map[FieldKey]*Mask{
										"e": {},
										"f": {},
									},
								},
								"g": {
									Any: &Mask{
										FieldParts: map[FieldKey]*Mask{
											"h": {},
											"i": {},
										},
									},
									FieldParts: map[FieldKey]*Mask{
										"j": {},
										"k": {},
									},
								},
							},
						},
					},
					FieldParts: map[FieldKey]*Mask{
						"l": {},
						"m": {},
						"n": {
							Any: &Mask{},
						},
						"o": {
							FieldParts: map[FieldKey]*Mask{
								"p": {},
								"q": {},
							},
						},
						"r": {
							Any: &Mask{
								FieldParts: map[FieldKey]*Mask{
									"s": {},
									"t": {},
								},
							},
							FieldParts: map[FieldKey]*Mask{
								"u": {
									Any: &Mask{},
								},
								"v": {},
							},
						},
					},
				},
				Result: "*.*.(a,b,c.*,d.(e,f),g.(*.(h,i),j,k)),l,m,n.*,o.(p,q),r.(*.(s,t),u.*,v)",
			},
			{
				Mask: &Mask{
					Any: &Mask{
						Any: &Mask{
							FieldParts: map[FieldKey]*Mask{
								"a":   {},
								`"b"`: {},
								`"c"`: {
									Any: &Mask{},
								},
								`"d"`: {
									FieldParts: map[FieldKey]*Mask{
										"e":   {},
										`"f"`: {},
									},
								},
								"g": {
									Any: &Mask{
										FieldParts: map[FieldKey]*Mask{
											`"h"`: {},
											"i":   {},
										},
									},
									FieldParts: map[FieldKey]*Mask{
										`"j"`: {},
										"k":   {},
									},
								},
							},
						},
					},
					FieldParts: map[FieldKey]*Mask{
						"l": {},
						"m": {},
						"n": {
							Any: &Mask{},
						},
						`"o"`: {
							FieldParts: map[FieldKey]*Mask{
								`"p"`: {},
								"q":   {},
								"*":   {},
							},
						},
						"r": {
							Any: &Mask{
								FieldParts: map[FieldKey]*Mask{
									"s": {},
									"t": {},
									"*": {},
								},
							},
							FieldParts: map[FieldKey]*Mask{
								"u": {
									Any: &Mask{},
								},
								"v": {},
								"*": {},
							},
						},
					},
				},
				Result: `"\"o\"".("*","\"p\"",q),*.*.("\"b\"","\"c\"".*,"\"d\"".("\"f\"",e),a,g.("\"j\"",*.("\"h\"",i),k)),l,m,n.*,r.("*",*.("*",s,t),u.*,v)`,
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
				mask, err := c.Mask.Copy()
				if c.Err != "" {
					assert.EqualError(t, err, c.Err)
				} else {
					assert.NoError(t, err)
					str, err := mask.Marshal()
					assert.NoError(t, err)
					assert.Equal(t, c.Result, str)
				}
			})
		}
	})
	t.Run("deeptest", func(t *testing.T) {
		t.Parallel()
		mask := Mask{
			Any: &Mask{},
			FieldParts: map[FieldKey]*Mask{
				"foo": {},
				"bar": nil, // should be removed in copy
			},
		}
		cp, err := mask.Copy()
		assert.NoError(t, err)
		_, ok := cp.FieldParts["bar"]
		assert.False(t, ok)
		str1, err := cp.Marshal()
		assert.NoError(t, err)
		mask.Any.Any = &Mask{}
		str2, err := cp.Marshal()
		assert.NoError(t, err)
		assert.Equal(t, str1, str2)
		mask.FieldParts["a"] = &Mask{}
		str3, err := cp.Marshal()
		assert.NoError(t, err)
		assert.Equal(t, str1, str3)
		mask.FieldParts["foo"].Any = &Mask{}
		str4, err := cp.Marshal()
		assert.NoError(t, err)
		assert.Equal(t, str1, str4)
	})
}

func TestMask_Merge(t *testing.T) {
	t.Parallel()
	infMask := New()
	infMask.Any = infMask
	infiniteMaskSrc := New()
	infiniteMaskSrc.Any = infiniteMaskSrc
	cases := []struct {
		A      *Mask
		B      *Mask
		Err    string
		Result string
	}{
		{
			A:   infiniteMaskSrc,
			B:   infMask,
			Err: "*:M " + strings.Repeat("*:M ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			A:   &Mask{},
			B:   infMask,
			Err: "*:C " + strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			A: &Mask{},
			B: &Mask{
				Any: infMask,
			},
			Err: "*:C " + strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			A:      &Mask{},
			B:      &Mask{},
			Result: "",
		},
		{
			A: &Mask{
				Any: &Mask{},
			},
			B:      &Mask{},
			Result: "*",
		},
		{
			A: &Mask{
				Any: &Mask{},
			},
			B: &Mask{
				Any: &Mask{},
			},
			Result: "*",
		},
		{
			A: &Mask{
				Any: &Mask{
					Any: &Mask{},
				},
			},
			B: &Mask{
				Any: &Mask{},
			},
			Result: "*.*",
		},
		{
			B: &Mask{
				Any: &Mask{
					Any: &Mask{},
				},
			},
			A: &Mask{
				Any: &Mask{},
			},
			Result: "*.*",
		},
		{
			A: &Mask{
				Any: &Mask{},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {},
				},
			},
			Result: "*,test",
		},
		{
			A: &Mask{
				Any: &Mask{},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {},
					"nil":  nil,
				},
			},
			Result: "*,test",
		},
		{
			A: &Mask{
				Any: &Mask{},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"nil": nil,
				},
			},
			Result: "*",
		},
		{
			A: &Mask{
				Any: &Mask{
					Any: &Mask{
						FieldParts: map[FieldKey]*Mask{
							"a": {},
							"b": {},
							"c": {
								Any: &Mask{},
							},
							"d": {
								FieldParts: map[FieldKey]*Mask{
									"e":   {},
									"f":   {},
									"nil": nil,
								},
							},
							"g": {
								Any: &Mask{
									FieldParts: map[FieldKey]*Mask{
										"h": {},
										"i": {},
									},
								},
								FieldParts: map[FieldKey]*Mask{
									"j":   {},
									"nil": nil,
									"k":   {},
								},
							},
						},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"l": {},
					"m": {},
					"n": {
						Any: &Mask{},
					},
					"o": {
						FieldParts: map[FieldKey]*Mask{
							"p":   {},
							"q":   {},
							"nil": nil,
						},
					},
					"r": {
						Any: &Mask{
							FieldParts: map[FieldKey]*Mask{
								"s": {},
								"t": {},
							},
						},
						FieldParts: map[FieldKey]*Mask{
							"u": {
								Any: &Mask{},
							},
							"v": {},
						},
					},
				},
			},
			B: &Mask{
				Any: &Mask{
					Any: &Mask{
						FieldParts: map[FieldKey]*Mask{
							"a": {},
							"b": {},
							"c": {
								Any: &Mask{},
							},
							"d": {
								FieldParts: map[FieldKey]*Mask{
									"e":   {},
									"f":   {},
									"nil": nil,
								},
							},
							"g": {
								Any: &Mask{
									FieldParts: map[FieldKey]*Mask{
										"h": {},
										"i": {},
									},
								},
								FieldParts: map[FieldKey]*Mask{
									"j":   {},
									"nil": nil,
									"k":   {},
								},
							},
						},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"l": {},
					"A": {},
					"B": {},
					"1": {},
					"o": {
						FieldParts: map[FieldKey]*Mask{
							"w":   {},
							"x":   {},
							"nil": nil,
						},
					},
					"z": {
						Any: &Mask{
							FieldParts: map[FieldKey]*Mask{
								"y": {},
							},
						},
						FieldParts: map[FieldKey]*Mask{
							"u": {
								Any: &Mask{},
							},
							"v": {},
						},
					},
				},
			},
			Result: "*.*.(a,b,c.*,d.(e,f),g.(*.(h,i),j,k)),1,A,B,l,m,n.*,o.(p,q,w,x),r.(*.(s,t),u.*,v),z.(*.y,u.*,v)",
		},
		{
			A: &Mask{},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"inf": infMask,
				},
			},
			Err: "inf:C " + strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			A: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"inf": {},
				},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"inf": infMask,
				},
			},
			Err: "inf:M *:C " + strings.Repeat("*: ", recursionTooDeep-2) + "recursion too deep",
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
			err := c.A.Merge(c.B)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				str, err := c.A.Marshal()
				assert.NoError(t, err)
				assert.Equal(t, c.Result, str)
			}
		})
	}
}

func TestMask_String(t *testing.T) {
	t.Parallel()

	infMask := New()
	infMask.Any = infMask
	cases := []struct {
		Mask   *Mask
		Result string
	}{
		{
			Mask:   infMask,
			Result: "Mask<not-marshalable " + strings.Repeat("*: ", recursionTooDeep) + "recursion too deep" + ">",
		},
		{
			Mask:   &Mask{},
			Result: "Mask<>",
		},
		{
			Mask:   New(),
			Result: "Mask<>",
		},
		{
			Mask: &Mask{
				Any: &Mask{},
			},
			Result: "Mask<*>",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"test": {},
				},
			},
			Result: "Mask<test>",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"a": {},
						"b": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"c": {
						Any: &Mask{},
					},
					"d": {},
				},
			},
			Result: "Mask<*.(a,b),c.*,d>",
		},
	}
	for i, c := range cases {
		name := fmt.Sprintf("case %d ", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.Result, c.Mask.String())
		})
	}
}

func TestMask_ToFieldPath(t *testing.T) {
	t.Parallel()
	infMask := New()
	infMask.FieldParts["x"] = infMask
	cases := []struct {
		Mask   *Mask
		Err    string
		Result FieldPath
	}{
		{
			Mask:   &Mask{},
			Result: nil,
		},
		{
			Mask: infMask,
			Err:  strings.Repeat("x: ", recursionTooDeep) + "recursion too deep",
		},
		{
			Mask: &Mask{
				Any: &Mask{},
			},
			Err: "wildcard in the mask",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {},
					"bar": {},
				},
			},
			Err: "multiple paths in the mask",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						Any: &Mask{},
					},
				},
			},
			Err: "foo: wildcard in the mask",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
							"baz": {},
						},
					},
				},
			},
			Err: "foo: multiple paths in the mask",
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {},
				},
			},
			Result: FieldPath{
				"foo",
			},
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Result: FieldPath{
				"foo",
				"bar",
			},
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
			fp, err := c.Mask.ToFieldPath()
			isFp := c.Mask.IsFieldPath()
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
				assert.False(t, isFp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.Result, fp)
				assert.True(t, isFp)
			}
		})
	}
}

func TestMask_GetSubMask(t *testing.T) {
	t.Parallel()
	infMask := New()
	infMask.FieldParts["x"] = infMask
	cases := []struct {
		Mask   *Mask
		Key    FieldKey
		Err    string
		Result *Mask
	}{
		{
			Mask:   &Mask{},
			Key:    "foo",
			Result: nil,
		},
		{
			Mask:   nil,
			Key:    "foo",
			Result: nil,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
			Key:    "foo",
			Result: nil,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Key: "foo",
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{},
			},
			Key: "foo",
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Key: "foo",
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Key: "foo",
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
					"baz": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": infMask,
						},
					},
				},
			},
			Key: "foo",
			Err: "failed to copy key mask: bar: " + strings.Repeat("x: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": infMask,
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Key: "foo",
			Err: "failed to merge key mask with wildcard mask: baz:C " + strings.Repeat("x: ", recursionTooDeep-1) + "recursion too deep",
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
			fm, err := c.Mask.GetSubMask(c.Key)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				if c.Result == nil {
					assert.Nil(t, fm)
				} else {
					assert.True(t, c.Result.Equal(fm), "not equal", fm, c.Result)
				}
			}
		})
	}
}

func TestMask_GetSubMaskByPath(t *testing.T) {
	t.Parallel()
	infMask := New()
	infMask.FieldParts["x"] = infMask
	cases := []struct {
		Mask   *Mask
		Path   FieldPath
		Err    string
		Result *Mask
	}{
		{
			Mask:   &Mask{},
			Path:   FieldPath{"foo"},
			Result: nil,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
			Path:   FieldPath{"foo"},
			Result: nil,
		},
		{
			Mask: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Path: FieldPath{"foo"},
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{},
			},
			Path: FieldPath{"foo"},
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Path: FieldPath{"foo"},
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Path: FieldPath{"foo"},
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"bar": {},
					"baz": {},
				},
			},
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": infMask,
						},
					},
				},
			},
			Path: FieldPath{"foo"},
			Err:  "failed to get field mask at path : failed to copy key mask: bar: " + strings.Repeat("x: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": infMask,
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Path: FieldPath{"foo"},
			Err:  "failed to get field mask at path : failed to merge key mask with wildcard mask: baz:C " + strings.Repeat("x: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"baz": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
				},
			},
			Path:   FieldPath{"foo", "bar", "baz"},
			Result: nil,
		},
		{
			Mask: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {
							Any: &Mask{
								FieldParts: map[FieldKey]*Mask{
									"a": {},
								},
							},
							FieldParts: map[FieldKey]*Mask{
								"baz": {
									FieldParts: map[FieldKey]*Mask{
										"b": {},
									},
								},
								"other": {
									FieldParts: map[FieldKey]*Mask{
										"nope": {},
									},
								},
							},
						},
						"hey": {
							Any: &Mask{
								FieldParts: map[FieldKey]*Mask{
									"never": {},
								},
							},
						},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"foo": {
						Any: &Mask{
							Any: &Mask{
								FieldParts: map[FieldKey]*Mask{
									"f": {},
								},
							},
							FieldParts: map[FieldKey]*Mask{
								"baz": {
									FieldParts: map[FieldKey]*Mask{
										"c": {},
									},
								},
							},
						},
						FieldParts: map[FieldKey]*Mask{
							"bar": {
								Any: &Mask{
									FieldParts: map[FieldKey]*Mask{
										"d": {},
									},
								},
								FieldParts: map[FieldKey]*Mask{
									"baz": {
										FieldParts: map[FieldKey]*Mask{
											"e": {},
										},
									},
								},
							},
						},
					},
				},
			},
			Path: FieldPath{"foo", "bar", "baz"},
			Result: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"a": {},
					"b": {},
					"c": {},
					"d": {},
					"e": {},
					"f": {},
				},
			},
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
			fm, err := c.Mask.GetSubMaskByPath(c.Path)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				if c.Result == nil {
					assert.Nil(t, fm)
				} else {
					assert.True(t, c.Result.Equal(fm), "not equal", fm, c.Result)
				}
			}
		})
	}
}

func TestMask_IntersectDumb(t *testing.T) {
	t.Parallel()
	infiMask := &Mask{}
	infiMask.Any = infiMask
	cases := []struct {
		A   *Mask
		B   *Mask
		Res *Mask
		Err string
	}{
		{
			A:   nil,
			B:   nil,
			Res: nil,
		},
		{
			A:   &Mask{},
			B:   nil,
			Res: nil,
		},
		{
			A:   nil,
			B:   &Mask{},
			Res: nil,
		},
		{
			A:   ParseMust("*.(a,b),x,y"),
			B:   ParseMust("*.(a,b),x,y"),
			Res: ParseMust("*.(a,b),x,y"),
		},
		{
			A:   ParseMust("a"),
			B:   ParseMust("*"),
			Res: ParseMust(""),
		},
		{
			A:   ParseMust("*.(a,b),x,y"),
			B:   ParseMust("*.(x,y),z,f"),
			Res: ParseMust("*"),
		},
		{
			A:   ParseMust("*.(a,b),x,y"),
			B:   ParseMust("*.(a,y),x,f"),
			Res: ParseMust("*.a,x"),
		},
		{
			A:   infiMask,
			B:   infiMask,
			Err: strings.Repeat("*: ", recursionTooDeep) + "recursion too deep",
		},
		{
			A: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			Err: "x: " + strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep",
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
			im, err := c.A.IntersectDumb(c.B)
			im2, err2 := c.B.IntersectDumb(c.A)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
				assert.EqualError(t, err2, c.Err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, err2)
				if c.Res == nil {
					assert.Nil(t, im)
					assert.Nil(t, im2)
				} else {
					assert.True(t, c.Res.Equal(im), "not equal", im, c.Res)
					assert.True(t, c.Res.Equal(im2), "not equal", im2, c.Res)
				}
			}
		})
	}
}

func TestMask_IntersectResetMask(t *testing.T) {
	t.Parallel()
	infiMask := &Mask{}
	infiMask.Any = infiMask
	cases := []struct {
		A   *Mask
		B   *Mask
		Res *Mask
		Err string
	}{
		{
			A:   nil,
			B:   nil,
			Res: nil,
		},
		{
			A:   &Mask{},
			B:   nil,
			Res: nil,
		},
		{
			A:   nil,
			B:   &Mask{},
			Res: nil,
		},
		{
			A:   ParseMust("a"),
			B:   ParseMust("*"),
			Res: ParseMust("a"),
		},
		{
			A:   ParseMust("a.(x,y,z),b.*"),
			B:   ParseMust("*.x,a.z"),
			Res: ParseMust("a.(x,z),b.x"),
		},
		{
			A:   infiMask,
			B:   infiMask,
			Err: strings.Repeat("*×*: ", recursionTooDeep) + "recursion too deep",
		},
		{
			A: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			Err: "x×x: " + strings.Repeat("*×*: ", recursionTooDeep-1) + "recursion too deep",
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
			im, err := c.A.IntersectResetMask(c.B)
			im2, err2 := c.B.IntersectResetMask(c.A)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
				assert.EqualError(t, err2, c.Err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, err2)
				if c.Res == nil {
					assert.Nil(t, im)
					assert.Nil(t, im2)
				} else {
					assert.True(t, c.Res.Equal(im), "not equal", im, c.Res)
					assert.True(t, c.Res.Equal(im2), "not equal", im2, c.Res)
				}
			}
		})
	}
}

func TestMask_SubtractDumb(t *testing.T) {
	t.Parallel()
	infiMask := &Mask{}
	infiMask.Any = infiMask
	cases := []struct {
		A   *Mask
		B   *Mask
		Res *Mask
		Err string
	}{
		{
			A:   nil,
			B:   nil,
			Res: nil,
		},
		{
			A:   &Mask{},
			B:   nil,
			Res: &Mask{},
		},
		{
			A:   nil,
			B:   &Mask{},
			Res: nil,
		},
		{
			A:   ParseMust("x.(a,b),*.(c,d),e,f"),
			B:   ParseMust("x.(a,b),*.(c,d),e,f"),
			Res: ParseMust(""),
		},
		{
			A:   ParseMust("x.(a,b),*.(c,d),e,f"),
			B:   ParseMust("x.(a),*.(c),e"),
			Res: ParseMust("x.b,*.d,f"),
		},
		{
			A:   ParseMust("a"),
			B:   ParseMust("*"),
			Res: ParseMust("a"),
		},
		{
			A: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			Err: "x: " + strings.Repeat("*: ", recursionTooDeep-1) + "recursion too deep",
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
			err := c.A.SubtractDumb(c.B)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				if c.Res == nil {
					assert.Nil(t, c.A)
				} else {
					assert.True(t, c.Res.Equal(c.A), "not equal", c.A, c.Res)
				}
			}
		})
	}
}

func TestMask_SubtractResetMask(t *testing.T) {
	t.Parallel()
	infiMask := &Mask{}
	infiMask.Any = infiMask
	cases := []struct {
		A   *Mask
		B   *Mask
		Res *Mask
		Err string
	}{
		{
			A:   nil,
			B:   nil,
			Res: nil,
		},
		{
			A:   &Mask{},
			B:   nil,
			Res: &Mask{},
		},
		{
			A:   nil,
			B:   &Mask{},
			Res: nil,
		},
		{
			A:   ParseMust("x.(a,b),*.(c,d),e,f"),
			B:   ParseMust("x.(a,b),*.(c,d),e,f"),
			Res: ParseMust(""),
		},
		{
			A:   ParseMust("x.(a,b),*.(c,d),e,f"),
			B:   ParseMust("x.(a),*.(c),e"),
			Res: ParseMust("x.b,*.d"),
		},
		{
			A: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			B: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			Err: "x\\x: " + strings.Repeat("*\\*: ", recursionTooDeep-1) + "recursion too deep",
		},
		{
			A: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"x": infiMask,
				},
			},
			B:   infiMask,
			Err: "x\\*: " + strings.Repeat("*\\*: ", recursionTooDeep-1) + "recursion too deep",
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
			err := c.A.SubtractResetMask(c.B)
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				if c.Res == nil {
					assert.Nil(t, c.A)
				} else {
					assert.True(t, c.Res.Equal(c.A), "not equal", c.A, c.Res)
				}
			}
		})
	}
}
