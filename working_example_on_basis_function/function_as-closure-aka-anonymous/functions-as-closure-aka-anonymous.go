package main

import (
	"fmt"
	"strings"
)

func greetingByCloser(s string) func() string {
	capitalized := strings.Title(s)
	return func() string {
		returnString := "Learning " + capitalized
		return returnString
	}
}

func main() {
	message := greetingByCloser("golang closer functions")
	fmt.Println(message())
}

// ╰─λ go run functions-as-closure-aka-anonymous.go                                                                                                                     0 (0.001s) < 02:58:00
// Learning Golang Closer Functions
