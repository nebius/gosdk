package conn_test

import (
	"context"
	"fmt"
	"log/slog"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk/conn"
)

func TestConventionResolver_positive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id      conn.ServiceID
		address conn.Address
	}{
		{
			id:      "nebius.compute.InstanceService",
			address: "compute.{domain}",
		},
		{
			id:      "nebius.msp.postgres.inner.v1.ClusterService",
			address: "msp.{domain}",
		},
	}
	for _, test := range tests {
		t.Run(string(test.id), func(t *testing.T) {
			t.Parallel()

			address, err := conn.NewConventionResolver().Resolve(t.Context(), test.id)
			require.NoError(t, err)
			assert.Equal(t, test.address, address)
		})
	}
}

type countingResolver struct {
	calls atomic.Int64
}

func (r *countingResolver) Resolve(context.Context, conn.ServiceID) (conn.Address, error) {
	return conn.Address(fmt.Sprintf("override-%d.example:443", r.calls.Add(1))), nil
}

func TestRoutingResolverCompositionAndCache(t *testing.T) {
	t.Parallel()

	id := conn.ServiceID("nebius.test.v1.CachedService")
	base := &countingResolver{}
	logger := slog.New(slog.DiscardHandler)
	resolver := conn.NewContextResolver(logger, conn.NewCachedResolver(logger, base))

	address, err := resolver.Resolve(t.Context(), id)
	require.NoError(t, err)
	assert.Equal(t, conn.Address("override-1.example:443"), address)

	address, err = resolver.Resolve(t.Context(), id)
	require.NoError(t, err)
	assert.Equal(t, conn.Address("override-1.example:443"), address)
	assert.Equal(t, int64(1), base.calls.Load())

	ctx := conn.ContextWithResolver(t.Context(), conn.NewConstantResolver("context.example:443"))
	address, err = resolver.Resolve(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, conn.Address("context.example:443"), address)
}

func TestConventionResolver_negative(t *testing.T) {
	t.Parallel()

	tests := []conn.ServiceID{
		"not.nebius.msp.postgres.inner.v1.ClusterService",
		"nebius.msp",
		"nebius.ClusterService",
	}
	for _, id := range tests {
		t.Run(string(id), func(t *testing.T) {
			t.Parallel()

			address, err := conn.NewConventionResolver().Resolve(t.Context(), id)
			require.EqualError(t, err, "unknown service "+string(id))
			unknownError := &conn.UnknownServiceError{}
			require.ErrorAs(t, err, &unknownError)
			assert.Equal(t, id, unknownError.ID)
			assert.Empty(t, address)
		})
	}
}
