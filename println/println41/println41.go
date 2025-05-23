package main

import (
	"fmt"
)

func main() {
	// Dictionary equivalent in Go (map)
	pizza := map[string]interface{}{
		"crust":    "thick",
		"toppings": []string{"mushroom", "extra cheese"},
	}

	// Print pizza details
	fmt.Println("You ordered a " + pizza["crust"].(string) + "-crust pizza with the following toppings:")

	for _, topping := range pizza["toppings"].([]string) {
		fmt.Println("\t" + topping)
	}

	// Function calls
	makePizza("pepperoni")
	makePizza("mushroom", "green peppers", "extra cheese")
	makePizzaWithSize(16, "pepperoni")
	makePizzaWithSize(12, "mushrooms", "green peppers", "extra cheese")
}

// Variadic function for toppings
func makePizza(toppings ...string) {
	fmt.Println("\nMaking a pizza with the following toppings:")
	for _, topping := range toppings {
		fmt.Println("- " + topping)
	}
}

// Variadic function with size argument
func makePizzaWithSize(size int, toppings ...string) {
	fmt.Printf("\nMaking a %d-inch pizza with the following toppings:\n", size)
	for _, topping := range toppings {
		fmt.Println("- " + topping)
	}
}

// You ordered a thick-crust pizza with the following toppings:
// 	mushroom
// 	extra cheese

// Making a pizza with the following toppings:
// - pepperoni

// Making a pizza with the following toppings:
// - mushroom
// - green peppers
// - extra cheese

// Making a 16-inch pizza with the following toppings:
// - pepperoni

// Making a 12-inch pizza with the following toppings:
// - mushrooms
// - green peppers
// - extra cheese
