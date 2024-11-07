package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/compute/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Disk() DiskService {
	return NewDiskService(s.sdk)
}

const DiskServiceID conn.ServiceID = "nebius.compute.v1.DiskService"

type DiskService interface {
	Get(context.Context, *v1.GetDiskRequest, ...grpc.CallOption) (*v1.Disk, error)
	GetByName(context.Context, *v11.GetByNameRequest, ...grpc.CallOption) (*v1.Disk, error)
	List(context.Context, *v1.ListDisksRequest, ...grpc.CallOption) (*v1.ListDisksResponse, error)
	Filter(context.Context, *v1.ListDisksRequest, ...grpc.CallOption) iter.Seq2[*v1.Disk, error]
	Create(context.Context, *v1.CreateDiskRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateDiskRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteDiskRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperationsByParent(context.Context, *v1.ListOperationsByParentRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type diskService struct {
	sdk iface.SDK
}

func NewDiskService(sdk iface.SDK) DiskService {
	return diskService{
		sdk: sdk,
	}
}

func (s diskService) Get(ctx context.Context, request *v1.GetDiskRequest, opts ...grpc.CallOption) (*v1.Disk, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewDiskServiceClient(con).Get(ctx, request, opts...)
}

func (s diskService) GetByName(ctx context.Context, request *v11.GetByNameRequest, opts ...grpc.CallOption) (*v1.Disk, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewDiskServiceClient(con).GetByName(ctx, request, opts...)
}

func (s diskService) List(ctx context.Context, request *v1.ListDisksRequest, opts ...grpc.CallOption) (*v1.ListDisksResponse, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewDiskServiceClient(con).List(ctx, request, opts...)
}

func (s diskService) Filter(ctx context.Context, request *v1.ListDisksRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Disk, error] {
	req := proto.Clone(request).(*v1.ListDisksRequest)
	return func(yield func(*v1.Disk, error) bool) {
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

func (s diskService) Create(ctx context.Context, request *v1.CreateDiskRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewDiskServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s diskService) Update(ctx context.Context, request *v1.UpdateDiskRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewDiskServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s diskService) Delete(ctx context.Context, request *v1.DeleteDiskRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewDiskServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s diskService) ListOperationsByParent(ctx context.Context, request *v1.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewDiskServiceClient(con).ListOperationsByParent(ctx, request, opts...)
}

func (s diskService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
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

func (s diskService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
