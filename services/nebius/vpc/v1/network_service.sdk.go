package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/vpc/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Network() NetworkService {
	return NewNetworkService(s.sdk)
}

const NetworkServiceID conn.ServiceID = "nebius.vpc.v1.NetworkService"

type NetworkService interface {
	Get(context.Context, *v1.GetNetworkRequest, ...grpc.CallOption) (*v1.Network, error)
	GetByName(context.Context, *v1.GetNetworkByNameRequest, ...grpc.CallOption) (*v1.Network, error)
	List(context.Context, *v1.ListNetworksRequest, ...grpc.CallOption) (*v1.ListNetworksResponse, error)
	Filter(context.Context, *v1.ListNetworksRequest, ...grpc.CallOption) iter.Seq2[*v1.Network, error]
}

type networkService struct {
	sdk iface.SDK
}

func NewNetworkService(sdk iface.SDK) NetworkService {
	return networkService{
		sdk: sdk,
	}
}

func (s networkService) Get(ctx context.Context, request *v1.GetNetworkRequest, opts ...grpc.CallOption) (*v1.Network, error) {
	address, err := s.sdk.Resolve(ctx, NetworkServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewNetworkServiceClient(con).Get(ctx, request, opts...)
}

func (s networkService) GetByName(ctx context.Context, request *v1.GetNetworkByNameRequest, opts ...grpc.CallOption) (*v1.Network, error) {
	address, err := s.sdk.Resolve(ctx, NetworkServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewNetworkServiceClient(con).GetByName(ctx, request, opts...)
}

func (s networkService) List(ctx context.Context, request *v1.ListNetworksRequest, opts ...grpc.CallOption) (*v1.ListNetworksResponse, error) {
	address, err := s.sdk.Resolve(ctx, NetworkServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewNetworkServiceClient(con).List(ctx, request, opts...)
}

func (s networkService) Filter(ctx context.Context, request *v1.ListNetworksRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Network, error] {
	req := proto.Clone(request).(*v1.ListNetworksRequest)
	return func(yield func(*v1.Network, error) bool) {
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
