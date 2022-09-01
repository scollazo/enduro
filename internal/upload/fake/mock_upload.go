// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/upload (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	io "io"
	reflect "reflect"

	upload "github.com/artefactual-sdps/enduro/internal/api/gen/upload"
	gomock "github.com/golang/mock/gomock"
	blob "gocloud.dev/blob"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Bucket mocks base method.
func (m *MockService) Bucket() *blob.Bucket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bucket")
	ret0, _ := ret[0].(*blob.Bucket)
	return ret0
}

// Bucket indicates an expected call of Bucket.
func (mr *MockServiceMockRecorder) Bucket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bucket", reflect.TypeOf((*MockService)(nil).Bucket))
}

// Close mocks base method.
func (m *MockService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockService)(nil).Close))
}

// Upload mocks base method.
func (m *MockService) Upload(arg0 context.Context, arg1 *upload.UploadPayload, arg2 io.ReadCloser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockServiceMockRecorder) Upload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockService)(nil).Upload), arg0, arg1, arg2)
}
