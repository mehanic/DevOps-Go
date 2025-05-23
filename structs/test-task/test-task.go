package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
//	"os"
)

// User represents the structure of a user in the randomuser.me API
type User struct {
	Name struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"name"`
	Location struct {
		Timezone struct {
			Offset string `json:"offset"`
		} `json:"timezone"`
	} `json:"location"`
}

// Response is the main structure of the randomuser.me API response
type Response struct {
	Results []User `json:"results"`
}

// fetchUsers fetches the users from the randomuser.me API
func fetchUsers() ([]User, error) {
	// Make request to the API
	response, err := http.Get("https://randomuser.me/api?results=50")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON
	var apiResponse Response
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return apiResponse.Results, nil
}

// filterUsersByTimezone filters the users by the specified timezone
func filterUsersByTimezone(users []User, timezone string) ([]User, int) {
	var filteredUsers []User
	totalUsers := 0

	for _, user := range users {
		if user.Location.Timezone.Offset == timezone {
			filteredUsers = append(filteredUsers, user)
		}
		totalUsers++
	}

	return filteredUsers, totalUsers
}

// printUsers prints the user information in the specified format
func printUsers(users []User) {
	for _, user := range users {
		fmt.Printf("%s %s %s\n", user.Name.Title, user.Name.First, user.Name.Last)
	}
}

func main() {
	// Parse command-line argument for timezone
	timezone := flag.String("t", "", "Specify the timezone to filter users (e.g., +10:00)")
	flag.Parse()

	// Fetch users
	users, err := fetchUsers()
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	if *timezone != "" {
		// Filter users by timezone
		filteredUsers, totalUsers := filterUsersByTimezone(users, *timezone)

		// Print summary and users in the specified timezone
		fmt.Printf("%d/%d users in timezone %s\n", len(filteredUsers), totalUsers, *timezone)
		printUsers(filteredUsers)
	} else {
		// No timezone specified, print all users without summary
		printUsers(users)
	}
}

