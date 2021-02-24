package main

type Board [0x77]Mover
type Game struct {
	Board
	Active        Color
	Castles       [2][2]bool
	EnPassant     Square
	HalfMoveClock uint
	FullMoves     uint
}

func Fen(fen string) (g Game) {
	return
}

func StartingPosition() Game {
	return Fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
}
