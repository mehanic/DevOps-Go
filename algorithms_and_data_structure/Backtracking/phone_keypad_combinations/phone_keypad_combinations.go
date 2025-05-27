package main

import (
	"fmt"
)

func phoneKeypadCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	keypadMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	var res []string
	backtrack(0, []byte{}, digits, keypadMap, &res)
	return res
}

func backtrack(i int, currCombination []byte, digits string, keypadMap map[byte]string, res *[]string) {
	// Termination condition: if current combination is complete
	if len(currCombination) == len(digits) {
		*res = append(*res, string(currCombination))
		return
	}

	letters := keypadMap[digits[i]]
	for j := 0; j < len(letters); j++ {
		currCombination = append(currCombination, letters[j])
		backtrack(i+1, currCombination, digits, keypadMap, res)
		// Backtrack
		currCombination = currCombination[:len(currCombination)-1]
	}
}

func main() {
	digits := "23"
	combinations := phoneKeypadCombinations(digits)
	fmt.Println("Combinations:", combinations)
}
