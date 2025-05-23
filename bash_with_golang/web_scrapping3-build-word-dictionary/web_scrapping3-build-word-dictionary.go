package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Function to calculate the sum of all values in the word list
func wordListSum(wordList map[string]int) int {
	sum := 0
	for _, value := range wordList {
		sum += value
	}
	return sum
}

// Function to retrieve a random word based on weights in the word list
func retrieveRandomWord(wordList map[string]int) string {
	randIndex := rand.Intn(wordListSum(wordList)) + 1
	for word, value := range wordList {
		randIndex -= value
		if randIndex <= 0 {
			return word
		}
	}
	return ""
}

// Function to build a word dictionary based on input text
func buildWordDict(text string) map[string]map[string]int {
	// Clean the text and split words
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\"", "")
	punctuation := []string{",", ".", ";", ":"}
	for _, symbol := range punctuation {
		text = strings.ReplaceAll(text, symbol, " "+symbol+" ")
	}
	words := strings.Fields(text)

	// Build the word dictionary
	wordDict := make(map[string]map[string]int)
	for i := 1; i < len(words); i++ {
		if _, exists := wordDict[words[i-1]]; !exists {
			wordDict[words[i-1]] = make(map[string]int)
		}
		wordDict[words[i-1]][words[i]]++
	}
	return wordDict
}

func main() {
	// Fetch text data from the URL
	resp, err := http.Get("http://pythonscraping.com/files/inaugurationSpeech.txt")
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Convert the body to string
	text := string(body)

	// Build the word dictionary
	wordDict := buildWordDict(text)

	// Generate a Markov chain of length 100
	rand.Seed(time.Now().UnixNano())
	length := 100
	currentWord := "I"
	chain := ""

	for i := 0; i < length; i++ {
		chain += currentWord + " "
		currentWord = retrieveRandomWord(wordDict[currentWord])
	}

	fmt.Println(chain)
}

// I shall enter upon the subject not only of course , therefore , not satisfy his objections . Those of the citizen derives from our system with a reference to commit to fill the subjects of h
// er laws , in the most purely democratic claims them which was impossible to Congress should be supposed was the Constitution the Executive will the ruling passion by their will of ascertaini
// ng the certain purposes at this day were alarmed at work by any other , for other . It would terminate in every excrescence which I have been granted to perform is not
