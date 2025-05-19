package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <word>")
		return
	}

	word := os.Args[1]
	dictionaryPath := "/usr/share/dict/words"

	file, err := os.Open(dictionaryPath)
	if err != nil {
		fmt.Printf("Error opening dictionary file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == word {
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading dictionary file: %v\n", err)
		return
	}

	if found {
		fmt.Printf("%s is a dictionary word\n", word)
	} else {
		fmt.Printf("%s is not a dictionary word\n", word)
	}
}

