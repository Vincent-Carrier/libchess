package chess


func (p Piece) rays(board *Board, from Sq, rays []Sq) (sqs []Sq) {
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

func (k King) Moves(board *Board, color Color, from Sq) []Sq {
	panic("implement me")
}

func (q Queen) Moves(board *Board, color Color, from Sq) []Sq {
	panic("implement me")
}

func (r Rook) Moves(board *Board, color Color, from Sq) []Sq {
	panic("implement me")
}

func (b Bishop) Moves(board *Board, color Color, from Sq) []Sq {
	panic("implement me")
}

func (k Knight) Moves(board *Board, color Color, from Sq) []Sq {
	panic("implement me")
}

func (p Pawn) Moves(board *Board, color Color, from Sq) []Sq {
	panic("implement me")
}
