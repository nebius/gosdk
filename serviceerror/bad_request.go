package serviceerror

import (
	"fmt"
	"strings"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*BadRequest)(nil)

type BadRequest struct {
	commonDetail
	badRequest *common.BadRequest
}

func newBadRequest(se *common.ServiceError, br *common.BadRequest) *BadRequest {
	return &BadRequest{
		commonDetail: commonDetail{
			source: se,
		},
		badRequest: br,
	}
}

func (e *BadRequest) Violations() []*common.BadRequest_Violation {
	return e.badRequest.GetViolations()
}

func (e *BadRequest) String() string {
	violations := make([]string, 0, len(e.badRequest.GetViolations()))
	for _, violation := range e.badRequest.GetViolations() {
		violations = append(violations, fmt.Sprintf(
			"  %s: %s",
			violation.GetField(),
			violation.GetMessage(),
		))
	}
	return fmt.Sprintf(
		"Bad Request %s: service %s, violations:\n%s",
		e.source.GetCode(),
		e.source.GetService(),
		strings.Join(violations, "\n"),
	)
}
