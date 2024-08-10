package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit string
type Rank string

const (
	Spades   Suit = "Spades"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Hearts   Suit = "Hearts"
)

const (
	Ace   Rank = "Ace"
	Two   Rank = "Two"
	Three Rank = "Three"
	Four  Rank = "Four"
	Five  Rank = "Five"
	Six   Rank = "Six"
	Seven Rank = "Seven"
	Eight Rank = "Eight"
	Nine  Rank = "Nine"
	Ten   Rank = "Ten"
	Jack  Rank = "Jack"
	Queen Rank = "Queen"
	King  Rank = "King"
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

func NewDeck() []Card {
	suits := []Suit{Spades, Diamonds, Clubs, Hearts}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	var deck []Card
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

func Shuffle(deck []Card) {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new Rand instance
	r.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}
