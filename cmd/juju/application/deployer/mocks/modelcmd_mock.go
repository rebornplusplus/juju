// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/modelcmd (interfaces: Filesystem)

// Package mocks is a generated GoMock package.
package mocks

import (
	fs "io/fs"
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	modelcmd "github.com/juju/juju/cmd/modelcmd"
)

// MockFilesystem is a mock of Filesystem interface.
type MockFilesystem struct {
	ctrl     *gomock.Controller
	recorder *MockFilesystemMockRecorder
}

// MockFilesystemMockRecorder is the mock recorder for MockFilesystem.
type MockFilesystemMockRecorder struct {
	mock *MockFilesystem
}

// NewMockFilesystem creates a new mock instance.
func NewMockFilesystem(ctrl *gomock.Controller) *MockFilesystem {
	mock := &MockFilesystem{ctrl: ctrl}
	mock.recorder = &MockFilesystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilesystem) EXPECT() *MockFilesystemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFilesystem) Create(arg0 string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFilesystemMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFilesystem)(nil).Create), arg0)
}

// Open mocks base method.
func (m *MockFilesystem) Open(arg0 string) (modelcmd.ReadSeekCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(modelcmd.ReadSeekCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockFilesystemMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockFilesystem)(nil).Open), arg0)
}

// OpenFile mocks base method.
func (m *MockFilesystem) OpenFile(arg0 string, arg1 int, arg2 fs.FileMode) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockFilesystemMockRecorder) OpenFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockFilesystem)(nil).OpenFile), arg0, arg1, arg2)
}

// RemoveAll mocks base method.
func (m *MockFilesystem) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockFilesystemMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockFilesystem)(nil).RemoveAll), arg0)
}

// Stat mocks base method.
func (m *MockFilesystem) Stat(arg0 string) (fs.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", arg0)
	ret0, _ := ret[0].(fs.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockFilesystemMockRecorder) Stat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockFilesystem)(nil).Stat), arg0)
}
