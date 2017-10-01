package main

import (
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt/startupmenu"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt/ui"
	"os"
	"time"
)

func setup() *ui.ConsoleTTTUI {
	console := consolettt.NewTTTConsole(os.Stdin, os.Stdout)
	startupMenu := startupmenu.MakeRunner(startupmenu.NewStartupMenu(console))
	startupMenu.Setup()
	player1, player2 := startupMenu.Players()
	players := games.MakePlayers(player1, player2)
	game := games.MakeGame(tictactoe.MakeTTTBoard(3), players, tictactoe.TTTRules{})
	styler := ui.ColorStyler{}
	boardRender := ui.MakeTTTBoardRender(styler)
	return ui.NewConsoleTTTUI(game, boardRender, console)
}

func pacer(fn func() error, minElapsedTime time.Duration) error {
	start := time.Now()
	err := fn()
	if err != nil {
		return err
	}
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(start)
	if elapsedTime < minElapsedTime {
		var duration = (minElapsedTime) - elapsedTime
		time.Sleep(duration)
	}
	return err
}

func main() {
	consoleUI := setup()
	var message string
	consoleUI.RenderBoard()
	message, playing := consoleUI.NextGameState()
	consoleUI.RenderMessage(message)
	for playing {
		err := pacer(consoleUI.GetMove, 1*time.Second)
		if err != nil {
			continue
		}
		consoleUI.RenderBoard()
		message, playing = consoleUI.NextGameState()
		consoleUI.RenderMessage(message)
	}
}
