// Code generated by MockGen. DO NOT EDIT.
// Source: supply.go

// Package supply_test is a generated GoMock package.
package supply_test

import (
	libbuildpack "github.com/SUSE/cf-libbuildpack"
	gomock "github.com/golang/mock/gomock"
	io "io"
	exec "os/exec"
	reflect "reflect"
	cache "ruby/cache"
)

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCommand) Execute(arg0 string, arg1, arg2 io.Writer, arg3 string, arg4 ...string) error {
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockCommandMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommand)(nil).Execute), varargs...)
}

// Output mocks base method
func (m *MockCommand) Output(arg0, arg1 string, arg2 ...string) (string, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Output", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Output indicates an expected call of Output
func (mr *MockCommandMockRecorder) Output(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockCommand)(nil).Output), varargs...)
}

// Run mocks base method
func (m *MockCommand) Run(arg0 *exec.Cmd) error {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockCommandMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCommand)(nil).Run), arg0)
}

// MockManifest is a mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *MockManifestMockRecorder
}

// MockManifestMockRecorder is the mock recorder for MockManifest
type MockManifestMockRecorder struct {
	mock *MockManifest
}

// NewMockManifest creates a new mock instance
func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &MockManifestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManifest) EXPECT() *MockManifestMockRecorder {
	return m.recorder
}

// AllDependencyVersions mocks base method
func (m *MockManifest) AllDependencyVersions(arg0 string) []string {
	ret := m.ctrl.Call(m, "AllDependencyVersions", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// AllDependencyVersions indicates an expected call of AllDependencyVersions
func (mr *MockManifestMockRecorder) AllDependencyVersions(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllDependencyVersions", reflect.TypeOf((*MockManifest)(nil).AllDependencyVersions), arg0)
}

// InstallDependency mocks base method
func (m *MockManifest) InstallDependency(arg0 libbuildpack.Dependency, arg1 string) error {
	ret := m.ctrl.Call(m, "InstallDependency", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependency indicates an expected call of InstallDependency
func (mr *MockManifestMockRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDependency", reflect.TypeOf((*MockManifest)(nil).InstallDependency), arg0, arg1)
}

// InstallOnlyVersion mocks base method
func (m *MockManifest) InstallOnlyVersion(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "InstallOnlyVersion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallOnlyVersion indicates an expected call of InstallOnlyVersion
func (mr *MockManifestMockRecorder) InstallOnlyVersion(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallOnlyVersion", reflect.TypeOf((*MockManifest)(nil).InstallOnlyVersion), arg0, arg1)
}

// DefaultVersion mocks base method
func (m *MockManifest) DefaultVersion(arg0 string) (libbuildpack.Dependency, error) {
	ret := m.ctrl.Call(m, "DefaultVersion", arg0)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultVersion indicates an expected call of DefaultVersion
func (mr *MockManifestMockRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultVersion", reflect.TypeOf((*MockManifest)(nil).DefaultVersion), arg0)
}

// MockVersions is a mock of Versions interface
type MockVersions struct {
	ctrl     *gomock.Controller
	recorder *MockVersionsMockRecorder
}

// MockVersionsMockRecorder is the mock recorder for MockVersions
type MockVersionsMockRecorder struct {
	mock *MockVersions
}

// NewMockVersions creates a new mock instance
func NewMockVersions(ctrl *gomock.Controller) *MockVersions {
	mock := &MockVersions{ctrl: ctrl}
	mock.recorder = &MockVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVersions) EXPECT() *MockVersionsMockRecorder {
	return m.recorder
}

// Engine mocks base method
func (m *MockVersions) Engine() (string, error) {
	ret := m.ctrl.Call(m, "Engine")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Engine indicates an expected call of Engine
func (mr *MockVersionsMockRecorder) Engine() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Engine", reflect.TypeOf((*MockVersions)(nil).Engine))
}

// Version mocks base method
func (m *MockVersions) Version() (string, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockVersionsMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockVersions)(nil).Version))
}

// JrubyVersion mocks base method
func (m *MockVersions) JrubyVersion() (string, error) {
	ret := m.ctrl.Call(m, "JrubyVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JrubyVersion indicates an expected call of JrubyVersion
func (mr *MockVersionsMockRecorder) JrubyVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JrubyVersion", reflect.TypeOf((*MockVersions)(nil).JrubyVersion))
}

// RubyEngineVersion mocks base method
func (m *MockVersions) RubyEngineVersion() (string, error) {
	ret := m.ctrl.Call(m, "RubyEngineVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RubyEngineVersion indicates an expected call of RubyEngineVersion
func (mr *MockVersionsMockRecorder) RubyEngineVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RubyEngineVersion", reflect.TypeOf((*MockVersions)(nil).RubyEngineVersion))
}

// HasGemVersion mocks base method
func (m *MockVersions) HasGemVersion(gem string, constraints ...string) (bool, error) {
	varargs := []interface{}{gem}
	for _, a := range constraints {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasGemVersion", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasGemVersion indicates an expected call of HasGemVersion
func (mr *MockVersionsMockRecorder) HasGemVersion(gem interface{}, constraints ...interface{}) *gomock.Call {
	varargs := append([]interface{}{gem}, constraints...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasGemVersion", reflect.TypeOf((*MockVersions)(nil).HasGemVersion), varargs...)
}

// VersionConstraint mocks base method
func (m *MockVersions) VersionConstraint(version string, constraints ...string) (bool, error) {
	varargs := []interface{}{version}
	for _, a := range constraints {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "VersionConstraint", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VersionConstraint indicates an expected call of VersionConstraint
func (mr *MockVersionsMockRecorder) VersionConstraint(version interface{}, constraints ...interface{}) *gomock.Call {
	varargs := append([]interface{}{version}, constraints...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionConstraint", reflect.TypeOf((*MockVersions)(nil).VersionConstraint), varargs...)
}

// HasWindowsGemfileLock mocks base method
func (m *MockVersions) HasWindowsGemfileLock() (bool, error) {
	ret := m.ctrl.Call(m, "HasWindowsGemfileLock")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasWindowsGemfileLock indicates an expected call of HasWindowsGemfileLock
func (mr *MockVersionsMockRecorder) HasWindowsGemfileLock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasWindowsGemfileLock", reflect.TypeOf((*MockVersions)(nil).HasWindowsGemfileLock))
}

// Gemfile mocks base method
func (m *MockVersions) Gemfile() string {
	ret := m.ctrl.Call(m, "Gemfile")
	ret0, _ := ret[0].(string)
	return ret0
}

// Gemfile indicates an expected call of Gemfile
func (mr *MockVersionsMockRecorder) Gemfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gemfile", reflect.TypeOf((*MockVersions)(nil).Gemfile))
}

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStager) EXPECT() *MockStagerMockRecorder {
	return m.recorder
}

// BuildDir mocks base method
func (m *MockStager) BuildDir() string {
	ret := m.ctrl.Call(m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDir", reflect.TypeOf((*MockStager)(nil).BuildDir))
}

// DepDir mocks base method
func (m *MockStager) DepDir() string {
	ret := m.ctrl.Call(m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepDir", reflect.TypeOf((*MockStager)(nil).DepDir))
}

// DepsIdx mocks base method
func (m *MockStager) DepsIdx() string {
	ret := m.ctrl.Call(m, "DepsIdx")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepsIdx indicates an expected call of DepsIdx
func (mr *MockStagerMockRecorder) DepsIdx() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepsIdx", reflect.TypeOf((*MockStager)(nil).DepsIdx))
}

// LinkDirectoryInDepDir mocks base method
func (m *MockStager) LinkDirectoryInDepDir(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "LinkDirectoryInDepDir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkDirectoryInDepDir indicates an expected call of LinkDirectoryInDepDir
func (mr *MockStagerMockRecorder) LinkDirectoryInDepDir(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LinkDirectoryInDepDir", reflect.TypeOf((*MockStager)(nil).LinkDirectoryInDepDir), arg0, arg1)
}

// WriteEnvFile mocks base method
func (m *MockStager) WriteEnvFile(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "WriteEnvFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteEnvFile indicates an expected call of WriteEnvFile
func (mr *MockStagerMockRecorder) WriteEnvFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteEnvFile", reflect.TypeOf((*MockStager)(nil).WriteEnvFile), arg0, arg1)
}

// WriteProfileD mocks base method
func (m *MockStager) WriteProfileD(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "WriteProfileD", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteProfileD indicates an expected call of WriteProfileD
func (mr *MockStagerMockRecorder) WriteProfileD(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteProfileD", reflect.TypeOf((*MockStager)(nil).WriteProfileD), arg0, arg1)
}

// SetStagingEnvironment mocks base method
func (m *MockStager) SetStagingEnvironment() error {
	ret := m.ctrl.Call(m, "SetStagingEnvironment")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStagingEnvironment indicates an expected call of SetStagingEnvironment
func (mr *MockStagerMockRecorder) SetStagingEnvironment() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStagingEnvironment", reflect.TypeOf((*MockStager)(nil).SetStagingEnvironment))
}

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Metadata mocks base method
func (m *MockCache) Metadata() *cache.Metadata {
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(*cache.Metadata)
	return ret0
}

// Metadata indicates an expected call of Metadata
func (mr *MockCacheMockRecorder) Metadata() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockCache)(nil).Metadata))
}

// Restore mocks base method
func (m *MockCache) Restore() error {
	ret := m.ctrl.Call(m, "Restore")
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockCacheMockRecorder) Restore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockCache)(nil).Restore))
}

// Save mocks base method
func (m *MockCache) Save() error {
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockCacheMockRecorder) Save() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCache)(nil).Save))
}
