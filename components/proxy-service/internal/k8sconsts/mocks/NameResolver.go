// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// NameResolver is an autogenerated mock type for the NameResolver type
type NameResolver struct {
	mock.Mock
}

// ExtractServiceId provides a mock function with given fields: host
func (_m *NameResolver) ExtractServiceId(host string) string {
	ret := _m.Called(host)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(host)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetResourceName provides a mock function with given fields: id
func (_m *NameResolver) GetResourceName(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
