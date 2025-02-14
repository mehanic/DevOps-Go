package main

import (
	"fmt"
)

func main() {
	// Initial toppings
	requestedToppings := []string{"mushrooms", "green peppers", "extra cheese"}

	if contains(requestedToppings, "mushrooms") {
		fmt.Println("Adding mushrooms.")
	}
	if contains(requestedToppings, "pepperoni") {
		fmt.Println("Adding pepperoni.")
	}
	if contains(requestedToppings, "extra cheese") {
		fmt.Println("Adding extra cheese.")

		// Assuming `requestedTopping` here should be checked
		if contains(requestedToppings, "green peppers") {
			fmt.Println("Sorry, we are run out of green peppers right now.")
		} else {
			fmt.Println("\nFinished making your pizza!")
		}
	}

	// Check if there are any requested toppings
	requestedToppings = []string{}
	if len(requestedToppings) > 0 {
		for _, topping := range requestedToppings {
			fmt.Println("Adding " + topping + ".")
		}
		fmt.Println("\nFinished making your pizza!")
	} else {
		fmt.Println("Are you sure you want a plain pizza?")
	}

	// Available and requested toppings
	availableToppings := []string{"mushrooms", "olives", "green peppers", "pepperoni", "pineapple", "extra cheese"}
	requestedToppings = []string{"mushrooms", "french fries", "extra cheese"}

	for _, topping := range requestedToppings {
		if contains(availableToppings, topping) {
			fmt.Println("Adding " + topping + ".")
		} else {
			fmt.Println("Sorry, we don't have " + topping + ".")
		}
	}

	fmt.Println("\nFinished making your pizza!")
}

// Helper function to check if a slice contains a specific value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

