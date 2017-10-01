package main

import (
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt"
	"os"
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	"github.com/sc2nomore/tic-tac-go/tttuis/consolettt/ui"
	"time"
	"fmt"
	players2 "github.com/sc2nomore/tic-tac-go/core/players"
)

func setup() *ui.ConsoleTTTUI {
	console := consolettt.NewTTTConsole(os.Stdin, os.Stdout)
	//startupMenu := startupmenu.MakeStartupMenuRunner(startupmenu.NewStartupMenu(console))
	//startupMenu.Setup()
	//player1, player2 := startupMenu.Players()
	//TODO remove only for quick setup to test game play
	player2 := players2.MakeComputerPlayer("O")
	player1 := players2.MakeConsolePlayer("X")
	players := games.MakePlayers(player1, player2)
	game := games.MakeGame(tictactoe.MakeTTTBoard(4), players, tictactoe.TTTRules{})
	styler := ui.ColorStyler{}
	boardRender := ui.MakeTTTBoardRender(styler)
	return ui.NewConsoleTTTUI(game, boardRender, console)
}

func pacer(fn func() error, minElapsedTime time.Duration) error {
	start := time.Now()
	err := fn()
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(start)
	fmt.Printf("%v", elapsedTime)
	switch {
	case err != nil:
		return err
	case elapsedTime < minElapsedTime:
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
	count := 0
	for playing {
		err := pacer(consoleUI.GetMove, 1*time.Second)
		count++
		if count == 2 {
			break
		}
		if err != nil {
			continue
		}
		consoleUI.RenderBoard()
		message, playing = consoleUI.NextGameState()
		consoleUI.RenderMessage(message)
	}
}
