package chess

const (
	WHITE Color = 1
	EMPTY Color = 0
	BLACK Color = -1
)

type (
	Color int

	Piecer interface {
		Moves(g *Game, from Sq) *Moves
		//Move(board *Board, color Color, from, to Sq) bool
	}

	Piece struct {
		Color
		Kind Piecer
	}

	Pawn   struct{}
	Knight struct{}
	Bishop struct{}
	Rook   struct{}
	Queen  struct{}
	King   struct{}
)

func (p Piece) Moves(g *Game, from Sq) Moves {
	return *p.Kind.Moves(g, from)
}

func (color Color) pawnRow() Sq {
	switch color {
	case WHITE:
		return 1
	case BLACK:
		return 6
	default:
		panic("invalid color")
	}
}
func (color Color) kingRow() Sq {
	switch color {
	case WHITE:
		return 0
	case BLACK:
		return 7
	default:
		panic("invalid color")
	}
}
