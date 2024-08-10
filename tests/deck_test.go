package test

import (
	"William-Young-97/s-head/game"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := game.NewDeck()

	// Test the length of the deck
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(deck))
	}

	// Test that the first card is the Ace of Spades
	if deck[0].String() != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %s", deck[0].String())
	}

	// Test that the last card is the King of Hearts
	if deck[len(deck)-1].String() != "King of Hearts" {
		t.Errorf("Expected last card to be King of Hearts, but got %s", deck[len(deck)-1].String())
	}

	// Test that all cards are unique
	seen := make(map[string]bool)
	for _, card := range deck {
		cardStr := card.String()
		if seen[cardStr] {
			t.Errorf("Found duplicate card in deck: %s", cardStr)
		}
		seen[cardStr] = true
	}
}
