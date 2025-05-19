package main

import (
	"fmt"
	"io/ioutil"
//	"os"
	"regexp"
	"strings"
)

func main() {
	// Create and write to log file
	ioutil.WriteFile("log", []byte("12345"), 0644)

	// 1. Transform digits (0 -> 1, 1 -> 2, ..., 9 -> 0)
	transformDigits()

	// 2. Remove digits
	removeDigits()

	// 3. Keep only digits and newlines
	keepOnlyDigits()

	// 4. Squeeze spaces
	squeezeSpaces()
}

func transformDigits() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	result := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return '0' + (r-'0'+1)%10 // Transform each digit
		}
		return r
	}, string(content))

	fmt.Println("Transformed Digits:", result)
}

func removeDigits() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	// Remove all digits
	result := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return -1 // Remove digits
		}
		return r
	}, string(content))

	fmt.Println("Removed Digits:", result)
}

func keepOnlyDigits() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	// Keep only digits and newlines
	result := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' || r == '\n' {
			return r // Keep digits and newlines
		}
		return -1 // Remove everything else
	}, string(content))

	fmt.Println("Kept Only Digits:", result)
}

func squeezeSpaces() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	// Squeeze multiple spaces into a single space
	re := regexp.MustCompile(` +`) // Regular expression to match multiple spaces
	result := re.ReplaceAllString(string(content), " ") // Replace with a single space

	fmt.Println("Squeezed Spaces:", result)
}

