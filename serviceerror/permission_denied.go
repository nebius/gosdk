package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*PermissionDenied)(nil)

type PermissionDenied struct {
	commonDetail
	denied *common.PermissionDenied
}

func newPermissionDenied(se *common.ServiceError, detail *common.PermissionDenied) *PermissionDenied {
	return &PermissionDenied{
		commonDetail: commonDetail{
			source: se,
		},
		denied: detail,
	}
}

func (e *PermissionDenied) ResourceID() string {
	return e.denied.GetResourceId()
}

func (e *PermissionDenied) String() string {
	return fmt.Sprintf(
		"Permission denied %s: service %s, resource ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.ResourceID(),
	)
}
