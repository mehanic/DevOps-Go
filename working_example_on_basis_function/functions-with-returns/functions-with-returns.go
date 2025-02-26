package main

import (
	"fmt"
	"strings"
)

func repeatWord(s string, times int) (string, int, error) {
	if len(s) <= 0 {
		return "", len(s), fmt.Errorf("length of string can't be less than or equal to zero")
	}
	repeatedWord := strings.Repeat(s, times)
	return repeatedWord, len(s), nil
}

func main() {
	var s = "Golang"
	if newWord, length, err := repeatWord(s, 5); err == nil {
		fmt.Printf("Repeated word [%s], length is [%d]\n", newWord, length)
	} else {
		fmt.Println(err)
	}
}

// Repeated word [GolangGolangGolangGolangGolang], length is [6]
