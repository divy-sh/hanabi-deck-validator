package main

import (
	"fmt"
	"time"

	"github.com/divy-sh/hanabi-deck-validator/engine"
	"github.com/divy-sh/hanabi-deck-validator/game"
)

func main() {
	// play()
	timeIt("evalIterate", evalIterate)
	// eval()
}

func timeIt(name string, f func()) {
	start := time.Now()
	f()
	fmt.Printf("%s took %v to complete\n", name, time.Since(start))
}

func eval() {
	g := game.NewGame()
	g.PrintBoard()
	score := engine.Eval(g)
	fmt.Printf("\nFinal Score: %d\n", score)
}

func evalIterate() {
	scoresDistribution := map[int]int{}
	totalIterations := 10000
	for i := range totalIterations {
		fmt.Printf("\rProgress: %f", float64(i)/float64(totalIterations-1)*100)
		g := game.NewGame()
		score := engine.Eval(g)
		scoresDistribution[score] += 1
	}
	totalScore := 0
	for i := range scoresDistribution {
		totalScore += i * scoresDistribution[i]
	}
	fmt.Printf("\r\nMean Score: %f\n", float64(totalScore)/float64(totalIterations))
	fmt.Printf("Score Distribution: %v\n", scoresDistribution)
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
