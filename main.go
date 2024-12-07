package main

import (
	"fmt"

	"github.com/divy-sh/hanabi-deck-validator/engine"
	"github.com/divy-sh/hanabi-deck-validator/game"
)

func main() {
	// play()
	evalIterate()
	// eval()
}

func eval() {
	g := game.NewGame()
	g.PrintBoard()
	_, score := engine.Eval(g)
	fmt.Printf("\nFinal Score: %d\n", score)
}

func evalIterate() {
	totalScore := 0
	totalIterations := 100000
	for range totalIterations {
		g := game.NewGame()
		_, score := engine.Eval(g)
		if score != g.MaxPossibleScore {
			g.PrintBoard()
			fmt.Printf("Max Score: %d\n\n", score)
		}
		totalScore += score
	}
	fmt.Printf("\nFinal Score: %f\n", float64(totalScore)/float64(totalIterations))
}

func play() {
	g := game.NewGame()
	for !g.IsGameOver() {
		g.PrintBoard()
		fmt.Printf("\nPlayer %d's turn.\n", g.CurrentPlayer)
		fmt.Println("Choose an action:")
		fmt.Println("1. Play a card")
		fmt.Println("2. Discard a card")
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
		case 1:
			move.Play = true
			move.SelectedCardIndex = cardIndex
		case 2:
			move.Discard = true
			move.SelectedCardIndex = cardIndex
		case 3:
			move.Hint = true
		}
		g, _ = g.PushMove(move)
	}
}
