package engine

import (
	"math"

	"github.com/divy-sh/hanabi-deck-validator/game"
)

func Eval(game game.Game) int {
	bestScore := 0
	moves := game.LegalMoves()

	if len(moves) == 0 {
		return 0
	}
	for _, move := range game.LegalMoves() {
		newBoard, _ := game.PushMove(move)
		score := maximize(newBoard)
		if score >= game.MaxPossibleScore {
			return score
		}
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}

func maximize(game game.Game) int {
	if game.IsGameOver() {
		return game.Score
	}
	bestScore := math.MinInt
	for _, move := range game.LegalMoves() {
		newBoard, _ := game.PushMove(move)
		score := maximize(newBoard)
		if score >= game.MaxPossibleScore {
			return score
		}
		bestScore = max(bestScore, score)
	}
	return bestScore
}
