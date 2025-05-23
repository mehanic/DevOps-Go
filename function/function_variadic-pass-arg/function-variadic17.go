package main

import (
	"fmt"
)

func shoppingList2(budget int, urgent bool, items ...string) {

	fmt.Println("Shopping budget is:", budget, "dollars")

	fmt.Println("Do We need to do the shopping asap?:", urgent)

	for _, item := range items {
		fmt.Println(item)
	}

}

func main() {

	shoppingList2(100, true, "milk", "bananas", "bread")

}

// Shopping budget is: 100 dollars
// Do We need to do the shopping asap?: true
// milk
// bananas
// bread
