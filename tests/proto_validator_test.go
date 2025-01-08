package tests_test

import (
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
	var vErr *protovalidate.ValidationError
	require.ErrorAs(t, err, &vErr)
	assertViolation(t, vErr, "metadata.parent_id", "required")
	assertViolation(t, vErr, "metadata.resource_version", "int64.gte")
	assertViolation(t, vErr, "spec.size", "required")
	assertViolation(t, vErr, "spec.type", "required")
}

func assertViolation(t *testing.T, err *protovalidate.ValidationError, field string, expectedConstraint string) {
	require.NotNil(t, err)
	for _, v := range err.Violations {
		if v.GetFieldPath() == field {
			assert.Equal(t, expectedConstraint, v.GetConstraintId())
			return
		}
	}
	assert.Fail(t, "violation not found", "field %s", field)
}
