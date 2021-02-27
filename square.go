package chess

import (
	"fmt"
)

type Sq int16

const (
	NIL_SQ Sq = -1
	ROW    Sq = 0x0F
)

func (sq Sq) Rank() int16 {
	return int16(sq&0x70) >> 4
}

func (sq Sq) File() int16 {
	return int16(sq & 0x07)
}

func (sq Sq) String() string {
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}

func (sq Sq) Inbounds() bool {
	return sq&0x88 == 0
}
