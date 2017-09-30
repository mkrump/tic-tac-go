package menus

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"github.com/sc2nomore/tic-tac-go/core/players"
)

func TestHumanOrComputerPromptValid(t *testing.T) {
	input := bytes.NewBufferString("2\r\n")
	output := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{input, output}

	choice, _ := consoleMenu.PlayerTypePrompt()

	assert.Equal(t, 2, choice)
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

	playerSelection, _ := consoleMenu.SelectPlayerType(2, "X")

	expectedPlayerType := players.MakeComputerPlayer("X")
	assert.Equal(t, expectedPlayerType, playerSelection)
}

func TestStartupMenuConsolePlayer(t *testing.T) {
	empty := bytes.NewBufferString("")
	consoleMenu := ConsoleMenu{empty, empty}

	playerSelection, _ := consoleMenu.SelectPlayerType(1, "O")

	expectedPlayerType := players.MakeConsolePlayer("O")
	assert.Equal(t, expectedPlayerType, playerSelection)
}

//func TestStartupMenu(t *testing.T) {
//	mockStartupMenus := &mocks.StartupMenu{}
//	mockStartupMenus.On("PlayerTypePrompt", mock.Anything).Return(1)
//	mockStartupMenus.On("PlayerSymbolPrompt", mock.Anything).Return(1)
//	mockStartupMenus.On("SelectPlayerType", mock.Anything).Return(1)
//
//	expectedPlayerType := players.MakeComputerPlayer("X")
//	assert.Equal(t, expectedPlayerType, playerSelection)
//}

//
//func TestPlayerSymbolPromptDuplicateSymbol(t *testing.T) {
//	input := bytes.NewBufferString("O\n")
//
//	_, err := PlayerSymbolPrompt(input)
//
//	assert.NotNil(t, err)
//}
