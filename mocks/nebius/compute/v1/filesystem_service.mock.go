// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/compute/v1/filesystem_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/compute/v1/filesystem_service.sdk.go -package v1 -typed
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v10 "github.com/nebius/gosdk/proto/nebius/compute/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockFilesystemService is a mock of FilesystemService interface.
type MockFilesystemService struct {
	ctrl     *gomock.Controller
	recorder *MockFilesystemServiceMockRecorder
}

// MockFilesystemServiceMockRecorder is the mock recorder for MockFilesystemService.
type MockFilesystemServiceMockRecorder struct {
	mock *MockFilesystemService
}

// NewMockFilesystemService creates a new mock instance.
func NewMockFilesystemService(ctrl *gomock.Controller) *MockFilesystemService {
	mock := &MockFilesystemService{ctrl: ctrl}
	mock.recorder = &MockFilesystemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilesystemService) EXPECT() *MockFilesystemServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFilesystemService) Create(arg0 context.Context, arg1 *v10.CreateFilesystemRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockFilesystemServiceMockRecorder) Create(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceCreateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFilesystemService)(nil).Create), varargs...)
	return &MockFilesystemServiceCreateCall{Call: call}
}

// MockFilesystemServiceCreateCall wrap *gomock.Call
type MockFilesystemServiceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceCreateCall) Return(arg0 operations.Operation, arg1 error) *MockFilesystemServiceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceCreateCall) Do(f func(context.Context, *v10.CreateFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceCreateCall) DoAndReturn(f func(context.Context, *v10.CreateFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockFilesystemService) Delete(arg0 context.Context, arg1 *v10.DeleteFilesystemRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockFilesystemServiceMockRecorder) Delete(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceDeleteCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFilesystemService)(nil).Delete), varargs...)
	return &MockFilesystemServiceDeleteCall{Call: call}
}

// MockFilesystemServiceDeleteCall wrap *gomock.Call
type MockFilesystemServiceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceDeleteCall) Return(arg0 operations.Operation, arg1 error) *MockFilesystemServiceDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceDeleteCall) Do(f func(context.Context, *v10.DeleteFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceDeleteCall) DoAndReturn(f func(context.Context, *v10.DeleteFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockFilesystemService) Filter(arg0 context.Context, arg1 *v10.ListFilesystemsRequest, arg2 ...grpc.CallOption) iter.Seq2[*v10.Filesystem, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v10.Filesystem, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockFilesystemServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockFilesystemService)(nil).Filter), varargs...)
	return &MockFilesystemServiceFilterCall{Call: call}
}

// MockFilesystemServiceFilterCall wrap *gomock.Call
type MockFilesystemServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceFilterCall) Return(arg0 iter.Seq2[*v10.Filesystem, error]) *MockFilesystemServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceFilterCall) Do(f func(context.Context, *v10.ListFilesystemsRequest, ...grpc.CallOption) iter.Seq2[*v10.Filesystem, error]) *MockFilesystemServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceFilterCall) DoAndReturn(f func(context.Context, *v10.ListFilesystemsRequest, ...grpc.CallOption) iter.Seq2[*v10.Filesystem, error]) *MockFilesystemServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockFilesystemService) Get(arg0 context.Context, arg1 *v10.GetFilesystemRequest, arg2 ...grpc.CallOption) (*v10.Filesystem, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v10.Filesystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFilesystemServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFilesystemService)(nil).Get), varargs...)
	return &MockFilesystemServiceGetCall{Call: call}
}

// MockFilesystemServiceGetCall wrap *gomock.Call
type MockFilesystemServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceGetCall) Return(arg0 *v10.Filesystem, arg1 error) *MockFilesystemServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceGetCall) Do(f func(context.Context, *v10.GetFilesystemRequest, ...grpc.CallOption) (*v10.Filesystem, error)) *MockFilesystemServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceGetCall) DoAndReturn(f func(context.Context, *v10.GetFilesystemRequest, ...grpc.CallOption) (*v10.Filesystem, error)) *MockFilesystemServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByName mocks base method.
func (m *MockFilesystemService) GetByName(arg0 context.Context, arg1 *v1.GetByNameRequest, arg2 ...grpc.CallOption) (*v10.Filesystem, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByName", varargs...)
	ret0, _ := ret[0].(*v10.Filesystem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockFilesystemServiceMockRecorder) GetByName(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceGetByNameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockFilesystemService)(nil).GetByName), varargs...)
	return &MockFilesystemServiceGetByNameCall{Call: call}
}

// MockFilesystemServiceGetByNameCall wrap *gomock.Call
type MockFilesystemServiceGetByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceGetByNameCall) Return(arg0 *v10.Filesystem, arg1 error) *MockFilesystemServiceGetByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceGetByNameCall) Do(f func(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v10.Filesystem, error)) *MockFilesystemServiceGetByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceGetByNameCall) DoAndReturn(f func(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v10.Filesystem, error)) *MockFilesystemServiceGetByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOperation mocks base method.
func (m *MockFilesystemService) GetOperation(arg0 context.Context, arg1 *v1.GetOperationRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockFilesystemServiceMockRecorder) GetOperation(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceGetOperationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockFilesystemService)(nil).GetOperation), varargs...)
	return &MockFilesystemServiceGetOperationCall{Call: call}
}

// MockFilesystemServiceGetOperationCall wrap *gomock.Call
type MockFilesystemServiceGetOperationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceGetOperationCall) Return(arg0 operations.Operation, arg1 error) *MockFilesystemServiceGetOperationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceGetOperationCall) Do(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceGetOperationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceGetOperationCall) DoAndReturn(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceGetOperationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockFilesystemService) List(arg0 context.Context, arg1 *v10.ListFilesystemsRequest, arg2 ...grpc.CallOption) (*v10.ListFilesystemsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v10.ListFilesystemsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockFilesystemServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFilesystemService)(nil).List), varargs...)
	return &MockFilesystemServiceListCall{Call: call}
}

// MockFilesystemServiceListCall wrap *gomock.Call
type MockFilesystemServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceListCall) Return(arg0 *v10.ListFilesystemsResponse, arg1 error) *MockFilesystemServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceListCall) Do(f func(context.Context, *v10.ListFilesystemsRequest, ...grpc.CallOption) (*v10.ListFilesystemsResponse, error)) *MockFilesystemServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceListCall) DoAndReturn(f func(context.Context, *v10.ListFilesystemsRequest, ...grpc.CallOption) (*v10.ListFilesystemsResponse, error)) *MockFilesystemServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperations mocks base method.
func (m *MockFilesystemService) ListOperations(arg0 context.Context, arg1 *v1.ListOperationsRequest, arg2 ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
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
func (mr *MockFilesystemServiceMockRecorder) ListOperations(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceListOperationsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockFilesystemService)(nil).ListOperations), varargs...)
	return &MockFilesystemServiceListOperationsCall{Call: call}
}

// MockFilesystemServiceListOperationsCall wrap *gomock.Call
type MockFilesystemServiceListOperationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceListOperationsCall) Return(arg0 *v1.ListOperationsResponse, arg1 error) *MockFilesystemServiceListOperationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceListOperationsCall) Do(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockFilesystemServiceListOperationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceListOperationsCall) DoAndReturn(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockFilesystemServiceListOperationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperationsByParent mocks base method.
func (m *MockFilesystemService) ListOperationsByParent(arg0 context.Context, arg1 *v10.ListOperationsByParentRequest, arg2 ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperationsByParent", varargs...)
	ret0, _ := ret[0].(*v1.ListOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperationsByParent indicates an expected call of ListOperationsByParent.
func (mr *MockFilesystemServiceMockRecorder) ListOperationsByParent(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceListOperationsByParentCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperationsByParent", reflect.TypeOf((*MockFilesystemService)(nil).ListOperationsByParent), varargs...)
	return &MockFilesystemServiceListOperationsByParentCall{Call: call}
}

// MockFilesystemServiceListOperationsByParentCall wrap *gomock.Call
type MockFilesystemServiceListOperationsByParentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceListOperationsByParentCall) Return(arg0 *v1.ListOperationsResponse, arg1 error) *MockFilesystemServiceListOperationsByParentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceListOperationsByParentCall) Do(f func(context.Context, *v10.ListOperationsByParentRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockFilesystemServiceListOperationsByParentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceListOperationsByParentCall) DoAndReturn(f func(context.Context, *v10.ListOperationsByParentRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockFilesystemServiceListOperationsByParentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *MockFilesystemService) Update(arg0 context.Context, arg1 *v10.UpdateFilesystemRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockFilesystemServiceMockRecorder) Update(arg0, arg1 any, arg2 ...any) *MockFilesystemServiceUpdateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFilesystemService)(nil).Update), varargs...)
	return &MockFilesystemServiceUpdateCall{Call: call}
}

// MockFilesystemServiceUpdateCall wrap *gomock.Call
type MockFilesystemServiceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockFilesystemServiceUpdateCall) Return(arg0 operations.Operation, arg1 error) *MockFilesystemServiceUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockFilesystemServiceUpdateCall) Do(f func(context.Context, *v10.UpdateFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockFilesystemServiceUpdateCall) DoAndReturn(f func(context.Context, *v10.UpdateFilesystemRequest, ...grpc.CallOption) (operations.Operation, error)) *MockFilesystemServiceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
