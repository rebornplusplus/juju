// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/deployer (interfaces: Resolver,Bundle)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	charm "github.com/juju/charm/v10"
	charm0 "github.com/juju/juju/api/common/charm"
)

// MockResolver is a mock of Resolver interface.
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver.
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance.
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// GetBundle mocks base method.
func (m *MockResolver) GetBundle(arg0 *charm.URL, arg1 charm0.Origin, arg2 string) (charm.Bundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBundle", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm.Bundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBundle indicates an expected call of GetBundle.
func (mr *MockResolverMockRecorder) GetBundle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBundle", reflect.TypeOf((*MockResolver)(nil).GetBundle), arg0, arg1, arg2)
}

// ResolveBundleURL mocks base method.
func (m *MockResolver) ResolveBundleURL(arg0 *charm.URL, arg1 charm0.Origin) (*charm.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBundleURL", arg0, arg1)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResolveBundleURL indicates an expected call of ResolveBundleURL.
func (mr *MockResolverMockRecorder) ResolveBundleURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBundleURL", reflect.TypeOf((*MockResolver)(nil).ResolveBundleURL), arg0, arg1)
}

// ResolveCharm mocks base method.
func (m *MockResolver) ResolveCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 bool) (*charm.URL, charm0.Origin, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveCharm indicates an expected call of ResolveCharm.
func (mr *MockResolverMockRecorder) ResolveCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveCharm", reflect.TypeOf((*MockResolver)(nil).ResolveCharm), arg0, arg1, arg2)
}

// MockBundle is a mock of Bundle interface.
type MockBundle struct {
	ctrl     *gomock.Controller
	recorder *MockBundleMockRecorder
}

// MockBundleMockRecorder is the mock recorder for MockBundle.
type MockBundleMockRecorder struct {
	mock *MockBundle
}

// NewMockBundle creates a new mock instance.
func NewMockBundle(ctrl *gomock.Controller) *MockBundle {
	mock := &MockBundle{ctrl: ctrl}
	mock.recorder = &MockBundleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBundle) EXPECT() *MockBundleMockRecorder {
	return m.recorder
}

// ContainsOverlays mocks base method.
func (m *MockBundle) ContainsOverlays() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsOverlays")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainsOverlays indicates an expected call of ContainsOverlays.
func (mr *MockBundleMockRecorder) ContainsOverlays() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsOverlays", reflect.TypeOf((*MockBundle)(nil).ContainsOverlays))
}

// Data mocks base method.
func (m *MockBundle) Data() *charm.BundleData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(*charm.BundleData)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockBundleMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockBundle)(nil).Data))
}

// ReadMe mocks base method.
func (m *MockBundle) ReadMe() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMe")
	ret0, _ := ret[0].(string)
	return ret0
}

// ReadMe indicates an expected call of ReadMe.
func (mr *MockBundleMockRecorder) ReadMe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMe", reflect.TypeOf((*MockBundle)(nil).ReadMe))
}
