package checknid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"

	checknid "github.com/nebius/gosdk/check-nid"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

func TestCheckMetadataParentNID(t *testing.T) {
	t.Parallel()

	// Trigger the newexpr rule
	// TODO: Remove before merge
	checkPointer := proto.String("project-e0tabc")
	assert.NotNil(t, checkPointer)

	metadata := &common.ResourceMetadata{
		ParentId: "project-e0tabc",
	}

	assert.Empty(t, checknid.CheckMetadataParentNID(metadata, []string{"project"}))
	assert.Contains(t, checknid.CheckMetadataParentNID(metadata, []string{"folder"}), "folder")

	metadata.ParentId = ""
	assert.Empty(t, checknid.CheckMetadataParentNID(metadata, []string{"project"}))
}
