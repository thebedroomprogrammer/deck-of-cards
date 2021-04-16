package deck

import (
	"reflect"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestCheckCards(t *testing.T) {
	allValidCards := []string{"AS", "KH", "QD", "10C"}
	allInvalidCards := []string{"aS", "sH", "33", "10"}
	someValidSomeInvalidCards := []string{"AS", "KH", "33", "2c"}
	validDuplicateCards := []string{"AS", "AS", "QD", "10C"}

	t.Log("Given the need to test the validity of cards")

	t.Logf("\tTest 0:\tWhen all cards are valid. Cards: %v", allValidCards)

	if err := CheckCards(allValidCards); err != nil {
		t.Errorf("\t%s\tFailed", failed)
	} else {
		t.Logf("\t%s\tPassed", succeed)

	}

	t.Logf("\tTest 1:\tWhen all cards are invalid. Cards: %v", allInvalidCards)

	if err := CheckCards(allInvalidCards); err != nil {
		t.Logf("\t%s\tPassed", succeed)
	} else {
		t.Errorf("\t%s\tFailed", failed)

	}

	t.Logf("\tTest 2:\tWhen some cards are valid and some cards are invalid. Cards: %v", someValidSomeInvalidCards)
	if err := CheckCards(someValidSomeInvalidCards); err != nil {
		t.Logf("\t%s\tPassed", succeed)
	} else {
		t.Errorf("\t%s\tFailed", failed)

	}

	t.Logf("\tTest 3:\tWhen all cards are valid but there are some duplicates. Cards: %v", validDuplicateCards)
	if err := CheckCards(validDuplicateCards); err != nil {
		t.Logf("\t%s\tPassed", succeed)
	} else {
		t.Errorf("\t%s\tFailed", failed)

	}

}

func TestShuffle(t *testing.T) {
	cards := []string{"AS", "3C", "2D", "4S", "JS"}

	t.Log("Given the need that the cards must be shuffled")
	t.Logf("\tTest 0:\tSupplied Cards: %v", cards)
	shuffledCards := Shuffle(cards)

	if reflect.DeepEqual(shuffledCards, cards) == false && len(shuffledCards) == len(cards) {
		t.Logf("\t%s\tSupplied: %v Receieved: %v", succeed, cards, shuffledCards)
	} else {
		t.Logf("\t%s\tSupplied: %v Receieved: %v", failed, cards, shuffledCards)
	}
}

func TestGetDetailedCards(t *testing.T) {
	validCardObjects := []Card{{Value: "ACE", Suit: "SPADES", Code: "AS"}, {Value: "3", Suit: "CLUBS", Code: "3C"}}
	validCardCodes := []string{"AS", "3C"}
	t.Log("Given the need that the function must return a valid slice of card objects")
	t.Logf("\tTest 0:\tWhen all cards are valid. Cards: %v", validCardCodes)
	cardSet := GetDetailedCards(validCardCodes)

	if reflect.DeepEqual(cardSet, validCardObjects) {
		t.Logf("\t%s\tExpected: %v Receieved: %v", succeed, validCardObjects, cardSet)
	} else {
		t.Logf("\t%s\tExpected: %v Receieved: %v", failed, validCardObjects, cardSet)
	}
}
