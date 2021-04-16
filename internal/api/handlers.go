package handler

import (
	"github.com/thebedroomprogrammer/deck-of-cards/internal/deck"
)

type API struct {
	Store map[string]deck.Deck
}
