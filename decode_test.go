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
	tests := []struct {
		want Piece
		Sq
	}{
		{Piece{BLACK, Rook{}}, square("a8")},
		{Piece{BLACK, Knight{}}, square("g8")},
		{Piece{BLACK, Pawn{}}, square("a7")},
		{Piece{WHITE, Pawn{}}, square("a2")},
		{Piece{WHITE, Rook{}}, square("h1")},
	}
	var b Board
	_, err := fmt.Sscan("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", &b)
	assert.Nil(t, err)

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := b.MustAt(tt.Sq)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGame_Scan(t *testing.T) {
	g := StartingPosition()
	assert.Equal(t, Piece{WHITE, Pawn{}}, g.MustAt(square("e2")))
	assert.Equal(t, WHITE, g.Active)
	assert.True(t, g.Castles[0][0])
	assert.Equal(t, NIL_SQ, g.EnPassant)
	assert.Equal(t, 0, g.HalfMoveClock)
	assert.Equal(t, 1, g.FullMoves)
}
