package console

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/uis"
)

type ConsoleUI struct {
	Game          core.Game
	BoardRenderer uis.BoardRender
}

func MakeConsoleUI(game core.Game, render uis.BoardRender) ConsoleUI {
	return ConsoleUI{
		Game:          game,
		BoardRenderer: render,
	}
}

//RenderMessage prints string to console
func (ConsoleUI ConsoleUI) RenderMessage(str string) {
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
	consoleUI.clearConsole()
	fmt.Println("\n\n" + consoleUI.BoardRenderer.RenderBoard(board, players))
}

func (consoleUI ConsoleUI) winMessage(marker string) string {
	return fmt.Sprintf("%s's win!", marker)
}

func (consoleUI ConsoleUI) tieMessage() string {
	return fmt.Sprintf("Tie...")
}

func (consoleUI ConsoleUI) clearConsole() {
	fmt.Print("\033c")
}

func (consoleUI ConsoleUI) NextGameState() (message string, endGame bool) {
	if consoleUI.Game.IsWin() {
		message := consoleUI.winMessage(consoleUI.Game.InActivePlayerMarker())
		return message, false
	}
	if consoleUI.Game.IsTie() {
		message := consoleUI.tieMessage()
		return message, false
	}
	return "", true
}
