package game

import (
	"fmt"
	"math"

	"golang.org/x/exp/rand"
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
	// variables needed for the game
	playerCount := 2
	colorCount := 5
	DeckBuild := map[int]int{
		1: 3,
		2: 2,
		3: 2,
		4: 2,
		5: 1,
	}

	// create all cards
	cards := []Card{}
	for color := range colorCount {
		for number, count := range DeckBuild {
			for range count {
				cards = append(cards, Card{Color: color, Number: number})
			}
		}
	}

	// shuffle cards
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	// create new board
	newBoard := []int{}
	for i := 0; i < colorCount; i++ {
		newBoard = append(newBoard, 0)
	}
	players := make([][]Card, playerCount)
	for i := range players {
		players[i] = []Card{}
		for j := range 5 {
			players[i] = append(players[i], cards[i*5+j])
		}
	}
	game := Game{
		Board:           newBoard,
		Players:         players,
		RemainingCards:  cards[playerCount*5:],
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
	if g.Hints > 0 {
		moves = append(moves, Move{Hint: true})
	}
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
				newGame.MisfiresAllowed -= 1
			} else {
				newGame.Board[newGame.Players[g.CurrentPlayer][move.SelectedCardIndex].Color] += 1
			}
		}
	} else {
		newGame.Hints -= 1
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
	if len(g.RemainingCards) <= 0 {
		return true
	}
	return false
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
