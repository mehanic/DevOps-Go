package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// err prints the error message with context information.
func err(message string) {
	// Get the caller information
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	// Get the function name
	funcName := "unknown function"
	pc, _, _, _ := runtime.Caller(1)
	if fn := runtime.FuncForPC(pc); fn != nil {
		funcName = filepath.Base(fn.Name())
	}

	// Print the error message with context
	fmt.Println()
	fmt.Println("**********************************************************")
	fmt.Println()
	fmt.Printf("error: Line %d in function %s ", line, funcName)
	fmt.Printf("which is in the file %s\n", filepath.Base(file))
	fmt.Printf("error: Message was: %s\n", message)
	fmt.Println()
	fmt.Println("**********************************************************")
	fmt.Println()
}

func main() {
	// Example usage of the err function
	err("This is a test error message.")
}

