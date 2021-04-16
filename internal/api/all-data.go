package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thebedroomprogrammer/deck-of-cards/internal/deck"
)

type SmallResp struct {
	DeckID    string      `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []deck.Card `json:"cards"`
}

type resp map[string]SmallResp

func (api API) AllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OPEN")

	decks := api.Store
	respMap := make(resp)

	for _, oneDeck := range decks {
		cards := deck.GetDetailedCards(oneDeck.Cards)
		respMap[oneDeck.DeckId] = SmallResp{DeckID: oneDeck.DeckId, Shuffled: oneDeck.Shuffled, Remaining: len(oneDeck.Cards), Cards: cards}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(respMap)

	return
}
