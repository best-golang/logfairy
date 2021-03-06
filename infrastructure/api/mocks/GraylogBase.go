// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// GraylogBase is an autogenerated mock type for the GraylogBase type
type GraylogBase struct {
	mock.Mock
}

// GetAuth provides a mock function with given fields: tokenName
func (_m *GraylogBase) GetAuth(tokenName string) ([]string, error) {
	ret := _m.Called(tokenName)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(tokenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleFailure provides a mock function with given fields: response, status
func (_m *GraylogBase) HandleFailure(response []byte, status int) error {
	ret := _m.Called(response, status)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, int) error); ok {
		r0 = rf(response, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
