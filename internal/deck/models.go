package deck

type Deck struct {
	DeckId   string   `json:"deck_id"`
	Shuffled bool     `json:"shuffled"`
	Cards    []string `json:"cards"`
}

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}
