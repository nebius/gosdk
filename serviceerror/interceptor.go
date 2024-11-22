package serviceerror

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return FromError(err)
}

func StreamClientInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	stream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, FromError(err)
	}
	return wrappedStream{ClientStream: stream}, nil
}

type wrappedStream struct {
	grpc.ClientStream
}

func (w wrappedStream) SendMsg(m any) error {
	return FromError(w.ClientStream.SendMsg(m))
}

func (w wrappedStream) RecvMsg(m any) error {
	return FromError(w.ClientStream.RecvMsg(m))
}
