package grpcheader

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
	"github.com/nebius/gosdk/proto/fieldmask/protobuf/testdata"
)

func TestIsResetMaskInOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("no mask at all", func(t *testing.T) {
		t.Parallel()
		assert.False(t, IsResetMaskInOutgoingContext(context.Background()))
	})
	t.Run("different mask", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddMaskToOutgoingContext(context.Background(), "other-mask", mask.ParseMust("a,b"))
		assert.NoError(t, err)
		assert.False(t, IsResetMaskInOutgoingContext(ctx))
	})
	t.Run("has mask", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddMaskToOutgoingContext(context.Background(), "x-resetMask", mask.ParseMust("a,b"))
		assert.NoError(t, err)
		assert.True(t, IsResetMaskInOutgoingContext(ctx))
	})
}

func TestAddResetMaskToOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		m := mask.New()
		m.Any = m
		_, err := AddResetMaskToOutgoingContext(context.Background(), m)
		assert.EqualError(t, err, "failed to marshal mask: "+strings.Repeat("*: ", 1000)+"recursion too deep")
	})
	t.Run("add once", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddResetMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.c"))
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.c"}, md["x-resetmask"])
	})
	t.Run("add twice", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddResetMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.c"))
		assert.NoError(t, err)
		ctx, err = AddResetMaskToOutgoingContext(ctx, mask.ParseMust("a,b.d"))
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.c", "a,b.d"}, md["x-resetmask"])
	})
	t.Run("add twice reverse order", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddResetMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.d"))
		assert.NoError(t, err)
		ctx, err = AddResetMaskToOutgoingContext(ctx, mask.ParseMust("a,b.c"))
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.d", "a,b.c"}, md["x-resetmask"])
	})
}

func TestGetResetMaskFromIncomingContext(t *testing.T) {
	t.Parallel()
	t.Run("no mask", func(t *testing.T) {
		t.Parallel()
		m, err := GetResetMaskFromIncomingContext(context.Background())
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("other mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-selectmask": []string{"a,b.c"},
		})
		m, err := GetResetMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("empty mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-resetmask": []string{},
		})
		m, err := GetResetMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("one mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-resetmask": []string{"a,b.c"},
		})
		m, err := GetResetMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("a,b.c")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("multiple masks", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-resetmask": []string{"a,b.c", "b.d"},
		})
		m, err := GetResetMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("a,b.(c,d)")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
}

func TestAddMessageResetMaskToOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		recursionTooDeep := 1000
		veryRecursive := &testdata.RecursiveStruct{
			SomeString: "foo",
		}
		for range recursionTooDeep + 42 {
			veryRecursive = &testdata.RecursiveStruct{
				Recursive: veryRecursive,
			}
		}
		_, err := AddMessageResetMaskToOutgoingContext(context.Background(), veryRecursive)
		assert.EqualError(t, err, "failed to get mask from message: failed to collect modification mask: "+
			strings.Repeat("RecursiveStruct: ", recursionTooDeep)+"recursion too deep")
	})
	t.Run("success one", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddMessageResetMaskToOutgoingContext(context.Background(), &testdata.TestRecursiveSingle{})
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"field"}, md["x-resetmask"])
	})
	t.Run("success two", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddResetMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.d"))
		assert.NoError(t, err)
		ctx, err = AddMessageResetMaskToOutgoingContext(ctx, &testdata.TestRecursiveSingle{})
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.d", "field"}, md["x-resetmask"])
	})
}
func TestEnsureMessageResetMaskInOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		recursionTooDeep := 1000
		veryRecursive := &testdata.RecursiveStruct{
			SomeString: "foo",
		}
		for range recursionTooDeep + 42 {
			veryRecursive = &testdata.RecursiveStruct{
				Recursive: veryRecursive,
			}
		}
		_, err := EnsureMessageResetMaskInOutgoingContext(context.Background(), veryRecursive)
		assert.EqualError(t, err, "failed to get mask from message: failed to collect modification mask: "+
			strings.Repeat("RecursiveStruct: ", recursionTooDeep)+"recursion too deep")
	})
	t.Run("insert", func(t *testing.T) {
		t.Parallel()
		ctx, err := EnsureMessageResetMaskInOutgoingContext(context.Background(), &testdata.TestRecursiveSingle{})
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"field"}, md["x-resetmask"])
	})
	t.Run("no insert", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddResetMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.d"))
		assert.NoError(t, err)
		ctx, err = EnsureMessageResetMaskInOutgoingContext(ctx, &testdata.TestRecursiveSingle{})
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.d"}, md["x-resetmask"])
	})
}
