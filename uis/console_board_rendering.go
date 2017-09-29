package uis

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"strings"
)

type BoardRender interface {
	RenderBoard(board boards.Board, players core.PlayerMap) string
}

type TTTBoardRender struct {
	styler Styler
}

func MakeTTTBoardRender(style Styler) TTTBoardRender {
	return TTTBoardRender{
		styler: style,
	}
}

//RenderBoard returns a string representation of a boards object
func (tttBoardRender TTTBoardRender) RenderBoard(board boards.Board, players core.PlayerMap) string {
	styledBoard := tttBoardRender.StyleSquares(board.BoardState(), players)
	return tttBoardRender.render(styledBoard, board.GridSize())
}

func (tttBoardRender TTTBoardRender) render(styledBoard []string, gridSize int) string {
	renderedBoard := ""
	for rowNumber := 0; rowNumber < gridSize; rowNumber++ {
		rowSlice := tttBoardRender.getRowSlice(rowNumber, gridSize, styledBoard)
		renderedBoard += tttBoardRender.renderRowDivider(rowNumber, gridSize)
		renderedBoard += tttBoardRender.renderRow(rowSlice)
	}
	return renderedBoard
}

func (tttBoardRender TTTBoardRender) StyleSquares(board []int, players core.PlayerMap) []string {
	consoleBoard := make([]string, len(board))
	for i, square := range board {
		marker, _ := players.PlayerSymbol(square)
		consoleBoard[i] = tttBoardRender.styler.Style(square, i, marker)
	}
	return consoleBoard
}

func (tttBoardRender TTTBoardRender) getRowSlice(rowIndex int, gridSize int, board []string) []string {
	rowStartIndex := rowIndex * gridSize
	row := board[rowStartIndex:(rowStartIndex + gridSize)]
	return row
}

func (tttBoardRender TTTBoardRender) renderRow(row []string) string {
	return "  " + strings.Join(row, "   |  ") + "   \n"
}

func (tttBoardRender TTTBoardRender) renderRowDivider(rowNumber int, gridSize int) string {
	if rowNumber > 0 {
		return strings.Repeat("- - - - ", gridSize-1) + "- - - -\n"
	}
	return ""
}
