package serviceerror

import (
	"fmt"
	"strings"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*QuotaFailure)(nil)
)

type QuotaFailure struct {
	commonServiceError
	failure *common.QuotaFailure
}

func NewQuotaFailure(se *common.ServiceError, detail *common.QuotaFailure) *QuotaFailure {
	return &QuotaFailure{
		commonServiceError: commonServiceError{
			source: se,
		},
		failure: detail,
	}
}

func (e *QuotaFailure) Violations() []*common.QuotaFailure_Violation {
	return e.failure.Violations
}

func (e *QuotaFailure) Error() string {
	violations := make([]string, 0, len(e.failure.Violations))
	for _, violation := range e.failure.Violations {
		violations = append(violations, fmt.Sprintf(
			"  %s (limit %s, requested %s): %s",
			violation.Quota,
			violation.Limit,
			violation.Requested,
			violation.Message,
		))
	}
	return fmt.Sprintf(
		"Quota failure %s: service %s, violations:\n%s",
		e.source.Code,
		e.source.Service,
		strings.Join(violations, "\n"),
	)
}
