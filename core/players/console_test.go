package players

import (
	"bytes"
	"github.com/sc2nomore/tic-tac-go/mocks"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConsoleInputMac(t *testing.T) {
	mockConsole := consolettt.NewTTTConsole(bytes.NewBufferString("1\n"), bytes.NewBufferString(""))
	consoleStrat := ConsoleStrategy{mockConsole}
	mockBoard := &mocks.Board{}
	move := consoleStrat.FindMove(mockBoard, 1).(string)
	assert.Equal(t, "1", move)
}

func TestGetConsoleInvalidInput(t *testing.T) {
	mockConsole := consolettt.NewTTTConsole(bytes.NewBufferString("A\r\n"), bytes.NewBufferString(""))
	consoleStrat := ConsoleStrategy{mockConsole}
	mockBoard := &mocks.Board{}
	move := consoleStrat.FindMove(mockBoard, 1).(string)
	assert.Equal(t, "A", move)
}

func TestGetConsoleInputWindows(t *testing.T) {
	mockConsole := consolettt.NewTTTConsole(bytes.NewBufferString("2\r\n"), bytes.NewBufferString(""))
	consoleStrat := ConsoleStrategy{mockConsole}
	mockBoard := &mocks.Board{}
	move := consoleStrat.FindMove(mockBoard, 1).(string)
	assert.Equal(t, "2", move)
}
