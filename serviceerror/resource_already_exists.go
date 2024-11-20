package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*ResourceAlreadyExists)(nil)
)

type ResourceAlreadyExists struct {
	commonServiceError
	exists *common.ResourceAlreadyExists
}

func NewResourceAlreadyExists(se *common.ServiceError, detail *common.ResourceAlreadyExists) *ResourceAlreadyExists {
	return &ResourceAlreadyExists{
		commonServiceError: commonServiceError{
			source: se,
		},
		exists: detail,
	}
}

func (e *ResourceAlreadyExists) ResourceID() string {
	return e.exists.GetResourceId()
}

func (e *ResourceAlreadyExists) Error() string {
	return fmt.Sprintf(
		"Resource already exists %s: service %s, resource ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
	)
}
