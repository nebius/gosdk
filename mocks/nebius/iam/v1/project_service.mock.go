// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/iam/v1/project_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/iam/v1/project_service.sdk.go -package v1 -typed
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/iam/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockProjectService is a mock of ProjectService interface.
type MockProjectService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceMockRecorder
}

// MockProjectServiceMockRecorder is the mock recorder for MockProjectService.
type MockProjectServiceMockRecorder struct {
	mock *MockProjectService
}

// NewMockProjectService creates a new mock instance.
func NewMockProjectService(ctrl *gomock.Controller) *MockProjectService {
	mock := &MockProjectService{ctrl: ctrl}
	mock.recorder = &MockProjectServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectService) EXPECT() *MockProjectServiceMockRecorder {
	return m.recorder
}

// Filter mocks base method.
func (m *MockProjectService) Filter(arg0 context.Context, arg1 *v1.ListProjectsRequest, arg2 ...grpc.CallOption) iter.Seq2[*v1.Container, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v1.Container, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockProjectServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockProjectServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockProjectService)(nil).Filter), varargs...)
	return &MockProjectServiceFilterCall{Call: call}
}

// MockProjectServiceFilterCall wrap *gomock.Call
type MockProjectServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectServiceFilterCall) Return(arg0 iter.Seq2[*v1.Container, error]) *MockProjectServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectServiceFilterCall) Do(f func(context.Context, *v1.ListProjectsRequest, ...grpc.CallOption) iter.Seq2[*v1.Container, error]) *MockProjectServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectServiceFilterCall) DoAndReturn(f func(context.Context, *v1.ListProjectsRequest, ...grpc.CallOption) iter.Seq2[*v1.Container, error]) *MockProjectServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockProjectService) Get(arg0 context.Context, arg1 *v1.GetProjectRequest, arg2 ...grpc.CallOption) (*v1.Container, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v1.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProjectServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockProjectServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProjectService)(nil).Get), varargs...)
	return &MockProjectServiceGetCall{Call: call}
}

// MockProjectServiceGetCall wrap *gomock.Call
type MockProjectServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectServiceGetCall) Return(arg0 *v1.Container, arg1 error) *MockProjectServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectServiceGetCall) Do(f func(context.Context, *v1.GetProjectRequest, ...grpc.CallOption) (*v1.Container, error)) *MockProjectServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectServiceGetCall) DoAndReturn(f func(context.Context, *v1.GetProjectRequest, ...grpc.CallOption) (*v1.Container, error)) *MockProjectServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByName mocks base method.
func (m *MockProjectService) GetByName(arg0 context.Context, arg1 *v1.GetProjectByNameRequest, arg2 ...grpc.CallOption) (*v1.Container, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByName", varargs...)
	ret0, _ := ret[0].(*v1.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockProjectServiceMockRecorder) GetByName(arg0, arg1 any, arg2 ...any) *MockProjectServiceGetByNameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockProjectService)(nil).GetByName), varargs...)
	return &MockProjectServiceGetByNameCall{Call: call}
}

// MockProjectServiceGetByNameCall wrap *gomock.Call
type MockProjectServiceGetByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectServiceGetByNameCall) Return(arg0 *v1.Container, arg1 error) *MockProjectServiceGetByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectServiceGetByNameCall) Do(f func(context.Context, *v1.GetProjectByNameRequest, ...grpc.CallOption) (*v1.Container, error)) *MockProjectServiceGetByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectServiceGetByNameCall) DoAndReturn(f func(context.Context, *v1.GetProjectByNameRequest, ...grpc.CallOption) (*v1.Container, error)) *MockProjectServiceGetByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockProjectService) List(arg0 context.Context, arg1 *v1.ListProjectsRequest, arg2 ...grpc.CallOption) (*v1.ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v1.ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProjectServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockProjectServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProjectService)(nil).List), varargs...)
	return &MockProjectServiceListCall{Call: call}
}

// MockProjectServiceListCall wrap *gomock.Call
type MockProjectServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockProjectServiceListCall) Return(arg0 *v1.ListProjectsResponse, arg1 error) *MockProjectServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockProjectServiceListCall) Do(f func(context.Context, *v1.ListProjectsRequest, ...grpc.CallOption) (*v1.ListProjectsResponse, error)) *MockProjectServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockProjectServiceListCall) DoAndReturn(f func(context.Context, *v1.ListProjectsRequest, ...grpc.CallOption) (*v1.ListProjectsResponse, error)) *MockProjectServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
