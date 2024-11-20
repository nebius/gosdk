package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*OutOfRange)(nil)
)

type OutOfRange struct {
	commonServiceError
	outOfRange *common.OutOfRange
}

func NewOutOfRange(se *common.ServiceError, detail *common.OutOfRange) *OutOfRange {
	return &OutOfRange{
		commonServiceError: commonServiceError{
			source: se,
		},
		outOfRange: detail,
	}
}

func (e *OutOfRange) Requested() string {
	return e.outOfRange.GetRequested()
}
func (e *OutOfRange) Limit() string {
	return e.outOfRange.GetLimit()
}

func (e *OutOfRange) Error() string {
	return fmt.Sprintf(
		"Out of range %s: service %s, requested: %s, limit: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.Requested(),
		e.Limit(),
	)
}
