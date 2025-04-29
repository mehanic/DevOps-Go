package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var sfile, dfile string

	fmt.Print("Enter your source file: ")
	fmt.Scanln(&sfile)

	fmt.Print("Enter your destination file: ")
	fmt.Scanln(&dfile)

	// Read the content from the source file
	content, err := ioutil.ReadFile(sfile)
	if err != nil {
		return
	}
	
	if err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}

	// Write the content to the destination file
	if err = ioutil.WriteFile(dfile, content, 0644); err != nil {
		return
	}

	if err != nil {
		fmt.Println("Error writing to destination file:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
