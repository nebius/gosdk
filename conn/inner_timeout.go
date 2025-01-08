package conn

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

type TimeoutError struct {
	wrapped error
}

func (t *TimeoutError) Error() string {
	return fmt.Sprintf("gosdk request timeout: %s", t.wrapped.Error())
}

func (t *TimeoutError) Unwrap() error {
	return t.wrapped
}

// UnaryClientTimeoutInterceptor returns a new unary client interceptor that sets a timeout on the request context.
func UnaryClientTimeoutInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		timedCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		err := invoker(timedCtx, method, req, reply, cc, opts...)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				select {
				case <-ctx.Done():
					return err
				default:
					return &TimeoutError{wrapped: err}
				}
			}
			// Otherwise, return the original error
			return err
		}
		return nil
	}
}
