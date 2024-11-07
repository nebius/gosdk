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

func (s Services) Preset() PresetService {
	return NewPresetService(s.sdk)
}

const PresetServiceID conn.ServiceID = "nebius.msp.v1alpha1.resource.PresetService"

type PresetService interface {
	List(context.Context, *resource.ListPresetsRequest, ...grpc.CallOption) (*resource.ListPresetsResponse, error)
	Filter(context.Context, *resource.ListPresetsRequest, ...grpc.CallOption) iter.Seq2[*resource.Preset, error]
}

type presetService struct {
	sdk iface.SDK
}

func NewPresetService(sdk iface.SDK) PresetService {
	return presetService{
		sdk: sdk,
	}
}

func (s presetService) List(ctx context.Context, request *resource.ListPresetsRequest, opts ...grpc.CallOption) (*resource.ListPresetsResponse, error) {
	address, err := s.sdk.Resolve(ctx, PresetServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return resource.NewPresetServiceClient(con).List(ctx, request, opts...)
}

func (s presetService) Filter(ctx context.Context, request *resource.ListPresetsRequest, opts ...grpc.CallOption) iter.Seq2[*resource.Preset, error] {
	req := proto.Clone(request).(*resource.ListPresetsRequest)
	return func(yield func(*resource.Preset, error) bool) {
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
