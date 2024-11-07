package tests_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/conn"
	commonpb "github.com/nebius/gosdk/proto/nebius/common/v1"
	computepb "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

func TestGoSDK(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	server := grpc.NewServer()
	computepb.RegisterDiskServiceServer(server, DiskServiceStub{})

	lis, err := net.Listen("tcp6", "[::1]:0")
	require.NoError(t, err)

	serveErr := make(chan error, 1)
	go func() {
		serveErr <- server.Serve(lis)
	}()
	t.Cleanup(func() {
		server.Stop()
		select {
		case errX := <-serveErr:
			require.NoError(t, errX)
		case <-time.After(5 * time.Second):
			require.Fail(t, "timeout")
		}
	})

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithResolvers(conn.NewConstantResolver(conn.Address(lis.Addr().String()))),
		gosdk.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		errX := sdk.Close()
		require.NoError(t, errX)
	})

	diskService := sdk.Services().Compute().V1().Disk()
	disk, err := diskService.Get(ctx, &computepb.GetDiskRequest{Id: "disk-id"})
	require.NoError(t, err)

	assert.EqualExportedValues(t, &computepb.Disk{
		Metadata: &commonpb.ResourceMetadata{
			Id:   "disk-id",
			Name: "disk-name",
		},
		Spec: &computepb.DiskSpec{
			Type: computepb.DiskSpec_NETWORK_SSD,
		},
	}, disk)
}

type DiskServiceStub struct {
	computepb.UnimplementedDiskServiceServer
}

func (DiskServiceStub) Get(context.Context, *computepb.GetDiskRequest) (*computepb.Disk, error) {
	return &computepb.Disk{
		Metadata: &commonpb.ResourceMetadata{
			Id:   "disk-id",
			Name: "disk-name",
		},
		Spec: &computepb.DiskSpec{
			Type: computepb.DiskSpec_NETWORK_SSD,
		},
	}, nil
}
