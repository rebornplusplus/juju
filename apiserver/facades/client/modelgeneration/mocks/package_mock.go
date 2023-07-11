// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/modelgeneration (interfaces: State,Model,Generation,Application,ModelCache)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	charm "github.com/juju/charm/v10"
	modelgeneration "github.com/juju/juju/apiserver/facades/client/modelgeneration"
	cache "github.com/juju/juju/core/cache"
	settings "github.com/juju/juju/core/settings"
	names "github.com/juju/names/v4"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockState) Application(arg0 string) (modelgeneration.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(modelgeneration.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockState)(nil).Application), arg0)
}

// ControllerTag mocks base method.
func (m *MockState) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockStateMockRecorder) ControllerTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockState)(nil).ControllerTag))
}

// Model mocks base method.
func (m *MockState) Model() (modelgeneration.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(modelgeneration.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Model indicates an expected call of Model.
func (mr *MockStateMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockState)(nil).Model))
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// AddBranch mocks base method.
func (m *MockModel) AddBranch(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBranch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBranch indicates an expected call of AddBranch.
func (mr *MockModelMockRecorder) AddBranch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBranch", reflect.TypeOf((*MockModel)(nil).AddBranch), arg0, arg1)
}

// Branch mocks base method.
func (m *MockModel) Branch(arg0 string) (modelgeneration.Generation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Branch", arg0)
	ret0, _ := ret[0].(modelgeneration.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Branch indicates an expected call of Branch.
func (mr *MockModelMockRecorder) Branch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Branch", reflect.TypeOf((*MockModel)(nil).Branch), arg0)
}

// Branches mocks base method.
func (m *MockModel) Branches() ([]modelgeneration.Generation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Branches")
	ret0, _ := ret[0].([]modelgeneration.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Branches indicates an expected call of Branches.
func (mr *MockModelMockRecorder) Branches() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Branches", reflect.TypeOf((*MockModel)(nil).Branches))
}

// Generation mocks base method.
func (m *MockModel) Generation(arg0 int) (modelgeneration.Generation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generation", arg0)
	ret0, _ := ret[0].(modelgeneration.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generation indicates an expected call of Generation.
func (mr *MockModelMockRecorder) Generation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generation", reflect.TypeOf((*MockModel)(nil).Generation), arg0)
}

// Generations mocks base method.
func (m *MockModel) Generations() ([]modelgeneration.Generation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generations")
	ret0, _ := ret[0].([]modelgeneration.Generation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generations indicates an expected call of Generations.
func (mr *MockModelMockRecorder) Generations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generations", reflect.TypeOf((*MockModel)(nil).Generations))
}

// ModelTag mocks base method.
func (m *MockModel) ModelTag() names.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockModelMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockModel)(nil).ModelTag))
}

// MockGeneration is a mock of Generation interface.
type MockGeneration struct {
	ctrl     *gomock.Controller
	recorder *MockGenerationMockRecorder
}

// MockGenerationMockRecorder is the mock recorder for MockGeneration.
type MockGenerationMockRecorder struct {
	mock *MockGeneration
}

// NewMockGeneration creates a new mock instance.
func NewMockGeneration(ctrl *gomock.Controller) *MockGeneration {
	mock := &MockGeneration{ctrl: ctrl}
	mock.recorder = &MockGenerationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeneration) EXPECT() *MockGenerationMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockGeneration) Abort(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Abort", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Abort indicates an expected call of Abort.
func (mr *MockGenerationMockRecorder) Abort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockGeneration)(nil).Abort), arg0)
}

// AssignAllUnits mocks base method.
func (m *MockGeneration) AssignAllUnits(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignAllUnits", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignAllUnits indicates an expected call of AssignAllUnits.
func (mr *MockGenerationMockRecorder) AssignAllUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignAllUnits", reflect.TypeOf((*MockGeneration)(nil).AssignAllUnits), arg0)
}

// AssignUnit mocks base method.
func (m *MockGeneration) AssignUnit(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignUnit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignUnit indicates an expected call of AssignUnit.
func (mr *MockGenerationMockRecorder) AssignUnit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignUnit", reflect.TypeOf((*MockGeneration)(nil).AssignUnit), arg0)
}

// AssignUnits mocks base method.
func (m *MockGeneration) AssignUnits(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignUnits", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignUnits indicates an expected call of AssignUnits.
func (mr *MockGenerationMockRecorder) AssignUnits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignUnits", reflect.TypeOf((*MockGeneration)(nil).AssignUnits), arg0, arg1)
}

// AssignedUnits mocks base method.
func (m *MockGeneration) AssignedUnits() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedUnits")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// AssignedUnits indicates an expected call of AssignedUnits.
func (mr *MockGenerationMockRecorder) AssignedUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedUnits", reflect.TypeOf((*MockGeneration)(nil).AssignedUnits))
}

// BranchName mocks base method.
func (m *MockGeneration) BranchName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BranchName")
	ret0, _ := ret[0].(string)
	return ret0
}

// BranchName indicates an expected call of BranchName.
func (mr *MockGenerationMockRecorder) BranchName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BranchName", reflect.TypeOf((*MockGeneration)(nil).BranchName))
}

// Commit mocks base method.
func (m *MockGeneration) Commit(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockGenerationMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockGeneration)(nil).Commit), arg0)
}

// Completed mocks base method.
func (m *MockGeneration) Completed() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Completed")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Completed indicates an expected call of Completed.
func (mr *MockGenerationMockRecorder) Completed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Completed", reflect.TypeOf((*MockGeneration)(nil).Completed))
}

// CompletedBy mocks base method.
func (m *MockGeneration) CompletedBy() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompletedBy")
	ret0, _ := ret[0].(string)
	return ret0
}

// CompletedBy indicates an expected call of CompletedBy.
func (mr *MockGenerationMockRecorder) CompletedBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompletedBy", reflect.TypeOf((*MockGeneration)(nil).CompletedBy))
}

// Config mocks base method.
func (m *MockGeneration) Config() map[string]settings.ItemChanges {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(map[string]settings.ItemChanges)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockGenerationMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockGeneration)(nil).Config))
}

// Created mocks base method.
func (m *MockGeneration) Created() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Created")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Created indicates an expected call of Created.
func (mr *MockGenerationMockRecorder) Created() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Created", reflect.TypeOf((*MockGeneration)(nil).Created))
}

// CreatedBy mocks base method.
func (m *MockGeneration) CreatedBy() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatedBy")
	ret0, _ := ret[0].(string)
	return ret0
}

// CreatedBy indicates an expected call of CreatedBy.
func (mr *MockGenerationMockRecorder) CreatedBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedBy", reflect.TypeOf((*MockGeneration)(nil).CreatedBy))
}

// GenerationId mocks base method.
func (m *MockGeneration) GenerationId() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerationId")
	ret0, _ := ret[0].(int)
	return ret0
}

// GenerationId indicates an expected call of GenerationId.
func (mr *MockGenerationMockRecorder) GenerationId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerationId", reflect.TypeOf((*MockGeneration)(nil).GenerationId))
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// DefaultCharmConfig mocks base method.
func (m *MockApplication) DefaultCharmConfig() (charm.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultCharmConfig")
	ret0, _ := ret[0].(charm.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultCharmConfig indicates an expected call of DefaultCharmConfig.
func (mr *MockApplicationMockRecorder) DefaultCharmConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultCharmConfig", reflect.TypeOf((*MockApplication)(nil).DefaultCharmConfig))
}

// UnitNames mocks base method.
func (m *MockApplication) UnitNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnitNames indicates an expected call of UnitNames.
func (mr *MockApplicationMockRecorder) UnitNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitNames", reflect.TypeOf((*MockApplication)(nil).UnitNames))
}

// MockModelCache is a mock of ModelCache interface.
type MockModelCache struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheMockRecorder
}

// MockModelCacheMockRecorder is the mock recorder for MockModelCache.
type MockModelCacheMockRecorder struct {
	mock *MockModelCache
}

// NewMockModelCache creates a new mock instance.
func NewMockModelCache(ctrl *gomock.Controller) *MockModelCache {
	mock := &MockModelCache{ctrl: ctrl}
	mock.recorder = &MockModelCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelCache) EXPECT() *MockModelCacheMockRecorder {
	return m.recorder
}

// Branch mocks base method.
func (m *MockModelCache) Branch(arg0 string) (cache.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Branch", arg0)
	ret0, _ := ret[0].(cache.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Branch indicates an expected call of Branch.
func (mr *MockModelCacheMockRecorder) Branch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Branch", reflect.TypeOf((*MockModelCache)(nil).Branch), arg0)
}
