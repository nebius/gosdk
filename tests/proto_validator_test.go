package tests_test

import (
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

func TestProtoValidator(t *testing.T) {
	t.Parallel()
	validator, err := protovalidate.New()
	require.NoError(t, err)
	err = validator.Validate(&compute.Disk{
		Metadata: &common.ResourceMetadata{
			ParentId:        "",   // required string
			ResourceVersion: -123, // >= 0
		},
		Spec: &compute.DiskSpec{
			Size: nil, // required oneof
			Type: 0,   // required enum
		},
	})
	require.Error(t, err)
	assertViolation(t, err, "metadata.parent_id", "required")
	assertViolation(t, err, "metadata.resource_version", "int64.gte")
	assertViolation(t, err, "spec.size", "required")
	assertViolation(t, err, "spec.type", "required")
}

func assertViolation(t *testing.T, err error, field string, expectedConstraint string) {
	for _, line := range strings.Split(err.Error(), "\n") {
		_, after, found := strings.Cut(line, field)
		if found && !strings.HasPrefix(after, ".") {
			assert.Contains(t, line, expectedConstraint, "field %s", field)
			return
		}
	}
	assert.Fail(t, "violation not found", "field %s\nerror: %s", field, err)
}
