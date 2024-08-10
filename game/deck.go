package game

import "fmt"

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
