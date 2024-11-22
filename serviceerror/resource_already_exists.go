package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*ResourceAlreadyExists)(nil)

type ResourceAlreadyExists struct {
	commonDetail
	exists *common.ResourceAlreadyExists
}

func newResourceAlreadyExists(se *common.ServiceError, detail *common.ResourceAlreadyExists) *ResourceAlreadyExists {
	return &ResourceAlreadyExists{
		commonDetail: commonDetail{
			source: se,
		},
		exists: detail,
	}
}

func (e *ResourceAlreadyExists) ResourceID() string {
	return e.exists.GetResourceId()
}

func (e *ResourceAlreadyExists) String() string {
	return fmt.Sprintf(
		"Resource already exists %s: service %s, resource ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
	)
}
