// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	grpcheader "github.com/nebius/gosdk/proto/fieldmask/grpcheader"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/applications/v1alpha1"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[K8SReleaseServiceID] = "deployment-manager.mkt"
}

func (s Services) K8sRelease() K8SReleaseService {
	return NewK8SReleaseService(s.sdk)
}

const K8SReleaseServiceID conn.ServiceID = "nebius.applications.v1alpha1.K8sReleaseService"

type K8SReleaseService interface {
	Get(context.Context, *v1alpha1.GetK8SReleaseRequest, ...grpc.CallOption) (*v1alpha1.K8SRelease, error)
	List(context.Context, *v1alpha1.ListK8SReleasesRequest, ...grpc.CallOption) (*v1alpha1.ListK8SReleasesResponse, error)
	Filter(context.Context, *v1alpha1.ListK8SReleasesRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.K8SRelease, error]
	Create(context.Context, *v1alpha1.CreateK8SReleaseRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1alpha1.UpdateK8SReleaseRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1alpha1.DeleteK8SReleaseRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)
}

type k8SReleaseService struct {
	sdk iface.SDK
}

func NewK8SReleaseService(sdk iface.SDK) K8SReleaseService {
	return k8SReleaseService{
		sdk: sdk,
	}
}

func (s k8SReleaseService) Get(ctx context.Context, request *v1alpha1.GetK8SReleaseRequest, opts ...grpc.CallOption) (*v1alpha1.K8SRelease, error) {
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewK8SReleaseServiceClient(con).Get(ctx, request, opts...)
}

func (s k8SReleaseService) List(ctx context.Context, request *v1alpha1.ListK8SReleasesRequest, opts ...grpc.CallOption) (*v1alpha1.ListK8SReleasesResponse, error) {
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewK8SReleaseServiceClient(con).List(ctx, request, opts...)
}

func (s k8SReleaseService) Filter(ctx context.Context, request *v1alpha1.ListK8SReleasesRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.K8SRelease, error] {
	req := proto.Clone(request).(*v1alpha1.ListK8SReleasesRequest)
	return func(yield func(*v1alpha1.K8SRelease, error) bool) {
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

func (s k8SReleaseService) Create(ctx context.Context, request *v1alpha1.CreateK8SReleaseRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewK8SReleaseServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s k8SReleaseService) Update(ctx context.Context, request *v1alpha1.UpdateK8SReleaseRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewK8SReleaseServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s k8SReleaseService) Delete(ctx context.Context, request *v1alpha1.DeleteK8SReleaseRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewK8SReleaseServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s k8SReleaseService) GetOperation(ctx context.Context, request *v1.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
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

func (s k8SReleaseService) ListOperations(ctx context.Context, request *v1.ListOperationsRequest, opts ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, K8SReleaseServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewOperationServiceClient(con).List(ctx, request, opts...)
}
