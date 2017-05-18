// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/anexia-it/go-cloudlog (interfaces: Event)

package cloudlog

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Event interface
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *_MockEventRecorder
}

// Recorder for MockEvent (not exported)
type _MockEventRecorder struct {
	mock *MockEvent
}

func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &_MockEventRecorder{mock}
	return mock
}

func (_m *MockEvent) EXPECT() *_MockEventRecorder {
	return _m.recorder
}

func (_m *MockEvent) Encode() map[string]interface{} {
	ret := _m.ctrl.Call(_m, "Encode")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

func (_mr *_MockEventRecorder) Encode() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Encode")
}
