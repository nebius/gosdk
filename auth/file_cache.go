package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gofrs/flock"
	"golang.org/x/sync/singleflight"
	"gopkg.in/yaml.v3"

	"github.com/nebius/gosdk/config"
)

const cacheFileName = "credentials.yaml"

type cacheContents struct {
	Tokens map[string]BearerToken `yaml:"tokens"`
}

type tokenCache struct {
	fileName string
	logger   *slog.Logger
	lockFile *flock.Flock
}

func newTokenCache(
	fileName string,
	logger *slog.Logger,
) *tokenCache {
	return &tokenCache{
		fileName: fileName,
		logger: logger.With(
			slog.String("name", "token_cache"),
			slog.String("file_name", fileName),
		),
		lockFile: flock.New(fileName),
	}
}

func (f *tokenCache) loadCache(ctx context.Context) (*cacheContents, error) {
	bytes, err := os.ReadFile(f.fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.logger.DebugContext(
				ctx,
				"cache file does not exist, returning empty cache",
			)
			return nil, nil //nolint:nilnil // it's fine to return nil, nil from private method
		}
		return nil, fmt.Errorf("read %q: %w", f.fileName, err)
	}

	var cache cacheContents
	if err = yaml.Unmarshal(bytes, &cache); err != nil {
		if f.logger != nil {
			f.logger.WarnContext(
				ctx,
				"failed to unmarshal cache file, purging cache",
				slog.Any("error", err),
			)
		}
		if err = os.Remove(f.fileName); err != nil {
			return nil, fmt.Errorf("remove %q: %w", f.fileName, err)
		}
		f.logger.DebugContext(
			ctx,
			"cache file purged due to unmarshal error",
		)
		return nil, nil //nolint:nilnil // it's fine to return nil, nil from private method
	}

	return &cache, nil
}

func (f *tokenCache) saveCache(ctx context.Context, cache *cacheContents) error {
	for k, v := range cache.Tokens { // always remove expired tokens
		if v.Empty() || v.IsExpired() {
			f.logger.DebugContext(
				ctx,
				"removing expired or empty token from cache",
				slog.String("token_name", k),
				slog.String("token", v.String()),
			)
			delete(cache.Tokens, k)
		}
	}
	bytes, err := yaml.Marshal(cache)
	if err != nil {
		return fmt.Errorf("yaml marshal: %w", err)
	}

	if err = os.WriteFile(f.fileName, bytes, 0600); err != nil {
		return fmt.Errorf("write to %q: %w", f.fileName, err)
	}
	return nil
}

func (f *tokenCache) Get(
	ctx context.Context,
	tokenName string,
) (_ BearerToken, deferError error) {
	// do not try flocking if the directory does not exist
	if _, err := os.Stat(filepath.Dir(f.fileName)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.logger.DebugContext(
				ctx,
				"cache file directory does not exist, returning empty token",
				slog.String("token_name", tokenName),
			)
			return BearerToken{}, nil
		}
		return BearerToken{}, fmt.Errorf("stat cache file directory: %w", err)
	}

	if err := f.lockFile.Lock(); err != nil {
		return BearerToken{}, fmt.Errorf("lock cache file: %w", err)
	}
	defer func() {
		if err := f.lockFile.Unlock(); err != nil {
			deferError = errors.Join(
				deferError,
				fmt.Errorf("unlock cache file: %w", err),
			)
		}
	}()
	cacheFile, err := f.loadCache(ctx)
	if err != nil {
		return BearerToken{}, fmt.Errorf("load cache: %w", err)
	}
	if cacheFile == nil {
		f.logger.DebugContext(
			ctx,
			"cache file is empty or does not exist, returning empty token",
			slog.String("token_name", tokenName),
		)
		return BearerToken{}, nil
	}

	cacheToken, ok := cacheFile.Tokens[tokenName]
	if !ok {
		f.logger.DebugContext(
			ctx,
			"token not found in cache",
			slog.String("token_name", tokenName),
		)
		return BearerToken{}, nil
	}

	f.logger.DebugContext(
		ctx,
		"token retrieved from cache",
		slog.String("token_name", tokenName),
		slog.String("token", cacheToken.String()),
	)

	return cacheToken, nil
}

func (f *tokenCache) Set(
	ctx context.Context,
	tokenName string,
	token BearerToken,
) (deferError error) {
	if token.Empty() {
		f.logger.DebugContext(
			ctx,
			"attempt to set an empty token, removing it from cache",
			slog.String("token_name", tokenName),
		)
		return f.Remove(ctx, tokenName)
	}

	if err := os.MkdirAll(filepath.Dir(f.fileName), 0750); err != nil {
		return fmt.Errorf("create cache file directory: %w", err)
	}

	if err := f.lockFile.Lock(); err != nil {
		return fmt.Errorf("lock cache file: %w", err)
	}
	defer func() {
		if err := f.lockFile.Unlock(); err != nil {
			deferError = errors.Join(
				deferError,
				fmt.Errorf("unlock cache file: %w", err),
			)
		}
	}()

	cache, err := f.loadCache(ctx)
	if err != nil {
		return fmt.Errorf("load cache: %w", err)
	}

	if cache == nil {
		cache = &cacheContents{
			Tokens: make(map[string]BearerToken),
		}
	}
	if cache.Tokens == nil {
		cache.Tokens = make(map[string]BearerToken)
	}

	cache.Tokens[tokenName] = token

	if err = f.saveCache(ctx, cache); err != nil {
		return fmt.Errorf("save cache: %w", err)
	}
	f.logger.DebugContext(
		ctx,
		"token set in cache",
		slog.String("token_name", tokenName),
		slog.String("token", token.String()),
	)
	return nil
}

func (f *tokenCache) Remove(
	ctx context.Context,
	tokenName string,
) (deferError error) {
	// do not try flocking if the directory does not exist
	if _, err := os.Stat(filepath.Dir(f.fileName)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.logger.DebugContext(
				ctx,
				"cache file directory does not exist, nothing to remove",
				slog.String("token_name", tokenName),
			)
			return nil
		}
		return fmt.Errorf("stat cache file directory: %w", err)
	}

	if err := f.lockFile.Lock(); err != nil {
		return fmt.Errorf("lock cache file: %w", err)
	}
	defer func() {
		if err := f.lockFile.Unlock(); err != nil {
			deferError = errors.Join(
				deferError,
				fmt.Errorf("unlock cache file: %w", err),
			)
		}
	}()

	cache, err := f.loadCache(ctx)
	if err != nil {
		return fmt.Errorf("load cache: %w", err)
	}

	if cache != nil && len(cache.Tokens) > 0 {
		f.logger.DebugContext(
			ctx,
			"removing token from cache",
			slog.String("token_name", tokenName),
			slog.String("token", cache.Tokens[tokenName].String()),
		)
		delete(cache.Tokens, tokenName)
	} else {
		f.logger.DebugContext(
			ctx,
			"no tokens in cache or token not found, nothing to remove",
			slog.String("token_name", tokenName),
		)
	}

	if cache != nil && len(cache.Tokens) == 0 {
		if err = os.Remove(f.fileName); err != nil {
			return fmt.Errorf("remove cache file %q: %w", f.fileName, err)
		}
		f.logger.DebugContext(
			ctx,
			"cache file is empty, removing it",
		)
	}

	if err = f.saveCache(ctx, cache); err != nil {
		return fmt.Errorf("save cache: %w", err)
	}
	f.logger.DebugContext(
		ctx,
		"token removed from cache",
		slog.String("token_name", tokenName),
	)
	return nil
}

func (f *tokenCache) RemoveIfEqual( //nolint:gocognit
	ctx context.Context,
	tokenName string,
	token BearerToken,
) (deferError error) {
	// do not try flocking if the directory does not exist
	if _, err := os.Stat(filepath.Dir(f.fileName)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.logger.DebugContext(
				ctx,
				"cache file directory does not exist, nothing to remove",
				slog.String("token_name", tokenName),
				slog.String("token", token.String()),
			)
			return nil
		}
		return fmt.Errorf("stat cache file directory: %w", err)
	}

	if err := f.lockFile.Lock(); err != nil {
		return fmt.Errorf("lock cache file: %w", err)
	}
	defer func() {
		if err := f.lockFile.Unlock(); err != nil {
			deferError = errors.Join(
				deferError,
				fmt.Errorf("unlock cache file: %w", err),
			)
		}
	}()

	cache, err := f.loadCache(ctx)
	if err != nil {
		return fmt.Errorf("load cache: %w", err)
	}

	if cache != nil { //nolint:nestif // it's ok to have so many ifs
		if existingToken, exists := cache.Tokens[tokenName]; exists {
			if existingToken.Equal(token) {
				f.logger.DebugContext(
					ctx,
					"removing token from cache as it is equal to the provided one",
					slog.String("token_name", tokenName),
					slog.String("token", existingToken.String()),
				)
				delete(cache.Tokens, tokenName)
			} else {
				f.logger.DebugContext(
					ctx,
					"token not removed from cache as it is not equal to the provided one",
					slog.String("token_name", tokenName),
					slog.String("existing_token", existingToken.String()),
					slog.String("provided_token", token.String()),
				)
			}
		} else {
			f.logger.DebugContext(
				ctx,
				"token not found in cache, nothing to remove",
				slog.String("token_name", tokenName),
				slog.String("token", token.String()),
			)
		}
	} else {
		f.logger.DebugContext(
			ctx,
			"cache file is empty or does not exist, nothing to remove",
			slog.String("token_name", tokenName),
			slog.String("token", token.String()),
		)
	}

	if cache != nil && len(cache.Tokens) == 0 {
		f.logger.DebugContext(
			ctx,
			"cache file is empty, removing it",
		)
		if err = os.Remove(f.fileName); err != nil {
			return fmt.Errorf("remove cache file %q: %w", f.fileName, err)
		}
	}

	if err = f.saveCache(ctx, cache); err != nil {
		return fmt.Errorf("save cache: %w", err)
	}
	f.logger.DebugContext(
		ctx,
		"token removed from cache",
		slog.String("token_name", tokenName),
	)
	return nil
}

type throttledTokenCache struct {
	cache       *tokenCache
	throttle    time.Duration
	name        string
	cachedToken BearerToken
	nextCheck   time.Time
}

func newThrottledTokenCache(
	fileName string,
	throttle time.Duration,
	name string,
	logger *slog.Logger,
) *throttledTokenCache {
	return &throttledTokenCache{
		cache:       newTokenCache(fileName, logger),
		throttle:    throttle,
		name:        name,
		cachedToken: BearerToken{},
		nextCheck:   time.Now(),
	}
}

func (t *throttledTokenCache) SetCacheThrottle(throttle time.Duration) {
	t.throttle = throttle
}

func (t *throttledTokenCache) Get(
	ctx context.Context,
) (BearerToken, error) {
	if t.cachedToken.NotEmpty() && !t.cachedToken.IsExpired() && time.Now().Before(t.nextCheck) {
		return t.cachedToken, nil
	}
	return t.Refresh(ctx)
}

func (t *throttledTokenCache) Refresh(
	ctx context.Context,
) (BearerToken, error) {
	tok, err := t.cache.Get(ctx, t.name)
	if err != nil {
		return BearerToken{}, fmt.Errorf("get from cache: %w", err)
	}
	t.cachedToken = tok
	t.nextCheck = time.Now().Add(t.throttle) // reset next check time
	return t.cachedToken, nil
}

func (t *throttledTokenCache) Set(
	ctx context.Context,
	token BearerToken,
) error {
	if t.cachedToken.Equal(token) {
		return nil // no change, nothing to do
	}
	t.cachedToken = token
	err := t.cache.Set(ctx, t.name, token)
	if err != nil {
		return fmt.Errorf("set cache: %w", err)
	}
	t.nextCheck = time.Now().Add(t.throttle) // reset next check time
	return nil
}

func (t *throttledTokenCache) Remove(
	ctx context.Context,
) error {
	if t.cachedToken.Empty() {
		return nil // nothing to remove
	}
	if err := t.cache.Remove(ctx, t.name); err != nil {
		return fmt.Errorf("remove from cache: %w", err)
	}
	t.cachedToken = BearerToken{} // reset cached token
	t.nextCheck = time.Now()      // reset next check time
	return nil
}

func (t *throttledTokenCache) RemoveIfEqual(
	ctx context.Context,
	token BearerToken,
) error {
	if t.cachedToken.Empty() {
		return nil // nothing to remove
	}
	if err := t.cache.RemoveIfEqual(ctx, t.name, token); err != nil {
		return fmt.Errorf("remove from cache: %w", err)
	}
	t.cachedToken = BearerToken{} // reset cached token
	t.nextCheck = time.Now()      // reset next check time
	return nil
}

type pureFileCachedTokener struct {
	cache *throttledTokenCache
	name  string
}

const defaultFileCacheThrottle = 5 * time.Minute // default throttle duration
func NewPureFileCachedTokener(
	cacheFile string,
	name string,
	logger *slog.Logger,
) NamedTokener {
	return &pureFileCachedTokener{
		name:  name,
		cache: newThrottledTokenCache(cacheFile, defaultFileCacheThrottle, name, logger),
	}
}

func (p *pureFileCachedTokener) SetCacheThrottle(throttle time.Duration) {
	p.cache.SetCacheThrottle(throttle)
}

func (p *pureFileCachedTokener) Name() string {
	return p.name
}
func (p *pureFileCachedTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	tok, err := p.cache.Get(ctx)
	if err != nil {
		return BearerToken{}, fmt.Errorf("get from cache: %w", err)
	}
	return tok, nil
}
func (p *pureFileCachedTokener) HandleError(
	ctx context.Context, token BearerToken, checkErr error,
) error {
	if checkErr == nil {
		return nil // no error to handle, token is valid
	}
	cacheTok, err := p.cache.Refresh(ctx)
	if err != nil {
		return errors.Join(checkErr, fmt.Errorf("refresh cache: %w", err))
	}
	if !cacheTok.Empty() && !cacheTok.Equal(token) {
		return nil
	}
	return fmt.Errorf("token requires refresh using another tool: %w", checkErr)
}

var _ NamedTokener = (*RenewableFileCachedTokener)(nil)

type RenewableFileCachedTokener struct {
	cache        *throttledTokenCache
	tokener      NamedTokener
	safetyMargin time.Duration
	logger       *slog.Logger
	name         string
}

func NewFileCachedTokener(
	cacheFile string,
	tokener NamedTokener,
	safetyMargin time.Duration,
	logger *slog.Logger,
) *RenewableFileCachedTokener {
	name := tokener.Name()
	return &RenewableFileCachedTokener{
		cache: newThrottledTokenCache(
			cacheFile, defaultFileCacheThrottle, name, logger,
		),
		tokener:      tokener,
		safetyMargin: safetyMargin,
		logger: logger.With(
			slog.String("name", "renewable_file_cached_tokener"),
			slog.String("file_cache", cacheFile),
			slog.String("token_name", name),
			slog.Duration("safety_margin", safetyMargin),
		),

		name: name, // should not change after creation
	}
}

func (f *RenewableFileCachedTokener) SetCacheThrottle(throttle time.Duration) {
	f.cache.SetCacheThrottle(throttle)
}

func (f *RenewableFileCachedTokener) Name() string {
	return f.name
}
func (f *RenewableFileCachedTokener) Unwrap() BearerTokener {
	return f.tokener
}

func (f *RenewableFileCachedTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	tok, err := f.cache.Get(ctx)
	if err != nil {
		return BearerToken{}, fmt.Errorf("get from cache: %w", err)
	}
	if tok.NotEmpty() && !tok.ExpiresAt.IsZero() {
		if f.safetyMargin > 0 && tok.ExpiresAt.Before(time.Now().Add(f.safetyMargin)) {
			f.logger.DebugContext(
				ctx,
				"cached token is about to expire, refreshing",
				slog.String("token", tok.String()),
			)
			tok = BearerToken{}
		}
	}
	f.safetyMargin = 0 // safety margin did its primary run check
	if tok.NotEmpty() && !tok.IsExpired() {
		return tok, nil
	}
	f.logger.DebugContext(
		ctx,
		"cached token is empty or expired, requesting a new one",
	)

	token, err := f.tokener.BearerToken(ctx)
	if err != nil {
		return token, fmt.Errorf("get from wrapped tokener: %w", err)
	}
	if token.Empty() {
		f.logger.DebugContext(
			ctx,
			"received empty token from wrapped tokener, not saving to cache",
		)
		return token, nil
	}
	if err = f.cache.Set(ctx, token); err != nil {
		f.logger.ErrorContext(
			ctx,
			"failed to save token to cache",
			slog.String("token_name", f.name),
			slog.Any("error", err),
		)
	} else {
		f.logger.DebugContext(
			ctx,
			"new token saved to cache",
			slog.String("token", token.String()),
		)
	}
	return token, nil
}

func (f *RenewableFileCachedTokener) HandleError(
	ctx context.Context, token BearerToken, checkErr error,
) error {
	if checkErr == nil { // no error to handle, token is valid
		return nil
	}
	f.logger.DebugContext(
		ctx,
		"handling error for token",
		slog.String("token", token.String()),
		slog.String("error", checkErr.Error()),
	)
	cachedTok, err := f.cache.Refresh(ctx)
	if err != nil {
		return errors.Join(
			checkErr,
			fmt.Errorf("refresh cache: %w", err),
		)
	}
	if cachedTok.NotEmpty() && !cachedTok.Equal(token) {
		f.logger.InfoContext(
			ctx,
			"token refreshed from cache",
			slog.String("token_name", f.name),
		)
		return nil // new token is available in cache, no need to handle error
	}
	if cachedTok.Equal(token) {
		err = f.cache.RemoveIfEqual(ctx, token) // remove the token from cache as it is no longer valid
		if err != nil {
			return errors.Join(
				checkErr,
				fmt.Errorf("remove from cache: %w", err),
			)
		}
	}
	return f.tokener.HandleError(ctx, token, checkErr)
}

type AsynchronouslyRenewableFileCachedTokener struct {
	cache        *throttledTokenCache
	tokener      NamedTokener
	safetyMargin time.Duration
	logger       *slog.Logger
	name         string

	ticker *time.Ticker
	group  singleflight.Group
	mutex  sync.Mutex // mutex to protect access to fetchedChannel and contextsWaiting
}

const (
	retryTimeout = 1 * time.Second // default retry timeout
)

func NewAsynchronouslyRenewableFileCachedTokener(
	cacheFile string,
	tokener NamedTokener,
	safetyMargin time.Duration,
	logger *slog.Logger,
) *AsynchronouslyRenewableFileCachedTokener {
	name := tokener.Name()
	return &AsynchronouslyRenewableFileCachedTokener{
		cache: newThrottledTokenCache(
			cacheFile, defaultFileCacheThrottle, name, logger,
		),
		tokener:      tokener,
		safetyMargin: safetyMargin,
		logger:       logger,

		ticker: time.NewTicker(0), // stopped ticker, will be reset on first token fetch
		group:  singleflight.Group{},
		name:   name, // should not change after creation
		mutex:  sync.Mutex{},
	}
}

func (f *AsynchronouslyRenewableFileCachedTokener) SetCacheThrottle(throttle time.Duration) {
	f.cache.SetCacheThrottle(throttle)
}

func (f *AsynchronouslyRenewableFileCachedTokener) Name() string {
	return f.name
}
func (f *AsynchronouslyRenewableFileCachedTokener) Unwrap() BearerTokener {
	return f.tokener
}
func (f *AsynchronouslyRenewableFileCachedTokener) BearerToken(
	ctx context.Context,
) (BearerToken, error) {
	tok, err := f.cache.Get(ctx)
	if err != nil {
		return BearerToken{}, fmt.Errorf("get from cache: %w", err)
	}
	if tok.NotEmpty() && !tok.IsExpired() {
		return tok, nil
	}
	return f.requestToken(ctx)
}
func (f *AsynchronouslyRenewableFileCachedTokener) HandleError(
	ctx context.Context, token BearerToken, checkErr error,
) error {
	if checkErr == nil { // no error to handle, token is valid
		return nil
	}
	cachedTok, err := f.cache.Refresh(ctx)
	if err != nil {
		return errors.Join(
			checkErr,
			fmt.Errorf("refresh cache: %w", err),
		)
	}
	if cachedTok.NotEmpty() && !cachedTok.Equal(token) {
		if f.logger != nil {
			f.logger.InfoContext(
				ctx,
				"token refreshed from cache",
				slog.String("token_name", f.name),
			)
		}
		return nil // new token is available in cache, no need to handle error
	}
	if cachedTok.Equal(token) {
		err = f.cache.RemoveIfEqual(ctx, token) // remove the token from cache as it is no longer valid
		if err != nil {
			return errors.Join(
				checkErr,
				fmt.Errorf("remove from cache: %w", err),
			)
		}
	}
	return f.tokener.HandleError(ctx, token, checkErr)
}

func (f *AsynchronouslyRenewableFileCachedTokener) Run(ctx context.Context) error { //nolint:gocognit
	runningUnderlying := false
	for {
		select {
		case <-ctx.Done():
			f.logger.InfoContext(
				ctx,
				"background token refresh stopped as the context is done",
			)
			return ctx.Err()
		case <-f.ticker.C:
			token, err := f.cache.Get(ctx)
			if err != nil {
				f.logger.ErrorContext(
					ctx,
					"failed to get token from cache",
					slog.String("token_name", f.name),
					slog.Any("error", err),
				)
				continue
			}

			if token.Empty() {
				continue
			}

			if token.ExpiresAt.IsZero() {
				continue
			}

			timeLeft := time.Until(token.ExpiresAt.Add(-f.safetyMargin))
			if timeLeft > 0 {
				f.ticker.Reset(timeLeft)
				continue
			}

			f.logger.InfoContext(
				ctx,
				"token is about to expire; initiating background token refresh",
				slog.Time("expires_at", token.ExpiresAt),
			)

			// Deferred until the first underlying tokener fetch
			// If cached token is fine, no need to run the underlying tokener
			if !runningUnderlying {
				runningUnderlying = true
				go func() {
					internalErr, _ := RunAsyncTokener(ctx, f.tokener)
					if internalErr != nil && !errors.Is(
						internalErr, context.Canceled,
					) {
						f.logger.ErrorContext(
							ctx,
							"background token refresh failed",
							slog.Any("error", internalErr),
						)
					}
				}()
			}

			_, err = f.requestToken(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				f.logger.ErrorContext(
					ctx,
					"background token refresh failed",
					slog.Any("error", err),
					slog.Duration("next_attempt", retryTimeout),
				)
				f.ticker.Reset(retryTimeout)
			}
		}
	}
}

func (f *AsynchronouslyRenewableFileCachedTokener) requestToken(
	ctx context.Context,
) (BearerToken, error) {
	res, err, _ := f.group.Do("", func() (any, error) {
		var refreshAfter time.Duration

		now := time.Now()
		token, err := f.tokener.BearerToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("get from wrapped tokener: %w", err)
		}

		if !token.ExpiresAt.IsZero() {
			ttl := token.ExpiresAt.Sub(now)
			if ttl > 0 {
				refreshAfter = ttl - f.safetyMargin
			} else {
				f.logger.ErrorContext(
					ctx,
					"received already expired token",
					slog.Time("expires_at", token.ExpiresAt),
				)
			}
		}

		if err = f.cache.Set(ctx, token); err != nil {
			if f.logger != nil {
				f.logger.ErrorContext(
					ctx,
					"failed to save token to cache",
					slog.String("token_name", f.name),
					slog.Any("error", err),
				)
			} else {
				return BearerToken{}, fmt.Errorf("save to cache: %w", err)
			}
		}

		if refreshAfter > 0 {
			f.ticker.Reset(refreshAfter)
		}

		return token, nil
	})
	if err != nil {
		return BearerToken{}, err
	}

	return res.(BearerToken), nil //nolint:errcheck // ok to panic
}

func GetDefaultCacheFileName() (string, error) {
	home, err := config.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, config.DefaultConfigDir, cacheFileName), nil
}
