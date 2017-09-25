package ui

import (
	"github.com/sc2nomore/tic-tac-go/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func RenderBoard(board core.Board) string {
	consoleBoard := boardToConsoleUI(board)
	return render(consoleBoard, board.Size)
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

func render(board []string, gridSize int) string {
	renderedBoard := "\n"
	for rowIndex := 0; rowIndex < gridSize; rowIndex++ {
		rowStartIndex := rowIndex * gridSize
		row := board[rowStartIndex:(rowStartIndex + gridSize)]
		renderedBoard = addRowDivider(rowIndex, renderedBoard, gridSize)
		renderedBoard += " " + strings.Join(row, " | ") + " \n"
	}
	return renderedBoard
}
func addRowDivider(i int, s string, gridSize int) string {
	if i > 0 {
		s += strings.Repeat(" -  ", gridSize-1) + " - \n"
	}
	return s
}

func TestRenderingEmpty3x3Board(t *testing.T) {
	expected := "\n" +
		"   |   |   \n" +
		" -   -   - \n" +
		"   |   |   \n" +
		" -   -   - \n" +
		"   |   |   \n"
	board := core.MakeBoard(3)
	assert.Equal(t, expected, RenderBoard(board), "")
}

func TestRenderingNonEmpty3x3Board(t *testing.T) {
	expected := "\n" +
		"   | X |   \n" +
		" -   -   - \n" +
		"   | O |   \n" +
		" -   -   - \n" +
		"   |   | X \n"
	board := core.MakeBoard(3)
	board.MakeMove(1, 1)
	board.MakeMove(4, -1)
	board.MakeMove(8, 1)
	assert.Equal(t, expected, RenderBoard(board), "")
}

func TestRenderingNonEmpty4x4Board(t *testing.T) {
	expected := "\n" +
		"   | X |   |   \n" +
		" -   -   -   - \n" +
		"   | O |   |   \n" +
		" -   -   -   - \n" +
		"   |   | X |   \n" +
		" -   -   -   - \n" +
		"   |   |   |   \n"
	board := core.MakeBoard(4)
	board.MakeMove(1, 1)
	board.MakeMove(5, -1)
	board.MakeMove(10, 1)
	assert.Equal(t, expected, RenderBoard(board), "")
}
