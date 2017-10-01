package main

import (
	"os"
	"github.com/sc2nomore/tic-tac-go/core/games"
	"github.com/sc2nomore/tic-tac-go/core/tictactoe"
	//"github.com/sc2nomore/tic-tac-go/consolettt/menus"
	"github.com/sc2nomore/tic-tac-go/consolettt/uis"
	"github.com/sc2nomore/tic-tac-go/consolettt"
	//"time"
	players2 "github.com/sc2nomore/tic-tac-go/core/players"
	"time"
)

func Setup() *uis.ConsoleTTTUI {
	console := consolettt.NewTTTConsole(os.Stdin, os.Stdout)
	//startupMenu := menus.MakeStartupMenuRunner(menus.NewStartupMenu(console))
	//startupMenu.Setup()
	//player1, player2 := startupMenu.Players()
	//players := games.MakePlayers(player1, player2)
	player1 := players2.MakeComputerPlayer("O")
	player2 := players2.MakeConsolePlayer("X")
	players := games.MakePlayers(player1, player2)
	game := games.MakeGame(tictactoe.MakeTTTBoard(3), players, tictactoe.TTTRules{})
	styler := uis.ColorStyler{}
	boardRender := uis.MakeTTTBoardRender(styler)
	return uis.MakeConsoleUI(game, boardRender, console)
}

func main() {
	consoleUI := Setup()
	playing := true
	var message string
	consoleUI.RenderBoard()
	message, playing = consoleUI.NextGameState()
	consoleUI.RenderMessage(message)
	for playing {
		start := time.Now()
		err := consoleUI.GetMove()
		if err != nil {
			continue
		}
		t := time.Now()
		elapsed := t.Sub(start)
		if elapsed < time.Second*1 {
			time.Sleep(time.Second*1 - elapsed)
		}
		consoleUI.RenderBoard()
		message, playing = consoleUI.NextGameState()
		consoleUI.RenderMessage(message)
	}
}
