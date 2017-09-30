package menus

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"github.com/sc2nomore/tic-tac-go/core/players"
)

func TestHumanOrComputerPromptValid(t *testing.T) {
	input := bytes.NewBufferString("2\r\n")

	choice, _ := PlayerTypePrompt(input)

	assert.Equal(t, 2, choice)
}

func TestHumanOrComputerPromptInvalid(t *testing.T) {
	input := bytes.NewBufferString("NOT VALID\r\n")

	_, err := PlayerTypePrompt(input)

	assert.NotNil(t, err)
}

func TestHumanOrComputerPromptInvalidChoice(t *testing.T) {
	input := bytes.NewBufferString("10\r\n")

	_, err := PlayerTypePrompt(input)

	assert.NotNil(t, err)
}

func TestPlayerSymbolPromptValid(t *testing.T) {
	input := bytes.NewBufferString("X\r\n")

	choice, _ := PlayerSymbolPrompt(input)

	assert.Equal(t, "X", choice)
}

func TestPlayerSymbolPromptInvalidValid(t *testing.T) {
	input := bytes.NewBufferString("1\r\n")

	_, err := PlayerSymbolPrompt(input)

	assert.NotNil(t, err)
}


func TestStartupMenuComputerPlayer(t *testing.T) {
	playerSelection, _ := SelectPlayerType(2, "X")

	expectedPlayerType := players.MakeComputerPlayer("X")

	assert.Equal(t, expectedPlayerType, playerSelection)
}

func TestStartupMenuConsolePlayer(t *testing.T) {
	playerSelection, _ := SelectPlayerType(1, "X")

	expectedPlayerType := players.MakeConsolePlayer("X")

	assert.Equal(t, expectedPlayerType, playerSelection)
}

//
//func TestPlayerSymbolPromptDuplicateSymbol(t *testing.T) {
//	input := bytes.NewBufferString("O\n")
//
//	_, err := PlayerSymbolPrompt(input)
//
//	assert.NotNil(t, err)
//}
