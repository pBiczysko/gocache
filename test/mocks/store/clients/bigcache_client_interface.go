// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BigcacheClientInterface is an autogenerated mock type for the BigcacheClientInterface type
type BigcacheClientInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: key
func (_m *BigcacheClientInterface) Delete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *BigcacheClientInterface) Get(key string) ([]byte, error) {
	ret := _m.Called(key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, entry
func (_m *BigcacheClientInterface) Set(key string, entry []byte) error {
	ret := _m.Called(key, entry)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(key, entry)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
