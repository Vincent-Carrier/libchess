package chess

import (
	"fmt"
	"strings"
)

func NewGame(fen string) (g *Game, err error) {
	parts := strings.Split(fen, " ")
	g = new(Game)
	_, err = fmt.Sscan(parts[0], &g.Board)
	_, err = fmt.Sscan(parts[1], &g.Active)
	_, err = fmt.Sscan(parts[2], &g.Castles)
	_, err = fmt.Sscan(parts[3], &g.EnPassant)
	_, err = fmt.Sscan(parts[4], &g.HalfMoveClock)
	_, err = fmt.Sscan(parts[5], &g.FullMoves)

	return
}

func StartingPosition() *Game {
	game, err := NewGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
	return game
}
