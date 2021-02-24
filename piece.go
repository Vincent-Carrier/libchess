package main

type Color int8

const (
	White = 0
	Black = 1
)

type Mover interface {
	Move(board Board, from, to Square) bool
	Color() Color
}

type Piece struct{ color Color }

func (p Piece) Color() Color { return p.color }

func (p Pawn) Move(board Board, from, to Square) bool   { panic("implement me") }
func (k Knight) Move(board Board, from, to Square) bool { panic("implement me") }
func (b Bishop) Move(board Board, from, to Square) bool { panic("implement me") }
func (r Rook) Move(board Board, from, to Square) bool   { panic("implement me") }
func (q Queen) Move(board Board, from, to Square) bool  { panic("implement me") }
func (k King) Move(board Board, from, to Square) bool   { panic("implement me") }

type Pawn struct{ Piece }
type Knight struct{ Piece }
type Bishop struct{ Piece }
type Rook struct{ Piece }
type Queen struct{ Piece }
type King struct {
	Piece
	Checked bool
}
