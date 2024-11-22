package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*BadResourceState)(nil)

type BadResourceState struct {
	commonDetail
	badResourceState *common.BadResourceState
}

func newBadResourceState(se *common.ServiceError, br *common.BadResourceState) *BadResourceState {
	return &BadResourceState{
		commonDetail: commonDetail{
			source: se,
		},
		badResourceState: br,
	}
}

func (e *BadResourceState) ResourceID() string {
	return e.badResourceState.GetResourceId()
}

func (e *BadResourceState) Message() string {
	return e.badResourceState.GetMessage()
}

func (e *BadResourceState) String() string {
	return fmt.Sprintf(
		"Bad resource state %s:\n  service %s:\n  resource ID: %s\n  message: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
		e.Message(),
	)
}
