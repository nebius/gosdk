package gosdk

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"

	"github.com/nebius/gosdk/conn"
)

type Option interface {
	option()
}

// WithCredentials enables client authentication. By default, [NoCredentials] is used.
//
// [Credentials] are applied under the following conditions:
//   - The outgoing gRPC metadata does not already include authorization information.
//   - The list of gRPC call options does not contain [github.com/nebius/gosdk/auth.DisableAuth].
//
// If either of these conditions is not met, the provided credentials will not be used.
func WithCredentials(creds Credentials) Option {
	return optionCredentials{creds: creds}
}

// WithLogger enables logging in the SDK. By default [NoopHandler] is used.
func WithLogger(logger *slog.Logger) Option {
	return WithLoggerHandler(logger.Handler())
}

// WithLoggerHandler enables logging in the SDK. By default [NoopHandler] is used.
func WithLoggerHandler(handler slog.Handler) Option {
	return optionLogger{handler: handler}
}

func WithLoggingOptions(opts ...logging.Option) Option {
	return optionLoggingOptions(opts)
}

func WithTracingOptions(opts ...otelgrpc.Option) Option {
	return optionTracingOptions(opts)
}

func WithDialOptions(opts ...grpc.DialOption) Option {
	return optionDialOpts(opts)
}

func WithResolvers(resolvers ...conn.Resolver) Option {
	return optionResolvers(resolvers)
}

func WithDomain(domain string) Option {
	return optionDomain(domain)
}

func WithAddressTemplate(find string, replace string) Option {
	return optionAddressTemplate{
		find:    find,
		replace: replace,
	}
}

// WithExplicitInit changes the [New] constructor behavior.
// If explicitInit is false (default), [SDK.Init] is called by [New].
// Otherwise, you must call [SDK.Init] by yourself.
// This option may be useful to separate the [SDK] instantiation from IO operations.
func WithExplicitInit(explicitInit bool) Option {
	return optionExplicitInit(explicitInit)
}

// WithInit adds an extra fn, which will be called on init [SDK].
func WithInit(fn func(context.Context, *SDK) error) Option {
	return optionInit(fn)
}

type (
	optionCredentials     struct{ creds Credentials }
	optionLogger          struct{ handler slog.Handler }
	optionLoggingOptions  []logging.Option
	optionTracingOptions  []otelgrpc.Option
	optionDialOpts        []grpc.DialOption
	optionResolvers       []conn.Resolver
	optionDomain          string
	optionAddressTemplate struct {
		find    string
		replace string
	}
	optionExplicitInit bool
	optionInit         func(context.Context, *SDK) error
)

func (optionCredentials) option()     {}
func (optionLogger) option()          {}
func (optionLoggingOptions) option()  {}
func (optionTracingOptions) option()  {}
func (optionDialOpts) option()        {}
func (optionResolvers) option()       {}
func (optionDomain) option()          {}
func (optionAddressTemplate) option() {}
func (optionExplicitInit) option()    {}
func (optionInit) option()            {}

type NoopHandler struct{}

func (NoopHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (NoopHandler) Handle(context.Context, slog.Record) error { return nil }
func (h NoopHandler) WithAttrs([]slog.Attr) slog.Handler      { return h }
func (h NoopHandler) WithGroup(string) slog.Handler           { return h }

func TracingStatsHandler(opts ...otelgrpc.Option) grpc.DialOption {
	options := []otelgrpc.Option{
		otelgrpc.WithPropagators(propagation.TraceContext{}),
	}
	options = append(options, opts...)
	return grpc.WithStatsHandler(otelgrpc.NewClientHandler(options...))
}

type slogAdapter struct {
	log *slog.Logger
}

func (a slogAdapter) Log(ctx context.Context, level logging.Level, msg string, fields ...any) {
	a.log.Log(ctx, slog.Level(level), msg, fields...)
}

func unaryLoggingInterceptor(log *slog.Logger, opts ...logging.Option) grpc.UnaryClientInterceptor {
	return logging.UnaryClientInterceptor(slogAdapter{log: log}, opts...)
}

func streamLoggingInterceptor(log *slog.Logger, opts ...logging.Option) grpc.StreamClientInterceptor {
	return logging.StreamClientInterceptor(slogAdapter{log: log}, opts...)
}
