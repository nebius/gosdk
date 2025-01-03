package grpcheader

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/nebius/gosdk/proto/fieldmask/mask"
)

func TestIsSelectMaskInOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("no mask at all", func(t *testing.T) {
		t.Parallel()
		assert.False(t, IsSelectMaskInOutgoingContext(context.Background()))
	})
	t.Run("different mask", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddMaskToOutgoingContext(context.Background(), "other-mask", mask.ParseMust("a,b"))
		assert.NoError(t, err)
		assert.False(t, IsSelectMaskInOutgoingContext(ctx))
	})
	t.Run("has mask", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddMaskToOutgoingContext(context.Background(), "x-selectMask", mask.ParseMust("a,b"))
		assert.NoError(t, err)
		assert.True(t, IsSelectMaskInOutgoingContext(ctx))
	})
}

func TestAddSelectMaskToOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		m := mask.New()
		m.Any = m
		_, err := AddSelectMaskToOutgoingContext(context.Background(), m)
		assert.EqualError(t, err, "failed to marshal mask: "+strings.Repeat("*: ", 1000)+"recursion too deep")
	})
	t.Run("add once", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddSelectMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.c"))
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.c"}, md["x-selectmask"])
	})
	t.Run("add twice", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddSelectMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.c"))
		assert.NoError(t, err)
		ctx, err = AddSelectMaskToOutgoingContext(ctx, mask.ParseMust("a,b.d"))
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.c", "a,b.d"}, md["x-selectmask"])
	})
	t.Run("add twice reverse order", func(t *testing.T) {
		t.Parallel()
		ctx, err := AddSelectMaskToOutgoingContext(context.Background(), mask.ParseMust("a,b.d"))
		assert.NoError(t, err)
		ctx, err = AddSelectMaskToOutgoingContext(ctx, mask.ParseMust("a,b.c"))
		assert.NoError(t, err)
		md, _ := metadata.FromOutgoingContext(ctx)
		assert.Equal(t, []string{"a,b.d", "a,b.c"}, md["x-selectmask"])
	})
}

func TestGetSelectMaskFromIncomingContext(t *testing.T) {
	t.Parallel()
	t.Run("no mask", func(t *testing.T) {
		t.Parallel()
		m, err := GetSelectMaskFromIncomingContext(context.Background())
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("other mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-resetmask": []string{"a,b.c"},
		})
		m, err := GetSelectMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("empty mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-selectmask": []string{},
		})
		m, err := GetSelectMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("one mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-selectmask": []string{"a,b.c"},
		})
		m, err := GetSelectMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("a,b.c")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("multiple masks", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"x-selectmask": []string{"a,b.c", "b.d"},
		})
		m, err := GetSelectMaskFromIncomingContext(ctx)
		assert.NoError(t, err)
		target := mask.ParseMust("a,b.(c,d)")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
}
