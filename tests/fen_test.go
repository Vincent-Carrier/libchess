package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Vincent-Carrier/libchess"
	. "github.com/Vincent-Carrier/libchess/engine"
)

func TestNewGame(t *testing.T) {
		t.Run("starting position", func(t *testing.T) {
			got := StartingPosition()
			assert.Equal(t, Piece{White, Pawn{}}, got.Board[ScanSq("e2")])
			assert.Equal(t, White, got.Active)
			assert.Equal(t, NullSq, got.EnPassant)
			assert.Equal(t, 0, got.HalfMoveClock)
			assert.Equal(t, 1, got.FullMoves)
		})
}
