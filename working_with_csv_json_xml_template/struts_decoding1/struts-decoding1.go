package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name        string          `json:"username"`
	Permissions map[string]bool `json:"perms"`
	Devices     []struct {
		Name    string `json:"name"`
		Battery int    `json:"battery"`
	} `json:"devices"`
}

func main() {
	// Read the file directly (change the file name as needed)
	input, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse JSON into the 'users' slice
	var users []user
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Iterate over the users and print information
	for _, u := range users {
		// Print user name
		fmt.Print("+ ", u.Name)

		// Check permissions
		switch {
		case u.Permissions == nil || len(u.Permissions) == 0:
			fmt.Print(" has no power.")
		case u.Permissions["admin"]:
			fmt.Print(" is an admin.")
		case u.Permissions["write"]:
			fmt.Print(" can write.")
		default:
			fmt.Print(" has unknown permissions.")
		}

		// Print devices
		for _, device := range u.Devices {
			fmt.Printf("\n\t[ %-10s: %d%% ]", device.Name, device.Battery)
		}
		fmt.Println() // New line after each user
	}
}


// + Alice is an admin.
// 	[ Phone     : 95% ]
// 	[ Laptop    : 85% ]
// + Bob can write.
// 	[ Tablet    : 60% ]
// + Charlie has no power.
