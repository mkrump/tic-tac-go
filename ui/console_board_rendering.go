package ui

import (
	"fmt"
	"github.com/sc2nomore/tic-tac-go/core"
	"strings"
)

//RenderBoard returns a string representation of a board object
func RenderBoard(board core.Playable) string {
	consoleBoard := boardToConsoleUI(board)
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

func boardToConsoleUI(board core.Playable) []string {
	consoleBoard := make([]string, len(board.BoardState()))
	for i, square := range board.BoardState() {
		switch square {
		case -1:
			//TODO add color maybe using callback?
			//consoleBoard[i] = fmt.Sprintf("\033[1;31m%s\033[0m", "X")
			consoleBoard[i] = fmt.Sprintf("%2s", "X")
		case 1:
			//TODO add color maybe using callback?
			//consoleBoard[i] = fmt.Sprintf("\033[1;32m%s\033[0m", "O")
			consoleBoard[i] = fmt.Sprintf("%2s", "O")
		default:
			consoleBoard[i] = fmt.Sprintf("%2d", i+1)
		}
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
