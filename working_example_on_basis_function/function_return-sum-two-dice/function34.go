package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Dice represents a pair of dice
type Dice struct {
	faces [2]int
}

// Roll simulates rolling the dice
func (d *Dice) Roll() {
	d.faces[0] = rand.Intn(6) + 1
	d.faces[1] = rand.Intn(6) + 1
}

// Total returns the sum of the dice faces
func (d *Dice) Total() int {
	return d.faces[0] + d.faces[1]
}

// Hardway returns true if both dice faces are the same
func (d *Dice) Hardway() bool {
	return d.faces[0] == d.faces[1]
}

// Easyway returns true if the dice faces are different
func (d *Dice) Easyway() bool {
	return d.faces[0] != d.faces[1]
}

// TestDice performs tests similar to Python doctests
func TestDice(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	d1 := Dice{}
	d1.Roll()
	total1 := d1.Total()
	faces1 := d1.faces

	if total1 != faces1[0]+faces1[1] {
		t.Errorf("Expected total %d, got %d", faces1[0]+faces1[1], total1)
	}

	d2 := Dice{}
	d2.Roll()
	total2 := d2.Total()
	hardway := d2.Hardway()
	faces2 := d2.faces

	if total2 != faces2[0]+faces2[1] {
		t.Errorf("Expected total %d, got %d", faces2[0]+faces2[1], total2)
	}

	if hardway && faces2[0] != faces2[1] {
		t.Errorf("Hardway should be false for different faces")
	}

	if !hardway && faces2[0] == faces2[1] {
		t.Errorf("Hardway should be true for the same faces")
	}
}

func main() {
	// Example usage
	rand.Seed(time.Now().UnixNano())
	d := Dice{}
	d.Roll()
	fmt.Printf("Dice rolled: %v\n", d.faces)
	fmt.Printf("Total: %d\n", d.Total())
	fmt.Printf("Hardway: %t\n", d.Hardway())
	fmt.Printf("Easyway: %t\n", d.Easyway())
}

// Dice rolled: [5 3]
// Total: 8
// Hardway: false
// Easyway: true
