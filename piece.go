package chess

const (
	WHITE Color = 1
	EMPTY Color = 0
	BLACK Color = -1
)

type (
	Color int

	Piecer interface {
		Moves(g *Game, from Sq) Moves
		//Move(board *Board, color Color, from, to Sq) bool
	}

	Piece struct {
		Color
		Piecer
	}

	Pawn   struct{}
	Knight struct{}
	Bishop struct{}
	Rook   struct{}
	Queen  struct{}
	King   struct{}
)


func (p Piece) Moves(g *Game, from Sq) Moves {
	return p.Piecer.Moves(g, from)
}

func (c Color) pawnRow() int16 {
	switch c {
	case WHITE:
		return 1
	case BLACK:
		return 6
	default:
		panic("invalid color")
	}
}
