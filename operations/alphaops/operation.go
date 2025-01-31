package alphaops

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
)

// Operation is a wrapper over protobuf operation to simplify some actions.
//
// Deprecated: migrate to common.v1.
type Operation struct {
	service      v1alpha1.OperationServiceClient
	mu           sync.RWMutex
	proto        *v1alpha1.Operation
	unmarshalled unmarshalled
}

type unmarshalled struct {
	request        proto.Message
	requestHeaders http.Header
	resource       proto.Message
	progressData   proto.Message
}

// Wrap validates an operation and returns Operation wrapper.
//
// Deprecated: migrate to common.v1.
func Wrap(operation *v1alpha1.Operation, service v1alpha1.OperationServiceClient) (*Operation, error) {
	msgs, err := validateAndUnmarshal(operation)
	if err != nil {
		return nil, err
	}

	return &Operation{
		service:      service,
		mu:           sync.RWMutex{},
		proto:        operation,
		unmarshalled: msgs,
	}, nil
}

// Done returns true if the operation is completed.
func (o *Operation) Done() bool {
	return o.Status() != nil
}

// Successful returns true if the operation is completed and status code is OK.
func (o *Operation) Successful() bool {
	s := o.Status()
	return s != nil && s.Code() == codes.OK
}

// Poll updates the operation from the server. It does nothing if the operation is done.
func (o *Operation) Poll(ctx context.Context, opts ...grpc.CallOption) (*Operation, error) {
	if o.Done() {
		return o, nil
	}

	operation, err := o.service.Get(ctx, &v1alpha1.GetOperationRequest{Id: o.ID()}, opts...)
	if err != nil {
		return nil, fmt.Errorf("can't get operation: %w", err)
	}

	msgs, err := validateAndUnmarshal(operation)
	if err != nil {
		return nil, err
	}

	o.mu.Lock()
	o.proto = operation
	o.unmarshalled = msgs
	o.mu.Unlock()
	return o, nil
}

// Wait polls the operation from the server until it's done.
//
// Important: It returns [*Error] if operation is not successful.
//
// Use [PollInterval] call option to override the [DefaultPollInterval].
func (o *Operation) Wait(ctx context.Context, opts ...grpc.CallOption) (*Operation, error) {
	if o.Done() {
		return o, NewError(o)
	}
	interval := DefaultPollInterval
	for _, opt := range opts {
		constant, isConstant := opt.(pollInterval)
		if isConstant && constant.interval > 0 {
			interval = constant.interval
		}
	}
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(interval):
			_, err := o.Poll(ctx, opts...)
			if err != nil {
				return nil, err
			}

			if o.Done() {
				return o, NewError(o)
			}
		}
	}
}

// ID returns operation ID.
func (o *Operation) ID() string {
	return o.Raw().GetId()
}

// Description returns a human-readable description of the operation.
func (o *Operation) Description() string {
	return o.Raw().GetDescription()
}

// CreatedAt returns the operation creation timestamp.
func (o *Operation) CreatedAt() time.Time {
	return o.Raw().GetCreatedAt().AsTime()
}

// CreatedBy returns the operation initiator ID.
func (o *Operation) CreatedBy() string {
	return o.Raw().GetCreatedBy()
}

// FinishedAt returns the completion timestamp if the operation is done and nil otherwise.
func (o *Operation) FinishedAt() *time.Time {
	finishedAt := o.Raw().GetFinishedAt()
	if finishedAt == nil {
		return nil
	}
	asTime := finishedAt.AsTime()
	return &asTime
}

// Request returns the request that generated this operation.
func (o *Operation) Request() proto.Message {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.request
}

// RequestHeaders returns headers of the request that generated this operation.
func (o *Operation) RequestHeaders() http.Header {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.requestHeaders.Clone()
}

// ResourceID returns ID of the resource that this operation creates, updates, deletes or otherwise changes.
// It may be empty, see API documentation.
func (o *Operation) ResourceID() string {
	return o.Raw().GetResourceId()
}

// Resource returns the snapshot of the resource at the moment this operation started.
// It returns nil if a proto message is not present or is the google.protobuf.Empty.
// See API documentation for details.
func (o *Operation) Resource() proto.Message {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.resource
}

// ProgressData returns additional information about the progress of the operation.
// It returns nil if a proto message is not present or is the google.protobuf.Empty.
// Must be nil after operation is done.
func (o *Operation) ProgressData() proto.Message {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.progressData
}

// Status returns the status of the completed operation and nil otherwise.
// See API documentation for details.
func (o *Operation) Status() *status.Status {
	s := o.Raw().GetStatus()
	if s == nil {
		return nil
	}
	return status.FromProto(s)
}

// Raw returns underlying protobuf operation.
func (o *Operation) Raw() *v1alpha1.Operation {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.proto
}

func validateAndUnmarshal(operation *v1alpha1.Operation) (unmarshalled, error) {
	var errs error
	var msgs unmarshalled

	if operation.GetId() == "" {
		errs = errors.Join(errs, errors.New("id: empty operation ID"))
	}

	err := operation.GetCreatedAt().CheckValid()
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("created_at: %w", err))
	}

	if operation.GetFinishedAt() != nil {
		err = operation.GetFinishedAt().CheckValid()
		if err != nil {
			errs = errors.Join(errs, fmt.Errorf("finished_at: %w", err))
		}
	}

	msgs.request = unmarshalNotEmpty(operation.GetRequest())
	msgs.resource = unmarshalNotEmpty(operation.GetResource())
	msgs.progressData = unmarshalNotEmpty(operation.GetProgressData())

	msgs.requestHeaders = http.Header{}
	for name, header := range operation.GetRequestHeaders() {
		msgs.requestHeaders[name] = header.GetValues()
	}

	return msgs, errs
}

func unmarshalNotEmpty(message *anypb.Any) proto.Message {
	if message.GetTypeUrl() == "" || message.MessageIs(&emptypb.Empty{}) {
		return nil
	}

	res, err := message.UnmarshalNew()
	if err != nil {
		return nil // ignore the error to be compatible with newer services that may send unknown messages inside Any
	}

	return res
}
