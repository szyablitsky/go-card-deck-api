package decorators

import "github.com/szyablitsky/go-card-deck-api/models"

type CreateDeckDecorator struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func NewCreateDeckDecorator(deck models.Deck) (decorator CreateDeckDecorator) {
	decorator.DeckId = deck.Id
	decorator.Shuffled = deck.Shuffled
	decorator.Remaining = len(deck.Cards) - deck.DrawCount
	return
}

type OpenDeckDecorator struct {
	DeckId    string        `json:"deck_id"`
	Shuffled  bool          `json:"shuffled"`
	Remaining int           `json:"remaining"`
	Cards     []models.Card `json:"cards"`
}

func NewOpenDeckDecorator(deck *models.Deck) (decorator OpenDeckDecorator) {
	decorator.DeckId = deck.Id
	decorator.Shuffled = deck.Shuffled
	decorator.Remaining = len(deck.Cards) - deck.DrawCount
	decorator.Cards = deck.Cards[deck.DrawCount:]
	return
}

type DrawCardsDecorator struct {
	Cards []models.Card `json:"cards"`
}

func NewDrawCardsDecorator(deck *models.Deck, count int) (decorator DrawCardsDecorator) {
	decorator.Cards = deck.Cards[deck.DrawCount-count : deck.DrawCount]
	return
}
