package chess

import (
	"bytes"
	"fmt"
	"unicode"
)

func (sq Sq) String() string {
	if !sq.Inbounds() {
		panic("tried to print an out-of-bounds square")
	}
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}



func (p Piece) char() (r rune) {
	switch p.Kind.(type) {
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
		panic("invalid argument")
	}
	if p.Color == WHITE {
		r = unicode.ToUpper(r)
	}
	return
}

func (b *Board) String() string {
	var buf bytes.Buffer
	var skip int
	for sq := Sq(0x70); sq != 0x07; {
		if p, ok := b.At(sq); ok && p.Color != EMPTY {
			if skip > 0 {
				buf.WriteString(fmt.Sprint(skip))
				skip = 0
			}
			buf.WriteRune(p.char())
		} else {
			skip++
		}

		if sq.File() >= 7 {
			if skip > 0 {
				buf.WriteString(fmt.Sprint(skip))
				skip = 0
			}
			buf.WriteRune('/')
			sq -= ROW
			sq -= sq.File()
		} else {
			sq++
		}
	}
	buf.WriteRune(b[0x07].char())
	return buf.String()
}
