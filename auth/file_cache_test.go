package auth_test

import (
	"context"
	"log/slog"
	"path/filepath"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/nebius/gosdk/auth"
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

func loggerAtPath(t *testing.T, obj any, path ...string) *slog.Logger {
	t.Helper()

	logger, ok := valueByPath(t, obj, path...).Interface().(*slog.Logger)
	if !ok {
		t.Fatalf("field %v is not *slog.Logger", path)
	}
	return logger
}

func TestPureFileCachedTokenerSetLoggerPropagatesToCache(t *testing.T) {
	t.Parallel()

	tokener := auth.NewPureFileCachedTokener(
		"test-token",
		auth.WithFileCacheFileName(filepath.Join(t.TempDir(), "credentials.yaml")),
	)

	logger := slog.New(&testSlogHandler{id: "initial"})
	setter, ok := tokener.(auth.LoggerSetter)
	if !ok {
		t.Fatal("tokener does not implement LoggerSetter")
	}
	setter.SetLogger(logger)

	if got := loggerAtPath(t, tokener, "logger").Handler().(*testSlogHandler).id; got != "initial" {
		t.Fatalf("logger id = %q, want %q", got, "initial")
	}

	if _, err := tokener.BearerToken(context.Background()); err != nil {
		t.Fatalf("bearer token: %v", err)
	}

	updatedLogger := slog.New(&testSlogHandler{id: "updated"})
	setter.SetLogger(updatedLogger)

	if got := loggerAtPath(t, tokener, "logger").Handler().(*testSlogHandler).id; got != "updated" {
		t.Fatalf("logger id after update = %q, want %q", got, "updated")
	}
	if got := loggerAtPath(t, tokener, "cache", "cache", "logger").Handler().(*testSlogHandler).id; got != "updated" {
		t.Fatalf("cache logger id = %q, want %q", got, "updated")
	}
}

func TestRenewableFileCachedTokenerSetLoggerPropagatesToCache(t *testing.T) {
	t.Parallel()

	tokener := auth.NewFileCacheTokener(
		auth.NewNameWrapper("test-token", auth.StaticBearerToken("token")),
		auth.WithFileCacheFileName(filepath.Join(t.TempDir(), "credentials.yaml")),
	)

	initialLogger := slog.New(&testSlogHandler{id: "initial"})
	tokener.SetLogger(initialLogger)

	if got := loggerAtPath(t, tokener, "logger").Handler().(*testSlogHandler).id; got != "initial" {
		t.Fatalf("logger id = %q, want %q", got, "initial")
	}

	if _, err := tokener.BearerToken(context.Background()); err != nil {
		t.Fatalf("bearer token: %v", err)
	}

	updatedLogger := slog.New(&testSlogHandler{id: "updated"})
	tokener.SetLogger(updatedLogger)

	if got := loggerAtPath(t, tokener, "logger").Handler().(*testSlogHandler).id; got != "updated" {
		t.Fatalf("logger id after update = %q, want %q", got, "updated")
	}
	if got := loggerAtPath(t, tokener, "cache", "cache", "logger").Handler().(*testSlogHandler).id; got != "updated" {
		t.Fatalf("cache logger id = %q, want %q", got, "updated")
	}
}

func TestWithFileCacheSafetyMarginUpdatesRenewableTokener(t *testing.T) {
	t.Parallel()

	tokener := auth.NewFileCacheTokener(
		auth.NewNameWrapper("test-token", auth.StaticBearerToken("token")),
		auth.WithFileCacheFileName(filepath.Join(t.TempDir(), "credentials.yaml")),
	)

	if got := valueByPath(t, tokener, "safetyMargin").Interface().(time.Duration); got != 0 {
		t.Fatalf("initial safety margin = %v, want %v", got, 0)
	}

	tokener = auth.NewFileCacheTokener(
		auth.NewNameWrapper("test-token", auth.StaticBearerToken("token")),
		auth.WithFileCacheFileName(filepath.Join(t.TempDir(), "credentials.yaml")),
		auth.WithFileCacheSafetyMargin(2*time.Minute),
	)

	if got := valueByPath(t, tokener, "safetyMargin").Interface().(time.Duration); got != 2*time.Minute {
		t.Fatalf("safety margin = %v, want %v", got, 2*time.Minute)
	}
	if loggerAtPath(t, tokener, "logger") == nil {
		t.Fatal("logger is nil after updating safety margin")
	}
}
