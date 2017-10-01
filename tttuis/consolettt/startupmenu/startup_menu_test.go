package startupmenu

import (
	"bytes"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
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


