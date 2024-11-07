package serviceerror

import (
	"fmt"
	"strings"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*BadRequest)(nil)
)

type BadRequest struct {
	commonServiceError
	badRequest *common.BadRequest
}

func NewBadRequest(se *common.ServiceError, br *common.BadRequest) *BadRequest {
	return &BadRequest{
		commonServiceError: commonServiceError{
			source: se,
		},
		badRequest: br,
	}
}

func (e *BadRequest) Violations() []*common.BadRequest_Violation {
	return e.badRequest.Violations
}

func (e *BadRequest) Error() string {
	violations := make([]string, 0, len(e.badRequest.Violations))
	for _, violation := range e.badRequest.Violations {
		violations = append(violations, fmt.Sprintf(
			"  %s: %s",
			violation.Field,
			violation.Message,
		))
	}
	return fmt.Sprintf(
		"Bad Request %s: service %s, violations:\n%s",
		e.source.Code,
		e.source.Service,
		strings.Join(violations, "\n"),
	)
}
