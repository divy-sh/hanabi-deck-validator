package main

import (
	"github.com/divy-sh/hanabi-deck-validator/engine"
	"github.com/divy-sh/hanabi-deck-validator/game"
)

func main() {
	g := game.NewGame()
	score := engine.Eval(g)
	print(score)
}
