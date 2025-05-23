package main

import (
	"fmt"
	"strings"
)

func main() {
	guests := []string{"chima", "daniel", "kosomo"} // Invited guest list

	for _, guest := range guests {
		invitation := fmt.Sprintf("Hi %s, would you care to join us for dinner tonight?", strings.Title(guest))
		fmt.Println(invitation)
	}
}

// Hi Chima, would you care to join us for dinner tonight?
// Hi Daniel, would you care to join us for dinner tonight?
// Hi Kosomo, would you care to join us for dinner tonight?
