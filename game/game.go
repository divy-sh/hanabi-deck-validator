package game

import (
	"fmt"
	"math/rand"
)

// Game holds the state of the game.
type Game struct {
	Board           []int
	Players         [][]Card
	RemainingCards  []Card
	DiscardedPile   []Card
	Hints           int
	MisfiresAllowed int
	TotalPlayers    int
	CurrentPlayer   int
	Score           int
}

// NewGame initializes a new Hanabi game.
func NewGame() Game {
	playerCount := 2
	colorCount := 5
	hints := 8
	misfires := 3

	// Create deck
	deck := []Card{}
	deckBuild := map[int]int{
		1: 3,
		2: 2,
		3: 2,
		4: 2,
		5: 1,
	}

	for color := 0; color < colorCount; color++ {
		for number, count := range deckBuild {
			for i := 0; i < count; i++ {
				deck = append(deck, Card{Color: color, Number: number})
			}
		}
	}

	// Shuffle deck
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	// Initialize hands
	players := make([][]Card, playerCount)
	for i := 0; i < playerCount; i++ {
		players[i] = deck[:5]
		deck = deck[5:]
	}

	// Initialize board
	board := make([]int, colorCount)

	return Game{
		Board:           board,
		Players:         players,
		RemainingCards:  deck,
		Hints:           hints,
		MisfiresAllowed: misfires,
		TotalPlayers:    playerCount,
		CurrentPlayer:   0,
		Score:           0,
	}
}

func (g *Game) LegalMoves() []Move {
	moves := []Move{}
	player := g.Players[g.CurrentPlayer]

	for i, card := range player {
		if g.Board[card.Color]+1 == card.Number {
			moves = append(moves, Move{SelectedCardIndex: i, Play: true})
		}
		moves = append(moves, Move{SelectedCardIndex: i, Discard: true})
	}

	if g.Hints > 0 {
		moves = append(moves, Move{Hint: true})
	}

	return moves
}

func (g *Game) PushMove(move Move) (Game, error) {
	newGame := g.deepCopy()

	player := newGame.CurrentPlayer
	card := newGame.Players[player][move.SelectedCardIndex]

	if move.Hint {
		if newGame.Hints <= 0 {
			return Game{}, fmt.Errorf("no hints remaining")
		}
		newGame.Hints--
	} else {
		newGame.Players[player] = append(newGame.Players[player][:move.SelectedCardIndex], newGame.Players[player][move.SelectedCardIndex+1:]...)

		if move.Play {
			if newGame.Board[card.Color]+1 == card.Number {
				newGame.Board[card.Color]++
			} else {
				newGame.MisfiresAllowed--
				if newGame.MisfiresAllowed < 0 {
					return Game{}, fmt.Errorf("game over due to misfires")
				}
			}
		} else if move.Discard {
			g.DiscardedPile = append(g.DiscardedPile, g.Players[g.CurrentPlayer][move.SelectedCardIndex])
			newGame.Hints++
		}

		if len(newGame.RemainingCards) > 0 {
			newGame.Players[player] = append(newGame.Players[player], newGame.RemainingCards[0])
			newGame.RemainingCards = newGame.RemainingCards[1:]
		}
	}

	newGame.updateGameScore()
	newGame.changePlayer()

	return newGame, nil
}

func (g *Game) deepCopy() Game {
	boardCopy := make([]int, len(g.Board))
	copy(boardCopy, g.Board)

	remainingCardsCopy := make([]Card, len(g.RemainingCards))
	copy(remainingCardsCopy, g.RemainingCards)

	playersCopy := make([][]Card, len(g.Players))
	for i := range g.Players {
		playersCopy[i] = make([]Card, len(g.Players[i]))
		copy(playersCopy[i], g.Players[i])
	}

	return Game{
		Board:           boardCopy,
		Players:         playersCopy,
		RemainingCards:  remainingCardsCopy,
		Hints:           g.Hints,
		MisfiresAllowed: g.MisfiresAllowed,
		TotalPlayers:    g.TotalPlayers,
		CurrentPlayer:   g.CurrentPlayer,
		Score:           g.Score,
	}
}

func (g *Game) PrintBoard() {
	reset := "\033[0m"
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m"}

	fmt.Println("Board Status:")
	for i, val := range g.Board {
		fmt.Printf("%s%d%s ", colors[i], val, reset)
	}
	fmt.Println()

	for i, player := range g.Players {
		fmt.Printf("Player %d's hand: ", i+1)
		for _, card := range player {
			fmt.Printf("%s%d%s ", colors[card.Color], card.Number, reset)
		}
		fmt.Println()
	}

	fmt.Println("Discard Pile:")
	for i, val := range g.DiscardedPile {
		fmt.Printf("%s%d%s ", colors[i], val, reset)
	}
	fmt.Println()

	fmt.Printf("Hints: %d\n", g.Hints)
	fmt.Printf("Misfires Remaining: %d\n", g.MisfiresAllowed)
	fmt.Printf("Remaining Cards: %d\n", len(g.RemainingCards))
}

func (g *Game) updateGameScore() {
	g.Score = 0
	for _, val := range g.Board {
		g.Score += val
	}
}

func (g *Game) changePlayer() {
	g.CurrentPlayer = (g.CurrentPlayer + 1) % g.TotalPlayers
}

func (g *Game) IsGameOver() bool {
	return g.MisfiresAllowed <= 0 || len(g.RemainingCards) == 0
}
