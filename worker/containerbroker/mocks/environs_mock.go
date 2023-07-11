// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/environs (interfaces: LXDProfiler,InstanceBroker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
	environs "github.com/juju/juju/environs"
	context "github.com/juju/juju/environs/context"
	instances "github.com/juju/juju/environs/instances"
)

// MockLXDProfiler is a mock of LXDProfiler interface.
type MockLXDProfiler struct {
	ctrl     *gomock.Controller
	recorder *MockLXDProfilerMockRecorder
}

// MockLXDProfilerMockRecorder is the mock recorder for MockLXDProfiler.
type MockLXDProfilerMockRecorder struct {
	mock *MockLXDProfiler
}

// NewMockLXDProfiler creates a new mock instance.
func NewMockLXDProfiler(ctrl *gomock.Controller) *MockLXDProfiler {
	mock := &MockLXDProfiler{ctrl: ctrl}
	mock.recorder = &MockLXDProfilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLXDProfiler) EXPECT() *MockLXDProfilerMockRecorder {
	return m.recorder
}

// AssignLXDProfiles mocks base method.
func (m *MockLXDProfiler) AssignLXDProfiles(arg0 string, arg1 []string, arg2 []lxdprofile.ProfilePost) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignLXDProfiles", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignLXDProfiles indicates an expected call of AssignLXDProfiles.
func (mr *MockLXDProfilerMockRecorder) AssignLXDProfiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignLXDProfiles", reflect.TypeOf((*MockLXDProfiler)(nil).AssignLXDProfiles), arg0, arg1, arg2)
}

// LXDProfileNames mocks base method.
func (m *MockLXDProfiler) LXDProfileNames(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfileNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LXDProfileNames indicates an expected call of LXDProfileNames.
func (mr *MockLXDProfilerMockRecorder) LXDProfileNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfileNames", reflect.TypeOf((*MockLXDProfiler)(nil).LXDProfileNames), arg0)
}

// MaybeWriteLXDProfile mocks base method.
func (m *MockLXDProfiler) MaybeWriteLXDProfile(arg0 string, arg1 lxdprofile.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybeWriteLXDProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaybeWriteLXDProfile indicates an expected call of MaybeWriteLXDProfile.
func (mr *MockLXDProfilerMockRecorder) MaybeWriteLXDProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeWriteLXDProfile", reflect.TypeOf((*MockLXDProfiler)(nil).MaybeWriteLXDProfile), arg0, arg1)
}

// MockInstanceBroker is a mock of InstanceBroker interface.
type MockInstanceBroker struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceBrokerMockRecorder
}

// MockInstanceBrokerMockRecorder is the mock recorder for MockInstanceBroker.
type MockInstanceBrokerMockRecorder struct {
	mock *MockInstanceBroker
}

// NewMockInstanceBroker creates a new mock instance.
func NewMockInstanceBroker(ctrl *gomock.Controller) *MockInstanceBroker {
	mock := &MockInstanceBroker{ctrl: ctrl}
	mock.recorder = &MockInstanceBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstanceBroker) EXPECT() *MockInstanceBrokerMockRecorder {
	return m.recorder
}

// AllInstances mocks base method.
func (m *MockInstanceBroker) AllInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllInstances indicates an expected call of AllInstances.
func (mr *MockInstanceBrokerMockRecorder) AllInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllInstances", reflect.TypeOf((*MockInstanceBroker)(nil).AllInstances), arg0)
}

// AllRunningInstances mocks base method.
func (m *MockInstanceBroker) AllRunningInstances(arg0 context.ProviderCallContext) ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRunningInstances", arg0)
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRunningInstances indicates an expected call of AllRunningInstances.
func (mr *MockInstanceBrokerMockRecorder) AllRunningInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRunningInstances", reflect.TypeOf((*MockInstanceBroker)(nil).AllRunningInstances), arg0)
}

// StartInstance mocks base method.
func (m *MockInstanceBroker) StartInstance(arg0 context.ProviderCallContext, arg1 environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstance", arg0, arg1)
	ret0, _ := ret[0].(*environs.StartInstanceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartInstance indicates an expected call of StartInstance.
func (mr *MockInstanceBrokerMockRecorder) StartInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstance", reflect.TypeOf((*MockInstanceBroker)(nil).StartInstance), arg0, arg1)
}

// StopInstances mocks base method.
func (m *MockInstanceBroker) StopInstances(arg0 context.ProviderCallContext, arg1 ...instance.Id) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopInstances", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopInstances indicates an expected call of StopInstances.
func (mr *MockInstanceBrokerMockRecorder) StopInstances(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstances", reflect.TypeOf((*MockInstanceBroker)(nil).StopInstances), varargs...)
}
