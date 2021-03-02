package chess

import (
	"fmt"
	"unicode"
)

func (sq *Sq) Scan(state fmt.ScanState, _ rune) error {
	var file, rank int16
	r, _, err := state.ReadRune()
	if err != nil {
		return err
	} else if r == '-' {
		*sq = NIL_SQ
		return nil
	} else if 'a' > r || r > 'h' {
		return fmt.Errorf("invalid character for square's file: %c", r)
	}
	file = int16(r - 'a')

	r, _, err = state.ReadRune()
	if err != nil {
		return err
	} else if '1' > r || r > '8' {
		return fmt.Errorf("invalid character for square's rank: %c", r)
	}
	rank = int16(r-'1') << 4

	*sq = Sq(rank | file)
	return nil
}

func (color *Color) Scan(state fmt.ScanState, _ rune) error {
	r, _, err := state.ReadRune()
	if err != nil {
		return err
	}
	switch r {
	case 'w':
		*color = WHITE
	case 'b':
		*color = BLACK
	case '-':
		*color = EMPTY
	default:
		return fmt.Errorf("invalid color rune %c", r)
	}
	return nil
}

func (p *Piece) Scan(state fmt.ScanState, verb rune) error {
	r, _, err := state.ReadRune()
	if err != nil {
		return err
	}
	switch unicode.ToUpper(r) {
	case 'P':
		p.Kind = Pawn{}
	case 'N':
		p.Kind = Knight{}
	case 'B':
		p.Kind = Bishop{}
	case 'R':
		p.Kind = Rook{}
	case 'Q':
		p.Kind = Queen{}
	case 'K':
		p.Kind = King{}
	default:
		return fmt.Errorf("invalid char %c", r)
	}
	if unicode.IsUpper(verb) {
		return nil
	}
	if unicode.IsLower(r) {
		p.Color = BLACK
	} else {
		p.Color = WHITE
	}
	return nil
}

func (b *Board) Scan(state fmt.ScanState, _ rune) (err error) {
	var p Piece
	for sq := Sq(0x70); sq != 0x08; {
		r, _, err := state.ReadRune()
		if '1' <= r && r <= '8' {
			sq += Sq(r - '0')
		} else if r == '/' {
			sq -= ROW
			sq -= sq & 0x08
		} else if _, err = fmt.Sscan(string(r), &p); err == nil {
			b[sq] = p
			sq++
		} else {
			return fmt.Errorf("invalid char %c", r)
		}
	}
	return
}

func (c *Castles) Scan(state fmt.ScanState, _ rune) error {
	for {
		r, _, err := state.ReadRune()
		if err != nil {
			return err
		}
		switch r {
		case 'Q':
			c[0][0] = true
		case 'K':
			c[0][1] = true
		case 'q':
			c[1][0] = true
		case 'k':
			c[1][1] = true
		default:
			return fmt.Errorf("invalid char %c", r)
		}
	}
}
