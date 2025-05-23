package main

import (
	"fmt"
	"strings"
)

func main() {
	// Part 1: Extracting name and file type from "sample.jpg"
	fileJpg := "sample.jpg"
	name := strings.TrimSuffix(fileJpg, ".jpg") // Remove the file extension
	fileType := strings.TrimPrefix(fileJpg[strings.LastIndex(fileJpg, "."):], ".") // Get file type

	fmt.Printf("name: %s\n", name)
	fmt.Printf("filetype: %s\n", fileType)

	// Part 2: String manipulations on "hack.fun.book.txt"
	var str = "hack.fun.book.txt"

	// Remove everything before the first dot
	firstPart := str[strings.Index(str, ".")+1:]

	// Extract everything after the last dot
	lastPart := str[strings.LastIndex(str, ".")+1:]

	// Remove the extension
	withoutExt := str[:strings.LastIndex(str, ".")]

	// Extract everything before the first dot
	beforeFirstDot := str[:strings.Index(str, ".")]

	// Printing the results
	fmt.Println(firstPart)         // Output: fun.book.txt
	fmt.Println(lastPart)          // Output: txt
	fmt.Println(withoutExt)        // Output: hack.fun.book
	fmt.Println(beforeFirstDot)    // Output: hack
}

