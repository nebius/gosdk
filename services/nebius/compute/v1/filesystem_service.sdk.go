// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	grpcheader "github.com/nebius/gosdk/proto/fieldmask/grpcheader"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/compute/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[FilesystemServiceID] = "compute"
}

func (s Services) Filesystem() FilesystemService {
	return NewFilesystemService(s.sdk)
}

const FilesystemServiceID conn.ServiceID = "nebius.compute.v1.FilesystemService"

type FilesystemService interface {
	Get(context.Context, *v1.GetFilesystemRequest, ...grpc.CallOption) (*v1.Filesystem, error)
	GetByName(context.Context, *v11.GetByNameRequest, ...grpc.CallOption) (*v1.Filesystem, error)
	List(context.Context, *v1.ListFilesystemsRequest, ...grpc.CallOption) (*v1.ListFilesystemsResponse, error)
	Filter(context.Context, *v1.ListFilesystemsRequest, ...grpc.CallOption) iter.Seq2[*v1.Filesystem, error]
	Create(context.Context, *v1.CreateFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperationsByParent(context.Context, *v1.ListOperationsByParentRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type filesystemService struct {
	sdk iface.SDK
}

func NewFilesystemService(sdk iface.SDK) FilesystemService {
	return filesystemService{
		sdk: sdk,
	}
}

func (s filesystemService) Get(ctx context.Context, request *v1.GetFilesystemRequest, opts ...grpc.CallOption) (*v1.Filesystem, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFilesystemServiceClient(con).Get(ctx, request, opts...)
}

func (s filesystemService) GetByName(ctx context.Context, request *v11.GetByNameRequest, opts ...grpc.CallOption) (*v1.Filesystem, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFilesystemServiceClient(con).GetByName(ctx, request, opts...)
}

func (s filesystemService) List(ctx context.Context, request *v1.ListFilesystemsRequest, opts ...grpc.CallOption) (*v1.ListFilesystemsResponse, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFilesystemServiceClient(con).List(ctx, request, opts...)
}

func (s filesystemService) Filter(ctx context.Context, request *v1.ListFilesystemsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Filesystem, error] {
	req := proto.Clone(request).(*v1.ListFilesystemsRequest)
	return func(yield func(*v1.Filesystem, error) bool) {
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

func (s filesystemService) Create(ctx context.Context, request *v1.CreateFilesystemRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFilesystemServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s filesystemService) Update(ctx context.Context, request *v1.UpdateFilesystemRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFilesystemServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s filesystemService) Delete(ctx context.Context, request *v1.DeleteFilesystemRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFilesystemServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s filesystemService) ListOperationsByParent(ctx context.Context, request *v1.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFilesystemServiceClient(con).ListOperationsByParent(ctx, request, opts...)
}

func (s filesystemService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
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

func (s filesystemService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, FilesystemServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
