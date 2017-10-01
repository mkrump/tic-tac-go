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

func (startupMenu StartupMenu) ClearMenu() {
	startupMenu.console.ClearConsole()
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

func (startupMenu StartupMenu) PlayerTypePrompt(playerNumber int) (string, error) {
	startupMenu.console.RenderMessage(
		fmt.Sprintf(
			"Choose a player type for player %d. \n\n"+
				"  1. Human Player\n"+
				"  2. Computer Player\n\n" +
					">: ", playerNumber))
	choice := startupMenu.console.ReadInput()
	switch choice {
	case "1":
		return "HUMAN", nil
	case "2":
		return "COMPUTER", nil
	default:
		startupMenu.console.RenderMessage(
			fmt.Sprintf(
				"%s is not valid.\n\n", choice))
		return "", consoles.ErrInvalidOption
	}
}

func (startupMenu StartupMenu) PlayerSymbolPrompt() (string, error) {
	re := regexp.MustCompile("^[A-Z]$")

	startupMenu.console.RenderMessage(
		"\nChoose a marker [A-Z].\n\n" +
			">: ")

	input := startupMenu.console.ReadInput()
	input = strings.ToUpper(input)
	switch {
	case re.MatchString(input):
		return input, nil
	default:
		startupMenu.console.RenderMessage(
			fmt.Sprintf(
				"%s is not valid.\n", input))
		return "", consoles.ErrInvalidOption
	}
}
