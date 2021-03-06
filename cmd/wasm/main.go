// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	chess "github.com/Vincent-Carrier/libchess"
)

func main() {
	done := make(chan bool)

	chessboard := js.Global().Get("ChessboardElement")
	chessboard.Get("prototype").Set("moves",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g, err := chess.NewGame(this.Call("fen").String())
			var sq chess.Sq
			_, err = fmt.Sscan(args[1].String(), &sq)
			if err != nil {
				return nil
			}

			var moves []interface{}
			for _, move := range g.MovesFrom(sq) {
				moves = append(moves, move.String())
			}
			return moves
		}))

	<-done
}
