package checknid

import (
	"testing"

	"github.com/stretchr/testify/assert"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

func TestCheckMetadataParentNID(t *testing.T) {
	t.Parallel()

	metadata := &common.ResourceMetadata{
		ParentId: "project-e0tabc",
	}

	assert.Empty(t, CheckMetadataParentNID(metadata, []string{"project"}))
	assert.Contains(t, CheckMetadataParentNID(metadata, []string{"folder"}), "folder")

	metadata.ParentId = ""
	assert.Empty(t, CheckMetadataParentNID(metadata, []string{"project"}))
}
