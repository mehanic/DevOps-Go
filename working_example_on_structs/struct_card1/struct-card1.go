package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Define suits as constants
const (
	Spades   = "♠"
	Hearts   = "♡"
	Diamonds = "♢"
	Clubs    = "♣"
)

// Card represents a playing card with a rank and a suit.
type Card struct {
	Rank int
	Suit string
}

// Implement the Stringer interface for custom string representation
func (c Card) String() string {
	return fmt.Sprintf("%2d %s", c.Rank, c.Suit)
}

// AceCard represents an Ace card.
type AceCard struct {
	Card
}

// Override String method for AceCard
func (a AceCard) String() string {
	return fmt.Sprintf(" A %s", a.Suit)
}

// FaceCard represents a face card (Jack, Queen, King).
type FaceCard struct {
	Card
}

// Override String method for FaceCard
func (f FaceCard) String() string {
	names := map[int]string{11: "J", 12: "Q", 13: "K"}
	return fmt.Sprintf(" %s %s", names[f.Rank], f.Suit)
}

// CribbagePoints represents points in Cribbage.
type CribbagePoints struct {
	Card
}

// Points returns the points for CribbagePoints.
func (c CribbagePoints) Points() int {
	return c.Rank
}

// CribbageFacePoints represents points for face cards in Cribbage.
type CribbageFacePoints struct {
	CribbagePoints
}

// Points returns 10 points for face cards.
func (c CribbageFacePoints) Points() int {
	return 10
}

// CribbageAce represents an Ace card with CribbagePoints.
type CribbageAce struct {
	AceCard
	CribbagePoints
}

// CribbageCard represents a normal card with CribbagePoints.
type CribbageCard struct {
	Card
	CribbagePoints
}

// CribbageFace represents a face card with CribbageFacePoints.
type CribbageFace struct {
	FaceCard
	CribbageFacePoints
}

// Create a card with Cribbage points based on rank.
func makeCard(rank int, suit string) interface{} {
	switch {
	case rank == 1:
		return CribbageAce{AceCard{Card{rank, suit}}, CribbagePoints{Card{rank, suit}}}
	case 2 <= rank && rank < 11:
		return CribbageCard{Card{rank, suit}, CribbagePoints{Card{rank, suit}}}
	case 11 <= rank:
		return CribbageFace{FaceCard{Card{rank, suit}}, CribbageFacePoints{CribbagePoints{Card{rank, suit}}}}
	default:
		return nil
	}
}

// Logged is a wrapper that adds logging functionality.
type Logged struct {
	Logger *log.Logger
}

// Points returns the points, logging the result.
func (l Logged) Points(c CribbagePoints) int {
	p := c.Points()
	l.Logger.Printf("points %d", p)
	return p
}

// LoggedCribbageAce represents a logged Ace card with CribbagePoints.
type LoggedCribbageAce struct {
	Logged
	CribbageAce
}

// LoggedCribbageCard represents a logged card with CribbagePoints.
type LoggedCribbageCard struct {
	Logged
	CribbageCard
}

// LoggedCribbageFace represents a logged face card with CribbageFacePoints.
type LoggedCribbageFace struct {
	Logged
	CribbageFace
}

// Create a logged card with Cribbage points based on rank.
func makeLoggedCard(rank int, suit string) interface{} {
	logger := log.New(log.Writer(), "", log.LstdFlags)
	switch {
	case rank == 1:
		return LoggedCribbageAce{Logged{logger}, CribbageAce{AceCard{Card{rank, suit}}, CribbagePoints{Card{rank, suit}}}}
	case 2 <= rank && rank < 11:
		return LoggedCribbageCard{Logged{logger}, CribbageCard{Card{rank, suit}, CribbagePoints{Card{rank, suit}}}}
	case 11 <= rank:
		return LoggedCribbageFace{Logged{logger}, CribbageFace{FaceCard{Card{rank, suit}}, CribbageFacePoints{CribbagePoints{Card{rank, suit}}}}}
	default:
		return nil
	}
}

func main() {
	// Create a deck of cards
	domain := make([]interface{}, 0, 52)
	for r := 1; r <= 13; r++ {
		for _, s := range []string{Spades, Hearts, Diamonds, Clubs} {
			domain = append(domain, makeCard(r, s))
		}
	}

	// Shuffle and deal the cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(domain), func(i, j int) {
		domain[i], domain[j] = domain[j], domain[i]
	})

	// Test card functions
	for i := 0; i < 5; i++ {
		if card, ok := domain[i].(Card); ok {
			fmt.Println(card)
		}
	}

	// Test logged card functions
	loggedDeck := make([]interface{}, len(domain))
	for i, card := range domain {
		switch c := card.(type) {
		case Card:
			loggedDeck[i] = makeLoggedCard(c.Rank, c.Suit)
		case CribbageFace:
			loggedDeck[i] = makeLoggedCard(c.Rank, c.Suit)
		}
	}

	for i := 0; i < 5; i++ {
		switch loggedCard := loggedDeck[i].(type) {
		case LoggedCribbageAce:
			fmt.Println(loggedCard)
			loggedCard.Points(loggedCard.CribbageAce.CribbagePoints)
		case LoggedCribbageCard:
			fmt.Println(loggedCard)
			loggedCard.Points(loggedCard.CribbageCard.CribbagePoints)
		case LoggedCribbageFace:
			// Pass the CribbagePoints from CribbageFace
			fmt.Println(loggedCard)
			loggedCard.Points(loggedCard.CribbageFace.CribbageFacePoints.CribbagePoints) // Corrected here
		}
	}
}


// J ♠
// 2025/02/13 16:23:20 points 11
//  K ♢
// 2025/02/13 16:23:20 points 13
