// +build js,wasm

package main

import (
	"fmt"
	chess "github.com/Vincent-Carrier/libchess"
	"syscall/js"
)

var matches map[string]chess.Match

func main() {
	js.Global().Set("chess", map[string]interface{}{
		"moves": js.FuncOf(moves),
	})

	wait := make(chan bool)
	<-wait
}

var moves = func(this js.Value, args []js.Value) interface{} {
	g, _ := chess.NewGame(args[0].String())
	var sq chess.Sq
	_, _ = fmt.Sscan(args[1].String(), &sq)
	var moves []interface{}
	for _, m := range g.MovesFrom(sq) {
		moves = append(moves, m.(*chess.Slide).To.String())
	}

	return moves
}

var move = func(this js.Value, args []js.Value) interface{} {
	return nil
}
