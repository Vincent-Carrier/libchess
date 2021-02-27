package chess

import (
	"fmt"
)

type Sq uint8

const (
	FileMask  Sq = 0x70
	RankMask  Sq = 0x07
	BoardMask Sq = 0x88

	NullSq Sq = 0xFF
	Row    Sq = 0x0F
)

func (sq Sq) Rank() uint8 {
	return uint8(sq&FileMask) >> 4
}

func (sq Sq) File() uint8 {
	return uint8(sq & RankMask)
}

func (sq Sq) String() string {
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}

func (sq Sq) Inbounds() bool {
	return sq&BoardMask == 0
}
