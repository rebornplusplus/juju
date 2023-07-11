// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/upgrades (interfaces: Context)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	agent "github.com/juju/juju/agent"
	base "github.com/juju/juju/api/base"
	upgrades "github.com/juju/juju/upgrades"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// APIContext mocks base method.
func (m *MockContext) APIContext() upgrades.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIContext")
	ret0, _ := ret[0].(upgrades.Context)
	return ret0
}

// APIContext indicates an expected call of APIContext.
func (mr *MockContextMockRecorder) APIContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIContext", reflect.TypeOf((*MockContext)(nil).APIContext))
}

// APIState mocks base method.
func (m *MockContext) APIState() base.APICaller {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIState")
	ret0, _ := ret[0].(base.APICaller)
	return ret0
}

// APIState indicates an expected call of APIState.
func (mr *MockContextMockRecorder) APIState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIState", reflect.TypeOf((*MockContext)(nil).APIState))
}

// AgentConfig mocks base method.
func (m *MockContext) AgentConfig() agent.ConfigSetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentConfig")
	ret0, _ := ret[0].(agent.ConfigSetter)
	return ret0
}

// AgentConfig indicates an expected call of AgentConfig.
func (mr *MockContextMockRecorder) AgentConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentConfig", reflect.TypeOf((*MockContext)(nil).AgentConfig))
}

// State mocks base method.
func (m *MockContext) State() upgrades.StateBackend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(upgrades.StateBackend)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockContextMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockContext)(nil).State))
}

// StateContext mocks base method.
func (m *MockContext) StateContext() upgrades.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateContext")
	ret0, _ := ret[0].(upgrades.Context)
	return ret0
}

// StateContext indicates an expected call of StateContext.
func (mr *MockContextMockRecorder) StateContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateContext", reflect.TypeOf((*MockContext)(nil).StateContext))
}
