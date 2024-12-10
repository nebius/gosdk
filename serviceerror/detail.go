package serviceerror

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

// Detail is a wrapper for protobuf message nebius.common.v1.ServiceError.
//
// Here is the list of possible detail types:
//   - [BadRequest]
//   - [BadResourceState]
//   - [Internal]
//   - [NotEnoughResources]
//   - [OperationAborted]
//   - [OutOfRange]
//   - [PermissionDenied]
//   - [QuotaFailure]
//   - [ResourceAlreadyExists]
//   - [ResourceConflict]
//   - [ResourceNotFound]
//   - [TooManyRequests]
//   - [Unknown]
type Detail interface {
	String() string
	Proto() *common.ServiceError
	Service() string
	Code() string
	RetryType() common.ServiceError_RetryType
}

// NewDetail converts protobuf message to [Detail].
func NewDetail(serviceError *common.ServiceError) Detail {
	switch seDetails := serviceError.GetDetails().(type) {
	case *common.ServiceError_BadRequest:
		return newBadRequest(serviceError, seDetails.BadRequest)
	case *common.ServiceError_BadResourceState:
		return newBadResourceState(serviceError, seDetails.BadResourceState)
	case *common.ServiceError_ResourceNotFound:
		return newResourceNotFound(serviceError, seDetails.ResourceNotFound)
	case *common.ServiceError_ResourceAlreadyExists:
		return newResourceAlreadyExists(serviceError, seDetails.ResourceAlreadyExists)
	case *common.ServiceError_OutOfRange:
		return newOutOfRange(serviceError, seDetails.OutOfRange)
	case *common.ServiceError_PermissionDenied:
		return newPermissionDenied(serviceError, seDetails.PermissionDenied)
	case *common.ServiceError_ResourceConflict:
		return newResourceConflict(serviceError, seDetails.ResourceConflict)
	case *common.ServiceError_OperationAborted:
		return newOperationAborted(serviceError, seDetails.OperationAborted)
	case *common.ServiceError_TooManyRequests:
		return newTooManyRequests(serviceError, seDetails.TooManyRequests)
	case *common.ServiceError_QuotaFailure:
		return newQuotaFailure(serviceError, seDetails.QuotaFailure)
	case *common.ServiceError_InternalError:
		return newInternal(serviceError, seDetails.InternalError)
	case *common.ServiceError_NotEnoughResources:
		return newNotEnoughResources(serviceError, seDetails.NotEnoughResources)
	default:
		return newUnknown(serviceError)
	}
}

// NewDetailFromAny converts protobuf [anypb.Any] message to [Detail].
func NewDetailFromAny(detail *anypb.Any) (Detail, error) {
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

	return NewDetail(serviceError), nil
}

type commonDetail struct {
	source *common.ServiceError
}

func (e *commonDetail) Code() string {
	return e.source.GetCode()
}

func (e *commonDetail) Proto() *common.ServiceError {
	return e.source
}

func (e *commonDetail) RetryType() common.ServiceError_RetryType {
	return e.source.GetRetryType()
}

func (e *commonDetail) Service() string {
	return e.source.GetService()
}
