// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/lib/persistedretry (interfaces: Store,Task,Executor,Manager)

// Package mockpersistedretry is a generated GoMock package.
package mockpersistedretry

import (
	persistedretry "code.uber.internal/infra/kraken/lib/persistedretry"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStore) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// GetFailed mocks base method
func (m *MockStore) GetFailed() ([]persistedretry.Task, error) {
	ret := m.ctrl.Call(m, "GetFailed")
	ret0, _ := ret[0].([]persistedretry.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFailed indicates an expected call of GetFailed
func (mr *MockStoreMockRecorder) GetFailed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFailed", reflect.TypeOf((*MockStore)(nil).GetFailed))
}

// GetPending mocks base method
func (m *MockStore) GetPending() ([]persistedretry.Task, error) {
	ret := m.ctrl.Call(m, "GetPending")
	ret0, _ := ret[0].([]persistedretry.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPending indicates an expected call of GetPending
func (mr *MockStoreMockRecorder) GetPending() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPending", reflect.TypeOf((*MockStore)(nil).GetPending))
}

// MarkDone mocks base method
func (m *MockStore) MarkDone(arg0 persistedretry.Task) error {
	ret := m.ctrl.Call(m, "MarkDone", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDone indicates an expected call of MarkDone
func (mr *MockStoreMockRecorder) MarkDone(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDone", reflect.TypeOf((*MockStore)(nil).MarkDone), arg0)
}

// MarkFailed mocks base method
func (m *MockStore) MarkFailed(arg0 persistedretry.Task) error {
	ret := m.ctrl.Call(m, "MarkFailed", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkFailed indicates an expected call of MarkFailed
func (mr *MockStoreMockRecorder) MarkFailed(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkFailed", reflect.TypeOf((*MockStore)(nil).MarkFailed), arg0)
}

// MarkPending mocks base method
func (m *MockStore) MarkPending(arg0 persistedretry.Task) error {
	ret := m.ctrl.Call(m, "MarkPending", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkPending indicates an expected call of MarkPending
func (mr *MockStoreMockRecorder) MarkPending(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkPending", reflect.TypeOf((*MockStore)(nil).MarkPending), arg0)
}

// MockTask is a mock of Task interface
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMockRecorder
}

// MockTaskMockRecorder is the mock recorder for MockTask
type MockTaskMockRecorder struct {
	mock *MockTask
}

// NewMockTask creates a new mock instance
func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &MockTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTask) EXPECT() *MockTaskMockRecorder {
	return m.recorder
}

// GetLastAttempt mocks base method
func (m *MockTask) GetLastAttempt() time.Time {
	ret := m.ctrl.Call(m, "GetLastAttempt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetLastAttempt indicates an expected call of GetLastAttempt
func (mr *MockTaskMockRecorder) GetLastAttempt() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAttempt", reflect.TypeOf((*MockTask)(nil).GetLastAttempt))
}

// MockExecutor is a mock of Executor interface
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Exec mocks base method
func (m *MockExecutor) Exec(arg0 persistedretry.Task) error {
	ret := m.ctrl.Call(m, "Exec", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec
func (mr *MockExecutorMockRecorder) Exec(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockExecutor)(nil).Exec), arg0)
}

// Name mocks base method
func (m *MockExecutor) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockExecutorMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockExecutor)(nil).Name))
}

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockManager) Add(arg0 persistedretry.Task) error {
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockManagerMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockManager)(nil).Add), arg0)
}

// Close mocks base method
func (m *MockManager) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockManagerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close))
}