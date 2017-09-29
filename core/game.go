package core

import (
	"github.com/sc2nomore/tic-tac-go/core/boards"
	"github.com/sc2nomore/tic-tac-go/core/playertypes"
	"github.com/sc2nomore/tic-tac-go/core/rules"
)

type Game interface {
	IsWin() bool
	IsTie() bool
	GetMove() interface{}
	MakeMove(int) error
}

type TTTGame struct {
	Players PlayerMap
	Board   boards.Board
	Rules   rules.Rules
}

func (game TTTGame) GetMove() interface{} {
	activePlayer := game.ActivePlayer()
	return activePlayer.Move(game.Board, game.boardActivePlayer())
}

//MakeGame constructor for game struct
func MakeGame(board boards.Board, players PlayerMap, rules rules.Rules) TTTGame {
	return TTTGame{
		Board:   board,
		Players: players,
		Rules:   rules,
	}
}

// IsWin checks if the game is a win for the current player
// and returns a boolean
func (game TTTGame) IsWin() bool {
	return game.Rules.IsWin(game.Board, -1*game.boardActivePlayer())
}

// IsWin checks if the game is a tie and returns a boolean
func (game TTTGame) IsTie() bool {
	v := game.Rules.IsTie(game.Board)
	return v
}

// MakeMove attempts make a move and updates the TTTBoard state
// if valid and returns an error if the move is invalid
func (game TTTGame) MakeMove(move int) error {
	return game.Board.MakeMove(move, game.boardActivePlayer())
}

func (game TTTGame) InActivePlayer() playertypes.Player {
	return game.Players.Player(-1 * game.boardActivePlayer())
}

func (game TTTGame) ActivePlayer() playertypes.Player {
	return game.Players.Player(game.boardActivePlayer())
}

func (game TTTGame) InActivePlayerMarker() string {
	player := game.InActivePlayer()
	return player.Symbol()
}

func (game TTTGame) ActivePlayerMarker() string {
	player := game.ActivePlayer()
	return player.Symbol()
}

func MakePlayers(player1 playertypes.Player, player2 playertypes.Player) PlayerMap {
	return PlayerMap{
		Players: map[int]playertypes.Player{
			-1: player1,
			1:  player2,
		},
	}
}

func (game TTTGame) boardActivePlayer() int {
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
