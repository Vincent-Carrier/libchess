package engine

import (
	"testing"

	. "github.com/Vincent-Carrier/libchess"
	"github.com/stretchr/testify/assert"
)

func TestSquare_String(t *testing.T) {
	tests := []struct {
		b    *Board
		want float32
	}{
		{StartingPosition().Board, 190.0},
	}
	for _, tt := range tests {
		t.Run("to "+tt.want, func(t *testing.T) {
			if got := tt.sq.String(); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
