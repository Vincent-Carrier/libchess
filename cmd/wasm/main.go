// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	chess "github.com/Vincent-Carrier/libchess"
)

type Human struct{}

func (h *Human) Validate(m chess.Mover, g *chess.Game) bool {
	return true
}

func (h *Human) InvalidMove(err error) {
	fmt.Println("invalid move: ", err)
}

func (h *Human) Prompt() {
	// str := <-h.move
	// _, err = fmt.Sscan(str, &move)
	return
}

func main() {
	m := chess.NewMatch(&Human{}, &Human{})
	chessboard := js.Global().Get("ChessBoardElement")
	chessboard.Get("prototype").Set(
		"moves",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				var sq chess.Sq
				_, err := fmt.Sscan(args[0].String(), &sq)
				if err != nil {
					return nil
				}

				var moves []interface{}
				if piece, ok := m.At(sq); !ok || piece.Color != m.Active {
					return nil
				}
				for _, move := range m.MovesFrom(sq) {
					var sq chess.Sq
					switch move := move.(type) {
					case *chess.Slide:
						sq = move.To
					case *chess.Capture:
						sq = move.To
					}
					moves = append(moves, sq.String())
				}
				return moves
			},
		),
	)
	chessboard.Get("prototype").Set(
		"move",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				move := chess.Slide{}
				move.UnmarshalText([]byte(args[0].String()))
				m.Input <- &move
				return true
			},
		),
	)
	m.Start()
}
