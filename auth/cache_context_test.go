package auth_test

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nebius/gosdk/auth"
)

func TestCachedServiceTokenerAlreadyCanceled(t *testing.T) {
	t.Parallel()
	underlying := newBlockingBearerTokener()
	cached := auth.NewCachedTokener(underlying)
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	_, err := cached.BearerToken(ctx)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("BearerToken() error = %v, want context.Canceled", err)
	}
	if calls := underlying.calls.Load(); calls != 0 {
		t.Fatalf("underlying calls = %d, want 0", calls)
	}
}

func TestCachedServiceTokenerWaitHonorsContext(t *testing.T) {
	t.Parallel()

	underlying := newBlockingBearerTokener()
	t.Cleanup(underlying.release)
	cached := auth.NewCachedTokener(underlying)
	leaderResult := make(chan tokenResult, 1)
	go func() {
		token, err := cached.BearerToken(t.Context())
		leaderResult <- tokenResult{token: token, err: err}
	}()
	<-underlying.started

	for _, test := range []struct {
		name string
		ctx  func() (context.Context, context.CancelFunc)
		want error
	}{
		{
			name: "canceled",
			ctx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(t.Context())
				cancel()
				return ctx, func() {}
			},
			want: context.Canceled,
		},
		{
			name: "deadline exceeded while waiting",
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(t.Context(), 25*time.Millisecond)
			},
			want: context.DeadlineExceeded,
		},
	} {
		ctx, cancel := test.ctx()
		result := make(chan error, 1)
		go func() {
			_, err := cached.BearerToken(ctx)
			result <- err
		}()

		select {
		case err := <-result:
			if !errors.Is(err, test.want) {
				t.Fatalf("%s: BearerToken() error = %v, want %v", test.name, err, test.want)
			}
		case <-time.After(time.Second):
			t.Fatalf("%s: BearerToken() did not stop waiting for the shared acquisition", test.name)
		}
		cancel()
	}

	if calls := underlying.calls.Load(); calls != 1 {
		t.Fatalf("underlying calls = %d, want 1", calls)
	}
	underlying.release()

	result := <-leaderResult
	if result.err != nil {
		t.Fatalf("leader BearerToken() error = %v", result.err)
	}
	if result.token.Token != "token" {
		t.Fatalf("leader token = %q, want token", result.token.Token)
	}

	cachedToken, err := cached.BearerToken(t.Context())
	if err != nil {
		t.Fatalf("cached BearerToken() error = %v", err)
	}
	if cachedToken.Token != "token" {
		t.Fatalf("cached token = %q, want token", cachedToken.Token)
	}
	if calls := underlying.calls.Load(); calls != 1 {
		t.Fatalf("underlying calls after cache hit = %d, want 1", calls)
	}
}

type tokenResult struct {
	token auth.BearerToken
	err   error
}

type blockingBearerTokener struct {
	started     chan struct{}
	releaseCh   chan struct{}
	startedOnce sync.Once
	releaseOnce sync.Once
	calls       atomic.Int32
}

func newBlockingBearerTokener() *blockingBearerTokener {
	return &blockingBearerTokener{
		started:   make(chan struct{}),
		releaseCh: make(chan struct{}),
	}
}

func (t *blockingBearerTokener) BearerToken(ctx context.Context) (auth.BearerToken, error) {
	t.calls.Add(1)
	t.startedOnce.Do(func() { close(t.started) })
	select {
	case <-ctx.Done():
		return auth.BearerToken{}, ctx.Err()
	case <-t.releaseCh:
	}
	if err := ctx.Err(); err != nil {
		return auth.BearerToken{}, err
	}
	return auth.BearerToken{Token: "token"}, nil
}

func (*blockingBearerTokener) HandleError(_ context.Context, _ auth.BearerToken, err error) error {
	return err
}

func (t *blockingBearerTokener) release() {
	t.releaseOnce.Do(func() { close(t.releaseCh) })
}

func TestCachedServiceTokenerLeaderCancellation(t *testing.T) {
	t.Parallel()
	underlying := newBlockingBearerTokener()
	t.Cleanup(underlying.release)
	cached := auth.NewCachedTokener(underlying)
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	leader := make(chan error, 1)
	go func() {
		_, err := cached.BearerToken(ctx)
		leader <- err
	}()
	select {
	case <-underlying.started:
	case <-time.After(time.Second):
		t.Fatal("acquisition did not start")
	}
	cancel()
	select {
	case err := <-leader:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("leader error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("leader did not stop waiting")
	}
	underlying.release()
	waitCtx, waitCancel := context.WithTimeout(t.Context(), time.Second)
	defer waitCancel()
	token, err := cached.BearerToken(waitCtx)
	if err != nil || token.Token != "token" {
		t.Fatalf("remaining caller got (%v, %v), want token", token, err)
	}
	if calls := underlying.calls.Load(); calls != 1 {
		t.Fatalf("underlying calls = %d, want 1", calls)
	}
}

func TestCachedServiceTokenerAcquisitionTimeout(t *testing.T) {
	t.Parallel()
	underlying := newBlockingBearerTokener()
	t.Cleanup(underlying.release)
	cached := auth.NewCachedTokener(underlying, auth.WithCachedTokenerAcquireTimeout(25*time.Millisecond))
	ctx, cancel := context.WithTimeout(t.Context(), time.Second)
	defer cancel()
	_, err := cached.BearerToken(ctx)
	if !errors.Is(err, context.DeadlineExceeded) || ctx.Err() != nil {
		t.Fatalf("acquisition error = %v, caller error = %v", err, ctx.Err())
	}
	underlying.release()
	if _, retryErr := cached.BearerToken(ctx); retryErr != nil {
		t.Fatalf("retry after acquisition timeout: %v", retryErr)
	}
	if calls := underlying.calls.Load(); calls != 2 {
		t.Fatalf("underlying calls = %d, want 2", calls)
	}
}
