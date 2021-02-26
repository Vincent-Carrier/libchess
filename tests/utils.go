package tests

import (
	"log"

	"github.com/Vincent-Carrier/libchess"
)

func Sq(s string) chess.Square {
	sq, err := chess.ScanSq(s)
	Crash(err)
	return sq
}

func Crash(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
