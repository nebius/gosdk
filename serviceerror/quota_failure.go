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
	return e.failure.GetViolations()
}

func (e *QuotaFailure) Error() string {
	violations := make([]string, 0, len(e.failure.GetViolations()))
	for _, violation := range e.failure.GetViolations() {
		violations = append(violations, fmt.Sprintf(
			"  %s (limit %s, requested %s): %s",
			violation.GetQuota(),
			violation.GetLimit(),
			violation.GetRequested(),
			violation.GetMessage(),
		))
	}
	return fmt.Sprintf(
		"Quota failure %s: service %s, violations:\n%s",
		e.source.GetCode(),
		e.source.GetService(),
		strings.Join(violations, "\n"),
	)
}
