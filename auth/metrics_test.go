package auth_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/nebius/gosdk/auth"
	iampb "github.com/nebius/gosdk/proto/nebius/iam/v1"
)

type lifetimeMetricsRecorder struct {
	acquireCount   int
	lifetimeCount  int
	refreshCount   int
	cacheMissCount int
	ttl            time.Duration
	providers      []string
	results        []string
	missResults    []string
}

func (m *lifetimeMetricsRecorder) TokenAcquire(
	_ context.Context,
	provider string,
	result string,
	_ time.Duration,
	_ int,
) {
	m.acquireCount++
	m.providers = append(m.providers, provider)
	m.results = append(m.results, result)
}

func (m *lifetimeMetricsRecorder) TokenLifetime(_ context.Context, _ string, ttl time.Duration) {
	m.lifetimeCount++
	m.ttl = ttl
}

func (m *lifetimeMetricsRecorder) TokenRefresh(context.Context, string, string, time.Duration, bool) {
	m.refreshCount++
}

func (*lifetimeMetricsRecorder) CacheHit(context.Context, string) {}

func (m *lifetimeMetricsRecorder) CacheMiss(_ context.Context, _ string, result string) {
	m.cacheMissCount++
	m.missResults = append(m.missResults, result)
}

func (*lifetimeMetricsRecorder) CacheStore(context.Context, string, string) {}

func (*lifetimeMetricsRecorder) CacheRefresh(context.Context, string, string) {}

func (*lifetimeMetricsRecorder) CacheInvalidate(context.Context, string) {}

type testExchangeTokenRequester struct{}

func (testExchangeTokenRequester) GetExchangeTokenRequest(context.Context) (*iampb.ExchangeTokenRequest, error) {
	return &iampb.ExchangeTokenRequest{}, nil
}

type testTokenExchangeClient struct{}

func (testTokenExchangeClient) Exchange(
	context.Context,
	*iampb.ExchangeTokenRequest,
	...grpc.CallOption,
) (*iampb.CreateTokenResponse, error) {
	return &iampb.CreateTokenResponse{
		AccessToken: "test-access-token",
		TokenType:   "Bearer",
		ExpiresIn:   60,
	}, nil
}

func TestExchangeableBearerTokenerReportsFetchedTokenLifetime(t *testing.T) {
	t.Parallel()

	metrics := &lifetimeMetricsRecorder{}
	tokener := auth.NewExchangeableBearerTokener(testExchangeTokenRequester{}, testTokenExchangeClient{})
	tokener.SetMetrics(metrics)

	token, err := tokener.BearerToken(context.Background())
	if err != nil {
		t.Fatalf("bearer token: %v", err)
	}
	if token.Token == "" {
		t.Fatal("token is empty")
	}
	if metrics.acquireCount != 1 {
		t.Fatalf("token acquire count = %d, want 1", metrics.acquireCount)
	}
	if got := metrics.providers[0]; got != "token-exchange" {
		t.Fatalf("provider = %q, want %q", got, "token-exchange")
	}
	if metrics.lifetimeCount != 1 {
		t.Fatalf("token lifetime count = %d, want 1", metrics.lifetimeCount)
	}
	if metrics.ttl < 55*time.Second || metrics.ttl > 60*time.Second {
		t.Fatalf("token lifetime = %s, want between 55s and 60s", metrics.ttl)
	}
}

func TestTypedNameWrapperPropagatesMetricsAndProviderType(t *testing.T) {
	t.Parallel()

	metrics := &lifetimeMetricsRecorder{}
	wrapped := auth.NewTypedNameWrapper(
		"service-account/test/key",
		"service-account",
		auth.NewExchangeableBearerTokener(testExchangeTokenRequester{}, testTokenExchangeClient{}),
	)
	wrapped.SetMetrics(metrics)

	token, err := wrapped.BearerToken(context.Background())
	if err != nil {
		t.Fatalf("bearer token: %v", err)
	}
	if token.Token == "" {
		t.Fatal("token is empty")
	}
	if metrics.acquireCount != 1 {
		t.Fatalf("token acquire count = %d, want 1", metrics.acquireCount)
	}
	if got, ok := auth.TypeOfTokener(wrapped); !ok || got != "service-account" {
		t.Fatalf("wrapped type = %q, %t, want %q, true", got, ok, "service-account")
	}
	if got := metrics.providers[0]; got != "token-exchange" {
		t.Fatalf("provider = %q, want %q", got, "token-exchange")
	}
	if metrics.lifetimeCount != 1 {
		t.Fatalf("token lifetime count = %d, want 1", metrics.lifetimeCount)
	}
}

func TestAsynchronouslyRenewableFileCachedTokenerForegroundAcquireDoesNotReportRefresh(t *testing.T) {
	t.Parallel()

	metrics := &lifetimeMetricsRecorder{}
	tokener := auth.NewAsynchronouslyRenewableFileCacheTokener(
		auth.NewTypedNameWrapper(
			"service-account/test/key",
			"service-account",
			auth.NewExchangeableBearerTokener(testExchangeTokenRequester{}, testTokenExchangeClient{}),
		),
		auth.WithFileCacheFileName(filepath.Join(t.TempDir(), "credentials.yaml")),
	)
	tokener.SetMetrics(metrics)

	token, err := tokener.BearerToken(context.Background())
	if err != nil {
		t.Fatalf("bearer token: %v", err)
	}
	if token.Token == "" {
		t.Fatal("token is empty")
	}
	if metrics.acquireCount != 1 {
		t.Fatalf("token acquire count = %d, want 1", metrics.acquireCount)
	}
	if metrics.refreshCount != 0 {
		t.Fatalf("token refresh count = %d, want 0", metrics.refreshCount)
	}
	if metrics.cacheMissCount != 1 {
		t.Fatalf("cache miss count = %d, want 1", metrics.cacheMissCount)
	}
	if got := metrics.missResults[0]; got != "success" {
		t.Fatalf("cache miss result = %q, want %q", got, "success")
	}
}

func TestFileTokenerReportsFailedMissAsError(t *testing.T) {
	t.Parallel()

	metrics := &lifetimeMetricsRecorder{}
	tokener, err := auth.NewFileTokener(
		filepath.Join(t.TempDir(), "missing-token"),
		0,
		auth.WithMetrics(metrics),
	)
	if err != nil {
		t.Fatalf("new file tokener: %v", err)
	}

	_, err = tokener.BearerToken(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	if metrics.cacheMissCount != 1 {
		t.Fatalf("cache miss count = %d, want 1", metrics.cacheMissCount)
	}
	if got := metrics.missResults[0]; got != "error" {
		t.Fatalf("cache miss result = %q, want %q", got, "error")
	}
}

func TestStaticTokenerReportsAcquireMetrics(t *testing.T) {
	t.Parallel()

	metrics := &lifetimeMetricsRecorder{}
	tokener := auth.NewStaticTokener("test-token", auth.WithMetrics(metrics))

	token, err := tokener.BearerToken(context.Background())
	if err != nil {
		t.Fatalf("bearer token: %v", err)
	}
	if token.Token != "test-token" {
		t.Fatalf("token = %q, want %q", token.Token, "test-token")
	}
	if metrics.acquireCount != 1 {
		t.Fatalf("token acquire count = %d, want 1", metrics.acquireCount)
	}
	if got := metrics.providers[0]; got != "static" {
		t.Fatalf("provider = %q, want %q", got, "static")
	}
	if got := metrics.results[0]; got != "success" {
		t.Fatalf("result = %q, want %q", got, "success")
	}
}

func TestInstrumentedBearerTokenerPropagatesMetricsToWrappedTokener(t *testing.T) {
	t.Parallel()

	tokenPath := filepath.Join(t.TempDir(), "token")
	if err := os.WriteFile(tokenPath, []byte("test-token"), 0600); err != nil {
		t.Fatalf("write token file: %v", err)
	}

	metrics := &lifetimeMetricsRecorder{}
	tokener, err := auth.NewFileTokener(tokenPath, 0)
	if err != nil {
		t.Fatalf("new file tokener: %v", err)
	}
	wrapped := auth.NewInstrumentedBearerTokener(tokener)
	wrapped.SetMetrics(metrics)

	if _, tokenErr := wrapped.BearerToken(context.Background()); tokenErr != nil {
		t.Fatalf("bearer token: %v", tokenErr)
	}
	if metrics.cacheMissCount != 1 {
		t.Fatalf("cache miss count = %d, want 1", metrics.cacheMissCount)
	}
	if got := metrics.missResults[0]; got != "success" {
		t.Fatalf("cache miss result = %q, want %q", got, "success")
	}
}

type trackingReadCloser struct {
	reader io.Reader
	closed bool
}

func (t *trackingReadCloser) Read(p []byte) (int, error) {
	return t.reader.Read(p)
}

func (t *trackingReadCloser) Close() error {
	t.closed = true
	return nil
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestIMDSTokenizerClosesRetryableResponseBodies(t *testing.T) {
	t.Parallel()

	retryBody := &trackingReadCloser{reader: strings.NewReader("retry later")}
	attempts := 0
	client := &http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			attempts++
			switch attempts {
			case 1:
				return &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Status:     "429 Too Many Requests",
					Body:       retryBody,
				}, nil
			default:
				return &http.Response{
					StatusCode: http.StatusOK,
					Status:     "200 OK",
					Body: io.NopCloser(strings.NewReader(
						`{"access_token":"token","expires_at":"2026-01-02T15:04:05Z"}`,
					)),
				}, nil
			}
		}),
	}
	tokener, err := auth.NewIMDSTokenizer(
		"http://metadata.internal/token",
		auth.WithHTTPClient(client),
		auth.WithHTTPBaseBackoff(0),
		auth.WithHTTPMaxAttempts(2),
	)
	if err != nil {
		t.Fatalf("new imds tokener: %v", err)
	}

	if _, tokenErr := tokener.BearerToken(context.Background()); tokenErr != nil {
		t.Fatalf("bearer token: %v", tokenErr)
	}
	if !retryBody.closed {
		t.Fatal("retryable response body was not closed")
	}
}
