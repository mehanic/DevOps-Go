package main

import (
	"fmt"
	"math/rand"
//	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(20) + 1
	fmt.Println("I am thinking of a number between 1 and 20.")

	var guess int
	var guessesTaken int

	for guessesTaken = 1; guessesTaken <= 6; guessesTaken++ {
		fmt.Print("Take a guess: ")
		// Read input from the user
		var input string
		fmt.Scan(&input)
		
		// Convert input to integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			guessesTaken-- // Do not count invalid guesses
			continue
		}

		if guess < secretNumber {
			fmt.Println("Your guess is too low.")
		} else if guess > secretNumber {
			fmt.Println("Your guess is too high.")
		} else {
			break // Correct guess
		}
	}

	if guess == secretNumber {
		fmt.Printf("Good job! You guessed my number in %d guesses!\n", guessesTaken)
	} else {
		fmt.Printf("Nope. The number I was thinking of was %d\n", secretNumber)
	}
}

// I am thinking of a number between 1 and 20.
// Take a guess: 6
// Your guess is too low.
// Take a guess: 9
// Your guess is too low.
// Take a guess: 14
// Nope. The number I was thinking of was 14
