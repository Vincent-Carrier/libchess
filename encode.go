package chess

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

func (sq Sq) String() string {
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}

func (m Move) String() string {
	return fmt.Sprintf("%v-%v", m.From, m.To)
}
func (m Capture) String() string {
	return fmt.Sprintf("%v-%v", m.From, m.To)
}
func (m EnPassant) String() string {
	return fmt.Sprintf("%v-%v", m.From, m.To)
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
		if p, _ := b.At(sq); p.Color != EMPTY {
			if skip > 0 {
				buf.WriteString(strconv.Itoa(skip))
				skip = 0
			}
			buf.WriteRune(p.char())
		} else {
			skip++
		}

		if sq.File() == 7 {
			buf.WriteRune('/')
			skip = 0
			sq -= ROW
			sq -= Sq(sq.File())
		} else {
			sq++
		}
	}
	return buf.String()
}
