package chess

import (
	"fmt"
	"strconv"
)

type Square uint8

const (
	FileMask  Square = 0x70
	RankMask  Square = 0x07
	BoardMask Square = 0x88

	NullSq Square = 0xFF
	Row    Square = 0x0F
)

func ScanSq(s string) (sq Square, err error) {
	if len(s) != 2 {
		err = fmt.Errorf("invalid string length, expected 2")
		return
	}
	c := s[0]
	n, _ := strconv.Atoi(s[1:])
	if c < 'a' || c > 'h' || n < 1 || n > 8 {
		err = fmt.Errorf("invalid format, expected standard notation e.g. 'e4'")
		return
	}
	file := c - 'a'
	rank := byte(n-1) << 4
	sq = Square(rank | file)
	return
}

func (sq Square) Rank() uint8 {
	return uint8(sq&FileMask) >> 4
}

func (sq Square) File() uint8 {
	return uint8(sq & RankMask)
}

func (sq Square) String() string {
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}

func (sq Square) Inbounds() bool {
	return sq&BoardMask == 0
}
