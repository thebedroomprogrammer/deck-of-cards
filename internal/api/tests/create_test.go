package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	handler "github.com/thebedroomprogrammer/deck-of-cards/internal/api"
	"github.com/thebedroomprogrammer/deck-of-cards/internal/store"
)

const succeed = "\u2713"
const failed = "\u2717"

func NewApi() handler.API {
	store := store.CreateStore()
	api := handler.API{Store: store}
	return api
}

func TestCreateDeckHandler(t *testing.T) {
	t.Log("Given the need that create deck api must work fine")
	t.Logf("\tTest 0:\tWhen the deck is created without any optional parameter")
	{
		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		response := new(handler.CreateResponse)
		api := NewApi()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Create)
		handler.ServeHTTP(w, req)
		RunCreateDeckSuccessChecks(w, response, t, false, 52)
	}

	t.Logf("\tTest 1:\tWhen a shuffled deck is created")
	{
		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		response := new(handler.CreateResponse)
		api := NewApi()
		q := req.URL.Query()
		q.Add("shuffle", "true")
		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Create)
		handler.ServeHTTP(w, req)
		RunCreateDeckSuccessChecks(w, response, t, true, 52)
	}

	t.Logf("\tTest 2:\tWhen a shuffled deck is created with optional cards parameter")
	{
		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		response := new(handler.CreateResponse)
		api := NewApi()
		q := req.URL.Query()
		q.Add("shuffle", "true")
		q.Add("cards", "AS,3S")

		req.URL.RawQuery = q.Encode()

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Create)
		handler.ServeHTTP(w, req)
		RunCreateDeckSuccessChecks(w, response, t, true, 2)
	}

	t.Logf("\tTest 3:\tWhen an unshuffled deck is created with optional cards parameter")
	{
		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		response := new(handler.CreateResponse)
		api := NewApi()
		q := req.URL.Query()
		q.Add("cards", "AS,3S")
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Create)
		handler.ServeHTTP(w, req)
		RunCreateDeckSuccessChecks(w, response, t, false, 2)
	}

	t.Logf("\tTest 4:\tWhen duplicate cards are supplied in the optional parameter")
	{
		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		api := NewApi()
		q := req.URL.Query()
		q.Add("cards", "AS,AS")
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Create)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}
	}

	t.Logf("\tTest 5:\tWhen invalid optional parameter is supplied")
	{
		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		api := NewApi()
		q := req.URL.Query()
		q.Add("cards", "AS,as")
		req.URL.RawQuery = q.Encode()
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(api.Create)
		handler.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Errorf("\t%s\tShould return status code %v but got %v", failed, http.StatusBadRequest, w.Code)
		} else {
			t.Logf("\t%s\tShould return response code %d", succeed, http.StatusBadRequest)
		}
	}
}

func RunCreateDeckSuccessChecks(w *httptest.ResponseRecorder, response *handler.CreateResponse, t *testing.T, shouldShuffle bool, remainingCards int) {
	success := true

	if w.Code != http.StatusOK {
		success = false
		t.Fatalf("\t%s\tShould return status code %v but got %v", failed, http.StatusOK, w.Code)
	}

	if err := json.NewDecoder(w.Result().Body).Decode(response); err != nil {
		success = false
		t.Fatalf("\t%s\tInvalid response received", failed)
	}

	if _, err := uuid.Parse(response.DeckId); err != nil {
		success = false
		t.Errorf("\t%s\tShould return a valid deck id. Received %v", failed, response.DeckId)
	}
	if response.Remaining != remainingCards {
		success = false
		t.Errorf("\t%s\tShould receive %v cards but received %v cards", failed, remainingCards, response.Remaining)
	}
	if response.Shuffled != shouldShuffle {
		success = false
		t.Errorf("\t%s\tShuffled parameter must be %v but got %v", failed, shouldShuffle, response.Shuffled)
	}

	if success {
		t.Logf("\t%s\tShould receive the expected response %v", succeed, w.Body.String())
	}
}
