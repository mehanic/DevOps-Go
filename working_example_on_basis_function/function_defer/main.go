package main

import "fmt"

func main() {
	runApplication(0)
}

func logging() {
	fmt.Println("Log success")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run Application")
	if value == 0 {
		fmt.Println("Error: division by zero")
		return
	}
	result := 10 / value
	fmt.Println(result)
}
