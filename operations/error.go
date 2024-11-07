package operations

import (
	"fmt"

	"google.golang.org/grpc/codes"

	"github.com/nebius/gosdk/serviceerror"
)

type Error struct {
	Operation Operation
	Code      codes.Code
	Message   string
}

func NewError(operation Operation) error {
	if operation.Successful() {
		return nil
	}
	s := operation.Status()
	return serviceerror.FromErrorAndAny(&Error{
		Operation: operation,
		Code:      s.Code(),
		Message:   s.Message(),
	}, s.Proto().GetDetails())
}

func (e *Error) Error() string {
	return fmt.Sprintf("operation failed [id=%s, code=%d]: %s", e.Operation.ID(), e.Code, e.Message)
}
