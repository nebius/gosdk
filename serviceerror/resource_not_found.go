package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*ResourceNotFound)(nil)

type ResourceNotFound struct {
	commonDetail
	notFound *common.ResourceNotFound
}

func newResourceNotFound(se *common.ServiceError, detail *common.ResourceNotFound) *ResourceNotFound {
	return &ResourceNotFound{
		commonDetail: commonDetail{
			source: se,
		},
		notFound: detail,
	}
}

func (e *ResourceNotFound) ResourceID() string {
	return e.notFound.GetResourceId()
}

func (e *ResourceNotFound) String() string {
	return fmt.Sprintf(
		"Resource not found %s: service %s, resource ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
	)
}
