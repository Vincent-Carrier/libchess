package chess

import "fmt"

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
