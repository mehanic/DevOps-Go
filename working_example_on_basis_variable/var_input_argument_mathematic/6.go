package main

import "fmt"

func main() {
	var znak string
	var a int
	var b int
	fmt.Print("Введите знак:")
	fmt.Scan(&znak)
	fmt.Print("Введите цифру А:")
	fmt.Scan(&a)
	fmt.Print("Введите цифру Б:")
	fmt.Scan(&b)
	if znak == "+" {
		fmt.Println(a + b)
	} else if znak == "-" {
		fmt.Println(a - b)
	} else if znak == "/" {
		fmt.Println(a / b)
	} else if znak == "*" {
		fmt.Println(a * b)
	} else {
		fmt.Println("неправильный знак")
	}
	fmt.Println("выход из программы")
}

// Введите знак:+
// Введите цифру А:56
// Введите цифру Б:34
// 90
// выход из программы
