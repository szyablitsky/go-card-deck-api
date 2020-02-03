package repository

import (
	"github.com/szyablitsky/go-card-deck-api/models"
	"math/rand"
	"time"
)

var values = [13]string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
var suits = [4]string{"CLUBS", "DIAMONDS", "HEARTS", "SPADES"}

type Cards [52]models.Card

func CreateCards(shuffled bool) Cards {
	var cards Cards

	for i, suit := range suits {
		for j, value := range values {
			cards[i*13+j] = models.NewWithCode(value, suit)
		}
	}

	if shuffled {
		shuffle(&cards)
	}

	return cards
}

func shuffle(cards *Cards) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}
