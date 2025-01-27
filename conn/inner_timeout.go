package conn

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			if errors.Is(err, context.DeadlineExceeded) || status.Code(err) == codes.DeadlineExceeded {
				// Return TimeoutError if timedCtx is done but ctx is not done
				select {
				case <-timedCtx.Done():
					select {
					case <-ctx.Done():
						// This is not timeout caused by this interceptor if the original context is done
					default:
						return &TimeoutError{wrapped: err}
					}
				default:
				}
			}
			// Otherwise, return the original error
			return err
		}
		return nil
	}
}
