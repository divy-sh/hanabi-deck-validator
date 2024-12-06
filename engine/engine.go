package engine

import (
	"fmt"
	"math"

	"github.com/divy-sh/hanabi-deck-validator/game"
)

func Eval(game game.Game) (*game.Move, int) {
	game.PrintBoard()

	bestScore := math.MinInt
	moves := game.LegalMoves()

	fmt.Println("legalMoves:", moves)

	if len(moves) == 0 {
		return nil, bestScore
	}
	bestMove := moves[0]
	for _, move := range moves {
		newBoard, _ := game.PushMove(move)
		fmt.Println("current move:", move)
		score := maximize(newBoard)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return &bestMove, bestScore
}

func maximize(game game.Game) int {
	game.PrintBoard()
	if game.IsGameOver() {
		return game.Score
	}
	bestScore := math.MinInt
	moves := game.LegalMoves()
	for _, move := range moves {
		newBoard, _ := game.PushMove(move)
		score := maximize(newBoard)
		bestScore = max(bestScore, score)
	}
	return bestScore
}
