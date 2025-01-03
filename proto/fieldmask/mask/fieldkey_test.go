package mask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldKey(t *testing.T) {
	t.Parallel()
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		fk := FieldKey("Simple_Key_123")
		str, err := fk.Marshal()
		assert.NoError(t, err)
		assert.Equal(t, "Simple_Key_123", str)
	})
	t.Run("quoted", func(t *testing.T) {
		t.Parallel()
		fk := FieldKey(`Key		with spaces, commas, and: ".\_-!@$%^"`)
		str, err := fk.Marshal()
		assert.NoError(t, err)
		assert.Equal(t, "\"Key\\t\\twith spaces, commas, and: \\\".\\\\_-!@$%^\\\"\"", str)
	})
	t.Run("marshal text", func(t *testing.T) {
		t.Parallel()
		fk := FieldKey("Simple_Key_123")
		bb, err := fk.MarshalText()
		assert.NoError(t, err)
		assert.Equal(t, []byte(`"Simple_Key_123"`), bb)
		n := FieldKey("")
		np := &n
		err = np.UnmarshalText(bb)
		assert.NoError(t, err)
		assert.Equal(t, fk, *np)
	})
}
