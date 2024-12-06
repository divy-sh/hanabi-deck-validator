package game

import (
	"fmt"
	"math"
)

type Game struct {
	Board           []int
	Players         [][]Card
	RemainingCards  []Card
	Hints           int
	MisfiresAllowed int
	TotalPlayers    int
	CurrentPlayer   int
	Score           int
}

func NewGame() Game {
	playerCount := 2
	newBoard := []int{}
	for i := 0; i < 5; i++ {
		newBoard = append(newBoard, 0)
	}
	cards := make([]Card)
	players := make([][]Card, playerCount)
	for i := range g.TotalPlayers {
		newGame.Players[i] = make([]Card, len(g.Players))
		copy(newGame.Players[i], g.Players[i])
	}
	game := Game{
		Board:           newBoard,
		Players:         players,
		RemainingCards:  []Card{},
		Hints:           math.MaxInt,
		MisfiresAllowed: math.MaxInt,
		TotalPlayers:    playerCount,
		CurrentPlayer:   0,
		Score:           0,
	}
	return game
}

func (g *Game) LegalMoves() []Move {
	moves := []Move{}
	player := g.Players[g.CurrentPlayer]
	for i, card := range player {
		if g.Board[card.Color]-card.Number == -1 {
			moves = append(moves, Move{SelectedCardIndex: i, Discard: false, Hint: false, Play: true})
		}
		moves = append(moves, Move{SelectedCardIndex: i, Discard: true})
	}
	moves = append(moves, Move{Hint: true})
	return moves
}

func (g *Game) PushMove(move Move) (Game, error) {
	newGame := *g

	newGame.Board = make([]int, len(g.Board))
	copy(newGame.Board, g.Board)
	newGame.Players = make([][]Card, len(g.Players))
	for i := range g.TotalPlayers {
		newGame.Players[i] = make([]Card, len(g.Players))
		copy(newGame.Players[i], g.Players[i])
	}
	newGame.RemainingCards = make([]Card, len(g.RemainingCards))
	copy(newGame.RemainingCards, g.RemainingCards)

	if !move.Hint {
		newGame.Players[g.CurrentPlayer] = append(g.Players[g.CurrentPlayer][:move.SelectedCardIndex],
			g.Players[g.CurrentPlayer][move.SelectedCardIndex+1:]...)
		if len(g.RemainingCards) > 0 {
			newGame.Players[g.CurrentPlayer] = append(g.Players[g.CurrentPlayer], g.RemainingCards[0])
			newGame.RemainingCards = g.RemainingCards[1:]
		}
		if move.Play {
			if newGame.Board[newGame.Players[g.CurrentPlayer][move.SelectedCardIndex].Color]+1 !=
				newGame.Players[g.CurrentPlayer][move.SelectedCardIndex].Number {
				g.MisfiresAllowed -= 1
			} else {
				newGame.Board[newGame.Players[g.CurrentPlayer][move.SelectedCardIndex].Color] += 1
			}
		}
	}

	newGame.updateGameScore()
	newGame.changePlayer()

	return newGame, nil
}

func (g *Game) PrintGameStatus() string {
	if !g.IsGameOver() {
		if g.CurrentPlayer == 1 {
			return "Player 1's turn"
		} else {
			return "Player 2's turn"
		}
	} else {
		return "Score: " + string(g.Score)
	}
}

func (g *Game) GetGameScore() int {
	return g.Score
}

func (g *Game) IsGameOver() bool {
	if g.MisfiresAllowed <= 0 {
		return true
	}
	for _, cards := range g.Players {
		if len(cards) > 0 {
			return false
		}
	}
	return true
}

func (g *Game) PrintBoard() string {
	fmt.Println(g.Board)
	return ""
}

func (g *Game) updateGameScore() {
	score := 0
	for _, val := range g.Board {
		score += val
	}
	g.Score = score
}

func (g *Game) changePlayer() {
	g.CurrentPlayer = (g.CurrentPlayer + 1) % g.TotalPlayers
}
