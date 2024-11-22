package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*OutOfRange)(nil)

type OutOfRange struct {
	commonDetail
	outOfRange *common.OutOfRange
}

func newOutOfRange(se *common.ServiceError, detail *common.OutOfRange) *OutOfRange {
	return &OutOfRange{
		commonDetail: commonDetail{
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

func (e *OutOfRange) String() string {
	return fmt.Sprintf(
		"Out of range %s: service %s, requested: %s, limit: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.Requested(),
		e.Limit(),
	)
}
