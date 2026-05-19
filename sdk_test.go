package gosdk_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/config"
	reader "github.com/nebius/gosdk/config/reader"
)

type sdkMetricsRecorder struct {
	configLoad         int
	credentialsResolve int
	tokenAcquire       int
	cacheMiss          int
}

type sdkAuthMetricsRecorder struct {
	tokenAcquire int
	cacheMiss    int
}

func (m *sdkMetricsRecorder) TokenAcquire(context.Context, string, string, time.Duration, int) {
	m.tokenAcquire++
}

func (*sdkMetricsRecorder) TokenLifetime(context.Context, string, time.Duration) {}

func (*sdkMetricsRecorder) TokenRefresh(context.Context, string, string, time.Duration, bool) {}

func (*sdkMetricsRecorder) CacheHit(context.Context, string) {}

func (m *sdkMetricsRecorder) CacheMiss(context.Context, string, string) {
	m.cacheMiss++
}

func (*sdkMetricsRecorder) CacheStore(context.Context, string, string) {}

func (*sdkMetricsRecorder) CacheRefresh(context.Context, string, string) {}

func (*sdkMetricsRecorder) CacheInvalidate(context.Context, string) {}

func (m *sdkMetricsRecorder) ConfigLoad(context.Context, string, string, time.Duration) {
	m.configLoad++
}

func (m *sdkMetricsRecorder) CredentialsResolve(context.Context, string, string, time.Duration) {
	m.credentialsResolve++
}

func (m *sdkAuthMetricsRecorder) TokenAcquire(context.Context, string, string, time.Duration, int) {
	m.tokenAcquire++
}

func (*sdkAuthMetricsRecorder) TokenLifetime(context.Context, string, time.Duration) {}

func (*sdkAuthMetricsRecorder) TokenRefresh(context.Context, string, string, time.Duration, bool) {}

func (*sdkAuthMetricsRecorder) CacheHit(context.Context, string) {}

func (m *sdkAuthMetricsRecorder) CacheMiss(context.Context, string, string) {
	m.cacheMiss++
}

func (*sdkAuthMetricsRecorder) CacheStore(context.Context, string, string) {}

func (*sdkAuthMetricsRecorder) CacheRefresh(context.Context, string, string) {}

func (*sdkAuthMetricsRecorder) CacheInvalidate(context.Context, string) {}

func TestConfigReaderMetricsPreservedWhenSDKMetricsUnset(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("test-token"), 0600))

	readerMetrics := &sdkMetricsRecorder{}
	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(&config.Config{
			Default: "test",
			Profiles: config.ProfilesConfig{
				"test": {
					Endpoint:  "api.example.test:443",
					TokenFile: tokenPath,
				},
			},
		}),
		reader.WithMetrics(readerMetrics),
	)

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithConfigReader(cfgReader),
		gosdk.WithExplicitInit(true),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, sdk.Close())
	})

	_, tokenErr := sdk.BearerToken(ctx)
	require.NoError(t, tokenErr)

	assert.Equal(t, 1, readerMetrics.configLoad)
	assert.Equal(t, 1, readerMetrics.credentialsResolve)
	assert.Equal(t, 1, readerMetrics.cacheMiss)
	assert.Equal(t, 1, readerMetrics.tokenAcquire)
}

func TestWithMetricsOverridesConfigReaderMetrics(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("test-token"), 0600))

	readerMetrics := &sdkMetricsRecorder{}
	sdkMetrics := &sdkMetricsRecorder{}
	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(&config.Config{
			Default: "test",
			Profiles: config.ProfilesConfig{
				"test": {
					Endpoint:  "api.example.test:443",
					TokenFile: tokenPath,
				},
			},
		}),
		reader.WithMetrics(readerMetrics),
	)

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithConfigReader(cfgReader),
		gosdk.WithMetrics(sdkMetrics),
		gosdk.WithExplicitInit(true),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, sdk.Close())
	})

	_, tokenErr := sdk.BearerToken(ctx)
	require.NoError(t, tokenErr)

	assert.Zero(t, readerMetrics.configLoad)
	assert.Zero(t, readerMetrics.credentialsResolve)
	assert.Zero(t, readerMetrics.cacheMiss)
	assert.Zero(t, readerMetrics.tokenAcquire)
	assert.Equal(t, 1, sdkMetrics.configLoad)
	assert.Equal(t, 1, sdkMetrics.credentialsResolve)
	assert.Equal(t, 1, sdkMetrics.cacheMiss)
	assert.Equal(t, 1, sdkMetrics.tokenAcquire)
}

func TestWithAuthMetricsUsesAuthOnlyMetrics(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tokenPath := filepath.Join(t.TempDir(), "token")
	require.NoError(t, os.WriteFile(tokenPath, []byte("test-token"), 0600))

	authMetrics := &sdkAuthMetricsRecorder{}
	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(&config.Config{
			Default: "test",
			Profiles: config.ProfilesConfig{
				"test": {
					Endpoint:  "api.example.test:443",
					TokenFile: tokenPath,
				},
			},
		}),
	)

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithConfigReader(cfgReader),
		gosdk.WithAuthMetrics(authMetrics),
		gosdk.WithExplicitInit(true),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, sdk.Close())
	})

	_, tokenErr := sdk.BearerToken(ctx)
	require.NoError(t, tokenErr)

	assert.Equal(t, 1, authMetrics.cacheMiss)
	assert.Equal(t, 1, authMetrics.tokenAcquire)
}
