// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/observer/metricobserver (interfaces: MetricsCollector,SummaryVec)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	metricobserver "github.com/juju/juju/apiserver/observer/metricobserver"
	prometheus "github.com/prometheus/client_golang/prometheus"
)

// MockMetricsCollector is a mock of MetricsCollector interface.
type MockMetricsCollector struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsCollectorMockRecorder
}

// MockMetricsCollectorMockRecorder is the mock recorder for MockMetricsCollector.
type MockMetricsCollectorMockRecorder struct {
	mock *MockMetricsCollector
}

// NewMockMetricsCollector creates a new mock instance.
func NewMockMetricsCollector(ctrl *gomock.Controller) *MockMetricsCollector {
	mock := &MockMetricsCollector{ctrl: ctrl}
	mock.recorder = &MockMetricsCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsCollector) EXPECT() *MockMetricsCollectorMockRecorder {
	return m.recorder
}

// APIRequestDuration mocks base method.
func (m *MockMetricsCollector) APIRequestDuration() metricobserver.SummaryVec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIRequestDuration")
	ret0, _ := ret[0].(metricobserver.SummaryVec)
	return ret0
}

// APIRequestDuration indicates an expected call of APIRequestDuration.
func (mr *MockMetricsCollectorMockRecorder) APIRequestDuration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIRequestDuration", reflect.TypeOf((*MockMetricsCollector)(nil).APIRequestDuration))
}

// MockSummaryVec is a mock of SummaryVec interface.
type MockSummaryVec struct {
	ctrl     *gomock.Controller
	recorder *MockSummaryVecMockRecorder
}

// MockSummaryVecMockRecorder is the mock recorder for MockSummaryVec.
type MockSummaryVecMockRecorder struct {
	mock *MockSummaryVec
}

// NewMockSummaryVec creates a new mock instance.
func NewMockSummaryVec(ctrl *gomock.Controller) *MockSummaryVec {
	mock := &MockSummaryVec{ctrl: ctrl}
	mock.recorder = &MockSummaryVecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSummaryVec) EXPECT() *MockSummaryVecMockRecorder {
	return m.recorder
}

// With mocks base method.
func (m *MockSummaryVec) With(arg0 prometheus.Labels) prometheus.Observer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "With", arg0)
	ret0, _ := ret[0].(prometheus.Observer)
	return ret0
}

// With indicates an expected call of With.
func (mr *MockSummaryVecMockRecorder) With(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockSummaryVec)(nil).With), arg0)
}
