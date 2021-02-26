package chess

type Board [0x78]Piece

type Game struct {
	*Board
	Active        Color
	Castles       [2][2]bool
	EnPassant     Square
	HalfMoveClock int
	FullMoves     int
}

func (b *Board) At(sq Square) (Piece, bool) {
	p := b[sq]
	return p, sq.Inbounds() && p != Piece{}
}
