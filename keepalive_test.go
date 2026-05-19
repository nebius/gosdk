package gosdk_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk"
)

type sdkKeepaliveEnvTestCase struct {
	env              map[string]string
	withoutKeepalive bool
	wantErrSub       string
}

func TestKeepalivePublicDefaults(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 20*time.Second, gosdk.DefaultKeepaliveTime)
	assert.Equal(t, 10*time.Second, gosdk.DefaultKeepaliveTimeout)
	assert.True(t, gosdk.DefaultKeepalivePermitWithoutStream)
	assert.Equal(t, "NEBIUS_GRPC_KEEPALIVE_TIME", gosdk.EnvGRPCKeepaliveTime)
	assert.Equal(t, "NEBIUS_GRPC_KEEPALIVE_TIMEOUT", gosdk.EnvGRPCKeepaliveTimeout)
	assert.Equal(t, "NEBIUS_GRPC_KEEPALIVE_PERMIT_WITHOUT_STREAM", gosdk.EnvGRPCKeepalivePermitWithoutStream)
}

func TestSDKKeepaliveEnv(t *testing.T) {
	for name, tt := range sdkKeepaliveEnvTestCases() {
		t.Run(name, func(t *testing.T) {
			clearKeepaliveEnv(t)
			for key, value := range tt.env {
				t.Setenv(key, value)
			}

			var opts []gosdk.Option
			if tt.withoutKeepalive {
				opts = append(opts, gosdk.WithoutKeepalive())
			}
			sdk, err := newExplicitSDK(opts...)
			if tt.wantErrSub != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.wantErrSub)
				if sdk != nil {
					require.NoError(t, sdk.Close())
				}
				return
			}

			require.NoError(t, err)
			require.NotNil(t, sdk)
			t.Cleanup(func() {
				require.NoError(t, sdk.Close())
			})
		})
	}
}

func sdkKeepaliveEnvTestCases() map[string]sdkKeepaliveEnvTestCase {
	return map[string]sdkKeepaliveEnvTestCase{
		"defaults": {},
		"override": {
			env: map[string]string{
				gosdk.EnvGRPCKeepaliveTime:                "45s",
				gosdk.EnvGRPCKeepaliveTimeout:             "12s",
				gosdk.EnvGRPCKeepalivePermitWithoutStream: "false",
			},
		},
		"disable": {
			env: map[string]string{
				gosdk.EnvGRPCKeepaliveTime:    "0",
				gosdk.EnvGRPCKeepaliveTimeout: "0",
			},
		},
		"explicit option ignores invalid env": {
			env: map[string]string{
				gosdk.EnvGRPCKeepaliveTimeout: "0",
			},
			withoutKeepalive: true,
		},
		"empty values are ignored": {
			env: map[string]string{
				gosdk.EnvGRPCKeepaliveTime:                "",
				gosdk.EnvGRPCKeepaliveTimeout:             "",
				gosdk.EnvGRPCKeepalivePermitWithoutStream: "",
			},
		},
		"negative timeout": {
			env: map[string]string{
				gosdk.EnvGRPCKeepaliveTimeout: "-1s",
			},
			wantErrSub: "must not be negative",
		},
		"zero timeout while enabled": {
			env: map[string]string{
				gosdk.EnvGRPCKeepaliveTimeout: "0",
			},
			wantErrSub: "must be positive",
		},
		"invalid permit without stream": {
			env: map[string]string{
				gosdk.EnvGRPCKeepalivePermitWithoutStream: "sometimes",
			},
			wantErrSub: gosdk.EnvGRPCKeepalivePermitWithoutStream,
		},
	}
}

func newExplicitSDK(opts ...gosdk.Option) (*gosdk.SDK, error) {
	baseOpts := []gosdk.Option{
		gosdk.WithExplicitInit(true),
		gosdk.WithDomain("api.example.test:443"),
		gosdk.WithCredentials(gosdk.IAMToken("test-token")),
	}
	baseOpts = append(baseOpts, opts...)
	return gosdk.New(context.Background(), baseOpts...)
}

func clearKeepaliveEnv(t *testing.T) {
	t.Helper()

	for _, key := range []string{
		gosdk.EnvGRPCKeepaliveTime,
		gosdk.EnvGRPCKeepaliveTimeout,
		gosdk.EnvGRPCKeepalivePermitWithoutStream,
	} {
		t.Setenv(key, "")
	}
}
