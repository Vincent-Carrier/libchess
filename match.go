package chess

type (
	Match struct {
		*Game
		ActivePlayer Player
		White, Black Player
		History      []Mover
		Outcome
		Input chan Mover
	}
	Player interface {
		Prompt()
		InvalidMove(err error)
		Validate(m Mover, g *Game) bool
	}
	Outcome int
)

const (
	UNKNOWN Outcome = iota
	WIN
	DRAW
)

func NewMatch(w, b Player) (m *Match) {
	return &Match{
		Game:         StartingPosition(),
		ActivePlayer: w,
		White:        w,
		Black:        b,
		Input:        make(chan Mover),
	}
}

func (m *Match) Start() {
	for move := range m.Input {
		m.Game.Exec(move)
		if m.ActivePlayer == m.White {
			m.ActivePlayer = m.Black
		} else {
			m.ActivePlayer = m.White
		}
	}
}
