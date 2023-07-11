// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/logger (interfaces: Logger,LoggerCloser)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	logger "github.com/juju/juju/core/logger"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLogger) Log(arg0 []logger.LogRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockLoggerMockRecorder) Log(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), arg0)
}

// MockLoggerCloser is a mock of LoggerCloser interface.
type MockLoggerCloser struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerCloserMockRecorder
}

// MockLoggerCloserMockRecorder is the mock recorder for MockLoggerCloser.
type MockLoggerCloserMockRecorder struct {
	mock *MockLoggerCloser
}

// NewMockLoggerCloser creates a new mock instance.
func NewMockLoggerCloser(ctrl *gomock.Controller) *MockLoggerCloser {
	mock := &MockLoggerCloser{ctrl: ctrl}
	mock.recorder = &MockLoggerCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggerCloser) EXPECT() *MockLoggerCloserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockLoggerCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLoggerCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLoggerCloser)(nil).Close))
}

// Log mocks base method.
func (m *MockLoggerCloser) Log(arg0 []logger.LogRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MockLoggerCloserMockRecorder) Log(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLoggerCloser)(nil).Log), arg0)
}
