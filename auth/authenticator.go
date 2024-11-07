package auth

import (
	"context"
)

// Authenticator is used in gRPC interceptors.
type Authenticator interface {
	// Auth returns a [context.Context] with necessary grpc metadata for authorization.
	Auth(context.Context) (context.Context, error)

	// HandleError is called with a [context.Context] received from [Authenticator.Auth]
	// and an error from a gRPC call if it has the Unauthenticated code.
	// If HandleError returns nil, a new auth will be requested to retry the gRPC call.
	// If the gRPC call should not be retried, HandleError must return the incoming error.
	HandleError(context.Context, error) error
}
