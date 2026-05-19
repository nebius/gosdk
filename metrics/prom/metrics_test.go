package prom_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"

	prommetrics "github.com/nebius/gosdk/metrics/prom"
)

func TestNewFromRegistry_DefaultJavaParityMetrics(t *testing.T) {
	t.Parallel()

	registry := prometheus.NewRegistry()
	metrics, err := prommetrics.NewFromRegistry(registry)
	if err != nil {
		t.Fatalf("NewFromRegistry() error = %v", err)
	}

	metrics.TokenRefresh(context.Background(), "federation", "success", time.Second, false)
	metrics.TokenRefresh(context.Background(), "federation", "error", time.Second, true)
	metrics.TokenLifetime(context.Background(), "federation", 45*time.Second)

	expected := `
# HELP ` + defaultTokenLifetimeMetricName() + ` Lifetime of the newly fetched token in seconds.
# TYPE ` + defaultTokenLifetimeMetricName() + ` gauge
` + defaultTokenLifetimeMetricName() + `{name="",provider="federation"} 45
# HELP ` + defaultTokenRefreshMetricName() + ` Total number of token refresh attempts.
# TYPE ` + defaultTokenRefreshMetricName() + ` counter
` + defaultTokenRefreshMetricName() + `{name="",provider="federation",success="false"} 1
` + defaultTokenRefreshMetricName() + `{name="",provider="federation",success="true"} 1
`
	if compareErr := testutil.GatherAndCompare(
		registry,
		strings.NewReader(expected),
		defaultTokenLifetimeMetricName(),
		defaultTokenRefreshMetricName(),
	); compareErr != nil {
		t.Fatalf("GatherAndCompare() error = %v", compareErr)
	}
}

func TestNewFromRegistry_PrefixAndMetricNameOverride(t *testing.T) {
	t.Parallel()

	registry := prometheus.NewRegistry()
	metrics, err := prommetrics.NewFromRegistry(
		registry,
		prommetrics.WithPrefix("sdk"),
		prommetrics.WithNames(prommetrics.Names{
			CacheHit: "cache_hits_total",
		}),
	)
	if err != nil {
		t.Fatalf("NewFromRegistry() error = %v", err)
	}

	metrics.CacheHit(context.Background(), "file_cache")

	expected := `
# HELP sdk_cache_hits_total Total number of cache hits.
# TYPE sdk_cache_hits_total counter
sdk_cache_hits_total{name="",provider="file_cache"} 1
`
	if compareErr := testutil.GatherAndCompare(
		registry,
		strings.NewReader(expected),
		"sdk_cache_hits_total",
	); compareErr != nil {
		t.Fatalf("GatherAndCompare() error = %v", compareErr)
	}
}

func TestMetricsWithNameSharesCollectors(t *testing.T) {
	t.Parallel()

	registry := prometheus.NewRegistry()
	metrics, err := prommetrics.NewFromRegistry(registry)
	if err != nil {
		t.Fatalf("NewFromRegistry() error = %v", err)
	}

	metrics.CacheHit(context.Background(), "file_cache")
	metrics.WithName("sdk-a").CacheHit(context.Background(), "file_cache")
	metrics.WithName("sdk-b").CacheHit(context.Background(), "file_cache")

	expected := `
# HELP ` + defaultCacheHitMetricName() + ` Total number of cache hits.
# TYPE ` + defaultCacheHitMetricName() + ` counter
` + defaultCacheHitMetricName() + `{name="",provider="file_cache"} 1
` + defaultCacheHitMetricName() + `{name="sdk-a",provider="file_cache"} 1
` + defaultCacheHitMetricName() + `{name="sdk-b",provider="file_cache"} 1
`
	if compareErr := testutil.GatherAndCompare(
		registry,
		strings.NewReader(expected),
		defaultCacheHitMetricName(),
	); compareErr != nil {
		t.Fatalf("GatherAndCompare() error = %v", compareErr)
	}
}

func TestNewFromRegistry_RejectsDuplicateCollectors(t *testing.T) {
	t.Parallel()

	registry := prometheus.NewRegistry()
	_, err := prommetrics.NewFromRegistry(registry)
	if err != nil {
		t.Fatalf("NewFromRegistry() error = %v", err)
	}
	if _, err = prommetrics.NewFromRegistry(registry); err == nil {
		t.Fatal("expected duplicate collector registration error")
	}
}

func TestNewAuthFromRegistry_DoesNotRegisterReaderMetrics(t *testing.T) {
	t.Parallel()

	registry := prometheus.NewRegistry()
	metrics, err := prommetrics.NewAuthFromRegistry(registry)
	if err != nil {
		t.Fatalf("NewAuthFromRegistry() error = %v", err)
	}

	metrics.CacheHit(context.Background(), "cache")

	gathered, err := registry.Gather()
	if err != nil {
		t.Fatalf("Gather() error = %v", err)
	}
	for _, metricFamily := range gathered {
		switch metricFamily.GetName() {
		case "config_load_seconds", "credentials_resolve_seconds":
			t.Fatalf("unexpected reader metric registered: %s", metricFamily.GetName())
		}
	}
}

func defaultTokenRefreshMetricName() string {
	return "gosdk_auth_token_refresh_total"
}

func defaultTokenLifetimeMetricName() string {
	return "gosdk_auth_token_lifetime_seconds"
}

func defaultCacheHitMetricName() string {
	return "gosdk_auth_cache_hit_total"
}
