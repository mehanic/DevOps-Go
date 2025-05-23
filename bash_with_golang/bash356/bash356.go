package main

import (
	"fmt"
)

func main() {
	var name1, name2 string

	// Read the first name from the user
	fmt.Print("Enter First name: ")
	fmt.Scan(&name1)

	// Read the second name from the user
	fmt.Print("Enter Second name: ")
	fmt.Scan(&name2)

	// Check if the two names are equal
	if name1 == name2 {
		fmt.Println("The names are equal.")
	} else {
		fmt.Println("The names are not equal.")
	}

	// Check if the length of the second name is greater than zero
	if len(name2) > 0 {
		fmt.Println("The second name is not empty.")
	} else {
		fmt.Println("The second name is empty.")
	}
}

