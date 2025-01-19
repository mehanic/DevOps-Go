package main

import (
	"fmt"
	"os"
)

func writeContent() {
	file, err := os.Create("num-series.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for x := 0; x < 50; x++ {
		content := fmt.Sprintf("%d\n", x)
		_, err := file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func appendContent() {
	file, err := os.OpenFile("num-series.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for x := 50; x < 100; x++ {
		content := fmt.Sprintf("%d\n", x)
		_, err := file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func main() {
	writeContent()
	appendContent()
}
