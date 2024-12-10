package auth

import (
	"context"
)

// Authenticator is an interface for adding authentication data to the [context.Context] in gRPC interceptors.
// There are two main implementations: [AuthenticatorFromBearerTokener] and [PropagateAuthorizationHeader],
// but custom implementations can also be provided.
type Authenticator interface {
	// Auth enriches the provided [context.Context] with necessary gRPC metadata for authentication.
	Auth(context.Context) (context.Context, error)

	// HandleError processes errors from gRPC calls that fail with the Unauthenticated status code.
	// It receives the [context.Context] returned by [Authenticator.Auth] and the error from the gRPC call.
	// If HandleError returns nil, the system will attempt to re-authenticate and retry the call.
	// If the call should not be retried, HandleError should return the original error.
	HandleError(context.Context, error) error
}
