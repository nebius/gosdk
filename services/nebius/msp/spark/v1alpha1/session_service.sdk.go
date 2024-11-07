package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/msp/spark/v1alpha1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[SessionServiceID] = "sp.msp"
}

func (s Services) Session() SessionService {
	return NewSessionService(s.sdk)
}

const SessionServiceID conn.ServiceID = "nebius.msp.spark.v1alpha1.SessionService"

type SessionService interface {
	Get(context.Context, *v1alpha1.GetSessionRequest, ...grpc.CallOption) (*v1alpha1.Session, error)
	GetByName(context.Context, *v1alpha1.GetSessionByNameRequest, ...grpc.CallOption) (*v1alpha1.Session, error)
	List(context.Context, *v1alpha1.ListSessionsRequest, ...grpc.CallOption) (*v1alpha1.ListSessionsResponse, error)
	Filter(context.Context, *v1alpha1.ListSessionsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Session, error]
	Create(context.Context, *v1alpha1.CreateSessionRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1alpha1.DeleteSessionRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)
}

type sessionService struct {
	sdk iface.SDK
}

func NewSessionService(sdk iface.SDK) SessionService {
	return sessionService{
		sdk: sdk,
	}
}

func (s sessionService) Get(ctx context.Context, request *v1alpha1.GetSessionRequest, opts ...grpc.CallOption) (*v1alpha1.Session, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSessionServiceClient(con).Get(ctx, request, opts...)
}

func (s sessionService) GetByName(ctx context.Context, request *v1alpha1.GetSessionByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Session, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSessionServiceClient(con).GetByName(ctx, request, opts...)
}

func (s sessionService) List(ctx context.Context, request *v1alpha1.ListSessionsRequest, opts ...grpc.CallOption) (*v1alpha1.ListSessionsResponse, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSessionServiceClient(con).List(ctx, request, opts...)
}

func (s sessionService) Filter(ctx context.Context, request *v1alpha1.ListSessionsRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Session, error] {
	req := proto.Clone(request).(*v1alpha1.ListSessionsRequest)
	return func(yield func(*v1alpha1.Session, error) bool) {
		for {
			res, err := s.List(ctx, req, opts...)
			if err != nil {
				yield(nil, err)
				return
			}

			for _, item := range res.GetItems() {
				if !yield(item, nil) {
					return
				}
			}

			if res.GetNextPageToken() == "" {
				return
			}

			req.PageToken = res.GetNextPageToken()
		}
	}
}

func (s sessionService) Create(ctx context.Context, request *v1alpha1.CreateSessionRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewSessionServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s sessionService) Delete(ctx context.Context, request *v1alpha1.DeleteSessionRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewSessionServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s sessionService) GetOperation(ctx context.Context, request *v1.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	client := v1.NewOperationServiceClient(con)
	op, err := client.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, client)
}

func (s sessionService) ListOperations(ctx context.Context, request *v1.ListOperationsRequest, opts ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, SessionServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewOperationServiceClient(con).List(ctx, request, opts...)
}
