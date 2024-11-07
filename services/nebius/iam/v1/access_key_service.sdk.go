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
	conn.ConventionResolverServiceIDToNameMap[AccessKeyServiceID] = "cpl.iam"
}

func (s Services) AccessKey() AccessKeyService {
	return NewAccessKeyService(s.sdk)
}

const AccessKeyServiceID conn.ServiceID = "nebius.iam.v1.AccessKeyService"

type AccessKeyService interface {
	Create(context.Context, *v1.CreateAccessKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	List(context.Context, *v1.ListAccessKeysRequest, ...grpc.CallOption) (*v1.ListAccessKeysResponse, error)
	Filter(context.Context, *v1.ListAccessKeysRequest, ...grpc.CallOption) iter.Seq2[*v1.AccessKey, error]
	ListByAccount(context.Context, *v1.ListAccessKeysByAccountRequest, ...grpc.CallOption) (*v1.ListAccessKeysResponse, error)
	Update(context.Context, *v1.UpdateAccessKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	GetById(context.Context, *v1.GetAccessKeyByIdRequest, ...grpc.CallOption) (*v1.AccessKey, error)
	GetByAwsId(context.Context, *v1.GetAccessKeyByAwsIdRequest, ...grpc.CallOption) (*v1.AccessKey, error)
	GetSecretOnce(context.Context, *v1.GetAccessKeySecretOnceRequest, ...grpc.CallOption) (*v1.GetAccessKeySecretOnceResponse, error)
	Activate(context.Context, *v1.ActivateAccessKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Deactivate(context.Context, *v1.DeactivateAccessKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteAccessKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type accessKeyService struct {
	sdk iface.SDK
}

func NewAccessKeyService(sdk iface.SDK) AccessKeyService {
	return accessKeyService{
		sdk: sdk,
	}
}

func (s accessKeyService) Create(ctx context.Context, request *v1.CreateAccessKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessKeyServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessKeyService) List(ctx context.Context, request *v1.ListAccessKeysRequest, opts ...grpc.CallOption) (*v1.ListAccessKeysResponse, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessKeyServiceClient(con).List(ctx, request, opts...)
}

func (s accessKeyService) Filter(ctx context.Context, request *v1.ListAccessKeysRequest, opts ...grpc.CallOption) iter.Seq2[*v1.AccessKey, error] {
	req := proto.Clone(request).(*v1.ListAccessKeysRequest)
	return func(yield func(*v1.AccessKey, error) bool) {
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

func (s accessKeyService) ListByAccount(ctx context.Context, request *v1.ListAccessKeysByAccountRequest, opts ...grpc.CallOption) (*v1.ListAccessKeysResponse, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessKeyServiceClient(con).ListByAccount(ctx, request, opts...)
}

func (s accessKeyService) Update(ctx context.Context, request *v1.UpdateAccessKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessKeyServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessKeyService) GetById(ctx context.Context, request *v1.GetAccessKeyByIdRequest, opts ...grpc.CallOption) (*v1.AccessKey, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessKeyServiceClient(con).GetById(ctx, request, opts...)
}

func (s accessKeyService) GetByAwsId(ctx context.Context, request *v1.GetAccessKeyByAwsIdRequest, opts ...grpc.CallOption) (*v1.AccessKey, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessKeyServiceClient(con).GetByAwsId(ctx, request, opts...)
}

func (s accessKeyService) GetSecretOnce(ctx context.Context, request *v1.GetAccessKeySecretOnceRequest, opts ...grpc.CallOption) (*v1.GetAccessKeySecretOnceResponse, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessKeyServiceClient(con).GetSecretOnce(ctx, request, opts...)
}

func (s accessKeyService) Activate(ctx context.Context, request *v1.ActivateAccessKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessKeyServiceClient(con).Activate(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessKeyService) Deactivate(ctx context.Context, request *v1.DeactivateAccessKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessKeyServiceClient(con).Deactivate(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessKeyService) Delete(ctx context.Context, request *v1.DeleteAccessKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessKeyServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessKeyService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
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

func (s accessKeyService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, AccessKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
