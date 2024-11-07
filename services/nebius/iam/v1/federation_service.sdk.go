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
	conn.ConventionResolverServiceIDToNameMap[FederationServiceID] = "cpl.iam"
}

func (s Services) Federation() FederationService {
	return NewFederationService(s.sdk)
}

const FederationServiceID conn.ServiceID = "nebius.iam.v1.FederationService"

type FederationService interface {
	Create(context.Context, *v1.CreateFederationRequest, ...grpc.CallOption) (operations.Operation, error)
	Get(context.Context, *v1.GetFederationRequest, ...grpc.CallOption) (*v1.Federation, error)
	GetByName(context.Context, *v11.GetByNameRequest, ...grpc.CallOption) (*v1.Federation, error)
	List(context.Context, *v1.ListFederationsRequest, ...grpc.CallOption) (*v1.ListFederationsResponse, error)
	Filter(context.Context, *v1.ListFederationsRequest, ...grpc.CallOption) iter.Seq2[*v1.Federation, error]
	Update(context.Context, *v1.UpdateFederationRequest, ...grpc.CallOption) (operations.Operation, error)
	Delete(context.Context, *v1.DeleteFederationRequest, ...grpc.CallOption) (operations.Operation, error)
	GetOperation(context.Context, *v11.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)
	ListOperations(context.Context, *v11.ListOperationsRequest, ...grpc.CallOption) (*v11.ListOperationsResponse, error)
}

type federationService struct {
	sdk iface.SDK
}

func NewFederationService(sdk iface.SDK) FederationService {
	return federationService{
		sdk: sdk,
	}
}

func (s federationService) Create(ctx context.Context, request *v1.CreateFederationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFederationServiceClient(con).Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s federationService) Get(ctx context.Context, request *v1.GetFederationRequest, opts ...grpc.CallOption) (*v1.Federation, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFederationServiceClient(con).Get(ctx, request, opts...)
}

func (s federationService) GetByName(ctx context.Context, request *v11.GetByNameRequest, opts ...grpc.CallOption) (*v1.Federation, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFederationServiceClient(con).GetByName(ctx, request, opts...)
}

func (s federationService) List(ctx context.Context, request *v1.ListFederationsRequest, opts ...grpc.CallOption) (*v1.ListFederationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v1.NewFederationServiceClient(con).List(ctx, request, opts...)
}

func (s federationService) Filter(ctx context.Context, request *v1.ListFederationsRequest, opts ...grpc.CallOption) iter.Seq2[*v1.Federation, error] {
	req := proto.Clone(request).(*v1.ListFederationsRequest)
	return func(yield func(*v1.Federation, error) bool) {
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

func (s federationService) Update(ctx context.Context, request *v1.UpdateFederationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	ctx, err := grpcheader.EnsureMessageResetMaskInOutgoingContext(ctx, request)
	if err != nil {
		return nil, err
	}
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFederationServiceClient(con).Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s federationService) Delete(ctx context.Context, request *v1.DeleteFederationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	op, err := v1.NewFederationServiceClient(con).Delete(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return operations.New(op, v11.NewOperationServiceClient(con))
}

func (s federationService) GetOperation(ctx context.Context, request *v11.GetOperationRequest, opts ...grpc.CallOption) (operations.Operation, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
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

func (s federationService) ListOperations(ctx context.Context, request *v11.ListOperationsRequest, opts ...grpc.CallOption) (*v11.ListOperationsResponse, error) {
	address, err := s.sdk.Resolve(ctx, FederationServiceID)
	if err != nil {
		return nil, err
	}
	con, err := s.sdk.Dial(ctx, address)
	if err != nil {
		return nil, err
	}
	return v11.NewOperationServiceClient(con).List(ctx, request, opts...)
}
