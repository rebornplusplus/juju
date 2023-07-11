// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/stateconverter (interfaces: Machiner,Machine)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	watcher "github.com/juju/juju/core/watcher"
	params "github.com/juju/juju/rpc/params"
	stateconverter "github.com/juju/juju/worker/stateconverter"
	names "github.com/juju/names/v4"
)

// MockMachiner is a mock of Machiner interface.
type MockMachiner struct {
	ctrl     *gomock.Controller
	recorder *MockMachinerMockRecorder
}

// MockMachinerMockRecorder is the mock recorder for MockMachiner.
type MockMachinerMockRecorder struct {
	mock *MockMachiner
}

// NewMockMachiner creates a new mock instance.
func NewMockMachiner(ctrl *gomock.Controller) *MockMachiner {
	mock := &MockMachiner{ctrl: ctrl}
	mock.recorder = &MockMachinerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachiner) EXPECT() *MockMachinerMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockMachiner) Machine(arg0 names.MachineTag) (stateconverter.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(stateconverter.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockMachinerMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockMachiner)(nil).Machine), arg0)
}

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// Jobs mocks base method.
func (m *MockMachine) Jobs() (*params.JobsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jobs")
	ret0, _ := ret[0].(*params.JobsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Jobs indicates an expected call of Jobs.
func (mr *MockMachineMockRecorder) Jobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jobs", reflect.TypeOf((*MockMachine)(nil).Jobs))
}

// Watch mocks base method.
func (m *MockMachine) Watch() (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockMachineMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMachine)(nil).Watch))
}
