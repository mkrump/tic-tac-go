package tictactoe

import "github.com/sc2nomore/tic-tac-go/core"

type TTTRules struct {
}

//IsWin returns boolean evaluation if a player has won for the supplied Board
func (ttt_rules TTTRules) IsWin(board core.Board, player int) bool {
	gridSize := board.GridSize()
	boardState := board.BoardState()
	switch {
	case upDiagWin(gridSize, boardState, player):
		return true
	case downDiagWin(gridSize, boardState, player):
		return true
	case rowWin(gridSize, boardState, player):
		return true
	case colWin(gridSize, boardState, player):
		return true
	default:
		return false
	}
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

//TODO factor out some of this
func rowWin(gridSize int, boardState []int, player int) bool {
	for i := 0; i < gridSize; i++ {
		var total int
		for j := 0; j < gridSize; j++ {
			if boardState[i*gridSize+j] == player {
				total++
			}
			if total == gridSize {
				return true
			}
		}
	}
	return false
}

func colWin(gridSize int, boardState []int, player int) bool {
	for i := 0; i < gridSize; i++ {
		var total int
		for j := 0; j < gridSize; j++ {
			if boardState[i+j*gridSize] == player {
				total++
			}
			if total == gridSize {
				return true
			}
		}
	}
	return false
}

func downDiagWin(gridSize int, boardState []int, player int) bool {
	var total int
	for i := 0; i < gridSize; i++ {
		//y = a + mx w/ a=0 m=gridsize+1
		if boardState[i*(gridSize+1)] == player {
			total++
		}
		if total == gridSize {
			return true
		}
	}
	return false
}

func upDiagWin(gridSize int, boardState []int, player int) bool {
	var total int
	for i := 0; i < gridSize; i++ {
		//y = a + mx w/ a=gridsize*(gridsize-1) m=-(gridsize-1)
		if boardState[gridSize*(gridSize-1)-(gridSize-1)*i] == player {
			total++
		}
		if total == gridSize {
			return true
		}
	}
	return false
}

func isOdd(n int) bool {
	return n%2 == 0
}

func isEven(n int) bool {
	return n%2 != 0
}
