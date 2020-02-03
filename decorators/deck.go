package decorators

import "github.com/szyablitsky/go-card-deck-api/models"

type DeckDecorator struct {
	DeckId string `json:"deck_id"`
	Shuffled bool `json:"shuffled"`
	Remaining int `json:"remaining"`
}

func NewCreateDeckDecorator(deck models.Deck) (decorator DeckDecorator) {
	decorator.DeckId = deck.Id
	decorator.Shuffled = deck.Shuffled
	decorator.Remaining = len(deck.Cards) - deck.DrawCount
	return
}