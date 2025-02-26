package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := []string{
		"pngs/cups-jpg.png",
		"pngs/forest-jpg.png",
		"pngs/golden.png",
		"pngs/work.png",
		"pngs/shakespeare-text.png",
		"pngs/empty.png",
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide a format (png or jpg).")
		return
	}

	list(args[0], files)
}

func list(format string, files []string) {
	valids, err := Detect(format, files)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Correct Files:\n")
	for _, valid := range valids {
		fmt.Printf(" + %s\n", valid)
	}
}

// Detect filters the files that match the provided format (either png or jpg)
func Detect(format string, files []string) ([]string, error) {
	if format != "png" && format != "jpg" {
		return nil, fmt.Errorf("unsupported format: %s. Please use 'png' or 'jpg'", format)
	}

	var validFiles []string
	for _, file := range files {
		// Check if the file ends with the given format
		if strings.HasSuffix(file, "."+format) {
			validFiles = append(validFiles, file)
		}
	}

	return validFiles, nil
}
