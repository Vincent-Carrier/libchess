package main

func PieceValue(p Mover) (f float32) {
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
