package reader_test

import (
	"context"
	"log/slog"
	"reflect"
	"testing"
	"unsafe"

	"github.com/nebius/gosdk/auth"
	"github.com/nebius/gosdk/config"
	reader "github.com/nebius/gosdk/config/reader"
)

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

	if got := loggerPointer(t, cacheTokener, "logger"); got == 0 {
		t.Fatal("cache logger is nil")
	}
	if got := loggerPointer(t, federationTokener, "logger"); got == 0 {
		t.Fatal("federation logger is nil")
	}
}
