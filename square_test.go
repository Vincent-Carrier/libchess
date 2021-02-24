package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare_String(t *testing.T) {
	tests := []struct {
		sq   Square
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

func TestSq(t *testing.T) {
	tests := []struct {
		want Square
		str  string
	}{
		{0x00, "a1"},
		{0x07, "h1"},
		{0x34, "e4"},
		{0x77, "h8"},
	}
	for _, tt := range tests {
		t.Run("from "+tt.str, func(t *testing.T) {
			if got := Sq(tt.str); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSquare_Inbounds(t *testing.T) {
	tests := []struct {
		sq   Square
		want bool
	}{
		{Sq("a1"), true},
		{Sq("h8"), true},
		{0x08, false},
		{0x78, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.sq), func(t *testing.T) {
			if got := tt.sq.Inbounds(); got != tt.want {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
