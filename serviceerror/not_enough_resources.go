package serviceerror

import (
	"fmt"
	"strings"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*NotEnoughResources)(nil)
)

type NotEnoughResources struct {
	commonServiceError
	notEnoughResources *common.NotEnoughResources
}

func NewNotEnoughResources(se *common.ServiceError, detail *common.NotEnoughResources) *NotEnoughResources {
	return &NotEnoughResources{
		commonServiceError: commonServiceError{
			source: se,
		},
		notEnoughResources: detail,
	}
}

func (e *NotEnoughResources) Violations() []*common.NotEnoughResources_Violation {
	return e.notEnoughResources.GetViolations()
}

func (e *NotEnoughResources) Error() string {
	violations := make([]string, 0, len(e.notEnoughResources.GetViolations()))
	for _, violation := range e.notEnoughResources.GetViolations() {
		violations = append(violations, fmt.Sprintf(
			"  %s (requested %s): %s",
			violation.GetResourceType(),
			violation.GetRequested(),
			violation.GetMessage(),
		))
	}
	return fmt.Sprintf(
		"Not enough resources %s: service %s, violations:\n%s",
		e.source.GetCode(),
		e.source.GetService(),
		strings.Join(violations, "\n"),
	)
}
