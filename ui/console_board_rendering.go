package ui

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"strings"
)

//RenderBoard returns a string representation of a boards object
func RenderBoard(board boards.Board, players core.PlayerMap, styler Styler) string {
	consoleBoard := boardToConsoleUI(board, players, styler)
	return render(consoleBoard, board.GridSize())
}

//ConsolePrint prints string to console
func ConsolePrint(str string) {
	fmt.Println(str)
}

func render(uiBoard []string, gridSize int) string {
	renderedBoard := ""
	for rowNumber := 0; rowNumber < gridSize; rowNumber++ {
		rowSlice := getRowSlice(rowNumber, gridSize, uiBoard)
		renderedBoard += renderRowDivider(rowNumber, gridSize)
		renderedBoard += renderRow(rowSlice)
	}
	return renderedBoard
}

func boardToConsoleUI(board boards.Board, players core.PlayerMap, styler Styler) []string {
	consoleBoard := make([]string, len(board.BoardState()))
	for i, square := range board.BoardState() {
		marker, _ := players.PlayerSymbol(square)
		consoleBoard[i] = styler.Style(square, i, marker)
	}
	return consoleBoard
}

func getRowSlice(rowIndex int, gridSize int, board []string) []string {
	rowStartIndex := rowIndex * gridSize
	row := board[rowStartIndex:(rowStartIndex + gridSize)]
	return row
}

func renderRow(row []string) string {
	return "  " + strings.Join(row, "   |  ") + "   \n"
}

func renderRowDivider(rowNumber int, gridSize int) string {
	if rowNumber > 0 {
		return strings.Repeat("- - - - ", gridSize-1) + "- - - -\n"
	}
	return ""
}
