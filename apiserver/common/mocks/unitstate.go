// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: UnitStateBackend,UnitStateUnit)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	common "github.com/juju/juju/apiserver/common"
	controller "github.com/juju/juju/controller"
	state "github.com/juju/juju/state"
)

// MockUnitStateBackend is a mock of UnitStateBackend interface.
type MockUnitStateBackend struct {
	ctrl     *gomock.Controller
	recorder *MockUnitStateBackendMockRecorder
}

// MockUnitStateBackendMockRecorder is the mock recorder for MockUnitStateBackend.
type MockUnitStateBackendMockRecorder struct {
	mock *MockUnitStateBackend
}

// NewMockUnitStateBackend creates a new mock instance.
func NewMockUnitStateBackend(ctrl *gomock.Controller) *MockUnitStateBackend {
	mock := &MockUnitStateBackend{ctrl: ctrl}
	mock.recorder = &MockUnitStateBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitStateBackend) EXPECT() *MockUnitStateBackendMockRecorder {
	return m.recorder
}

// ApplyOperation mocks base method.
func (m *MockUnitStateBackend) ApplyOperation(arg0 state.ModelOperation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyOperation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyOperation indicates an expected call of ApplyOperation.
func (mr *MockUnitStateBackendMockRecorder) ApplyOperation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyOperation", reflect.TypeOf((*MockUnitStateBackend)(nil).ApplyOperation), arg0)
}

// ControllerConfig mocks base method.
func (m *MockUnitStateBackend) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockUnitStateBackendMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockUnitStateBackend)(nil).ControllerConfig))
}

// Unit mocks base method.
func (m *MockUnitStateBackend) Unit(arg0 string) (common.UnitStateUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit", arg0)
	ret0, _ := ret[0].(common.UnitStateUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unit indicates an expected call of Unit.
func (mr *MockUnitStateBackendMockRecorder) Unit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockUnitStateBackend)(nil).Unit), arg0)
}

// MockUnitStateUnit is a mock of UnitStateUnit interface.
type MockUnitStateUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitStateUnitMockRecorder
}

// MockUnitStateUnitMockRecorder is the mock recorder for MockUnitStateUnit.
type MockUnitStateUnitMockRecorder struct {
	mock *MockUnitStateUnit
}

// NewMockUnitStateUnit creates a new mock instance.
func NewMockUnitStateUnit(ctrl *gomock.Controller) *MockUnitStateUnit {
	mock := &MockUnitStateUnit{ctrl: ctrl}
	mock.recorder = &MockUnitStateUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitStateUnit) EXPECT() *MockUnitStateUnitMockRecorder {
	return m.recorder
}

// SetStateOperation mocks base method.
func (m *MockUnitStateUnit) SetStateOperation(arg0 *state.UnitState, arg1 state.UnitStateSizeLimits) state.ModelOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStateOperation", arg0, arg1)
	ret0, _ := ret[0].(state.ModelOperation)
	return ret0
}

// SetStateOperation indicates an expected call of SetStateOperation.
func (mr *MockUnitStateUnitMockRecorder) SetStateOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateOperation", reflect.TypeOf((*MockUnitStateUnit)(nil).SetStateOperation), arg0, arg1)
}

// State mocks base method.
func (m *MockUnitStateUnit) State() (*state.UnitState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(*state.UnitState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockUnitStateUnitMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockUnitStateUnit)(nil).State))
}
