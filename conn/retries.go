package conn

import (
	"context"
	"errors"
	"net"
	"slices"
	"strconv"
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

const unexpectedHTTPStatusPrefix = "unexpected HTTP status code received from server: "

func IsRetriableNebiusError(err error) bool {
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

	if s, isGRPC := status.FromError(err); isGRPC {
		if slices.Contains(DefaultRetriableCodes, s.Code()) {
			return true
		}

		// grpc-go maps unknown HTTP status codes from proxies, such as
		// Cloudflare 52x responses, to codes.Unknown.
		return s.Code() == codes.Unknown && hasUnexpectedHTTP52xStatus(s.Message())
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

func hasUnexpectedHTTP52xStatus(message string) bool {
	statusIndex := strings.Index(message, unexpectedHTTPStatusPrefix)
	if statusIndex < 0 {
		return false
	}

	statusStart := statusIndex + len(unexpectedHTTPStatusPrefix)
	statusEnd := statusStart
	for statusEnd < len(message) && message[statusEnd] >= '0' && message[statusEnd] <= '9' {
		statusEnd++
	}
	if statusStart == statusEnd {
		return false
	}

	code, err := strconv.Atoi(message[statusStart:statusEnd])
	return err == nil && code >= 520 && code < 530
}
