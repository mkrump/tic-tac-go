package games

import "github.com/sc2nomore/tic-tac-go/core"

type TTTGame struct {
	Players core.PlayerMapper
	Board   core.Board
	Rules   core.Rules
}

func (game TTTGame) GameBoard() core.Board {
	return game.Board
}

func (game TTTGame) GamePlayers() core.PlayerMapper {
	return game.Players
}

func (game TTTGame) GetMove() interface{} {
	board := game.GameBoard()
	activePlayer := game.ActivePlayer()
	activePlayerNumber := game.Rules.ActivePlayerNumber(board)
	return activePlayer.Move(game.Board, activePlayerNumber)
}

//MakeGame constructor for game struct
func MakeGame(board core.Board, players core.PlayerMapper, rules core.Rules) TTTGame {
	return TTTGame{
		Board:   board,
		Players: players,
		Rules:   rules,
	}
}

// IsWin checks if the game is a win for the current player
// and returns a boolean
func (game TTTGame) IsWin() bool {
	board := game.GameBoard()
	InActivePlayerNumber := game.Rules.InActivePlayerNumber(board)
	return game.Rules.IsWin(game.Board, InActivePlayerNumber)
}

// IsWin checks if the game is a tie and returns a boolean
func (game TTTGame) IsTie() bool {
	return game.Rules.IsTie(game.Board)
}

// MakeMove attempts make a move and updates the TTTBoard state
// if valid and returns an error if the move is invalid
func (game TTTGame) MakeMove(move int) error {
	board := game.GameBoard()
	activePlayerNumber := game.Rules.ActivePlayerNumber(board)
	return game.Board.MakeMove(move, activePlayerNumber)
}

func (game TTTGame) InActivePlayer() core.Player {
	board := game.GameBoard()
	InActivePlayerNumber := game.Rules.InActivePlayerNumber(board)
	return game.Players.Player(InActivePlayerNumber)
}

func (game TTTGame) ActivePlayer() core.Player {
	board := game.GameBoard()
	activePlayerNumber := game.Rules.ActivePlayerNumber(board)
	return game.Players.Player(activePlayerNumber)
}

func (game TTTGame) InActivePlayerMarker() string {
	player := game.InActivePlayer()
	return player.Symbol()
}

func (game TTTGame) ActivePlayerMarker() string {
	player := game.ActivePlayer()
	return player.Symbol()
}
