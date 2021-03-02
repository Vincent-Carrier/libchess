package chess

import (
	"fmt"
)

func at(str string) (sq Sq) {
	_, err := fmt.Sscan(str, &sq)
	if err != nil {
		panic(err)
	}

	return sq
}
