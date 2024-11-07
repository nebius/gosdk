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

func (s Services) Subnet() SubnetService {
	return NewSubnetService(s.sdk)
}

const SubnetServiceID conn.ServiceID = "nebius.vpc.v1alpha1.SubnetService"

type SubnetService interface {
	Get(context.Context, *v1alpha1.GetSubnetRequest, ...grpc.CallOption) (*v1alpha1.Subnet, error)
	GetByName(context.Context, *v1alpha1.GetSubnetByNameRequest, ...grpc.CallOption) (*v1alpha1.Subnet, error)
	List(context.Context, *v1alpha1.ListSubnetsRequest, ...grpc.CallOption) (*v1alpha1.ListSubnetsResponse, error)
	Filter(context.Context, *v1alpha1.ListSubnetsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Subnet, error]
	ListByNetwork(context.Context, *v1alpha1.ListSubnetsByNetworkRequest, ...grpc.CallOption) (*v1alpha1.ListSubnetsResponse, error)
}

type subnetService struct {
	sdk iface.SDK
}

func NewSubnetService(sdk iface.SDK) SubnetService {
	return subnetService{
		sdk: sdk,
	}
}

func (s subnetService) Get(ctx context.Context, request *v1alpha1.GetSubnetRequest, opts ...grpc.CallOption) (*v1alpha1.Subnet, error) {
	address, err := s.sdk.Resolve(ctx, SubnetServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSubnetServiceClient(con).Get(ctx, request, opts...)
}

func (s subnetService) GetByName(ctx context.Context, request *v1alpha1.GetSubnetByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Subnet, error) {
	address, err := s.sdk.Resolve(ctx, SubnetServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSubnetServiceClient(con).GetByName(ctx, request, opts...)
}

func (s subnetService) List(ctx context.Context, request *v1alpha1.ListSubnetsRequest, opts ...grpc.CallOption) (*v1alpha1.ListSubnetsResponse, error) {
	address, err := s.sdk.Resolve(ctx, SubnetServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSubnetServiceClient(con).List(ctx, request, opts...)
}

func (s subnetService) Filter(ctx context.Context, request *v1alpha1.ListSubnetsRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Subnet, error] {
	req := proto.Clone(request).(*v1alpha1.ListSubnetsRequest)
	return func(yield func(*v1alpha1.Subnet, error) bool) {
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

func (s subnetService) ListByNetwork(ctx context.Context, request *v1alpha1.ListSubnetsByNetworkRequest, opts ...grpc.CallOption) (*v1alpha1.ListSubnetsResponse, error) {
	address, err := s.sdk.Resolve(ctx, SubnetServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewSubnetServiceClient(con).ListByNetwork(ctx, request, opts...)
}
