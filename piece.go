package chess

const (
	White Color = 1
	Black Color = -1
)

type (
	Color int

	Mover interface {
		Moves(board *Board, color Color, from Square) []Square
		//Move(board *Board, color Color, from, to Square) bool
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

func (p Piece) Moves(board *Board, from Square) []Square {
	return p.Mover.Moves(board, p.Color, from)
}

func (p Piece) rays(board *Board, from Square, rays []Square) (sqs []Square) {
	for _, towards := range rays {
		start := from
		for {
			start += towards
			if capture, ok := board.At(from); ok {
				switch capture.Color {
				case p.Color:
					sqs = append(sqs, from)
					continue
				case ^p.Color:
					continue
				case 0:
					sqs = append(sqs, from)
				}
			} else {
				break // outside bounds
			}
		}
	}
	return
}
