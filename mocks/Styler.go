// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Styler is an autogenerated mock type for the Styler type
type Styler struct {
	mock.Mock
}

// Style provides a mock function with given fields: _a0, _a1, _a2
func (_m *Styler) Style(_a0 int, _a1 int, _a2 string) string {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
