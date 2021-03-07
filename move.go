package chess

import (
	"encoding"
	"fmt"
)

type (
	Mover interface {
		Execute(g *Game)
		fmt.Stringer
		encoding.TextUnmarshaler
	}

	Moves []Mover

	Slide struct {
		From, To Sq
	}
	Capture struct {
		Slide
		Capture Piece
	}
	// EnPassant struct {
	// 	Slide
	// }
)

func (m *Slide) Execute(g *Game) {
	p, _ := g.At(m.From)
	g.Set(m.To, p)
	g.Set(m.From, Piece{})
}

func (m *Slide) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%v-%v", &m.From, &m.To)
	return err
}

func (m *Slide) String() string {
	return fmt.Sprintf("%v-%v", m.From, m.To)
}
func (m *Capture) String() string {
	return fmt.Sprintf("%v-%v", m.From, m.To)
}
// func (m EnPassant) String() string {
// 	return fmt.Sprintf("%v-%v", m.From, m.To)
// }
