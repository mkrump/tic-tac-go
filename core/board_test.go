package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveOpenSquare3x3(t *testing.T) {
	board := MakeBoard(3)
	board.MakeMove(0, 1)
	expected := make([]int, 3*3)
	expected[0] = 1
	assert.Equal(t, expected, board.BoardState, "")
}

func TestMoveOpenSquare4x4(t *testing.T) {
	board := MakeBoard(4)
	board.MakeMove(9, -1)
	expected := make([]int, 4*4)
	expected[9] = -1
	assert.Equal(t, expected, board.BoardState, "")
}

func TestMoveOccupiedSquare4x4(t *testing.T) {
	board := MakeBoard(4)
	board.MakeMove(9, -1)
	err := board.MakeMove(9, 1)
	expected := make([]int, 4*4)
	expected[9] = -1
	assert.NotNil(t, err)
	assert.Equal(t, expected, board.BoardState)
}

func TestMoveOutOfBounds(t *testing.T) {
	board := MakeBoard(3)
	err := board.MakeMove(100, 1)
	assert.NotNil(t, err)
}
