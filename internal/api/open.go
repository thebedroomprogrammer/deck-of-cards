package handler

import (
	"encoding/json"
	"net/http"

	"github.com/thebedroomprogrammer/deck-of-cards/internal/deck"
)

type OpenResponse struct {
	DeckID    string      `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []deck.Card `json:"cards"`
}

func (api API) Open(w http.ResponseWriter, r *http.Request) {
	deckID := r.URL.Query().Get("deck_id")

	if deckID == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	openedDeck, ok := api.Store[deckID]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cards := deck.GetDetailedCards(openedDeck.Cards)

	res := OpenResponse{DeckID: openedDeck.DeckId, Shuffled: openedDeck.Shuffled, Remaining: len(openedDeck.Cards), Cards: cards}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
