package checknid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	checknid "github.com/nebius/gosdk/check-nid"
)

func TestIsNIDAllowedForAutoFill(t *testing.T) {
	t.Parallel()

	assert.True(t, checknid.IsNIDAllowedForAutoFill("project-e0tabc", []string{"project"}))
	assert.False(t, checknid.IsNIDAllowedForAutoFill("computedisk-e0tabc", []string{"project"}))
	assert.False(t, checknid.IsNIDAllowedForAutoFill("invalid", []string{"project"}))
	assert.True(t, checknid.IsNIDAllowedForAutoFill("computedisk-e0tabc", nil))
	assert.True(t, checknid.IsNIDAllowedForAutoFill("computedisk-e0tabc", []string{"*"}))
}

func TestMetadataParentAutoFillTypes(t *testing.T) {
	t.Parallel()

	metadataParent := parentResourceSubfield("metadata", []string{"project"})
	settings := []*checknid.SubfieldSettings{
		parentResourceSubfield("spec", []string{"folder"}),
		metadataParent,
		nidSubfield("metadata.parent_id.extra", []string{"folder"}),
		nidSubfield("metadata.parent_id", []string{"organization"}),
	}

	want := []string{"project"}
	require.Equal(t, want, checknid.MetadataParentAutoFillTypes(settings, "metadata"))
}

func TestSubfieldSettingMatchesPath(t *testing.T) {
	t.Parallel()

	require.True(t, checknid.SubfieldSettingMatchesPath(
		"metadata.parent_id",
		[]string{"metadata", "parent_id"},
	))
	require.False(t, checknid.SubfieldSettingMatchesPath(
		"metadata.parent_id.extra",
		[]string{"metadata", "parent_id"},
	))
}

func nidSubfield(fieldPath string, resources []string) *checknid.SubfieldSettings {
	return &checknid.SubfieldSettings{
		FieldPath: fieldPath,
		Nid: &checknid.NIDFieldSettings{
			Resource: resources,
		},
	}
}

func parentResourceSubfield(fieldPath string, resources []string) *checknid.SubfieldSettings {
	return &checknid.SubfieldSettings{
		FieldPath: fieldPath,
		Nid: &checknid.NIDFieldSettings{
			ParentResource: resources,
		},
	}
}
