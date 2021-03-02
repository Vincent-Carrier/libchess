package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_String(t *testing.T) {
	t.Run("starting board", func(t *testing.T) {
		g := StartingPosition()
		fen := g.Board.String()
		assert.Equal(t, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", fen)

	})
	t.Run("e2-e4", func(t *testing.T) {
		g := StartingPosition()
		fen := g.Board.String()
		assert.Equal(t, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR", fen)
	})
}
