package main

import (
	"fmt"
	"strings"
)

func main() {
	magicians := []string{"alice", "david", "carolina"}

	for _, magician := range magicians {
		// Capitalize the first letter of the magician's name
		formattedName := strings.Title(magician)

		// Print the messages
		fmt.Printf("%s, that was a great trick!\n", formattedName)
		fmt.Printf("I can't wait to see your next trick, %s,\n\n", formattedName)
	}

	// Print the closing message
	fmt.Println("Thank you, everyone. That was a great magic show!")
}

// Alice, that was a great trick!
// I can't wait to see your next trick, Alice,

// David, that was a great trick!
// I can't wait to see your next trick, David,

// Carolina, that was a great trick!
// I can't wait to see your next trick, Carolina,

// Thank you, everyone. That was a great magic show!
