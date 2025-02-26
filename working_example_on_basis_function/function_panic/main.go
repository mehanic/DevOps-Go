package main

import "fmt"

func main() {
	runApplication(true)
}

func lastMessage() {
	fmt.Println("Program Ended")
}

func runApplication(error bool) {
	defer lastMessage()
	// error == true
	if error {
		panic("An error occurred in the program")
	}
	fmt.Println("Application is running smoothly")
}
