package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("dauren")
	defer fmt.Println("shalabayev")
	fmt.Println("Hello")
}

//после окончания работы метода, можно очищать ресурсы, принцип стека

// Hello
// shalabayev
// dauren
// world
