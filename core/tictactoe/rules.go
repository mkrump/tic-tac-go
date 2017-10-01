package tictactoe

import "github.com/sc2nomore/tic-tac-go/core"

type TTTRules struct {
}

//IsWin returns boolean evaluation if a player has won for the supplied Board
func (ttt_rules TTTRules) IsWin(board core.Board, player int) bool {
	gridSize := board.GridSize()
	boardState := board.BoardState()
	if isWin(gridSize, boardState, player) {
		return true
	}
	return false
}

//IsTie returns boolean evaluation is a playable is a tie
func (ttt_rules TTTRules) IsTie(board core.Board) bool {
	return len(board.OpenSquares()) == 0 &&
		!ttt_rules.IsWin(board, 1) && !ttt_rules.IsWin(board, -1)
}

func (ttt_rules TTTRules) ActivePlayerNumber(board core.Board) int {
	var currentPlayer int
	gridSize := board.GridSize()
	openSquaresCount := len(board.OpenSquares())
	switch {
	case isEven(gridSize) && isEven(openSquaresCount):
		currentPlayer = -1
	case isEven(gridSize) && isOdd(openSquaresCount):
		currentPlayer = 1
	case isOdd(gridSize) && isOdd(openSquaresCount):
		currentPlayer = -1
	case isOdd(gridSize) && isEven(openSquaresCount):
		currentPlayer = 1
	}
	return currentPlayer
}

func (ttt_rules TTTRules) InActivePlayerNumber(board core.Board) int {
	return -ttt_rules.ActivePlayerNumber(board)
}

func isWin(gridSize int, boardState []int, player int) bool {
	var nInDownDiag int
	var nInUpdDiag int
	for row := 0; row < gridSize; row++ {
		var nInRow int
		var nInCol int
		nInDownDiag += downDiagOccupied(row, player, boardState, gridSize)
		nInUpdDiag += upDiagOccupied(row, player, boardState, gridSize)
		if nEqualsGridSize(nInUpdDiag, gridSize) || nEqualsGridSize(nInDownDiag, gridSize) {
			return true
		}
		for col := 0; col < gridSize; col++ {
			nInRow += rowColOccupied(row, col, player, boardState, gridSize)
			nInCol += colRowOccupied(col, row, player, boardState, gridSize)
			if nEqualsGridSize(nInRow, gridSize) || nEqualsGridSize(nInCol, gridSize) {
				return true
			}
		}
	}
	return false
}
func nEqualsGridSize(count int, gridSize int) bool {
	if count == gridSize {
		return true
	}
	return false
}

func colRowOccupied(col int, row int, player int, boardState []int, gridSize int) int {
	if boardState[col*gridSize+row] == player {
		return 1
	}
	return 0
}

func rowColOccupied(row int, col int, player int, boardState []int, gridSize int) int {
	if boardState[row*gridSize+col] == player {
		return 1
	}
	return 0
}
func downDiagOccupied(row int, player int, boardState []int, gridSize int) int {
	//y = a + mx w/ a=0 m=gridsize+1
	if boardState[row*(gridSize+1)] == player {
		return 1
	}
	return 0
}

func upDiagOccupied(row int, player int, boardState []int, gridSize int) int {
	//y = a + mx w/ a=gridsize*(gridsize-1) m=-(gridsize-1)
	if boardState[gridSize*(gridSize-1)-(gridSize-1)*row] == player {
		return 1
	}
	return 0
}

func isOdd(n int) bool {
	return n%2 == 0
}

func isEven(n int) bool {
	return n%2 != 0
}
