// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/caasapplicationprovisioner (interfaces: CAASProvisionerFacade)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	charms "github.com/juju/juju/api/common/charms"
	caasapplicationprovisioner "github.com/juju/juju/api/controller/caasapplicationprovisioner"
	life "github.com/juju/juju/core/life"
	resources "github.com/juju/juju/core/resources"
	status "github.com/juju/juju/core/status"
	watcher "github.com/juju/juju/core/watcher"
	params "github.com/juju/juju/rpc/params"
)

// MockCAASProvisionerFacade is a mock of CAASProvisionerFacade interface.
type MockCAASProvisionerFacade struct {
	ctrl     *gomock.Controller
	recorder *MockCAASProvisionerFacadeMockRecorder
}

// MockCAASProvisionerFacadeMockRecorder is the mock recorder for MockCAASProvisionerFacade.
type MockCAASProvisionerFacadeMockRecorder struct {
	mock *MockCAASProvisionerFacade
}

// NewMockCAASProvisionerFacade creates a new mock instance.
func NewMockCAASProvisionerFacade(ctrl *gomock.Controller) *MockCAASProvisionerFacade {
	mock := &MockCAASProvisionerFacade{ctrl: ctrl}
	mock.recorder = &MockCAASProvisionerFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAASProvisionerFacade) EXPECT() *MockCAASProvisionerFacadeMockRecorder {
	return m.recorder
}

// ApplicationCharmInfo mocks base method.
func (m *MockCAASProvisionerFacade) ApplicationCharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationCharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationCharmInfo indicates an expected call of ApplicationCharmInfo.
func (mr *MockCAASProvisionerFacadeMockRecorder) ApplicationCharmInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationCharmInfo", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ApplicationCharmInfo), arg0)
}

// ApplicationOCIResources mocks base method.
func (m *MockCAASProvisionerFacade) ApplicationOCIResources(arg0 string) (map[string]resources.DockerImageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationOCIResources", arg0)
	ret0, _ := ret[0].(map[string]resources.DockerImageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationOCIResources indicates an expected call of ApplicationOCIResources.
func (mr *MockCAASProvisionerFacadeMockRecorder) ApplicationOCIResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationOCIResources", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ApplicationOCIResources), arg0)
}

// CharmInfo mocks base method.
func (m *MockCAASProvisionerFacade) CharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo.
func (mr *MockCAASProvisionerFacadeMockRecorder) CharmInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).CharmInfo), arg0)
}

// ClearApplicationResources mocks base method.
func (m *MockCAASProvisionerFacade) ClearApplicationResources(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearApplicationResources", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearApplicationResources indicates an expected call of ClearApplicationResources.
func (mr *MockCAASProvisionerFacadeMockRecorder) ClearApplicationResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearApplicationResources", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ClearApplicationResources), arg0)
}

// DestroyUnits mocks base method.
func (m *MockCAASProvisionerFacade) DestroyUnits(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyUnits", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyUnits indicates an expected call of DestroyUnits.
func (mr *MockCAASProvisionerFacadeMockRecorder) DestroyUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyUnits", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).DestroyUnits), arg0)
}

// Life mocks base method.
func (m *MockCAASProvisionerFacade) Life(arg0 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life", arg0)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Life indicates an expected call of Life.
func (mr *MockCAASProvisionerFacadeMockRecorder) Life(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).Life), arg0)
}

// ProvisionerConfig mocks base method.
func (m *MockCAASProvisionerFacade) ProvisionerConfig() (params.CAASApplicationProvisionerConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvisionerConfig")
	ret0, _ := ret[0].(params.CAASApplicationProvisionerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvisionerConfig indicates an expected call of ProvisionerConfig.
func (mr *MockCAASProvisionerFacadeMockRecorder) ProvisionerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvisionerConfig", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ProvisionerConfig))
}

// ProvisioningInfo mocks base method.
func (m *MockCAASProvisionerFacade) ProvisioningInfo(arg0 string) (caasapplicationprovisioner.ProvisioningInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvisioningInfo", arg0)
	ret0, _ := ret[0].(caasapplicationprovisioner.ProvisioningInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvisioningInfo indicates an expected call of ProvisioningInfo.
func (mr *MockCAASProvisionerFacadeMockRecorder) ProvisioningInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvisioningInfo", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ProvisioningInfo), arg0)
}

// ProvisioningState mocks base method.
func (m *MockCAASProvisionerFacade) ProvisioningState(arg0 string) (*params.CAASApplicationProvisioningState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvisioningState", arg0)
	ret0, _ := ret[0].(*params.CAASApplicationProvisioningState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvisioningState indicates an expected call of ProvisioningState.
func (mr *MockCAASProvisionerFacadeMockRecorder) ProvisioningState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvisioningState", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ProvisioningState), arg0)
}

// RemoveUnit mocks base method.
func (m *MockCAASProvisionerFacade) RemoveUnit(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUnit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUnit indicates an expected call of RemoveUnit.
func (mr *MockCAASProvisionerFacadeMockRecorder) RemoveUnit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUnit", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).RemoveUnit), arg0)
}

// SetOperatorStatus mocks base method.
func (m *MockCAASProvisionerFacade) SetOperatorStatus(arg0 string, arg1 status.Status, arg2 string, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOperatorStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOperatorStatus indicates an expected call of SetOperatorStatus.
func (mr *MockCAASProvisionerFacadeMockRecorder) SetOperatorStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOperatorStatus", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).SetOperatorStatus), arg0, arg1, arg2, arg3)
}

// SetPassword mocks base method.
func (m *MockCAASProvisionerFacade) SetPassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPassword indicates an expected call of SetPassword.
func (mr *MockCAASProvisionerFacadeMockRecorder) SetPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).SetPassword), arg0, arg1)
}

// SetProvisioningState mocks base method.
func (m *MockCAASProvisionerFacade) SetProvisioningState(arg0 string, arg1 params.CAASApplicationProvisioningState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProvisioningState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProvisioningState indicates an expected call of SetProvisioningState.
func (mr *MockCAASProvisionerFacadeMockRecorder) SetProvisioningState(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProvisioningState", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).SetProvisioningState), arg0, arg1)
}

// Units mocks base method.
func (m *MockCAASProvisionerFacade) Units(arg0 string) ([]params.CAASUnit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units", arg0)
	ret0, _ := ret[0].([]params.CAASUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units.
func (mr *MockCAASProvisionerFacadeMockRecorder) Units(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).Units), arg0)
}

// UpdateUnits mocks base method.
func (m *MockCAASProvisionerFacade) UpdateUnits(arg0 params.UpdateApplicationUnits) (*params.UpdateApplicationUnitsInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnits", arg0)
	ret0, _ := ret[0].(*params.UpdateApplicationUnitsInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUnits indicates an expected call of UpdateUnits.
func (mr *MockCAASProvisionerFacadeMockRecorder) UpdateUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnits", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).UpdateUnits), arg0)
}

// WatchApplication mocks base method.
func (m *MockCAASProvisionerFacade) WatchApplication(arg0 string) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplication", arg0)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplication indicates an expected call of WatchApplication.
func (mr *MockCAASProvisionerFacadeMockRecorder) WatchApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplication", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).WatchApplication), arg0)
}

// WatchApplications mocks base method.
func (m *MockCAASProvisionerFacade) WatchApplications() (watcher.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplications")
	ret0, _ := ret[0].(watcher.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplications indicates an expected call of WatchApplications.
func (mr *MockCAASProvisionerFacadeMockRecorder) WatchApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplications", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).WatchApplications))
}

// WatchProvisioningInfo mocks base method.
func (m *MockCAASProvisionerFacade) WatchProvisioningInfo(arg0 string) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchProvisioningInfo", arg0)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchProvisioningInfo indicates an expected call of WatchProvisioningInfo.
func (mr *MockCAASProvisionerFacadeMockRecorder) WatchProvisioningInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchProvisioningInfo", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).WatchProvisioningInfo), arg0)
}

// WatchUnits mocks base method.
func (m *MockCAASProvisionerFacade) WatchUnits(arg0 string) (watcher.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUnits", arg0)
	ret0, _ := ret[0].(watcher.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUnits indicates an expected call of WatchUnits.
func (mr *MockCAASProvisionerFacadeMockRecorder) WatchUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnits", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).WatchUnits), arg0)
}
