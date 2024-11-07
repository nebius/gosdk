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
	return e.notEnoughResources.Violations
}

func (e *NotEnoughResources) Error() string {
	violations := make([]string, 0, len(e.notEnoughResources.Violations))
	for _, violation := range e.notEnoughResources.Violations {
		violations = append(violations, fmt.Sprintf(
			"  %s (requested %s): %s",
			violation.ResourceType,
			violation.Requested,
			violation.Message,
		))
	}
	return fmt.Sprintf(
		"Not enough resources %s: service %s, violations:\n%s",
		e.source.Code,
		e.source.Service,
		strings.Join(violations, "\n"),
	)
}
