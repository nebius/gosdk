package auth

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// CachedServiceTokener is a [BearerTokener] decorator that enhances its functionality
// with [BearerToken] caching and automatic background refresh to ensure that token is always valid.
//
// Recommended parameters from the IAM team:
//   - lifetime: 0.9 of token lifespan (90% of the token lifespan)
//   - initial retry: 1 second
//   - max retry: 1 minute
type CachedServiceTokener struct {
	logger          *slog.Logger
	tokener         BearerTokener
	lifetime        float64 // lifetime is fraction of token lifespan after which the token is refreshed
	initialRetry    time.Duration
	retryMultiplier float64
	maxRetry        time.Duration
	ticker          *time.Ticker
	now             func() time.Time
	group           singleflight.Group

	mu         sync.RWMutex
	cache      *BearerToken
	refreshAt  time.Time
	retryCount int
}

var _ BearerTokener = (*CachedServiceTokener)(nil)

const (
	// Recommended by the IAM team: refresh after 90% of the token lifetime.
	defaultCachedTokenerLifetime = 0.9
	// Recommended by the IAM team: start retrying after 1 second.
	defaultCachedTokenerInitialRetry = time.Second
	// Recommended by the IAM team: increase retry delay by 1.5x on each attempt.
	defaultCachedTokenerRetryMultiplier = 1.5
	// Recommended by the IAM team: cap retry delay at 1 minute.
	defaultCachedTokenerMaxRetry = time.Minute

	stoppedTickerInterval = time.Minute
)

type CachedTokenerOptionFunc func(*CachedServiceTokener)

func (f CachedTokenerOptionFunc) apply(c BearerTokener) {
	if ct, ok := (c).(*CachedServiceTokener); ok {
		f(ct)
	}
}

func WithCachedTokenerLifetime(lifetime float64) Option {
	return CachedTokenerOptionFunc(func(c *CachedServiceTokener) {
		c.lifetime = lifetime
	})
}

func WithCachedTokenerInitialRetry(initialRetry time.Duration) Option {
	return CachedTokenerOptionFunc(func(c *CachedServiceTokener) {
		c.initialRetry = initialRetry
	})
}

func WithCachedTokenerRetryMultiplier(retryMultiplier float64) Option {
	return CachedTokenerOptionFunc(func(c *CachedServiceTokener) {
		c.retryMultiplier = retryMultiplier
	})
}

func WithCachedTokenerMaxRetry(maxRetry time.Duration) Option {
	return CachedTokenerOptionFunc(func(c *CachedServiceTokener) {
		c.maxRetry = maxRetry
	})
}

func NewCachedTokener(tokener BearerTokener, opts ...Option) *CachedServiceTokener {
	c := &CachedServiceTokener{
		tokener:         tokener,
		lifetime:        defaultCachedTokenerLifetime,
		initialRetry:    0,
		retryMultiplier: 0,
		maxRetry:        0,
		ticker:          time.NewTicker(stoppedTickerInterval), // will be stopped immediately
		now:             time.Now,

		mu:         sync.RWMutex{},
		cache:      nil,
		refreshAt:  time.Time{},
		retryCount: 0,
	}

	c.ticker.Stop()
	for _, opt := range opts {
		opt.apply(c)
	}

	if c.initialRetry <= 0 {
		c.initialRetry = defaultCachedTokenerInitialRetry
	}
	if c.retryMultiplier <= 0 {
		c.retryMultiplier = defaultCachedTokenerRetryMultiplier
	}
	if c.maxRetry <= 0 {
		c.maxRetry = defaultCachedTokenerMaxRetry
	}
	if c.logger == nil {
		c.logger = slog.New(slog.DiscardHandler)
	}
	return c
}

// NewCachedServiceTokener returns a decorated [BearerTokener] that enhances its functionality
// with [BearerToken] caching and automatic background refresh to ensure that token is always valid.
//
// Recommended parameters from the IAM team:
//   - lifetime: 0.9 of token lifespan (90% of the token lifespan)
//   - initial retry: 1 second
//   - max retry: 1 minute
//
// Note on the retry multiplier: in the new version, it is automatically increased by
// 1.5x on each retry, while in the deprecated version, it is fixed to 1.5x. The new
// approach is more resilient to transient errors while still preventing excessive retries.
// This function retains the 1x multiplier option for backward compatibility,
// but it is recommended to use the new version with automatic multiplier increase.
//
// Deprecated: use NewCachedTokener instead.
func NewCachedServiceTokener(
	logger *slog.Logger,
	tokener BearerTokener,
	lifetime float64,
	initialRetry time.Duration,
	retryMultiplier float64,
	maxRetry time.Duration,
) *CachedServiceTokener {
	logger.Warn("you're using a deprecated service tokener, consider switching to" +
		" NewCachedTokener with options for better retry behavior and easier maintenance." +
		" See the documentation for details.")
	if retryMultiplier <= 0 {
		retryMultiplier = 1 // for backward compatibility
	}
	return NewCachedTokener(
		tokener,
		WithLogger(logger),
		WithCachedTokenerLifetime(lifetime),
		WithCachedTokenerInitialRetry(initialRetry),
		WithCachedTokenerRetryMultiplier(retryMultiplier),
		WithCachedTokenerMaxRetry(maxRetry),
	)
}

func (c *CachedServiceTokener) SetLogger(logger *slog.Logger) {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)
	}
	c.logger = logger
}

func (c *CachedServiceTokener) Unwrap() BearerTokener {
	return c.tokener
}
func (c *CachedServiceTokener) Run(ctx context.Context) error { //nolint:gocognit
	go func() { // run underlying async tokener, if necessary
		err, _ := RunAsyncTokener(ctx, c.tokener)
		if err != nil && !errors.Is(err, context.Canceled) {
			c.logger.ErrorContext(
				ctx,
				"background token refresh failed",
				slog.Any("error", err),
			)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			c.logger.InfoContext(ctx, "background token refresh stopped as the context is done")
			return ctx.Err()
		case <-c.ticker.C:

			token, refreshAt, retryCount := c.getToken()

			if token == nil {
				continue
			}

			if token.ExpiresAt.IsZero() {
				continue
			}

			if refreshAt.IsZero() {
				continue
			}

			timeLeft := refreshAt.Sub(c.now())
			if timeLeft > 0 {
				c.ticker.Reset(timeLeft)
				continue
			}

			c.logger.InfoContext(
				ctx,
				"token is about to expire; initiating background token refresh",
				slog.Time("expires_at", token.ExpiresAt),
				slog.Int("attempt", retryCount+1),
			)

			_, err := c.requestToken(ctx, true)
			if err != nil && !errors.Is(err, context.Canceled) {
				var retry time.Duration
				if retryCount <= 0 || math.Abs(c.retryMultiplier-1) <= 1e-9 {
					retry = c.initialRetry
				} else {
					mul := math.Pow(c.retryMultiplier, float64(retryCount))
					retry = time.Duration(math.Min(float64(c.maxRetry), float64(c.initialRetry)*mul))
				}
				c.logger.ErrorContext(
					ctx,
					"background token refresh failed",
					slog.Any("error", err),
					slog.Int("attempt", retryCount+1),
					slog.Duration("next_attempt", retry),
				)
				c.ticker.Reset(retry)
			}
		}
	}
}

func (c *CachedServiceTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	token, _, _ := c.getToken()
	if token != nil {
		return *token, nil
	}

	return c.requestToken(ctx, false)
}

func (c *CachedServiceTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	if err == nil {
		return nil
	}

	c.mu.Lock()
	if c.cache != nil && c.cache.Token == token.Token {
		c.cache = nil
		c.refreshAt = time.Time{}
		c.ticker.Stop()
	}
	c.mu.Unlock()

	return c.tokener.HandleError(ctx, token, err)
}

func (c *CachedServiceTokener) getToken() (*BearerToken, time.Time, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache, c.refreshAt, c.retryCount
}

func (c *CachedServiceTokener) requestToken(ctx context.Context, background bool) (BearerToken, error) {
	res, err, _ := c.group.Do("", func() (any, error) {
		var refreshAfter time.Duration

		now := c.now()
		token, err := c.tokener.BearerToken(ctx)
		if err != nil {
			if background {
				c.mu.Lock()
				c.retryCount++
				c.mu.Unlock()
			}
			return nil, err
		}

		if !token.ExpiresAt.IsZero() {
			ttl := token.ExpiresAt.Sub(now)
			if ttl > 0 {
				refreshAfter = time.Duration(float64(ttl) * c.lifetime)
			} else {
				c.logger.ErrorContext(
					ctx,
					"received already expired token",
					slog.Time("expires_at", token.ExpiresAt),
				)
			}
		}

		c.mu.Lock()
		c.cache = &token
		if refreshAfter > 0 {
			c.refreshAt = now.Add(refreshAfter)
		} else {
			c.refreshAt = time.Time{}
		}
		c.retryCount = 0
		c.mu.Unlock()

		if refreshAfter > 0 {
			c.ticker.Reset(refreshAfter)
		}

		return token, nil
	})
	if err != nil {
		return BearerToken{}, err
	}

	return res.(BearerToken), nil //nolint:errcheck // ok to panic
}

// CachedBearerTokener is a [BearerTokener] decorator that caches the [BearerToken].
// Method [CachedBearerTokener.HandleError] invalidates the cache on any error.
type CachedBearerTokener struct {
	tokener BearerTokener
	group   singleflight.Group

	mu    sync.RWMutex
	cache *BearerToken
}

var _ BearerTokener = (*CachedBearerTokener)(nil)

// NewCachedBearerTokener returns a decorated [BearerTokener] that caches the [BearerToken].
func NewCachedBearerTokener(tokener BearerTokener) *CachedBearerTokener {
	return &CachedBearerTokener{
		tokener: tokener,
		group:   singleflight.Group{},

		mu:    sync.RWMutex{},
		cache: nil,
	}
}

func (c *CachedBearerTokener) Unwrap() BearerTokener {
	return c.tokener
}

func (c *CachedBearerTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	token := c.getToken()
	if token != nil {
		return *token, nil
	}

	return c.requestToken(ctx)
}

func (c *CachedBearerTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	if err == nil {
		return nil
	}

	c.mu.Lock()
	if c.cache != nil && c.cache.Token == token.Token {
		c.cache = nil
	}
	c.mu.Unlock()

	return c.tokener.HandleError(ctx, token, err)
}

func (c *CachedBearerTokener) getToken() *BearerToken {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache
}

func (c *CachedBearerTokener) requestToken(ctx context.Context) (BearerToken, error) {
	res, err, _ := c.group.Do("", func() (interface{}, error) {
		// Token could be already cached by another goroutine
		if token := c.getToken(); token != nil {
			return *token, nil
		}

		token, err := c.tokener.BearerToken(ctx)
		if err != nil {
			return nil, err
		}

		c.mu.Lock()
		c.cache = &token
		c.mu.Unlock()

		return token, nil
	})
	if err != nil {
		return BearerToken{}, err
	}

	return res.(BearerToken), nil //nolint:errcheck // ok to panic
}
