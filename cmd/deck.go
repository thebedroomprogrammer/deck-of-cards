package main

import (
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/thebedroomprogrammer/deck-of-cards/internal/api"
	"github.com/thebedroomprogrammer/deck-of-cards/internal/store"
)

func main() {
	r := mux.NewRouter()
	store := store.CreateStore()
	api := handler.API{Store: store}

	r.HandleFunc("/create", api.Create)
	r.HandleFunc("/open", api.Open)
	r.HandleFunc("/draw", api.Draw)
	http.ListenAndServe(":8080", r)
}
