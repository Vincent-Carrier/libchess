package chess

type (
	Mover interface {
		Validate(board *Board) bool
		Execute(board *Board)
	}

	BasicMove struct {
		From, To Sq
		Piece
	}
	CaptureMove struct {
		BasicMove
		Capture Piece
	}
)
