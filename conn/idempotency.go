package conn

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const IdempotencyKeyHeader = "X-Idempotency-Key"

type withIdempotencyKey struct {
	grpc.EmptyCallOption

	key string
}

// WithIdempotencyKey returns [grpc.CallOption] to provide a custom idempotency key for a specific request.
func WithIdempotencyKey(key string) grpc.CallOption {
	return withIdempotencyKey{key: key}
}

// AddCustomIdempotencyKeyToOutgoingContext adds custom idempotency key to the outgoing gRPC metadata.
// It overrides an existing value if any.
func AddCustomIdempotencyKeyToOutgoingContext(ctx context.Context, key string) context.Context {
	md, exists := metadata.FromOutgoingContext(ctx)
	if !exists {
		md = metadata.MD{}
	}
	md.Set(IdempotencyKeyHeader, key)
	return metadata.NewOutgoingContext(ctx, md)
}

// AddNewIdempotencyKeyToOutgoingContext adds a new random idempotency key to the outgoing gRPC metadata.
// It overrides an existing value if any.
func AddNewIdempotencyKeyToOutgoingContext(ctx context.Context) (context.Context, error) {
	key, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}
	return AddCustomIdempotencyKeyToOutgoingContext(ctx, key.String()), nil
}

// EnsureIdempotencyKeyInOutgoingContext adds a new random idempotency key to the outgoing gRPC metadata if necessary.
// It does nothing if the metadata already contains an idempotency key.
func EnsureIdempotencyKeyInOutgoingContext(ctx context.Context) (context.Context, error) {
	md, _ := metadata.FromOutgoingContext(ctx)
	if len(md.Get(IdempotencyKeyHeader)) != 0 {
		return ctx, nil
	}
	return AddNewIdempotencyKeyToOutgoingContext(ctx)
}

// IdempotencyKeyInterceptor is a unary gRPC client interceptor that adds idempotency key to the outgoing metadata.
// An idempotency key from [WithIdempotencyKey] call option is used if provided.
// A new random idempotency key is generated otherwise.
//
// The SDK uses the interceptor by default, you don't need to do it by yourself.
func IdempotencyKeyInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	var idempotencyKey string
	for _, opt := range opts {
		ik, isIK := opt.(withIdempotencyKey)
		if isIK && ik.key != "" {
			idempotencyKey = ik.key
		}
	}

	if idempotencyKey != "" {
		ctx = AddCustomIdempotencyKeyToOutgoingContext(ctx, idempotencyKey)
	} else {
		var err error
		ctx, err = EnsureIdempotencyKeyInOutgoingContext(ctx)
		if err != nil {
			return fmt.Errorf("idempotency key: %w", err)
		}
	}

	return invoker(ctx, method, req, reply, cc, opts...)
}
