package main

import (
	"fmt"
)

func shoppingList3(budget int, urgent bool, items []string) {

	fmt.Println("Shopping budget is:", budget, "dollars")

	fmt.Println("Do We need to do the shopping asap?:", urgent)

	for _, item := range items {
		fmt.Println(item)
	}

}

func main() {

	list_of_items := []string{"milk", "bananas", "bread"}

	shoppingList3(100, true, list_of_items)

}

// Shopping budget is: 100 dollars
// Do We need to do the shopping asap?: true
// milk
// bananas
// bread
