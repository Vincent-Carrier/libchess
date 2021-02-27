package chess

type Board [0x77]Piece

type Game struct {
	*Board
	Active        Color
	Castles       [2][2]bool
	EnPassant     Sq
	HalfMoveClock int
	FullMoves     int
}

func (b *Board) At(sq Sq) (Piece, bool) {
	p := b[sq]
	return p, sq.Inbounds()
}
