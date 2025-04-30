package tests_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk"
	computepb "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

func TestRealCall_WithoutCreds(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	sdk, err := gosdk.New(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, sdk.Close())
	})

	diskService := sdk.Services().Compute().V1().Disk()
	_, err = diskService.List(ctx, &computepb.ListDisksRequest{ParentId: "fake-e00id"})
	require.Error(t, err)
	require.Regexp(t, "^rpc error: code = Unauthenticated desc = failed to list disks: "+
		"authorization failed: failed to extract token from context: no token request = [-a-zA-Z0-9]+$", err.Error())
}
