package engine

import (
	"unicode"

	. "github.com/Vincent-Carrier/libchess"
)

func Char(p Piece) (r rune) {
	switch p.Piecer.(type) {
	case Pawn:
		r = 'p'
	case Knight:
		r = 'n'
	case Bishop:
		r = 'b'
	case Rook:
		r = 'r'
	case Queen:
		r = 'q'
	case King:
		r = 'k'
	default:
		panic("Invalid argument")
	}
	if p.Color == WHITE {
		r = unicode.ToUpper(r)
	}
	return
}
