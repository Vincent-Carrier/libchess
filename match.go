package chess

type (
	Match struct {
		game *Game
		White, Black Player
	}
	Player interface {
		Validate(m Mover, g *Game) bool
	}
	Outcome int
)

const (
	ONGOING Outcome = iota
	WIN
	DRAW
)

func (m *Match) Ply(move Mover) (o Outcome) {
	var p Player
	if m.game.Active == WHITE {
		p = m.White
	} else {
		p = m.Black
	}
	p.Validate(move, m.game)
	m.game.Exec(move)
	o = ONGOING
	return
}
