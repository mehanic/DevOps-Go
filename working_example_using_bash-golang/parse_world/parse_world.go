package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// parseWord parses the word and counts the number of digits, lowercase letters,
// uppercase letters, and symbols. Returns a map with the results.
func parseWord(word, symbols string) map[string]int {
	count := map[string]int{
		"d": 0, // digits
		"l": 0, // lowercase
		"u": 0, // uppercase
		"s": 0, // symbols
		"x": 0, // invalid characters
	}

	digits := "0123456789"
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for _, c := range word {
		char := string(c)
		switch {
		case strings.Contains(digits, char):
			count["d"]++
		case strings.Contains(lower, char):
			count["l"]++
		case strings.Contains(upper, char):
			count["u"]++
		case strings.Contains(symbols, char):
			count["s"]++
		default:
			count["x"]++
		}
	}

	return count
}

// parseRequirements determines which characters are required and their count.
func parseRequirements(reqStr string) map[string]int {
	req := map[string]int{
		"d": 0, // digits
		"l": 0, // lowercase
		"u": 0, // uppercase
		"s": 0, // symbols
	}

	for _, c := range reqStr {
		switch c {
		case 'd':
			req["d"]++
		case 'l':
			req["l"]++
		case 'u':
			req["u"]++
		case 's':
			req["s"]++
		}
	}

	return req
}

// complexPass checks if the word meets Windows complexity requirements.
func complexPass(count map[string]int) bool {
	if count["d"] > 0 && count["l"] > 0 && count["u"] > 0 {
		return true
	}
	if count["d"] > 0 && count["l"] > 0 && count["s"] > 0 {
		return true
	}
	if count["d"] > 0 && count["u"] > 0 && count["s"] > 0 {
		return true
	}
	if count["l"] > 0 && count["u"] > 0 && count["s"] > 0 {
		return true
	}
	return false
}

// meetsRequirements checks if the word meets the required character counts.
func meetsRequirements(count, req map[string]int) bool {
	return count["d"] >= req["d"] &&
		count["l"] >= req["l"] &&
		count["u"] >= req["u"] &&
		count["s"] >= req["s"]
}

func main() {
	// Command-line flags
	minLength := flag.Int("m", 3, "Minimum password length. (default: 3)")
	maxLength := flag.Int("x", 10, "Maximum password length. (default: 10)")
	symbols := flag.String("s", ` !"#$%&'()*+,-./:;<=>?@[\]^_`+"{|}~", `Symbols allowed in the password.`)
	wordlist := flag.String("f", "", "Wordlist to parse (default: stdin).")
	winComplexity := flag.Bool("w", false, "Passwords must meet Windows complexity requirements.")
	requirements := flag.String("r", "", "String representing the character groups and count required.")
	flag.Parse()

	// Open file or stdin
	var input *os.File
	var err error
	if *wordlist != "" {
		input, err = os.Open(*wordlist)
		if err != nil {
			fmt.Printf("Could not open file %s\n", *wordlist)
			os.Exit(1)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	// Prepare to read the wordlist
	scanner := bufio.NewScanner(input)

	// Regular expressions to match blank lines and comments
	blankLine := regexp.MustCompile(`^$`)
	commentLine := regexp.MustCompile(`^#.*$`)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip blank lines and comments
		if blankLine.MatchString(line) || commentLine.MatchString(line) {
			continue
		}

		word := strings.TrimSpace(line)
		if len(word) < *minLength || len(word) > *maxLength {
			continue
		}

		// Count character types
		count := parseWord(word, *symbols)

		// Skip words with invalid characters
		if count["x"] > 0 {
			continue
		}

		// Check requirements if provided
		if *requirements != "" {
			req := parseRequirements(*requirements)
			if meetsRequirements(count, req) {
				fmt.Println(word)
			}
			continue
		}

		// Check Windows complexity if requested
		if *winComplexity {
			if complexPass(count) {
				fmt.Println(word)
			}
			continue
		}

		// Print the word if no specific requirements or complexity are enforced
		fmt.Println(word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}

