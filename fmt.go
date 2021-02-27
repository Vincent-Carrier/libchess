package chess

import (
	"fmt"
)

func (sq *Sq) Scan(state fmt.ScanState, _ rune) error {
	var file, rank byte
	r, _, err := state.ReadRune()
	if err != nil {
		return err
	} else if r == '-' {
		*sq = NullSq
		return nil
	} else if 'a' > r || r > 'h' {
		return fmt.Errorf("invalid character for square's file: %c", r)
	}
	file = uint8(r - 'a')

	r, _, err = state.ReadRune()
	if err != nil {
		return err
	} else if '1' > r || r > '8' {
		return fmt.Errorf("invalid character for square's rank: %c", r)
	}
	rank = uint8(r-'1') << 4

	*sq = Sq(rank | file)
	return nil
}
