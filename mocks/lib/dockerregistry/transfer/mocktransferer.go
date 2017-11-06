// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/lib/dockerregistry/transfer (interfaces: ImageTransferer)

// Package mocktransferer is a generated GoMock package.
package mocktransferer

import (
	transfer "code.uber.internal/infra/kraken/lib/dockerregistry/transfer"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockImageTransferer is a mock of ImageTransferer interface
type MockImageTransferer struct {
	ctrl     *gomock.Controller
	recorder *MockImageTransfererMockRecorder
}

// MockImageTransfererMockRecorder is the mock recorder for MockImageTransferer
type MockImageTransfererMockRecorder struct {
	mock *MockImageTransferer
}

// NewMockImageTransferer creates a new mock instance
func NewMockImageTransferer(ctrl *gomock.Controller) *MockImageTransferer {
	mock := &MockImageTransferer{ctrl: ctrl}
	mock.recorder = &MockImageTransfererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageTransferer) EXPECT() *MockImageTransfererMockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockImageTransferer) Download(arg0 string) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockImageTransfererMockRecorder) Download(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockImageTransferer)(nil).Download), arg0)
}

// GetManifest mocks base method
func (m *MockImageTransferer) GetManifest(arg0, arg1 string) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "GetManifest", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManifest indicates an expected call of GetManifest
func (mr *MockImageTransfererMockRecorder) GetManifest(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManifest", reflect.TypeOf((*MockImageTransferer)(nil).GetManifest), arg0, arg1)
}

// PostManifest mocks base method
func (m *MockImageTransferer) PostManifest(arg0, arg1, arg2 string, arg3 io.Reader) error {
	ret := m.ctrl.Call(m, "PostManifest", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostManifest indicates an expected call of PostManifest
func (mr *MockImageTransfererMockRecorder) PostManifest(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostManifest", reflect.TypeOf((*MockImageTransferer)(nil).PostManifest), arg0, arg1, arg2, arg3)
}

// Upload mocks base method
func (m *MockImageTransferer) Upload(arg0 string, arg1 transfer.IOCloner, arg2 int64) error {
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockImageTransfererMockRecorder) Upload(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockImageTransferer)(nil).Upload), arg0, arg1, arg2)
}
