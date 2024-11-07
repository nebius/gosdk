package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*BadResourceState)(nil)
)

type BadResourceState struct {
	commonServiceError
	badResouceState *common.BadResourceState
}

func NewBadResourceState(se *common.ServiceError, br *common.BadResourceState) *BadResourceState {
	return &BadResourceState{
		commonServiceError: commonServiceError{
			source: se,
		},
		badResouceState: br,
	}
}

func (e *BadResourceState) ResourceID() string {
	return e.badResouceState.ResourceId
}

func (e *BadResourceState) Message() string {
	return e.badResouceState.Message
}

func (e *BadResourceState) Error() string {
	return fmt.Sprintf(
		"Bad resource state %s:\n  service %s:\n  resource ID: %s\n  message: %s",
		e.source.Code,
		e.source.Service,
		e.ResourceID(),
		e.Message(),
	)
}
