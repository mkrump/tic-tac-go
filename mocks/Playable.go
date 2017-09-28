// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Board is an autogenerated mock type for the Board type
type Playable struct {
	mock.Mock
}

// BoardState provides a mock function with given fields:
func (_m *Playable) BoardState() []int {
	ret := _m.Called()

	var r0 []int
	if rf, ok := ret.Get(0).(func() []int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// GridSize provides a mock function with given fields:
func (_m *Playable) GridSize() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MakeMove provides a mock function with given fields: _a0, _a1
func (_m *Playable) MakeMove(_a0 int, _a1 int) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenSquares provides a mock function with given fields:
func (_m *Playable) OpenSquares() []int {
	ret := _m.Called()

	var r0 []int
	if rf, ok := ret.Get(0).(func() []int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// UndoMove provides a mock function with given fields: _a0
func (_m *Playable) UndoMove(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
