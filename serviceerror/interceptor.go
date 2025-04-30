package serviceerror

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply any,
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	md := metadata.MD{}
	opts = append(opts, grpc.Header(&md))
	err := invoker(ctx, method, req, reply, cc, opts...)
	return FromRPCError(err, &md)
}

func StreamClientInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	md := metadata.MD{}
	opts = append(opts, grpc.Header(&md))
	stream, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, FromRPCError(err, &md)
	}
	return wrappedStream{
		md:           &md,
		ClientStream: stream,
	}, nil
}

type wrappedStream struct {
	grpc.ClientStream
	md *metadata.MD
}

func (w wrappedStream) SendMsg(m any) error {
	return FromRPCError(w.ClientStream.SendMsg(m), w.md)
}

func (w wrappedStream) RecvMsg(m any) error {
	return FromRPCError(w.ClientStream.RecvMsg(m), w.md)
}
