// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/msp/v1alpha1/resource/template_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/msp/v1alpha1/resource/template_service.sdk.go -package resource -typed
//

// Package resource is a generated GoMock package.
package resource

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	resource "github.com/nebius/gosdk/proto/nebius/msp/v1alpha1/resource"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTemplateService is a mock of TemplateService interface.
type MockTemplateService struct {
	ctrl     *gomock.Controller
	recorder *MockTemplateServiceMockRecorder
}

// MockTemplateServiceMockRecorder is the mock recorder for MockTemplateService.
type MockTemplateServiceMockRecorder struct {
	mock *MockTemplateService
}

// NewMockTemplateService creates a new mock instance.
func NewMockTemplateService(ctrl *gomock.Controller) *MockTemplateService {
	mock := &MockTemplateService{ctrl: ctrl}
	mock.recorder = &MockTemplateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTemplateService) EXPECT() *MockTemplateServiceMockRecorder {
	return m.recorder
}

// Filter mocks base method.
func (m *MockTemplateService) Filter(arg0 context.Context, arg1 *resource.ListTemplatesRequest, arg2 ...grpc.CallOption) iter.Seq2[*resource.Template, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*resource.Template, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockTemplateServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockTemplateServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockTemplateService)(nil).Filter), varargs...)
	return &MockTemplateServiceFilterCall{Call: call}
}

// MockTemplateServiceFilterCall wrap *gomock.Call
type MockTemplateServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTemplateServiceFilterCall) Return(arg0 iter.Seq2[*resource.Template, error]) *MockTemplateServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTemplateServiceFilterCall) Do(f func(context.Context, *resource.ListTemplatesRequest, ...grpc.CallOption) iter.Seq2[*resource.Template, error]) *MockTemplateServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTemplateServiceFilterCall) DoAndReturn(f func(context.Context, *resource.ListTemplatesRequest, ...grpc.CallOption) iter.Seq2[*resource.Template, error]) *MockTemplateServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockTemplateService) List(arg0 context.Context, arg1 *resource.ListTemplatesRequest, arg2 ...grpc.CallOption) (*resource.ListTemplatesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*resource.ListTemplatesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTemplateServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockTemplateServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTemplateService)(nil).List), varargs...)
	return &MockTemplateServiceListCall{Call: call}
}

// MockTemplateServiceListCall wrap *gomock.Call
type MockTemplateServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTemplateServiceListCall) Return(arg0 *resource.ListTemplatesResponse, arg1 error) *MockTemplateServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTemplateServiceListCall) Do(f func(context.Context, *resource.ListTemplatesRequest, ...grpc.CallOption) (*resource.ListTemplatesResponse, error)) *MockTemplateServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTemplateServiceListCall) DoAndReturn(f func(context.Context, *resource.ListTemplatesRequest, ...grpc.CallOption) (*resource.ListTemplatesResponse, error)) *MockTemplateServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
