package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*TooManyRequests)(nil)
)

type TooManyRequests struct {
	commonServiceError
	violation *common.TooManyRequests
}

func NewTooManyRequests(se *common.ServiceError, detail *common.TooManyRequests) *TooManyRequests {
	return &TooManyRequests{
		commonServiceError: commonServiceError{
			source: se,
		},
		violation: detail,
	}
}

func (e *TooManyRequests) Violation() string {
	return e.violation.GetViolation()
}

func (e *TooManyRequests) Error() string {
	return fmt.Sprintf(
		"Too many requests %s: service %s: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.violation.GetViolation(),
	)
}
