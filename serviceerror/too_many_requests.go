package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*TooManyRequests)(nil)

type TooManyRequests struct {
	commonDetail
	violation *common.TooManyRequests
}

func newTooManyRequests(se *common.ServiceError, detail *common.TooManyRequests) *TooManyRequests {
	return &TooManyRequests{
		commonDetail: commonDetail{
			source: se,
		},
		violation: detail,
	}
}

func (e *TooManyRequests) Violation() string {
	return e.violation.GetViolation()
}

func (e *TooManyRequests) String() string {
	return fmt.Sprintf(
		"Too many requests %s: service %s: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.violation.GetViolation(),
	)
}
