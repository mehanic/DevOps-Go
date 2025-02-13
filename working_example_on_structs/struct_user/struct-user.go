package main

import (
	"fmt"
)

// Define the User struct
type User struct {
	Name           string
	Salt           string
	HashedPassword string
}

// Sample user data
var users = []User{
	{Name: "svinci", Salt: "098u n4v04", HashedPassword: "q423uinf9304fh2u4nf3410uth1394hf"},
	{Name: "admin", Salt: "0198234nva", HashedPassword: "3894tumn13498ujc843jmcv92384vmqv"},
}

// Define the PublicUser struct without sensitive information
type PublicUser struct {
	Name string
}

// allUsers removes the sensitive information (salt and hashed_password)
func allUsers() []PublicUser {
	publicUsers := make([]PublicUser, len(users))

	// Iterate through each user and discard sensitive information
	for i, user := range users {
		publicUsers[i] = PublicUser{Name: user.Name}
	}
	return publicUsers
}

func main() {
	us := allUsers()
	for _, u := range us {
		fmt.Printf("%+v\n", u) // Print the public user struct
	}
}

// {Name:svinci}
// {Name:admin}
