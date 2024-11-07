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

func (s Services) Scope() ScopeService {
	return NewScopeService(s.sdk)
}

const ScopeServiceID conn.ServiceID = "nebius.vpc.v1alpha1.ScopeService"

type ScopeService interface {
	Get(context.Context, *v1alpha1.GetScopeRequest, ...grpc.CallOption) (*v1alpha1.Scope, error)
	GetByName(context.Context, *v1alpha1.GetScopeByNameRequest, ...grpc.CallOption) (*v1alpha1.Scope, error)
	List(context.Context, *v1alpha1.ListScopesRequest, ...grpc.CallOption) (*v1alpha1.ListScopesResponse, error)
	Filter(context.Context, *v1alpha1.ListScopesRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Scope, error]
}

type scopeService struct {
	sdk iface.SDK
}

func NewScopeService(sdk iface.SDK) ScopeService {
	return scopeService{
		sdk: sdk,
	}
}

func (s scopeService) Get(ctx context.Context, request *v1alpha1.GetScopeRequest, opts ...grpc.CallOption) (*v1alpha1.Scope, error) {
	address, err := s.sdk.Resolve(ctx, ScopeServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewScopeServiceClient(con).Get(ctx, request, opts...)
}

func (s scopeService) GetByName(ctx context.Context, request *v1alpha1.GetScopeByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Scope, error) {
	address, err := s.sdk.Resolve(ctx, ScopeServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewScopeServiceClient(con).GetByName(ctx, request, opts...)
}

func (s scopeService) List(ctx context.Context, request *v1alpha1.ListScopesRequest, opts ...grpc.CallOption) (*v1alpha1.ListScopesResponse, error) {
	address, err := s.sdk.Resolve(ctx, ScopeServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewScopeServiceClient(con).List(ctx, request, opts...)
}

func (s scopeService) Filter(ctx context.Context, request *v1alpha1.ListScopesRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Scope, error] {
	req := proto.Clone(request).(*v1alpha1.ListScopesRequest)
	return func(yield func(*v1alpha1.Scope, error) bool) {
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
