package grpcheader

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"

	"github.com/nebius/gosdk/fieldmask/mask"
)

func TestIsInOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("no metadata in headers", func(t *testing.T) {
		t.Parallel()
		assert.False(t, IsInOutgoingContext(context.Background(), "some-header"))
	})
	t.Run("different header", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.AppendToOutgoingContext(context.Background(), "Other-Header", "foo")
		assert.False(t, IsInOutgoingContext(ctx, "some-header"))
	})
	t.Run("has header same case", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.AppendToOutgoingContext(context.Background(), "Some-Header", "foo")
		assert.True(t, IsInOutgoingContext(ctx, "Some-Header"))
	})
	t.Run("has header diff case", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.AppendToOutgoingContext(context.Background(), "Some-Header", "foo")
		assert.True(t, IsInOutgoingContext(ctx, "some-heaDer"))
	})
}

func TestAddMaskToOutgoingContext(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()
		m := mask.New()
		m.Any = m
		_, err := AddMaskToOutgoingContext(context.Background(), "mymask", m)
		assert.EqualError(t, err, "failed to marshal mask: "+strings.Repeat("*: ", 1000)+"recursion too deep")
	})
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		m := mask.ParseMust("a,b,c.*")
		ctx, err := AddMaskToOutgoingContext(context.Background(), "My-Mask", m)
		assert.NoError(t, err)
		md, ok := metadata.FromOutgoingContext(ctx)
		assert.True(t, ok)
		assert.Equal(t, []string{"a,b,c.*"}, md["my-mask"])
	})
	t.Run("multi", func(t *testing.T) {
		t.Parallel()
		m := mask.ParseMust("a,b,c.*")
		ctx, err := AddMaskToOutgoingContext(context.Background(), "My-Mask", m)
		assert.NoError(t, err)
		m2 := mask.ParseMust("d,e.f")
		ctx, err = AddMaskToOutgoingContext(ctx, "My-Mask", m2)
		assert.NoError(t, err)
		md, ok := metadata.FromOutgoingContext(ctx)
		assert.True(t, ok)
		assert.Equal(t, []string{"a,b,c.*", "d,e.f"}, md["my-mask"])
	})
	t.Run("multi rev order", func(t *testing.T) {
		t.Parallel()
		m := mask.ParseMust("d,e.f")
		ctx, err := AddMaskToOutgoingContext(context.Background(), "My-Mask", m)
		assert.NoError(t, err)
		m2 := mask.ParseMust("a,b,c.*")
		ctx, err = AddMaskToOutgoingContext(ctx, "My-Mask", m2)
		assert.NoError(t, err)
		md, ok := metadata.FromOutgoingContext(ctx)
		assert.True(t, ok)
		assert.Equal(t, []string{"d,e.f", "a,b,c.*"}, md["my-mask"])
	})
}

func TestGetMaskFromIncomingContext(t *testing.T) {
	t.Parallel()
	t.Run("no metadata", func(t *testing.T) {
		t.Parallel()
		m, err := GetMaskFromIncomingContext(context.Background(), "my-mask")
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("metadata is nil", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), nil)
		m, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("metadata is empty", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{})
		m, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("metadata is without mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"other": []string{"foo", "bar"},
		})
		m, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("metadata is with empty mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"my-mask": []string{},
		})
		m, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.NoError(t, err)
		target := mask.ParseMust("")
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("metadata is with fail mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"my-mask": []string{"("},
		})
		_, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.EqualError(t, err, "failed parse mask 0: unclosed left brace at position 0 near \"⃞(\"")
	})
	t.Run("metadata is with one mask", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"my-mask": []string{"a,b.c"},
		})
		target := mask.ParseMust("a,b.c")
		m, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.NoError(t, err)
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
	t.Run("metadata is with couple masks", func(t *testing.T) {
		t.Parallel()
		ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{
			"my-mask": []string{"a,b.c", "b.d,*.e"},
		})
		target := mask.ParseMust("a,b.(c,d),*.e")
		m, err := GetMaskFromIncomingContext(ctx, "my-mask")
		assert.NoError(t, err)
		assert.True(t, target.Equal(m), "not equal", target, m)
	})
}
