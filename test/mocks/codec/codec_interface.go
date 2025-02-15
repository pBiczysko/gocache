// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import codec "github.com/eko/gocache/codec"
import mock "github.com/stretchr/testify/mock"
import store "github.com/eko/gocache/store"

// CodecInterface is an autogenerated mock type for the CodecInterface type
type CodecInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: key
func (_m *CodecInterface) Delete(key interface{}) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *CodecInterface) Get(key interface{}) (interface{}, error) {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields:
func (_m *CodecInterface) GetStats() *codec.Stats {
	ret := _m.Called()

	var r0 *codec.Stats
	if rf, ok := ret.Get(0).(func() *codec.Stats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*codec.Stats)
		}
	}

	return r0
}

// GetStore provides a mock function with given fields:
func (_m *CodecInterface) GetStore() store.StoreInterface {
	ret := _m.Called()

	var r0 store.StoreInterface
	if rf, ok := ret.Get(0).(func() store.StoreInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreInterface)
		}
	}

	return r0
}

// Invalidate provides a mock function with given fields: options
func (_m *CodecInterface) Invalidate(options store.InvalidateOptions) error {
	ret := _m.Called(options)

	var r0 error
	if rf, ok := ret.Get(0).(func(store.InvalidateOptions) error); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: key, value, options
func (_m *CodecInterface) Set(key interface{}, value interface{}, options *store.Options) error {
	ret := _m.Called(key, value, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, *store.Options) error); ok {
		r0 = rf(key, value, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
