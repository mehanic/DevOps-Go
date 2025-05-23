package main

import (
	"fmt"
	"io/ioutil"
//	"os"
	"regexp"
	"strings"
)

func main() {
	// 1. Replace "hello" with "goodbye" in a string
	fmt.Println("Before:", "hello world")
	str1 := strings.Replace("hello world", "hello", "goodbye", 1)
	fmt.Println("After:", str1)

	// Create 1.txt and write "hello world" to it
	err := ioutil.WriteFile("1.txt", []byte("hello world"), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// 2. Replace "hello" with "goodbye" in 1.txt
	content, err := ioutil.ReadFile("1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	contentStr := strings.Replace(string(content), "hello", "goodbye", -1) // -1 replaces all instances
	err = ioutil.WriteFile("1.txt", []byte(contentStr), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// 3. Replace all instances of "this" with "THIS"
	str2 := "thisthisthisthis"
	str2Replaced := strings.ReplaceAll(str2, "this", "THIS")
	fmt.Println("Replaced all 'this':", str2Replaced)

	// 4. Replace the second occurrence of "this" with "THIS"
	re := regexp.MustCompile(`(this)(.*?)(this)`)
	str2SecondReplaced := re.ReplaceAllString(str2, "THIS$2this")
	fmt.Println("Replaced second 'this':", str2SecondReplaced)

	// 5. Delete empty lines from diff.sh (for example purposes, we'll create a string)
	diffContent := "line 1\n\nline 2\n\nline 3\n"
	fmt.Println("Original diff content:\n", diffContent)
	lines := strings.Split(diffContent, "\n")
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	fmt.Println("After deleting empty lines:\n", strings.Join(nonEmptyLines, "\n"))

	// 6. Surround each word in a sentence with brackets
	sentence := "this is an example"
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = "[" + word + "]"
	}
	fmt.Println("Words surrounded by brackets:", strings.Join(words, " "))

	// 7. Replace "digit X" with just the digit
	digitSentence := "this is a digit 7 in a number"
	digitRe := regexp.MustCompile(`digit (\d)`)
	digitReplaced := digitRe.ReplaceAllString(digitSentence, "$1")
	fmt.Println("Replaced 'digit X':", digitReplaced)
}

