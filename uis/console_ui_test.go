package uis

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"github.com/sc2nomore/tic-tac-go/mocks"
)

func TestRequestUserMoveValid(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("9")
	mockGame.On("MakeMove", 8).Return(nil)
	mockRender := &mocks.BoardRender{}

	consoleUI := MakeConsoleUI(mockGame, mockRender)

	consoleUI.GetMove()

	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUserMoveInvalid(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("ABC")
	mockRender := &mocks.BoardRender{}

	consoleUI := MakeConsoleUI(mockGame, mockRender)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUserMoveOutBounds(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("1000")
	mockGame.On("MakeMove", mock.AnythingOfType("int")).Return(errors.New("SOME ERROR"))
	mockRender := &mocks.BoardRender{}

	consoleUI := MakeConsoleUI(mockGame, mockRender)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}

func TestRenderBoard(t *testing.T) {
	mockGame := &mocks.Game{}
	mockBoard := &mocks.Board{}
	mockPlayerMapper := &mocks.PlayerMapper{}
	mockGame.On("GameBoard").Return(mockBoard)
	mockGame.On("GamePlayers").Return(mockPlayerMapper)
	mockRender := &mocks.BoardRender{}
	mockRender.On("RenderBoard", mockBoard, mockPlayerMapper).Return("BOARD")

	consoleUI := MakeConsoleUI(mockGame, mockRender)

	consoleUI.RenderBoard()

	mockRender.AssertCalled(t, "RenderBoard", mockBoard, mockPlayerMapper)
}
