package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card represents a playing card with a rank and a suit.
type Card struct {
	Rank int
	Suit string
}

// Define suits as constants
const (
	Spades   = "♠"
	Hearts   = "♡"
	Diamonds = "♢"
	Clubs    = "♣"
)

// Deck_W represents a deck of cards with shuffle and deal methods.
type Deck_W struct {
	cards     []Card
	dealIndex int
}

// NewDeck_W creates a new Deck_W with the given cards.
func NewDeck_W(cards []Card) *Deck_W {
	return &Deck_W{
		cards:     append([]Card{}, cards...), // Copy the cards slice
		dealIndex: 0,
	}
}

// Shuffle shuffles the deck of cards.
func (d *Deck_W) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.dealIndex = 0
}

// Deal returns the next card from the deck.
func (d *Deck_W) Deal() Card {
	card := d.cards[d.dealIndex]
	d.dealIndex++
	return card
}

// Deck_X represents a deck of cards by embedding a slice of Card.
type Deck_X struct {
	cards     []Card
	dealIndex int
}

// NewDeck_X creates a new Deck_X with the given cards.
func NewDeck_X(cards []Card) *Deck_X {
	return &Deck_X{
		cards:     append([]Card{}, cards...), // Copy the cards slice
		dealIndex: 0,
	}
}

// Shuffle shuffles the deck of cards.
func (d *Deck_X) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.dealIndex = 0
}

// Deal returns the next card from the deck.
func (d *Deck_X) Deal() Card {
	card := d.cards[d.dealIndex]
	d.dealIndex++
	return card
}

func main() {
	// Create a deck of cards
	domain := make([]Card, 0, 52)
	for r := 1; r <= 13; r++ {
		for _, s := range []string{Spades, Hearts, Diamonds, Clubs} {
			domain = append(domain, Card{Rank: r, Suit: s})
		}
	}

	// Test Deck_W
	dw := NewDeck_W(domain)
	rand.Seed(1)
	dw.Shuffle()
	for i := 0; i < 5; i++ {
		fmt.Println(dw.Deal())
	}

	// Test Deck_X
	dx := NewDeck_X(domain)
	rand.Seed(1)
	dx.Shuffle()
	for i := 0; i < 5; i++ {
		fmt.Println(dx.Deal())
	}
}

// {11 ♡}
// {6 ♣}
// {13 ♣}
// {10 ♠}
// {3 ♠}
// {3 ♢}
// {8 ♠}
// {12 ♢}
// {13 ♢}
// {2 ♡}
