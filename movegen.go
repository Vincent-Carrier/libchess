package chess

var (
	BISHOP_RAYS = []Sq{-ROW - 1, -ROW + 1, +ROW - 1, +ROW+1}
	ROOK_RAYS   = []Sq{-1, 1, -ROW, +ROW}
	QUEEN_RAYS  = append(BISHOP_RAYS, ROOK_RAYS...)
	KNIGHT_RAYS = []Sq{
		2*ROW + 1, ROW + 2, -ROW + 2, -2*ROW + 1,
		2*ROW - 1, ROW - 2, -ROW - 2, -2*ROW - 1,
	}
)

func slide(g *Game, from, towards Sq) (Sq, Piece, bool) {
	sq := from + towards
	capture, ok := g.At(sq)
	return sq, capture, ok && g.Active != capture.Color
}

func (moves *Moves) append(move Mover) {
	*moves = append(*moves, move)
}

func (moves *Moves) push(from, to Sq, capture Piece, color Color) bool {
	slide := Slide{from, to}
	var move Mover = &slide
	if capture.Color == -color {
		move = &Capture{slide, capture}
		moves.append(move)
		return false
	}
	moves.append(move)
	return true
}

func ray(g *Game, from, towards Sq, moves *Moves) {
	if to, capture, ok := slide(g, from, towards); !ok {
		return
	} else {
		moves.push(from, to, capture, g.Active)
		if capture.Color != -g.Active {
			ray(g, to, towards, moves)
		}
	}
}

func rays(g *Game, from Sq, rays []Sq) *Moves {
	moves := new(Moves)
	for _, towards := range rays {
		ray(g, from, towards, moves)
	}
	return moves
}


func (p Pawn) Moves(g *Game, from Sq) *Moves {
	moves := new(Moves)
	fwd := Sq(g.Active)
	to, capture, ok := slide(g, from, fwd*ROW)
	if ok && capture.Color != -g.Active {
		moves.append(&Slide{from, to})
	}
	if from.Rank() == g.Active.pawnRow() {
		to, capture, ok := slide(g, from, 2*fwd*ROW)
		if ok && capture.Color != -g.Active {
			moves.append(&Slide{from, to})
		}
	}
	for _, x := range []int{-1, 1} {
		if to, capture, ok := slide(g, from, Sq(x)+fwd*ROW); ok && capture.Color == -g.Active {
			moves.push(from, to, capture, g.Active)
		}
	}
	return moves
}

func (k Knight) Moves(g *Game, from Sq) (moves *Moves) {
	moves = new(Moves)
	for _, ray := range KNIGHT_RAYS {
		if to, capture, ok := slide(g, from, ray); ok {
			moves.push(from, to, capture, g.Active)
		}
	}
	return
}

func (b Bishop) Moves(g *Game, from Sq) *Moves {
	return rays(g, from, BISHOP_RAYS)
}

func (r Rook) Moves(g *Game, from Sq) *Moves {
	return rays(g, from, ROOK_RAYS)
}

func (q Queen) Moves(g *Game, from Sq) *Moves {
	return rays(g, from, QUEEN_RAYS)
}

func (k King) Moves(g *Game, from Sq) (moves *Moves) {
	moves = new(Moves)
	for _, ray := range QUEEN_RAYS {
		if to, capture, ok := slide(g, from, ray); ok {
			moves.push(from, to, capture, g.Active)
		}
	}

	return
}
