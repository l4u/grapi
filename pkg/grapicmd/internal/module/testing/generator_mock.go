// Code generated by MockGen. DO NOT EDIT.
// Source: generator.go

// Package moduletesting is a generated GoMock package.
package moduletesting

import (
	gomock "github.com/golang/mock/gomock"
	module "github.com/izumin5210/grapi/pkg/grapicmd/internal/module"
	reflect "reflect"
)

// MockGenerator is a mock of Generator interface
type MockGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator
type MockGeneratorMockRecorder struct {
	mock *MockGenerator
}

// NewMockGenerator creates a new mock instance
func NewMockGenerator(ctrl *gomock.Controller) *MockGenerator {
	mock := &MockGenerator{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenerator) EXPECT() *MockGeneratorMockRecorder {
	return m.recorder
}

// GenerateProject mocks base method
func (m *MockGenerator) GenerateProject(rootDir, pkgName string) error {
	ret := m.ctrl.Call(m, "GenerateProject", rootDir, pkgName)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateProject indicates an expected call of GenerateProject
func (mr *MockGeneratorMockRecorder) GenerateProject(rootDir, pkgName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateProject", reflect.TypeOf((*MockGenerator)(nil).GenerateProject), rootDir, pkgName)
}

// GenerateService mocks base method
func (m *MockGenerator) GenerateService(name string, cfg module.ServiceGenerationConfig) error {
	ret := m.ctrl.Call(m, "GenerateService", name, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateService indicates an expected call of GenerateService
func (mr *MockGeneratorMockRecorder) GenerateService(name, cfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateService", reflect.TypeOf((*MockGenerator)(nil).GenerateService), name, cfg)
}

// ScaffoldService mocks base method
func (m *MockGenerator) ScaffoldService(name string, cfg module.ServiceGenerationConfig) error {
	ret := m.ctrl.Call(m, "ScaffoldService", name, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaffoldService indicates an expected call of ScaffoldService
func (mr *MockGeneratorMockRecorder) ScaffoldService(name, cfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaffoldService", reflect.TypeOf((*MockGenerator)(nil).ScaffoldService), name, cfg)
}

// DestroyService mocks base method
func (m *MockGenerator) DestroyService(name string) error {
	ret := m.ctrl.Call(m, "DestroyService", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyService indicates an expected call of DestroyService
func (mr *MockGeneratorMockRecorder) DestroyService(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyService", reflect.TypeOf((*MockGenerator)(nil).DestroyService), name)
}

// GenerateCommand mocks base method
func (m *MockGenerator) GenerateCommand(name string) error {
	ret := m.ctrl.Call(m, "GenerateCommand", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateCommand indicates an expected call of GenerateCommand
func (mr *MockGeneratorMockRecorder) GenerateCommand(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCommand", reflect.TypeOf((*MockGenerator)(nil).GenerateCommand), name)
}

// DestroyCommand mocks base method
func (m *MockGenerator) DestroyCommand(name string) error {
	ret := m.ctrl.Call(m, "DestroyCommand", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyCommand indicates an expected call of DestroyCommand
func (mr *MockGeneratorMockRecorder) DestroyCommand(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyCommand", reflect.TypeOf((*MockGenerator)(nil).DestroyCommand), name)
}

// MockProjectGenerator is a mock of ProjectGenerator interface
type MockProjectGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockProjectGeneratorMockRecorder
}

// MockProjectGeneratorMockRecorder is the mock recorder for MockProjectGenerator
type MockProjectGeneratorMockRecorder struct {
	mock *MockProjectGenerator
}

// NewMockProjectGenerator creates a new mock instance
func NewMockProjectGenerator(ctrl *gomock.Controller) *MockProjectGenerator {
	mock := &MockProjectGenerator{ctrl: ctrl}
	mock.recorder = &MockProjectGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectGenerator) EXPECT() *MockProjectGeneratorMockRecorder {
	return m.recorder
}

// GenerateProject mocks base method
func (m *MockProjectGenerator) GenerateProject(rootDir, pkgName string) error {
	ret := m.ctrl.Call(m, "GenerateProject", rootDir, pkgName)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateProject indicates an expected call of GenerateProject
func (mr *MockProjectGeneratorMockRecorder) GenerateProject(rootDir, pkgName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateProject", reflect.TypeOf((*MockProjectGenerator)(nil).GenerateProject), rootDir, pkgName)
}

// MockServiceGenerator is a mock of ServiceGenerator interface
type MockServiceGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockServiceGeneratorMockRecorder
}

// MockServiceGeneratorMockRecorder is the mock recorder for MockServiceGenerator
type MockServiceGeneratorMockRecorder struct {
	mock *MockServiceGenerator
}

// NewMockServiceGenerator creates a new mock instance
func NewMockServiceGenerator(ctrl *gomock.Controller) *MockServiceGenerator {
	mock := &MockServiceGenerator{ctrl: ctrl}
	mock.recorder = &MockServiceGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceGenerator) EXPECT() *MockServiceGeneratorMockRecorder {
	return m.recorder
}

// GenerateService mocks base method
func (m *MockServiceGenerator) GenerateService(name string, cfg module.ServiceGenerationConfig) error {
	ret := m.ctrl.Call(m, "GenerateService", name, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateService indicates an expected call of GenerateService
func (mr *MockServiceGeneratorMockRecorder) GenerateService(name, cfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateService", reflect.TypeOf((*MockServiceGenerator)(nil).GenerateService), name, cfg)
}

// ScaffoldService mocks base method
func (m *MockServiceGenerator) ScaffoldService(name string, cfg module.ServiceGenerationConfig) error {
	ret := m.ctrl.Call(m, "ScaffoldService", name, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaffoldService indicates an expected call of ScaffoldService
func (mr *MockServiceGeneratorMockRecorder) ScaffoldService(name, cfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaffoldService", reflect.TypeOf((*MockServiceGenerator)(nil).ScaffoldService), name, cfg)
}

// DestroyService mocks base method
func (m *MockServiceGenerator) DestroyService(name string) error {
	ret := m.ctrl.Call(m, "DestroyService", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyService indicates an expected call of DestroyService
func (mr *MockServiceGeneratorMockRecorder) DestroyService(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyService", reflect.TypeOf((*MockServiceGenerator)(nil).DestroyService), name)
}

// MockCommandGenerator is a mock of CommandGenerator interface
type MockCommandGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockCommandGeneratorMockRecorder
}

// MockCommandGeneratorMockRecorder is the mock recorder for MockCommandGenerator
type MockCommandGeneratorMockRecorder struct {
	mock *MockCommandGenerator
}

// NewMockCommandGenerator creates a new mock instance
func NewMockCommandGenerator(ctrl *gomock.Controller) *MockCommandGenerator {
	mock := &MockCommandGenerator{ctrl: ctrl}
	mock.recorder = &MockCommandGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommandGenerator) EXPECT() *MockCommandGeneratorMockRecorder {
	return m.recorder
}

// GenerateCommand mocks base method
func (m *MockCommandGenerator) GenerateCommand(name string) error {
	ret := m.ctrl.Call(m, "GenerateCommand", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateCommand indicates an expected call of GenerateCommand
func (mr *MockCommandGeneratorMockRecorder) GenerateCommand(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCommand", reflect.TypeOf((*MockCommandGenerator)(nil).GenerateCommand), name)
}

// DestroyCommand mocks base method
func (m *MockCommandGenerator) DestroyCommand(name string) error {
	ret := m.ctrl.Call(m, "DestroyCommand", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyCommand indicates an expected call of DestroyCommand
func (mr *MockCommandGeneratorMockRecorder) DestroyCommand(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyCommand", reflect.TypeOf((*MockCommandGenerator)(nil).DestroyCommand), name)
}
