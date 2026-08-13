package alphaops

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
)

const (
	DefaultPollInterval        = 1 * time.Second
	DefaultPollErrorBackoffMax = 30 * time.Second
	defaultPollErrorJitter     = 0.2
)

type pollInterval struct {
	grpc.EmptyCallOption

	interval time.Duration
}

type pollErrorBackoff struct {
	grpc.EmptyCallOption

	backoff retry.BackoffFunc
}

// PollInterval returns [grpc.CallOption] you can pass to [Operation.Wait] to override the [DefaultPollInterval].
//
// If an interval isn't positive, the [DefaultPollInterval] is used.
func PollInterval(interval time.Duration) grpc.CallOption {
	return pollInterval{
		interval: interval,
	}
}

// PollErrorBackoff returns [grpc.CallOption] you can pass to [Operation.Wait]
// to override the default exponential backoff between consecutive retriable
// polling errors. The first retry uses [DefaultPollInterval], jittered by 20%,
// and the delay is capped at [DefaultPollErrorBackoffMax].
//
// The attempt passed to backoff starts at one and resets after a successful
// poll. Passing nil disables retries after polling errors.
func PollErrorBackoff(backoff retry.BackoffFunc) grpc.CallOption {
	return pollErrorBackoff{
		backoff: backoff,
	}
}

func defaultPollErrorBackoff() retry.BackoffFunc {
	backoff := retry.BackoffExponentialWithJitter(DefaultPollInterval, defaultPollErrorJitter)
	return func(ctx context.Context, attempt uint) time.Duration {
		return min(backoff(ctx, attempt), DefaultPollErrorBackoffMax)
	}
}

func resolveWaitOptions(opts []grpc.CallOption) (time.Duration, retry.BackoffFunc) {
	interval := DefaultPollInterval
	backoff := defaultPollErrorBackoff()
	for _, opt := range opts {
		switch option := opt.(type) {
		case pollInterval:
			if option.interval > 0 {
				interval = option.interval
			}
		case pollErrorBackoff:
			backoff = option.backoff
		}
	}
	return interval, backoff
}
