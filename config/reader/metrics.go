package reader

import (
	"context"
	"time"

	"github.com/nebius/gosdk/auth"
)

const (
	metricResultSuccess = "success"
	metricResultError   = "error"
)

// Metrics receives config-reader observations and auth observations emitted by nested tokeners.
type Metrics interface {
	auth.Metrics
	ConfigLoad(context.Context, string, string, time.Duration)
	CredentialsResolve(context.Context, string, string, time.Duration)
}

// ExpandedAuthMetrics adapts auth-only metrics to [Metrics] while muting
// config-reader-specific callbacks.
type ExpandedAuthMetrics struct {
	AuthMetrics auth.Metrics
}

var _ Metrics = ExpandedAuthMetrics{}

var discardAuthMetrics auth.Metrics = auth.DiscardMetrics{}

// DiscardMetrics ignores all config-reader and auth events.
type DiscardMetrics struct{}

func (DiscardMetrics) TokenAcquire(context.Context, string, string, time.Duration, int) {}

func (DiscardMetrics) TokenLifetime(context.Context, string, time.Duration) {}

func (DiscardMetrics) TokenRefresh(context.Context, string, string, time.Duration, bool) {}

func (DiscardMetrics) CacheHit(context.Context, string) {}

func (DiscardMetrics) CacheMiss(context.Context, string, string) {}

func (DiscardMetrics) CacheStore(context.Context, string, string) {}

func (DiscardMetrics) CacheRefresh(context.Context, string, string) {}

func (DiscardMetrics) CacheInvalidate(context.Context, string) {}

func (DiscardMetrics) ConfigLoad(context.Context, string, string, time.Duration) {}

func (DiscardMetrics) CredentialsResolve(context.Context, string, string, time.Duration) {}

func (m ExpandedAuthMetrics) metrics() auth.Metrics {
	if m.AuthMetrics != nil {
		return m.AuthMetrics
	}
	return discardAuthMetrics
}

func (m ExpandedAuthMetrics) TokenAcquire(
	ctx context.Context,
	provider string,
	result string,
	duration time.Duration,
	attempt int,
) {
	m.metrics().TokenAcquire(ctx, provider, result, duration, attempt)
}

func (m ExpandedAuthMetrics) TokenLifetime(
	ctx context.Context,
	provider string,
	ttl time.Duration,
) {
	m.metrics().TokenLifetime(ctx, provider, ttl)
}

func (m ExpandedAuthMetrics) TokenRefresh(
	ctx context.Context,
	provider string,
	result string,
	duration time.Duration,
	background bool,
) {
	m.metrics().TokenRefresh(ctx, provider, result, duration, background)
}

func (m ExpandedAuthMetrics) CacheHit(ctx context.Context, provider string) {
	m.metrics().CacheHit(ctx, provider)
}

func (m ExpandedAuthMetrics) CacheMiss(
	ctx context.Context,
	provider string,
	result string,
) {
	m.metrics().CacheMiss(ctx, provider, result)
}

func (m ExpandedAuthMetrics) CacheStore(
	ctx context.Context,
	provider string,
	result string,
) {
	m.metrics().CacheStore(ctx, provider, result)
}

func (m ExpandedAuthMetrics) CacheRefresh(
	ctx context.Context,
	provider string,
	result string,
) {
	m.metrics().CacheRefresh(ctx, provider, result)
}

func (m ExpandedAuthMetrics) CacheInvalidate(ctx context.Context, provider string) {
	m.metrics().CacheInvalidate(ctx, provider)
}

func (ExpandedAuthMetrics) ConfigLoad(context.Context, string, string, time.Duration) {}

func (ExpandedAuthMetrics) CredentialsResolve(context.Context, string, string, time.Duration) {}

// ExpandAuthMetrics returns [Metrics] for auth-only metrics by muting
// config-reader-specific callbacks.
func ExpandAuthMetrics(metrics auth.Metrics) Metrics {
	if metrics == nil {
		return nil
	}
	if fullMetrics, ok := metrics.(Metrics); ok {
		return fullMetrics
	}
	return ExpandedAuthMetrics{
		AuthMetrics: metrics,
	}
}

func configLoad(ctx context.Context, metrics Metrics, source string, result string, duration time.Duration) {
	if metrics != nil {
		metrics.ConfigLoad(ctx, source, result, duration)
	}
}

func credentialsResolve(ctx context.Context, metrics Metrics, source string, result string, duration time.Duration) {
	if metrics != nil {
		metrics.CredentialsResolve(ctx, source, result, duration)
	}
}
