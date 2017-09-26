package core

//IsWin returns boolean evaluation if a player has won for the supplied Playable
func IsWin(playable Playable, player int) bool {
	gridSize := playable.GridSize()
	boardState := playable.BoardState()
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
func IsTie(playable Playable) bool {
	return len(openSquares(playable)) == 0 &&
		!IsWin(playable, 1) && !IsWin(playable, -1)
}

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

func openSquares(playable Playable) []int {
	var openSquaresIndices []int
	for i, square := range playable.BoardState() {
		if square == 0 {
			openSquaresIndices = append(openSquaresIndices, i)
		}
	}
	return openSquaresIndices
}
