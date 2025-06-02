package main

import ( 
	      "fmt"
        "unicode"
			)

func Palindrom(s string) bool{
	var filtered[] rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right]{
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	// Test cases
	examples := []string{
		"Was it a car or a cat I saw?", // true
		"tab a cat",                   // false
		"A man, a plan, a canal: Panama!", // true
		"RaceCar",                     // true
		"hello",                        // false
	}

	// Run test cases
	for _, example := range examples {
		fmt.Printf("Palindrome(%q) = %v\n", example, Palindrom(example))
	}
}
