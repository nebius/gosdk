package operations

import (
	"context"
	"net/http"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/nebius/gosdk/operations"
	commonpb "github.com/nebius/gosdk/proto/nebius/common/v1"
)

// Stub is an [operations.Operation] wrapper over a static [commonpb.Operation].
// It behaves like the real wrapper, mostly.
// It fails a test if the operation is not completed and Poll or Wait methods are called.
//
// If you need the full control over its behavior, use [MockOperation].
type Stub struct {
	t  require.TestingT
	op *commonpb.Operation
}

// NewStubOperation returns [Stub] operation.
// If you need the full control over its behavior, use [NewMockOperation].
func NewStubOperation(t require.TestingT, operation *commonpb.Operation) operations.Operation {
	return Stub{
		t:  t,
		op: operation,
	}
}

func (s Stub) Done() bool {
	return s.Status() != nil
}

func (s Stub) Successful() bool {
	st := s.Status()
	return st != nil && st.Code() == codes.OK
}

func (s Stub) Poll(context.Context, ...grpc.CallOption) (operations.Operation, error) {
	if s.Done() {
		return s, nil
	}
	require.Failf(
		s.t,
		"Poll is called on the stub with uncompleted operation",
		"Underlying operation: %s",
		protojson.Format(s.op),
	)
	return s, nil
}

func (s Stub) Wait(context.Context, ...grpc.CallOption) (operations.Operation, error) {
	if s.Done() {
		return s, operations.NewError(s)
	}
	require.Failf(
		s.t,
		"Wait is called on the stub with uncompleted operation",
		"Underlying operation: %s",
		protojson.Format(s.op),
	)
	return s, nil
}

func (s Stub) ID() string {
	return s.op.GetId()
}

func (s Stub) Description() string {
	return s.op.GetDescription()
}

func (s Stub) CreatedAt() time.Time {
	createdAt := s.op.GetCreatedAt()
	require.NoError(s.t, createdAt.CheckValid())
	return createdAt.AsTime()
}

func (s Stub) CreatedBy() string {
	return s.op.GetCreatedBy()
}

func (s Stub) FinishedAt() *time.Time {
	finishedAt := s.op.GetFinishedAt()
	if finishedAt == nil {
		return nil
	}
	require.NoError(s.t, finishedAt.CheckValid())
	asTime := finishedAt.AsTime()
	return &asTime
}

func (s Stub) Request() proto.Message {
	return unmarshalNotEmpty(s.t, s.op.GetRequest())
}

func (s Stub) RequestHeaders() http.Header {
	headers := make(http.Header, len(s.op.GetRequestHeaders()))
	for name, header := range s.op.GetRequestHeaders() {
		headers[name] = header.GetValues()
	}
	return headers
}

func (s Stub) ResourceID() string {
	return s.op.GetResourceId()
}

func (s Stub) ProgressData() proto.Message {
	return unmarshalNotEmpty(s.t, s.op.GetProgressData())
}

func (s Stub) Status() *status.Status {
	st := s.op.GetStatus()
	if st == nil {
		return nil
	}
	return status.FromProto(st)
}

func (s Stub) Raw() *commonpb.Operation {
	return s.op
}

func unmarshalNotEmpty(t require.TestingT, message *anypb.Any) proto.Message {
	if message == nil || message.MessageIs(&emptypb.Empty{}) {
		return nil
	}

	unmarshalled, err := message.UnmarshalNew()
	require.NoError(t, err)
	return unmarshalled
}
