package menus

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/uis"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/menus"
)

type StartupMenuRunner struct {
	startUpMenu menus.StartupMenu
	players     []core.Player
}

func MakeStartupMenuRunner(startupMenu menus.StartupMenu) *StartupMenuRunner {
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
			fmt.Sprintf("%s has already been chosen.\n"+
				"Choose a different symbol.", symbol))
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

func (startupMenuRunner *StartupMenuRunner) Run(playerNumber int) {
	playerType := retryUntilValid(func() (string, error) {
		return startupMenuRunner.startUpMenu.PlayerTypePrompt(playerNumber)
	})
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
		players := len(startupMenuRunner.players)
		startupMenuRunner.startUpMenu.ClearMenu()
		startupMenuRunner.Run(players + 1)
	}
}
