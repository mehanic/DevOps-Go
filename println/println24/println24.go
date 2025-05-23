package main

import (
    "fmt"
    "sort"
    "strings"
)

// breakWords splits a sentence into words
func breakWords(stuff string) []string {
    words := strings.Split(stuff, " ")
    return words
}

// sortWords sorts a slice of words
func sortWords(words []string) []string {
    sortedWords := make([]string, len(words))
    copy(sortedWords, words)
    sort.Strings(sortedWords)
    return sortedWords
}

// printFirstWord prints the first word from a slice and removes it
func printFirstWord(words []string) {
    if len(words) > 0 {
        word := words[0]
        words = words[1:]
        fmt.Println(word)
    }
}

// printLastWord prints the last word from a slice and removes it
func printLastWord(words []string) {
    if len(words) > 0 {
        word := words[len(words)-1]
        words = words[:len(words)-1]
        fmt.Println(word)
    }
}

// sortSentence splits the sentence into words and sorts them
func sortSentence(sentence string) []string {
    words := breakWords(sentence)
    return sortWords(words)
}

// printFirstAndLast prints the first and last words of a sentence
func printFirstAndLast(sentence string) {
    words := breakWords(sentence)
    printFirstWord(words)
    printLastWord(words)
}

// printFirstAndLastSorted sorts the words in a sentence and prints the first and last
func printFirstAndLastSorted(sentence string) {
    words := sortSentence(sentence)
    printFirstWord(words)
    printLastWord(words)
}

func main() {
    fmt.Println("Let's practice everything.")
    fmt.Println("You'd need to know 'bout escapes with \\ that do \n newlines and \t tabs.")

    poem := `
    \tThe lovely world
    with logic so firmly planted
    cannot discern \n the needs of love
    nor comprehend passion from intuition
    and requires an explanation
    \n\t\twhere there is none.
    `

    fmt.Println("--------------")
    fmt.Println(poem)
    fmt.Println("--------------")

    five := 10 - 2 + 3 - 6
    fmt.Printf("This should be five: %d\n", five)

    secretFormula := func(started int) (int, int, int) {
        jellyBeans := started * 500
        jars := jellyBeans / 1000
        crates := jars / 100
        return jellyBeans, jars, crates
    }

    startPoint := 10000
    beans, jars, crates := secretFormula(startPoint)

    fmt.Printf("With a starting point of: %d\n", startPoint)
    fmt.Printf("We'd have %d beans, %d jars, and %d crates.\n", beans, jars, crates)

    startPoint = startPoint / 10

    fmt.Println("We can also do that this way:")
    beans, jars, crates = secretFormula(startPoint)
    fmt.Printf("We'd have %d beans, %d jars, and %d crates.\n", beans, jars, crates)

    sentence := "All good things come to those who wait."

    words := breakWords(sentence)
    sortedWords := sortWords(words)

    printFirstWord(words)
    printLastWord(words)
    printFirstWord(sortedWords)
    printLastWord(sortedWords)
    sortedWords = sortSentence(sentence)
    fmt.Println(sortedWords)

    printFirstAndLast(sentence)
    printFirstAndLastSorted(sentence)
}

// Let's practice everything.
// You'd need to know 'bout escapes with \ that do 
//  newlines and 	 tabs.
// --------------

//     \tThe lovely world
//     with logic so firmly planted
//     cannot discern \n the needs of love
//     nor comprehend passion from intuition
//     and requires an explanation
//     \n\t\twhere there is none.
    
// --------------
// This should be five: 5
// With a starting point of: 10000
// We'd have 5000000 beans, 5000 jars, and 50 crates.
// We can also do that this way:
// We'd have 500000 beans, 500 jars, and 5 crates.
// All
// wait.
// All
// who
// [All come good things those to wait. who]
// All
// wait.
// All
// who
