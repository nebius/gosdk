package checknid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	example "github.com/nebius/gosdk/proto/nebius/example/compute/v1alpha1"
)

func TestCheckNIDFields(t *testing.T) {
	t.Parallel()

	instance := &example.Instance{
		Metadata: &common.ResourceMetadata{
			ParentId: "wrong-e0tabc",
		},
		Spec: &example.InstanceSpec{
			PlatformId: "other-e0tabc",
		},
	}

	warnings := CheckNIDFields(instance)
	require.Len(t, warnings, 1, warnings)
	require.Contains(t, warnings, "spec.platform_id")
	require.Contains(t, warnings["spec.platform_id"], "platform")
}

func TestCheckNIDFieldsValid(t *testing.T) {
	t.Parallel()

	instance := &example.Instance{
		Metadata: &common.ResourceMetadata{
			ParentId: "testresource-e0tabc",
		},
		Spec: &example.InstanceSpec{
			PlatformId: "platform-e0tabc",
		},
	}

	warnings := CheckNIDFields(instance)
	assert.Empty(t, warnings)
}

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
