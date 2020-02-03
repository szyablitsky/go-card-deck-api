package repository

import (
	"github.com/szyablitsky/go-card-deck-api/models"
	"reflect"
	"testing"
)

var want = [52]models.Card{
	models.Card{"ACE", "CLUBS", "AC"},
	models.Card{"2", "CLUBS", "2C"},
	models.Card{"3", "CLUBS", "3C"},
	models.Card{"4", "CLUBS", "4C"},
	models.Card{"5", "CLUBS", "5C"},
	models.Card{"6", "CLUBS", "6C"},
	models.Card{"7", "CLUBS", "7C"},
	models.Card{"8", "CLUBS", "8C"},
	models.Card{"9", "CLUBS", "9C"},
	models.Card{"10", "CLUBS", "10C"},
	models.Card{"JACK", "CLUBS", "JC"},
	models.Card{"QUEEN", "CLUBS", "QC"},
	models.Card{"KING", "CLUBS", "KC"},
	models.Card{"ACE", "DIAMONDS", "AD"},
	models.Card{"2", "DIAMONDS", "2D"},
	models.Card{"3", "DIAMONDS", "3D"},
	models.Card{"4", "DIAMONDS", "4D"},
	models.Card{"5", "DIAMONDS", "5D"},
	models.Card{"6", "DIAMONDS", "6D"},
	models.Card{"7", "DIAMONDS", "7D"},
	models.Card{"8", "DIAMONDS", "8D"},
	models.Card{"9", "DIAMONDS", "9D"},
	models.Card{"10", "DIAMONDS", "10D"},
	models.Card{"JACK", "DIAMONDS", "JD"},
	models.Card{"QUEEN", "DIAMONDS", "QD"},
	models.Card{"KING", "DIAMONDS", "KD"},
	models.Card{"ACE", "HEARTS", "AH"},
	models.Card{"2", "HEARTS", "2H"},
	models.Card{"3", "HEARTS", "3H"},
	models.Card{"4", "HEARTS", "4H"},
	models.Card{"5", "HEARTS", "5H"},
	models.Card{"6", "HEARTS", "6H"},
	models.Card{"7", "HEARTS", "7H"},
	models.Card{"8", "HEARTS", "8H"},
	models.Card{"9", "HEARTS", "9H"},
	models.Card{"10", "HEARTS", "10H"},
	models.Card{"JACK", "HEARTS", "JH"},
	models.Card{"QUEEN", "HEARTS", "QH"},
	models.Card{"KING", "HEARTS", "KH"},
	models.Card{"ACE", "SPADES", "AS"},
	models.Card{"2", "SPADES", "2S"},
	models.Card{"3", "SPADES", "3S"},
	models.Card{"4", "SPADES", "4S"},
	models.Card{"5", "SPADES", "5S"},
	models.Card{"6", "SPADES", "6S"},
	models.Card{"7", "SPADES", "7S"},
	models.Card{"8", "SPADES", "8S"},
	models.Card{"9", "SPADES", "9S"},
	models.Card{"10", "SPADES", "10S"},
	models.Card{"JACK", "SPADES", "JS"},
	models.Card{"QUEEN", "SPADES", "QS"},
	models.Card{"KING", "SPADES", "KS"},
}

func TestCreateCards(t *testing.T) {
	got := CreateCards(false)

	for i, card := range got {
		if !reflect.DeepEqual(card, want[i]) {
			t.Errorf("CreateDeck(true)[%d] == %v, want %v", i, card, want[i])
		}
	}
}
