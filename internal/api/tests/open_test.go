package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"
	handler "github.com/thebedroomprogrammer/deck-of-cards/internal/api"
	"github.com/thebedroomprogrammer/deck-of-cards/internal/deck"
)

func createADeck(cards string, shuffled string) (handler.API, string, error) {
	response := new(handler.CreateResponse)
	req := httptest.NewRequest(http.MethodGet, "/create", nil)
	api := NewApi()
	if cards != "" {
		q := req.URL.Query()
		q.Add("cards", cards)
		req.URL.RawQuery = q.Encode()
	}
	if shuffled != "" {
		q := req.URL.Query()
		q.Add("shuffle", shuffled)
		req.URL.RawQuery = q.Encode()
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Create)
	handler.ServeHTTP(w, req)
	json.NewDecoder(w.Result().Body).Decode(response)
	return api, response.DeckId, nil
}

func TestOpenDeckHandler(t *testing.T) {
	t.Log("Given the need that a deck is opened")
	t.Logf("\tTest 0:\tWhen the deck is opened with a valid deck id")
	{
		api, deckId, _ := createADeck("", "")
		req := httptest.NewRequest(http.MethodGet, "/open", nil)
		q := req.URL.Query()
		q.Add("deck_id", deckId)
		req.URL.RawQuery = q.Encode()
		response := new(handler.OpenResponse)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Open)
		handler.ServeHTTP(w, req)
		RunOpenDeckSuccessChecks(w, response, t, false, 52, deckId, deck.CARDS)
	}

	t.Logf("\tTest 1:\tWhen the deck is opened with an invalid deck id")
	{
		api, _, _ := createADeck("", "")
		req := httptest.NewRequest(http.MethodGet, "/open", nil)
		q := req.URL.Query()
		q.Add("deck_id", "invalid-deck-id")
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Open)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}

	}

	t.Logf("\tTest 2:\tWhen the deck is opened without")
	{
		api, _, _ := createADeck("", "")
		req := httptest.NewRequest(http.MethodGet, "/open", nil)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Open)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}

	}

	t.Logf("\tTest 3:\tWhen the deck is created with optional params and then opened")
	{
		api, deckId, _ := createADeck("AS,3S,KH", "true")
		req := httptest.NewRequest(http.MethodGet, "/open", nil)
		q := req.URL.Query()
		q.Add("deck_id", deckId)
		req.URL.RawQuery = q.Encode()
		response := new(handler.OpenResponse)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Open)
		handler.ServeHTTP(w, req)
		RunOpenDeckSuccessChecks(w, response, t, true, 3, deckId, []string{"AS", "3S", "KH"})

	}
}

func RunOpenDeckSuccessChecks(w *httptest.ResponseRecorder, response *handler.OpenResponse, t *testing.T, shouldShuffle bool, remainingCards int, deckId string, cards []string) {
	success := true
	if w.Code != http.StatusOK {
		success = false
		t.Fatalf("\t%s\tShould return status code %v but got %v", failed, http.StatusOK, w.Code)
	}

	if err := json.NewDecoder(w.Result().Body).Decode(response); err != nil {
		success = false
		t.Fatalf("\t%s\tInvalid response received", failed)
	}

	if _, err := uuid.Parse(response.DeckID); err != nil {
		success = false
		t.Errorf("\t%s\tShould return a valid deck id. Received %v", failed, response.DeckID)
	}

	if response.DeckID != deckId {
		success = false
		t.Errorf("\t%s\tShould return a data for same deck id supplied %v", failed, response.DeckID)
	}

	if response.Remaining != remainingCards {
		success = false
		t.Errorf("\t%s\tShould receive %v cards but received %v cards", failed, remainingCards, response.Remaining)
	}
	if response.Shuffled != shouldShuffle {
		success = false
		t.Errorf("\t%s\tShuffled parameter must be %v but got %v", failed, shouldShuffle, response.Shuffled)
	}

	if reflect.DeepEqual(response.Cards, deck.GetDetailedCards(cards)) == false {
		success = false
		t.Errorf("\t%s\tCards present at the time of creation are not same as in the opened deck", failed)
	}

	if success {
		t.Logf("\t%s\tShould receive the expected response ", succeed)
	}
}
