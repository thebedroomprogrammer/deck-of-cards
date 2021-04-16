package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/thebedroomprogrammer/deck-of-cards/internal/api"
)

func TestDrawHandler(t *testing.T) {
	t.Log("Given the need that cards need to drawn from a deck")
	t.Logf("\tTest 0:\tWhen a card is removed from a full deck")
	{
		api, deckId, _ := createADeck("", "")
		req := httptest.NewRequest(http.MethodGet, "/draw", nil)
		q := req.URL.Query()
		q.Add("count", "1")
		q.Add("deck_id", deckId)

		req.URL.RawQuery = q.Encode()
		response := new(handler.DrawResponse)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Draw)
		handler.ServeHTTP(w, req)
		RunDrawDeckSuccessChecks(w, response, t, 1)
	}

	t.Logf("\tTest 1:\tWhen more cards are retrieved from a deck with less cards")
	{
		api, deckId, _ := createADeck("AS,QS,KH", "")
		req := httptest.NewRequest(http.MethodGet, "/draw", nil)
		q := req.URL.Query()
		q.Add("count", "4")
		q.Add("deck_id", deckId)

		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Draw)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}
	}

	t.Logf("\tTest 2:\tWhen Invalid deck id is supplied")
	{
		api, _, _ := createADeck("AS,QS,KH", "")
		req := httptest.NewRequest(http.MethodGet, "/draw", nil)
		q := req.URL.Query()
		q.Add("count", "4")
		q.Add("deck_id", "invalid_deck_id")

		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Draw)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}
	}

	t.Logf("\tTest 3:\tWhen invalid count value is supplied")
	{
		api, deckId, _ := createADeck("AS,QS,KH", "")
		req := httptest.NewRequest(http.MethodGet, "/draw", nil)
		q := req.URL.Query()
		q.Add("count", "abc")
		q.Add("deck_id", deckId)

		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Draw)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}
	}
}

func RunDrawDeckSuccessChecks(w *httptest.ResponseRecorder, response *handler.DrawResponse, t *testing.T, count int) {
	success := true
	if w.Code != http.StatusOK {
		success = false
		t.Fatalf("\t%s\tShould return status code %v but got %v", failed, http.StatusOK, w.Code)
	}

	if err := json.NewDecoder(w.Result().Body).Decode(response); err != nil {
		success = false
		t.Fatalf("\t%s\tInvalid response received", failed)
	}

	if count != len(response.Cards) {
		success = false
		t.Fatalf("\t%s\tMore cards are drawn than the passed count", failed)
	}

	if success {
		t.Logf("\t%s\tShould receive the expected response ", succeed)
	}
}
