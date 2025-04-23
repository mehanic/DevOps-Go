package main

import (
	"fmt"
	"sort"
)

// 1. Sorting
// Function to check if two strings are anagrams
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Convert strings to rune slices for sorting
	sRunes, tRunes := []rune(s), []rune(t)
	sort.Slice(sRunes, func(i, j int) bool {
		return sRunes[i] < sRunes[j]
	})
	sort.Slice(tRunes, func(i, j int) bool {
		return tRunes[i] < tRunes[j]
	})

	// Compare sorted slices
	for i := range sRunes {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(isAnagram("listen", "silent"))   // true
	fmt.Println(isAnagram("hello", "world"))     // false
	fmt.Println(isAnagram("rat", "tar"))         // true
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("test", "ttew"))       // false
	fmt.Println("------------")
	main1()
	fmt.Println("------------")
	main2()
}

//2. Hash Map

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		countS[ch]++
		countT[rune(t[i])]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}

func main1() {
	// Test cases
	fmt.Println(isAnagram1("listen", "silent"))   // true
	fmt.Println(isAnagram1("hello", "world"))     // false
	fmt.Println(isAnagram1("rat", "tar"))         // true
	fmt.Println(isAnagram1("anagram", "nagaram")) // true
	fmt.Println(isAnagram1("test", "ttew"))       // false
}

//3. Hash Table (Using Array)

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}

func main2() {
	// Test cases
	fmt.Println(isAnagram2("listen", "silent"))   // true
	fmt.Println(isAnagram2("hello", "world"))     // false
	fmt.Println(isAnagram2("rat", "tar"))         // true
	fmt.Println(isAnagram2("anagram", "nagaram")) // true
	fmt.Println(isAnagram2("test", "ttew"))       // false
}
