package main

import (
	"fmt"
)

func displayInventory(inventory map[string]int) {
	fmt.Println("Inventory:")
	itemTotal := 0
	for k, v := range inventory {
		fmt.Printf("%d %s\n", v, k)
		itemTotal += v
	}
	fmt.Printf("Total number of items: %d\n", itemTotal)
}

func addToInventory(inventory map[string]int, addedItems []string) map[string]int {
	for _, item := range addedItems {
		// Check if the item exists in the map, if not set it to 0 and then increment
		inventory[item] = inventory[item] + 1
	}
	return inventory
}

func main() {
	// Initial inventory
	inv := map[string]int{
		"gold coin": 42,
		"rope":      1,
	}
	
	// Additional loot from dragon
	dragonLoot := []string{"gold coin", "dagger", "gold coin", "gold coin", "ruby"}

	// Add loot to inventory
	inv = addToInventory(inv, dragonLoot)

	// Display updated inventory
	displayInventory(inv)

	// Example stuff inventory
	stuff := map[string]int{
		"rope":      1,
		"torch":     6,
		"gold coin": 42,
		"dagger":    1,
		"arrow":     12,
	}

	// Display stuff inventory
	displayInventory(stuff)
}

// Inventory:
// 45 gold coin
// 1 rope
// 1 dagger
// 1 ruby
// Total number of items: 48
// Inventory:
// 42 gold coin
// 1 dagger
// 12 arrow
// 1 rope
// 6 torch
// Total number of items: 62
