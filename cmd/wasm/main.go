// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	chess "github.com/Vincent-Carrier/libchess"
)

func main() {
	done := make(chan bool)

	var Board = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		this.Set("fen", args[0])
		return nil
	})
	var BoardProto = map[string]interface{}{
		"moves": js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g, err := chess.NewGame(this.Get("fen").String())
			var sq chess.Sq
			_, err = fmt.Sscan(args[0].String(), &sq)
			if err != nil {
				return nil
			}
			piece, _ := g.At(sq)

			fmt.Println("inside Board.moves")
			var moves []string
			for _, move := range piece.Moves(g, sq) {
				fmt.Println(move.String())
				moves = append(moves, move.String())
			}
			return moves
		}),
	}
	Board.Set("prototype", BoardProto)

	js.Global().Set("Board", Board)

	<-done
}
