// +build js,wasm
package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/Vincent-Carrier/libchess"
)

var matches map[string]*chess.Match = make(map[string]*chess.Match)

func main() {
	js.Global().Set("chess", map[string]interface{}{
		"newMatch": js.FuncOf(newMatch),
		"moves":    js.FuncOf(moves),
		"exec":     js.FuncOf(exec),
	})

	wait := make(chan bool)
	<-wait
}

var newMatch = func(this js.Value, args []js.Value) interface{} {
	matchId := args[0].String()
	match := chess.NewMatch(&Human{}, &Human{})
	matches[matchId] = match
	this.Set("matchId", matchId)
	go match.Start()
	return nil
}

var moves = func(this js.Value, args []js.Value) interface{} {
	match := matches[this.Get("matchId").String()]
	var sq chess.Sq
	_, _ = fmt.Sscan(args[0].String(), &sq)
	var moves []interface{}
	if match.MustAt(sq).Color != match.Active {
		return moves
	}
	for _, m := range match.MovesFrom(sq) {
		moves = append(moves, m.(*chess.Slide).To.String())
	}

	return moves
}

var exec = func(this js.Value, args []js.Value) interface{} {
	match := matches[this.Get("matchId").String()]
	var move chess.Mover = new(chess.Slide)
	str := args[0].String()
	err := move.UnmarshalText([]byte(str))
	if err != nil {
		log.Fatalln("invalid move string: ", str)
	}
	match.Input <- move
	return nil
}
