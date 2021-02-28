package chess

type (
	Mover interface {
		//Validate(g *Game) bool
		Execute(g *Game)
	}
	Moves []Mover

	Move struct {
		From, To Sq
	}
	Capture struct {
		Move
		Capture Piece
	}
	EnPassant struct {
		Move
	}
)

func (m Move) Execute(g *Game) {
	panic("implement me")
}
func (m Capture) Execute(g *Game) {
	panic("implement me")
}
func (m EnPassant) Execute(g *Game) {
	panic("implement me")
}
