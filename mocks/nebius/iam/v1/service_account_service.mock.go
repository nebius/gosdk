// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/iam/v1/service_account_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/iam/v1/service_account_service.sdk.go -package v1 -typed
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v10 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockServiceAccountService is a mock of ServiceAccountService interface.
type MockServiceAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountServiceMockRecorder
}

// MockServiceAccountServiceMockRecorder is the mock recorder for MockServiceAccountService.
type MockServiceAccountServiceMockRecorder struct {
	mock *MockServiceAccountService
}

// NewMockServiceAccountService creates a new mock instance.
func NewMockServiceAccountService(ctrl *gomock.Controller) *MockServiceAccountService {
	mock := &MockServiceAccountService{ctrl: ctrl}
	mock.recorder = &MockServiceAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountService) EXPECT() *MockServiceAccountServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockServiceAccountService) Create(arg0 context.Context, arg1 *v10.CreateServiceAccountRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceAccountServiceMockRecorder) Create(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceCreateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceAccountService)(nil).Create), varargs...)
	return &MockServiceAccountServiceCreateCall{Call: call}
}

// MockServiceAccountServiceCreateCall wrap *gomock.Call
type MockServiceAccountServiceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceCreateCall) Return(arg0 operations.Operation, arg1 error) *MockServiceAccountServiceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceCreateCall) Do(f func(context.Context, *v10.CreateServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceCreateCall) DoAndReturn(f func(context.Context, *v10.CreateServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockServiceAccountService) Delete(arg0 context.Context, arg1 *v10.DeleteServiceAccountRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceAccountServiceMockRecorder) Delete(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceDeleteCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceAccountService)(nil).Delete), varargs...)
	return &MockServiceAccountServiceDeleteCall{Call: call}
}

// MockServiceAccountServiceDeleteCall wrap *gomock.Call
type MockServiceAccountServiceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceDeleteCall) Return(arg0 operations.Operation, arg1 error) *MockServiceAccountServiceDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceDeleteCall) Do(f func(context.Context, *v10.DeleteServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceDeleteCall) DoAndReturn(f func(context.Context, *v10.DeleteServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockServiceAccountService) Filter(arg0 context.Context, arg1 *v10.ListServiceAccountRequest, arg2 ...grpc.CallOption) iter.Seq2[*v10.ServiceAccount, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v10.ServiceAccount, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockServiceAccountServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockServiceAccountService)(nil).Filter), varargs...)
	return &MockServiceAccountServiceFilterCall{Call: call}
}

// MockServiceAccountServiceFilterCall wrap *gomock.Call
type MockServiceAccountServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceFilterCall) Return(arg0 iter.Seq2[*v10.ServiceAccount, error]) *MockServiceAccountServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceFilterCall) Do(f func(context.Context, *v10.ListServiceAccountRequest, ...grpc.CallOption) iter.Seq2[*v10.ServiceAccount, error]) *MockServiceAccountServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceFilterCall) DoAndReturn(f func(context.Context, *v10.ListServiceAccountRequest, ...grpc.CallOption) iter.Seq2[*v10.ServiceAccount, error]) *MockServiceAccountServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockServiceAccountService) Get(arg0 context.Context, arg1 *v10.GetServiceAccountRequest, arg2 ...grpc.CallOption) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceAccountServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceAccountService)(nil).Get), varargs...)
	return &MockServiceAccountServiceGetCall{Call: call}
}

// MockServiceAccountServiceGetCall wrap *gomock.Call
type MockServiceAccountServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceGetCall) Return(arg0 *v10.ServiceAccount, arg1 error) *MockServiceAccountServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceGetCall) Do(f func(context.Context, *v10.GetServiceAccountRequest, ...grpc.CallOption) (*v10.ServiceAccount, error)) *MockServiceAccountServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceGetCall) DoAndReturn(f func(context.Context, *v10.GetServiceAccountRequest, ...grpc.CallOption) (*v10.ServiceAccount, error)) *MockServiceAccountServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByName mocks base method.
func (m *MockServiceAccountService) GetByName(arg0 context.Context, arg1 *v10.GetServiceAccountByNameRequest, arg2 ...grpc.CallOption) (*v10.ServiceAccount, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByName", varargs...)
	ret0, _ := ret[0].(*v10.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockServiceAccountServiceMockRecorder) GetByName(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceGetByNameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockServiceAccountService)(nil).GetByName), varargs...)
	return &MockServiceAccountServiceGetByNameCall{Call: call}
}

// MockServiceAccountServiceGetByNameCall wrap *gomock.Call
type MockServiceAccountServiceGetByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceGetByNameCall) Return(arg0 *v10.ServiceAccount, arg1 error) *MockServiceAccountServiceGetByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceGetByNameCall) Do(f func(context.Context, *v10.GetServiceAccountByNameRequest, ...grpc.CallOption) (*v10.ServiceAccount, error)) *MockServiceAccountServiceGetByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceGetByNameCall) DoAndReturn(f func(context.Context, *v10.GetServiceAccountByNameRequest, ...grpc.CallOption) (*v10.ServiceAccount, error)) *MockServiceAccountServiceGetByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOperation mocks base method.
func (m *MockServiceAccountService) GetOperation(arg0 context.Context, arg1 *v1.GetOperationRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperation", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockServiceAccountServiceMockRecorder) GetOperation(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceGetOperationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockServiceAccountService)(nil).GetOperation), varargs...)
	return &MockServiceAccountServiceGetOperationCall{Call: call}
}

// MockServiceAccountServiceGetOperationCall wrap *gomock.Call
type MockServiceAccountServiceGetOperationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceGetOperationCall) Return(arg0 operations.Operation, arg1 error) *MockServiceAccountServiceGetOperationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceGetOperationCall) Do(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceGetOperationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceGetOperationCall) DoAndReturn(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceGetOperationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockServiceAccountService) List(arg0 context.Context, arg1 *v10.ListServiceAccountRequest, arg2 ...grpc.CallOption) (*v10.ListServiceAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v10.ListServiceAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceAccountServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceAccountService)(nil).List), varargs...)
	return &MockServiceAccountServiceListCall{Call: call}
}

// MockServiceAccountServiceListCall wrap *gomock.Call
type MockServiceAccountServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceListCall) Return(arg0 *v10.ListServiceAccountResponse, arg1 error) *MockServiceAccountServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceListCall) Do(f func(context.Context, *v10.ListServiceAccountRequest, ...grpc.CallOption) (*v10.ListServiceAccountResponse, error)) *MockServiceAccountServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceListCall) DoAndReturn(f func(context.Context, *v10.ListServiceAccountRequest, ...grpc.CallOption) (*v10.ListServiceAccountResponse, error)) *MockServiceAccountServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperations mocks base method.
func (m *MockServiceAccountService) ListOperations(arg0 context.Context, arg1 *v1.ListOperationsRequest, arg2 ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*v1.ListOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockServiceAccountServiceMockRecorder) ListOperations(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceListOperationsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockServiceAccountService)(nil).ListOperations), varargs...)
	return &MockServiceAccountServiceListOperationsCall{Call: call}
}

// MockServiceAccountServiceListOperationsCall wrap *gomock.Call
type MockServiceAccountServiceListOperationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceListOperationsCall) Return(arg0 *v1.ListOperationsResponse, arg1 error) *MockServiceAccountServiceListOperationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceListOperationsCall) Do(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockServiceAccountServiceListOperationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceListOperationsCall) DoAndReturn(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockServiceAccountServiceListOperationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *MockServiceAccountService) Update(arg0 context.Context, arg1 *v10.UpdateServiceAccountRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockServiceAccountServiceMockRecorder) Update(arg0, arg1 any, arg2 ...any) *MockServiceAccountServiceUpdateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceAccountService)(nil).Update), varargs...)
	return &MockServiceAccountServiceUpdateCall{Call: call}
}

// MockServiceAccountServiceUpdateCall wrap *gomock.Call
type MockServiceAccountServiceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockServiceAccountServiceUpdateCall) Return(arg0 operations.Operation, arg1 error) *MockServiceAccountServiceUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockServiceAccountServiceUpdateCall) Do(f func(context.Context, *v10.UpdateServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockServiceAccountServiceUpdateCall) DoAndReturn(f func(context.Context, *v10.UpdateServiceAccountRequest, ...grpc.CallOption) (operations.Operation, error)) *MockServiceAccountServiceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
