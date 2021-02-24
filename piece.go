package main

type Color int8

const (
	Black = -1
	White = 1
)

type Mover interface {
	Move(board Board, from, to Square) bool
}

type Piece struct{ Color }

type Pawn struct {
	Piece
	EnPassant bool
}

func (p Pawn) Move(board Board, from, to Square) bool {
	panic("implement me")
}

type Knight struct{ Piece }
type Bishop struct{ Piece }
type Rook struct {
	Piece
	Moved bool
}
type Queen struct{ Piece }
type King struct {
	Piece
	Moved bool
}

func PieceValue(p Mover) (f float32) {
	switch p.(type) {
	case Pawn:
		f=1.
	}
	return
}
