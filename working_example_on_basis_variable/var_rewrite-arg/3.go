package main

import "fmt"

func main() {
	var a int // 0 по default
	fmt.Println(a)
	a = 3
	fmt.Println(a)
	//замена значения или перезаписывание переменной
	a = 5
	fmt.Println(a)
}

// 0
// 3
// 5
