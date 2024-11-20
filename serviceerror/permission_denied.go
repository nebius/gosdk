package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*PermissionDenied)(nil)
)

type PermissionDenied struct {
	commonServiceError
	denied *common.PermissionDenied
}

func NewPermissionDenied(se *common.ServiceError, detail *common.PermissionDenied) *PermissionDenied {
	return &PermissionDenied{
		commonServiceError: commonServiceError{
			source: se,
		},
		denied: detail,
	}
}

func (e *PermissionDenied) ResourceID() string {
	return e.denied.GetResourceId()
}

func (e *PermissionDenied) Error() string {
	return fmt.Sprintf(
		"Permission denied %s: service %s, resource ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
	)
}
