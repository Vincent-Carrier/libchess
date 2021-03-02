package chess

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPiece_Scan(t *testing.T) {
	t.Run("scan rook", func(t *testing.T) {
		var p Piece
		_, err := fmt.Sscan("r", &p)
		assert.Nil(t, err)
		assert.Equal(t, Piece{BLACK, Rook{}}, p)
	})

	t.Run("scan pawn", func(t *testing.T) {
		var p Piece
		_, err := fmt.Sscan("P", &p)
		assert.Nil(t, err)
		assert.Equal(t, Piece{WHITE, Pawn{}}, p)
	})
}

func TestBoard_Scan(t *testing.T) {
	t.Run("scan board", func(t *testing.T) {
		var b Board
		_, err := fmt.Sscan("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", &b)
		assert.Nil(t, err)
		assert.Equal(t, Piece{BLACK, Rook{}}, b[at("a1")])
		assert.Equal(t, Piece{WHITE, Pawn{}}, b[at("e2")])
		assert.Equal(t, Piece{WHITE, Rook{}}, b[at("h8")])
		assert.Equal(t, Piece{}, b[at("h3")])
	})

	t.Run("starting position", func(t *testing.T) {
		got := StartingPosition()
		assert.Equal(t, Piece{WHITE, Pawn{}}, got.Board[at("e2")])
		assert.Equal(t, WHITE, got.Active)
		assert.Equal(t, NIL_SQ, got.EnPassant)
		assert.Equal(t, 0, got.HalfMoveClock)
		assert.Equal(t, 1, got.FullMoves)
	})
}
