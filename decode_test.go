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
		{Piece{BLACK, Rook{}}, at("a8")},
		{Piece{BLACK, Knight{}}, at("g8")},
		{Piece{BLACK, Pawn{}}, at("a7")},
		{Piece{WHITE, Pawn{}}, at("a2")},
		{Piece{WHITE, Rook{}}, at("h1")},
	}
	var b Board
	_, err := fmt.Sscan("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", &b)
	println(b.String())
	assert.Nil(t, err)

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := b[tt.Sq]
			assert.Equal(t, tt.want, got)
		})
	}
}
