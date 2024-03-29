package engine

import (
	. "github.com/Vincent-Carrier/libchess"
)

func pieceValue(p Piecer) (f float32) {
	switch p.(type) {
	case Pawn:
		f = 1.
	case Knight:
		f = 3.
	case Bishop:
		f = 3.
	case Rook:
		f = 5.
	case Queen:
		f = 9.
	case King:
		f = 100.
	}
	return
}

func Eval(g *Game) (score float32) {
	for _, pp := range g.PiecesOf(BLACK + WHITE) {
		score += pieceValue(pp.Piece)
	}

	return score
}
