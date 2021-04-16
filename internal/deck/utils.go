package deck

import (
	"errors"
	"math/rand"
	"time"
)

//check if the cards and valid and unique in a set
func CheckCards(cards []string) error {
	cardMap := make(map[string]int)

	for _, card := range CARDS {
		cardMap[card] = 0
	}

	for _, card := range cards {
		if cardCount, ok := cardMap[card]; !ok || cardCount > 0 {
			return errors.New("Invalid or duplicate card")
		}
		cardMap[card] = cardMap[card] + 1
	}

	return nil
}

//shuffle a deck
//check if the cards and valid and unique in a set
func Shuffle(cards []string) []string {
	shuffledCards := make([]string, len(cards))

	copy(shuffledCards, cards)

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(shuffledCards), func(i, j int) {
		shuffledCards[i], shuffledCards[j] = shuffledCards[j], shuffledCards[i]
	})

	return shuffledCards
}

func GetDetailedCards(cards []string) []Card {
	cardSet := make([]Card, len(cards))

	for i, cardCode := range cards {
		suitCode := cardCode[len(cardCode)-1]
		valueCode := cardCode[0 : len(cardCode)-1]
		value := valueCode
		if face, ok := FACES[valueCode]; ok {
			value = face
		}

		cardSet[i] = Card{Value: value, Code: cardCode, Suit: SUITS[string(suitCode)]}
	}
	return cardSet
}
