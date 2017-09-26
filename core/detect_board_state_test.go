package core

import (
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test3InRow0IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		1, 1, 1,
		0, 0, 0,
		0, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InRow1IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		0, 0, 0,
		1, 1, 1,
		0, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InRow2IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		0, 0, 0,
		0, 0, 0,
		1, 1, 1,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InColumn0IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InColumn1IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InColumn2IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InDownDiagIsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test3InUpDiagIsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		0, 0, 1,
		0, 1, 0,
		1, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test4InUpDiagIsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(4)
	mockBoard.On("BoardState").Return([]int{
		0, 0, 0, 1,
		0, 0, 1, 0,
		0, 1, 0, 0,
		1, 0, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, 1), "")
}

func Test4InDownDiagIsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(4)
	mockBoard.On("BoardState").Return([]int{
		-1, 0, 0, 0,
		0, -1, 0, 0,
		0, 0, -1, 0,
		0, 0, 0, -1,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, -1), "")
}

func Test4InColumn0IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(4)
	mockBoard.On("BoardState").Return([]int{
		-1, 0, 0, 0,
		-1, 0, 0, 0,
		-1, 0, 0, 0,
		-1, 0, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, -1), "")
}

func Test4InRow1IsWin(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(4)
	mockBoard.On("BoardState").Return([]int{
		0, 0, 0, 0,
		-1, -1, -1, -1,
		0, 0, 0, 0,
		0, 0, 0, 0,
	})

	assert.False(t, IsTie(mockBoard), "")
	assert.True(t, IsWin(mockBoard, -1), "")
}

func Test3IsTie(t *testing.T) {
	mockBoard := &mocks.Playable{}
	mockBoard.On("GridSize").Return(3)
	mockBoard.On("BoardState").Return([]int{
		-1, 1, -1,
		1, -1, -1,
		1, -1, 1,
	})

	assert.True(t, IsTie(mockBoard), "")
}
