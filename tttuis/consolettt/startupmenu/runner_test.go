package startupmenu

import (
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartupMenuValid(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("2", nil)
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil)
	expectedPlayerSelected := players.MakeComputerPlayer("X")
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeRunner(mockStartupMenus)

	startUpMenuRunner.Run(1)

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

func TestStartupMenuInvalidThenValid(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("", consolettt.ErrInvalidOption).Once()
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("1", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil)
	expectedPlayerSelected := players.MakeConsolePlayer("X")
	mockStartupMenus.On("SelectPlayerType", "1", "X").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeRunner(mockStartupMenus)

	startUpMenuRunner.Run(1)

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

func TestStartupMenuValidThenInvalid(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("1", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("", consolettt.ErrInvalidOption).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("O", nil).Once()
	expectedPlayerSelected := players.MakeConsolePlayer("O")
	mockStartupMenus.On("SelectPlayerType", "1", "O").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeRunner(mockStartupMenus)

	startUpMenuRunner.Run(1)

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

func TestStartupMenuSymbolTaken(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	expectedPlayer := players.MakeComputerPlayer("X")

	mockStartupMenus.On("ClearMenu", ).Return()
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("2", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil).Once()
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayer, nil).Once()
	startUpMenuRunner := MakeRunner(mockStartupMenus)

	startUpMenuRunner.Run(1)
	_, err := startUpMenuRunner.symbolsAlreadyTaken("X")

	assert.NotNil(t, err)
}

func TestStartupMenuEntireFlow(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	expectedPlayer1 := players.MakeComputerPlayer("X")
	expectedPlayer2 := players.MakeConsolePlayer("O")

	mockStartupMenus.On("ClearMenu", ).Return()
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("2", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil).Once()
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayer1, nil).Once()

	mockStartupMenus.On("PlayerTypePrompt", 2).Return("2", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("O", nil).Once()
	mockStartupMenus.On("SelectPlayerType", "2", "O").Return(expectedPlayer2, nil).Once()
	startUpMenuRunner := MakeRunner(mockStartupMenus)

	startUpMenuRunner.Setup()
	player1, player2 := startUpMenuRunner.Players()

	assert.Equal(t, expectedPlayer1, player1)
	assert.Equal(t, expectedPlayer2, player2)
}
