package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thebedroomprogrammer/deck-of-cards/internal/deck"
)

type DrawResponse struct {
	Cards []deck.Card `json:"cards"`
}

func (api API) Draw(w http.ResponseWriter, r *http.Request) {

	count, countErr := strconv.ParseInt(r.URL.Query().Get("count"), 10, 64)
	deckID := r.URL.Query().Get("deck_id")
	if countErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if deckID == "" || count < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if deck, ok := api.Store[deckID]; !ok || len(deck.Cards) < int(count) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	openedDeck := api.Store[deckID]

	cards := openedDeck.Cards[:count]
	openedDeck.Cards = openedDeck.Cards[count:]

	api.Store[deckID] = openedDeck

	drawnCards := deck.GetDetailedCards(cards)

	response := DrawResponse{Cards: drawnCards}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return

}
