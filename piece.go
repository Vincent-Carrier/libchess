package chess

const (
	WHITE Color = 1
	BLACK Color = -1
)

type (
	Color int

	Piecer interface {
		Moves(board *Board, color Color, from Sq) []Sq
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


func (p Piece) Moves(board *Board, from Sq) []Sq {
	return p.Piecer.Moves(board, p.Color, from)
}
