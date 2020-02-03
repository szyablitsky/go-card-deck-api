package models

type Deck struct {
	Id        string
	Shuffled  bool
	Cards     [52]Card
	DrawCount int
}
