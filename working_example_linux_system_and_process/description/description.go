package main

import (
	"fmt"
	"os"
)

func printContent() {
	content, err := os.ReadFile("descriptions/description-01.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))
}

func writeNewContent() {
	err := os.WriteFile("descriptions/description-01.txt", []byte("Yodel grew up in a family of singers and loud talkers and could never get a word in edgewise"), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func main() {
	writeNewContent()
	printContent()
}

//Yodel grew up in a family of singers and loud talkers and could never get a word in edgewise