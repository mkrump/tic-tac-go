package console

import (
	"errors"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"bytes"
)

func TestRequestUserMoveValid(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("9")
	mockGame.On("MakeMove", 8).Return(nil)
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	consoleUI.GetMove()

	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUserMoveInvalid(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("ABC")
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUserMoveOutBounds(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("1000")
	mockGame.On("MakeMove", mock.AnythingOfType("int")).Return(core.ErrOutOfBounds)
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestSquareOccupiedBounds(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("1")
	mockGame.On("MakeMove", mock.AnythingOfType("int")).Return(core.ErrSquareOccupied)
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	err := consoleUI.GetMove()

	assert.NotNil(t, err)
	mockGame.AssertCalled(t, "GetMove")
}

func TestRequestUnexpectedError(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("GetMove").Return("1")
	mockGame.On("MakeMove", mock.AnythingOfType("int")).Return(errors.New("SOME ERROR"))
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

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
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	consoleUI.RenderBoard()

	mockRender.AssertCalled(t, "RenderBoard", mockBoard, mockPlayerMapper)
}

func TestRenderNextGameStateWin(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("IsWin").Return(true)
	mockGame.On("IsTie").Return(false)
	mockGame.On("InActivePlayerMarker").Return("X")
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	message, playing := consoleUI.NextGameState()

	assert.Equal(t, consoleUI.winMessage("X"), message)
	assert.False(t, playing, "")
}

func TestRenderNextGameStateTie(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("IsWin").Return(false)
	mockGame.On("IsTie").Return(true)
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	message, playing := consoleUI.NextGameState()

	assert.Equal(t, consoleUI.tieMessage(), message)
	assert.False(t, playing, "")
}

func TestRenderNextGameStateNoWinorTie(t *testing.T) {
	mockGame := &mocks.Game{}
	mockGame.On("IsWin").Return(false)
	mockGame.On("IsTie").Return(false)
	mockRender := &mocks.BoardRender{}
	outBuffer := bytes.NewBufferString("")
	consoleUI := MakeConsoleUI(mockGame, mockRender, outBuffer)

	message, playing := consoleUI.NextGameState()

	assert.Equal(t, "", message)
	assert.True(t, playing, "")
}
