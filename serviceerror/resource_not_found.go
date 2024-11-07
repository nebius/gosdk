package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*ResourceNotFound)(nil)
)

type ResourceNotFound struct {
	commonServiceError
	notFound *common.ResourceNotFound
}

func NewResourceNotFound(se *common.ServiceError, detail *common.ResourceNotFound) *ResourceNotFound {
	return &ResourceNotFound{
		commonServiceError: commonServiceError{
			source: se,
		},
		notFound: detail,
	}
}

func (e *ResourceNotFound) ResourceID() string {
	return e.notFound.ResourceId
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf(
		"Resource not found %s: service %s, resource ID: %s",
		e.source.Code,
		e.source.Service,
		e.ResourceID(),
	)
}
