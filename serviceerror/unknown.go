package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*Unknown)(nil)
)

type Unknown struct {
	commonServiceError
}

func NewUnknown(se *common.ServiceError) *Unknown {
	return &Unknown{
		commonServiceError: commonServiceError{
			source: se,
		},
	}
}

func (e *Unknown) IsUnknown() bool {
	return true
}

func (e *Unknown) Error() string {
	return fmt.Sprintf(
		"Service %s error %s",
		e.source.GetService(),
		e.source.GetCode(),
	)
}
