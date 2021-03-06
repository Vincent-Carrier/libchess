// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	chess "github.com/Vincent-Carrier/libchess"
)

func main() {
	done := make(chan bool)

	chessboard := js.Global().Get("ChessBoardElement")
	chessboard.Get("prototype").Set("moves",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g, err := chess.NewGame(this.Call("fen").String())
			var sq chess.Sq
			_, err = fmt.Sscan(args[0].String(), &sq)
			if err != nil {
				return nil
			}

			var moves []interface{}
			if piece, ok := g.At(sq); !ok || piece.Color != g.Active {
				return moves
			}
			for _, move := range g.MovesFrom(sq) {
				var sq chess.Sq
				switch move := move.(type) {
				case chess.Move:
					sq = move.To
				case chess.Capture:
					sq = move.To
				}
				moves = append(moves, sq.String())
			}
			return moves
		}))
	chessboard.Get("prototype").Set("exec",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			g, err := chess.NewGame(this.Call("fen").String())
			var from, to chess.Sq
			_, err = fmt.Sscan(args[0].String(), &from)
			_, err = fmt.Sscan(args[1].String(), &to)
			if err != nil {
				return nil
			}

			piece, _ := g.At(from)
			if piece.Color != g.Active {
				return nil
			}
			move := chess.Move{from, to}
			g.Exec(move)

			return true
		}))

	<-done
}
