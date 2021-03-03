package chess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPawn_Moves(t *testing.T) {
	g := StartingPosition()
	moves := g.MovesFrom(square("e2"))
	fmt.Println(moves)
	assert.Contains(t, moves, square("e3"))
	assert.Contains(t, moves, square("e4"))
}
