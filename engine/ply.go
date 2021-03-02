package engine

import chess "github.com/Vincent-Carrier/libchess"

type Ply struct {
	kind  chess.Mover
	Score float32
}

func (ply Ply) MarshalText() ([]byte, error) {
	return []byte(ply.kind.String()), nil
}
