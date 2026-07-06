package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*OperationConflict)(nil)

type OperationConflict struct {
	commonDetail
	conflict *common.OperationConflict
}

func newOperationConflict(se *common.ServiceError, detail *common.OperationConflict) *OperationConflict {
	return &OperationConflict{
		commonDetail: commonDetail{
			source: se,
		},
		conflict: detail,
	}
}

func (e *OperationConflict) ResourceID() string {
	return e.conflict.GetResourceId()
}

func (e *OperationConflict) ConflictingOperationID() string {
	return e.conflict.GetConflictingOperationId()
}

func (e *OperationConflict) String() string {
	return fmt.Sprintf(
		"Operation conflict %s: service %s, resource: %s, conflicting operation ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
		e.ConflictingOperationID(),
	)
}
