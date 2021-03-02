package chess

type (
	Board   [0x78]Piece
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
