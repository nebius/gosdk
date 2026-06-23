package conn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk/conn"
)

func TestParseOverridesOptions(t *testing.T) {
	t.Parallel()

	overrides, err := conn.ParseOverrides(
		"nebius.compute.*=compute.example,nebius.mk8s.*=localhost:30080;no-tls-verify,nebius.msp.*=msp.example;insecure",
	)
	require.NoError(t, err)
	require.Len(t, overrides, 3)

	assert.Equal(t, "nebius.compute.*", overrides[0].ServiceID)
	assert.Equal(t, "compute.example", overrides[0].Address)
	assert.False(t, overrides[0].Insecure)
	assert.False(t, overrides[0].NoTLSVerify)

	assert.Equal(t, "nebius.mk8s.*", overrides[1].ServiceID)
	assert.Equal(t, "localhost:30080", overrides[1].Address)
	assert.False(t, overrides[1].Insecure)
	assert.True(t, overrides[1].NoTLSVerify)

	assert.Equal(t, "nebius.msp.*", overrides[2].ServiceID)
	assert.Equal(t, "msp.example", overrides[2].Address)
	assert.True(t, overrides[2].Insecure)
	assert.False(t, overrides[2].NoTLSVerify)
}

func TestParseOverridesRejectsConflictingTransportOptions(t *testing.T) {
	t.Parallel()

	_, err := conn.ParseOverrides(
		"nebius.compute.*=compute.example," +
			"nebius.mk8s.*=localhost:30080;insecure;no-tls-verify," +
			"nebius.iam.*=iam.example;insecure",
	)
	require.EqualError(t, err,
		`options "insecure" and "no-tls-verify" are mutually exclusive in `+
			`"nebius.mk8s.*=localhost:30080;insecure;no-tls-verify"`,
	)
}
