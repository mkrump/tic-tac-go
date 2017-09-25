package game

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"strings"
)

//RenderBoard returns a string representation of a Board object
func RenderBoard(board core.Board) string {
	consoleBoard := boardToConsoleUI(board)
	return render(consoleBoard, board.Size)
}

func ConsolePrint(str string) {
	print(str)
}

func render(uiBoard []string, gridSize int) string {
	renderedBoard := "\n"
	for rowNumber := 0; rowNumber < gridSize; rowNumber++ {
		rowSlice := getRowSlice(rowNumber, gridSize, uiBoard)
		renderedBoard += renderRowDivider(rowNumber, gridSize)
		renderedBoard += renderRow(rowSlice)
	}
	return renderedBoard
}

func boardToConsoleUI(board core.Board) []string {
	consoleBoard := make([]string, len(board.BoardState))
	for i, square := range board.BoardState {
		switch square {
		case 1:
			consoleBoard[i] = "X"
		case -1:
			consoleBoard[i] = "O"
		default:
			consoleBoard[i] = " "
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
	return "  " + strings.Join(row, "  |  ") + "  \n"
}

func renderRowDivider(rowNumber int, gridSize int) string {
	if rowNumber > 0 {
		return strings.Repeat("- - - ", gridSize-1) + "- - -\n"
	}
	return ""
}
