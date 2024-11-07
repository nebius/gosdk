package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	alphaops "github.com/nebius/gosdk/operations/alphaops"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha11 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/compute/v1alpha1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Disk() DiskService {
	return NewDiskService(s.sdk)
}

const DiskServiceID conn.ServiceID = "nebius.compute.v1alpha1.DiskService"

type DiskService interface {
	Get(context.Context, *v1alpha1.GetDiskRequest, ...grpc.CallOption) (*v1alpha1.Disk, error)
	GetByName(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v1alpha1.Disk, error)
	List(context.Context, *v1alpha1.ListDisksRequest, ...grpc.CallOption) (*v1alpha1.ListDisksResponse, error)
	Filter(context.Context, *v1alpha1.ListDisksRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Disk, error]
	Create(context.Context, *v1alpha1.CreateDiskRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Update(context.Context, *v1alpha1.UpdateDiskRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Delete(context.Context, *v1alpha1.DeleteDiskRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	ListOperationsByParent(context.Context, *v1alpha11.ListOperationsByParentRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
	GetOperation(context.Context, *v1alpha11.GetOperationRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	ListOperations(context.Context, *v1alpha11.ListOperationsRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
}

type diskService struct {
	sdk iface.SDK
}

func NewDiskService(sdk iface.SDK) DiskService {
	return diskService{
		sdk: sdk,
	}
}

func (s diskService) Get(ctx context.Context, request *v1alpha1.GetDiskRequest, opts ...grpc.CallOption) (*v1alpha1.Disk, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewDiskServiceClient(con).Get(ctx, request, opts...)
}

func (s diskService) GetByName(ctx context.Context, request *v1.GetByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Disk, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewDiskServiceClient(con).GetByName(ctx, request, opts...)
}

func (s diskService) List(ctx context.Context, request *v1alpha1.ListDisksRequest, opts ...grpc.CallOption) (*v1alpha1.ListDisksResponse, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewDiskServiceClient(con).List(ctx, request, opts...)
}

func (s diskService) Filter(ctx context.Context, request *v1alpha1.ListDisksRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Disk, error] {
	req := proto.Clone(request).(*v1alpha1.ListDisksRequest)
	return func(yield func(*v1alpha1.Disk, error) bool) {
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

func (s diskService) Create(ctx context.Context, request *v1alpha1.CreateDiskRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewDiskServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s diskService) Update(ctx context.Context, request *v1alpha1.UpdateDiskRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
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
	op, err := v1alpha1.NewDiskServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s diskService) Delete(ctx context.Context, request *v1alpha1.DeleteDiskRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewDiskServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s diskService) ListOperationsByParent(ctx context.Context, request *v1alpha11.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewDiskServiceClient(con).ListOperationsByParent(ctx, request, opts...)
}

func (s diskService) GetOperation(ctx context.Context, request *v1alpha11.GetOperationRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	client := v1alpha11.NewOperationServiceClient(con)
	op, err := client.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, client)
}

func (s diskService) ListOperations(ctx context.Context, request *v1alpha11.ListOperationsRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, DiskServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
