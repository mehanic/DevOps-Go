package main

import (
	"fmt"
	"log"
	//"net"
	//	"net/http/httptest"
)

// partial Spam
func spam(a, b, c, d int) {
	fmt.Println(a, b, c, d)
}

func outputResult(result int, logger *log.Logger) {
	if logger != nil {
		logger.Printf("Got: %d", result)
	}
}

func add(x, y int) int {
	return x + y
}

// Partial function application equivalent
func partialSpam(a int) func(b, c, d int) {
	return func(b, c, d int) {
		spam(a, b, c, d)
	}
}

func main() {
	// Partial functions
	s1 := partialSpam(1)
	s1(2, 3, 4)
	s1(4, 5, 6)

	s2 := partialSpam(0)
	s2(1, 2, 42) // Using default value for the first argument

	s3 := partialSpam(1)
	s3(2, 0, 42) // Set a default value for b

	// Logging and async-like behavior in Go
	logger := log.New(log.Writer(), "test variant: ", log.LstdFlags)

	// Simulate async operatio
	result := add(3, 4)
	outputResult(result, logger)
}
