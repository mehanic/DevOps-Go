package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
//	"strings"
	"time"
)

// User represents the structure of the user data from the API.
type User struct {
	Title  string `json:"title"`
	Name   string `json:"first"`
	Surname string `json:"last"`
	TimeZone string `json:"timezone"`
}

// Response represents the structure of the response from the API.
type Response struct {
	Results []User `json:"results"`
}

func main() {
	// Check command line arguments for timezone
	var timezone string
	if len(os.Args) > 2 && os.Args[1] == "-t" {
		timezone = os.Args[2]
	}

	// Fetch user data from the API
	users, err := fetchUsers()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}

	// Filter users based on the provided timezone
	var filteredUsers []User
	for _, user := range users {
		if timezone == "" || isInTimeZone(user.TimeZone, timezone) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	// Print the results
	if timezone != "" {
		fmt.Printf("%d/%d users in timezone %s\n", len(filteredUsers), len(users), timezone)
	}
	for _, user := range filteredUsers {
		fmt.Printf("%s %s %s\n", user.Title, user.Name, user.Surname)
	}
}

// fetchUsers fetches user data from the API.
func fetchUsers() ([]User, error) {
	resp, err := http.Get("https://randomuser.me/api?results=50")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response.Results, nil
}

// isInTimeZone checks if the user's timezone matches the given timezone.
func isInTimeZone(userTZ, queryTZ string) bool {
	userOffset, err := time.ParseDuration(userTZ)
	if err != nil {
		return false
	}
	queryOffset, err := time.ParseDuration(queryTZ)
	if err != nil {
		return false
	}

	return userOffset == queryOffset
}


//go run listPersons-timeZones.go  -t +10:00                                                                                                                       0 (0.003s) < 22:38:02
//0/50 users in timezone +10:00

