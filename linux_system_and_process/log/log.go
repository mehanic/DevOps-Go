package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	Log()
	fmt.Println("logging 'handled' errors:")
	FinalDestination()
}

//└> $ go run main.go
//basic logging and modification of logger:
//logger: 2024/04/14 log.go:19: test
//new logger: 2024/04/14 log.go:23: you can also add args(true) and use Fatalln to log and crash
//
//logging 'handled' errors:
//2024/04/14 19:54:46 an error occurred: in passthrougherror: error occurred

// Log uses the setup logger
func Log() {
	// we'll configure the logger to write
	// to a bytes.Buffer
	buf := bytes.Buffer{}

	// second argument is the prefix last argument is about options
	// you combine them with a logical or.
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)
	logger.Println("test")
	logger.SetPrefix("new logger: ")
	logger.Printf("you can also add args(%v) and use Fatalln to log and crash", true)
	fmt.Println(buf.String())
}

//---

// OriginalError returns the error original error
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError calls OriginalError and
// forwards the error along after wrapping.
func PassThroughError() error {
	err := OriginalError()
	// no need to check error
	// since this works with nil
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination deals with the error
// and doesn't forward it
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// we log because an unexpected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
