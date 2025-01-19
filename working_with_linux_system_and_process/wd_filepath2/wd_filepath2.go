package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if at least one command-line argument is provided
	if len(os.Args) < 1 {
		fmt.Println("No command-line arguments provided.")
		return
	}

	// sys.argv[0] equivalent in Go is os.Args[0]
	fmt.Println("os.Args[0] =", os.Args[0])

	// Get the directory name of the executable path
	pathname := os.DirFS(os.Args[0])
	fmt.Println("path =", pathname)

	// Get the full absolute path of the directory
	absPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("full path =", absPath)
}

