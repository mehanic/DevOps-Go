package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func findMostCommonWord() {
	files, _ := filepath.Glob("*.txt")
	tracker := make(map[string]int)
	fmt.Printf("Number of Txt Files: %d\n", len(files))

	for _, file := range files {
		content, _ := ioutil.ReadFile(file)
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			words := strings.Fields(line)
			for _, word := range words {
				word = strings.ReplaceAll(strings.ReplaceAll(word, ",", ""), ".", "")
				word = strings.ToLower(word)
				tracker[word]++
			}
		}
	}

	maxKey := ""
	maxValue := 0
	for key, value := range tracker {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}

	fmt.Printf("Most common word: %s\n", maxKey)
	fmt.Printf("How many times: %d\n", maxValue)
	fmt.Printf("Dictionary: %v\n", tracker)
}

func main() {
	findMostCommonWord()
}
