package test

import (
	"William-Young-97/s-head/game"
	"math/rand"
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

func TestShuffleWithFixedSeed(t *testing.T) {
	originalDeck := game.NewDeck()
	shuffledDeck := make([]game.Card, len(originalDeck))
	copy(shuffledDeck, originalDeck)

	// Use a fixed seed for reproducibility in the test
	r := rand.New(rand.NewSource(42))
	r.Shuffle(len(shuffledDeck), func(i, j int) {
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	})

	// Assertion 1: Check that the length of the deck remains the same
	if len(shuffledDeck) != len(originalDeck) {
		t.Fatalf("Shuffled deck has different length: got %d, want %d", len(shuffledDeck), len(originalDeck))
	}

	// Assertion 2: Check that the order of the cards is different
	sameOrder := true
	for i := range originalDeck {
		if originalDeck[i] != shuffledDeck[i] {
			sameOrder = false
			break
		}
	}
	if sameOrder {
		t.Error("Shuffle did not change the order of the deck")
	}

	// Assertion 3: Ensure all cards are still present and there are no duplicates
	cardCount := make(map[string]int)
	for _, card := range originalDeck {
		cardCount[card.String()]++
	}
	for _, card := range shuffledDeck {
		cardCount[card.String()]--
	}
	for card, count := range cardCount {
		if count != 0 {
			t.Errorf("Card %s has mismatched count after shuffling: expected 0, got %d", card, count)
		}
	}
}
