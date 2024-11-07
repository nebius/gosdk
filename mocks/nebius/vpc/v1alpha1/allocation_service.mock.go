// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/vpc/v1alpha1/allocation_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/vpc/v1alpha1/allocation_service.sdk.go -package v1alpha1 -typed
//

// Package v1alpha1 is a generated GoMock package.
package v1alpha1

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	alphaops "github.com/nebius/gosdk/operations/alphaops"
	v1alpha1 "github.com/nebius/gosdk/proto/nebius/common/v1alpha1"
	v1alpha10 "github.com/nebius/gosdk/proto/nebius/vpc/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAllocationService is a mock of AllocationService interface.
type MockAllocationService struct {
	ctrl     *gomock.Controller
	recorder *MockAllocationServiceMockRecorder
}

// MockAllocationServiceMockRecorder is the mock recorder for MockAllocationService.
type MockAllocationServiceMockRecorder struct {
	mock *MockAllocationService
}

// NewMockAllocationService creates a new mock instance.
func NewMockAllocationService(ctrl *gomock.Controller) *MockAllocationService {
	mock := &MockAllocationService{ctrl: ctrl}
	mock.recorder = &MockAllocationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllocationService) EXPECT() *MockAllocationServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAllocationService) Create(arg0 context.Context, arg1 *v1alpha10.CreateAllocationRequest, arg2 ...grpc.CallOption) (*alphaops.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*alphaops.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAllocationServiceMockRecorder) Create(arg0, arg1 any, arg2 ...any) *MockAllocationServiceCreateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAllocationService)(nil).Create), varargs...)
	return &MockAllocationServiceCreateCall{Call: call}
}

// MockAllocationServiceCreateCall wrap *gomock.Call
type MockAllocationServiceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceCreateCall) Return(arg0 *alphaops.Operation, arg1 error) *MockAllocationServiceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceCreateCall) Do(f func(context.Context, *v1alpha10.CreateAllocationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceCreateCall) DoAndReturn(f func(context.Context, *v1alpha10.CreateAllocationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockAllocationService) Delete(arg0 context.Context, arg1 *v1alpha10.DeleteAllocationRequest, arg2 ...grpc.CallOption) (*alphaops.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*alphaops.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAllocationServiceMockRecorder) Delete(arg0, arg1 any, arg2 ...any) *MockAllocationServiceDeleteCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAllocationService)(nil).Delete), varargs...)
	return &MockAllocationServiceDeleteCall{Call: call}
}

// MockAllocationServiceDeleteCall wrap *gomock.Call
type MockAllocationServiceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceDeleteCall) Return(arg0 *alphaops.Operation, arg1 error) *MockAllocationServiceDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceDeleteCall) Do(f func(context.Context, *v1alpha10.DeleteAllocationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceDeleteCall) DoAndReturn(f func(context.Context, *v1alpha10.DeleteAllocationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockAllocationService) Filter(arg0 context.Context, arg1 *v1alpha10.ListAllocationsRequest, arg2 ...grpc.CallOption) iter.Seq2[*v1alpha10.Allocation, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v1alpha10.Allocation, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockAllocationServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockAllocationServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockAllocationService)(nil).Filter), varargs...)
	return &MockAllocationServiceFilterCall{Call: call}
}

// MockAllocationServiceFilterCall wrap *gomock.Call
type MockAllocationServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceFilterCall) Return(arg0 iter.Seq2[*v1alpha10.Allocation, error]) *MockAllocationServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceFilterCall) Do(f func(context.Context, *v1alpha10.ListAllocationsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha10.Allocation, error]) *MockAllocationServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceFilterCall) DoAndReturn(f func(context.Context, *v1alpha10.ListAllocationsRequest, ...grpc.CallOption) iter.Seq2[*v1alpha10.Allocation, error]) *MockAllocationServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockAllocationService) Get(arg0 context.Context, arg1 *v1alpha10.GetAllocationRequest, arg2 ...grpc.CallOption) (*v1alpha10.Allocation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v1alpha10.Allocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAllocationServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockAllocationServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAllocationService)(nil).Get), varargs...)
	return &MockAllocationServiceGetCall{Call: call}
}

// MockAllocationServiceGetCall wrap *gomock.Call
type MockAllocationServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceGetCall) Return(arg0 *v1alpha10.Allocation, arg1 error) *MockAllocationServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceGetCall) Do(f func(context.Context, *v1alpha10.GetAllocationRequest, ...grpc.CallOption) (*v1alpha10.Allocation, error)) *MockAllocationServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceGetCall) DoAndReturn(f func(context.Context, *v1alpha10.GetAllocationRequest, ...grpc.CallOption) (*v1alpha10.Allocation, error)) *MockAllocationServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByName mocks base method.
func (m *MockAllocationService) GetByName(arg0 context.Context, arg1 *v1alpha10.GetAllocationByNameRequest, arg2 ...grpc.CallOption) (*v1alpha10.Allocation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByName", varargs...)
	ret0, _ := ret[0].(*v1alpha10.Allocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockAllocationServiceMockRecorder) GetByName(arg0, arg1 any, arg2 ...any) *MockAllocationServiceGetByNameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockAllocationService)(nil).GetByName), varargs...)
	return &MockAllocationServiceGetByNameCall{Call: call}
}

// MockAllocationServiceGetByNameCall wrap *gomock.Call
type MockAllocationServiceGetByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceGetByNameCall) Return(arg0 *v1alpha10.Allocation, arg1 error) *MockAllocationServiceGetByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceGetByNameCall) Do(f func(context.Context, *v1alpha10.GetAllocationByNameRequest, ...grpc.CallOption) (*v1alpha10.Allocation, error)) *MockAllocationServiceGetByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceGetByNameCall) DoAndReturn(f func(context.Context, *v1alpha10.GetAllocationByNameRequest, ...grpc.CallOption) (*v1alpha10.Allocation, error)) *MockAllocationServiceGetByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOperation mocks base method.
func (m *MockAllocationService) GetOperation(arg0 context.Context, arg1 *v1alpha1.GetOperationRequest, arg2 ...grpc.CallOption) (*alphaops.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperation", varargs...)
	ret0, _ := ret[0].(*alphaops.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockAllocationServiceMockRecorder) GetOperation(arg0, arg1 any, arg2 ...any) *MockAllocationServiceGetOperationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockAllocationService)(nil).GetOperation), varargs...)
	return &MockAllocationServiceGetOperationCall{Call: call}
}

// MockAllocationServiceGetOperationCall wrap *gomock.Call
type MockAllocationServiceGetOperationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceGetOperationCall) Return(arg0 *alphaops.Operation, arg1 error) *MockAllocationServiceGetOperationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceGetOperationCall) Do(f func(context.Context, *v1alpha1.GetOperationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceGetOperationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceGetOperationCall) DoAndReturn(f func(context.Context, *v1alpha1.GetOperationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceGetOperationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockAllocationService) List(arg0 context.Context, arg1 *v1alpha10.ListAllocationsRequest, arg2 ...grpc.CallOption) (*v1alpha10.ListAllocationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1alpha10.ListAllocationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAllocationServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockAllocationServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAllocationService)(nil).List), varargs...)
	return &MockAllocationServiceListCall{Call: call}
}

// MockAllocationServiceListCall wrap *gomock.Call
type MockAllocationServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceListCall) Return(arg0 *v1alpha10.ListAllocationsResponse, arg1 error) *MockAllocationServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceListCall) Do(f func(context.Context, *v1alpha10.ListAllocationsRequest, ...grpc.CallOption) (*v1alpha10.ListAllocationsResponse, error)) *MockAllocationServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceListCall) DoAndReturn(f func(context.Context, *v1alpha10.ListAllocationsRequest, ...grpc.CallOption) (*v1alpha10.ListAllocationsResponse, error)) *MockAllocationServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperations mocks base method.
func (m *MockAllocationService) ListOperations(arg0 context.Context, arg1 *v1alpha1.ListOperationsRequest, arg2 ...grpc.CallOption) (*v1alpha1.ListOperationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ListOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockAllocationServiceMockRecorder) ListOperations(arg0, arg1 any, arg2 ...any) *MockAllocationServiceListOperationsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockAllocationService)(nil).ListOperations), varargs...)
	return &MockAllocationServiceListOperationsCall{Call: call}
}

// MockAllocationServiceListOperationsCall wrap *gomock.Call
type MockAllocationServiceListOperationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceListOperationsCall) Return(arg0 *v1alpha1.ListOperationsResponse, arg1 error) *MockAllocationServiceListOperationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceListOperationsCall) Do(f func(context.Context, *v1alpha1.ListOperationsRequest, ...grpc.CallOption) (*v1alpha1.ListOperationsResponse, error)) *MockAllocationServiceListOperationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceListOperationsCall) DoAndReturn(f func(context.Context, *v1alpha1.ListOperationsRequest, ...grpc.CallOption) (*v1alpha1.ListOperationsResponse, error)) *MockAllocationServiceListOperationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *MockAllocationService) Update(arg0 context.Context, arg1 *v1alpha10.UpdateAllocationRequest, arg2 ...grpc.CallOption) (*alphaops.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*alphaops.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAllocationServiceMockRecorder) Update(arg0, arg1 any, arg2 ...any) *MockAllocationServiceUpdateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAllocationService)(nil).Update), varargs...)
	return &MockAllocationServiceUpdateCall{Call: call}
}

// MockAllocationServiceUpdateCall wrap *gomock.Call
type MockAllocationServiceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAllocationServiceUpdateCall) Return(arg0 *alphaops.Operation, arg1 error) *MockAllocationServiceUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAllocationServiceUpdateCall) Do(f func(context.Context, *v1alpha10.UpdateAllocationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAllocationServiceUpdateCall) DoAndReturn(f func(context.Context, *v1alpha10.UpdateAllocationRequest, ...grpc.CallOption) (*alphaops.Operation, error)) *MockAllocationServiceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
