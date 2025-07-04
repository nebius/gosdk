// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/compute/v1/maintenance_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/compute/v1/maintenance_service.sdk.go -package v1 -typed
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	v1 "github.com/nebius/gosdk/proto/nebius/compute/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockMaintenanceService is a mock of MaintenanceService interface.
type MockMaintenanceService struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceServiceMockRecorder
}

// MockMaintenanceServiceMockRecorder is the mock recorder for MockMaintenanceService.
type MockMaintenanceServiceMockRecorder struct {
	mock *MockMaintenanceService
}

// NewMockMaintenanceService creates a new mock instance.
func NewMockMaintenanceService(ctrl *gomock.Controller) *MockMaintenanceService {
	mock := &MockMaintenanceService{ctrl: ctrl}
	mock.recorder = &MockMaintenanceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceService) EXPECT() *MockMaintenanceServiceMockRecorder {
	return m.recorder
}

// GetByInstance mocks base method.
func (m *MockMaintenanceService) GetByInstance(arg0 context.Context, arg1 *v1.GetMaintenanceEventByInstanceRequest, arg2 ...grpc.CallOption) (*v1.MaintenanceEvent, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByInstance", varargs...)
	ret0, _ := ret[0].(*v1.MaintenanceEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByInstance indicates an expected call of GetByInstance.
func (mr *MockMaintenanceServiceMockRecorder) GetByInstance(arg0, arg1 any, arg2 ...any) *MockMaintenanceServiceGetByInstanceCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByInstance", reflect.TypeOf((*MockMaintenanceService)(nil).GetByInstance), varargs...)
	return &MockMaintenanceServiceGetByInstanceCall{Call: call}
}

// MockMaintenanceServiceGetByInstanceCall wrap *gomock.Call
type MockMaintenanceServiceGetByInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMaintenanceServiceGetByInstanceCall) Return(arg0 *v1.MaintenanceEvent, arg1 error) *MockMaintenanceServiceGetByInstanceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMaintenanceServiceGetByInstanceCall) Do(f func(context.Context, *v1.GetMaintenanceEventByInstanceRequest, ...grpc.CallOption) (*v1.MaintenanceEvent, error)) *MockMaintenanceServiceGetByInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMaintenanceServiceGetByInstanceCall) DoAndReturn(f func(context.Context, *v1.GetMaintenanceEventByInstanceRequest, ...grpc.CallOption) (*v1.MaintenanceEvent, error)) *MockMaintenanceServiceGetByInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListActive mocks base method.
func (m *MockMaintenanceService) ListActive(arg0 context.Context, arg1 *v1.ListMaintenanceEventsRequest, arg2 ...grpc.CallOption) (*v1.ListMaintenanceEventsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListActive", varargs...)
	ret0, _ := ret[0].(*v1.ListMaintenanceEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActive indicates an expected call of ListActive.
func (mr *MockMaintenanceServiceMockRecorder) ListActive(arg0, arg1 any, arg2 ...any) *MockMaintenanceServiceListActiveCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActive", reflect.TypeOf((*MockMaintenanceService)(nil).ListActive), varargs...)
	return &MockMaintenanceServiceListActiveCall{Call: call}
}

// MockMaintenanceServiceListActiveCall wrap *gomock.Call
type MockMaintenanceServiceListActiveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMaintenanceServiceListActiveCall) Return(arg0 *v1.ListMaintenanceEventsResponse, arg1 error) *MockMaintenanceServiceListActiveCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMaintenanceServiceListActiveCall) Do(f func(context.Context, *v1.ListMaintenanceEventsRequest, ...grpc.CallOption) (*v1.ListMaintenanceEventsResponse, error)) *MockMaintenanceServiceListActiveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMaintenanceServiceListActiveCall) DoAndReturn(f func(context.Context, *v1.ListMaintenanceEventsRequest, ...grpc.CallOption) (*v1.ListMaintenanceEventsResponse, error)) *MockMaintenanceServiceListActiveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
