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
	conn.ConventionResolverServiceIDToNameMap[JobServiceID] = "sp.msp"
}

func (s Services) Job() JobService {
	return NewJobService(s.sdk)
}

const JobServiceID conn.ServiceID = "nebius.msp.spark.v1alpha1.JobService"

type JobService interface {
	Get(context.Context, *v1alpha1.GetJobRequest, ...grpc.CallOption) (*v1alpha1.Job, error)
	List(context.Context, *v1alpha1.ListJobsRequest, ...grpc.CallOption) (*v1alpha1.ListJobsResponse, error)
	Filter(context.Context, *v1alpha1.ListJobsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Job, error]
	Create(context.Context, *v1alpha1.CreateJobRequest, ...grpc.CallOption) (operations.Operation, error)
	Cancel(context.Context, *v1alpha1.CancelJobRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)
}

type jobService struct {
	sdk iface.SDK
}

func NewJobService(sdk iface.SDK) JobService {
	return jobService{
		sdk: sdk,
	}
}

func (s jobService) Get(ctx context.Context, request *v1alpha1.GetJobRequest, opts ...grpc.CallOption) (*v1alpha1.Job, error) {
	address, err := s.sdk.Resolve(ctx, JobServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewJobServiceClient(con).Get(ctx, request, opts...)
}

func (s jobService) List(ctx context.Context, request *v1alpha1.ListJobsRequest, opts ...grpc.CallOption) (*v1alpha1.ListJobsResponse, error) {
	address, err := s.sdk.Resolve(ctx, JobServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewJobServiceClient(con).List(ctx, request, opts...)
}

func (s jobService) Filter(ctx context.Context, request *v1alpha1.ListJobsRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Job, error] {
	req := proto.Clone(request).(*v1alpha1.ListJobsRequest)
	return func(yield func(*v1alpha1.Job, error) bool) {
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

func (s jobService) Create(ctx context.Context, request *v1alpha1.CreateJobRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, JobServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewJobServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s jobService) Cancel(ctx context.Context, request *v1alpha1.CancelJobRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, JobServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1alpha1.NewJobServiceClient(con).Cancel(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v1.NewOperationServiceClient(con))
}

func (s jobService) GetOperation(ctx context.Context, request *v1.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, JobServiceID)
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

func (s jobService) ListOperations(ctx context.Context, request *v1.ListOperationsRequest, opts ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, JobServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewOperationServiceClient(con).List(ctx, request, opts...)
}
