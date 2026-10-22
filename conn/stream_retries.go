package conn

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// StreamClientRetryInterceptor bounds each server-streaming attempt with its retry timeout.
// It accepts the same options as the gRPC retry interceptor. Client streams must disable retries.
// The timeout includes all responses. Set retry.WithPerRetryTimeout(0) to use only the caller deadline.
// Retries stop after the first delivered response.
func StreamClientRetryInterceptor(options ...retry.CallOption) grpc.StreamClientInterceptor {
	streamRetry := retry.StreamClientInterceptor(options...)
	attemptOptions := retry.UnaryClientInterceptor(options...)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string,
		streamer grpc.Streamer, opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		var currentStream grpc.ClientStream
		boundedStreamer := func(attemptCtx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
			method string, _ ...grpc.CallOption,
		) (grpc.ClientStream, error) {
			var stream grpc.ClientStream
			var streamErr error
			// The middleware resolves opaque retry options and creates deadlines only for unary calls.
			// Copy that deadline to a stream context that survives the unary interceptor's return.
			err := attemptOptions(
				attemptCtx,
				method,
				nil,
				nil,
				cc,
				func(callCtx context.Context, _ string, _, _ any, _ *grpc.ClientConn, callOpts ...grpc.CallOption) error {
					streamCtx := attemptCtx
					cancel := context.CancelFunc(func() {})
					if deadline, ok := callCtx.Deadline(); ok {
						if parentDeadline, bounded := attemptCtx.Deadline(); !bounded ||
							deadline.Before(parentDeadline) {
							streamCtx, cancel = context.WithDeadline(attemptCtx, deadline)
						}
					}
					stream, streamErr = streamer(streamCtx, desc, cc, method, callOpts...)
					if streamErr != nil {
						cancel()
					} else {
						stream = &deadlineClientStream{ClientStream: stream, cancel: cancel}
					}
					// The stream retry interceptor handles failures. Do not retry stream creation here.
					return nil
				},
				opts...)
			if err != nil {
				return nil, err
			}
			currentStream = stream
			return stream, streamErr
		}
		stream, err := streamRetry(ctx, desc, cc, method, boundedStreamer, opts...)
		if err != nil {
			return nil, err
		}
		return &responseClientStream{ClientStream: stream, current: &currentStream}, nil
	}
}

// Read from the successful native attempt to prevent response replay.
// RecvMsg follows the gRPC single-reader contract.
type responseClientStream struct {
	grpc.ClientStream
	current   *grpc.ClientStream
	committed grpc.ClientStream
}

func (s *responseClientStream) RecvMsg(message any) error {
	if s.committed != nil {
		return s.committed.RecvMsg(message)
	}
	err := s.ClientStream.RecvMsg(message)
	if err == nil {
		s.committed = *s.current
	}
	return err
}

type deadlineClientStream struct {
	grpc.ClientStream
	cancel context.CancelFunc
}

func (s *deadlineClientStream) RecvMsg(message any) error {
	err := s.ClientStream.RecvMsg(message)
	if err != nil {
		s.cancel()
	}
	return err
}

func (s *deadlineClientStream) SendMsg(message any) error {
	err := s.ClientStream.SendMsg(message)
	if err != nil {
		s.cancel()
	}
	return err
}

func (s *deadlineClientStream) Header() (metadata.MD, error) {
	header, err := s.ClientStream.Header()
	if err != nil {
		s.cancel()
	}
	return header, err
}

func (s *deadlineClientStream) CloseSend() error {
	err := s.ClientStream.CloseSend()
	if err != nil {
		s.cancel()
	}
	return err
}
