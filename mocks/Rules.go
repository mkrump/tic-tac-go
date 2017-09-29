// Code generated by mockery v1.0.0
package mocks

import core "github.com/sc2nomore/tic-tac-go/core"
import mock "github.com/stretchr/testify/mock"

// Rules is an autogenerated mock type for the Rules type
type Rules struct {
	mock.Mock
}

// ActivePlayerNumber provides a mock function with given fields: board
func (_m *Rules) ActivePlayerNumber(board core.Board) int {
	ret := _m.Called(board)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.Board) int); ok {
		r0 = rf(board)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InActivePlayerNumber provides a mock function with given fields: board
func (_m *Rules) InActivePlayerNumber(board core.Board) int {
	ret := _m.Called(board)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.Board) int); ok {
		r0 = rf(board)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IsTie provides a mock function with given fields: board
func (_m *Rules) IsTie(board core.Board) bool {
	ret := _m.Called(board)

	var r0 bool
	if rf, ok := ret.Get(0).(func(core.Board) bool); ok {
		r0 = rf(board)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsWin provides a mock function with given fields: board, player
func (_m *Rules) IsWin(board core.Board, player int) bool {
	ret := _m.Called(board, player)

	var r0 bool
	if rf, ok := ret.Get(0).(func(core.Board, int) bool); ok {
		r0 = rf(board, player)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}