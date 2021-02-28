package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Vincent-Carrier/libchess"
)

func TestNewGame(t *testing.T) {
		t.Run("starting position", func(t *testing.T) {
			got := StartingPosition()
			assert.Equal(t, Piece{WHITE, Pawn{}}, got.Board[Square("e2")])
			assert.Equal(t, WHITE, got.Active)
			assert.Equal(t, NIL_SQ, got.EnPassant)
			assert.Equal(t, 0, got.HalfMoveClock)
			assert.Equal(t, 1, got.FullMoves)
		})
}
