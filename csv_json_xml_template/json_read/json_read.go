package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func displayJSON() {
	f, err := os.Open("monster.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	contentJSON := make(map[string]interface{})
	err = json.NewDecoder(f).Decode(&contentJSON)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println(contentJSON)
}

func displayNameFromJSON() {
	f, err := os.Open("monster.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	contentJSON := make(map[string]interface{})
	err = json.NewDecoder(f).Decode(&contentJSON)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Welcome:", contentJSON["monsterName"])
	fmt.Println("Additions:", contentJSON["description"])
}

func main() {
	// displayJSON()
	displayNameFromJSON()
}

// Welcome: Spook
// Additions: Cracking code and battling hackers is Spook's forte. She holds a prominent position as Head of Cyber Security for the Department of Monster Defense, where she thwarts attacks on government computer systems as often as she blinks. When not at work, Spook delights in serving up a fright at haunted mansions and ghost walks.
