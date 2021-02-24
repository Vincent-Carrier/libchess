package main

import "log"

func Sq(s string) Square {
	sq, err := ParseSq(s)
	Crash(err)
	return sq
}

func Crash(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
