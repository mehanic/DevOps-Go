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
    sentence := "The quick brown fox jumps over the lazy dog"
    
    fmt.Println("Breaking and printing words:")
    printFirstAndLast(sentence)
    
    fmt.Println("\nSorted and printing words:")
    printFirstAndLastSorted(sentence)
}

