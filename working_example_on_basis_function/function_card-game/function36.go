package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card represents a playing card
type Card struct {
	Rank int
	Suit string
}

// Hand represents a collection of cards
type Hand struct {
	Hand []Card
	Bet  int
}

// NewHand creates a new Hand instance
func NewHand(bet int) *Hand {
	return &Hand{
		Hand: []Card{},
		Bet:  bet,
	}
}

// Deal adds a card to the Hand
func (h *Hand) Deal(card Card) {
	h.Hand = append(h.Hand, card)
}

// String returns a string representation of the Hand
func (h *Hand) String() string {
	return fmt.Sprintf("Hand(%d, %v)", h.Bet, h.Hand)
}

func main() {
	// Define suits and ranks
	suits := []string{
		"♠", // black spade suit
		"♡", // white heart suit
		"♦", // white diamond suit
		"♣", // black club suit
	}
	
	var deck []Card
	for rank := 1; rank <= 13; rank++ {
		for _, suit := range suits {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	
	// Shuffle the deck
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	
	// Create a new hand and deal cards
	dealer := 0
	h := NewHand(2)
	h.Deal(deck[dealer])
	dealer++
	h.Deal(deck[dealer])
	
	// Print the hand
	fmt.Println(h)
}

// Hand(2, [{7 ♦} {11 ♠}])


