// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	grpcheader "github.com/nebius/gosdk/fieldmask/grpcheader"
	iface "github.com/nebius/gosdk/internal/iface"
	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[InvitationServiceID] = "cpl.iam"
}

func (s Services) Invitation() InvitationService {
	return NewInvitationService(s.sdk)
}

const InvitationServiceID conn.ServiceID = "nebius.iam.v1.InvitationService"

type InvitationService interface {
	Create(context.Context, *v1.CreateInvitationRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetInvitationRequest, ...grpc.CallOption) (*v1.Invitation, error)
	List(context.Context, *v1.ListInvitationsRequest, ...grpc.CallOption) (*v1.ListInvitationsResponse, error)
	Filter(context.Context, *v1.ListInvitationsRequest, ...grpc.CallOption) iter.Seq2[*v1.Invitation, error]
	Delete(context.Context, *v1.DeleteInvitationRequest, ...grpc.CallOption) (operations.Operation, error)
	Update(context.Context, *v1.UpdateInvitationRequest, ...grpc.CallOption) (operations.Operation, error)
	Resend(context.Context, *v1.ResendInvitationRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type invitationService struct {
	sdk iface.SDK
}

func NewInvitationService(sdk iface.SDK) InvitationService {
	return invitationService{
		sdk: sdk,
	}
}

func (s invitationService) Create(ctx context.Context, request *v1.CreateInvitationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewInvitationServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s invitationService) Get(ctx context.Context, request *v1.GetInvitationRequest, opts ...grpc.CallOption) (*v1.Invitation, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewInvitationServiceClient(con).Get(ctx, request, opts...)
}

func (s invitationService) List(ctx context.Context, request *v1.ListInvitationsRequest, opts ...grpc.CallOption) (*v1.ListInvitationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewInvitationServiceClient(con).List(ctx, request, opts...)
}

func (s invitationService) Filter(ctx context.Context, request *v1.ListInvitationsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Invitation, error] {
	req := proto.Clone(request).(*v1.ListInvitationsRequest)
	return func(yield func(*v1.Invitation, error) bool) {
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

func (s invitationService) Delete(ctx context.Context, request *v1.DeleteInvitationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewInvitationServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s invitationService) Update(ctx context.Context, request *v1.UpdateInvitationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewInvitationServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s invitationService) Resend(ctx context.Context, request *v1.ResendInvitationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewInvitationServiceClient(con).Resend(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s invitationService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
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

func (s invitationService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, InvitationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
