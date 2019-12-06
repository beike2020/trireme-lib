// Code generated by MockGen. DO NOT EDIT.
// Source: controller/interfaces.go

// Package mockcontroller is a generated GoMock package.
package mockcontroller

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	packettracing "go.aporeto.io/trireme-lib/controller/pkg/packettracing"
	secrets "go.aporeto.io/trireme-lib/controller/pkg/secrets"
	runtime "go.aporeto.io/trireme-lib/controller/runtime"
	policy "go.aporeto.io/trireme-lib/policy"
)

// MockTriremeController is a mock of TriremeController interface
// nolint
type MockTriremeController struct {
	ctrl     *gomock.Controller
	recorder *MockTriremeControllerMockRecorder
}

// MockTriremeControllerMockRecorder is the mock recorder for MockTriremeController
// nolint
type MockTriremeControllerMockRecorder struct {
	mock *MockTriremeController
}

// NewMockTriremeController creates a new mock instance
// nolint
func NewMockTriremeController(ctrl *gomock.Controller) *MockTriremeController {
	mock := &MockTriremeController{ctrl: ctrl}
	mock.recorder = &MockTriremeControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockTriremeController) EXPECT() *MockTriremeControllerMockRecorder {
	return m.recorder
}

// Run mocks base method
// nolint
func (m *MockTriremeController) Run(ctx context.Context) error {
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
// nolint
func (mr *MockTriremeControllerMockRecorder) Run(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTriremeController)(nil).Run), ctx)
}

// CleanUp mocks base method
// nolint
func (m *MockTriremeController) CleanUp() error {
	ret := m.ctrl.Call(m, "CleanUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUp indicates an expected call of CleanUp
// nolint
func (mr *MockTriremeControllerMockRecorder) CleanUp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUp", reflect.TypeOf((*MockTriremeController)(nil).CleanUp))
}

// Enforce mocks base method
// nolint
func (m *MockTriremeController) Enforce(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error {
	ret := m.ctrl.Call(m, "Enforce", ctx, puID, policy, runtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enforce indicates an expected call of Enforce
// nolint
func (mr *MockTriremeControllerMockRecorder) Enforce(ctx, puID, policy, runtime interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enforce", reflect.TypeOf((*MockTriremeController)(nil).Enforce), ctx, puID, policy, runtime)
}

// UnEnforce mocks base method
// nolint
func (m *MockTriremeController) UnEnforce(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error {
	ret := m.ctrl.Call(m, "UnEnforce", ctx, puID, policy, runtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnEnforce indicates an expected call of UnEnforce
// nolint
func (mr *MockTriremeControllerMockRecorder) UnEnforce(ctx, puID, policy, runtime interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnEnforce", reflect.TypeOf((*MockTriremeController)(nil).UnEnforce), ctx, puID, policy, runtime)
}

// UpdatePolicy mocks base method
// nolint
func (m *MockTriremeController) UpdatePolicy(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error {
	ret := m.ctrl.Call(m, "UpdatePolicy", ctx, puID, policy, runtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePolicy indicates an expected call of UpdatePolicy
// nolint
func (mr *MockTriremeControllerMockRecorder) UpdatePolicy(ctx, puID, policy, runtime interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicy", reflect.TypeOf((*MockTriremeController)(nil).UpdatePolicy), ctx, puID, policy, runtime)
}

// UpdateSecrets mocks base method
// nolint
func (m *MockTriremeController) UpdateSecrets(secrets secrets.Secrets) error {
	ret := m.ctrl.Call(m, "UpdateSecrets", secrets)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecrets indicates an expected call of UpdateSecrets
// nolint
func (mr *MockTriremeControllerMockRecorder) UpdateSecrets(secrets interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecrets", reflect.TypeOf((*MockTriremeController)(nil).UpdateSecrets), secrets)
}

// UpdateConfiguration mocks base method
// nolint
func (m *MockTriremeController) UpdateConfiguration(cfg *runtime.Configuration) error {
	ret := m.ctrl.Call(m, "UpdateConfiguration", cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration
// nolint
func (mr *MockTriremeControllerMockRecorder) UpdateConfiguration(cfg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockTriremeController)(nil).UpdateConfiguration), cfg)
}

// EnableDatapathPacketTracing mocks base method
// nolint
func (m *MockTriremeController) EnableDatapathPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, direction packettracing.TracingDirection, interval time.Duration) error {
	ret := m.ctrl.Call(m, "EnableDatapathPacketTracing", ctx, puID, policy, runtime, direction, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDatapathPacketTracing indicates an expected call of EnableDatapathPacketTracing
// nolint
func (mr *MockTriremeControllerMockRecorder) EnableDatapathPacketTracing(ctx, puID, policy, runtime, direction, interval interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDatapathPacketTracing", reflect.TypeOf((*MockTriremeController)(nil).EnableDatapathPacketTracing), ctx, puID, policy, runtime, direction, interval)
}

// EnableIPTablesPacketTracing mocks base method
// nolint
func (m *MockTriremeController) EnableIPTablesPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, interval time.Duration) error {
	ret := m.ctrl.Call(m, "EnableIPTablesPacketTracing", ctx, puID, policy, runtime, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableIPTablesPacketTracing indicates an expected call of EnableIPTablesPacketTracing
// nolint
func (mr *MockTriremeControllerMockRecorder) EnableIPTablesPacketTracing(ctx, puID, policy, runtime, interval interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPTablesPacketTracing", reflect.TypeOf((*MockTriremeController)(nil).EnableIPTablesPacketTracing), ctx, puID, policy, runtime, interval)
}

// RunDiagnostics mocks base method
// nolint
func (m *MockTriremeController) RunDiagnostics(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, diagnosticsInfo *policy.DiagnosticsConfig) error {
	ret := m.ctrl.Call(m, "RunDiagnostics", ctx, puID, policy, runtime, diagnosticsInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDiagnostics indicates an expected call of RunDiagnostics
// nolint
func (mr *MockTriremeControllerMockRecorder) RunDiagnostics(ctx, puID, policy, runtime, diagnosticsInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDiagnostics", reflect.TypeOf((*MockTriremeController)(nil).RunDiagnostics), ctx, puID, policy, runtime, diagnosticsInfo)
}

// MockDebugInfo is a mock of DebugInfo interface
// nolint
type MockDebugInfo struct {
	ctrl     *gomock.Controller
	recorder *MockDebugInfoMockRecorder
}

// MockDebugInfoMockRecorder is the mock recorder for MockDebugInfo
// nolint
type MockDebugInfoMockRecorder struct {
	mock *MockDebugInfo
}

// NewMockDebugInfo creates a new mock instance
// nolint
func NewMockDebugInfo(ctrl *gomock.Controller) *MockDebugInfo {
	mock := &MockDebugInfo{ctrl: ctrl}
	mock.recorder = &MockDebugInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockDebugInfo) EXPECT() *MockDebugInfoMockRecorder {
	return m.recorder
}

// EnableDatapathPacketTracing mocks base method
// nolint
func (m *MockDebugInfo) EnableDatapathPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, direction packettracing.TracingDirection, interval time.Duration) error {
	ret := m.ctrl.Call(m, "EnableDatapathPacketTracing", ctx, puID, policy, runtime, direction, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDatapathPacketTracing indicates an expected call of EnableDatapathPacketTracing
// nolint
func (mr *MockDebugInfoMockRecorder) EnableDatapathPacketTracing(ctx, puID, policy, runtime, direction, interval interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDatapathPacketTracing", reflect.TypeOf((*MockDebugInfo)(nil).EnableDatapathPacketTracing), ctx, puID, policy, runtime, direction, interval)
}

// EnableIPTablesPacketTracing mocks base method
// nolint
func (m *MockDebugInfo) EnableIPTablesPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, interval time.Duration) error {
	ret := m.ctrl.Call(m, "EnableIPTablesPacketTracing", ctx, puID, policy, runtime, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableIPTablesPacketTracing indicates an expected call of EnableIPTablesPacketTracing
// nolint
func (mr *MockDebugInfoMockRecorder) EnableIPTablesPacketTracing(ctx, puID, policy, runtime, interval interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPTablesPacketTracing", reflect.TypeOf((*MockDebugInfo)(nil).EnableIPTablesPacketTracing), ctx, puID, policy, runtime, interval)
}

// RunDiagnostics mocks base method
// nolint
func (m *MockDebugInfo) RunDiagnostics(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, diagnosticsInfo *policy.DiagnosticsConfig) error {
	ret := m.ctrl.Call(m, "RunDiagnostics", ctx, puID, policy, runtime, diagnosticsInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunDiagnostics indicates an expected call of RunDiagnostics
// nolint
func (mr *MockDebugInfoMockRecorder) RunDiagnostics(ctx, puID, policy, runtime, diagnosticsInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunDiagnostics", reflect.TypeOf((*MockDebugInfo)(nil).RunDiagnostics), ctx, puID, policy, runtime, diagnosticsInfo)
}
