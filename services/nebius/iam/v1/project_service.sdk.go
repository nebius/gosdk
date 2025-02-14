// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[ProjectServiceID] = "cpl.iam"
}

func (s Services) Project() ProjectService {
	return NewProjectService(s.sdk)
}

const ProjectServiceID conn.ServiceID = "nebius.iam.v1.ProjectService"

type ProjectService interface {
	Get(context.Context, *v1.GetProjectRequest, ...grpc.CallOption) (*v1.Container, error)
	GetByName(context.Context, *v1.GetProjectByNameRequest, ...grpc.CallOption) (*v1.Container, error)
	List(context.Context, *v1.ListProjectsRequest, ...grpc.CallOption) (*v1.ListProjectsResponse, error)
	Filter(context.Context, *v1.ListProjectsRequest, ...grpc.CallOption) iter.Seq2[*v1.Container, error]
}

type projectService struct {
	sdk iface.SDK
}

func NewProjectService(sdk iface.SDK) ProjectService {
	return projectService{
		sdk: sdk,
	}
}

func (s projectService) Get(ctx context.Context, request *v1.GetProjectRequest, opts ...grpc.CallOption) (*v1.Container, error) {
	address, err := s.sdk.Resolve(ctx, ProjectServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewProjectServiceClient(con).Get(ctx, request, opts...)
}

func (s projectService) GetByName(ctx context.Context, request *v1.GetProjectByNameRequest, opts ...grpc.CallOption) (*v1.Container, error) {
	address, err := s.sdk.Resolve(ctx, ProjectServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewProjectServiceClient(con).GetByName(ctx, request, opts...)
}

func (s projectService) List(ctx context.Context, request *v1.ListProjectsRequest, opts ...grpc.CallOption) (*v1.ListProjectsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ProjectServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewProjectServiceClient(con).List(ctx, request, opts...)
}

func (s projectService) Filter(ctx context.Context, request *v1.ListProjectsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Container, error] {
	req := proto.Clone(request).(*v1.ListProjectsRequest)
	return func(yield func(*v1.Container, error) bool) {
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
