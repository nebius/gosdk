// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/registry/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[ArtifactServiceID] = "registry"
}

func (s Services) Artifact() ArtifactService {
	return NewArtifactService(s.sdk)
}

const ArtifactServiceID conn.ServiceID = "nebius.registry.v1.ArtifactService"

type ArtifactService interface {
	Get(context.Context, *v1.GetArtifactRequest, ...grpc.CallOption) (*v1.Artifact, error)
	List(context.Context, *v1.ListArtifactsRequest, ...grpc.CallOption) (*v1.ListArtifactsResponse, error)
	Filter(context.Context, *v1.ListArtifactsRequest, ...grpc.CallOption) iter.Seq2[*v1.Artifact, error]
	Delete(context.Context, *v1.DeleteArtifactRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type artifactService struct {
	sdk iface.SDK
}

func NewArtifactService(sdk iface.SDK) ArtifactService {
	return artifactService{
		sdk: sdk,
	}
}

func (s artifactService) Get(ctx context.Context, request *v1.GetArtifactRequest, opts ...grpc.CallOption) (*v1.Artifact, error) {
	address, err := s.sdk.Resolve(ctx, ArtifactServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewArtifactServiceClient(con).Get(ctx, request, opts...)
}

func (s artifactService) List(ctx context.Context, request *v1.ListArtifactsRequest, opts ...grpc.CallOption) (*v1.ListArtifactsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ArtifactServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewArtifactServiceClient(con).List(ctx, request, opts...)
}

func (s artifactService) Filter(ctx context.Context, request *v1.ListArtifactsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Artifact, error] {
	req := proto.Clone(request).(*v1.ListArtifactsRequest)
	return func(yield func(*v1.Artifact, error) bool) {
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

func (s artifactService) Delete(ctx context.Context, request *v1.DeleteArtifactRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ArtifactServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewArtifactServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s artifactService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, ArtifactServiceID)
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

func (s artifactService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ArtifactServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
