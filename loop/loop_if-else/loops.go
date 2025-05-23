package main

import "fmt"

func main() {
	number := 80
	guess := 6
	// for i <= 5 {
	// 	fmt.Printf("test")
	// 	i = i + 1
	// }
	fmt.Scanln(&guess)
	for guess != number {
		if guess < number {
			fmt.Printf("Larger")
			fmt.Scanln(&guess)
		} else if guess > number {
			fmt.Printf("Smaller")
			fmt.Scanln(&guess)
		}
	}
	fmt.Printf("You guessed it!")
}


// 6
// Larger90
// Smaller89
// Smaller81
// Smaller79
// Larger80
// You guessed it!%    