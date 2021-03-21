package main

import (
	chess "github.com/Vincent-Carrier/libchess"
)

type Human struct{}

func (h *Human) Validate(m chess.Mover, g *chess.Game) bool {
	return true
}

// func (h *Human) InvalidMove(err error) {
// 	fmt.Println("invalid move: ", err)
// }

func (h *Human) Prompt(match *chess.Match) {
	return
}
