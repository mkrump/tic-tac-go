package uis

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
)

type UI interface {
	GetMove() error
}

type ConsoleUI struct {
	Game core.Game
}

func MakeConsoleUI(game core.Game) UI {
	return ConsoleUI{
		Game: game,
	}
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

