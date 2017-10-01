package uis

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/consolettt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/uis"
)

type ConsoleTTTUI struct {
	Game          core.Game
	BoardRenderer uis.BoardRender
	console       consolettt.Console
}

func NewConsoleTTTUI(game core.Game, render uis.BoardRender, console consolettt.Console) *ConsoleTTTUI {
	return &ConsoleTTTUI{
		Game:          game,
		BoardRenderer: render,
		console:       console,
	}
}

func (ui ConsoleTTTUI) RenderMessage(str string) {
	ui.console.RenderMessage(str)
}

func (ui ConsoleTTTUI) GetMove() error {
	userMove := ui.Game.GetMove()
	move, err := ValidateMove(userMove)
	if err != nil {
		ui.RenderMessage(fmt.Sprintf("%s is Invalid Input \n\n", userMove))
		return err
	}
	if err := ui.Game.MakeMove(move); err != nil {
		switch err {
		case core.ErrOutOfBounds:
			ui.RenderMessage(fmt.Sprintf("%s is Out of Bounds \n\n", userMove))
			return err
		case core.ErrSquareOccupied:
			ui.RenderMessage(fmt.Sprintf("%s is Occupied \n\n", userMove))
			return err
		default:
			ui.RenderMessage(fmt.Sprint(err))
			return err
		}
	}
	return nil
}

func (ui ConsoleTTTUI) RenderBoard() {
	board := ui.Game.GameBoard()
	players := ui.Game.GamePlayers()
	ui.console.ClearConsole()
	ui.RenderMessage(fmt.Sprintf("\n\n"+ui.BoardRenderer.RenderBoard(board, players)) + "\n")
}

func (ui ConsoleTTTUI) winMessage(marker string) string {
	return fmt.Sprintf("%8s's Win! \n\n", marker)
}

func (ui ConsoleTTTUI) tieMessage() string {
	return fmt.Sprintf("%8sTie... \n\n", "")
}

func (ui ConsoleTTTUI) stillPlayingMessage() string {
	return fmt.Sprintf("%8s's Turn \n\n", ui.Game.ActivePlayerMarker())
}

func (ui ConsoleTTTUI) NextGameState() (message string, endGame bool) {
	if ui.Game.IsWin() {
		message := ui.winMessage(ui.Game.InActivePlayerMarker())
		return message, false
	}
	if ui.Game.IsTie() {
		message := ui.tieMessage()
		return message, false
	}
	return ui.stillPlayingMessage(), true
}
