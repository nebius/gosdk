package serviceerror

import (
	"fmt"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

var _ Detail = (*Internal)(nil)

type Internal struct {
	commonDetail
	internal *common.InternalError
}

func newInternal(se *common.ServiceError, detail *common.InternalError) *Internal {
	return &Internal{
		commonDetail: commonDetail{
			source: se,
		},
		internal: detail,
	}
}

func (e *Internal) RequestID() string {
	return e.internal.GetRequestId()
}

func (e *Internal) TraceID() string {
	return e.internal.GetTraceId()
}

func (e *Internal) String() string {
	return fmt.Sprintf(
		"Internal error %s: service %s, request ID: %s, trace ID: %s",
		e.source.GetCode(),
		e.source.GetService(),
		e.RequestID(),
		e.TraceID(),
	)
}
