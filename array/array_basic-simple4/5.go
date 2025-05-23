package main

import "fmt"

func main() {
	numbers := [...]int{100, 12, 3, 53}
	fmt.Println(numbers)
	numbers[2] = 15
	fmt.Println(numbers)
}
