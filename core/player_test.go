package core

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConsoleInputMac(t *testing.T) {
	consolePlayer := ConsolePlayer{in: bytes.NewBufferString("1\n")}
	move := consolePlayer.Move().(string)
	assert.Equal(t, "1", move)
}

func TestGetConsoleInvalidInput(t *testing.T) {
	consolePlayer := ConsolePlayer{in: bytes.NewBufferString("A\r\n")}
	move := consolePlayer.Move().(string)
	assert.Equal(t, "A", move)
}

func TestGetConsoleInputWindows(t *testing.T) {
	consolePlayer := ConsolePlayer{in: bytes.NewBufferString("2\r\n")}
	move := consolePlayer.Move().(string)
	assert.Equal(t, "2", move)
}
