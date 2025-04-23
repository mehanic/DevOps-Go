package main

import (
	"fmt"
)
//2. Hash Map
func isAnagram(s string, t string) bool {
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

func main() {
	// Test cases
	fmt.Println(isAnagram("racecar", "carrace")) // true
	fmt.Println(isAnagram("jar", "jam"))         // false
	fmt.Println(isAnagram("listen", "silent"))   // true
	fmt.Println(isAnagram("hello", "world"))     // false
}
