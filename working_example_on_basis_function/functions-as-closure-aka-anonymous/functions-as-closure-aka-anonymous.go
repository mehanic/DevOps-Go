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

// Learning Golang Closer Functions
