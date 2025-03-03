// Code generated by MockGen. DO NOT EDIT.
// Source: services/nebius/storage/v1/bucket_service.sdk.go
//
// Generated by this command:
//
//	mockgen -source services/nebius/storage/v1/bucket_service.sdk.go -package v1 -typed
//

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	iter "github.com/nebius/gosdk/iter"
	operations "github.com/nebius/gosdk/operations"
	v1 "github.com/nebius/gosdk/proto/nebius/common/v1"
	v10 "github.com/nebius/gosdk/proto/nebius/storage/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockBucketService is a mock of BucketService interface.
type MockBucketService struct {
	ctrl     *gomock.Controller
	recorder *MockBucketServiceMockRecorder
}

// MockBucketServiceMockRecorder is the mock recorder for MockBucketService.
type MockBucketServiceMockRecorder struct {
	mock *MockBucketService
}

// NewMockBucketService creates a new mock instance.
func NewMockBucketService(ctrl *gomock.Controller) *MockBucketService {
	mock := &MockBucketService{ctrl: ctrl}
	mock.recorder = &MockBucketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucketService) EXPECT() *MockBucketServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBucketService) Create(arg0 context.Context, arg1 *v10.CreateBucketRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockBucketServiceMockRecorder) Create(arg0, arg1 any, arg2 ...any) *MockBucketServiceCreateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBucketService)(nil).Create), varargs...)
	return &MockBucketServiceCreateCall{Call: call}
}

// MockBucketServiceCreateCall wrap *gomock.Call
type MockBucketServiceCreateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceCreateCall) Return(arg0 operations.Operation, arg1 error) *MockBucketServiceCreateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceCreateCall) Do(f func(context.Context, *v10.CreateBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceCreateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceCreateCall) DoAndReturn(f func(context.Context, *v10.CreateBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceCreateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockBucketService) Delete(arg0 context.Context, arg1 *v10.DeleteBucketRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockBucketServiceMockRecorder) Delete(arg0, arg1 any, arg2 ...any) *MockBucketServiceDeleteCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBucketService)(nil).Delete), varargs...)
	return &MockBucketServiceDeleteCall{Call: call}
}

// MockBucketServiceDeleteCall wrap *gomock.Call
type MockBucketServiceDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceDeleteCall) Return(arg0 operations.Operation, arg1 error) *MockBucketServiceDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceDeleteCall) Do(f func(context.Context, *v10.DeleteBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceDeleteCall) DoAndReturn(f func(context.Context, *v10.DeleteBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockBucketService) Filter(arg0 context.Context, arg1 *v10.ListBucketsRequest, arg2 ...grpc.CallOption) iter.Seq2[*v10.Bucket, error] {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].(iter.Seq2[*v10.Bucket, error])
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockBucketServiceMockRecorder) Filter(arg0, arg1 any, arg2 ...any) *MockBucketServiceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockBucketService)(nil).Filter), varargs...)
	return &MockBucketServiceFilterCall{Call: call}
}

// MockBucketServiceFilterCall wrap *gomock.Call
type MockBucketServiceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceFilterCall) Return(arg0 iter.Seq2[*v10.Bucket, error]) *MockBucketServiceFilterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceFilterCall) Do(f func(context.Context, *v10.ListBucketsRequest, ...grpc.CallOption) iter.Seq2[*v10.Bucket, error]) *MockBucketServiceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceFilterCall) DoAndReturn(f func(context.Context, *v10.ListBucketsRequest, ...grpc.CallOption) iter.Seq2[*v10.Bucket, error]) *MockBucketServiceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockBucketService) Get(arg0 context.Context, arg1 *v10.GetBucketRequest, arg2 ...grpc.CallOption) (*v10.Bucket, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*v10.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBucketServiceMockRecorder) Get(arg0, arg1 any, arg2 ...any) *MockBucketServiceGetCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBucketService)(nil).Get), varargs...)
	return &MockBucketServiceGetCall{Call: call}
}

// MockBucketServiceGetCall wrap *gomock.Call
type MockBucketServiceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceGetCall) Return(arg0 *v10.Bucket, arg1 error) *MockBucketServiceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceGetCall) Do(f func(context.Context, *v10.GetBucketRequest, ...grpc.CallOption) (*v10.Bucket, error)) *MockBucketServiceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceGetCall) DoAndReturn(f func(context.Context, *v10.GetBucketRequest, ...grpc.CallOption) (*v10.Bucket, error)) *MockBucketServiceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByName mocks base method.
func (m *MockBucketService) GetByName(arg0 context.Context, arg1 *v10.GetBucketByNameRequest, arg2 ...grpc.CallOption) (*v10.Bucket, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByName", varargs...)
	ret0, _ := ret[0].(*v10.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockBucketServiceMockRecorder) GetByName(arg0, arg1 any, arg2 ...any) *MockBucketServiceGetByNameCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockBucketService)(nil).GetByName), varargs...)
	return &MockBucketServiceGetByNameCall{Call: call}
}

// MockBucketServiceGetByNameCall wrap *gomock.Call
type MockBucketServiceGetByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceGetByNameCall) Return(arg0 *v10.Bucket, arg1 error) *MockBucketServiceGetByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceGetByNameCall) Do(f func(context.Context, *v10.GetBucketByNameRequest, ...grpc.CallOption) (*v10.Bucket, error)) *MockBucketServiceGetByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceGetByNameCall) DoAndReturn(f func(context.Context, *v10.GetBucketByNameRequest, ...grpc.CallOption) (*v10.Bucket, error)) *MockBucketServiceGetByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetOperation mocks base method.
func (m *MockBucketService) GetOperation(arg0 context.Context, arg1 *v1.GetOperationRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockBucketServiceMockRecorder) GetOperation(arg0, arg1 any, arg2 ...any) *MockBucketServiceGetOperationCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockBucketService)(nil).GetOperation), varargs...)
	return &MockBucketServiceGetOperationCall{Call: call}
}

// MockBucketServiceGetOperationCall wrap *gomock.Call
type MockBucketServiceGetOperationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceGetOperationCall) Return(arg0 operations.Operation, arg1 error) *MockBucketServiceGetOperationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceGetOperationCall) Do(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceGetOperationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceGetOperationCall) DoAndReturn(f func(context.Context, *v1.GetOperationRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceGetOperationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockBucketService) List(arg0 context.Context, arg1 *v10.ListBucketsRequest, arg2 ...grpc.CallOption) (*v10.ListBucketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*v10.ListBucketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockBucketServiceMockRecorder) List(arg0, arg1 any, arg2 ...any) *MockBucketServiceListCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBucketService)(nil).List), varargs...)
	return &MockBucketServiceListCall{Call: call}
}

// MockBucketServiceListCall wrap *gomock.Call
type MockBucketServiceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceListCall) Return(arg0 *v10.ListBucketsResponse, arg1 error) *MockBucketServiceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceListCall) Do(f func(context.Context, *v10.ListBucketsRequest, ...grpc.CallOption) (*v10.ListBucketsResponse, error)) *MockBucketServiceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceListCall) DoAndReturn(f func(context.Context, *v10.ListBucketsRequest, ...grpc.CallOption) (*v10.ListBucketsResponse, error)) *MockBucketServiceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListOperations mocks base method.
func (m *MockBucketService) ListOperations(arg0 context.Context, arg1 *v1.ListOperationsRequest, arg2 ...grpc.CallOption) (*v1.ListOperationsResponse, error) {
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
func (mr *MockBucketServiceMockRecorder) ListOperations(arg0, arg1 any, arg2 ...any) *MockBucketServiceListOperationsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockBucketService)(nil).ListOperations), varargs...)
	return &MockBucketServiceListOperationsCall{Call: call}
}

// MockBucketServiceListOperationsCall wrap *gomock.Call
type MockBucketServiceListOperationsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceListOperationsCall) Return(arg0 *v1.ListOperationsResponse, arg1 error) *MockBucketServiceListOperationsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceListOperationsCall) Do(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockBucketServiceListOperationsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceListOperationsCall) DoAndReturn(f func(context.Context, *v1.ListOperationsRequest, ...grpc.CallOption) (*v1.ListOperationsResponse, error)) *MockBucketServiceListOperationsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Purge mocks base method.
func (m *MockBucketService) Purge(arg0 context.Context, arg1 *v10.PurgeBucketRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Purge", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Purge indicates an expected call of Purge.
func (mr *MockBucketServiceMockRecorder) Purge(arg0, arg1 any, arg2 ...any) *MockBucketServicePurgeCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Purge", reflect.TypeOf((*MockBucketService)(nil).Purge), varargs...)
	return &MockBucketServicePurgeCall{Call: call}
}

// MockBucketServicePurgeCall wrap *gomock.Call
type MockBucketServicePurgeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServicePurgeCall) Return(arg0 operations.Operation, arg1 error) *MockBucketServicePurgeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServicePurgeCall) Do(f func(context.Context, *v10.PurgeBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServicePurgeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServicePurgeCall) DoAndReturn(f func(context.Context, *v10.PurgeBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServicePurgeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Undelete mocks base method.
func (m *MockBucketService) Undelete(arg0 context.Context, arg1 *v10.UndeleteBucketRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Undelete", varargs...)
	ret0, _ := ret[0].(operations.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Undelete indicates an expected call of Undelete.
func (mr *MockBucketServiceMockRecorder) Undelete(arg0, arg1 any, arg2 ...any) *MockBucketServiceUndeleteCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Undelete", reflect.TypeOf((*MockBucketService)(nil).Undelete), varargs...)
	return &MockBucketServiceUndeleteCall{Call: call}
}

// MockBucketServiceUndeleteCall wrap *gomock.Call
type MockBucketServiceUndeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceUndeleteCall) Return(arg0 operations.Operation, arg1 error) *MockBucketServiceUndeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceUndeleteCall) Do(f func(context.Context, *v10.UndeleteBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceUndeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceUndeleteCall) DoAndReturn(f func(context.Context, *v10.UndeleteBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceUndeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m *MockBucketService) Update(arg0 context.Context, arg1 *v10.UpdateBucketRequest, arg2 ...grpc.CallOption) (operations.Operation, error) {
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
func (mr *MockBucketServiceMockRecorder) Update(arg0, arg1 any, arg2 ...any) *MockBucketServiceUpdateCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBucketService)(nil).Update), varargs...)
	return &MockBucketServiceUpdateCall{Call: call}
}

// MockBucketServiceUpdateCall wrap *gomock.Call
type MockBucketServiceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBucketServiceUpdateCall) Return(arg0 operations.Operation, arg1 error) *MockBucketServiceUpdateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBucketServiceUpdateCall) Do(f func(context.Context, *v10.UpdateBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBucketServiceUpdateCall) DoAndReturn(f func(context.Context, *v10.UpdateBucketRequest, ...grpc.CallOption) (operations.Operation, error)) *MockBucketServiceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
