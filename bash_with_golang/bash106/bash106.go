package main

import (
	"fmt"
)

func main() {
	var no string
	var name string

	// Read number input
	fmt.Print("Enter number: ")
	_, err := fmt.Scanln(&no)
	if err != nil {
		fmt.Println("Error reading number:", err)
		return
	}

	// Read name input
	fmt.Print("Enter name: ")
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	// Output the entered values
	fmt.Printf("You have entered %s %s\n", no, name)
}

