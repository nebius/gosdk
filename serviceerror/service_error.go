package serviceerror

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

type ServiceError interface {
	error
	Proto() *common.ServiceError
	Service() string
	Code() string
	RetryType() common.ServiceError_RetryType
}

func ServiceErrorFromAny(detail *anypb.Any) (ServiceError, error) { //nolint:revive // don't want to name it FromAny
	serviceError := &common.ServiceError{}
	err := anypb.UnmarshalTo(
		detail,
		serviceError,
		proto.UnmarshalOptions{
			DiscardUnknown: true,
		},
	)
	if err != nil {
		return nil, err
	}
	switch seDetails := serviceError.Details.(type) {
	case *common.ServiceError_BadRequest:
		return NewBadRequest(serviceError, seDetails.BadRequest), nil
	case *common.ServiceError_BadResourceState:
		return NewBadResourceState(serviceError, seDetails.BadResourceState), nil
	case *common.ServiceError_ResourceNotFound:
		return NewResourceNotFound(serviceError, seDetails.ResourceNotFound), nil
	case *common.ServiceError_ResourceAlreadyExists:
		return NewResourceAlreadyExists(serviceError, seDetails.ResourceAlreadyExists), nil
	case *common.ServiceError_OutOfRange:
		return NewOutOfRange(serviceError, seDetails.OutOfRange), nil
	case *common.ServiceError_PermissionDenied:
		return NewPermissionDenied(serviceError, seDetails.PermissionDenied), nil
	case *common.ServiceError_ResourceConflict:
		return NewResourceConflict(serviceError, seDetails.ResourceConflict), nil
	case *common.ServiceError_OperationAborted:
		return NewOperationAborted(serviceError, seDetails.OperationAborted), nil
	case *common.ServiceError_TooManyRequests:
		return NewTooManyRequests(serviceError, seDetails.TooManyRequests), nil
	case *common.ServiceError_QuotaFailure:
		return NewQuotaFailure(serviceError, seDetails.QuotaFailure), nil
	case *common.ServiceError_InternalError:
		return NewInternal(serviceError, seDetails.InternalError), nil
	case *common.ServiceError_NotEnoughResources:
		return NewNotEnoughResources(serviceError, seDetails.NotEnoughResources), nil
	default:
		return NewUnknown(serviceError), nil
	}
}

type commonServiceError struct {
	source *common.ServiceError
}

// Code implements ServiceError.
func (e *commonServiceError) Code() string {
	return e.source.GetCode()
}

// Proto implements ServiceError.
func (e *commonServiceError) Proto() *common.ServiceError {
	return e.source
}

// RetryType implements ServiceError.
func (e *commonServiceError) RetryType() common.ServiceError_RetryType {
	return e.source.GetRetryType()
}

// Service implements ServiceError.
func (e *commonServiceError) Service() string {
	return e.source.GetService()
}
