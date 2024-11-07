package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[AuthPublicKeyServiceID] = "cpl.iam"
}

func (s Services) AuthPublicKey() AuthPublicKeyService {
	return NewAuthPublicKeyService(s.sdk)
}

const AuthPublicKeyServiceID conn.ServiceID = "nebius.iam.v1.AuthPublicKeyService"

type AuthPublicKeyService interface {
	Create(context.Context, *v1.CreateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetAuthPublicKeyRequest, ...grpc.CallOption) (*v1.AuthPublicKey, error)
	List(context.Context, *v1.ListAuthPublicKeyRequest, ...grpc.CallOption) (*v1.ListAuthPublicKeyResponse, error)
	Filter(context.Context, *v1.ListAuthPublicKeyRequest, ...grpc.CallOption) iter.Seq2[*v1.AuthPublicKey, error]
	ListByAccount(context.Context, *v1.ListAuthPublicKeyByAccountRequest, ...grpc.CallOption) (*v1.ListAuthPublicKeyResponse, error)
	Update(context.Context, *v1.UpdateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Activate(context.Context, *v1.ActivateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Deactivate(context.Context, *v1.DeactivateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type authPublicKeyService struct {
	sdk iface.SDK
}

func NewAuthPublicKeyService(sdk iface.SDK) AuthPublicKeyService {
	return authPublicKeyService{
		sdk: sdk,
	}
}

func (s authPublicKeyService) Create(ctx context.Context, request *v1.CreateAuthPublicKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAuthPublicKeyServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s authPublicKeyService) Get(ctx context.Context, request *v1.GetAuthPublicKeyRequest, opts ...grpc.CallOption) (*v1.AuthPublicKey, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthPublicKeyServiceClient(con).Get(ctx, request, opts...)
}

func (s authPublicKeyService) List(ctx context.Context, request *v1.ListAuthPublicKeyRequest, opts ...grpc.CallOption) (*v1.ListAuthPublicKeyResponse, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthPublicKeyServiceClient(con).List(ctx, request, opts...)
}

func (s authPublicKeyService) Filter(ctx context.Context, request *v1.ListAuthPublicKeyRequest, opts ...grpc.CallOption) iter.Seq2[*v1.AuthPublicKey, error] {
	req := proto.Clone(request).(*v1.ListAuthPublicKeyRequest)
	return func(yield func(*v1.AuthPublicKey, error) bool) {
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

func (s authPublicKeyService) ListByAccount(ctx context.Context, request *v1.ListAuthPublicKeyByAccountRequest, opts ...grpc.CallOption) (*v1.ListAuthPublicKeyResponse, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAuthPublicKeyServiceClient(con).ListByAccount(ctx, request, opts...)
}

func (s authPublicKeyService) Update(ctx context.Context, request *v1.UpdateAuthPublicKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAuthPublicKeyServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s authPublicKeyService) Activate(ctx context.Context, request *v1.ActivateAuthPublicKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAuthPublicKeyServiceClient(con).Activate(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s authPublicKeyService) Deactivate(ctx context.Context, request *v1.DeactivateAuthPublicKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAuthPublicKeyServiceClient(con).Deactivate(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s authPublicKeyService) Delete(ctx context.Context, request *v1.DeleteAuthPublicKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAuthPublicKeyServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s authPublicKeyService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	client := v11.NewOperationServiceClient(con)
	op, err := client.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, client)
}

func (s authPublicKeyService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, AuthPublicKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
