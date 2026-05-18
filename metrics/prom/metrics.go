package prom

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	reader "github.com/nebius/gosdk/config/reader"
)

const (
	resultSuccess = "success"
	resultError   = "error"
)

// AuthMetrics implements gosdk auth metrics using Prometheus.
type AuthMetrics struct {
	name                 string
	tokenAcquireDuration *prometheus.HistogramVec
	tokenRefresh         *prometheus.CounterVec
	tokenRefreshDuration *prometheus.HistogramVec
	tokenLifetime        *prometheus.GaugeVec
	cacheHit             *prometheus.CounterVec
	cacheMiss            *prometheus.CounterVec
	cacheStore           *prometheus.CounterVec
	cacheRefresh         *prometheus.CounterVec
	cacheInvalidate      *prometheus.CounterVec
}

// Metrics implements gosdk auth and config-reader metrics using Prometheus.
type Metrics struct {
	*AuthMetrics
	name               string
	configLoad         *prometheus.HistogramVec
	credentialsResolve *prometheus.HistogramVec
}

var _ reader.Metrics = (*Metrics)(nil)

// Names configures the Prometheus metric names emitted by this package.
//
// Prefix, if set via [WithPrefix], is prepended to each base name.
type Names struct {
	TokenAcquireDuration string
	TokenRefresh         string
	TokenRefreshDuration string
	TokenLifetime        string
	CacheHit             string
	CacheMiss            string
	CacheStore           string
	CacheRefresh         string
	CacheInvalidate      string
	ConfigLoad           string
	CredentialsResolve   string
}

// DefaultNames returns the default metric base names.
//
// Gosdk applies the prefix separately, so these names stay prefix-free.
// The default prefix is `gosdk_`.
func DefaultNames() Names {
	return Names{
		TokenAcquireDuration: "auth_token_acquire_seconds",

		TokenRefresh: "auth_token_refresh_total", TokenRefreshDuration: "auth_token_refresh_seconds",

		TokenLifetime: "auth_token_lifetime_seconds", CacheHit: "auth_cache_hit_total", CacheMiss: "auth_cache_miss_total",
		CacheStore:         "auth_cache_store_total",
		CacheRefresh:       "auth_cache_refresh_total",
		CacheInvalidate:    "auth_cache_invalidate_total",
		ConfigLoad:         "config_load_seconds",
		CredentialsResolve: "credentials_resolve_seconds",
	}
}

type config struct {
	prefix string
	names  Names
	name   string
}

// Option customizes Prometheus metric registration.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (f optionFunc) apply(cfg *config) {
	f(cfg)
}

// WithPrefix prepends a prefix to every metric base name.
//
// If the prefix is non-empty and does not end with `_`, one is added.
func WithPrefix(prefix string) Option {
	return optionFunc(func(cfg *config) {
		cfg.prefix = prefix
	})
}

// WithNames overrides any subset of metric base names.
func WithNames(names Names) Option {
	return optionFunc(func(cfg *config) {
		if names.TokenAcquireDuration != "" {
			cfg.names.TokenAcquireDuration = names.TokenAcquireDuration
		}
		if names.TokenRefresh != "" {
			cfg.names.TokenRefresh = names.TokenRefresh
		}
		if names.TokenRefreshDuration != "" {
			cfg.names.TokenRefreshDuration = names.TokenRefreshDuration
		}
		if names.TokenLifetime != "" {
			cfg.names.TokenLifetime = names.TokenLifetime
		}
		if names.CacheHit != "" {
			cfg.names.CacheHit = names.CacheHit
		}
		if names.CacheMiss != "" {
			cfg.names.CacheMiss = names.CacheMiss
		}
		if names.CacheStore != "" {
			cfg.names.CacheStore = names.CacheStore
		}
		if names.CacheRefresh != "" {
			cfg.names.CacheRefresh = names.CacheRefresh
		}
		if names.CacheInvalidate != "" {
			cfg.names.CacheInvalidate = names.CacheInvalidate
		}
		if names.ConfigLoad != "" {
			cfg.names.ConfigLoad = names.ConfigLoad
		}
		if names.CredentialsResolve != "" {
			cfg.names.CredentialsResolve = names.CredentialsResolve
		}
	})
}

// WithName sets a low-cardinality connector name label on all emitted metrics.
//
// The default name is empty. Use [Metrics.WithName] or [AuthMetrics.WithName]
// to create several named connectors that share the same registered collectors.
func WithName(name string) Option {
	return optionFunc(func(cfg *config) {
		cfg.name = name
	})
}

// NewAuthFromRegistry registers and returns a Prometheus-backed auth metrics
// recorder.
func NewAuthFromRegistry(registerer prometheus.Registerer, opts ...Option) (*AuthMetrics, error) {
	if registerer == nil {
		return nil, errors.New("nil Prometheus registerer")
	}
	cfg := newConfig(opts...)
	return newAuthFromRegistry(registerer, cfg)
}

// NewFromRegistry registers and returns a Prometheus-backed metrics recorder.
func NewFromRegistry(registerer prometheus.Registerer, opts ...Option) (*Metrics, error) {
	if registerer == nil {
		return nil, errors.New("nil Prometheus registerer")
	}
	cfg := newConfig(opts...)
	authMetrics, err := newAuthFromRegistry(registerer, cfg)
	if err != nil {
		return nil, err
	}

	metrics := &Metrics{
		AuthMetrics: authMetrics,
		name:        cfg.name,
	}

	metrics.configLoad, err = register(registerer, prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricName(cfg.prefix, cfg.names.ConfigLoad),
			Help:    "Duration of config reader load operations in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"name", "source", "result"},
	))
	if err != nil {
		return nil, err
	}

	metrics.credentialsResolve, err = register(registerer, prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricName(cfg.prefix, cfg.names.CredentialsResolve),
			Help:    "Duration of credentials resolution operations in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"name", "source", "result"},
	))
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func newConfig(opts ...Option) config {
	cfg := config{
		prefix: defaultPrefix(),
		names:  DefaultNames(),
	}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	return cfg
}

func newAuthFromRegistry(registerer prometheus.Registerer, cfg config) (*AuthMetrics, error) {
	metrics := &AuthMetrics{
		name: cfg.name,
	}
	var err error

	metrics.tokenAcquireDuration, err = register(registerer, prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricName(cfg.prefix, cfg.names.TokenAcquireDuration),
			Help:    "Duration of token acquisition attempts in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"name", "provider", "result", "attempt"},
	))
	if err != nil {
		return nil, err
	}

	metrics.tokenRefresh, err = register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName(cfg.prefix, cfg.names.TokenRefresh),
			Help: "Total number of token refresh attempts.",
		},
		[]string{"name", "provider", "success"},
	))
	if err != nil {
		return nil, err
	}

	metrics.tokenRefreshDuration, err = register(registerer, prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metricName(cfg.prefix, cfg.names.TokenRefreshDuration),
			Help:    "Duration of token refresh attempts in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"name", "provider", "result", "background"},
	))
	if err != nil {
		return nil, err
	}

	metrics.tokenLifetime, err = register(registerer, prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricName(cfg.prefix, cfg.names.TokenLifetime),
			Help: "Lifetime of the newly fetched token in seconds.",
		},
		[]string{"name", "provider"},
	))
	if err != nil {
		return nil, err
	}

	metrics.cacheHit, err = register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName(cfg.prefix, cfg.names.CacheHit),
			Help: "Total number of cache hits.",
		},
		[]string{"name", "provider"},
	))
	if err != nil {
		return nil, err
	}

	metrics.cacheMiss, err = register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName(cfg.prefix, cfg.names.CacheMiss),
			Help: "Total number of cache misses.",
		},
		[]string{"name", "provider", "result"},
	))
	if err != nil {
		return nil, err
	}

	metrics.cacheStore, err = register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName(cfg.prefix, cfg.names.CacheStore),
			Help: "Total number of cache store attempts.",
		},
		[]string{"name", "provider", "result"},
	))
	if err != nil {
		return nil, err
	}

	metrics.cacheRefresh, err = register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName(cfg.prefix, cfg.names.CacheRefresh),
			Help: "Total number of cache refresh attempts.",
		},
		[]string{"name", "provider", "result"},
	))
	if err != nil {
		return nil, err
	}

	metrics.cacheInvalidate, err = register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricName(cfg.prefix, cfg.names.CacheInvalidate),
			Help: "Total number of cache invalidations.",
		},
		[]string{"name", "provider"},
	))
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

// WithName returns a connector that shares the registered auth collectors and
// emits metrics with the provided low-cardinality connector name.
func (m *AuthMetrics) WithName(name string) *AuthMetrics {
	if m == nil {
		return nil
	}
	clone := *m
	clone.name = name
	return &clone
}

// WithName returns a connector that shares the registered collectors and emits
// metrics with the provided low-cardinality connector name.
func (m *Metrics) WithName(name string) *Metrics {
	if m == nil {
		return nil
	}
	clone := *m
	clone.name = name
	clone.AuthMetrics = m.AuthMetrics.WithName(name)
	return &clone
}

func (m *AuthMetrics) TokenAcquire(_ context.Context, provider string, result string, duration time.Duration, attempt int) {
	m.tokenAcquireDuration.WithLabelValues(m.name, provider, result, strconv.Itoa(attempt)).Observe(duration.Seconds())
}

func (m *AuthMetrics) TokenLifetime(_ context.Context, provider string, ttl time.Duration) {
	m.tokenLifetime.WithLabelValues(m.name, provider).Set(ttl.Seconds())
}

func (m *AuthMetrics) TokenRefresh(_ context.Context, provider string, result string, duration time.Duration, background bool) {
	m.tokenRefresh.WithLabelValues(m.name, provider, strconv.FormatBool(result == resultSuccess)).Inc()
	m.tokenRefreshDuration.WithLabelValues(m.name, provider, result, strconv.FormatBool(background)).Observe(duration.Seconds())
}

func (m *AuthMetrics) CacheHit(_ context.Context, provider string) {
	m.cacheHit.WithLabelValues(m.name, provider).Inc()
}

func (m *AuthMetrics) CacheMiss(_ context.Context, provider string, result string) {
	m.cacheMiss.WithLabelValues(m.name, provider, result).Inc()
}

func (m *AuthMetrics) CacheStore(_ context.Context, provider string, result string) {
	m.cacheStore.WithLabelValues(m.name, provider, result).Inc()
}

func (m *AuthMetrics) CacheRefresh(_ context.Context, provider string, result string) {
	m.cacheRefresh.WithLabelValues(m.name, provider, result).Inc()
}

func (m *AuthMetrics) CacheInvalidate(_ context.Context, provider string) {
	m.cacheInvalidate.WithLabelValues(m.name, provider).Inc()
}

func (m *Metrics) ConfigLoad(_ context.Context, source string, result string, duration time.Duration) {
	m.configLoad.WithLabelValues(m.name, source, result).Observe(duration.Seconds())
}

func (m *Metrics) CredentialsResolve(_ context.Context, source string, result string, duration time.Duration) {
	m.credentialsResolve.WithLabelValues(m.name, source, result).Observe(duration.Seconds())
}

func metricName(prefix string, name string) string {
	if prefix == "" {
		return name
	}
	if !strings.HasSuffix(prefix, "_") {
		prefix += "_"
	}
	return prefix + name
}

func defaultPrefix() string {

	return "gosdk"
}
func register[T prometheus.Collector](registerer prometheus.Registerer, collector T) (T, error) {
	if err := registerer.Register(collector); err != nil {
		var zero T
		return zero, err
	}
	return collector, nil
}
