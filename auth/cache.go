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
//   - lifetime: 90% of token lifespan
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

// NewCachedServiceTokener returns a decorated [BearerTokener] that enhances its functionality
// with [BearerToken] caching and automatic background refresh to ensure that token is always valid.
//
// Recommended parameters from the IAM team:
//   - lifetime: 90% of token lifespan
//   - initial retry: 1 second
//   - max retry: 1 minute
func NewCachedServiceTokener(
	logger *slog.Logger,
	tokener BearerTokener,
	lifetime float64,
	initialRetry time.Duration,
	retryMultiplier float64,
	maxRetry time.Duration,
) *CachedServiceTokener {
	stoppedTicker := time.NewTicker(time.Minute)
	stoppedTicker.Stop()
	return &CachedServiceTokener{
		logger:          logger,
		tokener:         tokener,
		lifetime:        lifetime,
		initialRetry:    initialRetry,
		retryMultiplier: retryMultiplier,
		maxRetry:        maxRetry,
		ticker:          stoppedTicker,
		now:             time.Now,
		group:           singleflight.Group{},

		mu:         sync.RWMutex{},
		cache:      nil,
		refreshAt:  time.Time{},
		retryCount: 0,
	}
}

func (c *CachedServiceTokener) Run(ctx context.Context) error { //nolint:gocognit
	for {
		select {
		case <-ctx.Done():
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
				"Token is about to expire; initiating background token refresh",
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
					"Background token refresh failed",
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
	res, err, _ := c.group.Do("", func() (interface{}, error) {
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
					"Received already expired token",
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
