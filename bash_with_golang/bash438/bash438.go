package main

import (
	"fmt"
	"io/ioutil"
//	"os"
//	"path/filepath"
//	"strings"
	"strconv"
)

var (
	title   = "Select file menu"
	prompt  = "Pick a task:"
	options = []string{"list", "delete", "modify", "create"}
)

func mainMenu() {
	fmt.Println(title)
	for {
		fmt.Println(prompt)
		for i, opt := range options {
			fmt.Printf("%d) %s\n", i+1, opt)
		}
		fmt.Println("5) quit")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > len(options)+1 {
			fmt.Println("Invalid option. Try another one.")
			continue
		}

		switch choice {
		case 1:
			listFiles()
		case 2:
			fmt.Println("Delete functionality not implemented yet.")
		case 3:
			fmt.Println("Modify functionality not implemented yet.")
		case 4:
			fmt.Println("Create functionality not implemented yet.")
		case 5:
			fmt.Println("Exiting!")
			return
		}
	}
}

func listFiles() {
	fmt.Println("Choose a file from the list or type 'back' to go back to the main menu:")
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	// Display file options
	for i, fileName := range fileNames {
		fmt.Printf("%d) %s\n", i+1, fileName)
	}
	fmt.Println("back) Go back to main menu")

	var selection string
	for {
		fmt.Scan(&selection)
		if selection == "back" {
			return
		}

		// Check if the selection is a valid file
		if idx, err := strconv.Atoi(selection); err == nil && idx > 0 && idx <= len(fileNames) {
			selectedFile := fileNames[idx-1]
			fmt.Printf("%s was selected\n", selectedFile)
			return
		}

		fmt.Println("Invalid option. Try again or type 'back' to return.")
	}
}

func main() {
	mainMenu()
}

