package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define a slice of messages
	messages := []string{
		"It is certain",
		"It is decidedly so",
		"Yes definitely",
		"Reply hazy try again",
		"Ask again later",
		"Concentrate and ask again",
		"My reply is no",
		"Outlook not so good",
		"Very doubtful",
	}

	// Seed the random number generator to get different results each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random index and print the message at that index
	randomIndex := rand.Intn(len(messages))
	fmt.Println(messages[randomIndex])
}

//Yes definitely