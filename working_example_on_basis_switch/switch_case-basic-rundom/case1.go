package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to get the answer based on the answer number
func getAnswer(answerNumber int) string {
	switch answerNumber {
	case 1:
		return "It is certain"
	case 2:
		return "It is decidedly so"
	case 3:
		return "Yes"
	case 4:
		return "Reply hazy, try again"
	case 5:
		return "Ask again later"
	case 6:
		return "Concentrate and ask again"
	case 7:
		return "My reply is no"
	case 8:
		return "Outlook not so good"
	case 9:
		return "Very doubtful"
	default:
		return "Invalid number"
	}
}

func main() {
	// Seed the random number generator to get different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 9
	r := rand.Intn(9) + 1

	// Get the fortune based on the random number
	fortune := getAnswer(r)

	// Print the fortune
	fmt.Println(fortune)
}

//Yes