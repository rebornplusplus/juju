// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/caas/kubernetes/provider/exec (interfaces: Executor)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	exec "github.com/juju/juju/caas/kubernetes/provider/exec"
	kubernetes "k8s.io/client-go/kubernetes"
	reflect "reflect"
)

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

// Copy mocks base method
func (m *MockExecutor) Copy(arg0 exec.CopyParams, arg1 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Copy indicates an expected call of Copy
func (mr *MockExecutorMockRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockExecutor)(nil).Copy), arg0, arg1)
}

// Exec mocks base method
func (m *MockExecutor) Exec(arg0 exec.ExecParams, arg1 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exec indicates an expected call of Exec
func (mr *MockExecutorMockRecorder) Exec(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockExecutor)(nil).Exec), arg0, arg1)
}

// NameSpace mocks base method
func (m *MockExecutor) NameSpace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NameSpace")
	ret0, _ := ret[0].(string)
	return ret0
}

// NameSpace indicates an expected call of NameSpace
func (mr *MockExecutorMockRecorder) NameSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NameSpace", reflect.TypeOf((*MockExecutor)(nil).NameSpace))
}

// RawClient mocks base method
func (m *MockExecutor) RawClient() kubernetes.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawClient")
	ret0, _ := ret[0].(kubernetes.Interface)
	return ret0
}

// RawClient indicates an expected call of RawClient
func (mr *MockExecutorMockRecorder) RawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawClient", reflect.TypeOf((*MockExecutor)(nil).RawClient))
}

// Status mocks base method
func (m *MockExecutor) Status(arg0 exec.StatusParams) (*exec.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(*exec.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockExecutorMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockExecutor)(nil).Status), arg0)
}
