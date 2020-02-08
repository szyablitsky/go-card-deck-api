package handlers

import (
	"github.com/gorilla/mux"
	"github.com/szyablitsky/go-card-deck-api/decorators"
	"github.com/szyablitsky/go-card-deck-api/repository"
	"net/http"
)

func OpenDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck, present := repository.OpenDeck(id)
	if !present {
		respondWithError(w, http.StatusNotFound, "The deck you are requesting is not found")
		return
	}

	decorator := decorators.NewOpenDeckDecorator(deck)
	respondWithJSON(w, http.StatusOK, decorator)
}
