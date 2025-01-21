package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	// Define favoriteLanguages map
	favoriteLanguages := map[string]interface{}{
		"jen":    []string{"python", "ruby"},
		"sarah":  "c",
		"edward": []string{"ruby", "go"},
		"phil":   []string{"python", "haskell"},
		"Jack":   "c",
	}

	// Printing Sarah's favorite language
	if language, ok := favoriteLanguages["sarah"].(string); ok {
		fmt.Printf("Sarah's favorite language is %s.\n", strings.Title(language))
	}

	// Iterating over the favoriteLanguages map and printing each person's favorite language(s)
	for name, languages := range favoriteLanguages {
		switch lang := languages.(type) {
		case string:
			fmt.Printf("%s's favorite language is %s.\n", strings.Title(name), strings.Title(lang))
		case []string:
			fmt.Printf("%s's favorite languages are:\n", strings.Title(name))
			for _, language := range lang {
				fmt.Printf("\t%s\n", strings.Title(language))
			}
		}
	}

	// Printing only the names of people in favoriteLanguages
	for name := range favoriteLanguages {
		fmt.Println(strings.Title(name))
	}

	// Checking for friends' favorite languages
	friends := []string{"phil", "sarah"}
	for name := range favoriteLanguages {
		fmt.Println(strings.Title(name))

		if contains(friends, name) {
			switch lang := favoriteLanguages[name].(type) {
			case string:
				fmt.Printf("Hi %s, I see your favorite language is %s!\n", strings.Title(name), strings.Title(lang))
			case []string:
				fmt.Printf("Hi %s, I see your favorite languages are: %s!\n", strings.Title(name), strings.Join(lang, ", "))
			}
		}
	}

	// Poll for people not in favoriteLanguages
	if _, ok := favoriteLanguages["erin"]; !ok {
		fmt.Println("Erin, please take our poll!")
	}

	// Sorting the names of the people and printing a thank you message
	names := make([]string, 0, len(favoriteLanguages))
	for name := range favoriteLanguages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s, thank you for taking the poll.\n", strings.Title(name))
	}

	// Printing all mentioned languages
	mentionedLanguages := make(map[string]bool)
	for _, lang := range favoriteLanguages {
		switch languages := lang.(type) {
		case string:
			mentionedLanguages[languages] = true
		case []string:
			for _, language := range languages {
				mentionedLanguages[language] = true
			}
		}
	}

	fmt.Println("The following languages have been mentioned:")
	for language := range mentionedLanguages {
		fmt.Println(strings.Title(language))
	}
}

// Helper function to check if a slice contains a specific value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Sarah's favorite language is C.
// Sarah's favorite language is C.
// Edward's favorite languages are:
// 	Ruby
// 	Go
// Phil's favorite languages are:
// 	Python
// 	Haskell
// Jack's favorite language is C.
// Jen's favorite languages are:
// 	Python
// 	Ruby
// Jen
// Sarah
// Edward
// Phil
// Jack
// Jen
// Sarah
// Hi Sarah, I see your favorite language is C!
// Edward
// Phil
// Hi Phil, I see your favorite languages are: python, haskell!
// Jack
// Erin, please take our poll!
// Jack, thank you for taking the poll.
// Edward, thank you for taking the poll.
// Jen, thank you for taking the poll.
// Phil, thank you for taking the poll.
// Sarah, thank you for taking the poll.
// The following languages have been mentioned:
// Ruby
// Go
// Python
// Haskell
// C
