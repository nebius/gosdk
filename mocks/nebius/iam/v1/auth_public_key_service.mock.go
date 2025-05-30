// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/iam/v1/auth_public_key_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/iam/v1/auth_public_key_service.sdk.go -package v1 -typed
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

// MockAuthPublicKeyService is a mock of AuthPublicKeyService interface.
type MockAuthPublicKeyService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthPublicKeyServiceMockRecorder
}

// MockAuthPublicKeyServiceMockRecorder is the mock recorder for MockAuthPublicKeyService.
type MockAuthPublicKeyServiceMockRecorder struct {
	mock *MockAuthPublicKeyService
}

// NewMockAuthPublicKeyService creates a new mock instance.
func NewMockAuthPublicKeyService(ctrl *gomock.Controller) *MockAuthPublicKeyService {
	mock := &MockAuthPublicKeyService{ctrl: ctrl}
	mock.recorder = &MockAuthPublicKeyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthPublicKeyService) EXPECT() *MockAuthPublicKeyServiceMockRecorder {
	return m.recorder
}

// Activate mocks base method.
func (m *MockAuthPublicKeyService) Activate(arg0 context.Context, arg1 *v10.ActivateAuthPublicKeyRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Activate", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Activate indicates an expected call of Activate.
func (mr *MockAuthPublicKeyServiceMockRecorder) Activate(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceActivateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Activate), varargs...)
	return &MockAuthPublicKeyServiceActivateCall{Call: call}
}

// MockAuthPublicKeyServiceActivateCall wrap *gomock.Call
type MockAuthPublicKeyServiceActivateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceActivateCall) Return(arg0 operations.Operation, arg1 error) *MockAuthPublicKeyServiceActivateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceActivateCall) Do(f func(context.Context, *v10.ActivateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceActivateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceActivateCall) DoAndReturn(f func(context.Context, *v10.ActivateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceActivateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Create mocks base method.
func (m *MockAuthPublicKeyService) Create(arg0 context.Context, arg1 *v10.CreateAuthPublicKeyRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockAuthPublicKeyServiceMockRecorder) Create(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceCreateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Create), varargs...)
	return &MockAuthPublicKeyServiceCreateCall{Call: call}
}

// MockAuthPublicKeyServiceCreateCall wrap *gomock.Call
type MockAuthPublicKeyServiceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceCreateCall) Return(arg0 operations.Operation, arg1 error) *MockAuthPublicKeyServiceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceCreateCall) Do(f func(context.Context, *v10.CreateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceCreateCall) DoAndReturn(f func(context.Context, *v10.CreateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Deactivate mocks base method.
func (m *MockAuthPublicKeyService) Deactivate(arg0 context.Context, arg1 *v10.DeactivateAuthPublicKeyRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Deactivate", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deactivate indicates an expected call of Deactivate.
func (mr *MockAuthPublicKeyServiceMockRecorder) Deactivate(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceDeactivateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Deactivate), varargs...)
	return &MockAuthPublicKeyServiceDeactivateCall{Call: call}
}

// MockAuthPublicKeyServiceDeactivateCall wrap *gomock.Call
type MockAuthPublicKeyServiceDeactivateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceDeactivateCall) Return(arg0 operations.Operation, arg1 error) *MockAuthPublicKeyServiceDeactivateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceDeactivateCall) Do(f func(context.Context, *v10.DeactivateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceDeactivateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceDeactivateCall) DoAndReturn(f func(context.Context, *v10.DeactivateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceDeactivateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockAuthPublicKeyService) Delete(arg0 context.Context, arg1 *v10.DeleteAuthPublicKeyRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockAuthPublicKeyServiceMockRecorder) Delete(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceDeleteCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Delete), varargs...)
	return &MockAuthPublicKeyServiceDeleteCall{Call: call}
}

// MockAuthPublicKeyServiceDeleteCall wrap *gomock.Call
type MockAuthPublicKeyServiceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceDeleteCall) Return(arg0 operations.Operation, arg1 error) *MockAuthPublicKeyServiceDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceDeleteCall) Do(f func(context.Context, *v10.DeleteAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceDeleteCall) DoAndReturn(f func(context.Context, *v10.DeleteAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockAuthPublicKeyService) Filter(arg0 context.Context, arg1 *v10.ListAuthPublicKeyRequest, arg2 ...grpc.CallOption) iter.Seq2[*v10.AuthPublicKey, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v10.AuthPublicKey, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockAuthPublicKeyServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Filter), varargs...)
	return &MockAuthPublicKeyServiceFilterCall{Call: call}
}

// MockAuthPublicKeyServiceFilterCall wrap *gomock.Call
type MockAuthPublicKeyServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceFilterCall) Return(arg0 iter.Seq2[*v10.AuthPublicKey, error]) *MockAuthPublicKeyServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceFilterCall) Do(f func(context.Context, *v10.ListAuthPublicKeyRequest, ...grpc.CallOption) iter.Seq2[*v10.AuthPublicKey, error]) *MockAuthPublicKeyServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceFilterCall) DoAndReturn(f func(context.Context, *v10.ListAuthPublicKeyRequest, ...grpc.CallOption) iter.Seq2[*v10.AuthPublicKey, error]) *MockAuthPublicKeyServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockAuthPublicKeyService) Get(arg0 context.Context, arg1 *v10.GetAuthPublicKeyRequest, arg2 ...grpc.CallOption) (*v10.AuthPublicKey, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v10.AuthPublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAuthPublicKeyServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Get), varargs...)
	return &MockAuthPublicKeyServiceGetCall{Call: call}
}

// MockAuthPublicKeyServiceGetCall wrap *gomock.Call
type MockAuthPublicKeyServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceGetCall) Return(arg0 *v10.AuthPublicKey, arg1 error) *MockAuthPublicKeyServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceGetCall) Do(f func(context.Context, *v10.GetAuthPublicKeyRequest, ...grpc.CallOption) (*v10.AuthPublicKey, error)) *MockAuthPublicKeyServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceGetCall) DoAndReturn(f func(context.Context, *v10.GetAuthPublicKeyRequest, ...grpc.CallOption) (*v10.AuthPublicKey, error)) *MockAuthPublicKeyServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOperation mocks base method.
func (m *MockAuthPublicKeyService) GetOperation(arg0 context.Context, arg1 *v1.GetOperationRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockAuthPublicKeyServiceMockRecorder) GetOperation(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceGetOperationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockAuthPublicKeyService)(nil).GetOperation), varargs...)
	return &MockAuthPublicKeyServiceGetOperationCall{Call: call}
}

// MockAuthPublicKeyServiceGetOperationCall wrap *gomock.Call
type MockAuthPublicKeyServiceGetOperationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceGetOperationCall) Return(arg0 operations.Operation, arg1 error) *MockAuthPublicKeyServiceGetOperationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceGetOperationCall) Do(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceGetOperationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceGetOperationCall) DoAndReturn(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceGetOperationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockAuthPublicKeyService) List(arg0 context.Context, arg1 *v10.ListAuthPublicKeyRequest, arg2 ...grpc.CallOption) (*v10.ListAuthPublicKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v10.ListAuthPublicKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAuthPublicKeyServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthPublicKeyService)(nil).List), varargs...)
	return &MockAuthPublicKeyServiceListCall{Call: call}
}

// MockAuthPublicKeyServiceListCall wrap *gomock.Call
type MockAuthPublicKeyServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceListCall) Return(arg0 *v10.ListAuthPublicKeyResponse, arg1 error) *MockAuthPublicKeyServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceListCall) Do(f func(context.Context, *v10.ListAuthPublicKeyRequest, ...grpc.CallOption) (*v10.ListAuthPublicKeyResponse, error)) *MockAuthPublicKeyServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceListCall) DoAndReturn(f func(context.Context, *v10.ListAuthPublicKeyRequest, ...grpc.CallOption) (*v10.ListAuthPublicKeyResponse, error)) *MockAuthPublicKeyServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListByAccount mocks base method.
func (m *MockAuthPublicKeyService) ListByAccount(arg0 context.Context, arg1 *v10.ListAuthPublicKeyByAccountRequest, arg2 ...grpc.CallOption) (*v10.ListAuthPublicKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByAccount", varargs...)
	ret0, _ := ret[0].(*v10.ListAuthPublicKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByAccount indicates an expected call of ListByAccount.
func (mr *MockAuthPublicKeyServiceMockRecorder) ListByAccount(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceListByAccountCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByAccount", reflect.TypeOf((*MockAuthPublicKeyService)(nil).ListByAccount), varargs...)
	return &MockAuthPublicKeyServiceListByAccountCall{Call: call}
}

// MockAuthPublicKeyServiceListByAccountCall wrap *gomock.Call
type MockAuthPublicKeyServiceListByAccountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceListByAccountCall) Return(arg0 *v10.ListAuthPublicKeyResponse, arg1 error) *MockAuthPublicKeyServiceListByAccountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceListByAccountCall) Do(f func(context.Context, *v10.ListAuthPublicKeyByAccountRequest, ...grpc.CallOption) (*v10.ListAuthPublicKeyResponse, error)) *MockAuthPublicKeyServiceListByAccountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceListByAccountCall) DoAndReturn(f func(context.Context, *v10.ListAuthPublicKeyByAccountRequest, ...grpc.CallOption) (*v10.ListAuthPublicKeyResponse, error)) *MockAuthPublicKeyServiceListByAccountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperations mocks base method.
func (m *MockAuthPublicKeyService) ListOperations(arg0 context.Context, arg1 *v1.ListOperationsRequest, arg2 ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
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
func (mr *MockAuthPublicKeyServiceMockRecorder) ListOperations(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceListOperationsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockAuthPublicKeyService)(nil).ListOperations), varargs...)
	return &MockAuthPublicKeyServiceListOperationsCall{Call: call}
}

// MockAuthPublicKeyServiceListOperationsCall wrap *gomock.Call
type MockAuthPublicKeyServiceListOperationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceListOperationsCall) Return(arg0 *v1.ListOperationsResponse, arg1 error) *MockAuthPublicKeyServiceListOperationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceListOperationsCall) Do(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockAuthPublicKeyServiceListOperationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceListOperationsCall) DoAndReturn(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockAuthPublicKeyServiceListOperationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *MockAuthPublicKeyService) Update(arg0 context.Context, arg1 *v10.UpdateAuthPublicKeyRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockAuthPublicKeyServiceMockRecorder) Update(arg0, arg1 any, arg2 ...any) *MockAuthPublicKeyServiceUpdateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthPublicKeyService)(nil).Update), varargs...)
	return &MockAuthPublicKeyServiceUpdateCall{Call: call}
}

// MockAuthPublicKeyServiceUpdateCall wrap *gomock.Call
type MockAuthPublicKeyServiceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAuthPublicKeyServiceUpdateCall) Return(arg0 operations.Operation, arg1 error) *MockAuthPublicKeyServiceUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAuthPublicKeyServiceUpdateCall) Do(f func(context.Context, *v10.UpdateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAuthPublicKeyServiceUpdateCall) DoAndReturn(f func(context.Context, *v10.UpdateAuthPublicKeyRequest, ...grpc.CallOption) (operations.Operation, error)) *MockAuthPublicKeyServiceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
