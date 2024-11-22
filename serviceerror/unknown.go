package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*Unknown)(nil)

type Unknown struct {
	commonDetail
}

func newUnknown(se *common.ServiceError) *Unknown {
	return &Unknown{
		commonDetail: commonDetail{
			source: se,
		},
	}
}

func (e *Unknown) IsUnknown() bool {
	return true
}

func (e *Unknown) String() string {
	return fmt.Sprintf(
		"Service %s error %s",
		e.source.GetService(),
		e.source.GetCode(),
	)
}
