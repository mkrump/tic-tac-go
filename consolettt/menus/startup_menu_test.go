package menus

import (
	"bytes"
	"github.com/sc2nomore/tic-tac-go/consolettt"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanOrComputerPromptValid(t *testing.T) {
	input := bytes.NewBufferString("2\r\n")
	output := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(input, output)
	startupMenu := NewStartupMenu(*mockConsole)

	choice, _ := startupMenu.PlayerTypePrompt(1)

	assert.Equal(t, "COMPUTER", choice)
}

func TestHumanOrComputerPromptInvalid(t *testing.T) {
	input := bytes.NewBufferString("NOT VALID\r\n")
	output := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(input, output)
	startupMenu := NewStartupMenu(*mockConsole)

	_, err := startupMenu.PlayerTypePrompt(1)

	assert.NotNil(t, err)
}

func TestHumanOrComputerPromptInvalidChoice(t *testing.T) {
	input := bytes.NewBufferString("10\r\n")
	output := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(input, output)
	startupMenu := NewStartupMenu(*mockConsole)

	_, err := startupMenu.PlayerTypePrompt(1)

	assert.NotNil(t, err)
}

func TestPlayerSymbolPromptValid(t *testing.T) {
	input := bytes.NewBufferString("X\r\n")
	output := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(input, output)
	startupMenu := NewStartupMenu(*mockConsole)

	choice, _ := startupMenu.PlayerSymbolPrompt()

	assert.Equal(t, "X", choice)
}

func TestPlayerSymbolPromptInvalidValid(t *testing.T) {
	input := bytes.NewBufferString("1\r\n")
	output := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(input, output)
	startupMenu := NewStartupMenu(*mockConsole)

	_, err := startupMenu.PlayerSymbolPrompt()

	assert.NotNil(t, err)
}

func TestStartupMenuComputerPlayer(t *testing.T) {
	empty := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(empty, empty)
	startupMenu := NewStartupMenu(*mockConsole)

	playerSelection, _ := startupMenu.SelectPlayerType("COMPUTER", "X")

	expectedPlayerType := players.MakeComputerPlayer("X")
	assert.Equal(t, expectedPlayerType, playerSelection)
}

func TestStartupMenuConsolePlayer(t *testing.T) {
	empty := bytes.NewBufferString("")
	mockConsole := consolettt.NewTTTConsole(empty, empty)
	startupMenu := NewStartupMenu(*mockConsole)

	playerSelection, _ := startupMenu.SelectPlayerType("HUMAN", "O")

	expectedPlayerType := players.MakeConsolePlayer("O")
	assert.Equal(t, expectedPlayerType, playerSelection)
}

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