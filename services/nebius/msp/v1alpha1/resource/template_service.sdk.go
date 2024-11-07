package resource

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	resource "github.com/nebius/gosdk/proto/nebius/msp/v1alpha1/resource"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Template() TemplateService {
	return NewTemplateService(s.sdk)
}

const TemplateServiceID conn.ServiceID = "nebius.msp.v1alpha1.resource.TemplateService"

type TemplateService interface {
	List(context.Context, *resource.ListTemplatesRequest, ...grpc.CallOption) (*resource.ListTemplatesResponse, error)
	Filter(context.Context, *resource.ListTemplatesRequest, ...grpc.CallOption) iter.Seq2[*resource.Template, error]
}

type templateService struct {
	sdk iface.SDK
}

func NewTemplateService(sdk iface.SDK) TemplateService {
	return templateService{
		sdk: sdk,
	}
}

func (s templateService) List(ctx context.Context, request *resource.ListTemplatesRequest, opts ...grpc.CallOption) (*resource.ListTemplatesResponse, error) {
	address, err := s.sdk.Resolve(ctx, TemplateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return resource.NewTemplateServiceClient(con).List(ctx, request, opts...)
}

func (s templateService) Filter(ctx context.Context, request *resource.ListTemplatesRequest, opts ...grpc.CallOption) iter.Seq2[*resource.Template, error] {
	req := proto.Clone(request).(*resource.ListTemplatesRequest)
	return func(yield func(*resource.Template, error) bool) {
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
