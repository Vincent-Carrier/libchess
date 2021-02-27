package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/Vincent-Carrier/libchess"
)

//func TestBishop_Moves(t *testing.T) {
//	type args struct {
//		board *Board
//		color Color
//		from  Sq
//	}
//	tests := []struct {
//		name string
//		args args
//		want []Sq
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := Bishop{}
//			if got := b.Moves(tt.args.board, tt.args.color, tt.args.from); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Moves() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestKing_Moves(t *testing.T) {
//	type args struct {
//		board *Board
//		color Color
//		from  Sq
//	}
//	tests := []struct {
//		name string
//		args args
//		want []Sq
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := King{}
//			if got := k.Moves(tt.args.board, tt.args.color, tt.args.from); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Moves() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
func TestKnight_Moves(t *testing.T) {
	board := Board{}
	board.Setup(WHITE, "Ne4", "f2")
	board.Setup(BLACK, "d2")
	k := board[Square("e4")]

	tests := []struct {
		sq  Sq
		got bool
	}{
		{Square("f2"), false},
		{Square("d2"), true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("capture %v", tt.sq), func(t *testing.T) {
			got := k.Moves(&board, Square("e4"))
			assert.Contains(t, got, tt.sq)
		})
	}
}

//
//func TestPawn_Moves(t *testing.T) {
//	type args struct {
//		board *Board
//		color Color
//		from  Sq
//	}
//	tests := []struct {
//		name string
//		args args
//		want []Sq
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := Pawn{}
//			if got := p.Moves(tt.args.board, tt.args.color, tt.args.from); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Moves() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestQueen_Moves(t *testing.T) {
//	type args struct {
//		board *Board
//		color Color
//		from  Sq
//	}
//	tests := []struct {
//		name string
//		args args
//		want []Sq
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			q := Queen{}
//			if got := q.Moves(tt.args.board, tt.args.color, tt.args.from); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Moves() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestRook_Moves(t *testing.T) {
//	type args struct {
//		board *Board
//		color Color
//		from  Sq
//	}
//	tests := []struct {
//		name string
//		args args
//		want []Sq
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := Rook{}
//			if got := r.Moves(tt.args.board, tt.args.color, tt.args.from); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Moves() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
