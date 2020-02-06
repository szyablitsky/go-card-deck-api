package app

import (
	"github.com/gorilla/mux"
	"github.com/szyablitsky/go-card-deck-api/handlers"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/decks", handlers.CreateDeckHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/decks/{id}", handlers.OpenDeckHandler).Methods(http.MethodGet)
	a.Router.HandleFunc("/decks/{id}/draw", handlers.DrawCardsHandler).Methods(http.MethodPost)

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
