package conn

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/nebius/gosdk/fieldmask/grpcheader"
)

const IdempotencyKeyHeader = "X-Idempotency-Key"

type IdempotencyKey string

func NewIdempotencyKey() (IdempotencyKey, error) {
	keyUUID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("generate key: %w", err)
	}
	return IdempotencyKey(keyUUID.String()), nil
}

func AddCustomIdempotencyKeyToOutgoingContext(ctx context.Context, key IdempotencyKey) context.Context {
	return metadata.AppendToOutgoingContext(ctx, IdempotencyKeyHeader, string(key))
}
func AddNewIdempotencyKeyToOutgoingContext(ctx context.Context) (context.Context, error) {
	ik, err := NewIdempotencyKey()
	if err != nil {
		return ctx, err
	}
	return AddCustomIdempotencyKeyToOutgoingContext(ctx, ik), nil
}

func EnsureIdempotencyKeyInOutgoingContext(ctx context.Context) (context.Context, error) {
	if grpcheader.IsInOutgoingContext(ctx, IdempotencyKeyHeader) {
		return ctx, nil
	}
	return AddNewIdempotencyKeyToOutgoingContext(ctx)
}

func IdempotencyKeyInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx, err := EnsureIdempotencyKeyInOutgoingContext(ctx)
	if err != nil {
		return fmt.Errorf("idempotency key: %w", err)
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}
