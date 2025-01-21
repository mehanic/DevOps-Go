package main

import (
	"fmt"
)

func displayInventory(inventory map[string]int) {
	fmt.Println("Inventory:")
	itemTotal := 0

	// Loop through the inventory items
	for k, v := range inventory {
		fmt.Printf("%d %s\n", v, k)
		itemTotal += v
	}

	// Display total number of items
	fmt.Println("Total number of items:", itemTotal)
}

func main() {
	// Inventory items
	stuff := map[string]int{
		"rope":      1,
		"torch":     6,
		"gold coin": 42,
		"dagger":    1,
		"arrow":     12,
	}

	// Call displayInventory function to show the inventory
	displayInventory(stuff)
}

// Inventory:
// 12 arrow
// 1 rope
// 6 torch
// 42 gold coin
// 1 dagger
// Total number of items: 62
