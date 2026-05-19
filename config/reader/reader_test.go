package reader_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/config"
	reader "github.com/nebius/gosdk/config/reader"
)

type metricsRecorder struct {
	configLoad         int
	credentialsResolve int
	tokenAcquire       int
	cacheMiss          int
}

type authOnlyMetricsRecorder struct {
	tokenAcquire int
}

func (m *metricsRecorder) TokenAcquire(context.Context, string, string, time.Duration, int) {
	m.tokenAcquire++
}

func (*metricsRecorder) TokenLifetime(context.Context, string, time.Duration) {}

func (*metricsRecorder) TokenRefresh(context.Context, string, string, time.Duration, bool) {}

func (*metricsRecorder) CacheHit(context.Context, string) {}

func (m *metricsRecorder) CacheMiss(context.Context, string, string) {
	m.cacheMiss++
}

func (*metricsRecorder) CacheStore(context.Context, string, string) {}

func (*metricsRecorder) CacheRefresh(context.Context, string, string) {}

func (*metricsRecorder) CacheInvalidate(context.Context, string) {}

func (m *metricsRecorder) ConfigLoad(context.Context, string, string, time.Duration) {
	m.configLoad++
}

func (m *metricsRecorder) CredentialsResolve(context.Context, string, string, time.Duration) {
	m.credentialsResolve++
}

func (m *authOnlyMetricsRecorder) TokenAcquire(context.Context, string, string, time.Duration, int) {
	m.tokenAcquire++
}

func (*authOnlyMetricsRecorder) TokenLifetime(context.Context, string, time.Duration) {}

func (*authOnlyMetricsRecorder) TokenRefresh(context.Context, string, string, time.Duration, bool) {}

func (*authOnlyMetricsRecorder) CacheHit(context.Context, string) {}

func (*authOnlyMetricsRecorder) CacheMiss(context.Context, string, string) {}

func (*authOnlyMetricsRecorder) CacheStore(context.Context, string, string) {}

func (*authOnlyMetricsRecorder) CacheRefresh(context.Context, string, string) {}

func (*authOnlyMetricsRecorder) CacheInvalidate(context.Context, string) {}

type testSlogHandler struct {
	id string
}

func (h *testSlogHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (h *testSlogHandler) Handle(context.Context, slog.Record) error {
	return nil
}

func (h *testSlogHandler) WithAttrs([]slog.Attr) slog.Handler {
	return &testSlogHandler{id: h.id}
}

func (h *testSlogHandler) WithGroup(string) slog.Handler {
	return &testSlogHandler{id: h.id}
}

func valueByPath(t *testing.T, obj any, path ...string) reflect.Value {
	t.Helper()

	v := reflect.ValueOf(obj)
	for _, name := range path {
		for v.Kind() == reflect.Interface || v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		field := v.FieldByName(name)
		if !field.IsValid() {
			t.Fatalf("field %q not found", name)
		}
		v = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	}
	return v
}

func loggerPointer(t *testing.T, obj any, path ...string) uintptr {
	t.Helper()

	return valueByPath(t, obj, path...).Pointer()
}

func loggerHandlerID(t *testing.T, obj any, path ...string) string {
	t.Helper()

	logger, ok := valueByPath(t, obj, path...).Interface().(*slog.Logger)
	if !ok {
		t.Fatalf("field %v is not *slog.Logger", path)
	}
	handler, ok := logger.Handler().(*testSlogHandler)
	if !ok {
		t.Fatalf("logger handler at %v is %T, want *testSlogHandler", path, logger.Handler())
	}
	return handler.id
}

func TestGetCredentialsAuthOptionsLoggerOverridesReaderLogger(t *testing.T) {
	t.Parallel()

	overrideReaderLogger := slog.New(&testSlogHandler{id: "reader"})
	overrideAuthLogger := slog.New(&testSlogHandler{id: "auth"})

	testCases := []struct {
		name          string
		readerLogger  *slog.Logger
		authOptions   []auth.Option
		wantHandlerID string
	}{
		{
			name:          "reader logger is used by default",
			readerLogger:  slog.New(&testSlogHandler{id: "default"}),
			wantHandlerID: "default",
		},
		{
			name:         "auth options logger overrides reader logger",
			readerLogger: overrideReaderLogger,
			authOptions: []auth.Option{
				auth.WithLogger(overrideAuthLogger),
			},
			wantHandlerID: "auth",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cfg := &config.Config{
				Default: "test",
				Profiles: config.ProfilesConfig{
					"test": {
						AuthType:           config.AuthTypeFederation,
						FederationEndpoint: "https://federation.example.test",
						FederationID:       "fed-id",
					},
				},
			}

			cfgReader := reader.NewConfigReader(
				reader.WithPreloadedConfig(cfg),
				reader.WithClientID("client-id"),
				reader.WithLogger(tc.readerLogger),
				reader.WithAuthOptions(tc.authOptions...),
			)

			if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
				t.Fatalf("load config: %v", err)
			}

			tokener, err := cfgReader.GetCredentials(context.Background())
			if err != nil {
				t.Fatalf("get credentials: %v", err)
			}

			syncTokener := tokener.(*auth.InAppSyncTokener)
			cacheTokener := syncTokener.Unwrap().(*auth.RenewableFileCachedTokener)
			federationTokener := cacheTokener.Unwrap().(*auth.FederationTokener)

			if got := loggerHandlerID(t, cacheTokener, "logger"); got != tc.wantHandlerID {
				t.Fatalf(
					"cache logger handler id = %q, want %q",
					got,
					tc.wantHandlerID,
				)
			}

			if got := loggerHandlerID(t, federationTokener, "logger"); got != tc.wantHandlerID {
				t.Fatalf(
					"federation logger handler id = %q, want %q",
					got,
					tc.wantHandlerID,
				)
			}
		})
	}
}

func TestWithLoggerNilUsesNoopLogger(t *testing.T) {
	t.Parallel()

	cfg := &config.Config{
		Default: "test",
		Profiles: config.ProfilesConfig{
			"test": {
				AuthType:           config.AuthTypeFederation,
				FederationEndpoint: "https://federation.example.test",
				FederationID:       "fed-id",
			},
		},
	}

	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(cfg),
		reader.WithClientID("client-id"),
		reader.WithLogger(nil),
	)

	if loadErr := cfgReader.LoadIfNeeded(context.Background()); loadErr != nil {
		t.Fatalf("load config: %v", loadErr)
	}

	tokener, err := cfgReader.GetCredentials(context.Background())
	if err != nil {
		t.Fatalf("get credentials: %v", err)
	}

	syncTokener := tokener.(*auth.InAppSyncTokener)
	cacheTokener := syncTokener.Unwrap().(*auth.RenewableFileCachedTokener)
	federationTokener := cacheTokener.Unwrap().(*auth.FederationTokener)

	if got := loggerPointer(t, cacheTokener, "logger"); got == 0 {
		t.Fatal("cache logger is nil")
	}
	if got := loggerPointer(t, federationTokener, "logger"); got == 0 {
		t.Fatal("federation logger is nil")
	}
}

func TestGetCredentialsTokenFilePropagatesMetrics(t *testing.T) {
	t.Parallel()

	tokenFile := t.TempDir() + "/token"
	if err := os.WriteFile(tokenFile, []byte("test-token"), 0600); err != nil {
		t.Fatalf("write token file: %v", err)
	}

	cfg := &config.Config{
		Default: "test",
		Profiles: config.ProfilesConfig{
			"test": {
				TokenFile: tokenFile,
			},
		},
	}
	metrics := &metricsRecorder{}

	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(cfg),
		reader.WithMetrics(metrics),
	)

	if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
		t.Fatalf("load config: %v", err)
	}

	tokener, err := cfgReader.GetCredentials(context.Background())
	if err != nil {
		t.Fatalf("get credentials: %v", err)
	}

	if _, tokenErr := tokener.BearerToken(context.Background()); tokenErr != nil {
		t.Fatalf("read token: %v", tokenErr)
	}

	if metrics.configLoad != 1 {
		t.Fatalf("config load count = %d, want 1", metrics.configLoad)
	}
	if metrics.credentialsResolve != 1 {
		t.Fatalf("credentials resolve count = %d, want 1", metrics.credentialsResolve)
	}
	if metrics.cacheMiss != 1 {
		t.Fatalf("cache miss count = %d, want 1", metrics.cacheMiss)
	}
	if metrics.tokenAcquire != 1 {
		t.Fatalf("token acquire count = %d, want 1", metrics.tokenAcquire)
	}
}

func TestLoadIfNeededDoesNotEmitMetricsWhenConfigAlreadyParsed(t *testing.T) {
	t.Parallel()

	metrics := &metricsRecorder{}
	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(&config.Config{
			Default: "test",
			Profiles: config.ProfilesConfig{
				"test": {
					Endpoint: "api.example.test:443",
				},
			},
		}),
		reader.WithMetrics(metrics),
	)

	if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
		t.Fatalf("first load: %v", err)
	}
	if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
		t.Fatalf("second load: %v", err)
	}

	if metrics.configLoad != 1 {
		t.Fatalf("config load count = %d, want 1", metrics.configLoad)
	}
}

func TestExpandAuthMetricsMutesReaderCallbacks(t *testing.T) {
	t.Parallel()

	authMetrics := &authOnlyMetricsRecorder{}
	expanded := reader.ExpandAuthMetrics(authMetrics)

	expanded.ConfigLoad(context.Background(), "preloaded", "success", time.Second)
	expanded.CredentialsResolve(context.Background(), "env", "success", time.Second)
	expanded.TokenAcquire(context.Background(), "static", "success", time.Second, 1)

	if authMetrics.tokenAcquire != 1 {
		t.Fatalf("token acquire count = %d, want 1", authMetrics.tokenAcquire)
	}
}

func TestGetCredentialsTokenFileDefersFileAccessToTokener(t *testing.T) {
	t.Parallel()

	home, err := os.UserHomeDir()
	if err != nil {
		t.Skipf("cannot resolve home directory: %v", err)
	}
	missingTokenFile := ".nebius-gosdk-test-missing-token-" + filepath.Base(t.TempDir())
	expandedTokenFile := filepath.Join(home, missingTokenFile)

	cfg := &config.Config{
		Default: "test",
		Profiles: config.ProfilesConfig{
			"test": {
				TokenFile: "~/" + missingTokenFile,
			},
		},
	}

	cfgReader := reader.NewConfigReader(
		reader.WithPreloadedConfig(cfg),
	)

	if loadErr := cfgReader.LoadIfNeeded(context.Background()); loadErr != nil {
		t.Fatalf("load config: %v", loadErr)
	}

	tokener, err := cfgReader.GetCredentials(context.Background())
	if err != nil {
		t.Fatalf("get credentials: %v", err)
	}

	if _, err = tokener.BearerToken(context.Background()); err == nil {
		t.Fatal("expected token-file read error")
	} else if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("read token error = %v, want os.ErrNotExist", err)
	} else if !strings.Contains(err.Error(), expandedTokenFile) {
		t.Fatalf("error = %q, want expanded token path %q", err.Error(), expandedTokenFile)
	}
}
