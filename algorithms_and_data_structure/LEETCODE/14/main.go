package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, ch := range s {
		counts[ch]++
	}

	for _, ch := range t {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "listen"
	s2 := "silent"

	result := isAnagram(s1, s2)
	fmt.Printf("isAnagram(%q, %q) = %v\n", s1, s2, result)

	s3 := "hello"
	s4 := "world"

	result2 := isAnagram(s3, s4)
	fmt.Printf("isAnagram(%q, %q) = %v\n", s3, s4, result2)
}
