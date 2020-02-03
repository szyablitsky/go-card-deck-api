package repository

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/szyablitsky/go-card-deck-api/models"
)

// in memory repository
// can be changed to database storage if needed
var decks = make(map[string]models.Deck)

func CreateDeck(shuffled bool) models.Deck {
	id := fmt.Sprint(uuid.NewV4())

	deck := models.Deck{
		Id:       id,
		Shuffled: shuffled,
		Cards:    CreateCards(shuffled),
	}

	decks[id] = deck

	return deck
}
