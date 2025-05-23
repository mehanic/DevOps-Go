package main

import "fmt"

func printHello() {
	fmt.Println("Hello world!!!")
}

func printElements(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello")
	}
}

func getSum(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	fmt.Println("It is main")
	printHello()
	printElements(1)
	getSum(30, 4)
}

// It is main
// Hello world!!!
// Hello
// 34
