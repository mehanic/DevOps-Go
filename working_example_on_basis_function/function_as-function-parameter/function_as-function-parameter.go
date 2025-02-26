package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Golang"
	var (
		repeatedWord, revertedWord string
		err                        error
	)
	if repeatedWord, err = applyFunction(repeatOnce, s); err == nil {
		fmt.Printf("Repeated the word [%s] as [%s]\n", s, repeatedWord)
	} else {
		fmt.Println(err)
	}
	if revertedWord, err = applyFunction(revertOnce, repeatedWord); err == nil {
		fmt.Printf("Reverted the word [%s] as [%s]\n", repeatedWord, revertedWord)
	} else {
		fmt.Println(err)
	}
}

func repeatOnce(s string) (string, error) {
	if len(s) <= 0 {
		return "", fmt.Errorf("Length of string can't less than equal to zero")
	}
	repeatedWord := s + " " + s
	return repeatedWord, nil
}

func revertOnce(s string) (string, error) {
	if len(s) <= 0 {
		return "", fmt.Errorf("Length of string can't less than equal to zero")
	}
	revertedWord := strings.Split(s, " ")
	return revertedWord[len(revertedWord)-1], nil
}

func applyFunction(function func(string) (string, error), s string) (string, error) {
	result, err := function(s)
	if err != nil {
		return "", err
	}
	return result, nil
}

// ╰─λ go run functions-as-function-parameter.go                                                                                                                        0 (0.001s) < 02:59:06
// Repeated the word [Golang] as [Golang Golang]
// Reverted the word [Golang Golang] as [Golang]
