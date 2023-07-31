// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/application (interfaces: Application,Charm,UpdateBaseState,UpdateBaseValidator,CharmhubClient)

// Package application is a generated GoMock package.
package application

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/charm/v10"
	charmhub "github.com/juju/juju/charmhub"
	transport "github.com/juju/juju/charmhub/transport"
	base "github.com/juju/juju/core/base"
	config "github.com/juju/juju/core/config"
	constraints "github.com/juju/juju/core/constraints"
	state "github.com/juju/juju/state"
	tools "github.com/juju/juju/tools"
	names "github.com/juju/names/v4"
	schema "github.com/juju/schema"
	gomock "go.uber.org/mock/gomock"
	environschema "gopkg.in/juju/environschema.v1"
)

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

// AddUnit mocks base method.
func (m *MockApplication) AddUnit(arg0 state.AddUnitParams) (Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUnit", arg0)
	ret0, _ := ret[0].(Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUnit indicates an expected call of AddUnit.
func (mr *MockApplicationMockRecorder) AddUnit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnit", reflect.TypeOf((*MockApplication)(nil).AddUnit), arg0)
}

// AgentTools mocks base method.
func (m *MockApplication) AgentTools() (*tools.Tools, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentTools")
	ret0, _ := ret[0].(*tools.Tools)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentTools indicates an expected call of AgentTools.
func (mr *MockApplicationMockRecorder) AgentTools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentTools", reflect.TypeOf((*MockApplication)(nil).AgentTools))
}

// AllUnits mocks base method.
func (m *MockApplication) AllUnits() ([]Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllUnits")
	ret0, _ := ret[0].([]Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllUnits indicates an expected call of AllUnits.
func (mr *MockApplicationMockRecorder) AllUnits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllUnits", reflect.TypeOf((*MockApplication)(nil).AllUnits))
}

// ApplicationConfig mocks base method.
func (m *MockApplication) ApplicationConfig() (config.ConfigAttributes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationConfig")
	ret0, _ := ret[0].(config.ConfigAttributes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationConfig indicates an expected call of ApplicationConfig.
func (mr *MockApplicationMockRecorder) ApplicationConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationConfig", reflect.TypeOf((*MockApplication)(nil).ApplicationConfig))
}

// ApplicationTag mocks base method.
func (m *MockApplication) ApplicationTag() names.ApplicationTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationTag")
	ret0, _ := ret[0].(names.ApplicationTag)
	return ret0
}

// ApplicationTag indicates an expected call of ApplicationTag.
func (mr *MockApplicationMockRecorder) ApplicationTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationTag", reflect.TypeOf((*MockApplication)(nil).ApplicationTag))
}

// ChangeScale mocks base method.
func (m *MockApplication) ChangeScale(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeScale", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeScale indicates an expected call of ChangeScale.
func (mr *MockApplicationMockRecorder) ChangeScale(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeScale", reflect.TypeOf((*MockApplication)(nil).ChangeScale), arg0)
}

// Charm mocks base method.
func (m *MockApplication) Charm() (Charm, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm")
	ret0, _ := ret[0].(Charm)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Charm indicates an expected call of Charm.
func (mr *MockApplicationMockRecorder) Charm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockApplication)(nil).Charm))
}

// CharmConfig mocks base method.
func (m *MockApplication) CharmConfig(arg0 string) (charm.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmConfig", arg0)
	ret0, _ := ret[0].(charm.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmConfig indicates an expected call of CharmConfig.
func (mr *MockApplicationMockRecorder) CharmConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmConfig", reflect.TypeOf((*MockApplication)(nil).CharmConfig), arg0)
}

// CharmOrigin mocks base method.
func (m *MockApplication) CharmOrigin() *state.CharmOrigin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmOrigin")
	ret0, _ := ret[0].(*state.CharmOrigin)
	return ret0
}

// CharmOrigin indicates an expected call of CharmOrigin.
func (mr *MockApplicationMockRecorder) CharmOrigin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmOrigin", reflect.TypeOf((*MockApplication)(nil).CharmOrigin))
}

// CharmURL mocks base method.
func (m *MockApplication) CharmURL() (*string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CharmURL indicates an expected call of CharmURL.
func (mr *MockApplicationMockRecorder) CharmURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockApplication)(nil).CharmURL))
}

// ClearExposed mocks base method.
func (m *MockApplication) ClearExposed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearExposed")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearExposed indicates an expected call of ClearExposed.
func (mr *MockApplicationMockRecorder) ClearExposed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearExposed", reflect.TypeOf((*MockApplication)(nil).ClearExposed))
}

// Constraints mocks base method.
func (m *MockApplication) Constraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints.
func (mr *MockApplicationMockRecorder) Constraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockApplication)(nil).Constraints))
}

// Destroy mocks base method.
func (m *MockApplication) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockApplicationMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockApplication)(nil).Destroy))
}

// DestroyOperation mocks base method.
func (m *MockApplication) DestroyOperation() *state.DestroyApplicationOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyOperation")
	ret0, _ := ret[0].(*state.DestroyApplicationOperation)
	return ret0
}

// DestroyOperation indicates an expected call of DestroyOperation.
func (mr *MockApplicationMockRecorder) DestroyOperation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyOperation", reflect.TypeOf((*MockApplication)(nil).DestroyOperation))
}

// EndpointBindings mocks base method.
func (m *MockApplication) EndpointBindings() (Bindings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointBindings")
	ret0, _ := ret[0].(Bindings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndpointBindings indicates an expected call of EndpointBindings.
func (mr *MockApplicationMockRecorder) EndpointBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointBindings", reflect.TypeOf((*MockApplication)(nil).EndpointBindings))
}

// Endpoints mocks base method.
func (m *MockApplication) Endpoints() ([]state.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoints")
	ret0, _ := ret[0].([]state.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Endpoints indicates an expected call of Endpoints.
func (mr *MockApplicationMockRecorder) Endpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoints", reflect.TypeOf((*MockApplication)(nil).Endpoints))
}

// ExposedEndpoints mocks base method.
func (m *MockApplication) ExposedEndpoints() map[string]state.ExposedEndpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExposedEndpoints")
	ret0, _ := ret[0].(map[string]state.ExposedEndpoint)
	return ret0
}

// ExposedEndpoints indicates an expected call of ExposedEndpoints.
func (mr *MockApplicationMockRecorder) ExposedEndpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExposedEndpoints", reflect.TypeOf((*MockApplication)(nil).ExposedEndpoints))
}

// IsExposed mocks base method.
func (m *MockApplication) IsExposed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExposed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExposed indicates an expected call of IsExposed.
func (mr *MockApplicationMockRecorder) IsExposed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExposed", reflect.TypeOf((*MockApplication)(nil).IsExposed))
}

// IsPrincipal mocks base method.
func (m *MockApplication) IsPrincipal() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrincipal")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrincipal indicates an expected call of IsPrincipal.
func (mr *MockApplicationMockRecorder) IsPrincipal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrincipal", reflect.TypeOf((*MockApplication)(nil).IsPrincipal))
}

// IsRemote mocks base method.
func (m *MockApplication) IsRemote() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRemote")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRemote indicates an expected call of IsRemote.
func (mr *MockApplicationMockRecorder) IsRemote() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRemote", reflect.TypeOf((*MockApplication)(nil).IsRemote))
}

// Life mocks base method.
func (m *MockApplication) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockApplicationMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockApplication)(nil).Life))
}

// MergeBindings mocks base method.
func (m *MockApplication) MergeBindings(arg0 *state.Bindings, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeBindings", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MergeBindings indicates an expected call of MergeBindings.
func (mr *MockApplicationMockRecorder) MergeBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeBindings", reflect.TypeOf((*MockApplication)(nil).MergeBindings), arg0, arg1)
}

// MergeExposeSettings mocks base method.
func (m *MockApplication) MergeExposeSettings(arg0 map[string]state.ExposedEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeExposeSettings", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MergeExposeSettings indicates an expected call of MergeExposeSettings.
func (mr *MockApplicationMockRecorder) MergeExposeSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeExposeSettings", reflect.TypeOf((*MockApplication)(nil).MergeExposeSettings), arg0)
}

// Name mocks base method.
func (m *MockApplication) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockApplicationMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockApplication)(nil).Name))
}

// Relations mocks base method.
func (m *MockApplication) Relations() ([]Relation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Relations")
	ret0, _ := ret[0].([]Relation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Relations indicates an expected call of Relations.
func (mr *MockApplicationMockRecorder) Relations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Relations", reflect.TypeOf((*MockApplication)(nil).Relations))
}

// SetCharm mocks base method.
func (m *MockApplication) SetCharm(arg0 state.SetCharmConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharm indicates an expected call of SetCharm.
func (mr *MockApplicationMockRecorder) SetCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockApplication)(nil).SetCharm), arg0)
}

// SetConstraints mocks base method.
func (m *MockApplication) SetConstraints(arg0 constraints.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConstraints", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints.
func (mr *MockApplicationMockRecorder) SetConstraints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockApplication)(nil).SetConstraints), arg0)
}

// SetMetricCredentials mocks base method.
func (m *MockApplication) SetMetricCredentials(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMetricCredentials", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMetricCredentials indicates an expected call of SetMetricCredentials.
func (mr *MockApplicationMockRecorder) SetMetricCredentials(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetricCredentials", reflect.TypeOf((*MockApplication)(nil).SetMetricCredentials), arg0)
}

// SetMinUnits mocks base method.
func (m *MockApplication) SetMinUnits(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMinUnits", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMinUnits indicates an expected call of SetMinUnits.
func (mr *MockApplicationMockRecorder) SetMinUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMinUnits", reflect.TypeOf((*MockApplication)(nil).SetMinUnits), arg0)
}

// SetScale mocks base method.
func (m *MockApplication) SetScale(arg0 int, arg1 int64, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetScale", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetScale indicates an expected call of SetScale.
func (mr *MockApplicationMockRecorder) SetScale(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetScale", reflect.TypeOf((*MockApplication)(nil).SetScale), arg0, arg1, arg2)
}

// UnsetExposeSettings mocks base method.
func (m *MockApplication) UnsetExposeSettings(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsetExposeSettings", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsetExposeSettings indicates an expected call of UnsetExposeSettings.
func (mr *MockApplicationMockRecorder) UnsetExposeSettings(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsetExposeSettings", reflect.TypeOf((*MockApplication)(nil).UnsetExposeSettings), arg0)
}

// UpdateApplicationBase mocks base method.
func (m *MockApplication) UpdateApplicationBase(arg0 state.Base, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationBase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationBase indicates an expected call of UpdateApplicationBase.
func (mr *MockApplicationMockRecorder) UpdateApplicationBase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationBase", reflect.TypeOf((*MockApplication)(nil).UpdateApplicationBase), arg0, arg1)
}

// UpdateApplicationConfig mocks base method.
func (m *MockApplication) UpdateApplicationConfig(arg0 config.ConfigAttributes, arg1 []string, arg2 environschema.Fields, arg3 schema.Defaults) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationConfig indicates an expected call of UpdateApplicationConfig.
func (mr *MockApplicationMockRecorder) UpdateApplicationConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationConfig", reflect.TypeOf((*MockApplication)(nil).UpdateApplicationConfig), arg0, arg1, arg2, arg3)
}

// UpdateCharmConfig mocks base method.
func (m *MockApplication) UpdateCharmConfig(arg0 string, arg1 charm.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCharmConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCharmConfig indicates an expected call of UpdateCharmConfig.
func (mr *MockApplicationMockRecorder) UpdateCharmConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCharmConfig", reflect.TypeOf((*MockApplication)(nil).UpdateCharmConfig), arg0, arg1)
}

// MockCharm is a mock of Charm interface.
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm.
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance.
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockCharm) Config() *charm.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*charm.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockCharmMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCharm)(nil).Config))
}

// Manifest mocks base method.
func (m *MockCharm) Manifest() *charm.Manifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*charm.Manifest)
	return ret0
}

// Manifest indicates an expected call of Manifest.
func (mr *MockCharmMockRecorder) Manifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*MockCharm)(nil).Manifest))
}

// Meta mocks base method.
func (m *MockCharm) Meta() *charm.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(*charm.Meta)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockCharmMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockCharm)(nil).Meta))
}

// String mocks base method.
func (m *MockCharm) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockCharmMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockCharm)(nil).String))
}

// URL mocks base method.
func (m *MockCharm) URL() *charm.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(*charm.URL)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockCharmMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockCharm)(nil).URL))
}

// MockUpdateBaseState is a mock of UpdateBaseState interface.
type MockUpdateBaseState struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateBaseStateMockRecorder
}

// MockUpdateBaseStateMockRecorder is the mock recorder for MockUpdateBaseState.
type MockUpdateBaseStateMockRecorder struct {
	mock *MockUpdateBaseState
}

// NewMockUpdateBaseState creates a new mock instance.
func NewMockUpdateBaseState(ctrl *gomock.Controller) *MockUpdateBaseState {
	mock := &MockUpdateBaseState{ctrl: ctrl}
	mock.recorder = &MockUpdateBaseStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateBaseState) EXPECT() *MockUpdateBaseStateMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockUpdateBaseState) Application(arg0 string) (Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockUpdateBaseStateMockRecorder) Application(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockUpdateBaseState)(nil).Application), arg0)
}

// MockUpdateBaseValidator is a mock of UpdateBaseValidator interface.
type MockUpdateBaseValidator struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateBaseValidatorMockRecorder
}

// MockUpdateBaseValidatorMockRecorder is the mock recorder for MockUpdateBaseValidator.
type MockUpdateBaseValidatorMockRecorder struct {
	mock *MockUpdateBaseValidator
}

// NewMockUpdateBaseValidator creates a new mock instance.
func NewMockUpdateBaseValidator(ctrl *gomock.Controller) *MockUpdateBaseValidator {
	mock := &MockUpdateBaseValidator{ctrl: ctrl}
	mock.recorder = &MockUpdateBaseValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateBaseValidator) EXPECT() *MockUpdateBaseValidatorMockRecorder {
	return m.recorder
}

// ValidateApplication mocks base method.
func (m *MockUpdateBaseValidator) ValidateApplication(arg0 Application, arg1 base.Base, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateApplication", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateApplication indicates an expected call of ValidateApplication.
func (mr *MockUpdateBaseValidatorMockRecorder) ValidateApplication(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateApplication", reflect.TypeOf((*MockUpdateBaseValidator)(nil).ValidateApplication), arg0, arg1, arg2)
}

// MockCharmhubClient is a mock of CharmhubClient interface.
type MockCharmhubClient struct {
	ctrl     *gomock.Controller
	recorder *MockCharmhubClientMockRecorder
}

// MockCharmhubClientMockRecorder is the mock recorder for MockCharmhubClient.
type MockCharmhubClientMockRecorder struct {
	mock *MockCharmhubClient
}

// NewMockCharmhubClient creates a new mock instance.
func NewMockCharmhubClient(ctrl *gomock.Controller) *MockCharmhubClient {
	mock := &MockCharmhubClient{ctrl: ctrl}
	mock.recorder = &MockCharmhubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmhubClient) EXPECT() *MockCharmhubClientMockRecorder {
	return m.recorder
}

// Refresh mocks base method.
func (m *MockCharmhubClient) Refresh(arg0 context.Context, arg1 charmhub.RefreshConfig) ([]transport.RefreshResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0, arg1)
	ret0, _ := ret[0].([]transport.RefreshResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockCharmhubClientMockRecorder) Refresh(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockCharmhubClient)(nil).Refresh), arg0, arg1)
}
