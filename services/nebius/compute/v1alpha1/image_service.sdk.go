package v1alpha1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1alpha11 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/compute/v1alpha1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func (s Services) Image() ImageService {
	return NewImageService(s.sdk)
}

const ImageServiceID conn.ServiceID = "nebius.compute.v1alpha1.ImageService"

type ImageService interface {
	Get(context.Context, *v1alpha1.GetImageRequest, ...grpc.CallOption) (*v1alpha1.Image, error)
	GetByName(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v1alpha1.Image, error)
	GetLatestByFamily(context.Context, *v1alpha1.GetImageLatestByFamilyRequest, ...grpc.CallOption) (*v1alpha1.Image, error)
	List(context.Context, *v1alpha1.ListImagesRequest, ...grpc.CallOption) (*v1alpha1.ListImagesResponse, error)
	Filter(context.Context, *v1alpha1.ListImagesRequest, ...grpc.CallOption) iter.Seq2[*v1alpha1.Image, error]
	ListOperationsByParent(context.Context, *v1alpha11.ListOperationsByParentRequest, ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error)
}

type imageService struct {
	sdk iface.SDK
}

func NewImageService(sdk iface.SDK) ImageService {
	return imageService{
		sdk: sdk,
	}
}

func (s imageService) Get(ctx context.Context, request *v1alpha1.GetImageRequest, opts ...grpc.CallOption) (*v1alpha1.Image, error) {
	address, err := s.sdk.Resolve(ctx, ImageServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewImageServiceClient(con).Get(ctx, request, opts...)
}

func (s imageService) GetByName(ctx context.Context, request *v1.GetByNameRequest, opts ...grpc.CallOption) (*v1alpha1.Image, error) {
	address, err := s.sdk.Resolve(ctx, ImageServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewImageServiceClient(con).GetByName(ctx, request, opts...)
}

func (s imageService) GetLatestByFamily(ctx context.Context, request *v1alpha1.GetImageLatestByFamilyRequest, opts ...grpc.CallOption) (*v1alpha1.Image, error) {
	address, err := s.sdk.Resolve(ctx, ImageServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewImageServiceClient(con).GetLatestByFamily(ctx, request, opts...)
}

func (s imageService) List(ctx context.Context, request *v1alpha1.ListImagesRequest, opts ...grpc.CallOption) (*v1alpha1.ListImagesResponse, error) {
	address, err := s.sdk.Resolve(ctx, ImageServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewImageServiceClient(con).List(ctx, request, opts...)
}

func (s imageService) Filter(ctx context.Context, request *v1alpha1.ListImagesRequest, opts ...grpc.CallOption) iter.Seq2[*v1alpha1.Image, error] {
	req := proto.Clone(request).(*v1alpha1.ListImagesRequest)
	return func(yield func(*v1alpha1.Image, error) bool) {
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

func (s imageService) ListOperationsByParent(ctx context.Context, request *v1alpha11.ListOperationsByParentRequest, opts ...grpc.CallOption) (*v1alpha11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, ImageServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1alpha1.NewImageServiceClient(con).ListOperationsByParent(ctx, request, opts...)
}
