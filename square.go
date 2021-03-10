package chess

type Sq int16

const (
	NIL_SQ Sq = -1
	ROW    Sq = 0x10
)

// Rank The square's Y position, where 0 corresponds to the white king's row
func (sq Sq) Rank() Sq {
	return (sq & 0x70) >> 4
}

// File The square's X position, where 0 corresponds the 'a' file
func (sq Sq) File() Sq {
	return sq & 0x07
}

func Coords(x, y int) Sq {
	return Sq(x | (y << 4))
}

func (sq Sq) Inbounds() bool {
	return sq&0x88 == 0
}

func (sq Sq) threatened(g *Game) {

}
