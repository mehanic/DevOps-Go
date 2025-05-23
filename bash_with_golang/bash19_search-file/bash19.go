package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Function to search for a term in the contents of files
func searchInFiles(files []string, searchTerm string) []string {
	var results []string

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		if strings.Contains(string(data), searchTerm) {
			results = append(results, file)
		}
	}

	return results
}

// Function to search for a term in a specific directory recursively
func recursiveSearch(directory, searchTerm string) []string {
	var results []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if strings.Contains(string(data), searchTerm) {
				results = append(results, path)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %v: %v\n", directory, err)
	}

	return results
}

func main() {
	directory := "www.packtpub.com/"
	searchTerm := "Packt"

	// Can we use grep? (simulate)
	result1 := searchInFiles([]string{"~/*"}, searchTerm) // Note: Adjust the path for practical use

	if len(result1) > 0 {
		err := ioutil.WriteFile("result1.txt", []byte(strings.Join(result1, "\n")), 0644)
		if err != nil {
			log.Fatalf("Error writing to result1.txt: %v\n", err)
		}
	}

	// Recursive check
	result2 := recursiveSearch(directory, searchTerm)
	err := ioutil.WriteFile("result2.txt", []byte(strings.Join(result2, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to result2.txt: %v\n", err)
	}

	// Check for multiple terms
	result3 := recursiveSearch(directory, "Publishing")
	result3 = append(result2, result3...) // Combine results
	err = ioutil.WriteFile("result3.txt", []byte(strings.Join(result3, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to result3.txt: %v\n", err)
	}

	// Using find with xargs
	result4 := searchInFiles(result2, searchTerm)
	err = ioutil.WriteFile("result4.txt", []byte(strings.Join(result4, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to result4.txt: %v\n", err)
	}

	// Looking for a specific type of content
	var xmlFiles []string
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".xml" && !strings.Contains(path, ".css") {
			xmlFiles = append(xmlFiles, path)
		}
		return nil
	})

	result5 := searchInFiles(xmlFiles, searchTerm)
	err = ioutil.WriteFile("result5.txt", []byte(strings.Join(result5, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to result5.txt: %v\n", err)
	}

	// Can this also be achieved with wildcards and subshell?
	result6 := searchInFiles([]string{"*.html", "*.txt"}, searchTerm)
	err = ioutil.WriteFile("result6.txt", []byte(strings.Join(result6, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to result6.txt: %v\n", err)
	}

	if len(result6) > 0 {
		fmt.Println("We found results!")
	} else {
		fmt.Println("It broke - it shouldn't happen (Packt is everywhere)!")
	}

	// History command simulation (last part)
	cmd := exec.Command("bash", "-c", "history | grep ls")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error executing command: %v\n", err)
	}
	fmt.Println(string(output))

	// Lesson
	fmt.Println("We can do a lot with grep!")
}

