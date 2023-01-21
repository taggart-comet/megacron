// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EmergencyBrakeInterface is an autogenerated mock type for the EmergencyBrakeInterface type
type EmergencyBrakeInterface struct {
	mock.Mock
}

// ShouldStop provides a mock function with given fields:
func (_m *EmergencyBrakeInterface) ShouldStop() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *EmergencyBrakeInterface) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *EmergencyBrakeInterface) Stop() {
	_m.Called()
}

// Wait provides a mock function with given fields:
func (_m *EmergencyBrakeInterface) Wait() {
	_m.Called()
}

type mockConstructorTestingTNewEmergencyBrakeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmergencyBrakeInterface creates a new instance of EmergencyBrakeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmergencyBrakeInterface(t mockConstructorTestingTNewEmergencyBrakeInterface) *EmergencyBrakeInterface {
	mock := &EmergencyBrakeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
