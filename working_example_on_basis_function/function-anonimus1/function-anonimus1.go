package main

import (
	"fmt"
	"log"

	//"net"
	"net/http"
	//	"net/http/httptest"
	"time"
)

// Basic function definitions

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

// EchoHandler-like functionality using HTTP
type EchoHandler struct {
	ack []byte
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for {
		_, err := w.Write(h.ack)
		if err != nil {
			break
		}
		_, err = w.Write([]byte("Hello World")) // Simulating reading from r.Body
		if err != nil {
			break
		}
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
	logger := log.New(log.Writer(), "test: ", log.LstdFlags)

	// Simulate async operation
	result := add(3, 4)
	outputResult(result, logger)

	// Set up a simple TCP server
	ack := []byte("RECEIVED:")
	http.Handle("/echo", &EchoHandler{ack: ack})
	server := &http.Server{
		Addr:         ":15000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
