package chess

import (
	"fmt"
	"strings"
)

type (
	Board   [128]Piece
	Castles [2][2]bool

	Game struct {
		Board
		Active Color
		Castles
		EnPassant     Sq
		HalfMoveClock int
		FullMoves     int
	}
)

func NewGame(fen string) (g *Game, err error) {
	g = new(Game)
	//TODO: figure out why a single Sscan doesn't work
	parts := strings.Split(fen, " ")
	_, err = fmt.Sscan(parts[0], &g.Board)
	_, err = fmt.Sscan(parts[1], &g.Active)
	_, err = fmt.Sscan(parts[2], &g.Castles)
	_, err = fmt.Sscan(parts[3], &g.EnPassant)
	_, err = fmt.Sscan(parts[4], &g.HalfMoveClock)
	_, err = fmt.Sscan(parts[5], &g.FullMoves)

	return
}

func StartingPosition() *Game {
	game, err := NewGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
	return game
}

func (b *Board) At(sq Sq) (Piece, bool) {
	p := b[sq]
	return p, sq.Inbounds()
}

func (g *Game) move(s string) error {
	var from, to Sq
	_, err := fmt.Sscanf(s, "%v-%v", &from, &to)
	p, _ := g.At(from)
	g.Board[to] = p
	g.Board[from] = Piece{}
	return err
}
