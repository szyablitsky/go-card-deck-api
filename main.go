package main

import (
	"github.com/gorilla/mux"
	"github.com/szyablitsky/go-card-deck-api/handlers"
	"log"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/decks", handlers.CreateDeckHandler).Methods(http.MethodPost)
	r.HandleFunc("/decks/{id}", handlers.OpenDeckHandler).Methods(http.MethodGet)
	r.HandleFunc("/decks/{id}/draw", handlers.DrawCardsHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}