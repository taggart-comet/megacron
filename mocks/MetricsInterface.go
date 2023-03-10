// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MetricsInterface is an autogenerated mock type for the MetricsInterface type
type MetricsInterface struct {
	mock.Mock
}

// IncFailedCronRun provides a mock function with given fields: command
func (_m *MetricsInterface) IncFailedCronRun(command string) {
	_m.Called(command)
}

// IncSuccessfulCronRun provides a mock function with given fields: command
func (_m *MetricsInterface) IncSuccessfulCronRun(command string) {
	_m.Called(command)
}

// MarkCronRunningTime provides a mock function with given fields: command, runningTime
func (_m *MetricsInterface) MarkCronRunningTime(command string, runningTime time.Duration) {
	_m.Called(command, runningTime)
}

// Push provides a mock function with given fields:
func (_m *MetricsInterface) Push() {
	_m.Called()
}

// SetMemory provides a mock function with given fields: containerName, valueInMegabytes
func (_m *MetricsInterface) SetMemory(containerName string, valueInMegabytes int) {
	_m.Called(containerName, valueInMegabytes)
}

// ToggleRunningCron provides a mock function with given fields: command, isUp
func (_m *MetricsInterface) ToggleRunningCron(command string, isUp bool) {
	_m.Called(command, isUp)
}

type mockConstructorTestingTNewMetricsInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMetricsInterface creates a new instance of MetricsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMetricsInterface(t mockConstructorTestingTNewMetricsInterface) *MetricsInterface {
	mock := &MetricsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
