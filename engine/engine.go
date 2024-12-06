package engine

import (
	"math"

	"github.com/divy-sh/hanabi-deck-validator/game"
)

func Eval(game game.Game) *game.Move {
	bestScore := math.MinInt
	moves := game.LegalMoves()
	if len(moves) == 0 {
		return nil
	}
	bestMove := moves[0]
	for _, move := range moves {
		newBoard, _ := game.PushMove(move)
		score := maximize(newBoard, math.MinInt)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	return &bestMove
}

func maximize(game game.Game, alpha int) int {
	game.PrintBoard()
	if game.IsGameOver() {
		return game.Score
	}
	bestScore := math.MinInt
	moves := game.LegalMoves()
	for _, move := range moves {
		newBoard, _ := game.PushMove(move)
		score := maximize(newBoard, alpha)
		bestScore = max(bestScore, score)
		alpha = max(alpha, score)
	}
	return bestScore
}
