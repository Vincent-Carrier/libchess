package main

import (
	"fmt"
	"strconv"
)

type Square byte

const (
	FileMask  = 0b0111_0000
	RankMask  = 0b0000_0111
	BoardMask = 0b1000_1000
)

func (sq Square) Rank() byte {
	return byte(sq&FileMask) >> 4
}

func (sq Square) File() byte {
	return byte(sq & RankMask)
}

func (sq Square) String() string {
	return fmt.Sprintf("%c%d", 'a'+sq.File(), sq.Rank()+1)
}

func ParseSq(s string) (sq Square, err error) {
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

func (sq Square) Inbounds() bool {
	return sq&BoardMask == 0
}
