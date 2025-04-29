package main

import (
	"bufio"
	"fmt"
	"os"
)

func printFileContent() {
	file, err := os.Open("descriptions/description-01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content, _ := reader.Peek(10)
	fmt.Println(string(content))
}

func printFileContentReadlines() {
	file, err := os.Open("descriptions/description-01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) > 1 {
		fmt.Println(lines[1])
	}
}

func printFileContentOneLineAtTime() {
	file, err := os.Open("descriptions/description-01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	printFileContentOneLineAtTime()
}

// go run description1.go                                                                                                                                        ──(Fri,Jan17)─┘
// Mingle excels at doing twice the work in half the time, with pinpoint accuracy.
// These skills serve her well in her role as Senior Data Analyst for an international cloud computing company.
// She's also got a penchant for ballroom dance, line dancing, and pretty much any kind of activity that lets her groove to music.
