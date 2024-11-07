// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/compute/v1/image_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/compute/v1/image_service.sdk.go -package v1 -typed
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v10 "github.com/nebius/gosdk/proto/nebius/compute/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockImageService is a mock of ImageService interface.
type MockImageService struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceMockRecorder
}

// MockImageServiceMockRecorder is the mock recorder for MockImageService.
type MockImageServiceMockRecorder struct {
	mock *MockImageService
}

// NewMockImageService creates a new mock instance.
func NewMockImageService(ctrl *gomock.Controller) *MockImageService {
	mock := &MockImageService{ctrl: ctrl}
	mock.recorder = &MockImageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageService) EXPECT() *MockImageServiceMockRecorder {
	return m.recorder
}

// Filter mocks base method.
func (m *MockImageService) Filter(arg0 context.Context, arg1 *v10.ListImagesRequest, arg2 ...grpc.CallOption) iter.Seq2[*v10.Image, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v10.Image, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockImageServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockImageServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockImageService)(nil).Filter), varargs...)
	return &MockImageServiceFilterCall{Call: call}
}

// MockImageServiceFilterCall wrap *gomock.Call
type MockImageServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImageServiceFilterCall) Return(arg0 iter.Seq2[*v10.Image, error]) *MockImageServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImageServiceFilterCall) Do(f func(context.Context, *v10.ListImagesRequest, ...grpc.CallOption) iter.Seq2[*v10.Image, error]) *MockImageServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImageServiceFilterCall) DoAndReturn(f func(context.Context, *v10.ListImagesRequest, ...grpc.CallOption) iter.Seq2[*v10.Image, error]) *MockImageServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockImageService) Get(arg0 context.Context, arg1 *v10.GetImageRequest, arg2 ...grpc.CallOption) (*v10.Image, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v10.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockImageServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockImageServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockImageService)(nil).Get), varargs...)
	return &MockImageServiceGetCall{Call: call}
}

// MockImageServiceGetCall wrap *gomock.Call
type MockImageServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImageServiceGetCall) Return(arg0 *v10.Image, arg1 error) *MockImageServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImageServiceGetCall) Do(f func(context.Context, *v10.GetImageRequest, ...grpc.CallOption) (*v10.Image, error)) *MockImageServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImageServiceGetCall) DoAndReturn(f func(context.Context, *v10.GetImageRequest, ...grpc.CallOption) (*v10.Image, error)) *MockImageServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByName mocks base method.
func (m *MockImageService) GetByName(arg0 context.Context, arg1 *v1.GetByNameRequest, arg2 ...grpc.CallOption) (*v10.Image, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByName", varargs...)
	ret0, _ := ret[0].(*v10.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockImageServiceMockRecorder) GetByName(arg0, arg1 any, arg2 ...any) *MockImageServiceGetByNameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockImageService)(nil).GetByName), varargs...)
	return &MockImageServiceGetByNameCall{Call: call}
}

// MockImageServiceGetByNameCall wrap *gomock.Call
type MockImageServiceGetByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImageServiceGetByNameCall) Return(arg0 *v10.Image, arg1 error) *MockImageServiceGetByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImageServiceGetByNameCall) Do(f func(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v10.Image, error)) *MockImageServiceGetByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImageServiceGetByNameCall) DoAndReturn(f func(context.Context, *v1.GetByNameRequest, ...grpc.CallOption) (*v10.Image, error)) *MockImageServiceGetByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLatestByFamily mocks base method.
func (m *MockImageService) GetLatestByFamily(arg0 context.Context, arg1 *v10.GetImageLatestByFamilyRequest, arg2 ...grpc.CallOption) (*v10.Image, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatestByFamily", varargs...)
	ret0, _ := ret[0].(*v10.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestByFamily indicates an expected call of GetLatestByFamily.
func (mr *MockImageServiceMockRecorder) GetLatestByFamily(arg0, arg1 any, arg2 ...any) *MockImageServiceGetLatestByFamilyCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestByFamily", reflect.TypeOf((*MockImageService)(nil).GetLatestByFamily), varargs...)
	return &MockImageServiceGetLatestByFamilyCall{Call: call}
}

// MockImageServiceGetLatestByFamilyCall wrap *gomock.Call
type MockImageServiceGetLatestByFamilyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImageServiceGetLatestByFamilyCall) Return(arg0 *v10.Image, arg1 error) *MockImageServiceGetLatestByFamilyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImageServiceGetLatestByFamilyCall) Do(f func(context.Context, *v10.GetImageLatestByFamilyRequest, ...grpc.CallOption) (*v10.Image, error)) *MockImageServiceGetLatestByFamilyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImageServiceGetLatestByFamilyCall) DoAndReturn(f func(context.Context, *v10.GetImageLatestByFamilyRequest, ...grpc.CallOption) (*v10.Image, error)) *MockImageServiceGetLatestByFamilyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockImageService) List(arg0 context.Context, arg1 *v10.ListImagesRequest, arg2 ...grpc.CallOption) (*v10.ListImagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v10.ListImagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockImageServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockImageServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockImageService)(nil).List), varargs...)
	return &MockImageServiceListCall{Call: call}
}

// MockImageServiceListCall wrap *gomock.Call
type MockImageServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImageServiceListCall) Return(arg0 *v10.ListImagesResponse, arg1 error) *MockImageServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImageServiceListCall) Do(f func(context.Context, *v10.ListImagesRequest, ...grpc.CallOption) (*v10.ListImagesResponse, error)) *MockImageServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImageServiceListCall) DoAndReturn(f func(context.Context, *v10.ListImagesRequest, ...grpc.CallOption) (*v10.ListImagesResponse, error)) *MockImageServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperationsByParent mocks base method.
func (m *MockImageService) ListOperationsByParent(arg0 context.Context, arg1 *v10.ListOperationsByParentRequest, arg2 ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
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
func (mr *MockImageServiceMockRecorder) ListOperationsByParent(arg0, arg1 any, arg2 ...any) *MockImageServiceListOperationsByParentCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperationsByParent", reflect.TypeOf((*MockImageService)(nil).ListOperationsByParent), varargs...)
	return &MockImageServiceListOperationsByParentCall{Call: call}
}

// MockImageServiceListOperationsByParentCall wrap *gomock.Call
type MockImageServiceListOperationsByParentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImageServiceListOperationsByParentCall) Return(arg0 *v1.ListOperationsResponse, arg1 error) *MockImageServiceListOperationsByParentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImageServiceListOperationsByParentCall) Do(f func(context.Context, *v10.ListOperationsByParentRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockImageServiceListOperationsByParentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImageServiceListOperationsByParentCall) DoAndReturn(f func(context.Context, *v10.ListOperationsByParentRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockImageServiceListOperationsByParentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
