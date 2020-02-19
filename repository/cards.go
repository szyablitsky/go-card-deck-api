package repository

import (
	"github.com/szyablitsky/go-card-deck-api/utils"
	"math/rand"
	"time"
)

var FullDeckCodes = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "TC", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "TD", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "TH", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "TS", "JS", "QS", "KS",
}

type Cards []string

func ValidateCards(cards Cards) bool {
	for _, card := range cards {
		if !utils.Contains(FullDeckCodes[:], card) {
			return false
		}
	}
	return true
}

func CreateCards(shuffled bool, cards Cards) Cards {
	if len(cards) == 0 {
		copy := FullDeckCodes
		cards = copy[:]
	}

	if shuffled {
		cards.shuffle()
	}

	return cards
}

func (cards Cards)shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}
