package chess

var (
	BISHOP_RAYS = []Sq{-ROW - 1, -ROW + 1, +ROW - 1, +1}
	ROOK_RAYS   = []Sq{-1, 1, -ROW, +ROW}
	QUEEN_RAYS  = append(BISHOP_RAYS, ROOK_RAYS...)
	KNIGHT_RAYS = []Sq{
		2*ROW + 1, ROW + 2, -ROW + 2, -2*ROW + 1,
		2*ROW - 1, ROW - 2, -ROW - 2, -2*ROW - 1,
	}
)

func slide(board *Board, c Color, from Sq, towards Sq) (Sq, Piece, bool) {
	sq := from + towards
	capture, ok := board.At(from)
	return sq, capture, ok && c != capture.Color
}

func rays(board *Board, c Color, from Sq, rays []Sq) (sqs []Sq) {
	for _, towards := range rays {
		for {
			if sq, capture, ok := slide(board, c, from, towards); !ok {
				break
			} else {
				sqs = append(sqs, sq)
				if capture.Color == -c {
					break
				}
			}
		}
	}
	return
}

func (p Pawn) Moves(board *Board, color Color, from Sq) (sqs []Sq) {
	for i := 0; i < 2; i += int(color) {
		sq, capture, ok := slide(board, color, from, Sq(i+1)*ROW)
		if ok && capture.Color != -color {
			sqs = append(sqs, sq)
		}
	}
	for _, x := range []int{-1, 1} {
		if sq, _, ok := slide(board, color, from, Sq(x)+ROW*Sq(color)); ok {
			sqs = append(sqs, sq)
		}
	}
	return
}

func (k Knight) Moves(board *Board, color Color, from Sq) (sqs []Sq) {
	for _, ray := range KNIGHT_RAYS {
		if sq, _, ok := slide(board, color, from, ray); ok {
			sqs = append(sqs, sq)
		}
	}
	return
}

func (b Bishop) Moves(board *Board, color Color, from Sq) []Sq {
	return rays(board, color, from, BISHOP_RAYS)
}
func (r Rook) Moves(board *Board, color Color, from Sq) []Sq {
	return rays(board, color, from, ROOK_RAYS)
}
func (q Queen) Moves(board *Board, color Color, from Sq) []Sq {
	return rays(board, color, from, QUEEN_RAYS)
}

func (k King) Moves(board *Board, color Color, from Sq) (sqs []Sq) {
	for _, ray := range QUEEN_RAYS {
		if sq, _, ok := slide(board, color, from, ray); ok {
			sqs = append(sqs, sq)
		}
	}
	return
}
