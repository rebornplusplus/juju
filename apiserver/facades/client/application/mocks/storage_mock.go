// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/storage (interfaces: ProviderRegistry)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	storage "github.com/juju/juju/storage"
)

// MockProviderRegistry is a mock of ProviderRegistry interface.
type MockProviderRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockProviderRegistryMockRecorder
}

// MockProviderRegistryMockRecorder is the mock recorder for MockProviderRegistry.
type MockProviderRegistryMockRecorder struct {
	mock *MockProviderRegistry
}

// NewMockProviderRegistry creates a new mock instance.
func NewMockProviderRegistry(ctrl *gomock.Controller) *MockProviderRegistry {
	mock := &MockProviderRegistry{ctrl: ctrl}
	mock.recorder = &MockProviderRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProviderRegistry) EXPECT() *MockProviderRegistryMockRecorder {
	return m.recorder
}

// StorageProvider mocks base method.
func (m *MockProviderRegistry) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider.
func (mr *MockProviderRegistryMockRecorder) StorageProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockProviderRegistry)(nil).StorageProvider), arg0)
}

// StorageProviderTypes mocks base method.
func (m *MockProviderRegistry) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes.
func (mr *MockProviderRegistryMockRecorder) StorageProviderTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockProviderRegistry)(nil).StorageProviderTypes))
}
