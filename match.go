package chess

type (
	Match struct {
		game *Game
		active Player
		White, Black Player
		Outcome
	}
	Player interface {
		Prompt() (Mover, error)
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
		game:  StartingPosition(),
		White: w,
		Black: b,
	}
}

func (m *Match) Start() {
	for m.Outcome == UNKNOWN {
		for {
			if move, err := m.active.Prompt(); err != nil {
				m.active.InvalidMove(err)
			} else {
				m.game.Exec(move)
				if m.active == m.White {
					m.active = m.Black
				} else {
					m.active = m.White
				}
				break
			}
		}
	}
}
