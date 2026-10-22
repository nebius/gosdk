package conn_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/nebius/gosdk/conn"
)

type attemptStream struct {
	grpc.ClientStream
	ctx      context.Context
	wait     bool
	response bool
}

func (s *attemptStream) Context() context.Context { return s.ctx }
func (s *attemptStream) SendMsg(any) error        { return nil }
func (s *attemptStream) CloseSend() error         { return nil }
func (s *attemptStream) RecvMsg(any) error {
	if s.response {
		s.response = false
		return nil
	}
	if s.wait {
		<-s.ctx.Done()
		return status.FromContextError(s.ctx.Err()).Err()
	}
	return io.EOF
}

func TestStreamRetryAttemptDeadline(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		overall      time.Duration
		perAttempt   time.Duration
		callOverride []grpc.CallOption
		expected     time.Duration
	}{
		{name: "attempt", overall: 5 * time.Second, perAttempt: time.Second, expected: time.Second},
		{name: "overall", overall: time.Second, perAttempt: 5 * time.Second, expected: time.Second},
		{name: "no overall", perAttempt: time.Second, expected: time.Second},
		{name: "unlimited"},
		{
			name:         "call override",
			perAttempt:   5 * time.Second,
			callOverride: []grpc.CallOption{retry.WithPerRetryTimeout(time.Second)},
			expected:     time.Second,
		},
		{name: "disabled", perAttempt: time.Second, callOverride: []grpc.CallOption{retry.Disable()}},
		{
			name:         "disabled with caller deadline",
			overall:      time.Second,
			perAttempt:   time.Second,
			callOverride: []grpc.CallOption{retry.Disable()},
			expected:     time.Second,
		},
		{name: "only caller deadline", overall: time.Second, expected: time.Second},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			if tc.overall > 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.overall)
				defer cancel()
			}
			interceptor := conn.StreamClientRetryInterceptor(retry.WithMax(2), retry.WithPerRetryTimeout(tc.perAttempt))
			var native *attemptStream
			stream, err := interceptor(
				ctx,
				&grpc.StreamDesc{ServerStreams: true},
				nil,
				"/test.Service/Watch",
				func(ctx context.Context, _ *grpc.StreamDesc, _ *grpc.ClientConn, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
					native = &attemptStream{ctx: ctx}
					return native, nil
				},
				tc.callOverride...)
			require.NoError(t, err)
			nativeCtx := native.Context()
			require.NoError(t, nativeCtx.Err(), "the deadline must survive interceptor return")
			deadline, bounded := nativeCtx.Deadline()
			assert.Equal(t, tc.expected > 0, bounded)
			if bounded {
				assert.InDelta(t, tc.expected.Seconds(), time.Until(deadline).Seconds(), 0.2)
			}
			if tc.overall > 0 && tc.expected == tc.overall {
				assert.Same(t, ctx, nativeCtx, "an inherited deadline needs no child context")
			}
			require.ErrorIs(t, stream.RecvMsg(nil), io.EOF)
			if nativeCtx != ctx {
				require.ErrorIs(t, nativeCtx.Err(), context.Canceled)
			} else {
				require.NoError(t, ctx.Err(), "the caller context must remain active")
			}
		})
	}
}

func TestStreamRetryAfterAttemptTimeout(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("x-audit", "original"))
	contexts := []context.Context{}
	interceptor := conn.StreamClientRetryInterceptor(retry.WithMax(2), retry.WithPerRetryTimeout(10*time.Millisecond),
		retry.WithBackoff(func(context.Context, uint) time.Duration { return 0 }))
	stream, err := interceptor(
		ctx,
		&grpc.StreamDesc{ServerStreams: true},
		nil,
		"/test.Service/Watch",
		func(ctx context.Context, _ *grpc.StreamDesc, _ *grpc.ClientConn, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
			contexts = append(contexts, ctx)
			return &attemptStream{ctx: ctx, wait: len(contexts) == 1}, nil
		},
	)
	require.NoError(t, err)
	require.NoError(t, stream.SendMsg("request"))
	require.NoError(t, stream.CloseSend())
	require.ErrorIs(t, stream.RecvMsg(nil), io.EOF)
	require.Len(t, contexts, 2)
	require.ErrorIs(t, contexts[0].Err(), context.DeadlineExceeded)
	require.NoError(t, ctx.Err())
	md, _ := metadata.FromOutgoingContext(contexts[1])
	assert.Equal(t, []string{"original"}, md.Get("x-audit"))
}

func TestStreamRetryDoesNotRepeatFailedCreationInsideAttempt(t *testing.T) {
	t.Parallel()
	attempts := 0
	interceptor := conn.StreamClientRetryInterceptor(retry.WithMax(2), retry.WithPerRetryTimeout(time.Second),
		retry.WithBackoff(func(context.Context, uint) time.Duration { return 0 }))
	_, err := interceptor(
		context.Background(),
		&grpc.StreamDesc{ServerStreams: true},
		nil,
		"/test.Service/Watch",
		func(context.Context, *grpc.StreamDesc, *grpc.ClientConn, string, ...grpc.CallOption) (grpc.ClientStream, error) {
			attempts++
			return nil, status.Error(codes.Unavailable, "transport failed")
		},
	)
	assert.Equal(t, codes.Unavailable, status.Code(err))
	assert.Equal(t, 2, attempts)
}

func TestStreamRetryDoesNotReplayDeliveredResponse(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	attempts := 0
	interceptor := conn.StreamClientRetryInterceptor(retry.WithMax(2), retry.WithPerRetryTimeout(10*time.Millisecond),
		retry.WithBackoff(func(context.Context, uint) time.Duration { return 0 }))
	stream, err := interceptor(
		ctx,
		&grpc.StreamDesc{ServerStreams: true},
		nil,
		"/test.Service/Watch",
		func(ctx context.Context, _ *grpc.StreamDesc, _ *grpc.ClientConn, _ string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
			attempts++
			return &attemptStream{ctx: ctx, response: true, wait: true}, nil
		},
	)
	require.NoError(t, err)
	require.NoError(t, stream.SendMsg("request"))
	require.NoError(t, stream.CloseSend())
	require.NoError(t, stream.RecvMsg(nil))
	assert.Equal(t, codes.DeadlineExceeded, status.Code(stream.RecvMsg(nil)))
	assert.Equal(t, 1, attempts)
	require.NoError(t, ctx.Err())
}
