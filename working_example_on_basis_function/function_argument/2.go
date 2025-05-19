package main

import "fmt"

func printData(word string) {
	fmt.Println(word)
}

func getSum(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	printData("Hello Mike")
	getSum(3, 4)
}

// Hello yerassyl
// 7
