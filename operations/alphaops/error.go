package alphaops

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

type Error struct {
	Operation *Operation
	Code      codes.Code
	Message   string
}

func NewError(operation *Operation) error {
	if operation.Successful() {
		return nil
	}
	s := operation.Status()
	return &Error{
		Operation: operation,
		Code:      s.Code(),
		Message:   s.Message(),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("operation failed [id=%s, code=%d]: %s", e.Operation.ID(), e.Code, e.Message)
}
