package handlers

import (
	"encoding/json"
	"github.com/szyablitsky/go-card-deck-api/decorators"
	"github.com/szyablitsky/go-card-deck-api/repository"
	"net/http"
)

type createDeckParams struct {
	Shuffled bool     `json:"shuffled,omitempty"`
	Cards    []string `json:"cards,omitempty"`
}

func CreateDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var params createDeckParams
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "This API endpoint should be called with 'shuffled' and 'cards' params"}`))
		return
	}

	deck := repository.CreateDeck(params.Shuffled)
	decorator := decorators.NewCreateDeckDecorator(deck)

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(decorator)
}
