package main

import (
	"fmt"

	"github.com/divy-sh/hanabi-deck-validator/engine"
	"github.com/divy-sh/hanabi-deck-validator/game"
)

func main() {
	// play()
	eval()
}

func eval() {
	g := game.NewGame()
	score := engine.Eval(g)
	fmt.Println(score)
}

func play() {
	g := game.NewGame()
	for !g.IsGameOver() {
		g.PrintBoard()
		fmt.Printf("\nPlayer %d's turn.\n", g.CurrentPlayer)
		fmt.Println("Choose an action:")
		fmt.Println("1. Discard a card")
		fmt.Println("2. Play a card")
		fmt.Println("3. Give a hint")
		var action int
		var cardIndex int
		fmt.Scanln(&action)
		if action == 1 || action == 2 {
			fmt.Println("CardIndex:")
			fmt.Scanln(&cardIndex)
		}
		var move game.Move
		switch action {
		case 1: // Discard
			move.Discard = true
			move.Play = false
			move.Hint = false
			move.SelectedCardIndex = cardIndex
		case 2: // Play
			move.Play = true
			move.Discard = false
			move.Hint = false
			move.SelectedCardIndex = cardIndex
		case 3: // Hint
			move.Hint = true
			move.Discard = false
			move.Play = false
		default:
			fmt.Println("Invalid action, please choose again.")
			continue
		}
		g, _ = g.PushMove(move)
	}
}
