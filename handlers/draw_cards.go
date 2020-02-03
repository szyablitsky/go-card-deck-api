package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/szyablitsky/go-card-deck-api/decorators"
	"github.com/szyablitsky/go-card-deck-api/repository"
	"net/http"
)

type drawCardsParams struct {
	Count int `json:"count,omitempty"`
}

func DrawCardsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	decoder := json.NewDecoder(r.Body)
	var params drawCardsParams
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "This API endpoint should be called with 'count' param"}`))
		return
	}

	fmt.Print(params)

	deck, drawErr := repository.DrawCards(id, params.Count)
	if drawErr != nil {
		w.WriteHeader(drawErr.Status())
		str := fmt.Sprintf(`{"error": %q}`, drawErr.Error())
		_, _ = w.Write([]byte(str))
		return
	}

	decorator := decorators.NewDrawCardsDecorator(deck, params.Count)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(decorator)
}
