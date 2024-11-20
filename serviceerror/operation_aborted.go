package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*OperationAborted)(nil)
)

type OperationAborted struct {
	commonServiceError
	aborted *common.OperationAborted
}

func NewOperationAborted(se *common.ServiceError, detail *common.OperationAborted) *OperationAborted {
	return &OperationAborted{
		commonServiceError: commonServiceError{
			source: se,
		},
		aborted: detail,
	}
}

func (e *OperationAborted) ResourceID() string {
	return e.aborted.GetResourceId()
}
func (e *OperationAborted) AbortedOperationID() string {
	return e.aborted.GetOperationId()
}
func (e *OperationAborted) NewOperationID() string {
	return e.aborted.GetAbortedByOperationId()
}

func (e *OperationAborted) Error() string {
	return fmt.Sprintf(
		"Operation aborted %s: service %s, resource: %s, aborted operation ID: %s, new operation ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
		e.AbortedOperationID(),
		e.NewOperationID(),
	)
}
