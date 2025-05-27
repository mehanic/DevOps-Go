package main

import (
	"fmt"
)

func substringAnagrams(s string, t string) int {
	lenS, lenT := len(s), len(t)
	if lenT > lenS {
		return 0
	}

	count := 0
	expectedFreqs := make([]int, 26)
	windowFreqs := make([]int, 26)

	for _, c := range t {
		expectedFreqs[c-'a']++
	}

	left, right := 0, 0
	for right < lenS {
		windowFreqs[s[right]-'a']++

		if right-left+1 == lenT {
			if equalFreqs(expectedFreqs, windowFreqs) {
				count++
			}
			windowFreqs[s[left]-'a']--
			left++
		}
		right++
	}

	return count
}

func equalFreqs(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "cbaebabacd"
	t := "abc"
	fmt.Println("Count of anagram substrings:", substringAnagrams(s, t)) // Output: 2
}
