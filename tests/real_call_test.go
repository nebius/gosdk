package tests_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nebius/gosdk"
	computepb "github.com/nebius/gosdk/proto/nebius/compute/v1"
	"github.com/nebius/gosdk/serviceerror"
)

func TestRealCall_WithoutCreds(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	sdk, err := gosdk.New(ctx, gosdk.WithUserAgentPrefix("test-real-call-no-creds"))
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, sdk.Close())
	})

	diskService := sdk.Services().Compute().V1().Disk()
	_, err = diskService.List(ctx, &computepb.ListDisksRequest{ParentId: "project-e00aaaaaaaaaaaaaaa"})
	require.Error(t, err)
	require.Equal(t, codes.Unauthenticated, status.Code(err))
	var serviceErr *serviceerror.Error
	require.ErrorAs(t, err, &serviceErr)
	require.NotEmpty(t, serviceErr.RequestID)
}
