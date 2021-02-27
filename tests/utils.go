package tests

import (
	"fmt"

	. "github.com/Vincent-Carrier/libchess"
)

func ScanSq(str string) (sq Sq) {
	_, err := fmt.Sscan(str, &sq)
	if err != nil {
		panic(err)
	}

	return sq
}
