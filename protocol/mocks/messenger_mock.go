// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/messenger.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/theobitoproject/kankuro/protocol"
)

// MockMessenger is a mock of Messenger interface.
type MockMessenger struct {
	ctrl     *gomock.Controller
	recorder *MockMessengerMockRecorder
}

// MockMessengerMockRecorder is the mock recorder for MockMessenger.
type MockMessengerMockRecorder struct {
	mock *MockMessenger
}

// NewMockMessenger creates a new mock instance.
func NewMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	mock := &MockMessenger{ctrl: ctrl}
	mock.recorder = &MockMessengerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessenger) EXPECT() *MockMessengerMockRecorder {
	return m.recorder
}

// WriteLog mocks base method.
func (m *MockMessenger) WriteLog(level protocol.LogLevel, message string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteLog", level, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteLog indicates an expected call of WriteLog.
func (mr *MockMessengerMockRecorder) WriteLog(level, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteLog", reflect.TypeOf((*MockMessenger)(nil).WriteLog), level, message)
}

// WriteRecord mocks base method.
func (m *MockMessenger) WriteRecord(data interface{}, stream, namespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRecord", data, stream, namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRecord indicates an expected call of WriteRecord.
func (mr *MockMessengerMockRecorder) WriteRecord(data, stream, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRecord", reflect.TypeOf((*MockMessenger)(nil).WriteRecord), data, stream, namespace)
}

// WriteState mocks base method.
func (m *MockMessenger) WriteState(data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteState", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteState indicates an expected call of WriteState.
func (mr *MockMessengerMockRecorder) WriteState(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteState", reflect.TypeOf((*MockMessenger)(nil).WriteState), data)
}
