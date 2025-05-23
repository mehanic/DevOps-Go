package main

import "fmt"

// Function to display inventory
func displayInventory(inventory map[string]int) {
	fmt.Println("Inventory:")
	itemTotal := 0

	// Iterate over the map
	for item, count := range inventory {
		fmt.Printf("%d %s\n", count, item)
		itemTotal += count
	}

	fmt.Printf("Total number of items: %d\n", itemTotal)
}

func main() {
	// Define the inventory map
	stuff := map[string]int{
		"rope":      1,
		"torch":     6,
		"gold coin": 42,
		"dagger":    1,
		"arrow":     12,
	}

	// Call the function to display the inventory
	displayInventory(stuff)
}

