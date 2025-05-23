package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// create a temporary directory
	name, err := os.MkdirTemp("", "dir") // in Go version older than 1.17 you can use ioutil.TempDir
	if err != nil {
		log.Fatal(err)
	}

	// remove the temporary directory at the end of the program
	defer os.RemoveAll(name)

	// print path of the directory
	fmt.Println(name)
}

