package gosdk

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

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
	if err := os.WriteFile(tokenPath, []byte("test-token"), 0600); err != nil {
		t.Fatalf("write token file: %v", err)
	}

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

	sdk, err := New(
		ctx,
		WithConfigReader(cfgReader),
		WithExplicitInit(true),
	)
	if err != nil {
		t.Fatalf("create sdk: %v", err)
	}
	t.Cleanup(func() {
		if err := sdk.Close(); err != nil {
			t.Fatalf("close sdk: %v", err)
		}
	})

	if _, err := sdk.BearerToken(ctx); err != nil {
		t.Fatalf("read token: %v", err)
	}

	if readerMetrics.configLoad != 1 {
		t.Fatalf("config load count = %d, want 1", readerMetrics.configLoad)
	}
	if readerMetrics.credentialsResolve != 1 {
		t.Fatalf("credentials resolve count = %d, want 1", readerMetrics.credentialsResolve)
	}
	if readerMetrics.cacheMiss != 1 {
		t.Fatalf("cache miss count = %d, want 1", readerMetrics.cacheMiss)
	}
	if readerMetrics.tokenAcquire != 1 {
		t.Fatalf("token acquire count = %d, want 1", readerMetrics.tokenAcquire)
	}
}

func TestWithMetricsOverridesConfigReaderMetrics(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tokenPath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenPath, []byte("test-token"), 0600); err != nil {
		t.Fatalf("write token file: %v", err)
	}

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

	sdk, err := New(
		ctx,
		WithConfigReader(cfgReader),
		WithMetrics(sdkMetrics),
		WithExplicitInit(true),
	)
	if err != nil {
		t.Fatalf("create sdk: %v", err)
	}
	t.Cleanup(func() {
		if err := sdk.Close(); err != nil {
			t.Fatalf("close sdk: %v", err)
		}
	})

	if _, err := sdk.BearerToken(ctx); err != nil {
		t.Fatalf("read token: %v", err)
	}

	if readerMetrics.configLoad != 0 || readerMetrics.credentialsResolve != 0 || readerMetrics.cacheMiss != 0 || readerMetrics.tokenAcquire != 0 {
		t.Fatalf("reader metrics were updated unexpectedly: %+v", *readerMetrics)
	}
	if sdkMetrics.configLoad != 1 {
		t.Fatalf("sdk config load count = %d, want 1", sdkMetrics.configLoad)
	}
	if sdkMetrics.credentialsResolve != 1 {
		t.Fatalf("sdk credentials resolve count = %d, want 1", sdkMetrics.credentialsResolve)
	}
	if sdkMetrics.cacheMiss != 1 {
		t.Fatalf("sdk cache miss count = %d, want 1", sdkMetrics.cacheMiss)
	}
	if sdkMetrics.tokenAcquire != 1 {
		t.Fatalf("sdk token acquire count = %d, want 1", sdkMetrics.tokenAcquire)
	}
}

func TestWithAuthMetricsUsesAuthOnlyMetrics(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tokenPath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenPath, []byte("test-token"), 0600); err != nil {
		t.Fatalf("write token file: %v", err)
	}

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

	sdk, err := New(
		ctx,
		WithConfigReader(cfgReader),
		WithAuthMetrics(authMetrics),
		WithExplicitInit(true),
	)
	if err != nil {
		t.Fatalf("create sdk: %v", err)
	}
	t.Cleanup(func() {
		if err := sdk.Close(); err != nil {
			t.Fatalf("close sdk: %v", err)
		}
	})

	if _, err := sdk.BearerToken(ctx); err != nil {
		t.Fatalf("read token: %v", err)
	}

	if authMetrics.cacheMiss != 1 {
		t.Fatalf("auth metrics cache miss count = %d, want 1", authMetrics.cacheMiss)
	}
	if authMetrics.tokenAcquire != 1 {
		t.Fatalf("auth metrics token acquire count = %d, want 1", authMetrics.tokenAcquire)
	}
}
