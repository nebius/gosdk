// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[StaticKeyServiceID] = "cpl.iam"
}

func (s Services) StaticKey() StaticKeyService {
	return NewStaticKeyService(s.sdk)
}

const StaticKeyServiceID conn.ServiceID = "nebius.iam.v1.StaticKeyService"

type StaticKeyService interface {
	Issue(context.Context, *v1.IssueStaticKeyRequest, ...grpc.CallOption) (*v1.IssueStaticKeyResponse, error)
	List(context.Context, *v1.ListStaticKeysRequest, ...grpc.CallOption) (*v1.ListStaticKeysResponse, error)
	Filter(context.Context, *v1.ListStaticKeysRequest, ...grpc.CallOption) iter.Seq2[*v1.StaticKey, error]
	Get(context.Context, *v1.GetStaticKeyRequest, ...grpc.CallOption) (*v1.StaticKey, error)
	GetByName(context.Context, *v1.GetStaticKeyByNameRequest, ...grpc.CallOption) (*v1.StaticKey, error)
	Delete(context.Context, *v1.DeleteStaticKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	Find(context.Context, *v1.FindStaticKeyRequest, ...grpc.CallOption) (*v1.FindStaticKeyResponse, error)
	Revoke(context.Context, *v1.RevokeStaticKeyRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type staticKeyService struct {
	sdk iface.SDK
}

func NewStaticKeyService(sdk iface.SDK) StaticKeyService {
	return staticKeyService{
		sdk: sdk,
	}
}

func (s staticKeyService) Issue(ctx context.Context, request *v1.IssueStaticKeyRequest, opts ...grpc.CallOption) (*v1.IssueStaticKeyResponse, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewStaticKeyServiceClient(con).Issue(ctx, request, opts...)
}

func (s staticKeyService) List(ctx context.Context, request *v1.ListStaticKeysRequest, opts ...grpc.CallOption) (*v1.ListStaticKeysResponse, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewStaticKeyServiceClient(con).List(ctx, request, opts...)
}

func (s staticKeyService) Filter(ctx context.Context, request *v1.ListStaticKeysRequest, opts ...grpc.CallOption) iter.Seq2[*v1.StaticKey, error] {
	req := proto.Clone(request).(*v1.ListStaticKeysRequest)
	return func(yield func(*v1.StaticKey, error) bool) {
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

func (s staticKeyService) Get(ctx context.Context, request *v1.GetStaticKeyRequest, opts ...grpc.CallOption) (*v1.StaticKey, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewStaticKeyServiceClient(con).Get(ctx, request, opts...)
}

func (s staticKeyService) GetByName(ctx context.Context, request *v1.GetStaticKeyByNameRequest, opts ...grpc.CallOption) (*v1.StaticKey, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewStaticKeyServiceClient(con).GetByName(ctx, request, opts...)
}

func (s staticKeyService) Delete(ctx context.Context, request *v1.DeleteStaticKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewStaticKeyServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s staticKeyService) Find(ctx context.Context, request *v1.FindStaticKeyRequest, opts ...grpc.CallOption) (*v1.FindStaticKeyResponse, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewStaticKeyServiceClient(con).Find(ctx, request, opts...)
}

func (s staticKeyService) Revoke(ctx context.Context, request *v1.RevokeStaticKeyRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewStaticKeyServiceClient(con).Revoke(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s staticKeyService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
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

func (s staticKeyService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, StaticKeyServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
