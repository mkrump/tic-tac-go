package startupmenu

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/tttuis"
)

type Runner struct {
	startUpMenu tttuis.StartupMenu
	players     []core.Player
}

func MakeRunner(startupMenu tttuis.StartupMenu) *Runner {
	return &Runner{
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

func (runner *Runner) playerSymbols() []string {
	playerSymbols := []string{}
	for _, player := range runner.players {
		playerSymbols = append(playerSymbols, player.Symbol())
	}
	return playerSymbols
}

func (runner *Runner) symbolsAlreadyTaken(symbol string) (string, error) {
	alreadyChosen := runner.playerSymbols()
	if contains(alreadyChosen, symbol) {
		fmt.Println(
			fmt.Sprintf("%s has already been chosen.\n"+
				"Choose a different symbol.", symbol))
		return "", tttuis.ErrInvalidOption
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

func (runner *Runner) Players() (player1 core.Player, player2 core.Player) {
	return runner.players[0], runner.players[1]
}

func (runner *Runner) Run(playerNumber int) {
	playerType := retryUntilValid(func() (string, error) {
		return runner.startUpMenu.PlayerTypePrompt(playerNumber)
	})
	playerSymbol := retryUntilValid(func() (string, error) {
		val, err := runner.startUpMenu.PlayerSymbolPrompt()
		if err == nil {
			return runner.symbolsAlreadyTaken(val)
		}
		return val, err
	})
	player, _ := runner.startUpMenu.SelectPlayerType(playerType, playerSymbol)
	runner.players = append(runner.players, player)
}

func (runner *Runner) Setup() {
	for len(runner.players) < 2 {
		players := len(runner.players)
		runner.startUpMenu.ClearMenu()
		runner.Run(players + 1)
	}
}
