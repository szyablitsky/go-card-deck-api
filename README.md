# go-card-deck-api

Sample Golang API server

To build and install use `go install`

Start server by typing `go-card-deck-api`. The server will be available at `http://localhost:8080`

Avalable API methods are:
- `POST /decks` to create deck. Optional params (`shuffled` and `cards`) are expected as json in the request body.
- `GET /decks/{id}` to open deck.
- `POST /decks/{id}/draw` to draw card(s). Parameter `count` is expected as json in the request body.


Run tests with `go test ./...`