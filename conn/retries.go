package conn

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"golang.org/x/net/http2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	"github.com/nebius/gosdk/serviceerror"
)

// DefaultRetriableCodes is a set of well known types gRPC codes that should be retriable.
//
// `ResourceExhausted` means that the user quota, e.g. per-RPC limits, have been reached.
// `Unavailable` means that system is currently unavailable and the client should retry again.
//
//nolint:gochecknoglobals // const
var DefaultRetriableCodes = []codes.Code{codes.ResourceExhausted, codes.Unavailable}

func IsRetriableNebiusError(err error) bool { //nolint:gocognit
	if err == nil {
		return false
	}

	// Do not retry on context errors or if the error is definitively unrecoverable
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	var serr *serviceerror.Error
	if errors.As(err, &serr); serr != nil {
		for _, detail := range serr.Details {
			// We were specifically asked to retry this call
			if detail.RetryType() == common.ServiceError_CALL {
				return true
			}
			// We were specifically asked to not retry this particular call
			if detail.RetryType() == common.ServiceError_NOTHING ||
				detail.RetryType() == common.ServiceError_UNIT_OF_WORK {
				return false
			}
			// In any other case, we have to decide based on other statement, eg
			// grpc code.
		}
	}
	if status, ok := status.FromError(err); ok {
		errCode := status.Code()
		for _, code := range DefaultRetriableCodes {
			if code == errCode {
				return true
			}
		}
		return false
	}

	// Retry on transport, TLS, and network-related errors
	if isNetworkError(err) || isTransportError(err) || isDNSError(err) {
		return true
	}

	return false
}

func WithServiceRetry() retry.CallOption {
	return retry.WithRetriable(IsRetriableNebiusError)
}

// Helper function to check for DNS errors.
func isDNSError(err error) bool {
	// Check for DNS resolution failures
	return strings.Contains(err.Error(), "name resolution")
}

// Helper function to check for common network errors.
func isNetworkError(err error) bool {
	var netErr net.Error
	if errors.As(err, &netErr) {
		return netErr.Timeout()
	}
	return false
}

// Helper function to check for transport-related errors.
func isTransportError(err error) bool {
	// Check for connection reset or refused errors
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		if opErr.Op == "dial" {
			if opErr.Err.Error() == "connection refused" || opErr.Err.Error() == "connection reset" {
				return true
			}
		}
	}

	// Check for gRPC stream errors
	var streamErr http2.StreamError
	return errors.As(err, &streamErr)
}
