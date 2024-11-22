package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nebius/gosdk/serviceerror"
)

// Disable is a [grpc.CallOption] that disables authentication on the client for a specific request.
type Disable struct {
	grpc.EmptyCallOption
}

// The following selectors can be used in common scenarios.
//
//nolint:gochecknoglobals // const
var /* const */ (
	Base      = Select("base")
	Propagate = Select("propagate")
)

// Selector is a [grpc.CallOption] to select one of configured authenticators for a specific request.
type Selector struct {
	grpc.EmptyCallOption
	Name string
}

// Select returns [Selector] to select one of configured authenticators for a specific request.
func Select(name string) Selector {
	return Selector{
		Name: name,
	}
}

// Error wraps any authentication error.
// It doesn't have an Unwrap method to avoid bugs.
// To get an underlying error you should do the following:
//
//	var authErr *auth.Error
//	if errors.As(err, &authErr) {
//		underlyingErr := authErr.Cause
//	}
type Error struct {
	Cause error
}

func newError(err error) *Error {
	return &Error{
		Cause: serviceerror.FromError(err),
	}
}

func (e *Error) Error() string {
	return e.Cause.Error()
}

func UnaryClientInterceptor(
	logger *slog.Logger,
	auths map[Selector]Authenticator,
	explicit bool,
) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		auth, err := selectAuth(ctx, logger, auths, explicit, method, opts)
		if err != nil {
			return newError(err)
		}

		if auth == nil {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		ctxAuth, err := auth.Auth(ctx)
		if err != nil {
			return newError(fmt.Errorf("get auth data on the client: %w", err))
		}

		err = invoker(ctxAuth, method, req, reply, cc, opts...)

		if status.Code(err) == codes.Unauthenticated {
			errX := auth.HandleError(ctxAuth, err)
			if errX == err { //nolint:errorlint // don't use errors.Is as we need an exact match
				// In this case, we should not wrap an error, see [Authenticator.HandleError] docs
				return errX
			}
			if errX != nil {
				return newError(fmt.Errorf("handle error: %w", errX))
			}

			newCtxAuth, errX := auth.Auth(ctx)
			if errX != nil {
				return newError(fmt.Errorf("get auth data on the client (second attempt): %w", errX))
			}

			logger.InfoContext(
				ctx,
				"retry grpc call with new credentials",
				slog.Any("previous_error", err),
				slog.String("method", method),
			)

			err = invoker(newCtxAuth, method, req, reply, cc, opts...)
		}

		return err
	}
}

func StreamClientInterceptor(
	logger *slog.Logger,
	auths map[Selector]Authenticator,
	explicit bool,
) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		auth, err := selectAuth(ctx, logger, auths, explicit, method, opts)
		if err != nil {
			return nil, newError(err)
		}

		if auth == nil {
			return streamer(ctx, desc, cc, method, opts...)
		}

		ctxAuth, err := auth.Auth(ctx)
		if err != nil {
			return nil, newError(fmt.Errorf("get auth data on the client: %w", err))
		}

		stream, err := streamer(ctxAuth, desc, cc, method, opts...)

		if status.Code(err) == codes.Unauthenticated {
			errX := auth.HandleError(ctxAuth, err)
			if errX == err { //nolint:errorlint // don't use errors.Is as we need an exact match
				// In this case, we should not wrap an error, see [Authenticator.HandleError] docs
				return nil, errX
			}
			if errX != nil {
				return nil, newError(fmt.Errorf("handle error: %w", errX))
			}

			newCtxAuth, errX := auth.Auth(ctx)
			if errX != nil {
				return nil, newError(fmt.Errorf("get auth data on the client (second attempt): %w", errX))
			}

			logger.InfoContext(
				ctx,
				"retry grpc stream with new credentials",
				slog.Any("previous_error", err),
				slog.String("method", method),
			)

			stream, err = streamer(newCtxAuth, desc, cc, method, opts...)
		}

		return stream, err
	}
}

func selectAuth( //nolint:gocognit
	ctx context.Context,
	logger *slog.Logger,
	auths map[Selector]Authenticator,
	explicit bool,
	method string,
	opts []grpc.CallOption,
) (Authenticator, error) {
	var auth Authenticator
	var selected bool
	var name string
	for _, opt := range opts {
		switch o := opt.(type) {
		case Disable:
			logger.DebugContext(ctx, "authenticator disabled", slog.String("method", method))
			return nil, nil //nolint:nilnil // don't want to introduce a sentinel error in private method
		case Selector:
			a, found := auths[o]
			if !found {
				names := make([]string, 0, len(auths))
				for selector := range auths {
					names = append(names, selector.Name)
				}
				return nil, fmt.Errorf(
					"unknown authenticator %q, known names are [%s]",
					o.Name,
					strings.Join(names, ", "),
				)
			}
			if selected {
				return nil, fmt.Errorf("ambigious selectors %s and %s", name, o.Name)
			}
			auth = a
			selected = true
			name = o.Name
		}
	}
	if !selected {
		if !explicit && len(auths) == 1 {
			for _, singleAuth := range auths {
				return singleAuth, nil
			}
		} else {
			return nil, errors.New("missing auth selector call option")
		}
	}
	logger.DebugContext(
		ctx,
		"authenticator selected",
		slog.String("method", method),
		slog.String("name", name),
	)
	return auth, nil
}
