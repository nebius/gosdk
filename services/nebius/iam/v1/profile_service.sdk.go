package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[ProfileServiceID] = "cpl.iam"
}

func (s Services) Profile() ProfileService {
	return NewProfileService(s.sdk)
}

const ProfileServiceID conn.ServiceID = "nebius.iam.v1.ProfileService"

type ProfileService interface {
	Get(context.Context, *v1.GetProfileRequest, ...grpc.CallOption) (*v1.GetProfileResponse, error)
}

type profileService struct {
	sdk iface.SDK
}

func NewProfileService(sdk iface.SDK) ProfileService {
	return profileService{
		sdk: sdk,
	}
}

func (s profileService) Get(ctx context.Context, request *v1.GetProfileRequest, opts ...grpc.CallOption) (*v1.GetProfileResponse, error) {
	address, err := s.sdk.Resolve(ctx, ProfileServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewProfileServiceClient(con).Get(ctx, request, opts...)
}
