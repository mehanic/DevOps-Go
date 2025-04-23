package main

import (
    "fmt"
    "sort"

    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func main() {
    // Define favoriteLanguages map
    favoriteLanguages := map[string]interface{}{
        "jen":    "python",
        "sarah":  "c",
        "edward": "ruby",
        "phil":   "python",
    }

    // Initialize a title caser
    titleCaser := cases.Title(language.English)

    // Print Sarah's favorite language
    fmt.Println("Sarah's favorite language is " + titleCaser.String(favoriteLanguages["sarah"].(string)) + ".")

    // Iterate through favoriteLanguages and print favorite language for each person
    for name, language := range favoriteLanguages {
        fmt.Println(titleCaser.String(name) + "'s favorite language is " + titleCaser.String(language.(string)) + ".")
    }

    // Print names of everyone in the map
    for name := range favoriteLanguages {
        fmt.Println(titleCaser.String(name))
    }

    // Friends list
    friends := []string{"phil", "sarah"}

    // Iterate through the names in favoriteLanguages and check if they are in the friends list
    for name := range favoriteLanguages {
        fmt.Println(titleCaser.String(name))

        // Check if the name is in the friends list
        for _, friend := range friends {
            if name == friend {
                fmt.Println("Hi " + titleCaser.String(name) + ", I see your favorite language is " + titleCaser.String(favoriteLanguages[name].(string)) + "!")
            }
        }
    }

    // Check if 'erin' is not in the favoriteLanguages map
    if _, ok := favoriteLanguages["erin"]; !ok {
        fmt.Println("Erin, please take our poll!")
    }

    // Print names in sorted order
    names := make([]string, 0, len(favoriteLanguages))
    for name := range favoriteLanguages {
        names = append(names, name)
    }
    sort.Strings(names)

    for _, name := range names {
        fmt.Println(titleCaser.String(name) + ", thank you for taking the poll.")
    }

    // Print the languages that have been mentioned
    fmt.Println("The following languages have been mentioned:")

    // Create a set-like behavior to avoid duplicates
    languageSet := make(map[string]bool)

    for _, language := range favoriteLanguages {
        languageSet[language.(string)] = true
    }

    for language := range languageSet {
        fmt.Println(titleCaser.String(language))
    }
}
// Sarah's favorite language is C.
// Jen's favorite language is Python.
// Sarah's favorite language is C.
// Edward's favorite language is Ruby.
// Phil's favorite language is Python.
// Jen
// Sarah
// Edward
// Phil
// Jen
// Sarah
// Hi Sarah, I see your favorite language is C!
// Edward
// Phil
// Hi Phil, I see your favorite language is Python!
// Erin, please take our poll!
// Edward, thank you for taking the poll.
// Jen, thank you for taking the poll.
// Phil, thank you for taking the poll.
// Sarah, thank you for taking the poll.
// The following languages have been mentioned:
// Python
// C
// Ruby
