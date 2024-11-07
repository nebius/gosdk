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

func (s Services) Instance() InstanceService {
	return NewInstanceService(s.sdk)
}

const InstanceServiceID conn.ServiceID = "nebius.compute.v1alpha1.InstanceService"

type InstanceService interface {
	Get(context.Context, *v1alpha1.GetInstanceRequest, ...grpc.CallOption) (*v1alpha1.Instance, error)
	GetByName(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v1alpha1.Instance, error)
	List(context.Context, *v1alpha1.ListInstancesRequest, ...grpc.CallOption) (*v1alpha1.ListInstancesResponse, error)
	Filter(context.Context, *v1alpha1.ListInstancesRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Instance, error]
	Create(context.Context, *v1alpha1.CreateInstanceRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Update(context.Context, *v1alpha1.UpdateInstanceRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Delete(context.Context, *v1alpha1.DeleteInstanceRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Start(context.Context, *v1alpha1.StartInstanceRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	Stop(context.Context, *v1alpha1.StopInstanceRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	ListOperationsByParent(context.Context, *v1alpha11.ListOperationsByParentRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
	GetOperation(context.Context, *v1alpha11.GetOperationRequest, ...grpc.CallOption) (*alphaops.Operation, error)
	ListOperations(context.Context, *v1alpha11.ListOperationsRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
}

type instanceService struct {
	sdk iface.SDK
}

func NewInstanceService(sdk iface.SDK) InstanceService {
	return instanceService{
		sdk: sdk,
	}
}

func (s instanceService) Get(ctx context.Context, request *v1alpha1.GetInstanceRequest, opts ...grpc.CallOption) (*v1alpha1.Instance, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewInstanceServiceClient(con).Get(ctx, request, opts...)
}

func (s instanceService) GetByName(ctx context.Context, request *v1.GetByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Instance, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewInstanceServiceClient(con).GetByName(ctx, request, opts...)
}

func (s instanceService) List(ctx context.Context, request *v1alpha1.ListInstancesRequest, opts ...grpc.CallOption) (*v1alpha1.ListInstancesResponse, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewInstanceServiceClient(con).List(ctx, request, opts...)
}

func (s instanceService) Filter(ctx context.Context, request *v1alpha1.ListInstancesRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Instance, error] {
	req := proto.Clone(request).(*v1alpha1.ListInstancesRequest)
	return func(yield func(*v1alpha1.Instance, error) bool) {
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

func (s instanceService) Create(ctx context.Context, request *v1alpha1.CreateInstanceRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewInstanceServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s instanceService) Update(ctx context.Context, request *v1alpha1.UpdateInstanceRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewInstanceServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s instanceService) Delete(ctx context.Context, request *v1alpha1.DeleteInstanceRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewInstanceServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s instanceService) Start(ctx context.Context, request *v1alpha1.StartInstanceRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewInstanceServiceClient(con).Start(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s instanceService) Stop(ctx context.Context, request *v1alpha1.StopInstanceRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewInstanceServiceClient(con).Stop(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return alphaops.Wrap(op, v1alpha11.NewOperationServiceClient(con))
}

func (s instanceService) ListOperationsByParent(ctx context.Context, request *v1alpha11.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewInstanceServiceClient(con).ListOperationsByParent(ctx, request, opts...)
}

func (s instanceService) GetOperation(ctx context.Context, request *v1alpha11.GetOperationRequest, opts ...grpc.CallOption) (*alphaops.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
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

func (s instanceService) ListOperations(ctx context.Context, request *v1alpha11.ListOperationsRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, InstanceServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
