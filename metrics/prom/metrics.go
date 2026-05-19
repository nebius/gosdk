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

	labelAttempt    = "attempt"
	labelBackground = "background"
	labelName       = "name"
	labelProvider   = "provider"
	labelResult     = "result"
	labelSource     = "source"
	labelSuccess    = "success"

	defaultTokenRefreshMetricName  = "auth_token_refresh_total"    //nolint:gosec // metric name, not a credential.
	defaultTokenLifetimeMetricName = "auth_token_lifetime_seconds" //nolint:gosec // metric name, not a credential.
	defaultPrefixValue             = "gosdk"
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
	return Names{ //nolint:gosec // Prometheus metric names contain "token"; they are not credentials.
		TokenAcquireDuration: "auth_token_acquire_seconds",
		TokenRefresh:         defaultTokenRefreshMetricName,
		TokenRefreshDuration: "auth_token_refresh_seconds",
		TokenLifetime:        defaultTokenLifetimeMetricName,
		CacheHit:             "auth_cache_hit_total",
		CacheMiss:            "auth_cache_miss_total",
		CacheStore:           "auth_cache_store_total",
		CacheRefresh:         "auth_cache_refresh_total",
		CacheInvalidate:      "auth_cache_invalidate_total",
		ConfigLoad:           "config_load_seconds",
		CredentialsResolve:   "credentials_resolve_seconds",
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
		for _, override := range nameOverrides(&cfg.names, names) {
			if override.value != "" {
				*override.target = override.value
			}
		}
	})
}

type nameOverride struct {
	target *string
	value  string
}

func nameOverrides(target *Names, values Names) []nameOverride {
	return []nameOverride{
		{target: &target.TokenAcquireDuration, value: values.TokenAcquireDuration},
		{target: &target.TokenRefresh, value: values.TokenRefresh},
		{target: &target.TokenRefreshDuration, value: values.TokenRefreshDuration},
		{target: &target.TokenLifetime, value: values.TokenLifetime},
		{target: &target.CacheHit, value: values.CacheHit},
		{target: &target.CacheMiss, value: values.CacheMiss},
		{target: &target.CacheStore, value: values.CacheStore},
		{target: &target.CacheRefresh, value: values.CacheRefresh},
		{target: &target.CacheInvalidate, value: values.CacheInvalidate},
		{target: &target.ConfigLoad, value: values.ConfigLoad},
		{target: &target.CredentialsResolve, value: values.CredentialsResolve},
	}
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
		[]string{labelName, labelSource, labelResult},
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
		[]string{labelName, labelSource, labelResult},
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

	metrics.tokenAcquireDuration, err = registerHistogram(
		registerer,
		metricName(cfg.prefix, cfg.names.TokenAcquireDuration),
		"Duration of token acquisition attempts in seconds.",
		[]string{labelName, labelProvider, labelResult, labelAttempt},
	)
	if err != nil {
		return nil, err
	}

	metrics.tokenRefresh, err = registerCounter(
		registerer,
		metricName(cfg.prefix, cfg.names.TokenRefresh),
		"Total number of token refresh attempts.",
		[]string{labelName, labelProvider, labelSuccess},
	)
	if err != nil {
		return nil, err
	}

	metrics.tokenRefreshDuration, err = registerHistogram(
		registerer,
		metricName(cfg.prefix, cfg.names.TokenRefreshDuration),
		"Duration of token refresh attempts in seconds.",
		[]string{labelName, labelProvider, labelResult, labelBackground},
	)
	if err != nil {
		return nil, err
	}

	metrics.tokenLifetime, err = registerGauge(
		registerer,
		metricName(cfg.prefix, cfg.names.TokenLifetime),
		"Lifetime of the newly fetched token in seconds.",
		[]string{labelName, labelProvider},
	)
	if err != nil {
		return nil, err
	}

	metrics.cacheHit, err = registerCounter(
		registerer,
		metricName(cfg.prefix, cfg.names.CacheHit),
		"Total number of cache hits.",
		[]string{labelName, labelProvider},
	)
	if err != nil {
		return nil, err
	}

	metrics.cacheMiss, err = registerCounter(
		registerer,
		metricName(cfg.prefix, cfg.names.CacheMiss),
		"Total number of cache misses.",
		[]string{labelName, labelProvider, labelResult},
	)
	if err != nil {
		return nil, err
	}

	metrics.cacheStore, err = registerCounter(
		registerer,
		metricName(cfg.prefix, cfg.names.CacheStore),
		"Total number of cache store attempts.",
		[]string{labelName, labelProvider, labelResult},
	)
	if err != nil {
		return nil, err
	}

	metrics.cacheRefresh, err = registerCounter(
		registerer,
		metricName(cfg.prefix, cfg.names.CacheRefresh),
		"Total number of cache refresh attempts.",
		[]string{labelName, labelProvider, labelResult},
	)
	if err != nil {
		return nil, err
	}

	metrics.cacheInvalidate, err = registerCounter(
		registerer,
		metricName(cfg.prefix, cfg.names.CacheInvalidate),
		"Total number of cache invalidations.",
		[]string{labelName, labelProvider},
	)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func registerHistogram(
	registerer prometheus.Registerer,
	name string,
	help string,
	labels []string,
) (*prometheus.HistogramVec, error) {
	return register(registerer, prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    name,
			Help:    help,
			Buckets: prometheus.DefBuckets,
		},
		labels,
	))
}

func registerCounter(
	registerer prometheus.Registerer,
	name string,
	help string,
	labels []string,
) (*prometheus.CounterVec, error) {
	return register(registerer, prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	))
}

func registerGauge(
	registerer prometheus.Registerer,
	name string,
	help string,
	labels []string,
) (*prometheus.GaugeVec, error) {
	return register(registerer, prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		},
		labels,
	))
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

func (m *AuthMetrics) TokenAcquire(
	_ context.Context,
	provider string,
	result string,
	duration time.Duration,
	attempt int,
) {
	m.tokenAcquireDuration.WithLabelValues(m.name, provider, result, strconv.Itoa(attempt)).Observe(duration.Seconds())
}

func (m *AuthMetrics) TokenLifetime(_ context.Context, provider string, ttl time.Duration) {
	m.tokenLifetime.WithLabelValues(m.name, provider).Set(ttl.Seconds())
}

func (m *AuthMetrics) TokenRefresh(
	_ context.Context,
	provider string,
	result string,
	duration time.Duration,
	background bool,
) {
	m.tokenRefresh.WithLabelValues(m.name, provider, strconv.FormatBool(result == resultSuccess)).Inc()
	m.tokenRefreshDuration.
		WithLabelValues(m.name, provider, result, strconv.FormatBool(background)).
		Observe(duration.Seconds())
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
	return defaultPrefixValue
}

func register[T prometheus.Collector](registerer prometheus.Registerer, collector T) (T, error) {
	if err := registerer.Register(collector); err != nil {
		var zero T
		return zero, err
	}
	return collector, nil
}
