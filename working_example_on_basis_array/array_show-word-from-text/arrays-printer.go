package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please give me a word to search.")
		return
	}

	filter := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	words := strings.Fields(corpus)

queries:
	for _, q := range query {
		q = strings.ToLower(q)

		for _, v := range filter {
			if q == v {
				continue queries
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}

// #4 : "again"
