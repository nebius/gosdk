package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gofrs/flock"
)

const syncLockFilePrefix = "tokener-sync-"

// InAppSyncTokener is a [BearerTokener] decorator that synchronizes
// requests to [BearerTokener.BearerToken] within a single process.
type InAppSyncTokener struct {
	tokener BearerTokener
	mu      sync.Mutex
}

var _ BearerTokener = (*InAppSyncTokener)(nil)
var _ Wrapper = (*InAppSyncTokener)(nil)

// NewInAppSyncTokener returns a decorated [BearerTokener] that synchronizes
// requests to [BearerTokener.BearerToken] within a single process.
func NewInAppSyncTokener(tokener BearerTokener) *InAppSyncTokener {
	return &InAppSyncTokener{
		tokener: tokener,
		mu:      sync.Mutex{},
	}
}

func (t *InAppSyncTokener) Unwrap() BearerTokener {
	return t.tokener
}

func (t *InAppSyncTokener) BearerToken(ctx context.Context) (BearerToken, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.tokener.BearerToken(ctx)
}

func (t *InAppSyncTokener) HandleError(
	ctx context.Context,
	token BearerToken,
	err error,
) error {
	return t.tokener.HandleError(ctx, token, err)
}

// MultiprocessSyncTokener is a [BearerTokener] decorator that synchronizes
// requests to [BearerTokener.BearerToken] across multiple processes using
// a lock file in the default cache directory.
type MultiprocessSyncTokener struct {
	tokener  BearerTokener
	lock     *flock.Flock
	lockPath string
}

var _ BearerTokener = (*MultiprocessSyncTokener)(nil)
var _ Wrapper = (*MultiprocessSyncTokener)(nil)

// NewMultiprocessSyncTokener returns a decorated [BearerTokener] that synchronizes
// requests to [BearerTokener.BearerToken] across multiple processes using
// a lock file in the default cache directory.
func NewMultiprocessSyncTokener(tokener BearerTokener) (*MultiprocessSyncTokener, error) {
	lockPath, err := lockFilePathForTokener(tokener)
	if err != nil {
		return nil, err
	}
	return &MultiprocessSyncTokener{
		tokener:  tokener,
		lock:     flock.New(lockPath),
		lockPath: lockPath,
	}, nil
}

func (t *MultiprocessSyncTokener) Unwrap() BearerTokener {
	return t.tokener
}

func (t *MultiprocessSyncTokener) BearerToken(
	ctx context.Context,
) (_ BearerToken, deferError error) {
	if err := t.ensureLockDir(); err != nil {
		return BearerToken{}, err
	}
	if err := t.lock.Lock(); err != nil {
		return BearerToken{}, fmt.Errorf("lock tokener sync: %w", err)
	}
	defer func() {
		if err := t.lock.Unlock(); err != nil {
			deferError = errors.Join(
				deferError,
				fmt.Errorf("unlock tokener sync: %w", err),
			)
		}
	}()

	return t.tokener.BearerToken(ctx)
}

func (t *MultiprocessSyncTokener) HandleError(
	ctx context.Context,
	token BearerToken,
	err error,
) error {
	return t.tokener.HandleError(ctx, token, err)
}

func (t *MultiprocessSyncTokener) ensureLockDir() error {
	if err := os.MkdirAll(filepath.Dir(t.lockPath), 0750); err != nil {
		return fmt.Errorf("create sync lock dir: %w", err)
	}
	return nil
}

func lockFilePathForTokener(tokener BearerTokener) (string, error) {
	cacheFile, err := GetDefaultCacheFileName()
	if err != nil {
		return "", fmt.Errorf("get default cache file name: %w", err)
	}
	dir := filepath.Dir(cacheFile)
	name := tokenerLockName(tokener)
	hash := sha256.Sum256([]byte(name))
	return filepath.Join(dir, syncLockFilePrefix+hex.EncodeToString(hash[:])+".lock"), nil
}

func tokenerLockName(tokener BearerTokener) string {
	if name, ok := NameOfTokener(tokener); ok && name != "" {
		return name
	}
	return fmt.Sprintf("%T", tokener)
}
