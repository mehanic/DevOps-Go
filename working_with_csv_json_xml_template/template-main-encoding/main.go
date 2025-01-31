package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID int
}

var data = map[string][]byte{
	"ok":   []byte(`{"ID": 27}`),
	"fail": []byte(`{"ID": 27`),
}

func GetUser(key string) (*User, error) {
	if jsonStr, ok := data[key]; ok {
		user := &User{}
		err := json.Unmarshal(jsonStr, user)
		if err != nil {
			return nil, fmt.Errorf("Cant decode json")
		}
		return user, nil
	}
	return nil, fmt.Errorf("User doesnt exist")
}

func main() {
	// Trying to fetch a valid user
	user, err := GetUser("ok")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}

	// Trying to fetch a user with invalid JSON
	user, err = GetUser("fail")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}

	// Trying to fetch a non-existent user
	user, err = GetUser("nonexistent")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}
}

// User: &{ID:27}
// Error: Cant decode json
// Error: User doesnt exist
