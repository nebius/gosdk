package auth

import "context"

type Runnable interface {
	Run(ctx context.Context) error
}

func RunAsyncTokener(ctx context.Context, tokener BearerTokener) (error, bool) {
	if namedTokener, ok := tokener.(Runnable); ok {
		return namedTokener.Run(ctx), true
	}
	if wrapper, ok := tokener.(Wrapper); ok {
		return RunAsyncTokener(ctx, wrapper.Unwrap())
	}
	return nil, false
}
