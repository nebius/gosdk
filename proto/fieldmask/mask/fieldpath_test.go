package mask

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldPath_Join(t *testing.T) {
	t.Parallel()
	cases := []struct {
		A FieldPath
		B FieldPath
		R FieldPath
	}{
		{
			A: FieldPath{},
			B: FieldPath{},
			R: FieldPath{},
		},
		{
			A: FieldPath{"foo"},
			B: FieldPath{},
			R: FieldPath{"foo"},
		},
		{
			A: FieldPath{"foo", "bar"},
			B: FieldPath{},
			R: FieldPath{"foo", "bar"},
		},
		{
			A: FieldPath{"foo", "bar"},
			B: FieldPath{"baz"},
			R: FieldPath{"foo", "bar", "baz"},
		},
		{
			A: FieldPath{"foo"},
			B: FieldPath{"bar", "baz"},
			R: FieldPath{"foo", "bar", "baz"},
		},
		{
			A: FieldPath{},
			B: FieldPath{"bar", "baz"},
			R: FieldPath{"bar", "baz"},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.R, c.A.Join(c.B...))
		})
	}
}

func TestFieldPath_Parent(t *testing.T) {
	t.Parallel()
	cases := []struct {
		A FieldPath
		R FieldPath
	}{
		{
			A: FieldPath{},
			R: nil,
		},
		{
			A: FieldPath{"foo"},
			R: FieldPath{},
		},
		{
			A: FieldPath{"foo", "bar"},
			R: FieldPath{"foo"},
		},
		{
			A: FieldPath{"foo", "bar", "baz"},
			R: FieldPath{"foo", "bar"},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.R, c.A.Parent())
		})
	}
}

func TestFieldPath_CopyEqualNew(t *testing.T) {
	t.Parallel()
	cases := []struct {
		A FieldPath
	}{
		{
			A: FieldPath{},
		},
		{
			A: FieldPath{"foo"},
		},
		{
			A: FieldPath{"foo", "bar"},
		},
		{
			A: FieldPath{"foo", "bar", "baz"},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			cp := c.A.Copy()
			assert.Equal(t, c.A, cp)
			assert.Equal(t, c.A, NewFieldPath(c.A...))
			assert.True(t, c.A.Equal(cp))
			assert.True(t, cp.Equal(c.A))
			if len(c.A) > 0 {
				c.A[0] = "changed"
				assert.NotEqual(t, c.A, cp)
				assert.False(t, c.A.Equal(cp))
				assert.False(t, cp.Equal(c.A))
			}
		})
	}
}
func TestFieldPath_Equal(t *testing.T) {
	t.Parallel()
	paths := []FieldPath{
		{},
		{"foo"},
		{"bar"},
		{"foo", "foo"},
		{"foo", "bar"},
		{"foo", "baz"},
		{"foe", "bar"},
		{"foe", "baz"},
		{"foo", "bar", "baz"},
		{"foo", "bae", "baz"},
		{"foe", "bar", "baz"},
		{"foo", "bar", "bax"},
		{"foo", "foo", "bar"},
		{"foo", "foo", "baz"},
		{"foo", "foo", "foo"},
	}
	for i, A := range paths {
		for j, B := range paths {
			t.Run(fmt.Sprintf("case %d %d", i, j), func(t *testing.T) {
				t.Parallel()
				assert.True(t, A.Equal(A)) //nolint:gocritic // we want to call with the same argument and receiver
				assert.True(t, B.Equal(B)) //nolint:gocritic // we want to call with the same argument and receiver
				if i == j {
					assert.True(t, A.Equal(B))
					assert.True(t, B.Equal(A))
				} else {
					assert.False(t, A.Equal(B))
					assert.False(t, B.Equal(A))
				}
			})
		}
	}
}

func TestFieldPath_ToMask(t *testing.T) {
	t.Parallel()
	cases := []struct {
		FP  FieldPath
		Res string
	}{
		{
			FP:  FieldPath{},
			Res: "",
		},
		{
			FP:  FieldPath{"foo"},
			Res: "foo",
		},
		{
			FP:  FieldPath{"foo", "bar"},
			Res: "foo.bar",
		},
		{
			FP:  FieldPath{"foo", "bar", "baz"},
			Res: "foo.bar.baz",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			m := c.FP.ToMask()
			str, err := m.Marshal()
			assert.NoError(t, err)
			assert.Equal(t, c.Res, str)
		})
	}
}

func TestFieldPath_IsPrefixOf(t *testing.T) {
	t.Parallel()
	cases := []struct {
		A   FieldPath
		B   FieldPath
		Res bool
	}{
		{
			A:   FieldPath{},
			B:   FieldPath{},
			Res: false,
		},
		{
			A:   FieldPath{"foo"},
			B:   FieldPath{},
			Res: false,
		},
		{
			A:   FieldPath{"foo"},
			B:   FieldPath{"bar"},
			Res: false,
		},
		{
			A:   FieldPath{"foo", "baz"},
			B:   FieldPath{"foo", "bar"},
			Res: false,
		},
		{
			A:   FieldPath{"foo", "baz", "abc"},
			B:   FieldPath{"foo", "bar", "abc", "def"},
			Res: false,
		},
		{
			A:   FieldPath{"baz", "foo"},
			B:   FieldPath{"bar", "foo", "abc", "def"},
			Res: false,
		},
		{
			A:   FieldPath{"baz", "foo", "abc"},
			B:   FieldPath{"bar", "foo", "abc", "def"},
			Res: false,
		},
		{
			A:   FieldPath{"bar", "foo", ""},
			B:   FieldPath{"bar", "foo", "abc", "def"},
			Res: false,
		},
		{
			A:   FieldPath{},
			B:   FieldPath{"foo"},
			Res: true,
		},
		{
			A:   FieldPath{"bar"},
			B:   FieldPath{"bar", "foo"},
			Res: true,
		},
		{
			A:   FieldPath{"bar", "baz"},
			B:   FieldPath{"bar", "baz", "foo"},
			Res: true,
		},
		{
			A:   FieldPath{},
			B:   FieldPath{"foo", "bar", "baz"},
			Res: true,
		},
		{
			A:   FieldPath{"bar"},
			B:   FieldPath{"bar", "foo", "baz"},
			Res: true,
		},
		{
			A:   FieldPath{"bar", "baz"},
			B:   FieldPath{"bar", "baz", "foo", "abc"},
			Res: true,
		},
		{
			A:   FieldPath{"bar"},
			B:   FieldPath{"bar", "baz", "foo", "abc"},
			Res: true,
		},
		{
			A:   FieldPath{},
			B:   FieldPath{"bar", "baz", "foo", "abc"},
			Res: true,
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			res := c.A.IsPrefixOf(c.B)
			assert.Equal(t, c.Res, res)
		})
	}
}

func TestFieldPath_MatchesResetMask(t *testing.T) {
	t.Parallel()
	cases := []struct {
		FP    FieldPath
		M     *Mask
		Res   bool
		Final bool
	}{
		{
			FP: FieldPath{},
			M:  nil,
		},
		{
			FP: FieldPath{"abc"},
			M:  nil,
		},
		{
			FP:    FieldPath{},
			M:     &Mask{},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc"},
			M:  &Mask{},
		},
		{
			FP: FieldPath{"abc"},
			M: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"abc": {},
				},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc"},
			M: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						Any: &Mask{},
					},
				},
			},
			Res: true,
		},
		{
			FP: FieldPath{"abc"},
			M: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						FieldParts: map[FieldKey]*Mask{
							"foo": {},
						},
					},
				},
			},
			Res: true,
		},
		{
			FP: FieldPath{"abc"},
			M: &Mask{
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						FieldParts: map[FieldKey]*Mask{
							"foo": {},
						},
					},
					"def": {
						FieldParts: map[FieldKey]*Mask{
							"foo": {},
						},
					},
				},
			},
			Res: true,
		},
		{
			FP: FieldPath{"abc", "foo"},
			M: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"foo": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
					"def": {
						FieldParts: map[FieldKey]*Mask{
							"baz": {},
						},
					},
				},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc", "bar"},
			M: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {
							Any: &Mask{},
						},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
					"def": {
						FieldParts: map[FieldKey]*Mask{
							"baz": {},
						},
					},
				},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc", "bar"},
			M: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {
								Any: &Mask{},
							},
						},
					},
					"def": {
						FieldParts: map[FieldKey]*Mask{
							"baz": {},
						},
					},
				},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc", "bar"},
			M: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"bar": {},
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"abc": {
						FieldParts: map[FieldKey]*Mask{
							"bar": {},
						},
					},
					"def": {
						FieldParts: map[FieldKey]*Mask{
							"baz": {},
						},
					},
				},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc"},
			M: &Mask{
				Any: &Mask{},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"abc"},
			M: &Mask{
				Any: &Mask{
					Any: &Mask{},
				},
			},
			Res: true,
		},
		{
			FP: FieldPath{"x"},
			M: &Mask{
				Any: New(),
				FieldParts: map[FieldKey]*Mask{
					"x": {
						FieldParts: map[FieldKey]*Mask{
							"y": New(),
						},
					},
				},
			},
			Res:   true,
			Final: true,
		},
		{
			FP: FieldPath{"x"},
			M: &Mask{
				Any: &Mask{
					FieldParts: map[FieldKey]*Mask{
						"y": New(),
					},
				},
				FieldParts: map[FieldKey]*Mask{
					"x": {},
				},
			},
			Res:   true,
			Final: true,
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			res := c.FP.MatchesResetMask(c.M)
			resFinal := c.FP.MatchesResetMaskFinal(c.M)
			assert.Equal(t, c.Res, res)
			assert.Equal(t, c.Final, resFinal)
		})
	}
}

func TestFieldPath_MatchesSelectMask(t *testing.T) {
	t.Parallel()
	cases := []struct {
		FP    FieldPath
		M     *Mask
		Res   bool
		Inner bool
	}{
		{
			FP:  FieldPath{},
			M:   nil,
			Res: true,
		},
		{
			FP:    FieldPath{"foo"},
			M:     nil,
			Res:   true,
			Inner: true,
		},
		{
			FP:    FieldPath{"foo", "bar"},
			M:     nil,
			Res:   true,
			Inner: true,
		},
		{
			FP:    FieldPath{"foo", "bar"},
			M:     New(),
			Res:   true,
			Inner: true,
		},
		{
			FP:    FieldPath{"foo"},
			M:     New(),
			Res:   true,
			Inner: true,
		},
		{
			FP:  FieldPath{},
			M:   New(),
			Res: true,
		},
		{
			FP:  FieldPath{},
			M:   ParseMust("a,b.c"),
			Res: true,
		},
		{
			FP:  FieldPath{"a"},
			M:   ParseMust("a,b.c"),
			Res: true,
		},
		{
			FP:  FieldPath{"b"},
			M:   ParseMust("a,b.c"),
			Res: true,
		},
		{
			FP:  FieldPath{"b", "c"},
			M:   ParseMust("a,b.c"),
			Res: true,
		},
		{
			FP:    FieldPath{"b", "c", "d"},
			M:     ParseMust("a,b.c"),
			Res:   true,
			Inner: true,
		},
		{
			FP:    FieldPath{"a", "c", "d"},
			M:     ParseMust("a,b.c"),
			Res:   true,
			Inner: true,
		},
		{
			FP: FieldPath{"b", "d"},
			M:  ParseMust("a,b.c"),
		},
		{
			FP: FieldPath{"d"},
			M:  ParseMust("a,b.c"),
		},
		{
			FP:  FieldPath{"a", "d"},
			M:   ParseMust("a.*"),
			Res: true,
		},
		{
			FP:    FieldPath{"a", "d", "e"},
			M:     ParseMust("a.*"),
			Res:   true,
			Inner: true,
		},
		{
			FP: FieldPath{"a", "d", "e"},
			M:  ParseMust("a.*.c"),
		},
		{
			FP:  FieldPath{"a", "d", "c"},
			M:   ParseMust("a.*.c"),
			Res: true,
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			res := c.FP.MatchesSelectMask(c.M)
			res2, inner := c.FP.MatchesSelectMaskInner(c.M)
			assert.Equal(t, c.Res, res)
			assert.Equal(t, c.Res, res2)
			assert.Equal(t, c.Inner, inner)
		})
	}
}

func TestFieldPath_Marshal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		FP  FieldPath
		M   string
		Err string
	}{
		{
			FP: FieldPath{},
			M:  "",
		},
		{
			FP: FieldPath{"foo"},
			M:  "foo",
		},
		{
			FP: FieldPath{"foo", "bar"},
			M:  "foo.bar",
		},
		{
			FP: FieldPath{"foo.bar"},
			M:  "\"foo.bar\"",
		},
		{
			FP: FieldPath{"foo.bar", "baz"},
			M:  "\"foo.bar\".baz",
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			res, err := c.FP.Marshal()
			r2, e2 := c.FP.MarshalText()
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
				assert.EqualError(t, e2, c.Err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, e2)
				assert.Equal(t, c.M, res)
				assert.Equal(t, c.M, string(r2))
			}
		})
	}
}

func TestFieldPath_UnarshalText(t *testing.T) {
	t.Parallel()
	cases := []struct {
		Mask      string
		NoStarter bool
		Err       string
		Result    FieldPath
	}{
		{
			NoStarter: true,
			Err:       "passed nil into unmarshal",
		},
		{
			Mask: "(",
			Err:  "unclosed left brace at position 0 near \"⃞(\"",
		},
		{
			Mask: "a,b",
			Err:  "multiple paths in the mask",
		},
		{
			Mask:   "a.b",
			Result: FieldPath{"a", "b"},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			t.Parallel()
			var s *FieldPath
			if !c.NoStarter {
				s = &FieldPath{}
			}
			err := s.UnmarshalText([]byte(c.Mask))
			if c.Err != "" {
				assert.EqualError(t, err, c.Err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.Result, *s)
			}
		})
	}
}
