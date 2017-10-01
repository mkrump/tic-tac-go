package menus

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"github.com/sc2nomore/tic-tac-go/consoles"
	"regexp"
	"strings"
	"fmt"
)

type StartupMenu struct {
	console consoles.Console
}

func NewStartupMenu(console consoles.Console) *StartupMenu {
	return &StartupMenu{
		console: console,
	}
}

func (startupMenu StartupMenu) SelectPlayerType(playerType string, playerSymbol string) (core.Player, error) {
	switch playerType {
	case "HUMAN":
		return players.MakeConsolePlayer(playerSymbol), nil
	case "COMPUTER":
		return players.MakeComputerPlayer(playerSymbol), nil
	default:
		return nil, consoles.ErrInvalidOption
	}
}

func (startupMenu StartupMenu) PlayerTypePrompt() (string, error) {
	startupMenu.console.RenderMessage(
		"Choose a player type: \n\n" +
			"  1. Human Player\n" +
			"  2. Computer Player\n\n")
	choice := startupMenu.console.ReadInput()
	switch choice {
	case "1":
		return "HUMAN", nil
	case "2":
		return "COMPUTER", nil
	default:
		return "", consoles.ErrInvalidOption
	}
}

func (startupMenu StartupMenu) PlayerSymbolPrompt() (string, error) {
	re := regexp.MustCompile("^[A-Z]$")

	startupMenu.console.RenderMessage(
		"Choose a marker [A-Z]: ")

	input := startupMenu.console.ReadInput()
	input = strings.ToUpper(input)
	switch {
	case re.MatchString(input):
		return input, nil
	default:
		startupMenu.console.RenderMessage(
			fmt.Sprintf("%s is invalid. Please choose marker A-Z.", input))
		return "", consoles.ErrInvalidOption
	}
}
