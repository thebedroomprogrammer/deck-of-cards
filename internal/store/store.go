package store

import "github.com/thebedroomprogrammer/deck-of-cards/internal/deck"

func CreateStore() map[string]deck.Deck {
	store := make(map[string]deck.Deck)
	return store
}
