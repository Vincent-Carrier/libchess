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
	var m *chess.Match
	js.Global().Set("Match", js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				m = chess.NewMatch(&Human{}, &Human{})
				m.Start()
				this.Set("board", args[0])
				return nil
			},
	))
	js.Global().Get("Match").Get("prototype").Set(
		"moves",
		js.FuncOf(
			func(this js.Value, args []js.Value) interface{} {
				var from chess.Sq
				_, err := fmt.Sscan(args[0].String(), &from)
				if err != nil {
					return nil
				}

				var moves []interface{}
				if piece, ok := m.At(from); !ok || piece.Color != m.Active {
					return nil
				}
				for _, move := range m.MovesFrom(from) {
					var sq chess.Sq
					switch move := move.(type) {
					case *chess.Slide:
						sq = move.To
					}
					moves = append(moves, sq.String())
				}
				return moves
			},
		),
	)
	js.Global().Get("Match").Get("prototype").Set(
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
}
