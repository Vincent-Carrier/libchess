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

func (p *Piece) Scan(state fmt.ScanState, verb rune) error {
	r, _, err := state.ReadRune()
	if err != nil {
		return err
	}
	colorblind := unicode.IsUpper(verb)
	if !colorblind {
		r = unicode.ToUpper(r)
	}
	switch r {
	case 'P':
		p.Piecer = Pawn{}
	case 'N':
		p.Piecer = Knight{}
	case 'B':
		p.Piecer = Bishop{}
	case 'R':
		p.Piecer = Rook{}
	case 'Q':
		p.Piecer = Queen{}
	case 'K':
		p.Piecer = King{}
	default:
		return fmt.Errorf("invalid char %c", r)
	}
	if colorblind {
		return nil
	} else if unicode.IsLower(r) {
		p.Color = BLACK
	} else {
		p.Color = WHITE
	}
	return nil
}

func (m *Move) Scan(state fmt.ScanState, _ rune) error {
	r, _, err := state.ReadRune()
	if err != nil {
		return err
	}
	switch r {
	case 'N':
		m.Piecer = Knight{}
	case 'B':
		m.Piecer = Bishop{}
	case 'R':
		m.Piecer = Rook{}
	case 'Q':
		m.Piecer = Queen{}
	case 'K':
		m.Piecer = King{}
	default:
		m.Piecer = Pawn{}
		err = state.UnreadRune()
	}

	r1, _, err := state.ReadRune()
	r2, _, err := state.ReadRune()

	if _, err = fmt.Sscan(string([]rune{r1, r2}), &m.To); err != nil {
		return err
	}
	return nil
}
