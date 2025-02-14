package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
}

func main() {
	// Read all input at once
	input, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	// Parse JSON
	var users []user
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	// Process users
	for _, u := range users {
		fmt.Printf("+ %s", u.Name)

		p := u.Permissions
		if len(p) == 0 {
			fmt.Print(" has no power.")
		} else {
			if p["admin"] {
				fmt.Print(" is an admin.")
			} else if p["write"] {
				fmt.Print(" can write.")
			} else {
				fmt.Print(" has unknown permissions.")
			}
		}
		fmt.Println()
	}
}

// + inanc has no power.
// + god is an admin.
// + devil can write.
