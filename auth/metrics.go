package auth

import (
	"context"
	"sync/atomic"
	"time"
)

// Metrics receives auth-related observations from tokeners and auth interceptors.
//
// Provider labels are low-cardinality token provider types. Decorators report
// cache and refresh events using the decorated provider type, while the inner
// token acquisition may report a lower-level implementation type. In
// particular, exchange-based providers can report cache events as
// "service-account" or "workload-identity" while their token acquisition is
// reported as "token-exchange".
type Metrics interface {
	TokenAcquire(context.Context, string, string, time.Duration, int)
	TokenLifetime(context.Context, string, time.Duration)
	TokenRefresh(context.Context, string, string, time.Duration, bool)
	CacheHit(context.Context, string)
	CacheMiss(context.Context, string, string)
	CacheStore(context.Context, string, string)
	CacheRefresh(context.Context, string, string)
	CacheInvalidate(context.Context, string)
}

// DiscardMetrics ignores all auth events.
type DiscardMetrics struct{}

func (DiscardMetrics) TokenAcquire(context.Context, string, string, time.Duration, int) {}

func (DiscardMetrics) TokenLifetime(context.Context, string, time.Duration) {}

func (DiscardMetrics) TokenRefresh(context.Context, string, string, time.Duration, bool) {}

func (DiscardMetrics) CacheHit(context.Context, string) {}

func (DiscardMetrics) CacheMiss(context.Context, string, string) {}

func (DiscardMetrics) CacheStore(context.Context, string, string) {}

func (DiscardMetrics) CacheRefresh(context.Context, string, string) {}

func (DiscardMetrics) CacheInvalidate(context.Context, string) {}

// MetricsSetter propagates auth metrics to tokeners that support instrumentation.
type MetricsSetter interface {
	SetMetrics(Metrics)
}

type metricsOption struct {
	metrics Metrics
}

type tokenAcquireAttemptContextKey struct{}

const (
	metricResultSuccess = "success"
	metricResultError   = "error"
)

var discardMetrics Metrics = DiscardMetrics{}

type atomicMetrics struct {
	metrics atomic.Pointer[Metrics]
}

func (m *atomicMetrics) Store(metrics Metrics) {
	if metrics == nil {
		metrics = discardMetrics
	}
	m.metrics.Store(&metrics)
}

func (m *atomicMetrics) Load() Metrics {
	if m == nil {
		return discardMetrics
	}
	metrics := m.metrics.Load()
	if metrics == nil || *metrics == nil {
		return discardMetrics
	}
	return *metrics
}

func (o metricsOption) apply(t BearerTokener) {
	applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
		setter, ok := t.(MetricsSetter)
		if ok {
			setter.SetMetrics(o.metrics)
		}
		return ok
	})
}

// WithMetrics sets metrics for token providers that support instrumentation.
func WithMetrics(metrics Metrics) Option {
	return metricsOption{metrics: metrics}
}

func contextWithTokenAcquireAttempt(ctx context.Context, attempt int) context.Context {
	if attempt <= 0 {
		return ctx
	}
	return context.WithValue(ctx, tokenAcquireAttemptContextKey{}, attempt)
}

func tokenAcquireAttempt(ctx context.Context, attempt int) int {
	if attempt > 0 {
		return attempt
	}
	if ctxAttempt, ok := ctx.Value(tokenAcquireAttemptContextKey{}).(int); ok && ctxAttempt > 0 {
		return ctxAttempt
	}
	return 1
}

func metricProvider(t BearerTokener) string {
	if typ, ok := TypeOfTokener(t); ok {
		return typ
	}
	return "custom"
}

func (m *atomicMetrics) tokenAcquireError(ctx context.Context, t BearerTokener, duration time.Duration, attempt int) {
	m.Load().TokenAcquire(ctx, metricProvider(t), metricResultError, duration, tokenAcquireAttempt(ctx, attempt))
}

func (m *atomicMetrics) tokenAcquireSuccess(
	ctx context.Context,
	t BearerTokener,
	duration time.Duration,
	attempt int,
	token BearerToken,
	now time.Time,
) {
	m.Load().TokenAcquire(ctx, metricProvider(t), metricResultSuccess, duration, tokenAcquireAttempt(ctx, attempt))
	m.tokenLifetime(ctx, t, token, now)
}

func (m *atomicMetrics) tokenLifetime(ctx context.Context, t BearerTokener, token BearerToken, now time.Time) {
	if token.ExpiresAt.IsZero() {
		return
	}
	ttl := max(token.ExpiresAt.Sub(now), 0)
	m.Load().TokenLifetime(ctx, metricProvider(t), ttl)
}

func (m *atomicMetrics) tokenRefresh(ctx context.Context, t BearerTokener, result string, duration time.Duration, background bool) {
	m.Load().TokenRefresh(ctx, metricProvider(t), result, duration, background)
}

func (m *atomicMetrics) cacheHit(ctx context.Context, t BearerTokener) {
	m.Load().CacheHit(ctx, metricProvider(t))
}

func (m *atomicMetrics) cacheMiss(ctx context.Context, t BearerTokener, result string) {
	m.Load().CacheMiss(ctx, metricProvider(t), result)
}

func (m *atomicMetrics) cacheStore(ctx context.Context, t BearerTokener, result string) {
	m.Load().CacheStore(ctx, metricProvider(t), result)
}

func (m *atomicMetrics) cacheRefresh(ctx context.Context, t BearerTokener, result string) {
	m.Load().CacheRefresh(ctx, metricProvider(t), result)
}

func (m *atomicMetrics) cacheInvalidate(ctx context.Context, t BearerTokener) {
	m.Load().CacheInvalidate(ctx, metricProvider(t))
}
