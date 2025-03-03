// Code generated by protoc-gen-gosdk. DO NOT EDIT.

package v1

import (
	context "context"
	conn "github.com/nebius/gosdk/conn"
	iface "github.com/nebius/gosdk/internal/iface"
	operations "github.com/nebius/gosdk/operations"
	grpcheader "github.com/nebius/gosdk/proto/fieldmask/grpcheader"
	v11 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	grpc "google.golang.org/grpc"
)

func init() {
	conn.ConventionResolverServiceIDToNameMap[FederationCertificateServiceID] = "cpl.iam"
}

func (s Services) FederationCertificate() FederationCertificateService {
	return NewFederationCertificateService(s.sdk)
}

const FederationCertificateServiceID conn.ServiceID = "nebius.iam.v1.FederationCertificateService"

type FederationCertificateService interface {
	Create(context.Context, *v1.CreateFederationCertificateRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetFederationCertificateRequest, ...grpc.CallOption) (*v1.FederationCertificate, error)
	ListByFederation(context.Context, *v1.ListFederationCertificateByFederationRequest, ...grpc.CallOption) (*v1.ListFederationCertificateResponse, error)
	Update(context.Context, *v1.UpdateFederationCertificateRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteFederationCertificateRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type federationCertificateService struct {
	sdk iface.SDK
}

func NewFederationCertificateService(sdk iface.SDK) FederationCertificateService {
	return federationCertificateService{
		sdk: sdk,
	}
}

func (s federationCertificateService) Create(ctx context.Context, request *v1.CreateFederationCertificateRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFederationCertificateServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s federationCertificateService) Get(ctx context.Context, request *v1.GetFederationCertificateRequest, opts ...grpc.CallOption) (*v1.FederationCertificate, error) {
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFederationCertificateServiceClient(con).Get(ctx, request, opts...)
}

func (s federationCertificateService) ListByFederation(ctx context.Context, request *v1.ListFederationCertificateByFederationRequest, opts ...grpc.CallOption) (*v1.ListFederationCertificateResponse, error) {
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFederationCertificateServiceClient(con).ListByFederation(ctx, request, opts...)
}

func (s federationCertificateService) Update(ctx context.Context, request *v1.UpdateFederationCertificateRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFederationCertificateServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s federationCertificateService) Delete(ctx context.Context, request *v1.DeleteFederationCertificateRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFederationCertificateServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s federationCertificateService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
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

func (s federationCertificateService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, FederationCertificateServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
