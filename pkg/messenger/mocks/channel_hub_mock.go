// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/messenger/channel_hub.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	messenger "github.com/theobitoproject/kankuro/pkg/messenger"
)

// MockChannelHub is a mock of ChannelHub interface.
type MockChannelHub struct {
	ctrl     *gomock.Controller
	recorder *MockChannelHubMockRecorder
}

// MockChannelHubMockRecorder is the mock recorder for MockChannelHub.
type MockChannelHubMockRecorder struct {
	mock *MockChannelHub
}

// NewMockChannelHub creates a new mock instance.
func NewMockChannelHub(ctrl *gomock.Controller) *MockChannelHub {
	mock := &MockChannelHub{ctrl: ctrl}
	mock.recorder = &MockChannelHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelHub) EXPECT() *MockChannelHubMockRecorder {
	return m.recorder
}

// GetErrorChannel mocks base method.
func (m *MockChannelHub) GetErrorChannel() messenger.ErrorChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErrorChannel")
	ret0, _ := ret[0].(messenger.ErrorChannel)
	return ret0
}

// GetErrorChannel indicates an expected call of GetErrorChannel.
func (mr *MockChannelHubMockRecorder) GetErrorChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorChannel", reflect.TypeOf((*MockChannelHub)(nil).GetErrorChannel))
}

// GetRecordChannel mocks base method.
func (m *MockChannelHub) GetRecordChannel() messenger.RecordChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordChannel")
	ret0, _ := ret[0].(messenger.RecordChannel)
	return ret0
}

// GetRecordChannel indicates an expected call of GetRecordChannel.
func (mr *MockChannelHubMockRecorder) GetRecordChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordChannel", reflect.TypeOf((*MockChannelHub)(nil).GetRecordChannel))
}
