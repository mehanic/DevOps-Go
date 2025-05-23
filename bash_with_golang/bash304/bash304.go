package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Directory path
	dirPath := "/var/"

	// Read the directory contents
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err) // Log any errors
	}

	// Loop through the directory contents and print file names
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

