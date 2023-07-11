// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/containeragent/initialize (interfaces: ApplicationAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	caasapplication "github.com/juju/juju/api/agent/caasapplication"
)

// MockApplicationAPI is a mock of ApplicationAPI interface.
type MockApplicationAPI struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationAPIMockRecorder
}

// MockApplicationAPIMockRecorder is the mock recorder for MockApplicationAPI.
type MockApplicationAPIMockRecorder struct {
	mock *MockApplicationAPI
}

// NewMockApplicationAPI creates a new mock instance.
func NewMockApplicationAPI(ctrl *gomock.Controller) *MockApplicationAPI {
	mock := &MockApplicationAPI{ctrl: ctrl}
	mock.recorder = &MockApplicationAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationAPI) EXPECT() *MockApplicationAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockApplicationAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockApplicationAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockApplicationAPI)(nil).Close))
}

// UnitIntroduction mocks base method.
func (m *MockApplicationAPI) UnitIntroduction(arg0, arg1 string) (*caasapplication.UnitConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitIntroduction", arg0, arg1)
	ret0, _ := ret[0].(*caasapplication.UnitConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitIntroduction indicates an expected call of UnitIntroduction.
func (mr *MockApplicationAPIMockRecorder) UnitIntroduction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitIntroduction", reflect.TypeOf((*MockApplicationAPI)(nil).UnitIntroduction), arg0, arg1)
}
