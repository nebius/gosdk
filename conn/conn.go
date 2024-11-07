package conn

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceID string
type Address string

type Resolver interface {
	// Resolve returns grpc service [Address] by [ServiceID].
	// It returns [*UnknownServiceError] if the service is unknown.
	Resolve(context.Context, ServiceID) (Address, error)
}

type Dialer interface {
	Dial(context.Context, Address) (grpc.ClientConnInterface, error)
	Close() error
}

type UnknownServiceError struct {
	ID ServiceID
}

func NewUnknownServiceError(id ServiceID) *UnknownServiceError {
	return &UnknownServiceError{
		ID: id,
	}
}

func (e *UnknownServiceError) Error() string {
	return fmt.Sprintf("unknown service %s", e.ID)
}
