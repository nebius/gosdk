package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/vpc/v1alpha1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Pool() PoolService {
	return NewPoolService(s.sdk)
}

const PoolServiceID conn.ServiceID = "nebius.vpc.v1alpha1.PoolService"

type PoolService interface {
	Get(context.Context, *v1alpha1.GetPoolRequest, ...grpc.CallOption) (*v1alpha1.Pool, error)
	GetByName(context.Context, *v1alpha1.GetPoolByNameRequest, ...grpc.CallOption) (*v1alpha1.Pool, error)
	List(context.Context, *v1alpha1.ListPoolsRequest, ...grpc.CallOption) (*v1alpha1.ListPoolsResponse, error)
	Filter(context.Context, *v1alpha1.ListPoolsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Pool, error]
}

type poolService struct {
	sdk iface.SDK
}

func NewPoolService(sdk iface.SDK) PoolService {
	return poolService{
		sdk: sdk,
	}
}

func (s poolService) Get(ctx context.Context, request *v1alpha1.GetPoolRequest, opts ...grpc.CallOption) (*v1alpha1.Pool, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewPoolServiceClient(con).Get(ctx, request, opts...)
}

func (s poolService) GetByName(ctx context.Context, request *v1alpha1.GetPoolByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Pool, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewPoolServiceClient(con).GetByName(ctx, request, opts...)
}

func (s poolService) List(ctx context.Context, request *v1alpha1.ListPoolsRequest, opts ...grpc.CallOption) (*v1alpha1.ListPoolsResponse, error) {
	address, err := s.sdk.Resolve(ctx, PoolServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewPoolServiceClient(con).List(ctx, request, opts...)
}

func (s poolService) Filter(ctx context.Context, request *v1alpha1.ListPoolsRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Pool, error] {
	req := proto.Clone(request).(*v1alpha1.ListPoolsRequest)
	return func(yield func(*v1alpha1.Pool, error) bool) {
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
