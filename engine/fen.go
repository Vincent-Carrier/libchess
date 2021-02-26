package engine

import (
	"fmt"
	"strconv"
	"unicode"

	. "github.com/Vincent-Carrier/libchess"
)

type Fen string

func NewGame(fen Fen) (g Game, err error) {
	var (
		sq       Square = 0x70
		p        Piece
		stateIdx int
	)
	for i, r := range []rune(fen) {
		//fmt.Printf("i: %d, r: %c, sq: %#x\n", i, r, uint8(sq))
		if unicode.IsDigit(r) {
			skip, _ := strconv.Atoi(string(r))
			sq += Square(skip) - 1
		} else if unicode.IsLetter(r) {
			p, err = newPiece(r)
			g.Board[sq] = p
			sq++
		} else if r == '/' {
			sq -= Row + 8
		} else {
			stateIdx = i + 1
			break
		}
	}

	var (
		active             rune
		castles, enPassant string
	)
	_, err = fmt.Sscanf(string(fen[stateIdx:]), "%c %s %s %d %d",
		&active, &castles, &enPassant, &(g.HalfMoveClock), &(g.FullMoves))

	switch active {
	case 'w': g.Active = White
	case 'b': g.Active = Black
	}
	for _, r := range []rune(castles) {
		switch r {
		case 'Q':
			g.Castles[0][0] = true
		case 'K':
			g.Castles[0][1] = true
		case 'q':
			g.Castles[1][0] = true
		case 'k':
			g.Castles[1][1] = true
		}
	}
	if sq, err := ScanSq(enPassant); err != nil {
		g.EnPassant = NullSq
	} else {
		g.EnPassant = sq
	}

	return
}

func StartingPosition() Game {
	game, err := NewGame("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}
	return game
}

func newPiece(r rune) (Piece, error) {
	var m Mover
	var c Color
	switch unicode.ToLower(r) {
	case 'p':
		m = Pawn{}
	case 'n':
		m = Knight{}
	case 'b':
		m = Bishop{}
	case 'r':
		m = Rook{}
	case 'q':
		m = Queen{}
	case 'k':
		m = King{}
	default:
		return Piece{}, fmt.Errorf("invalid char %c", m)
	}
	if unicode.IsLower(r) {
		c = Black
	}

	return Piece{c, m}, nil
}
