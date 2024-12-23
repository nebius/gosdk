package gosdk

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/nebius/gosdk/conn"
)

// Option is an interface combining all options for [New] func.
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

// WithLoggingOptions allows you to specify additional configurations for the grpc-ecosystem logging interceptor.
func WithLoggingOptions(opts ...logging.Option) Option {
	return optionLoggingOptions(opts)
}

// WithDialOptions allows you to specify additional grpc dial options for all services.
// You can use [conn.WithAddressDialOptions] to use different options for a specific [conn.Address].
func WithDialOptions(opts ...grpc.DialOption) Option {
	return optionDialOpts(opts)
}

// WithResolvers customizes service address resolution.
func WithResolvers(resolvers ...conn.Resolver) Option {
	return optionResolvers(resolvers)
}

// WithDomain changes the default "api.eu.nebius.cloud:443" domain.
func WithDomain(domain string) Option {
	return optionDomain(domain)
}

// WithAddressTemplate customizes service address resolution.
func WithAddressTemplate(find string, replace string) Option {
	return optionAddressTemplate{
		find:    find,
		replace: replace,
	}
}

// WithExplicitInit alters the behavior of the [New] constructor.
// If explicitInit is false (the default), [SDK.Init] is automatically called by [New].
// If explicitInit is true, you must manually call [SDK.Init] yourself.
// This option is useful for separating the [SDK] instantiation from I/O operations.
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
