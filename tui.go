package main

import "unicode"

func Char(m Mover) (r rune) {
	switch m.(type) {
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
	if m.Color() == White {
		r = unicode.ToUpper(r)
	}
	return
}
