package main

import "fmt"

func main() {
	fmt.Println("This is a lesson on functions")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 8, 7, 3)
	fmt.Println(myMessage)
	fmt.Println("Pro result is: ", proResult)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "This is a pro result function"
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Hello from GoLang")
}
