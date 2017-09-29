package uis

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
)

type UI interface {
	GetMove() error
	RenderBoard()
}

type ConsoleUI struct {
	Game          core.Game
	BoardRenderer BoardRender
}

func MakeConsoleUI(game core.Game, render BoardRender) UI {
	return ConsoleUI{
		Game:          game,
		BoardRenderer: render,
	}
}

//ConsolePrint prints string to console
func ConsolePrint(str string) {
	fmt.Println(str)
}

func (consoleUI ConsoleUI) GetMove() error {
	userMove := consoleUI.Game.GetMove()
	move, err := ValidateMove(userMove)
	if err != nil {
		fmt.Print(err)
		return err
	}
	if err := consoleUI.Game.MakeMove(move); err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func (consoleUI ConsoleUI) RenderBoard() {
	board := consoleUI.Game.GameBoard()
	players := consoleUI.Game.GamePlayers()
	fmt.Println("\n\n" + consoleUI.BoardRenderer.RenderBoard(board, players))
}
