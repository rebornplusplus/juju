// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: EntityFinder,Entity,Lifer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
)

// MockEntityFinder is a mock of EntityFinder interface.
type MockEntityFinder struct {
	ctrl     *gomock.Controller
	recorder *MockEntityFinderMockRecorder
}

// MockEntityFinderMockRecorder is the mock recorder for MockEntityFinder.
type MockEntityFinderMockRecorder struct {
	mock *MockEntityFinder
}

// NewMockEntityFinder creates a new mock instance.
func NewMockEntityFinder(ctrl *gomock.Controller) *MockEntityFinder {
	mock := &MockEntityFinder{ctrl: ctrl}
	mock.recorder = &MockEntityFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntityFinder) EXPECT() *MockEntityFinderMockRecorder {
	return m.recorder
}

// FindEntity mocks base method.
func (m *MockEntityFinder) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockEntityFinderMockRecorder) FindEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockEntityFinder)(nil).FindEntity), arg0)
}

// MockEntity is a mock of Entity interface.
type MockEntity struct {
	ctrl     *gomock.Controller
	recorder *MockEntityMockRecorder
}

// MockEntityMockRecorder is the mock recorder for MockEntity.
type MockEntityMockRecorder struct {
	mock *MockEntity
}

// NewMockEntity creates a new mock instance.
func NewMockEntity(ctrl *gomock.Controller) *MockEntity {
	mock := &MockEntity{ctrl: ctrl}
	mock.recorder = &MockEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEntity) EXPECT() *MockEntityMockRecorder {
	return m.recorder
}

// Tag mocks base method.
func (m *MockEntity) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockEntityMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockEntity)(nil).Tag))
}

// MockLifer is a mock of Lifer interface.
type MockLifer struct {
	ctrl     *gomock.Controller
	recorder *MockLiferMockRecorder
}

// MockLiferMockRecorder is the mock recorder for MockLifer.
type MockLiferMockRecorder struct {
	mock *MockLifer
}

// NewMockLifer creates a new mock instance.
func NewMockLifer(ctrl *gomock.Controller) *MockLifer {
	mock := &MockLifer{ctrl: ctrl}
	mock.recorder = &MockLiferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLifer) EXPECT() *MockLiferMockRecorder {
	return m.recorder
}

// Life mocks base method.
func (m *MockLifer) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockLiferMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockLifer)(nil).Life))
}
