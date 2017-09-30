package menus

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/sc2nomore/mocks"
	"github.com/sc2nomore/tic-tac-go/uis"
)

func TestHumanOrComputerPromptValid(t *testing.T) {
	input := bytes.NewBufferString("2\r\n")
	output := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{input, output}

	choice, _ := consoleMenu.PlayerTypePrompt()

	assert.Equal(t, "2", choice)
}

func TestHumanOrComputerPromptInvalid(t *testing.T) {
	input := bytes.NewBufferString("NOT VALID\r\n")
	output := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{input, output}

	_, err := consoleMenu.PlayerTypePrompt()

	assert.NotNil(t, err)
}

func TestHumanOrComputerPromptInvalidChoice(t *testing.T) {
	input := bytes.NewBufferString("10\r\n")
	output := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{input, output}

	_, err := consoleMenu.PlayerTypePrompt()

	assert.NotNil(t, err)
}

func TestPlayerSymbolPromptValid(t *testing.T) {
	input := bytes.NewBufferString("X\r\n")
	output := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{input, output}

	choice, _ := consoleMenu.PlayerSymbolPrompt()

	assert.Equal(t, "X", choice)
}

func TestPlayerSymbolPromptInvalidValid(t *testing.T) {
	input := bytes.NewBufferString("1\r\n")
	output := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{input, output}

	_, err := consoleMenu.PlayerSymbolPrompt()

	assert.NotNil(t, err)
}

func TestStartupMenuComputerPlayer(t *testing.T) {
	empty := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{empty, empty}

	playerSelection, _ := consoleMenu.SelectPlayerType("2", "X")

	expectedPlayerType := players.MakeComputerPlayer("X")
	assert.Equal(t, expectedPlayerType, playerSelection)
}

func TestStartupMenuConsolePlayer(t *testing.T) {
	empty := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{empty, empty}

	playerSelection, _ := consoleMenu.SelectPlayerType("1", "O")

	expectedPlayerType := players.MakeConsolePlayer("O")
	assert.Equal(t, expectedPlayerType, playerSelection)
}

func TestStartupMenu(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt").Return("2", nil)
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil)
	expectedPlayerSelected := players.MakeComputerPlayer("X")
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeStartupMenuRunner(mockStartupMenus)

	startUpMenuRunner.Run()

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

func TestStartupMenuInvalidThenValid(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt").Return("", uis.ErrInvalidOption).Once()
	mockStartupMenus.On("PlayerTypePrompt").Return("2", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("X", nil)
	expectedPlayerSelected := players.MakeComputerPlayer("X")
	mockStartupMenus.On("SelectPlayerType", "2", "X").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeStartupMenuRunner(mockStartupMenus)

	startUpMenuRunner.Run()

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

func TestStartupMenuValidThenInvalid(t *testing.T) {
	mockStartupMenus := &mocks.StartupMenu{}
	mockStartupMenus.On("PlayerTypePrompt").Return("1", nil).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("", uis.ErrInvalidOption).Once()
	mockStartupMenus.On("PlayerSymbolPrompt").Return("O", nil).Once()
	expectedPlayerSelected := players.MakeConsolePlayer("O")
	mockStartupMenus.On("SelectPlayerType", "1", "O").Return(expectedPlayerSelected, nil)
	startUpMenuRunner := MakeStartupMenuRunner(mockStartupMenus)

	startUpMenuRunner.Run()

	assert.Equal(t, expectedPlayerSelected, startUpMenuRunner.players[0])
}

//
//func TestPlayerSymbolPromptDuplicateSymbol(t *testing.T) {
//	input := bytes.NewBufferString("O\n")
//
//	_, err := PlayerSymbolPrompt(input)
//
//	assert.NotNil(t, err)
//}
