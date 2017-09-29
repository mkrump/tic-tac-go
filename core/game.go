package core

type TTTGame struct {
	Players PlayerMap
	Board   Board
	Rules   Rules
}

func (tttGame TTTGame) GameBoard() Board {
	return tttGame.Board
}

func (tttGame TTTGame) GamePlayers() PlayerMap {
	return tttGame.Players
}

func (tttGame TTTGame) GetMove() interface{} {
	activePlayer := tttGame.ActivePlayer()
	return activePlayer.Move(tttGame.Board, tttGame.boardActivePlayer())
}

//MakeGame constructor for game struct
func MakeGame(board Board, players PlayerMap, rules Rules) TTTGame {
	return TTTGame{
		Board:   board,
		Players: players,
		Rules:   rules,
	}
}

// IsWin checks if the game is a win for the current player
// and returns a boolean
func (tttGame TTTGame) IsWin() bool {
	return tttGame.Rules.IsWin(tttGame.Board, -1*tttGame.boardActivePlayer())
}

// IsWin checks if the game is a tie and returns a boolean
func (tttGame TTTGame) IsTie() bool {
	v := tttGame.Rules.IsTie(tttGame.Board)
	return v
}

// MakeMove attempts make a move and updates the TTTBoard state
// if valid and returns an error if the move is invalid
func (tttGame TTTGame) MakeMove(move int) error {
	return tttGame.Board.MakeMove(move, tttGame.boardActivePlayer())
}

func (tttGame TTTGame) InActivePlayer() Player {
	return tttGame.Players.Player(-1 * tttGame.boardActivePlayer())
}

func (tttGame TTTGame) ActivePlayer() Player {
	return tttGame.Players.Player(tttGame.boardActivePlayer())
}

func (tttGame TTTGame) InActivePlayerMarker() string {
	player := tttGame.InActivePlayer()
	return player.Symbol()
}

func (tttGame TTTGame) ActivePlayerMarker() string {
	player := tttGame.ActivePlayer()
	return player.Symbol()
}

func MakePlayers(player1 Player, player2 Player) PlayerMap {
	return PlayerMap{
		Players: map[int]Player{
			-1: player1,
			1:  player2,
		},
	}
}

func (tttGame TTTGame) boardActivePlayer() int {
	var currentPlayer int
	gridSize := tttGame.Board.GridSize()
	openSquaresCount := len(tttGame.Board.OpenSquares())
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

func isOdd(n int) bool {
	return n%2 == 0
}

func isEven(n int) bool {
	return n%2 != 0
}
