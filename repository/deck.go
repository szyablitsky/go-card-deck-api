package repository

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/szyablitsky/go-card-deck-api/models"
	"net/http"
	"sync"
)

// in memory repository
// can be changed to database storage if needed
var decks = make(map[string]*models.Deck)
var mux = sync.Mutex{}

func CreateDeck(shuffled bool) models.Deck {
	id := fmt.Sprint(uuid.NewV4())

	deck := models.Deck{
		Id:       id,
		Shuffled: shuffled,
		Cards:    CreateCards(shuffled),
	}

	decks[id] = &deck

	return deck
}

func OpenDeck(id string) (*models.Deck, bool) {
	deck, present := decks[id]
	return deck, present
}

type drawError struct {
	status int    // http status
	err    string // error description
}

func (e *drawError) Error() string {
	return e.err
}

func (e *drawError) Status() int {
	return e.status
}

func DrawCards(id string, count int) (*models.Deck, *drawError) {
	deck, present := OpenDeck(id)
	if !present {
		return deck, &drawError{http.StatusNotFound, "Deck not found"}
	}

	if count+deck.DrawCount > len(deck.Cards) {
		return deck, &drawError{http.StatusBadRequest, "Not enough cards are remaining"}
	}

	mux.Lock()
	deck.DrawCount += count
	mux.Unlock()

	return deck, nil
}
