package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Get the basename of the current executable without the extension
func getExecutableName() string {
	execPath, err := os.Executable()
	if err != nil {
		return "program" // Fallback name in case of an error
	}
	return filepath.Base(execPath)
}

// error function to print an error message and exit
func errorMsg(msg string) {
	execName := getExecutableName()
	fmt.Fprintf(os.Stderr, "%s: %s\n", execName, msg)
	os.Exit(1)
}

func main() {
	// Example usage of the error function
	errorMsg("An error occurred while processing the request")
}

