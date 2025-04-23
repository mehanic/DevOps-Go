package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Check if a given string is a palindrome
func isPalindrome(s string) bool {
	size := len(s)
	halfSize := size / 2
	for i := 0; i < halfSize; i++ {
		if s[i] != s[size-i-1] {
			return false
		}
	}
	return true
}

// Convert a string to a palindrome by appending characters as needed
func convertToPalindrome(s string) string {
	var action func(string, int) string
	action = func(str string, chars int) string {
		charsToAppend := []rune(str[:chars])
		// Reverse the slice of characters
		for i, j := 0, len(charsToAppend)-1; i < j; i, j = i+1, j-1 {
			charsToAppend[i], charsToAppend[j] = charsToAppend[j], charsToAppend[i]
		}
		newValue := str + string(charsToAppend)
		if !isPalindrome(newValue) {
			newValue = action(str, chars+1)
		}
		return newValue
	}
	return action(s, 0)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("string to convert to palindrome (exit to terminate program): ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			break
		}

		result := convertToPalindrome(userInput)
		fmt.Println(result)
	}
}


// go run /home/mehanic/structure/function/function30/function30.go
// string to convert to palindrome (exit to terminate program): munchen
// munchenehcnum
// string to convert to palindrome (exit to terminate program): paris
// parisirap
// string to convert to palindrome (exit to terminate program): ^Csignal: interrupt

