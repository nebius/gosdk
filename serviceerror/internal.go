package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var (
	_ ServiceError = (*Internal)(nil)
)

type Internal struct {
	commonServiceError
	internal *common.InternalError
}

func NewInternal(se *common.ServiceError, detail *common.InternalError) *Internal {
	return &Internal{
		commonServiceError: commonServiceError{
			source: se,
		},
		internal: detail,
	}
}

func (e *Internal) RequestID() string {
	return e.internal.RequestId
}

func (e *Internal) TraceID() string {
	return e.internal.TraceId
}

func (e *Internal) Error() string {
	return fmt.Sprintf(
		"Internal error %s: service %s, request ID: %s, trace ID: %s",
		e.source.Code,
		e.source.Service,
		e.RequestID(),
		e.TraceID(),
	)
}
