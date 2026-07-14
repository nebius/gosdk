//nolint:testpackage // tests intentionally configure package-local IMDS probe knobs.
package reader

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nebius/gosdk/config"
	"github.com/nebius/gosdk/config/paths"
)

const metadataPath = "/v1/iam/sa/token"

func withMetadataProbe(probe metadataProbe) config.Option {
	return optionFunc(func(c *configReader) {
		c.metadataProbe = probe
	})
}

func newMetadataProbe(t *testing.T, statusCode int) (metadataProbe, *atomic.Bool) {
	t.Helper()

	var sawValidProbeRequest atomic.Bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != metadataPath {
			http.NotFound(w, r)
			return
		}
		if r.Header.Get("Metadata") != "true" {
			http.Error(w, "missing Metadata: true header", http.StatusBadRequest)
			return
		}
		sawValidProbeRequest.Store(true)
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(`{"access_token":"token","expires_at":"2030-01-01T00:00:00Z"}`))
	}))

	t.Cleanup(server.Close)

	return metadataProbe{
		tokenEndpoint: server.URL + metadataPath,
		tokenFilePath: filepath.Join(t.TempDir(), "missing-token"),
		httpClient:    server.Client(),
	}, &sawValidProbeRequest
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestLoadMissingConfigUsesVirtualProfileWhenMetadataAvailable(t *testing.T) {
	t.Parallel()

	probe, sawValidProbeRequest := newMetadataProbe(t, http.StatusOK)

	cfgReader := NewConfigReader(
		WithPath(filepath.Join(t.TempDir(), "missing.yaml")),
		withMetadataProbe(probe),
	)

	if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
		t.Fatalf("load config: %v", err)
	}
	if !sawValidProbeRequest.Load() {
		t.Fatal("metadata token probe request did not include Metadata: true header")
	}
	if got := cfgReader.CurrentProfileName(); got != virtualProfileName {
		t.Fatalf("current profile = %q, want %s", got, virtualProfileName)
	}
	if got := cfgReader.Endpoint(); got != paths.DefaultAPIEndpoint {
		t.Fatalf("endpoint = %q, want %q", got, paths.DefaultAPIEndpoint)
	}
	profile := cfgReader.GetConfig().Profiles[virtualProfileName]
	if profile == nil {
		t.Fatal("virtual profile is missing")
	}
	if got := profile.TokenEndpoint; got != probe.tokenEndpoint {
		t.Fatalf("token endpoint = %q, want %q", got, probe.tokenEndpoint)
	}
}

func TestLoadMissingConfigReturnsMissingConfigWhenMetadataUnavailable(t *testing.T) {
	t.Parallel()

	probe, sawValidProbeRequest := newMetadataProbe(t, http.StatusNotFound)

	cfgReader := NewConfigReader(
		WithPath(filepath.Join(t.TempDir(), "missing.yaml")),
		withMetadataProbe(probe),
	)

	err := cfgReader.LoadIfNeeded(context.Background())
	if !sawValidProbeRequest.Load() {
		t.Fatal("metadata token probe request did not include Metadata: true header")
	}
	var missingConfigErr *config.MissingConfigError
	if !errors.As(err, &missingConfigErr) {
		t.Fatalf("load config error = %v, want MissingConfigError", err)
	}
	if got := cfgReader.CurrentProfileName(); got != "" {
		t.Fatalf("current profile = %q, want empty", got)
	}
}

func TestLoadMissingConfigRetriesRetriableMetadataStatus(t *testing.T) {
	t.Parallel()

	const requestTimeout = 10 * time.Millisecond

	var attempts atomic.Int32
	client := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if _, ok := r.Context().Deadline(); !ok {
				return nil, errors.New("metadata probe request has no timeout")
			}

			statusCode := http.StatusOK
			switch {
			case r.URL.Path != metadataPath:
				statusCode = http.StatusNotFound
			case r.Header.Get("Metadata") != "true":
				statusCode = http.StatusBadRequest
			case attempts.Add(1) == 1:
				statusCode = http.StatusServiceUnavailable
			}

			return &http.Response{
				StatusCode: statusCode,
				Body:       http.NoBody,
			}, nil
		}),
	}

	probe := metadataProbe{
		tokenEndpoint: "http://metadata.nebius.internal" + metadataPath,
		tokenFilePath: filepath.Join(t.TempDir(), "missing-token"),
		httpClient:    client,
		timeout:       requestTimeout,
		maxAttempts:   2,
		baseBackoff:   requestTimeout * 2,
	}
	cfgReader := NewConfigReader(
		WithPath(filepath.Join(t.TempDir(), "missing.yaml")),
		withMetadataProbe(probe),
	)

	if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
		t.Fatalf("load config: %v", err)
	}
	if got := attempts.Load(); got != 2 {
		t.Fatalf("metadata token probe attempts = %d, want 2", got)
	}
	profile := cfgReader.GetConfig().Profiles[virtualProfileName]
	if profile == nil {
		t.Fatal("virtual profile is missing")
	}
	if got := profile.TokenEndpoint; got != probe.tokenEndpoint {
		t.Fatalf("token endpoint = %q, want %q", got, probe.tokenEndpoint)
	}
}

func TestLoadMissingConfigUsesVirtualProfileTokenFileWhenMetadataUnavailable(t *testing.T) {
	t.Parallel()

	probe, sawValidProbeRequest := newMetadataProbe(t, http.StatusNotFound)
	tokenFilePath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenFilePath, []byte("token"), 0o600); err != nil {
		t.Fatalf("write token file: %v", err)
	}
	probe.tokenFilePath = tokenFilePath

	cfgReader := NewConfigReader(
		WithPath(filepath.Join(t.TempDir(), "missing.yaml")),
		withMetadataProbe(probe),
	)

	if err := cfgReader.LoadIfNeeded(context.Background()); err != nil {
		t.Fatalf("load config: %v", err)
	}
	if !sawValidProbeRequest.Load() {
		t.Fatal("metadata token probe request did not include Metadata: true header")
	}
	if got := cfgReader.CurrentProfileName(); got != virtualProfileName {
		t.Fatalf("current profile = %q, want %s", got, virtualProfileName)
	}
	profile := cfgReader.GetConfig().Profiles[virtualProfileName]
	if profile == nil {
		t.Fatal("virtual profile is missing")
	}
	if got := profile.TokenFile; got != tokenFilePath {
		t.Fatalf("token file = %q, want %q", got, tokenFilePath)
	}
	if got := profile.TokenEndpoint; got != "" {
		t.Fatalf("token endpoint = %q, want empty", got)
	}
}
