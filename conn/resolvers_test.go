package conn_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk/conn"
)

//nolint:paralleltest,tparallel,reassign // tests operate with the global state
func TestConventionResolver_positive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		id      conn.ServiceID
		address conn.Address
		id2name map[conn.ServiceID]string
	}{
		{
			id:      "nebius.compute.InstanceService",
			address: "compute.{domain}",
			id2name: map[conn.ServiceID]string{},
		},
		{
			id:      "nebius.msp.postgres.inner.v1.ClusterService",
			address: "msp.{domain}",
			id2name: map[conn.ServiceID]string{},
		},
		{
			id:      "nebius.msp.postgres.inner.v1.ClusterService",
			address: "postgres.msp.{domain}",
			id2name: map[conn.ServiceID]string{
				"nebius.msp.postgres.inner.v1.ClusterService": "postgres.msp",
			},
		},
	}
	for _, test := range tests {
		t.Run(string(test.id), func(t *testing.T) {
			conn.ConventionResolverServiceIDToNameMap = test.id2name
			address, err := conn.NewConventionResolver().Resolve(context.Background(), test.id)
			require.NoError(t, err)
			assert.Equal(t, test.address, address)
		})
	}
}

func TestConventionResolver_negative(t *testing.T) {
	t.Parallel()
	tests := []struct {
		id conn.ServiceID
	}{
		{id: "not.nebius.msp.postgres.inner.v1.ClusterService"},
		{id: "nebius.msp.postgres.inner.v1.ClusterServiceNot"},
		{id: "nebius.msp"},
		{id: "nebius.ClusterService"},
	}
	for _, test := range tests {
		t.Run(string(test.id), func(t *testing.T) {
			t.Parallel()
			address, err := conn.NewConventionResolver().Resolve(context.Background(), test.id)
			require.EqualError(t, err, string("unknown service "+test.id))
			unknownError := &conn.UnknownServiceError{}
			require.ErrorAs(t, err, &unknownError)
			assert.Equal(t, test.id, unknownError.ID)
			assert.Empty(t, address)
		})
	}
}
