package menus

import (
	"regexp"
	"github.com/sc2nomore/tic-tac-go/uis"
	"strings"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/players"
	"fmt"
)

type StartupMenu interface {
	PlayerTypePrompt() (string, error)
	PlayerSymbolPrompt() (string, error)
	SelectPlayerType(string, string) (core.Player, error)
}

func (consoleMenu ConsoleMenu) PlayerSymbolPrompt() (string, error) {
	re := regexp.MustCompile("^[A-Z]$")

	consoleMenu.RenderMessage(
		"Choose a marker [A-Z]: ")

	input := consoleMenu.ReadInput()
	input = strings.ToUpper(input)
	switch {
	case re.MatchString(input):
		return input, nil
	default:
		consoleMenu.RenderMessage(
			fmt.Sprintf("%s is invalid. Please choose marker A-Z.", input))
		return "", uis.ErrInvalidOption
	}
}

func (consoleMenu ConsoleMenu) PlayerTypePrompt() (string, error) {
	consoleMenu.RenderMessage(
		"Choose a player type: \n\n" +
			"  1. Human Player\n" +
			"  2. Computer Player\n\n")
	choice := consoleMenu.ReadInput()
	switch choice {
	case "1":
		return "HUMAN", nil
	case "2":
		return "COMPUTER", nil
	default:
		return "", uis.ErrInvalidOption
	}
}

func (consoleMenu ConsoleMenu) SelectPlayerType(playerType string, playerSymbol string) (core.Player, error) {
	switch playerType {
	case "HUMAN":
		return players.MakeConsolePlayer(playerSymbol), nil
	case "COMPUTER":
		return players.MakeComputerPlayer(playerSymbol), nil
	default:
		return nil, uis.ErrInvalidOption
	}
}

type StartupMenuRunner struct {
	startUpMenu StartupMenu
	players     []core.Player
}

func MakeStartupMenuRunner(startupMenu StartupMenu) *StartupMenuRunner {
	return &StartupMenuRunner{
		startUpMenu: startupMenu,
		players:     make([]core.Player, 0, 2),
	}
}

func contains(symbols []string, symbol string) bool {
	for _, e := range symbols {
		if symbol == e {
			return true
		}
	}
	return false
}

func (startupMenuRunner *StartupMenuRunner) playerSymbols() []string {
	playerSymbols := []string{}
	for _, player := range startupMenuRunner.players {
		playerSymbols = append(playerSymbols, player.Symbol())
	}
	return playerSymbols
}

func (startupMenuRunner *StartupMenuRunner) symbolsAlreadyTaken(symbol string) (string, error) {
	alreadyChosen := startupMenuRunner.playerSymbols()
	if contains(alreadyChosen, symbol) {
		fmt.Println(
			fmt.Sprintf("%s has already been chosen. "+
				"Please choose a different symbol", symbol))
		return "", uis.ErrInvalidOption
	}
	return symbol, nil
}

func retryUntilValid(fn func() (string, error)) string {
	for {
		if val, err := fn(); err == nil {
			return val
		}
	}
}

func (startupMenuRunner *StartupMenuRunner) Players() (player1 core.Player, player2 core.Player) {
	return startupMenuRunner.players[0], startupMenuRunner.players[1]
}

func (startupMenuRunner *StartupMenuRunner) Run() {
	playerType := retryUntilValid(startupMenuRunner.startUpMenu.PlayerTypePrompt)
	playerSymbol := retryUntilValid(func() (string, error) {
		val, err := startupMenuRunner.startUpMenu.PlayerSymbolPrompt()
		if err == nil {
			return startupMenuRunner.symbolsAlreadyTaken(val)
		}
		return val, err
	})
	player, _ := startupMenuRunner.startUpMenu.SelectPlayerType(playerType, playerSymbol)
	startupMenuRunner.players = append(startupMenuRunner.players, player)
}

func (startupMenuRunner *StartupMenuRunner) Setup() {
	for len(startupMenuRunner.players) < 2 {
		startupMenuRunner.Run()
	}
}
