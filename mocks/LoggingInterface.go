// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	runner "megacron/system/runner"

	mock "github.com/stretchr/testify/mock"
)

// LoggingInterface is an autogenerated mock type for the LoggingInterface type
type LoggingInterface struct {
	mock.Mock
}

// LogCommandsFromChannelForever provides a mock function with given fields: feedBackChannel
func (_m *LoggingInterface) LogCommandsFromChannelForever(feedBackChannel <-chan runner.RunResult) {
	_m.Called(feedBackChannel)
}

// LogCommandsFromClosedChannel provides a mock function with given fields: feedBackChannel
func (_m *LoggingInterface) LogCommandsFromClosedChannel(feedBackChannel <-chan runner.RunResult) {
	_m.Called(feedBackChannel)
}

// PushMetrics provides a mock function with given fields:
func (_m *LoggingInterface) PushMetrics() {
	_m.Called()
}

// StartPushingMetrics provides a mock function with given fields:
func (_m *LoggingInterface) StartPushingMetrics() {
	_m.Called()
}

// TrackMemory provides a mock function with given fields: container
func (_m *LoggingInterface) TrackMemory(container string) {
	_m.Called(container)
}

// WatchCrons provides a mock function with given fields:
func (_m *LoggingInterface) WatchCrons() {
	_m.Called()
}

type mockConstructorTestingTNewLoggingInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoggingInterface creates a new instance of LoggingInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoggingInterface(t mockConstructorTestingTNewLoggingInterface) *LoggingInterface {
	mock := &LoggingInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
