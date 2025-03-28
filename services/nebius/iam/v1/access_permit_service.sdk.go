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
	conn.ConventionResolverServiceIDToNameMap[AccessPermitServiceID] = "cpl.iam"
}

func (s Services) AccessPermit() AccessPermitService {
	return NewAccessPermitService(s.sdk)
}

const AccessPermitServiceID conn.ServiceID = "nebius.iam.v1.AccessPermitService"

type AccessPermitService interface {
	Create(context.Context, *v1.CreateAccessPermitRequest, ...grpc.CallOption) (operations.Operation, error)
	List(context.Context, *v1.ListAccessPermitRequest, ...grpc.CallOption) (*v1.ListAccessPermitResponse, error)
	Filter(context.Context, *v1.ListAccessPermitRequest, ...grpc.CallOption) iter.Seq2[*v1.AccessPermit, error]
	Delete(context.Context, *v1.DeleteAccessPermitRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetAccessPermitRequest, ...grpc.CallOption) (*v1.AccessPermit, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type accessPermitService struct {
	sdk iface.SDK
}

func NewAccessPermitService(sdk iface.SDK) AccessPermitService {
	return accessPermitService{
		sdk: sdk,
	}
}

func (s accessPermitService) Create(ctx context.Context, request *v1.CreateAccessPermitRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessPermitServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessPermitServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessPermitService) List(ctx context.Context, request *v1.ListAccessPermitRequest, opts ...grpc.CallOption) (*v1.ListAccessPermitResponse, error) {
	address, err := s.sdk.Resolve(ctx, AccessPermitServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessPermitServiceClient(con).List(ctx, request, opts...)
}

func (s accessPermitService) Filter(ctx context.Context, request *v1.ListAccessPermitRequest, opts ...grpc.CallOption) iter.Seq2[*v1.AccessPermit, error] {
	req := proto.Clone(request).(*v1.ListAccessPermitRequest)
	return func(yield func(*v1.AccessPermit, error) bool) {
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

func (s accessPermitService) Delete(ctx context.Context, request *v1.DeleteAccessPermitRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessPermitServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewAccessPermitServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s accessPermitService) Get(ctx context.Context, request *v1.GetAccessPermitRequest, opts ...grpc.CallOption) (*v1.AccessPermit, error) {
	address, err := s.sdk.Resolve(ctx, AccessPermitServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewAccessPermitServiceClient(con).Get(ctx, request, opts...)
}

func (s accessPermitService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, AccessPermitServiceID)
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

func (s accessPermitService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, AccessPermitServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
