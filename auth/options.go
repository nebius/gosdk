package auth

import "log/slog"

// Option configures one or more BearerTokener implementations.
//
// Options are applied on a best-effort basis: an option may affect only the
// concrete tokeners that recognize it and will be ignored by others.
type Option interface {
	apply(t BearerTokener)
}

type optionFunc func(t BearerTokener)

func (f optionFunc) apply(t BearerTokener) {
	f(t)
}

func applyOptions(t BearerTokener, opts ...Option) {
	for _, opt := range opts {
		opt.apply(t)
	}
}

func applyToTokenerOrWrapped(t BearerTokener, apply func(BearerTokener) bool) {
	for t != nil {
		if apply(t) {
			return
		}
		wrapper, ok := t.(Wrapper)
		if !ok {
			return
		}
		t = wrapper.Unwrap()
	}
}

type LoggerSetter interface {
	SetLogger(logger *slog.Logger)
}

// LoggerOption propagates the logger to tokeners that support logging.
type LoggerOption struct {
	Logger *slog.Logger
}

func (o LoggerOption) apply(t BearerTokener) {
	applyToTokenerOrWrapped(t, func(t BearerTokener) bool {
		setter, ok := t.(LoggerSetter)
		if ok {
			setter.SetLogger(o.Logger)
		}
		return ok
	})
}

// WithLogger sets the logger for the token provider. It will be used by both
// CachedServiceTokener and FileCacheTokener if they are used.
// It is also used in the federation Tokener if it is used.
// If the logger is not set, usually a no-op logger will be used.
func WithLogger(logger *slog.Logger) Option {
	return LoggerOption{Logger: logger}
}
