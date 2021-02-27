package chess

import (
	"fmt"
)

func Square(str string) (sq Sq) {
	_, err := fmt.Sscan(str, &sq)
	if err != nil {
		panic(err)
	}

	return sq
}

func (b *Board) Setup(color Color, moves... string) {
	var move BasicMove
	for _, str := range moves {
		_, err := fmt.Sscan(str, &move)
		if err != nil {
			panic(err)
		}
		move.Piece.Color = color
		b[move.To] = move.Piece
	}
}
