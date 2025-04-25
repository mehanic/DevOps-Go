package main

import "fmt"


func saySomething(s string) (string, string) {
	return s, "World"
}

func itIsNotEnd(p string) (string, string) {
	return p, "wait me"
}

func main() {
	// Using saySomething
	greeting, context := saySomething("Hello")
	fmt.Println("saySomething output:", greeting, context)

	// Using itIsNotEnd
	part1, part2 := itIsNotEnd("This is not the end")
	fmt.Println("itIsNotEnd output:", part1, part2)
	fmt.Println("---------------------------------")
	// Combining both functions
	firstMessage, secondMessage := saySomething("Start")
	finalMessage, additionalMessage := itIsNotEnd(secondMessage)

	fmt.Println("Combined function output:")
	fmt.Println("First message:", firstMessage)
	fmt.Println("Second message:", secondMessage)
	fmt.Println("Final message:", finalMessage)
	fmt.Println("Additional message:", additionalMessage)
}
