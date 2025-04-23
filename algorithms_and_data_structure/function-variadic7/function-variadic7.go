package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings
func joinstr1(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	// pass a slice in variadic function
	elements := []string{"geeks", "FOR", "geeks"}
	fmt.Println(joinstr1(elements...))
}
