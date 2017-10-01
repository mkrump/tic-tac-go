package menus

import (
	"testing"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/stretchr/testify/assert"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/sc2nomore/tic-tac-go/consolettt"
)

func TestStartupMenu(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("2", nil)
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil)
	expectedPlayerSelected := players.MakeComputerPlayer("X")
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeStartupMenuRunner(mockStartupMenus)

	startUpMenuRunner.Run(1)

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

func TestStartupMenuInvalidThenValid(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("", consolettt.ErrInvalidOption).Once()
	mockStartupMenus.On("PlayerTypePrompt", 1).Return("2", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil)
	expectedPlayerSelected := players.MakeComputerPlayer("X")
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeStartupMenuRunner(mockStartupMenus)

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
	startUpMenuRunner := MakeStartupMenuRunner(mockStartupMenus)

	startUpMenuRunner.Run(1)

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}
