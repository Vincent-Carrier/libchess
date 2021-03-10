package chess

import (
	"fmt"
	"strings"
)

type (
	Board struct {
		squares [128]Piece
		pieces  [2][]Sq
	}
	Pieces []struct {
		Piece
		Sq
	}
	Castles [2][2]bool // [White, Black][Queen, King]

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
	// TODO: figure out why a single Sscan doesn't work
	parts := strings.Split(fen, " ")
	_, err = fmt.Sscan(parts[0], &g.Board)
	if len(parts) == 1 {
		g.Active = WHITE
		return
	}
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

func (b *Board) At(sq Sq) (p Piece, ok bool) {
	if !sq.Inbounds() {
		return
	}
	return b.squares[sq], true
}

func (b *Board) Set(sq Sq, p Piece) {
	if !sq.Inbounds() {
		panic("invalid square")
	}
	b.squares[sq] = p
}

func (g *Game) MovesFrom(sq Sq) (moves Moves) {
	p, ok := g.At(sq)
	if !ok {
		panic("invalid square")
	}
	moves = p.Moves(g, sq)
	return
}

func (g *Game) MovesOf(color Color) (moves Moves) {
	for i := 0; i <= 7; i++ {
		for j := 0; i <= 7; j++ {
			if p, _ := g.At(Coords(x, y)); p.Color == color {
				moves = append(moves, p.Moves())
			}
		}
	}
	return
}

func (g *Game) Exec(m Mover) (o Outcome) {
	m.Execute(g)
	g.Active *= -1
	g.HalfMoveClock += 1
	if g.Active == WHITE {
		g.FullMoves += 1
	}
	return
}
