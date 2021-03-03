package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPawn_Moves(t *testing.T) {
	g := StartingPosition()
	moves := g.MovesFrom(square("e2"))
	assert.Contains(t, moves, move("e2-e3"))
	assert.Contains(t, moves, move("e2-e4"))
	assert.Equal(t, 2, len(moves) )
}

func TestKnight_Moves(t *testing.T) {
	g := StartingPosition()
	moves := g.MovesFrom(square("b1"))
	assert.Contains(t, moves, move("b1-a3"))
	assert.Contains(t, moves, move("b1-c3"))
	assert.Equal(t, 2, len(moves))
}
