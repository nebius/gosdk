package operations

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

	"github.com/nebius/gosdk/conn"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
)

// Operation is a wrapper over protobuf operation message.
// All mutating service methods returns it.
// An operations can be either synchronous or asynchronous.
// Synchronous operations are completed immediately, while asynchronous operations may take time to finish.
// The client should call Wait to ensure an operation is fully completed.
type Operation interface {
	// Done returns true if the operation is completed.
	Done() bool

	// Successful returns true if the operation is completed and status code is OK.
	Successful() bool

	// Poll fetches an updated operation from the server.
	// It does nothing if the operation is done.
	Poll(context.Context, ...grpc.CallOption) (Operation, error)

	// Wait polls the operation from the server until it's done.
	//
	// Important: It returns [*Error] if operation is not successful.
	//
	// Use [PollInterval] call option to override the [DefaultPollInterval].
	Wait(context.Context, ...grpc.CallOption) (Operation, error)

	// ID returns operation ID.
	ID() string

	// Description returns a human-readable description of the operation.
	Description() string

	// CreatedAt returns the operation creation timestamp.
	CreatedAt() time.Time

	// CreatedBy returns the operation initiator ID.
	CreatedBy() string

	// FinishedAt returns the completion timestamp if the operation is done and nil otherwise.
	FinishedAt() *time.Time

	// Request returns the request that generated this operation.
	Request() proto.Message

	// RequestHeaders returns headers of the request that generated this operation.
	RequestHeaders() http.Header

	// ResourceID returns ID of the resource that this operation creates, updates, deletes or otherwise changes.
	// It may be empty, see API documentation.
	ResourceID() string

	// ProgressData returns additional information about the progress of the operation.
	// It returns nil if a proto message is not present or is the google.protobuf.Empty.
	// Must be nil after operation is done.
	ProgressData() proto.Message

	// Status returns the status of the completed operation and nil otherwise.
	// The operation is successful if `Status() != nil && Status().Code() == codes.OK`.
	Status() *status.Status

	// Raw returns underlying protobuf operation.
	Raw() *common.Operation
}

type opWrapper struct {
	service      common.OperationServiceClient
	mu           sync.RWMutex
	proto        *common.Operation
	unmarshalled unmarshalled
}

type unmarshalled struct {
	request        proto.Message
	requestHeaders http.Header
	progressData   proto.Message
}

// New validates an operation and returns Operation wrapper.
func New(operation *common.Operation, service common.OperationServiceClient) (Operation, error) {
	msgs, err := validateAndUnmarshal(operation)
	if err != nil {
		return nil, err
	}

	return &opWrapper{
		service:      service,
		mu:           sync.RWMutex{},
		proto:        operation,
		unmarshalled: msgs,
	}, nil
}

// NewCompleted accepts completed operation and returns Operation wrapper.
// Unlike [New] it doesn't accept [common.OperationServiceClient] as the operation is done.
// It returns an error if the operation is incomplete.
func NewCompleted(operation *common.Operation) (Operation, error) {
	wrapper, err := New(operation, nil)
	if err != nil {
		return nil, err
	}

	if !wrapper.Done() {
		return nil, errors.New("operation is not completed")
	}

	return wrapper, nil
}

// Done returns true if the operation is completed.
func (o *opWrapper) Done() bool {
	// if you change this, do not forget to change Stub
	return o.Status() != nil
}

// Successful returns true if the operation is completed and status code is OK.
func (o *opWrapper) Successful() bool {
	// if you change this, do not forget to change Stub
	s := o.Status()
	return s != nil && s.Code() == codes.OK
}

// Poll fetches an updated operation from the server.
// It does nothing if the operation is done.
func (o *opWrapper) Poll(ctx context.Context, opts ...grpc.CallOption) (Operation, error) {
	// if you change this, do not forget to change Stub
	if o.Done() {
		return o, nil
	}

	operation, err := o.service.Get(ctx, &common.GetOperationRequest{Id: o.ID()}, opts...)
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
func (o *opWrapper) Wait(ctx context.Context, opts ...grpc.CallOption) (Operation, error) {
	// if you change this, do not forget to change Stub
	if o.Done() {
		return o, NewError(o)
	}
	interval := DefaultPollInterval
	for _, opt := range opts {
		// TODO: add more options to support backoff
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
				tErr := &conn.TimeoutError{}
				if errors.As(err, &tErr) {
					continue
				}
				return nil, err
			}

			if o.Done() {
				return o, NewError(o)
			}
		}
	}
}

// ID returns operation ID.
func (o *opWrapper) ID() string {
	return o.Raw().GetId()
}

// Description returns a human-readable description of the operation.
func (o *opWrapper) Description() string {
	return o.Raw().GetDescription()
}

// CreatedAt returns the operation creation timestamp.
func (o *opWrapper) CreatedAt() time.Time {
	return o.Raw().GetCreatedAt().AsTime()
}

// CreatedBy returns the operation initiator ID.
func (o *opWrapper) CreatedBy() string {
	return o.Raw().GetCreatedBy()
}

// FinishedAt returns the completion timestamp if the operation is done and nil otherwise.
func (o *opWrapper) FinishedAt() *time.Time {
	finishedAt := o.Raw().GetFinishedAt()
	if finishedAt == nil {
		return nil
	}
	asTime := finishedAt.AsTime()
	return &asTime
}

// Request returns the request that generated this operation.
func (o *opWrapper) Request() proto.Message {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.request
}

// RequestHeaders returns headers of the request that generated this operation.
func (o *opWrapper) RequestHeaders() http.Header {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.requestHeaders.Clone()
}

// ResourceID returns ID of the resource that this operation creates, updates, deletes or otherwise changes.
// It may be empty, see API documentation.
func (o *opWrapper) ResourceID() string {
	return o.Raw().GetResourceId()
}

// ProgressData returns additional information about the progress of the operation.
// It returns nil if a proto message is not present or is the google.protobuf.Empty.
// Must be nil after operation is done.
func (o *opWrapper) ProgressData() proto.Message {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.unmarshalled.progressData
}

// Status returns the status of the completed operation and nil otherwise.
// The operation is successful if `Status() != nil && Status().Code() == codes.OK`.
func (o *opWrapper) Status() *status.Status {
	s := o.Raw().GetStatus()
	if s == nil {
		return nil
	}
	return status.FromProto(s)
}

// Raw returns underlying protobuf operation.
func (o *opWrapper) Raw() *common.Operation {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.proto
}

func validateAndUnmarshal(operation *common.Operation) (unmarshalled, error) {
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
