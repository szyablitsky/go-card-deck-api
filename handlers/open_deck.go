package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/szyablitsky/go-card-deck-api/decorators"
	"github.com/szyablitsky/go-card-deck-api/repository"
	"net/http"
)

func OpenDeckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

    deck, present := repository.OpenDeck(id)
    if !present {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "The deck you are requesting is not found"}`))
		return
	}

	decorator := decorators.NewOpenDeckDecorator(deck)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(decorator)
}
