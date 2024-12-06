package game

type Move struct {
	Discard           bool
	Hint              bool
	Play              bool
	SelectedCardIndex int
}
