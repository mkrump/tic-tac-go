package games

import "github.com/sc2nomore/tic-tac-go/core"

type Game struct {
	Players core.PlayerMap
	Board   core.Board
	Rules   core.Rules
}

func (game Game) GameBoard() core.Board {
	return game.Board
}

func (game Game) GamePlayers() core.PlayerMap {
	return game.Players
}

func (game Game) GetMove() interface{} {
	activePlayer := game.ActivePlayer()
	return activePlayer.Move(game.Board, game.boardActivePlayer())
}

//MakeGame constructor for game struct
func MakeGame(board core.Board, players core.PlayerMap, rules core.Rules) Game {
	return Game{
		Board:   board,
		Players: players,
		Rules:   rules,
	}
}

// IsWin checks if the game is a win for the current player
// and returns a boolean
func (game Game) IsWin() bool {
	return game.Rules.IsWin(game.Board, -1*game.boardActivePlayer())
}

// IsWin checks if the game is a tie and returns a boolean
func (game Game) IsTie() bool {
	v := game.Rules.IsTie(game.Board)
	return v
}

// MakeMove attempts make a move and updates the TTTBoard state
// if valid and returns an error if the move is invalid
func (game Game) MakeMove(move int) error {
	return game.Board.MakeMove(move, game.boardActivePlayer())
}

func (game Game) InActivePlayer() core.Player {
	return game.Players.Player(-1 * game.boardActivePlayer())
}

func (game Game) ActivePlayer() core.Player {
	return game.Players.Player(game.boardActivePlayer())
}

func (game Game) InActivePlayerMarker() string {
	player := game.InActivePlayer()
	return player.Symbol()
}

func (game Game) ActivePlayerMarker() string {
	player := game.ActivePlayer()
	return player.Symbol()
}

func MakePlayers(player1 core.Player, player2 core.Player) core.PlayerMap {
	return core.PlayerMap{
		Players: map[int]core.Player{
			-1: player1,
			1:  player2,
		},
	}
}

func (game Game) boardActivePlayer() int {
	var currentPlayer int
	gridSize := game.Board.GridSize()
	openSquaresCount := len(game.Board.OpenSquares())
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
