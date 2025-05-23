package main

import (
	"fmt"
)

// buildProfile builds a map containing everything we know about a user.
func buildProfile(first, last string, userInfo map[string]string) map[string]string {
	profile := make(map[string]string)
	profile["first name"] = first
	profile["last name"] = last

	// Print the profile dictionary
	fmt.Println("Profile:")
	for key, value := range profile {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Print user info details
	for key, value := range userInfo {
		fmt.Printf("\nkey: %s\n", key)
		fmt.Printf("value: %s\n", value)
		profile[key] = value
	}
	return profile
}

func main() {
	// Define additional user information as a map
	userInfo := map[string]string{
		"location": "princeton",
		"field":    "physics",
	}

	// Build the profile
	userProfile := buildProfile("albert", "einstein", userInfo)

	// Print the complete user profile
	fmt.Println("\nComplete user profile:")
	for key, value := range userProfile {
		fmt.Printf("%s: %s\n", key, value)
	}
}

// Profile:
// first name: albert
// last name: einstein

// key: location
// value: princeton

// key: field
// value: physics

// Complete user profile:
// first name: albert
// last name: einstein
// location: princeton
// field: physics
