// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/edwardmelvin/quick-core/features/outbound (interfaces: Manager,HandlerSelector)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	outbound "github.com/edwardmelvin/quick-core/features/outbound"
)

// OutboundManager is a mock of Manager interface
type OutboundManager struct {
	ctrl     *gomock.Controller
	recorder *OutboundManagerMockRecorder
}

// OutboundManagerMockRecorder is the mock recorder for OutboundManager
type OutboundManagerMockRecorder struct {
	mock *OutboundManager
}

// NewOutboundManager creates a new mock instance
func NewOutboundManager(ctrl *gomock.Controller) *OutboundManager {
	mock := &OutboundManager{ctrl: ctrl}
	mock.recorder = &OutboundManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *OutboundManager) EXPECT() *OutboundManagerMockRecorder {
	return m.recorder
}

// AddHandler mocks base method
func (m *OutboundManager) AddHandler(arg0 context.Context, arg1 outbound.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddHandler", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddHandler indicates an expected call of AddHandler
func (mr *OutboundManagerMockRecorder) AddHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*OutboundManager)(nil).AddHandler), arg0, arg1)
}

// Close mocks base method
func (m *OutboundManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *OutboundManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*OutboundManager)(nil).Close))
}

// GetDefaultHandler mocks base method
func (m *OutboundManager) GetDefaultHandler() outbound.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultHandler")
	ret0, _ := ret[0].(outbound.Handler)
	return ret0
}

// GetDefaultHandler indicates an expected call of GetDefaultHandler
func (mr *OutboundManagerMockRecorder) GetDefaultHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultHandler", reflect.TypeOf((*OutboundManager)(nil).GetDefaultHandler))
}

// GetHandler mocks base method
func (m *OutboundManager) GetHandler(arg0 string) outbound.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandler", arg0)
	ret0, _ := ret[0].(outbound.Handler)
	return ret0
}

// GetHandler indicates an expected call of GetHandler
func (mr *OutboundManagerMockRecorder) GetHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandler", reflect.TypeOf((*OutboundManager)(nil).GetHandler), arg0)
}

// RemoveHandler mocks base method
func (m *OutboundManager) RemoveHandler(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveHandler", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveHandler indicates an expected call of RemoveHandler
func (mr *OutboundManagerMockRecorder) RemoveHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveHandler", reflect.TypeOf((*OutboundManager)(nil).RemoveHandler), arg0, arg1)
}

// Start mocks base method
func (m *OutboundManager) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *OutboundManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*OutboundManager)(nil).Start))
}

// Type mocks base method
func (m *OutboundManager) Type() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Type indicates an expected call of Type
func (mr *OutboundManagerMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*OutboundManager)(nil).Type))
}

// OutboundHandlerSelector is a mock of HandlerSelector interface
type OutboundHandlerSelector struct {
	ctrl     *gomock.Controller
	recorder *OutboundHandlerSelectorMockRecorder
}

// OutboundHandlerSelectorMockRecorder is the mock recorder for OutboundHandlerSelector
type OutboundHandlerSelectorMockRecorder struct {
	mock *OutboundHandlerSelector
}

// NewOutboundHandlerSelector creates a new mock instance
func NewOutboundHandlerSelector(ctrl *gomock.Controller) *OutboundHandlerSelector {
	mock := &OutboundHandlerSelector{ctrl: ctrl}
	mock.recorder = &OutboundHandlerSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *OutboundHandlerSelector) EXPECT() *OutboundHandlerSelectorMockRecorder {
	return m.recorder
}

// Select mocks base method
func (m *OutboundHandlerSelector) Select(arg0 []string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Select indicates an expected call of Select
func (mr *OutboundHandlerSelectorMockRecorder) Select(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*OutboundHandlerSelector)(nil).Select), arg0)
}
