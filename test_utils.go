package chess

import (
	"fmt"
)

func square(str string) (sq Sq) {
	_, err := fmt.Sscan(str, &sq)
	if err != nil {
		panic(err)
	}

	return sq
}

func move(str string) Mover {
	var move Slide
	_, err := fmt.Sscanf(str, "%v-%v", &move.From, &move.To)
	if err != nil {
		panic(err)
	}

	return &move
}

func (g *Game) at(sq string) Piece {
	p := g.MustAt(square(sq))
	return p
}

func (g *Game) move(s string) error {
	var from, to Sq
	_, err := fmt.Sscanf(s, "%v-%v", &from, &to)
	p, _ := g.At(from)
	g.Set(to, p)
	g.Set(from, Piece{})
	return err
}
