// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/secretbackendmanager (interfaces: BackendRotate)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	state "github.com/juju/juju/state"
)

// MockBackendRotate is a mock of BackendRotate interface.
type MockBackendRotate struct {
	ctrl     *gomock.Controller
	recorder *MockBackendRotateMockRecorder
}

// MockBackendRotateMockRecorder is the mock recorder for MockBackendRotate.
type MockBackendRotateMockRecorder struct {
	mock *MockBackendRotate
}

// NewMockBackendRotate creates a new mock instance.
func NewMockBackendRotate(ctrl *gomock.Controller) *MockBackendRotate {
	mock := &MockBackendRotate{ctrl: ctrl}
	mock.recorder = &MockBackendRotateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendRotate) EXPECT() *MockBackendRotateMockRecorder {
	return m.recorder
}

// WatchSecretBackendRotationChanges mocks base method.
func (m *MockBackendRotate) WatchSecretBackendRotationChanges() (state.SecretBackendRotateWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchSecretBackendRotationChanges")
	ret0, _ := ret[0].(state.SecretBackendRotateWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchSecretBackendRotationChanges indicates an expected call of WatchSecretBackendRotationChanges.
func (mr *MockBackendRotateMockRecorder) WatchSecretBackendRotationChanges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchSecretBackendRotationChanges", reflect.TypeOf((*MockBackendRotate)(nil).WatchSecretBackendRotationChanges))
}
