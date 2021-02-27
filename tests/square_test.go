package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Vincent-Carrier/libchess"
)

func TestSquare_String(t *testing.T) {
	tests := []struct {
		sq   Sq
		want string
	}{
		{0x00, "a1"},
		{0x07, "h1"},
		{0x34, "e4"},
		{0x77, "h8"},
	}
	for _, tt := range tests {
		t.Run("to "+tt.want, func(t *testing.T) {
			if got := tt.sq.String(); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSquare_Scan(t *testing.T) {
	tests := []struct {
		want Sq
		str  string
	}{
		{0x00, "a1"},
		{0x07, "h1"},
		{0x34, "e4"},
		{0x77, "h8"},
		{0xFF, "-"},
	}
	for _, tt := range tests {
		t.Run("from "+tt.str, func(t *testing.T) {
			var got Sq
			if _, err := fmt.Sscan(tt.str, &got); got != tt.want {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSquare_Inbounds(t *testing.T) {
	tests := []struct {
		sq   Sq
		want bool
	}{
		{Square("a1"), true},
		{Square("h8"), true},
		{0x08, false},
		{0x78, false},
		{0xF0, false},
		{0x0F, false},
		{0xFF, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.sq), func(t *testing.T) {
			if got := tt.sq.Inbounds(); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
