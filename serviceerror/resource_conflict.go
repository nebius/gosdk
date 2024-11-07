package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*ResourceConflict)(nil)
)

type ResourceConflict struct {
	commonServiceError
	conflict *common.ResourceConflict
}

func NewResourceConflict(se *common.ServiceError, detail *common.ResourceConflict) *ResourceConflict {
	return &ResourceConflict{
		commonServiceError: commonServiceError{
			source: se,
		},
		conflict: detail,
	}
}

func (e *ResourceConflict) ResourceID() string {
	return e.conflict.ResourceId
}

func (e *ResourceConflict) Message() string {
	return e.conflict.Message
}

func (e *ResourceConflict) Error() string {
	return fmt.Sprintf(
		"Resource conflict %s: service %s, resource ID %s: %s",
		e.source.Code,
		e.source.Service,
		e.ResourceID(),
		e.Message(),
	)
}
