package main

import (
	"fmt"
	"os"
)

// The program checks if a directory argument is provided. If not, it prints a message and exits.
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}
	//The program reads the contents of the specified directory.
	//If there's an error (e.g., the directory doesn't exist), it prints the error and exits.
	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	//The program iterates over the files in the directory.
	//It skips directories.
	//It gets the file info and checks if the file size is 0 bytes.
	//If the file size is 0 bytes, it appends the file name to the names slice.
	var names []byte

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		info, err := file.Info()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if info.Size() == 0 {
			name := file.Name()

			fmt.Println(cap(names))
			names = append(names, name...)
			names = append(names, '\n')
		}
	}
	//The program writes the names of the empty files to out.txt. If there's an error, it prints the error and exits.
	err = os.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", names)
}

// go run main.go files
// 0
// 16
// 32
// empty1.txt
// empty2.txt
// empty3.txt


// Here's what each part means:

// Numbers (0, 16, 32):

// These numbers are the capacity of the names slice at different points in time. The fmt.Println(cap(names)) statements inside the loop print the current capacity of the names slice before each append operation.
// The names slice starts with a capacity of 0, then grows as needed to accommodate new data. The capacities 16 and 32 indicate that the slice's capacity is being increased to handle more data.
// File Names (empty1.txt, empty2.txt, empty3.txt):

// These are the names of the empty files (files with size 0 bytes) found in the specified directory (files).
// The program checks each file in the directory, and if the file size is 0, it appends the file name to the names slice and prints it.