package main

import (
	"fmt"
	"log"
)

func main() {
	var age int

	fmt.Print("Please enter your age under here: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		log.Fatal("Invalid input. Please enter a number.")
	}

	if age < 3 {
		fmt.Println("Your ticket is free")
	} else if age >= 3 && age <= 12 {
		fmt.Println("Your ticket needs 10 dollars")
	} else if age > 12 {
		fmt.Println("Your ticket needs 15 dollars")
	}
}

// Please enter your age under here: 30
// Your ticket needs 15 dollars
