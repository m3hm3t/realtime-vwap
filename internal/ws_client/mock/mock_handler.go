// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock_handler is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/m3hm3t/realtime-vwap/internal/ws_client/dto"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockHandler) Handle(receiver chan dto.Response) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", receiver)
}

// Handle indicates an expected call of Handle.
func (mr *MockHandlerMockRecorder) Handle(receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), receiver)
}
