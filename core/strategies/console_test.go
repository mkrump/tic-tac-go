package strategies

import (
	"bytes"
	"github.com/sc2nomore/tic-tac-go/core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConsoleInputMac(t *testing.T) {
	consoleStrat := ConsoleStrategy{in: bytes.NewBufferString("1\n")}
	mockBoard := &mocks.Playable{}
	move := consoleStrat.FindMove(mockBoard, 1).(string)
	assert.Equal(t, "1", move)
}

func TestGetConsoleInvalidInput(t *testing.T) {
	consoleStrat := ConsoleStrategy{in: bytes.NewBufferString("A\r\n")}
	mockBoard := &mocks.Playable{}
	move := consoleStrat.FindMove(mockBoard, 1).(string)
	assert.Equal(t, "A", move)
}

func TestGetConsoleInputWindows(t *testing.T) {
	consoleStrat := ConsoleStrategy{in: bytes.NewBufferString("2\r\n")}
	mockBoard := &mocks.Playable{}
	move := consoleStrat.FindMove(mockBoard, 1).(string)
	assert.Equal(t, "2", move)
}
