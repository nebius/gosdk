package operations

import (
	"fmt"

	"google.golang.org/grpc/codes"

	"github.com/nebius/gosdk/serviceerror"
)

// Error is returned by [Operation.Wait] if the operation is failed.
type Error struct {
	Operation Operation
	Code      codes.Code
	Message   string
}

// NewError returns nil if the operation is successful and [*Error] otherwise.
func NewError(operation Operation) error {
	if operation.Successful() {
		return nil
	}
	s := operation.Status()
	err := &Error{
		Operation: operation,
		Code:      s.Code(),
		Message:   s.Message(),
	}
	return serviceerror.FromError(err, s.Proto().GetDetails()...)
}

func (e *Error) Error() string {
	return fmt.Sprintf("operation failed [id=%s, code=%d]: %s", e.Operation.ID(), e.Code, e.Message)
}
