package repository

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	wantShuffled := true
	deck, _ := CreateDeck(wantShuffled, []string{})

	if deck.DrawCount != 0 {
		t.Errorf("CreateDeck().Shuffled == 0, got %d", deck.DrawCount)
	}

	if deck.Shuffled != wantShuffled {
		t.Errorf("CreateDeck().Shuffled == %v, got %v", wantShuffled, deck.Shuffled)
	}

	if _, err := uuid.FromString(deck.Id); err != nil {
		t.Errorf("CreateDeck().Id should be valid UUID, error %v", err)
	}
}
