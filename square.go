package chess

import (
	"fmt"
)

type Sq int16

const (
	FILE_MASK Sq = 0x70
	RankMask  Sq = 0x07
	BoardMask Sq = 0x88

	NullSq Sq = 0xFF
	Row    Sq = 0x0F
)

func (sq Sq) Rank() int16 {
	return int16(sq&FILE_MASK) >> 4
}

func (sq Sq) File() int16 {
	return int16(sq & RankMask)
}

func (sq Sq) String() string {
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}

func (sq Sq) Inbounds() bool {
	return sq&BoardMask == 0
}
