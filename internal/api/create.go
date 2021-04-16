package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/thebedroomprogrammer/deck-of-cards/internal/deck"
)

/*
create the standard 52-card deck of French playing cards, It includes
all thirteen ranks in each of the four suits: clubs (♣), diamonds (♦), hearts (♥)
and spades (♠).
Options:
-> Deck can be shuffled or unshuffled. By default a deck is unshuffled
-> Deck can be partioal or complete. By default it has all 52 cards. For partial deck a user must provide comma seperated card codes in the query param
*/

type CreateResponse struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func (api API) Create(w http.ResponseWriter, r *http.Request) {
	shuffleParam := r.URL.Query().Get("shuffle")
	shouldShuffle := false

	if shuffleParam == "true" || shuffleParam == "false" || shuffleParam == "" {
		shouldShuffle, _ = strconv.ParseBool(shuffleParam)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cardsParam := r.URL.Query().Get("cards")
	cardSet := []string{}

	if cardsParam != "" {
		if err := deck.CheckCards(strings.Split(cardsParam, ",")); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cardSet = strings.Split(cardsParam, ",")
	} else {
		cardSet = deck.CARDS
	}

	if shouldShuffle == true {
		deck.Shuffle(cardSet)
	}

	deckId := uuid.New().String()

	api.Store[deckId] = deck.Deck{DeckId: deckId, Shuffled: shouldShuffle, Cards: cardSet}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(CreateResponse{
		DeckId:    deckId,
		Shuffled:  shouldShuffle,
		Remaining: len(cardSet),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	return

}
