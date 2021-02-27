package chess

const (
	White Color = 1
	Black Color = -1
)

type (
	Color int

	Mover interface {
		Moves(board *Board, color Color, from Sq) []Sq
		//Move(board *Board, color Color, from, to Sq) bool
	}

	Piece struct {
		Color
		Mover
	}

	Pawn   struct{}
	Knight struct{}
	Bishop struct{}
	Rook   struct{}
	Queen  struct{}
	King   struct{}
)


func (p Piece) Moves(board *Board, from Sq) []Sq {
	return p.Mover.Moves(board, p.Color, from)
}
