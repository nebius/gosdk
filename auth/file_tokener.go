package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/nebius/gosdk/config/paths"
)

// fileTokener is a [BearerTokener] that serves token from file, reading it every time.
type fileTokener struct {
	path    string
	metrics atomicMetrics
}

var _ BearerTokener = (*fileTokener)(nil)
var _ MetricsSetter = (*fileTokener)(nil)
var _ TypedTokener = (*fileTokener)(nil)

// newFileTokener returns a [BearerTokener] that serves token from file, reading it every time.
// The path is expanded at construction, but file existence, type, permissions,
// and token contents are validated lazily when [BearerTokener.BearerToken] is called.
func newFileTokener(path string) (*fileTokener, error) {
	path, err := paths.ExpandHomeDir(path)
	if err != nil {
		return nil, fmt.Errorf("expand home: %w", err)
	}
	return &fileTokener{path: path}, nil
}

func (t *fileTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	start := time.Now()
	file := t.path
	data, err := os.ReadFile(file)
	if err != nil {
		t.metrics.tokenAcquireError(ctx, t, time.Since(start), 0)
		return BearerToken{}, fmt.Errorf("read file: %w", err)
	}
	token := strings.TrimSpace(string(data))
	if strings.Contains(token, "\n") {
		t.metrics.tokenAcquireError(ctx, t, time.Since(start), 0)
		return BearerToken{}, fmt.Errorf("invalid token file: %s contains newline", file)
	}
	bearerToken := BearerToken{
		Token:     token,
		ExpiresAt: time.Time{},
	}
	t.metrics.tokenAcquireSuccess(ctx, t, time.Since(start), 0, bearerToken, time.Now())
	return bearerToken, nil
}

func (t *fileTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	rt, tokErr := t.BearerToken(ctx)
	if tokErr != nil {
		return errors.Join(tokErr, err)
	}
	if rt.Token != token.Token {
		return nil // token has changed
	}

	return err // Same token, but it is invalid, fail
}

func (t *fileTokener) Type() string {
	return "file"
}

func (t *fileTokener) SetMetrics(metrics Metrics) {
	t.metrics.Store(metrics)
}

type FileTokener struct {
	tokener       *fileTokener
	mu            sync.RWMutex
	metrics       atomicMetrics
	token         *BearerToken
	refreshPeriod time.Duration
	refreshAt     time.Time // cache freshness deadline, not BearerToken.ExpiresAt
	generation    uint64
}

var _ BearerTokener = (*FileTokener)(nil)
var _ MetricsSetter = (*FileTokener)(nil)
var _ Wrapper = (*FileTokener)(nil)

func NewFileTokener(
	path string,
	refreshPeriod time.Duration,
	opts ...Option,
) (*FileTokener, error) {
	tokener, err := newFileTokener(path)
	if err != nil {
		return nil, err
	}
	res := &FileTokener{
		tokener:       tokener,
		refreshPeriod: refreshPeriod,
	}
	applyOptions(res, opts...)
	return res, nil
}

func (c *FileTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	cachedToken, generation, ok := c.cachedToken(time.Now())
	if ok {
		c.metrics.cacheHit(ctx, c)
		return cachedToken, nil
	}

	token, err := c.tokener.BearerToken(ctx)
	if err != nil {
		c.metrics.cacheMiss(ctx, c, metricResultError)
		return BearerToken{}, err
	}
	c.storeToken(token, time.Now().Add(c.refreshPeriod), generation)
	c.metrics.cacheMiss(ctx, c, metricResultSuccess)
	return token, nil
}

func (c *FileTokener) HandleError(ctx context.Context, token BearerToken, err error) error {
	c.mu.Lock()
	c.token = nil // force refresh on next call
	c.refreshAt = time.Time{}
	c.generation++
	c.mu.Unlock()

	c.metrics.cacheInvalidate(ctx, c)
	return c.tokener.HandleError(ctx, token, err)
}

func (c *FileTokener) SetMetrics(metrics Metrics) {
	c.metrics.Store(metrics)

	if c.tokener != nil {
		c.tokener.SetMetrics(metrics)
	}
}

func (c *FileTokener) Unwrap() BearerTokener {
	return c.tokener
}

func (c *FileTokener) cachedToken(now time.Time) (BearerToken, uint64, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	generation := c.generation
	if c.token == nil {
		return BearerToken{}, generation, false
	}
	if c.refreshPeriod != 0 && now.After(c.refreshAt) {
		return BearerToken{}, generation, false
	}
	return *c.token, generation, true
}

func (c *FileTokener) storeToken(token BearerToken, refreshAt time.Time, generation uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.generation != generation {
		return
	}
	c.token = &token
	c.refreshAt = refreshAt
	c.generation++
}
